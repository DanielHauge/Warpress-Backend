package BlizzardOpenAPI

import (
	log "../../Utility/Logrus"
	"../../Utility/Monitoring"
	"../Gojax"
	"github.com/avelino/slugify"
	"github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

var json = jsoniter.ConfigFastest

func GetBlizzardChar(realm string, name string, region string) (FullCharInfo, error) {

	url := "https://" + region + ".api.battle.net/wow/" + "character/" + realm + "/" + name + "?fields=guild+items+talents&locale=en_GB&apikey=" + os.Getenv("BLIZZARD_APIKEY")
	log.WithFields(logrus.Fields{"Character": name, "Realm": realm, "Region": region, "Url": url}).Info("Gojaxing Blizzard Open for character")
	var fullChar FullCharInfo
	now := time.Now()
	e := Gojax.Get(url, &fullChar)
	Monitoring.JaxObserveBlizzardOpen(time.Since(now).Seconds())
	return fullChar, e
}

func GetBlizzardGuildMembers(guildname string, realm string, region string) (GuildWithMembers, error) {
	realm = slugify.Slugify(realm)
	urlreadyGuildname := strings.Replace(guildname, " ", "%20", -1)
	url := "https://" + region + ".api.battle.net/wow/" + "guild/" + realm + "/" + urlreadyGuildname + "?fields=members&locale=en_GB&apikey=" + os.Getenv("BLIZZARD_APIKEY")
	log.WithFields(logrus.Fields{"Guild": guildname, "Realm": realm, "Region": region, "Url": url}).Info("Gojaxing Blizzard Open for guild members")
	var guild GuildWithMembers

	now := time.Now()
	e := Gojax.Get(url, &guild)
	Monitoring.JaxObserveBlizzardOpen(time.Since(now).Seconds())
	return guild, e

}
