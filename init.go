package app

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

var Dir = &directory{}

type directory struct {
	Root   string
	Config string
	Data   string
	Assets string
}

func (d *directory) SetRoot(root string) {
	d.Root = root
	d.Config = root + "configs/"
	d.Data = root + "data/"
	d.Assets = root + "assets/"
}

func init() {

	// path
	setPath()

	// rand seed
	rand.Seed(time.Now().UnixNano())

	// env
	if len(os.Getenv("APP_ENV")) == 0 {
		log.Fatal("env is missing, try to execute 'export `cat .env`' before run the app")
	}

	// app.yaml
	bs, err := ioutil.ReadFile(Dir.Config + "app.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if yaml.Unmarshal(bs, Yaml) != nil {
		log.Fatal(err)
	}

	// logs
	f, err := os.OpenFile(Dir.Data+"system.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(f)
	}
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{})

	// env & libs
	setEnv()
	Libs = NewLibs()
}

func setPath() {

	// TODO:
	// Root Path - relative to the main.go file
	// make sure it execute before other init()
	_, err := os.Stat("./go.mod")
	if err != nil {
		Dir.SetRoot("../")
	} else {
		Dir.SetRoot("./")
	}
}

func setEnv() {

	for _, item := range Yaml.Access {
		if len(item) != 2 {
			continue
		}
		os.Setenv("AK_"+item[0], item[1])
	}
}
