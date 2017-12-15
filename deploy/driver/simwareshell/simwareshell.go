package simwareshell

import (
	"pilot/deploy/driver"
	"os/exec"
	"github.com/sirupsen/logrus"
)

const (
	drivername = "simwareshell"
	driverapp = "cloudware"
)

type Driver struct {

}

func (d *Driver) String() string {
	return drivername
}

func (d *Driver) StartContainer(name string, opts *driver.ContainerOpts) error {
	bprojName := opts.CreateOpts["bname"].(string)
	btype := opts.CreateOpts["btype"].(string)
	bchassis := opts.CreateOpts["bchassis"].(string)
	bslot := opts.CreateOpts["bslot"].(string)
	bcpu := opts.CreateOpts["bcpu"].(string)
	username := opts.CreateOpts["username"].(string)

	cmd := exec.Command("cloudware", "create", username, bprojName, btype, bchassis+","+bslot+","+bcpu)
	logrus.Debugf("cloudware start :%v", cmd)
	err := cmd.Run()
	if err != nil {
		logrus.Debugf("cloudware start cntainer failed:%v", err)
		return err
	}
	return nil
}

func (d *Driver) ListContainers() (*driver.ContainerList, error) {
	return nil, nil
}

func (d *Driver) StopContainer(name string) error {
	return nil
}

func (d *Driver) RemoveContainer(name string) error {
	cmd := exec.Command("cloudware", "delete", name)
	logrus.Debugf("simwareshell remove:%v", cmd)
	err := cmd.Run()
	if err != nil {
		logrus.Debugf("simwareshell remove cntainer failed:%v", err)
		return err
	}
	return nil
}

func Init() (driver.Driver, error) {
	return &Driver{}, nil
}
