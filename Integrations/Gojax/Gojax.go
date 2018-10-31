package Gojax

import (
	log "../../Utility/Logrus"
	"encoding/json"
	"net/http"
)

func Get(url string, obj interface{}) error {

	resp, e := http.Get(url)
	if e != nil {
		log.WithLocation().WithError(e).WithField("URL", url).Error("gojax failed")
	}
	defer resp.Body.Close()

	e = json.NewDecoder(resp.Body).Decode(&obj)
	if e != nil {
		log.WithLocation().WithError(e).WithField("Struct", obj).WithField("URL", url).Error("json decoding failed")
	}

	return e
}
