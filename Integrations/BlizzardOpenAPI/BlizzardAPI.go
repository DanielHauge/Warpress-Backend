package BlizzardOpenAPI

import (
	"github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strings"
)

var json = jsoniter.ConfigFastest


func GetBlizzardChar(realm string, name string, region string) (FullCharInfo, error){

	log.Infof("Fetching Blizzardchar for: {Realm: %s - Name: %s - Region: %s",realm, name, "en_GB")
	url := "https://"+region+".api.battle.net/wow/"+"character/"+realm+"/"+name+"?fields=guild+items&locale=en_GB&apikey="+os.Getenv("BLIZZARD_APIKEY")
	resp, e := http.Get(url)
	defer resp.Body.Close()


	var fullChar FullCharInfo
	e = json.NewDecoder(resp.Body).Decode(&fullChar)
	if e != nil { log.Error(e, " -> Something went wrong with decoding blizzard char") }

	return fullChar, e
}

func GetBlizzardGuildMembers(guildname string, region string, realm string) (GuildWithMembers, error){
	urlreadyGuildname := strings.Replace(guildname, " ",  "%20", -1)

	log.Infof("Fetching Blizzard Guild members for: {Guildname: %s - Region: %s - realm: %s", guildname, region, realm )
	url := "https://"+region+".api.battle.net/wow/"+"guild/"+realm+"/"+urlreadyGuildname+"?fields=members&locale=en_GB&apikey="+os.Getenv("BLIZZARD_APIKEY")
	resp, e := http.Get(url)
	defer resp.Body.Close()


	var guild GuildWithMembers
	e = json.NewDecoder(resp.Body).Decode(&guild)
	if e != nil { log.Error(e, " -> Something went wrong with decoding blizzard guild") }

	return guild, e

}
