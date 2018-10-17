package Prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

var promRequests = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "requests_total",
			Help: "The ammount of requests that has occured since start",
		},
)

var promLogins = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "logins_total",
			Help: "The ammount of logins that has occured since start",
		},
)

var promRequestDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name: "request_duration_all",
	Help: "The duration of the requests from any route",
})



func init(){
	prometheus.MustRegister(promRequests)
	prometheus.MustRegister(promLogins)
	prometheus.MustRegister(promRequestDuration)

}

func LoginInc(){
	promLogins.Inc()
}

func RequestInc(){
	promRequests.Inc()
}

func ReqDurationObserve(dur float64){
	promRequestDuration.Observe(dur)
}





