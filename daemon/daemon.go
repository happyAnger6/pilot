package daemon

import (
	"sync"

	"pilot/deploy/driver"
	_ "pilot/deploy/driver/k8s"
	"pilot/models/deploy/board"
	"path/filepath"
	"pilot/models"
	 "pilot/deploy/driver/stub"
	_ "pilot/deploy/driver/simwareshell"
	"pilot/users"
	"pilot/users/driver/simwareshelluser"
	"github.com/sirupsen/logrus"
)

const (
	storerootdir="/var/pilot/store"
)

type Daemon struct {
	mux sync.Mutex
	driver.Driver
	users.UserManagerDriver
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
	logrus.Debugf("Daemon initizlie")
	driver, err := stub.Init(); if err != nil {
		return nil, err
	}

	userDriver, err := simwareshelluser.Init(); if err != nil {
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

	daemon = &Daemon{mux: sync.Mutex{}, Driver: driver,
				 UserManagerDriver: userDriver, BoardStore: bs}
	return daemon, nil
}
