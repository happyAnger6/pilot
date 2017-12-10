package models

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/pkg/errors"
)

// DigestWalkFunc is function called by StoreBackend.Walk
type NameWalkFunc func(name string) error

// StoreBackend provides interface for image.Store persistence
type StoreBackend interface {
	Walk(f NameWalkFunc) error
	Get(name string) ([]byte, error)
	Set(name string, data []byte) (error)
	Delete(name string) error
}

// fs implements StoreBackend using the filesystem.
type fs struct {
	sync.RWMutex
	root string
}

const (
	contentDirName  = "content"
	metadataDirName = "metadata"
)

// NewFSStoreBackend returns new filesystem based backend for image.Store
func NewFSStoreBackend(root string) (StoreBackend, error) {
	return newFSStore(root)
}

func newFSStore(root string) (*fs, error) {
	s := &fs{
		root: root,
	}
	if err := os.MkdirAll(filepath.Join(root, contentDirName), 0700); err != nil {
		return nil, errors.Wrap(err, "failed to create storage backend")
	}
	return s, nil
}

func (s *fs) contentFile(name string) string {
	return filepath.Join(s.root, contentDirName, name)
}

func (s *fs) metadataDir(name string) string {
	return filepath.Join(s.root, metadataDirName, name)
}

// Walk calls the supplied callback for each image ID in the storage backend.
func (s *fs) Walk(f NameWalkFunc) error {
	// Only Canonical digest (sha256) is currently supported
	s.RLock()
	dir, err := ioutil.ReadDir(filepath.Join(s.root, contentDirName))
	s.RUnlock()
	if err != nil {
		return err
	}
	for _, v := range dir {
		if err := f(v.Name()); err != nil {
			return err
		}
	}
	return nil
}

// Get returns the content stored under a given digest.
func (s *fs) Get(name string) ([]byte, error) {
	s.RLock()
	defer s.RUnlock()

	return s.get(name)
}

func (s *fs) get(name string) ([]byte, error) {
	content, err := ioutil.ReadFile(s.contentFile(name))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get name %s", name)
	}

	return content, nil
}

// Set stores content
func (s *fs) Set(name string, data []byte) error {
	s.Lock()
	defer s.Unlock()

	if len(data) == 0 {
		return fmt.Errorf("invalid empty data")
	}

	if err := ioutil.WriteFile(s.contentFile(name), data, 0600); err != nil {
		return errors.Wrap(err, "failed to write content data")
	}

	return nil
}

// Delete removes content and metadata files associated with the digest.
func (s *fs) Delete(name string) error {
	s.Lock()
	defer s.Unlock()

	if err := os.Remove(s.contentFile(name)); err != nil {
		return err
	}
	return nil
}
