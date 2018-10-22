package Filters

import (
	"../Monitoring"
	"net/http"
	"time"
)

func Monitor(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		end := time.Since(start)
		Monitoring.ReqDurationObserve(end.Seconds())
		Monitoring.RequestInc()

	})
}
