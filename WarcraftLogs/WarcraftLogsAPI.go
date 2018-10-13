package WarcraftLogs

import (
	"github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var json = jsoniter.ConfigFastest

var warcraftLogsAPIURL = "https://www.warcraftlogs.com:443/v1"

func GetWarcraftLogsRanks(input CharInput) ([]Encounter, error){
	fullUrl := warcraftLogsAPIURL+"/rankings/character/"+input.Name+"/"+input.Realm+"/"+input.Region+"?api_key="+os.Getenv("PUBLIC_LOGS")

	resp, e := http.Get(fullUrl)
	if e != nil{
		log.Error(e)
		return []Encounter{}, e
	}

	var rankings []Encounter
	e = json.NewDecoder(resp.Body).Decode(&rankings)
	if e != nil { log.Error(e) }

	return rankings, e
}