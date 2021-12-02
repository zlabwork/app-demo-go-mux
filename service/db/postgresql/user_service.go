package postgresql

import (
	"app"
	"database/sql"
)

type UserService struct {
	Conn *sql.DB
}

func (us *UserService) Get(id int64) (*app.User, error) {
	return nil, nil
}

func (us *UserService) Delete(id int64) error {
	return nil
}

func (us *UserService) Create(user *app.User) error {
	return nil
}

func (us *UserService) Modify(user *app.User) error {
	return nil
}
