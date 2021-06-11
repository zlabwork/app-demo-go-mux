package postgresql

import (
    "app"
    "database/sql"
)

type UserService struct {
    Db *sql.DB
}

func (s *UserService) User(id int64) (*app.User, error) {
    return &app.User{Id: id}, nil
}

func (s *UserService) Users() ([]*app.User, error) {
    return nil, nil
}

func (s *UserService) CreateUser(u *app.User) error {
    return nil
}

func (s *UserService) DeleteUser(id int64) error {
    return nil
}
