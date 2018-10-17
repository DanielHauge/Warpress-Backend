package BlizzardOpenAPI

type GuildWithMembers struct {
	GuildName string        `json:"name"`
	Realm     string        `json:"realm"`
	Members   []GuildMember `json:"members"`
}

type GuildMember struct {
	Character Character `json:"character"`
	Rank      int       `json:"rank"`
}

type Character struct {
	Name  string `json:"name"`
	Class int    `json:"class"`
	Spec  Spec   `json:"spec"`
}

type Spec struct {
	Name string `json:"name"`
	Role string `json:"role"`
}
