package main

import (
    "app/service/postgres"
    "fmt"
)

func main() {
    s := postgres.UserService{}
    fmt.Println(s.Users())
}
