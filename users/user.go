package users

type UserManagerDriver interface{
	String() string
	AddUser(name string) error
	DelUser(name string) error
	ListUser() ([]string, error)
}
