package storage

type PlayerType int

const (
	XPlayer PlayerType = iota
	OPlayer
)

type User struct {
	id         int
	playerType PlayerType
}

type Storage interface {
	AddUser(id int)
	GetUser(id int) (*User, error)
}
