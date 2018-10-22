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

func GetWarcraftLogsRanks(input CharInput) ([]Encounter, error) {
	log.WithFields(logrus.Fields{"Character": input.Name, "Realm": input.Realm, "Region": input.Region}).Info("Gojaxing warcraftlogs ranks for character")
	fullUrl := warcraftLogsAPIURL + "/rankings/character/" + input.Name + "/" + input.Realm + "/" + input.Region + "?api_key=" + os.Getenv("PUBLIC_LOGS")

	var rankings []Encounter

	now := time.Now()
	e := Gojax.Get(fullUrl, &rankings)
	Monitoring.JaxObserveWarcraftlogs(time.Since(now).Seconds())
	return rankings, e
}

func GetWarcraftLogsReport(ReportId string) (Report, error) {
	log.WithField("ReportID", ReportId).Info("Gojaxing warcraftlogs report")
	fullUrl := warcraftLogsAPIURL + "/report/fights/" + ReportId + "?api_key=" + os.Getenv("PUBLIC_LOGS")

	var report Report
	now := time.Now()
	e := Gojax.Get(fullUrl, &report)
	Monitoring.JaxObserveWarcraftlogs(time.Since(now).Seconds())
	return report, e
}

func GetWarcraftGuildReports(guildname string, realm string, region string, startime int64, endtime int64) ([]GuildReports, error) {
	log.WithFields(logrus.Fields{"Guild": guildname, "Realm": realm, "Region": region, "From": startime, "To": endtime}).Info("Gojaxing guild reports")
	urlguildname := strings.Replace(guildname, " ", "%20", -1)
	fullUrl := warcraftLogsAPIURL + "/reports/guild/" + urlguildname + "/" + slugify.Slugify(realm) + "/" + region + "?start=" + strconv.FormatInt(startime, 10) + "&end=" + strconv.FormatInt(endtime, 10) + "&api_key=" + os.Getenv("PUBLIC_LOGS")

	var report []GuildReports

	now := time.Now()
	e := Gojax.Get(fullUrl, &report)
	Monitoring.JaxObserveWarcraftlogs(time.Since(now).Seconds())
	return report, e
}
