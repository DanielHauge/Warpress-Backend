package DataModel

type Addon struct {
	Name string `json:"name"`
	TwitchLink string `json:"twitch_link"`
	GuildId int `json:"guild_id"`
	Id int `json:"id"`
}

type Weakaura struct {
	Name string `json:"name"`
	Link string `json:"link"`
	Import []byte `json:"import"`
	GuildId int `json:"guild_id"`
	Id int `json:"id"`
}