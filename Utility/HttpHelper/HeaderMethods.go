package HttpHelper

import "net/http"

func SuccessHeader(w http.ResponseWriter, msg []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
	w.Write(msg)
}

func InterErrorHeader(w http.ResponseWriter, e error) {
	w.WriteHeader(500)
	w.Write([]byte(e.Error()))
}
