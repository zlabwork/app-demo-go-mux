package main

import (
    "app"
    "app/service/db/postgresql"
    "fmt"
    "github.com/joho/godotenv"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)

func main() {

    // env
    err := godotenv.Load("../.env")
    if err != nil {
        log.Fatal(err)
    }

    // yaml
    bs, err := ioutil.ReadFile("../config/app.yaml")
    err = yaml.Unmarshal(bs, &app.Cfg)
    if err != nil {
        log.Fatalf("error: %v", err)
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
