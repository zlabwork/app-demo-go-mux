package service

import (
	"app"
	"app/service/repository/postgres"
	"context"
)

type UserService struct {
	ctx  context.Context
	Repo app.UserRepository
}

func NewUserService(ctx context.Context) (*UserService, error) {
	// TODO:
	return &UserService{ctx: ctx, Repo: &postgres.UserRepository{}}, nil
}

func (us *UserService) GetOne(id int64) (*app.User, error) {
	return us.Repo.GetOne(us.ctx, id)
}

func (us *UserService) GetMany(id []int64) ([]*app.User, error) {
	return us.Repo.GetMany(us.ctx, id)
}

func (us *UserService) Delete(id int64) error {
	return us.Repo.Delete(us.ctx, id)
}

func (us *UserService) Create(u *app.User) error {
	return us.Repo.Create(us.ctx, u)
}

func (us *UserService) Modify(u *app.User) error {
	return us.Repo.Modify(us.ctx, u)
}
