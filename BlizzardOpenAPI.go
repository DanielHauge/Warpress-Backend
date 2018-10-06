package main

import (
	"./WoWDataStructure"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

var ApiURL = "https://eu.api.battle.net/wow/"

func getCharInfo(char CharInfo) WoWDataStructure.FullCharInfo{

	url := ApiURL+"/character/"+char.Realm+"/"+char.Name+"?locale="+char.Locale+"apikey="+os.Args[4]
	resp, e := http.Get(url)

	var fullChar WoWDataStructure.FullCharInfo
	e = json.NewDecoder(resp.Body).Decode(fullChar)
	if e != nil { log.Println(e.Error()) }

	return fullChar
}