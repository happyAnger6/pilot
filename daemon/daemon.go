package daemon

import (
	"sync"

	"pilot/deploy/driver"
	"pilot/deploy/driver/k8s"
)

type Daemon struct {
	mux sync.Mutex
	driver.Driver
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
	driver, err := k8s.Init(); if err != nil {
		return nil, err
	}
	return &Daemon{mux: sync.Mutex{}, Driver:driver}, nil;
}
