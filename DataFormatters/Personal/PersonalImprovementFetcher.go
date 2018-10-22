package Personal

import (
	"../../Integrations/BlizzardOauthAPI"
	"../../Integrations/WarcraftLogs"
	"../../Redis"
	log "../../Utility/Logrus"
	"github.com/jinzhu/copier"
	"strconv"
)

var RaidBotUrl = "https://www.raidbots.com/simbot/"

func FetchPersonalImprovementsFull(id int, improvements *interface{}) error {

	var persImprov PersonalImprovement

	persImprov.SimulationURLS = MakeSimBotUrls(id)
	bossimprovements, e := GenerateWarcraftLogs(id)
	persImprov.BossImprovements = bossimprovements
	copier.Copy(improvements, persImprov)
	return e
}

func GenerateWarcraftLogs(id int) ([]BossImprovement, error) {
	var logs []WarcraftLogs.Encounter
	var interfaceLogs interface{}
	var improvements []BossImprovement
	e := FetchWarcraftlogsPersonal(id, &interfaceLogs)
	if e != nil {
		log.WithLocation().WithError(e).Error("Hov!")
	}
	copier.Copy(logs, interfaceLogs)

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
	url := "https://www.warcraftlogs.com/reports/" + ReportID + "#fight=" + fightId + "&type=damage-done&comparesearchplayer=" + GetCharId(ReportID, Name, mapOfCharIds) + "&comparesearch=2.10.2.28"

	return url
}

func GenerateAnalyserLink(ReportID string, FightID int, Name string, mapOfCharIds map[string]int) string {
	fightId := strconv.Itoa(FightID)
	url := "https://wowanalyzer.com/report/" + ReportID + "/" + fightId + "/" + GetCharId(ReportID, Name, mapOfCharIds)

	return url
}

func GetCharId(ReportID string, Name string, ints map[string]int) string {
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

func MakeSimBotUrls(id int) RaidBotSimulations {
	charMap, e := Redis.GetStruct("MAIN:" + strconv.Itoa(id))
	if e != nil {
		log.WithLocation().WithError(e).WithField("User", id).Error("There is no main registered to the user!")
		return RaidBotSimulations{}
	}
	char := BlizzardOauthAPI.CharacterMinimalFromMap(charMap)

	rest := "?region=" + char.Region + "&realm=" + char.Realm + "&name=" + char.Name
	return RaidBotSimulations{
		GearSim:   RaidBotUrl + "gear" + rest,
		TalentSim: RaidBotUrl + "talent" + rest,
		QuickSim:  RaidBotUrl + "quick" + rest,
		StatSim:   RaidBotUrl + "stat" + rest,
	}
}
