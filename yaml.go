package app

var Cfg = Yaml{}

type Yaml struct {
	Db struct {
		Mysql struct {
			Host string
			Port string
			User string
			Pass string
		}
		Redis struct {
			Host string
			Port string
		}
	}
}
