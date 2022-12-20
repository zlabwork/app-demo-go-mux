package restful

import (
	"app"
	"app/msg"
	"app/response"
	"net/http"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {

	response.Raw(w, response.JsonData{
		Code:    msg.OK,
		Message: msg.Text(msg.OK),
		Data:    app.VersionName + "@" + app.VersionNumber,
	})
}
