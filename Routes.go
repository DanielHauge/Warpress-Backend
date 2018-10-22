package main

import (
	"./Integrations/BlizzardOauthAPI"
	. "./Utility/Filters"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type Route struct {
	Name           string
	Method         string
	Pattern        string
	HandlerFunc    http.HandlerFunc
	ExpectedInput  interface{}
	ExpectedOutput interface{}
}

type Routes []Route

var routes = Routes{
	Route{
		"Index -> This page",
		"GET",
		"/",
		Index,
		nil,
		nil,
	},
	Route{
		"Get all character for requesting account",
		"GET",
		"/chars",
		RequireAuthentication(BlizzardOauthAPI.GetCharactersForRegistration),
		nil,
		nil,
	},
	Route{
		"Set main for requesting account",
		"POST",
		"/main",
		RequireAuthentication(BlizzardOauthAPI.SetMainCharacter),
		ExampleCharVar,
		nil,
	},
	Route{
		"Get main for requesting account",
		"GET",
		"/main",
		RequireAuthentication(BlizzardOauthAPI.GetMainCharacter),
		nil,
		ExampleCharVar,
	},
	Route{
		"Get Full Personal View, Includes: (BlizzardOpenAPI, Raider.io, warcraftlogs, wowprogress) profiles",
		"GET",
		"/personal",
		RequireAuthentication(HandleGetPersonalFull),
		nil,
		ExamplePersonalProfile,
	},
	Route{
		"Get Blizzards character profile",
		"GET",
		"/personal/blizzard",
		RequireAuthentication(HandleGetPersonalBlizzardChar),
		nil,
		ExampleFullBlizzChar,
	},
	Route{
		"Get Raider.IO character profile",
		"GET",
		"/personal/raiderio",
		RequireAuthentication(HandleGetPersonalRaiderio),
		nil,
		ExampleRaiderioProfile,
	},
	Route{
		"Get Warcraftlogs character profile",
		"GET",
		"/personal/warcraftlogs",
		RequireAuthentication(HandleGetPersonalWarcraftLogs),
		nil,
		nil,
	},
	Route{
		"Get personal improvements",
		"GET",
		"/personal/improvements",
		RequireAuthentication(HandleGetPersonalImprovements),
		nil,
		ExamplePersonalImprovement,
	},
	Route{
		"Get Guild overview",
		"GET",
		"/guild",
		RequireAuthentication(HandleGetGuildOverview),
		nil,
		ExampleGuildOverviewInfo,
	},
	Route{
		"Log out",
		"GET",
		"/bnet/logout",
		RequireAuthentication(HandleLogout),
		nil,
		nil,
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
		"Monitoring Metrics",
		"GET",
		"/metrics",
		promhttp.Handler().ServeHTTP,
		nil,
		nil,
	},
}
