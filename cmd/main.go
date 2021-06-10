package main

import (
    "app"
    "app/service/postgres"
    "fmt"
)

func main() {
    // read from database
    s := postgres.UserService{}
    u, _ := s.User(1111)
    fmt.Println(u)

    // read from cache
    cache := app.NewUserCache(&s)
    u, _ = cache.User(2222)
    fmt.Println(u)
}
