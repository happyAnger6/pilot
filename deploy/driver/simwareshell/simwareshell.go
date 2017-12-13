package simwareshell

import (
	"pilot/deploy/driver"
	"os/exec"
	"github.com/Sirupsen/logrus"
)

const (
	drivername = "simwareshell"
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

	cmd := exec.Command("simware", "create", bprojName, btype, bchassis+","+bslot+","+bcpu)
	logrus.Debugf("simwareshell start :%v", cmd)
	err := cmd.Run()
	if err != nil {
		logrus.Debugf("simwareshell start cntainer failed:%v", err)
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
	cmd := exec.Command("simware", "delete", name)
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
