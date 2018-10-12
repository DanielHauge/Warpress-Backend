package main

import (
	"./Blizzard"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
		"Index -> This page",
		"GET",
		"/",
		Index,
		ExampleInput{"hello", 5},
		ExampleOutput{[]int{1,2,3,4,5,6}},
	},
	Route{
		"Get all character for requesting account",
		"GET",
		"/chars",
		GetCharactersForRegistration,
		nil,
		nil,
	},
	Route{
		"Set main for requesting account",
		"POST",
		"/main",
		SetMainCharacter,
		Blizzard.CharacterMinimal{Name: "Rakhoal", Realm:"Twisting-Nether", Locale:"en_GB"},
		nil,
	},
	Route{
		"Get main for requesting account",
		"GET",
		"/main",
		GetMainCharacter,
		nil,
		Blizzard.CharacterMinimal{Name: "Rakhoal", Realm:"Twisting-Nether", Locale:"en_GB"},
	},
	Route{
		"Get Full Personal View, Includes: (Blizzard, Raider.io, warcraftlogs, wowprogress) profiles",
		"GET",
		"/personal",
		GetPersonalFull,
		nil,
		ExamplePleaseTryIt{AlotOfJson:"Please Try It And see."},
	},
	Route{
		"Get Blizzards character profile",
		"GET",
		"/personal/blizzard",
		GetPersonalBlizzardChar,
		nil,
		ExamplePleaseTryIt{AlotOfJson:"Please Try It And see."},
	},
	Route{
		"Get Raider.IO character profile",
		"GET",
		"/personal/raiderio",
		GetPersonalRaiderio,
		nil,
		ExamplePleaseTryIt{AlotOfJson:"Please Try It And see."},
	},
	Route{
		"Get Warcraftlogs character profile",
		"GET",
		"/personal/warcraftlogs",
		GetPersonalWarcraftLogs,
		nil,
		ExamplePleaseTryIt{AlotOfJson:"Please Try It And see."},
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
	Route{
		"Prometheus Metrics",
		"GET",
		"/metrics",
		promhttp.Handler().ServeHTTP,
		nil,
		nil,
	},

}