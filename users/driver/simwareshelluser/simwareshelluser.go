package simwareshelluser

import (
	"pilot/users"
	"github.com/Sirupsen/logrus"
)

const (
	drivername = "simwareshelluser"
)

type Driver struct {

}

func Init() (users.UserManagerDriver, error) {
	return &Driver{}, nil
}

func (d *Driver) String() string {
	return drivername
}

func (d *Driver) AddUser(name string) error {
	logrus.Debugf("Add user :%s ", name)
	return nil
}

func (d *Driver) DelUser(name string) error {
	return nil
}

func (d *Driver) ListUser() ([]string, error) {
	return []string{"za"}, nil
}
