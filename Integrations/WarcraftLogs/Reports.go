package WarcraftLogs

type GuildReports struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Owner string `json:"owner"`
	Start int64  `json:"start"`
	End   int64  `json:"end"`
	Zone  int    `json:"zone"`
}
