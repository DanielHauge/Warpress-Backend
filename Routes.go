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
		CharacterMinimal{Name: "Rakhoal", Realm:"Twisting-Nether", Locale:"en_GB"},
		nil,
	},
	Route{
		"GetMain",
		"GET",
		"/main",
		GetMainCharacter,
		nil,
		CharacterMinimal{Name: "Rakhoal", Realm:"Twisting-Nether", Locale:"en_GB"},
	},
	Route{
		"Get All Characters for an account",
		"POST",
		"/char",
		GetFullCharHandle,
		CharacterMinimal{Name: "Rakhoal", Realm:"Twisting-Nether", Locale:"en_GB"},
		ExampleFullBlizzChar,
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

}