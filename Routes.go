package main

import (
	"./DataFormatters/Guild"
	"./DataFormatters/Personal"
	"./Integrations/BlizzardOauthAPI"
	"./Integrations/BlizzardOpenAPI"
	"./Integrations/Raider.io"
	"./Integrations/WarcraftLogs"
	"./Postgres/DataModel"
	. "./Utility/Filters"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
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
		struct {
			Name   string
			Reealm string
		}{},
		nil,
	},
	Route{
		"Get main for requesting account",
		"GET",
		"/main",
		RequireAuthentication(BlizzardOauthAPI.GetMainCharacter),
		nil,
		struct {
			Name   string
			Reealm string
			Region string
		}{},
	},
	Route{
		"Get Full Personal View, Includes: (BlizzardOpenAPI, Raider.io, warcraftlogs, wowprogress) profiles",
		"GET",
		"/personal",
		RequireAuthentication(HandleGetPersonalFull),
		nil,
		Personal.Overview{},
	},
	Route{
		"Get Blizzards character profile unsanitized",
		"GET",
		"/personal/blizzard",
		RequireAuthentication(HandleGetPersonalBlizzardChar),
		nil,
		BlizzardOpenAPI.FullCharInfo{},
	},
	Route{
		"Get Raider.IO character profile unsanitized",
		"GET",
		"/personal/raiderio",
		RequireAuthentication(HandleGetPersonalRaiderio),
		nil,
		Raider_io.CharacterProfile{},
	},
	Route{
		"Get Warcraftlogs character profile unsanitized",
		"GET",
		"/personal/warcraftlogs",
		RequireAuthentication(HandleGetPersonalWarcraftLogs),
		nil,
		WarcraftLogs.Encounters{},
	},
	Route{
		"Get personal improvements",
		"GET",
		"/personal/improvements",
		RequireAuthentication(HandleGetPersonalImprovements),
		nil,
		Personal.Improvements{},
	},
	Route{
		"Get guild overview",
		"GET",
		"/guild",
		RequireAuthentication(RequireTrial(HandleGetGuildOverview)),
		nil,
		Guild.Overview{},
	},
	Route{
		"Log out",
		"GET",
		"/bnet/logout",
		RequireAuthentication(HandleLogout),
		nil,
		nil,
	},
	Route{
		"Registrer guild",
		"POST",
		"/guild",
		RequireAuthentication(RequireGuildMaster(HandleGuildRegistration)),
		struct {
			Officerrank int `json:"officerrank"`
			Raiderrank  int `json:"raiderrank"`
			Trialrank   int `json:"trialrank"`
		}{},
		nil,
	},
	Route{
		"Add Raidnight",
		"POST",
		"/guild/raids",
		RequireAuthentication(RequireOfficer(HandleAddRaidNight)),
		struct {
			Duration time.Duration `json:"duration"`
			Start    time.Duration     `json:"start"`
			Day      int           `json:"day"`
		}{},
		nil,
	},
	Route{
		"Edit Raidnight",
		"PUT",
		"/guild/raids",
		RequireAuthentication(RequireOfficer(HandleEditRaidNight)),
		struct {
			Duration    time.Duration `json:"duration"`
			Start       time.Duration     `json:"start"`
			Day         int           `json:"day"`
			RaidnightId int           `json:"raidnight_id"`
		}{},
		nil,
	},
	Route{
		"Delete Raidnight",
		"DELETE",
		"/guild/raids",
		RequireAuthentication(RequireOfficer(HandleDeleteRaidNight)),
		struct {
			URLParameter string
		}{URLParameter: "id"},
		nil,
	},
	Route{
		"Get Raidnights",
		"GET",
		"/guild/raids",
		RequireAuthentication(RequireTrial(HandleGetRaidNights)),
		nil,
		DataModel.RaidNight{},
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
