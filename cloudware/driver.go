package cloudware


type ContainerOpts struct{
	CreateOpts map[string]interface{}
}

type ContainerItem struct {
	BoardName string
	BoardType string
	BoardId	  string
}

type UserInfo struct {
	userName string
}

type UserList struct {
	Users []UserInfo
}

type ContainerList struct {
	Items []ContainerItem
}

type Driver interface {
	String() string

	AddUser(userName string) error

	DelUser(userName string) error

	ListUser() (*UserList, error)

	StartContainer(userName, boardName, btype, bchassis, bslot, bcpu string) error

	ListContainers(userName string) (*ContainerList, error)

	StopContainer(userName, boardName string) error

	RemoveContainer(userName, boardName string) error
}