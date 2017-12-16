package cloudware


type ContainerOpts struct{
	CreateOpts map[string]interface{}
}

type DeviceItem struct {
	DeviceName string
	Type string
	CSC string
}

type DeviceList struct {
	Items []DeviceItem
}

type ContainerItem struct {
	BoardName string
	BoardType string
	BoardId	  string
	Status	  string
	RunNode	  string
}

type ContainerList struct {
	Items []ContainerItem
}

type UserInfo struct {
	userName string
}

type UserList struct {
	Users []UserInfo
}

type ImageInfo struct {
	Type string
	Name string
}

type ImageList struct {
	Items []ImageInfo
}

type ConnectionInfo struct {
	DeviceName string
	PortName string
	PeerDevice string
	PeerPort string
}

type ConnectionInfoList struct {
	Items []ConnectionInfo
}

type Driver interface {
	String() string

	SetImage(userName, iType, name string) error

	ListImages(userName string) (*ImageList, error)

	AddUser(userName string) error

	AddConnection(userName, devName, devPort, PeerName, PeerPort string) error

	DelUser(userName string) error

	ListUser() (*UserList, error)

	StartContainer(userName, boardName, btype, bchassis, bslot, bcpu string) error

	ListContainers(userName string) (*ContainerList, error)

	ListDevices(userName string) (*DeviceList, error)

	ListConnections(userName, devName string) (*ConnectionInfoList, error)

	StopContainer(userName, boardName string) error

	RemoveContainer(userName, boardName string) error

	RemoveConnection(userName, devName, portName string) error
}