package msg

const (
	OK  = 200
	Err = 400
)

var statusText = map[int]string{
	OK:  "success",
	Err: "error",
}

func Text(code int) string {
	return statusText[code]
}
