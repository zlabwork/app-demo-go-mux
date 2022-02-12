package restful

import (
	"app"
	"app/msg"
	"net/http"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	app.ResponseRaw(w, app.JsonOK{
		Code:    msg.OK,
		Message: msg.Text(msg.OK),
		Data:    app.VersionName + "@" + app.VersionNumber,
	})
}
