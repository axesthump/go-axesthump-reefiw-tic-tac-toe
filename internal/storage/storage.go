package storage

type PlayerType int

const (
	XPlayer PlayerType = iota
	OPlayer
)

type User struct {
	ID         int
	PlayerType PlayerType
}

type Storage interface {
	AddUser(id int)
	GetUser(id int) (*User, error)
}
