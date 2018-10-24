package Personal

import (
	"../../Integrations/WarcraftLogs"
	"../../Postgres"
	log "../../Utility/Logrus"
	"github.com/jinzhu/copier"
	"strconv"
)

var RaidBotUrl = "https://www.raidbots.com/simbot/"

func FetchPersonalImprovementsFull(id int, improvements *interface{}) error {

	var persImprov Improvements
	persImprov.SimulationURLS = makeSimBotUrls(id)
	bossimprovements, e := generateWarcraftLogs(id)
	persImprov.BossImprovements = bossimprovements
	copier.Copy(improvements, persImprov)
	return e
}

func generateWarcraftLogs(id int) ([]BossImprovement, error) {
	var logs WarcraftLogs.Encounters
	var interfaceLogs interface{}
	var improvements []BossImprovement
	e := FetchWarcraftlogsPersonal(id, &interfaceLogs)
	if e != nil {
		log.WithLocation().WithError(e).Error("Hov!")
	}
	copier.Copy(logs, interfaceLogs)

	mapOfCharIds := map[string]int{}
	for _, value := range logs.Encounters {

		comparelink := generateCompareLink(value.ReportID, value.FightID, value.CharacterName, mapOfCharIds)
		analyselink := generateAnalyserLink(value.ReportID, value.FightID, value.CharacterName, mapOfCharIds)

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

func generateCompareLink(ReportID string, FightID int, Name string, mapOfCharIds map[string]int) string {
	fightId := strconv.Itoa(FightID)
	url := "https://www.warcraftlogs.com/reports/" + ReportID + "#fight=" + fightId + "&type=damage-done&comparesearchplayer=" + getCharId(ReportID, Name, mapOfCharIds) + "&comparesearch=2.10.2.28"

	return url
}

func generateAnalyserLink(ReportID string, FightID int, Name string, mapOfCharIds map[string]int) string {
	fightId := strconv.Itoa(FightID)
	url := "https://wowanalyzer.com/report/" + ReportID + "/" + fightId + "/" + getCharId(ReportID, Name, mapOfCharIds)

	return url
}

func getCharId(ReportID string, Name string, ints map[string]int) string {
	if ints[ReportID] != 0 {
		return strconv.Itoa(ints[ReportID])
	} else {
		reports, e := WarcraftLogs.GetWarcraftLogsReport(ReportID)
		if e != nil {
			log.WithLocation().WithError(e).Error("Hov!")
			return ""
		}
		for _, value := range reports.Friendlies {
			if value.Name == Name {
				ints[ReportID] = value.Id
				return strconv.Itoa(value.Id)
			}
		}
	}
	return ""
}

func makeSimBotUrls(id int) RaidBotSimulations {
	name, realm, region, e := Postgres.GetMain(id)
	if e != nil {
		log.WithLocation().WithError(e).WithField("User", id).Error("There is no main registered to the user!")
		return RaidBotSimulations{}
	}

	rest := "?region=" + region + "&realm=" + realm + "&name=" + name
	return RaidBotSimulations{
		GearSim:   RaidBotUrl + "gear" + rest,
		TalentSim: RaidBotUrl + "talent" + rest,
		QuickSim:  RaidBotUrl + "quick" + rest,
		StatSim:   RaidBotUrl + "stat" + rest,
	}
}
