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

func (*driver) AddConnection(userName, devName, devPort, PeerName, PeerPort string) error {
	cmd := exec.Command(cloudwarecmd,  userName, "connect", devName, devPort, "to", PeerName, PeerPort)
	cmd.Stdin = strings.NewReader(userName)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run();
	if err != nil {
		logrus.Errorf("cloudware user:%s Add connection:%s-%s to %s-%s failed:%v", userName,
			devName, devPort, PeerName, PeerPort, err)
		return err
	}
	logrus.Debugf("cloudware add connection ret:%q", out.String())
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

func (*driver) ListDevices(userName string) (*DeviceList, error) {
	cmd := exec.Command(cloudwarecmd, userName, "list", "device")
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
	if nums < 2 {
		return nil, nil
	}

	allDevices := []DeviceItem{}
	for i, line := range lines {
		if i >= 1 && i < (nums - 1) {
			seps := strings.Fields(line)
			device := DeviceItem{
				DeviceName: seps[0],
				Type: seps[1],
				CSC: seps[2],
			}
			allDevices = append(allDevices, device)
		}
	}
	logrus.Debugf("cloudware ListDevices ret:%v", allDevices)
	return &DeviceList{Items: allDevices}, nil
}

func (*driver) ListConnections(userName, devName string) (*ConnectionInfoList, error) {
	cmd := exec.Command(cloudwarecmd, userName, "list", "connection")
	logrus.Debugf("cmd: %v", cmd)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run(); if err != nil {
		logrus.Errorf("cloudware ListConnections user:%s failed:%v", userName, err)
		return nil, err
	}
	outputs := out.String()
	lines := strings.Split(outputs, "\n")
	nums := len(lines)
	if nums < 3 {
		return nil, nil
	}

	allConnections := []ConnectionInfo{}
	for i, line := range lines {
		if i >= 2 && i < (nums - 1) {
			seps := strings.Fields(line)
			if devName == "" || devName == seps[0] || devName == seps[2] {
				connection := ConnectionInfo{
					DeviceName: seps[0],
					PortName: seps[1],
					PeerDevice: seps[2],
					PeerPort: seps[3],
				}
				allConnections = append(allConnections, connection)
			}
		}
	}
	logrus.Debugf("cloudware ListConnections ret:%v", allConnections)
	return &ConnectionInfoList{Items: allConnections}, nil
}

func (*driver) StopContainer(userName, boardName string) error {
	return nil

}

func (*driver) RemoveContainer(userName, boardName string) error {
	cmd := exec.Command(cloudwarecmd, userName, "delete", boardName)
	logrus.Debugf("cmd: %v", cmd)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run(); if err != nil {
		logrus.Errorf("cloudware remove container:%s failed:%v", boardName, err)
		return err
	}
	logrus.Debugf("cloudware remove container ret:%q", out.String())
	return nil
}

func (*driver) RemoveConnection(userName, devName, portName string) error {
	cmd := exec.Command(cloudwarecmd, userName, "disconnect", devName, portName)
	logrus.Debugf("remove connection cmd:%v", cmd)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run(); if err != nil {
		logrus.Errorf("cloudware remove connection:%s--%s failed:%v", devName, portName, err)
		return err
	}
	logrus.Debugf("cloudware remove connection ret:%q", out.String())
	return nil
}
