package Wowprogress

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
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
		log.Error(e, " -> Something went wrong in getting data from wowprogress")
		return GuildRank{}, e
	}
	defer resp.Body.Close()

	var rankings GuildRank
	e = json.NewDecoder(resp.Body).Decode(&rankings)
	if e != nil { log.Error(e, "Something went wrong in decoding wowprogress") }

	return rankings, e
}