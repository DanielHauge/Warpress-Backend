package Wowprogress

import (
	log "../../Logrus"
	"../../Prometheus"
	"../Gojax"
	"github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
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
	log.WithFields(logrus.Fields{"Guild":input.Guild,"Realm":input.Realm,"Region":input.Region}).Info("Gojaxing wowprogress ranks")
	fullUrl := "https://www.wowprogress.com/guild/" + input.Region + "/" + input.Realm + "/" + strings.Replace(input.Guild, " ", "+", -1) + "/json_rank"

	var rankings GuildRank

	now := time.Now()
	e := Gojax.Get(fullUrl, &rankings)
	Prometheus.JaxObserveWowprogress(time.Since(now).Seconds())

	return rankings, e
}
