package main

import (
	"github.com/gorilla/mux"
	"net/http"
	. "./Filters"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	combinedRoutes := append(routes, restrictedRoutes...)
	for _, route := range combinedRoutes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		handler = Monitor(handler)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}