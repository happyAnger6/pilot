package driver

type ContainerOpts struct{
	CreateOpts map[string]string
}

type Driver interface {
	String() string

	StartContainer(name string, opts *ContainerOpts) error

	StopContainer(name string) error

	RemoveContainer(name string) error
}
