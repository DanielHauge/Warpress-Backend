package main

import (
	"./Integrations/BlizzardOauthAPI"
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
		BlizzardOauthAPI.GetCharactersForRegistration,
		nil,
		nil,
	},
	Route{
		"Set main for requesting account",
		"POST",
		"/main",
		BlizzardOauthAPI.SetMainCharacter,
		BlizzardOauthAPI.CharacterMinimal{Name: "Rakhoal", Realm:"Twisting-Nether", Locale:"en_GB"},
		nil,
	},
	Route{
		"Get main for requesting account",
		"GET",
		"/main",
		BlizzardOauthAPI.GetMainCharacter,
		nil,
		BlizzardOauthAPI.CharacterMinimal{Name: "Rakhoal", Realm:"Twisting-Nether", Locale:"en_GB"},
	},
	Route{
		"Get Full Personal View, Includes: (BlizzardOpenAPI, Raider.io, warcraftlogs, wowprogress) profiles",
		"GET",
		"/personal",
		HandleGetPersonalFull,
		nil,
		ExamplePleaseTryIt{AlotOfJson:"Please Try It And see."},
	},
	Route{
		"Get Blizzards character profile",
		"GET",
		"/personal/blizzard",
		HandleGetPersonalBlizzardChar,
		nil,
		ExamplePleaseTryIt{AlotOfJson:"Please Try It And see."},
	},
	Route{
		"Get Raider.IO character profile",
		"GET",
		"/personal/raiderio",
		HandleGetPersonalRaiderio,
		nil,
		ExamplePleaseTryIt{AlotOfJson:"Please Try It And see."},
	},
	Route{
		"Get Warcraftlogs character profile",
		"GET",
		"/personal/warcraftlogs",
		HandleGetPersonalWarcraftLogs,
		nil,
		ExamplePleaseTryIt{AlotOfJson:"Please Try It And see."},
	},
	Route{
		"Get personal improvements",
		"GET",
		"/personal/improvements",
		HandleGetPersonalImprovements,
		nil,
		ExamplePleaseTryIt{AlotOfJson:"Please Try It And see"},
	},
}

var restrictedRoutes = Routes{
	Route{
		"Authenticate",
		"GET",
		"/bnet/auth",
		BlizzardOauthAPI.HandleAuthenticate,
		nil,
		nil,
	},
	Route{
		"Callback",
		"GET",
		"/bnet/auth/callback",
		BlizzardOauthAPI.HandleOauthCallback,
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