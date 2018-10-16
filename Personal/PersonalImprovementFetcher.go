package Personal

import (
	"../Integrations/BlizzardOauthAPI"
	"../Integrations/WarcraftLogs"
	"../Redis"
	log "github.com/sirupsen/logrus"
	"strconv"
)

var RaidBotUrl = "https://www.raidbots.com/simbot/"

func FetchPersonalImprovementsFull(id int, improvements *PersonalImprovement) error{


	improvements.SimulationURLS = MakeSimBotUrls(id)
	bossimprov, e := GenerateWarcraftLogs(id)
	improvements.BossImprovements = bossimprov

	return e
}

func GenerateWarcraftLogs(id int) ([]BossImprovement, error) {
	var logs []WarcraftLogs.Encounter
	var improvements []BossImprovement
	e := FetchWarcraftlogsPersonal(id, &logs)
	if e != nil{
		log.Error(e, " Something went wrong in Fetching warcraftlogs ranks")
	}

	mapOfCharIds := map[string]int{}
	for _, value := range logs {

		comparelink := GenerateCompareLink(value.ReportID, value.FightID, value.CharacterName, mapOfCharIds)
		analyselink := GenerateAnalyserLink(value.ReportID, value.FightID, value.CharacterName, mapOfCharIds)

		improvements = append(improvements, BossImprovement{
			value.EncounterName,
			value.Difficulty,
			value.Total,
			value.Rank,
			value.Percentile,
			comparelink,
			analyselink,
		})
	}

	return improvements, e
}

func GenerateCompareLink(ReportID string, FightID int, Name string, mapOfCharIds map[string]int) string {
	fightId := strconv.Itoa(FightID)
	url := "https://www.warcraftlogs.com/reports/"+ReportID+"#fight="+fightId+"&type=damage-done&comparesearchplayer="+GetCharId(ReportID, Name, mapOfCharIds)+"&comparesearch=2.10.2.28"

	return url
}

func GenerateAnalyserLink(ReportID string, FightID int, Name string, mapOfCharIds map[string]int) string {
	fightId := strconv.Itoa(FightID)
	url := "https://wowanalyzer.com/report/"+ReportID+"/"+fightId+"/"+GetCharId(ReportID, Name, mapOfCharIds)

	return url

}

func GetCharId(ReportID string, Name string, ints map[string]int) string {
	if ints[ReportID] != 0 {
		log.Debug("Found id from existing search")
		return strconv.Itoa(ints[ReportID])
	} else {
		log.Debug("Did not find char id, have to ask for reports on warcraftlogs")
		reports, e := WarcraftLogs.GetWarcraftLogsReport(ReportID)
		if e != nil{
			return ""
		}
		for _, value := range reports.Friendlies {
			if value.Name == Name{
				ints[ReportID] = value.Id
				return strconv.Itoa(value.Id)
			}
		}
	}
	return ""
}


func MakeSimBotUrls(id int) RaidBotSimulations {
	charMap, e := Redis.GetStruct("MAIN:"+strconv.Itoa(id))
	if e != nil{
		log.Error(e, " -> It seems there is no main registered to the requesting user")
		return RaidBotSimulations{}
	}
	char := BlizzardOauthAPI.CharacterMinimalFromMap(charMap)

	rest := "?region="+char.Region+"&realm="+char.Realm+"&name="+char.Name
	return RaidBotSimulations{
		GearSim: RaidBotUrl+"gear"+rest,
		TalentSim: RaidBotUrl+"talent"+rest,
		QuickSim: RaidBotUrl+"quick"+rest,
		StatSim: RaidBotUrl+"stat"+rest,
	}
}
