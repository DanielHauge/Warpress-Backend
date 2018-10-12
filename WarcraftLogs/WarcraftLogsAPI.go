package WarcraftLogs

import (
	"log"
	"github.com/json-iterator/go"
	"net/http"
	"os"
)

var json = jsoniter.ConfigFastest

var warcraftLogsAPIURL = "https://www.warcraftlogs.com:443/v1"

func GetWarcraftLogsRanks(input CharInput) ([]Encounter, error){
	fullUrl := warcraftLogsAPIURL+"/rankings/character/"+input.Name+"/"+input.Realm+"/"+input.Region+"?api_key="+os.Getenv("PUBLIC_LOGS")

	resp, e := http.Get(fullUrl)
	if e != nil{
		log.Println(e.Error())
		return []Encounter{}, e
	}

	var rankings []Encounter
	e = json.NewDecoder(resp.Body).Decode(&rankings)
	if e != nil { log.Println(e.Error()) }

	return rankings, e
}