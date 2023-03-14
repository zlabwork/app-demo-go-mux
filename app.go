package app

import (
	"github.com/bwmarrin/snowflake"
	"os"
	"strconv"
)

const (
	AuthKey  = "_auth_key"
	TraceKey = "_trace_id"
)

var Libs *libs

type libs struct {
	Snow *snowflake.Node
}

func NewLibs() *libs {

	i, _ := strconv.ParseInt(os.Getenv("APP_NODE"), 10, 64)
	snowflake.Epoch = 1498612200000 // 2017-06-28 09:10:00
	snowflake.NodeBits = 8
	snowflake.StepBits = 14
	sn, _ := snowflake.NewNode(i)
	return &libs{
		Snow: sn,
	}
}
