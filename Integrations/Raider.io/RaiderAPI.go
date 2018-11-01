package Raider_io

import (
	log "../../Utility/Logrus"
	"../../Utility/Monitoring"
	"../Gojax"
	"github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	url2 "net/url"
	"strings"
	"time"
)

var json = jsoniter.ConfigFastest

func GetRaiderIORank(name string, realm string, region string) (CharacterProfile, error) {

	url := "https://raider.io/api/v1/characters/profile?region="+region+"&realm="+realm+"&name="+url2.QueryEscape(name)+"&fields=mythic_plus_scores%2Cmythic_plus_ranks%2Cmythic_plus_recent_runs%2Cmythic_plus_highest_level_runs%2Cmythic_plus_weekly_highest_level_runs%2C"

	log.WithFields(logrus.Fields{"Character": name, "Realm": realm, "Region": region, "Url": url}).Info("Gojaxing Raider.io ranks for character")
	var rankings CharacterProfile
	now := time.Now()
	e := Gojax.Get(url, &rankings)
	Monitoring.JaxObserveRaiderio(time.Since(now).Seconds())
	return rankings, e
}

func GetRaiderIOGuild(region string, realm string, guildname string) (GuildInfo, error) {

	urlguildname := strings.Replace(guildname, " ", "%20", -1)
	urlrealm := strings.Replace(realm, " ", "%20", -1)
	url := "https://raider.io/api/v1/guilds/profile?region=" + region + "&realm=" + urlrealm + "&name=" + urlguildname + "&fields=raid_progression%2Craid_rankings"
	log.WithFields(logrus.Fields{"Guild": guildname, "Realm": realm, "Region": region, "Url": url}).Info("Gojaxing Raider.io ranks for guild")
	var guildinfo GuildInfo
	now := time.Now()
	e := Gojax.Get(url, &guildinfo)
	Monitoring.JaxObserveRaiderio(time.Since(now).Seconds())
	return guildinfo, e

}
