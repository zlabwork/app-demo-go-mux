package service

import (
	"app"
	"app/service/repository/postgres"
)

// UserRepository Interfaces that the repository layer needs to implement
type UserRepository interface {
	Get(alias string) (*app.User, error)
	GetOne(id int64) (*app.User, error)
	GetMany(id []int64) ([]*app.User, error)
	Delete(id int64) error
	Create(u *app.User) error
	Modify(u *app.User) error
}

type UserService struct {
	Repo UserRepository
}

func (us *UserService) Get(alias string) (*app.User, error) {
	return us.Repo.Get(alias)
}

func (us *UserService) GetOne(id int64) (*app.User, error) {
	return us.Repo.GetOne(id)
}

func (us *UserService) GetMany(id []int64) ([]*app.User, error) {
	return us.Repo.GetMany(id)
}

func (us *UserService) Delete(id int64) error {
	return us.Repo.Delete(id)
}

func (us *UserService) Create(u *app.User) error {
	return us.Repo.Create(u)
}

func (us *UserService) Modify(u *app.User) error {
	return us.Repo.Modify(u)
}

func NewUserService() (*UserService, error) {
	// TODO:
	return &UserService{Repo: &postgres.UserRepository{}}, nil
}
