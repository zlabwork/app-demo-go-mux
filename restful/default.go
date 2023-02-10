package restful

import (
	"app"
	"app/msg"
	"app/response"
	"net/http"
	"os"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {

	response.Data(r.Context(), w, msg.OK, os.Getenv("APP_NAME")+"@"+app.VersionNumber)
}
