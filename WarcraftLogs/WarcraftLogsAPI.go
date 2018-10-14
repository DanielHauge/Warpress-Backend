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
	log.Info("Fetching warcraftlogs ranks for: ", input)
	fullUrl := warcraftLogsAPIURL+"/rankings/character/"+input.Name+"/"+input.Realm+"/"+input.Region+"?api_key="+os.Getenv("PUBLIC_LOGS")
	resp, e := http.Get(fullUrl)
	if e != nil{
		log.Error(e, " -> Something went wrong in getting data from warcraftlogs")
		return []Encounter{}, e
	}
	defer resp.Body.Close()

	var rankings []Encounter
	e = json.NewDecoder(resp.Body).Decode(&rankings)
	if e != nil { log.Error(e, "-> Something went wrong with decoding it from warcraftlogs") }


	return rankings, e
}

func GetWarcraftLogsReport(ReportId string) (Report, error){
	log.Info("Fetching warcraftlogs reports for "+ReportId)
	fullUrl := warcraftLogsAPIURL+"/report/fights/"+ReportId+"?api_key="+os.Getenv("PUBLIC_LOGS")
	resp, e := http.Get(fullUrl)
	if e != nil{
		log.Error(e, " -> Something went wrong in getting reports from warcraftlogs")
		return Report{}, e
	}
	defer resp.Body.Close()

	var report Report
	e = json.NewDecoder(resp.Body).Decode(&report)
	if e != nil { log.Error(e, "-> Something went wrong with decoding it from warcraftlogs") }

	return report, e
}