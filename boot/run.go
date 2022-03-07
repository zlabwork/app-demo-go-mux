package boot

import (
	"app"
	"os"
)

func setEnv() error {

	for _, item := range app.Yaml.Access {
		if len(item) != 2 {
			continue
		}
		os.Setenv("AK_"+item[0], item[1])
	}
	return nil
}

func Run() {
	setEnv()
}
