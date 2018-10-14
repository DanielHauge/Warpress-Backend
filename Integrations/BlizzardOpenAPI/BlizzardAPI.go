package BlizzardOpenAPI

import (
	"github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var json = jsoniter.ConfigFastest

var ApiURL = "https://eu.api.battle.net/wow/"

func GetBlizzardChar(realm string, name string, locale string) (FullCharInfo, error){
	log.Infof("Fetching Blizzardchar for: {Realm: %s - Name: %s - Locale: %s",realm, name, locale)
	url := ApiURL+"character/"+realm+"/"+name+"?fields=guild+items&locale="+locale+"&apikey="+os.Getenv("BLIZZARD_APIKEY")
	resp, e := http.Get(url)
	defer resp.Body.Close()


	var fullChar FullCharInfo
	e = json.NewDecoder(resp.Body).Decode(&fullChar)
	if e != nil { log.Error(e, " -> Something went wrong with decoding blizzard") }

	return fullChar, e
}
