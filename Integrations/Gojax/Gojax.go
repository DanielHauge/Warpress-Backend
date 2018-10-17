package Gojax

import (
	"encoding/json"
	"github.com/prometheus/common/log"
	"net/http"
)

func Get(url string, obj interface{}) error{

	resp, e := http.Get(url)
	if e != nil{
		log.Error(e, " -> Something went wrong in getting data from: ", url)
	}
	defer resp.Body.Close()

	e = json.NewDecoder(resp.Body).Decode(&obj)
	if e != nil { log.Error(e, "Something went wrong in decoding", obj) }

	return e
}
