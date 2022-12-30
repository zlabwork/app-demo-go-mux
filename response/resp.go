package response

import (
	"app"
	"app/msg"
	"context"
	"encoding/json"
	"net/http"
)

type JsonData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	TraceId interface{} `json:"trace_id,omitempty"`
}

func Raw(ctx context.Context, w http.ResponseWriter, data *JsonData) {

	w.Header().Set("Content-Type", "application/json")
	data.TraceId = ctx.Value(app.TraceKey)
	bs, _ := json.Marshal(data)
	w.Write(bs)
}

func Code(ctx context.Context, w http.ResponseWriter, code int) {

	Raw(ctx, w, &JsonData{
		Code:    code,
		Message: msg.Text(code),
	})
}

func Message(ctx context.Context, w http.ResponseWriter, code int, message string) {

	Raw(ctx, w, &JsonData{
		Code:    code,
		Message: message,
	})
}

func Data(ctx context.Context, w http.ResponseWriter, code int, data interface{}) {

	Raw(ctx, w, &JsonData{
		Code:    code,
		Message: msg.Text(code),
		Data:    data,
	})
}
