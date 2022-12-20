package app

import (
	"os"
)

type directory struct {
	Root   string
	Config string
	Data   string
	Assets string
}

var Dir = &directory{}

func (d *directory) SetRoot(root string) {
	d.Root = root
	d.Config = root + "config/"
	d.Data = root + "data/"
	d.Assets = root + "assets/"
}

func init() {

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

func Env() {

	for _, item := range Yaml.Access {
		if len(item) != 2 {
			continue
		}
		os.Setenv("AK_"+item[0], item[1])
	}
}
