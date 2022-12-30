package postgres

import (
	"app"
	"context"
	"database/sql"
)

func NewUserRepository() (*UserRepository, error) {

	h, err := getHandle()
	if err != nil {
		return nil, err
	}
	return &UserRepository{Conn: h.Conn}, nil
}

type UserRepository struct {
	Conn *sql.DB
}

func (ur *UserRepository) GetOne(ctx context.Context, id int64) (*app.User, error) {
	// TODO:
	return nil, nil
}
