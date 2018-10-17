package Filters

import (
	"net/http"
	"time"
	"../Prometheus"
)

func Monitor(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		end := time.Since(start)
		Prometheus.ReqDurationObserve(end.Seconds())
		Prometheus.RequestInc()

	})
}
