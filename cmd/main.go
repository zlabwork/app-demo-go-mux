package main

import (
    "app"
    "app/service/database/postgresql"
    "fmt"
    "github.com/joho/godotenv"
    "log"
)

func main() {

    // env
    err := godotenv.Load("../.env.local")
    if err != nil {
        err = godotenv.Load("../.env")
        if err != nil {
            log.Fatal(err)
        }
    }

    // read from database
    s := postgresql.UserService{}
    u, _ := s.User(1111)
    fmt.Println(u)

    // read from cache
    cache := app.NewUserCache(&s)
    u, _ = cache.User(2222)
    fmt.Println(u)
}
