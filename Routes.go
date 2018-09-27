package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	ExpectedInput interface{}
	ExpectedOutput interface{}
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
		ExampleInput{"hello", 5},
		ExampleOutput{[]int{1,2,3,4,5,6}},

	},

}