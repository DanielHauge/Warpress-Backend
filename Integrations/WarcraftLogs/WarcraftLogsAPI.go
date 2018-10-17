package WarcraftLogs

import (
	"../Gojax"
	"github.com/avelino/slugify"
	"github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

var json = jsoniter.ConfigFastest

var warcraftLogsAPIURL = "https://www.warcraftlogs.com:443/v1"

func GetWarcraftLogsRanks(input CharInput) ([]Encounter, error){
	log.Info("Fetching warcraftlogs ranks for: ", input)
	fullUrl := warcraftLogsAPIURL+"/rankings/character/"+input.Name+"/"+input.Realm+"/"+input.Region+"?api_key="+os.Getenv("PUBLIC_LOGS")

	var rankings []Encounter

	e := Gojax.Get(fullUrl, &rankings)


	return rankings, e
}

func GetWarcraftLogsReport(ReportId string) (Report, error){
	log.Info("Fetching warcraftlogs reports for "+ReportId)
	fullUrl := warcraftLogsAPIURL+"/report/fights/"+ReportId+"?api_key="+os.Getenv("PUBLIC_LOGS")


	var report Report
	e := Gojax.Get(fullUrl, &report)

	return report, e
}

func GetWarcraftGuildReports(guildname string, realm string, region string, startime int64, endtime int64) ([]GuildReports, error){
	log.Infof("Fetching warcraftlogs guild reports for {Guild: %s - Realm: %s - Region: %s - Starttime %s }", guildname, realm, region, startime)
	urlguildname := strings.Replace(guildname, " ", "%20", -1)
	fullUrl := warcraftLogsAPIURL+"/reports/guild/"+urlguildname+"/"+slugify.Slugify(realm)+"/"+region+"?start="+strconv.FormatInt(startime, 10)+"&end="+strconv.FormatInt(endtime, 10)+"&api_key="+os.Getenv("PUBLIC_LOGS")


	var report []GuildReports

	e := Gojax.Get(fullUrl, &report)

	return report, e
}