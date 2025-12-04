package types

type Users interface {
	GetUsers() ([]Users, error)
}
