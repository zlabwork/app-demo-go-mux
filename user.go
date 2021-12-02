package app

import "time"

type User struct {
	Id     int64
	Name   string
	Gender int
	State  int
	Ctime  time.Time
	Mtime  time.Time
}

type UserService interface {
	Get(id int64) (*User, error)
	Delete(id int64) error
	Create(user *User) error
	Modify(user *User) error
}
