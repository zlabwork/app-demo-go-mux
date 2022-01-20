package main

import (
	"app"
	"app/service"
	"context"
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
	err = yaml.Unmarshal(bs, app.Yaml)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// libs
	app.Libs = app.NewLibs()

	// banner
	app.Banner("This is a demo app")

	// read from database
	srv, _ := service.NewUserService(context.TODO())
	user, _ := srv.GetOne(111)
	fmt.Println(user)
}
