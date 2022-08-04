package mysql

import (
	"app"
	"context"
	"database/sql"
)

type UserRepository struct {
	Conn *sql.DB
}

func (ur *UserRepository) GetOne(ctx context.Context, id int64) (*app.User, error) {
	// TODO:
	return nil, nil
}

func (ur *UserRepository) GetMany(ctx context.Context, id []int64) ([]*app.User, error) {
	// TODO:
	return nil, nil
}

func (ur *UserRepository) Delete(ctx context.Context, id int64) error {
	// TODO:
	return nil
}

func (ur *UserRepository) Create(ctx context.Context, u *app.User) error {
	// TODO:
	return nil
}

func (ur *UserRepository) Modify(ctx context.Context, u *app.User) error {
	// TODO:
	return nil
}
