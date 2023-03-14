package app

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math/rand"
	"os"
	"os/user"
	"strings"
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
	d.Config = root + "config" + string(os.PathSeparator)
	d.Assets = root + "assets" + string(os.PathSeparator)

	dataPath := os.Getenv("APP_DATA")
	if dataPath != "" {
		if strings.HasPrefix(dataPath, string(os.PathSeparator)) {
			d.Data = strings.TrimRight(dataPath, string(os.PathSeparator)) + string(os.PathSeparator)
		} else if strings.HasPrefix(dataPath, ".") {
			d.Data = root + strings.Trim(strings.Trim(dataPath, "."), string(os.PathSeparator)) + string(os.PathSeparator)
		} else {
			d.Data = root + strings.TrimRight(dataPath, string(os.PathSeparator)) + string(os.PathSeparator)
		}
	} else {
		u, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		d.Data = u.HomeDir + string(os.PathSeparator) + os.Getenv("APP_NAME") + string(os.PathSeparator)
	}
	os.MkdirAll(d.Data, 0755)
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
