package cloudware

import "os/exec"

const (
	drivername="cloudwareshell"
)

type driver struct {

}


func Init()(Driver, error) {
	return &driver{}, nil
}

func (*driver) String() string {
	return drivername
}

func (*driver) AddUser(userName string) error {
	exec.Command("cloudware", "init")
	return nil
}

func (*driver) DelUser(userName string) error {
	return nil
}

func (*driver) ListUser() (*UserList, error) {
	return nil, nil
}

func (*driver) StartContainer(userName, boardName, btype, bchassis, bslot, bcpu string) error {
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
