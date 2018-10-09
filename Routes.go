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
	Route{
		"GetChars",
		"GET",
		"/chars",
		GetCharactersForRegistration,
		nil,
		nil,
	},
	Route{
		"SetMain",
		"POST",
		"/main",
		SetMainCharacter,
		charRequest{Name:"Rakhoal", Realm:"Twisting-Nether", Locale:"en_GB"},
		nil,
	},
	Route{
		"GetMain",
		"GET",
		"/main",
		GetMainCharacter,
		nil,
		charRequest{Name:"Rakhoal", Realm:"Twisting-Nether", Locale:"en_GB"},
	},
	Route{
		"Get Full Char Information",
		"POST",
		"/char",
		GetFullCharHandle,
		nil,
		nil,
	},
	Route{
		"Get Full Personal View",
		"GET",
		"/personal",
		GetPersonalCharInfo,
		nil,
		nil,
	},
}

var restrictedRoutes = Routes{
	Route{
		"Authenticate",
		"GET",
		"/bnet/auth",
		HandleAuthenticate,
		nil,
		nil,
	},
	Route{
		"Callback",
		"GET",
		"/bnet/auth/callback",
		HandleOauthCallback,
		nil,
		nil,
	},

}