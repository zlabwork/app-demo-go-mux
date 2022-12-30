package restful

import (
	"app"
	"app/msg"
	"app/response"
	"net/http"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {

	response.Data(r.Context(), w, msg.OK, app.VersionName+"@"+app.VersionNumber)
}
