package app

import (
	"context"
	"time"
)

type User struct {
	Id        int64
	Alias     string
	Name      string
	Gender    int
	Birth     time.Time
	State     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// UserRepository Interfaces that the repository layer needs to implement
type UserRepository interface {
	GetOne(ctx context.Context, id int64) (*User, error)
	GetMany(ctx context.Context, id []int64) ([]*User, error)
	Delete(ctx context.Context, id int64) error
	Create(ctx context.Context, u *User) error
	Modify(ctx context.Context, u *User) error
}
