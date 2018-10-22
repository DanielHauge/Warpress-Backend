package Filters

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.WithFields(log.Fields{
			"Method": r.Method,
			"Route": r.RequestURI,
			"Name":name,
			"Duration":time.Since(start),
		}).Info("Request")
	})
}
