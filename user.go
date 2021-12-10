package app

import "time"

type User struct {
	Id     int64
	Alias  string
	Name   string
	Gender int
	Birth  time.Time
	State  int
	Ctime  time.Time
	Mtime  time.Time
}
