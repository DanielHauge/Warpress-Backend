package Blizzard

import (
	"github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var json = jsoniter.ConfigFastest

var ApiURL = "https://eu.api.battle.net/wow/"

func GetBlizzardChar(char CharacterMinimal) (FullCharInfo, error){

	url := ApiURL+"character/"+char.Realm+"/"+char.Name+"?fields=guild+items&locale="+char.Locale+"&apikey="+os.Getenv("BLIZZARD_APIKEY")
	resp, e := http.Get(url)
	defer resp.Body.Close()

	var fullChar FullCharInfo
	e = json.NewDecoder(resp.Body).Decode(&fullChar)
	if e != nil { log.Error(e, " -> Something went wrong with decoding blizzard") }

	return fullChar, e
}
