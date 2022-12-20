package response

import (
	"app/msg"
	"encoding/json"
	"net/http"
)

type JsonCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type JsonData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Raw(w http.ResponseWriter, data interface{}) {

	w.Header().Set("Content-Type", "application/json")
	bs, _ := json.Marshal(data)
	w.Write(bs)
}

func Code(w http.ResponseWriter, code int) {

	w.Header().Set("Content-Type", "application/json")
	bs, _ := json.Marshal(JsonCode{
		Code:    code,
		Message: msg.Text(code),
	})
	w.Write(bs)
}

func Message(w http.ResponseWriter, code int, message string) {

	w.Header().Set("Content-Type", "application/json")
	bs, _ := json.Marshal(JsonCode{
		Code:    code,
		Message: message,
	})
	w.Write(bs)
}

func Data(w http.ResponseWriter, code int, data interface{}) {

	w.Header().Set("Content-Type", "application/json")
	bs, _ := json.Marshal(JsonData{
		Code:    code,
		Message: msg.Text(code),
		Data:    data,
	})
	w.Write(bs)
}
