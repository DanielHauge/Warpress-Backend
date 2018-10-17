package BlizzardOpenAPI

import (
	"../Gojax"
	"github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

var json = jsoniter.ConfigFastest

func GetBlizzardChar(realm string, name string, region string) (FullCharInfo, error) {

	log.Infof("Fetching Blizzardchar for: {Realm: %s - Name: %s - Region: %s", realm, name, "en_GB")
	url := "https://" + region + ".api.battle.net/wow/" + "character/" + realm + "/" + name + "?fields=guild+items&locale=en_GB&apikey=" + os.Getenv("BLIZZARD_APIKEY")

	var fullChar FullCharInfo
	e := Gojax.Get(url, &fullChar)

	return fullChar, e
}

func GetBlizzardGuildMembers(guildname string, region string, realm string) (GuildWithMembers, error) {
	urlreadyGuildname := strings.Replace(guildname, " ", "%20", -1)

	log.Infof("Fetching Blizzard Guild members for: {Guildname: %s - Region: %s - realm: %s", guildname, region, realm)
	url := "https://" + region + ".api.battle.net/wow/" + "guild/" + realm + "/" + urlreadyGuildname + "?fields=members&locale=en_GB&apikey=" + os.Getenv("BLIZZARD_APIKEY")

	var guild GuildWithMembers

	e := Gojax.Get(url, &guild)

	return guild, e

}
