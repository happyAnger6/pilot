package driver

type ContainerOpts struct{
	CreateOpts map[string]interface{}
}

type ContainerList struct {
	Items []interface{}
}

type Driver interface {
	String() string

	StartContainer(name string, opts *ContainerOpts) error

	ListContainers() (*ContainerList, error)

	StopContainer(name string) error

	RemoveContainer(name string) error
}
