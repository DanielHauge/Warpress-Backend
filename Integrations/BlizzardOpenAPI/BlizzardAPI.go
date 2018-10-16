package BlizzardOpenAPI

import (
	"github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var json = jsoniter.ConfigFastest


func GetBlizzardChar(realm string, name string, region string) (FullCharInfo, error){

	log.Infof("Fetching Blizzardchar for: {Realm: %s - Name: %s - Region: %s",realm, name, "en_GB")
	url := "https://"+region+".api.battle.net/wow/"+"character/"+realm+"/"+name+"?fields=guild+items&locale=en_GB&apikey="+os.Getenv("BLIZZARD_APIKEY")
	resp, e := http.Get(url)
	defer resp.Body.Close()


	var fullChar FullCharInfo
	e = json.NewDecoder(resp.Body).Decode(&fullChar)
	if e != nil { log.Error(e, " -> Something went wrong with decoding blizzard") }

	return fullChar, e
}
