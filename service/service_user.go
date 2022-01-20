package service

import (
	"app"
	"app/service/repository/postgres"
	"context"
)

// UserRepository Interfaces that the repository layer needs to implement
type UserRepository interface {
	GetOne(ctx context.Context, id int64) (*app.User, error)
	GetMany(ctx context.Context, id []int64) ([]*app.User, error)
	Delete(ctx context.Context, id int64) error
	Create(ctx context.Context, u *app.User) error
	Modify(ctx context.Context, u *app.User) error
}

type UserService struct {
	ctx  context.Context
	Repo UserRepository
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
