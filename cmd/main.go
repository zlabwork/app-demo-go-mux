package main

import (
	"app"
	"app/service"
	"context"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
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
	log.Println(user)
}

func init() {
	f, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(f)
	}
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{})
}
