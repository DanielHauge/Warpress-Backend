package Raider_io

import (
	"github.com/json-iterator/go"
	"log"
	"net/http"
)

var json = jsoniter.ConfigFastest

func GetRaiderIORank(input CharInput) (CharacterProfile, error){
	url := "https://raider.io/api/v1/characters/profile?region="+input.Region+"&realm="+input.Realm+"&name="+input.Name+"&fields=mythoc_plus_scores%2Cmythic_plus_ranks%2Cmythic_plus_recent_runs%2Cmythic_plus_highest_level_runs%2Cmythic_plus_weekly_highest_level_runs%2C"

	resp, e := http.Get(url)
	if e != nil{
		log.Println(e.Error())
		return CharacterProfile{}, e
	}

	var rankings CharacterProfile
	e = json.NewDecoder(resp.Body).Decode(&rankings)
	if e != nil { log.Println(e.Error()) }

	return rankings, e
}

