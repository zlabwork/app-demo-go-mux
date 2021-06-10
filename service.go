package app

type User struct {
    Id   int64
    Name string
    Age  int
}

type UserService interface {
    User(id int64) (*User, error)
    Users() ([]*User, error)
    CreateUser(u *User) error
    DeleteUser(id int64) error
}
