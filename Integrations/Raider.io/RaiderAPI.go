package Raider_io

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

func GetRaiderIORank(input CharInput) (CharacterProfile, error) {
	log.WithFields(logrus.Fields{"Character":input.Name,"Realm":input.Realm,"Region":input.Region}).Info("Gojaxing Raider.io ranks for character")
	url := "https://raider.io/api/v1/characters/profile?region=" + input.Region + "&realm=" + input.Realm + "&name=" + input.Name + "&fields=mythoc_plus_scores%2Cmythic_plus_ranks%2Cmythic_plus_recent_runs%2Cmythic_plus_highest_level_runs%2Cmythic_plus_weekly_highest_level_runs%2C"

	var rankings CharacterProfile
	now := time.Now()
	e := Gojax.Get(url, &rankings)
	Prometheus.JaxObserveRaiderio(time.Since(now).Seconds())
	return rankings, e
}

func GetRaiderIOGuild(region string, realm string, guildname string) (GuildInfo, error) {
	log.WithFields(logrus.Fields{"Guild":guildname,"Realm":realm,"Region":region}).Info("Gojaxing Raider.io ranks for guild")
	urlguildname := strings.Replace(guildname, " ", "%20", -1)
	urlrealm := strings.Replace(realm, " ", "%20", -1)
	url := "https://raider.io/api/v1/guilds/profile?region=" + region + "&realm=" + urlrealm + "&name=" + urlguildname + "&fields=raid_progression%2Craid_rankings"

	var guildinfo GuildInfo
	now := time.Now()
	e := Gojax.Get(url, &guildinfo)
	Prometheus.JaxObserveRaiderio(time.Since(now).Seconds())
	return guildinfo, e

}
