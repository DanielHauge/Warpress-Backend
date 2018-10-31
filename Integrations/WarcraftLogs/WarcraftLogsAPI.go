package WarcraftLogs

import (
	log "../../Utility/Logrus"
	"../../Utility/Monitoring"
	"../Gojax"
	"github.com/avelino/slugify"
	"github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
	"time"
)

var json = jsoniter.ConfigFastest

var warcraftLogsAPIURL = "https://www.warcraftlogs.com:443/v1"

func GetWarcraftLogsRanks(input CharInput) (Encounters, error) {

	fullUrl := warcraftLogsAPIURL + "/rankings/character/" + input.Name + "/" + input.Realm + "/" + input.Region + "?api_key=" + os.Getenv("PUBLIC_LOGS")
	log.WithFields(logrus.Fields{"Character": input.Name, "Realm": input.Realm, "Region": input.Region, "Url": fullUrl}).Info("Gojaxing warcraftlogs ranks for character")
	var rankings []Encounter

	now := time.Now()
	e := Gojax.Get(fullUrl, &rankings)
	Monitoring.JaxObserveWarcraftlogs(time.Since(now).Seconds())
	return Encounters{Encounters: rankings}, e
}

func GetWarcraftLogsReport(ReportId string) (Report, error) {

	fullUrl := warcraftLogsAPIURL + "/report/fights/" + ReportId + "?api_key=" + os.Getenv("PUBLIC_LOGS")
	log.WithField("ReportID", ReportId).WithField("Url", fullUrl).Info("Gojaxing warcraftlogs report")

	var report Report
	now := time.Now()
	e := Gojax.Get(fullUrl, &report)
	Monitoring.JaxObserveWarcraftlogs(time.Since(now).Seconds())
	return report, e
}

func GetWarcraftGuildReports(guildname string, realm string, region string, startime int64, endtime int64) ([]GuildReports, error) {

	urlguildname := strings.Replace(guildname, " ", "%20", -1)
	fullUrl := warcraftLogsAPIURL + "/reports/guild/" + urlguildname + "/" + slugify.Slugify(realm) + "/" + region + "?start=" + strconv.FormatInt(startime*1000, 10) + "&end=" + strconv.FormatInt(endtime*1000, 10) + "&api_key=" + os.Getenv("PUBLIC_LOGS")
	log.WithFields(logrus.Fields{"Guild": guildname, "Realm": realm, "Region": region, "From": startime * 1000, "To": endtime * 1000, "Url": fullUrl}).Info("Gojaxing guild reports")
	var report []GuildReports

	now := time.Now()
	e := Gojax.Get(fullUrl, &report)
	Monitoring.JaxObserveWarcraftlogs(time.Since(now).Seconds())
	return report, e
}
