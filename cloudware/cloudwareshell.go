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
		return err
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
	cmd := exec.Command(cloudwarecmd, userName, "create", boardName, btype, bchassis + "," + bslot + "," + bcpu)
	logrus.Debugf("cmd: %v", cmd)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run(); if err != nil {
		logrus.Errorf("cloudware start container:%s failed:%v", boardName, err)
		return err
	}
	logrus.Debugf("cloudware start container ret:%q", out.String())
	return nil
}

func (*driver) ListContainers(userName string) (*ContainerList, error) {
	cmd := exec.Command(cloudwarecmd, userName, "list", "pod")
	logrus.Debugf("cmd: %v", cmd)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run(); if err != nil {
		logrus.Errorf("cloudware ListContainers user:%s failed:%v", userName, err)
		return nil, err
	}
	outputs := out.String()
	lines := strings.Split(outputs, "\n")
	nums := len(lines)
	if nums < 3 {
		return nil, nil
	}

	allContainers := []ContainerItem{}
	for i, line := range lines {
		if i >= 2 && i < (nums - 1) {
			seps := strings.Fields(line)
			container := ContainerItem{
				BoardName: seps[0],
				Status: seps[1],
				RunNode: seps[2],
			}
			allContainers = append(allContainers, container)
		}
	}
	logrus.Debugf("cloudware ListContainers ret:%v", allContainers)
	return &ContainerList{Items: allContainers}, nil
}

func (*driver) StopContainer(userName, boardName string) error {
	return nil

}

func (*driver) RemoveContainer(userName, boardName string) error {
	return nil
}
