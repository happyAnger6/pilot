package cloudware

import (
	"os/exec"
	"strings"
	"bytes"
	"github.com/sirupsen/logrus"
)

const (
	drivername="cloudwareshell"
	cloudwarecmd="cloudware"
)

type driver struct {

}

func Init()(Driver, error) {
	return &driver{}, nil
}

func (*driver) String() string {
	return drivername
}

func (*driver) SetImage(userName, iType, name string) error {
	return nil
}

func (*driver) ListImages(userName string) (*ImageList, error) {
	return nil, nil
}

func (*driver) AddUser(userName string) error {
	cmd := exec.Command(cloudwarecmd, "init")
	cmd.Stdin = strings.NewReader(userName)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run(); if err != nil {
		logrus.Errorf("cloudware Add user:%s failed:%v", userName, err)
	}
	logrus.Debugf("cloudware init ret:%q", out.String())
	return nil
}

func (*driver) DelUser(userName string) error {
	return nil
}

func (*driver) ListUser() (*UserList, error) {
	return nil, nil
}

func (*driver) StartContainer(userName, boardName, btype, bchassis, bslot, bcpu string) error {
	cmd := exec.Command(cloudwarecmd, userName, boardName, btype, bchassis + "," + bslot + "," + bcpu)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run(); if err != nil {
		logrus.Errorf("cloudware Add user:%s failed:%v", userName, err)
	}
	logrus.Debugf("cloudware start container ret:%q", out.String())
	return nil
}

func (*driver) ListContainers(userName string) (*ContainerList, error) {
	return nil, nil
}

func (*driver) StopContainer(userName, boardName string) error {
	return nil

}

func (*driver) RemoveContainer(userName, boardName string) error {
	return nil
}
