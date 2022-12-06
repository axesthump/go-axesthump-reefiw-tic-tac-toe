package storage

import (
	"fmt"
	"log"
)

type InMemoryStorage struct {
	users []User
}

func NewInMemoryStorage() *InMemoryStorage {
	users := make([]User, 2)
	users[0].ID = -1
	users[1].ID = -1
	return &InMemoryStorage{
		users: users,
	}
}

func (s *InMemoryStorage) AddUser(id int) {
	switch {
	case s.users[0].ID == -1:
		s.users[0].ID = id
		s.users[0].PlayerType = OPlayer
	case s.users[1].ID == -1:
		s.users[1].ID = id
		s.users[1].PlayerType = XPlayer
	default:
		log.Println("All users alredy register")
	}
}

func (s *InMemoryStorage) GetUser(id int) (*User, error) {
	if id == 0 || id == 1 {
		return &s.users[id], nil
	} else {
		return nil, fmt.Errorf("user dont exist")
	}
}
