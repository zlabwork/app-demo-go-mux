package postgres

import (
	"app"
	"database/sql"
)

type UserRepository struct {
	Conn *sql.DB
}

func (ur *UserRepository) Get(alias string) (*app.User, error) {
	// TODO:
	return nil, nil
}

func (ur *UserRepository) GetOne(id int64) (*app.User, error) {
	// TODO:
	return nil, nil
}

func (ur *UserRepository) GetMany(id []int64) ([]*app.User, error) {
	// TODO:
	return nil, nil
}

func (ur *UserRepository) Delete(id int64) error {
	// TODO:
	return nil
}

func (ur *UserRepository) Create(u *app.User) error {
	// TODO:
	return nil
}

func (ur *UserRepository) Modify(u *app.User) error {
	// TODO:
	return nil
}
