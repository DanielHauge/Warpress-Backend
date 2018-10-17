package Wowprogress

import (
	"../Gojax"
	"github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"strings"
)

var json = jsoniter.ConfigFastest

type Input struct {
	Region string
	Realm  string
	Guild  string
}

type GuildRank struct {
	Score     int `json:"score"`
	WorldRank int `json:"world_rank"`
	AreaRank  int `json:"area_rank"`
	RealmRank int `json:"realm_rank"`
}

func GetGuildRank(input Input) (GuildRank, error) {
	log.Info("Fetching wowprogress Guildrank for: ", input)
	fullUrl := "https://www.wowprogress.com/guild/" + input.Region + "/" + input.Realm + "/" + strings.Replace(input.Guild, " ", "+", -1) + "/json_rank"

	var rankings GuildRank

	e := Gojax.Get(fullUrl, &rankings)

	return rankings, e
}
