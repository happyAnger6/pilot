package simwareshelluser

import "pilot/users"

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
	return nil
}

func (d *Driver) DelUser(name string) error {
	return nil
}

func (d *Driver) ListUser() ([]string, error) {
	return []string{"za"}, nil
}
