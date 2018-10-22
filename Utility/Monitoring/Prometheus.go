package Monitoring

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (

	promRequests = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "requests_total",
		Help: "The ammount of requests that has occured since start",
	})

	promLogins = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "logins_total",
		Help: "The ammount of logins that has occured since start",
	})

	promRequestDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "request_duration_all",
		Help: "The duration of the requests from any route",
	})

	promJaxDurationBlizOpen = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "gojax_duration_blizzardopen",
		Help: "The duration of the gojax request to blizzards open api",
	})

	promJaxDurationRaiderIO = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "gojax_duration_raiderio",
		Help: "The duration of the gojax request to raider.io api",
	})

	promJaxDurationWarcraftLogs = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "gojax_duration_warcraftlogs",
		Help: "The duration of the gojax request to warcraftlogs api",
	})

	promJaxDurationWowprogress = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "gojax_duration_wowprogress",
		Help: "The duration of the gojax request to wowprogress api",
	})

)

func init() {
	prometheus.MustRegister(promRequests)
	prometheus.MustRegister(promLogins)
	prometheus.MustRegister(promRequestDuration)
	prometheus.MustRegister(promJaxDurationWowprogress, promJaxDurationWarcraftLogs, promJaxDurationRaiderIO, promJaxDurationBlizOpen)
}

func LoginInc() {
	promLogins.Inc()
}

func RequestInc() {
	promRequests.Inc()
}

func ReqDurationObserve(dur float64) {
	promRequestDuration.Observe(dur)
}

func JaxObserveBlizzardOpen(dur float64){
	promJaxDurationBlizOpen.Observe(dur)
}

func JaxObserveRaiderio(dur float64){
	promJaxDurationRaiderIO.Observe(dur)
}

func JaxObserveWarcraftlogs(dur float64){
	promJaxDurationWarcraftLogs.Observe(dur)
}

func JaxObserveWowprogress(dur float64){
	promJaxDurationWowprogress.Observe(dur)
}
