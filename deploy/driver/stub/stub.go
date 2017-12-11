package stub

import (
	"fmt"

	"pilot/deploy/driver"
)

const (
	driverName = "stub"
)

type Driver struct{
}

func (d *Driver) String() string {
	return driverName
}

func (d *Driver) StartContainer(name string, opts *driver.ContainerOpts) error {
	fmt.Printf("StartContainer: %v\r\n", opts)

	return nil
}

type board struct{
	Name string
	Image string
	Status string
}

func (d *Driver) ListContainers() (*driver.ContainerList, error) {
	lists := &driver.ContainerList{}

	boards := []board{{Name: "sim01-mpu-0-1-0", Image: "v9trunk:d001", Status:"Up"},
		{Name: "sim01-lpu-0-3-0", Image: "v9trunk:d002", Status:"Down"}}

	for _, b := range boards {
		lists.Items = append(lists.Items, b)
	}

	return  lists, nil
}

func (d *Driver) StopContainer(name string) error {
	return nil
}

func (d *Driver) RemoveContainer(name string) error {
	return nil
}

// use the current context in kubeconfig
func Init() (driver.Driver, error){
	d := &Driver{
	}
	return d, nil
}
