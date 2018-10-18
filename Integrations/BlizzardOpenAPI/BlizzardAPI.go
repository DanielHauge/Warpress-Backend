package BlizzardOpenAPI

import (
	log "../../Logrus"
	"../../Prometheus"
	"../Gojax"
	"github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

var json = jsoniter.ConfigFastest

func GetBlizzardChar(realm string, name string, region string) (FullCharInfo, error) {
	log.WithFields(logrus.Fields{"Character":name,"Realm":realm,"Region":region}).Info("Gojaxing Blizzard Open for character")
	url := "https://" + region + ".api.battle.net/wow/" + "character/" + realm + "/" + name + "?fields=guild+items&locale=en_GB&apikey=" + os.Getenv("BLIZZARD_APIKEY")

	var fullChar FullCharInfo
	now := time.Now()
	e := Gojax.Get(url, &fullChar)
	Prometheus.JaxObserveBlizzardOpen(time.Since(now).Seconds())
	return fullChar, e
}

func GetBlizzardGuildMembers(guildname string, region string, realm string) (GuildWithMembers, error) {
	log.WithFields(logrus.Fields{"Guild":guildname,"Realm":realm,"Region":region}).Info("Gojaxing Blizzard Open for guild members")
	urlreadyGuildname := strings.Replace(guildname, " ", "%20", -1)
	url := "https://" + region + ".api.battle.net/wow/" + "guild/" + realm + "/" + urlreadyGuildname + "?fields=members&locale=en_GB&apikey=" + os.Getenv("BLIZZARD_APIKEY")

	var guild GuildWithMembers


	now := time.Now()
	e := Gojax.Get(url, &guild)
	Prometheus.JaxObserveBlizzardOpen(time.Since(now).Seconds())
	return guild, e

}
