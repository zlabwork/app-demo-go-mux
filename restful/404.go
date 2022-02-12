package restful

import (
	"app"
	"app/msg"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	app.ResponseCode(w, msg.ErrNotFound)
}
