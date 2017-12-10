package board

import (
	"encoding/json"
	"fmt"
	"sync"

	"pilot/models"
	"github.com/sirupsen/logrus"
)

// Store is an interface for creating and accessing boards
type BoardStore interface {
	Store(name string, board *Board) error
	Get(name string) (*Board, error)
	Delete(name string) error
}

type store struct {
	sync.RWMutex
	boards	  map[string]*Board
	fs        models.StoreBackend
}

// NewBoardStore returns new store object
func NewBoardStore(fs models.StoreBackend) (BoardStore, error) {
	is := &store{
		boards:    make(map[string]*Board),
		fs:        fs,
	}

	// load all current boards
	if err := is.restore(); err != nil {
		return nil, err
	}

	return is, nil
}

func (is *store) restore() error {
	err := is.fs.Walk(func(name string) error {
		board, err := is.Get(name)
		if err != nil {
			logrus.Errorf("invalid board name %v, %v", name, err)
			return nil
		}

		is.boards[name] = board

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (is *store) Store(name string, board *Board) error{
	is.Lock()
	defer is.Unlock()

	if _, exists := is.boards[name]; exists {
		return nil
	}

	data, err := json.Marshal(*board)
	if err != nil {
		return err
	}

	err = is.fs.Set(name, data)
	if err != nil {
		return err
	}
	is.boards[name] = board

	return nil
}

type imageNotFoundError string

func (e imageNotFoundError) Error() string {
	return "No such image: " + string(e)
}

func (imageNotFoundError) NotFound() {}

func (is *store) Get(name string) (*Board, error) {
	config, err := is.fs.Get(name)
	if err != nil {
		return nil, err
	}

	board, err := NewFromJSON(config)
	if err != nil {
		return nil, err
	}

	return board, nil
}

func (is *store) Delete(name string) error {
	is.Lock()
	defer is.Unlock()

	b := is.boards[name]
	if b == nil {
		return fmt.Errorf("unrecognized board name %s", name)
	}

	delete(is.boards, name)
	is.fs.Delete(name)

	return nil
}
