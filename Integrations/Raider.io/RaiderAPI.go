package Raider_io

import (
	"github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

var json = jsoniter.ConfigFastest

func GetRaiderIORank(input CharInput) (CharacterProfile, error){
	log.Info("Fetching RaiderIO profile for: ",input)
	url := "https://raider.io/api/v1/characters/profile?region="+input.Region+"&realm="+input.Realm+"&name="+input.Name+"&fields=mythoc_plus_scores%2Cmythic_plus_ranks%2Cmythic_plus_recent_runs%2Cmythic_plus_highest_level_runs%2Cmythic_plus_weekly_highest_level_runs%2C"

	resp, e := http.Get(url)
	if e != nil{
		log.Error(e, " -> Something went wrong in getting raider data from RaiderIO")
		return CharacterProfile{}, e
	}
	defer resp.Body.Close()

	var rankings CharacterProfile
	e = json.NewDecoder(resp.Body).Decode(&rankings)
	if e != nil { log.Error(e, "Something went wrong in decoding data from RaiderIO") }

	return rankings, e
}

func GetRaiderIOGuild(region string, realm string, guildname string) (GuildInfo, error){
	log.Info("Fetching RaiderIO Guild Profile for: {Guild: %s - Realm: %s - Region: %s", guildname, realm, region)
	urlguildname := strings.Replace(guildname, " ", "%20", -1)
	urlrealm := strings.Replace(realm, " ", "%20", -1)
	url := "https://raider.io/api/v1/guilds/profile?region="+region+"&realm="+urlrealm+"&name="+urlguildname+"&fields=raid_progression%2Craid_rankings"

	resp, e := http.Get(url)
	if e != nil{
		log.Error(e, " -> Something went wrong in getting guild data from RaiderIO")
		return GuildInfo{}, e
	}
	defer resp.Body.Close()

	var guildinfo GuildInfo
	e = json.NewDecoder(resp.Body).Decode(&guildinfo)
	if e != nil { log.Error(e, "Something went wrong in decoding data from RaiderIO") }

	return guildinfo, e

}

