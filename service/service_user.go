package service

import (
	"app"
	"app/service/repository/mysql"
	"context"
)

type UserService struct {
	Repo app.UserRepository
}

func NewUserService() (*UserService, error) {
	// TODO:
	return &UserService{Repo: &mysql.UserRepository{}}, nil
}

func (us *UserService) GetOne(ctx context.Context, id int64) (*app.User, error) {

	return us.Repo.GetOne(ctx, id)
}

func (us *UserService) GetMany(ctx context.Context, id []int64) ([]*app.User, error) {

	return us.Repo.GetMany(ctx, id)
}

func (us *UserService) Delete(ctx context.Context, id int64) error {

	return us.Repo.Delete(ctx, id)
}

func (us *UserService) Create(ctx context.Context, u *app.User) error {

	return us.Repo.Create(ctx, u)
}

func (us *UserService) Modify(ctx context.Context, u *app.User) error {

	return us.Repo.Modify(ctx, u)
}
