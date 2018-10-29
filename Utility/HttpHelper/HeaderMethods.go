package HttpHelper

import (
	"github.com/json-iterator/go"
	"net/http"
)

var json = jsoniter.ConfigFastest

func SuccessHeader(w http.ResponseWriter, msg []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
	w.Write(msg)
}

func InterErrorHeader(w http.ResponseWriter, e error, message string, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	msg, _ := json.Marshal(jsonError{Message:message, Error:e.Error(), Status:status})
	w.Write(msg)
}

type jsonError struct {
	Message string `json:"message"`
	Error string `json:"error"`
	Status int `json:"status"`
}
