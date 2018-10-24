package HttpHelper

import (
	log "../Logrus"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func ReadFromRequest(w http.ResponseWriter, r *http.Request, obj interface{}) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.WithLocation().WithError(err).Error("Hov!")
		w.WriteHeader(400)
		w.Write([]byte("Could not read body"))
		return
	}
	if err := r.Body.Close(); err != nil {
		log.WithLocation().WithError(err).Error("Hov!")
	}
	if err := json.Unmarshal(body, &obj); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.WithLocation().WithError(err).Error("Hov!")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

}
