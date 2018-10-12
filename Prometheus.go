package main

import "github.com/prometheus/client_golang/prometheus"


var (

	promRequests = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "requests",
			Help: "The ammount of requests that has occured since start",
		},
	)
	promLogins = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "logins",
			Help: "The ammount of logins that has occured since start",
		})



)



func init(){
	prometheus.MustRegister(promRequests)
}
