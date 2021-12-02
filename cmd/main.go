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

	// libs
	app.Libs = app.NewLibs()

	// read from database
	us := postgresql.UserService{}
	u, _ := us.Get(1111)
	fmt.Println(u)

	// read from cache
	uc := app.NewUserCache(&us)
	u, _ = uc.Get(2222)
	fmt.Println(u)
}
