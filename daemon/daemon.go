package daemon

import (
	"sync"

	"pilot/deploy/driver"
	"pilot/deploy/driver/k8s"
	"pilot/models/deploy/board"
	"path/filepath"
	"pilot/models"
	"pilot/deploy/driver/stub"
)

const (
	storerootdir="/var/pilot/store"
)

type Daemon struct {
	mux sync.Mutex
	driver.Driver
	BoardStore board.BoardStore
}

var daemon *Daemon
var mux = sync.Mutex{}

func GetInstance() (*Daemon, error) {
	mux.Lock()
	defer mux.Unlock()
	if daemon == nil {
		daemon, err := initialize(); if err != nil {
			return daemon, err
		}
	}
	return daemon, nil
}

func initialize()(*Daemon, error) {
	driver, err := stub.Init(); if err != nil {
		return nil, err
	}

	boardRoot := filepath.Join(storerootdir, "board")
	bfs, err := models.NewFSStoreBackend(boardRoot)
	if err != nil {
		return nil, err
	}

	bs, err := board.NewBoardStore(bfs)
	if err != nil {
		return nil, err
	}

	return &Daemon{mux: sync.Mutex{}, Driver: driver,
				BoardStore: bs}, nil;
}
