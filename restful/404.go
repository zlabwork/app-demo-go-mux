package restful

import (
	"app/msg"
	"app/response"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {

	response.Code(r.Context(), w, msg.ErrNotFound)
}
