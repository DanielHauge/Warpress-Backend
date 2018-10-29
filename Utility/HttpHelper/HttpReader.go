package HttpHelper

import (
	log "../Logrus"
	"io"
	"io/ioutil"
	"net/http"
)

func ReadFromRequest(w http.ResponseWriter, r *http.Request, obj interface{}) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.WithLocation().WithError(err).Error("Hov!")
		InterErrorHeader(w,err, "Tried to read request, but could not read request.", http.StatusRequestEntityTooLarge)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.WithLocation().WithError(err).Error("Hov!")
	}
	if err := json.Unmarshal(body, &obj); err != nil {
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.WithLocation().WithError(err).Error("Hov!")
			InterErrorHeader(w,err, "Tried to read request, but was not formatted as expected", 400)
			return
		}
	}

}
