package Personal

type Improvements struct {
	SimulationURLS   RaidBotSimulations `json:"simulation_urls"`
	BossImprovements []BossImprovement  `json:"boss_improvements"`
}

type RaidBotSimulations struct {
	GearSim   string `json:"gear_sim"`
	TalentSim string `json:"talent_sim"`
	QuickSim  string `json:"quick_sim"`
	StatSim   string `json:"stat_sim"`
}

type BossImprovement struct {
	BossName               string  `json:"boss_name"`
	Difficulty             int     `json:"difficulty"`
	Amount                 float64 `json:"amount"`
	Rank                   int     `json:"rank"`
	Percentile             int     `json:"percentile"`
	WarcraftLogsCompareUrl string  `json:"warcraft_logs_compare_url"`
	WowAnalyserUrl         string  `json:"wow_analyser_url"`
}
