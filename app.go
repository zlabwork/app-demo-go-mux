package app

const (
	TraceKey = "_trace_id"
)

var Libs *libs

type libs struct {
}

func NewLibs() *libs {
	return &libs{}
}
