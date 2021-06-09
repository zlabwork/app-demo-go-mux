package postgres

import (
    "app/service"
    "database/sql"
)

type UserService struct {
    Db *sql.DB
}

func (s *UserService) User(id int64) (*service.User, error) {
    return &service.User{Id: id}, nil
}

func (s *UserService) Users() ([]*service.User, error) {
    return nil, nil
}

func (s *UserService) CreateUser(u *service.User) error {
    return nil
}

func (s *UserService) DeleteUser(id int64) error {
    return nil
}
