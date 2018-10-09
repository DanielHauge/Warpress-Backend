package Wowprogress

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Input struct {
	Region string
	Realm  string
	Guild  string
}

type GuildRank struct {
	Score int `json:"score"`
	WorldRank int `json:"world_rank"`
	AreaRank int `json:"area_rank"`
	RealmRank int `json:"realm_rank"`
}

func GetGuildRank(input Input) (GuildRank, error){

	fullUrl := "https://www.wowprogress.com/guild/"+input.Region +"/"+input.Realm+"/"+strings.Replace(input.Guild, " ", "+", -1)+"/json_rank"

	resp, e := http.Get(fullUrl)
	if e != nil{
		log.Println(e.Error())
		return GuildRank{}, e
	}

	var rankings GuildRank
	e = json.NewDecoder(resp.Body).Decode(&rankings)
	if e != nil { log.Println(e.Error()) }

	return rankings, e
}