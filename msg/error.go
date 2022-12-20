package msg

const (
	OK          = 200
	Err         = 400
	ErrNotFound = 404
)

const (
	ErrDefault = iota + 20000
	ErrTimeout
	ErrSignature
	ErrAccess
	ErrEncode
	ErrHeader
	ErrParameter
	ErrProcess
	ErrNoData
	ErrModify
)

var statusText = map[int]string{
	OK:           "success",
	Err:          "error",
	ErrNotFound:  "page not found",
	ErrTimeout:   "error request time",
	ErrSignature: "error request signature",
	ErrAccess:    "error access",
	ErrEncode:    "error encode",
	ErrHeader:    "error request headers",
	ErrParameter: "error parameter",
	ErrProcess:   "error in execute process",
	ErrNoData:    "can not find",
	ErrModify:    "error when modify data",
}

func Text(code int) string {
	return statusText[code]
}
