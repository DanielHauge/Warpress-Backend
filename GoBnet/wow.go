package bnet


type ByLevel []WOWCharacter

func (c ByLevel) Len() int 				{ return len(c) }
func (c ByLevel) Less(i, j int) bool 	{ return c[i].Level > c[j].Level }
func (c ByLevel) Swap(i, j int) 		{ c[i], c[j] = c[j], c[i] }

type WOWCharacter struct {
	Name				string		`json:"name"`
	Realm 				string		`json:"realm"`
	BattleGroup 		string		`json:"battlegroup"`
	Class 				int			`json:"class"`
	Race 				int			`json:"race"`
	Gender 				int			`json:"gender"`
	Level				int			`json:"level"`
	AchievementPoints	int			`json:"achievementPoints"`
	Thumbnail			string		`json:"thumbnail"`
	Spec 				WowSpec		`json:"spec"`
	Guild 				string		`json:"guild"`
	GuildRealm 			string		`json:"guildRealm"`
	LastModified 		int			`json:"last_modified"`
}

type WowSpec struct {
	Name 				string		`json:"name"`
	Role 				string		`json:"role"`
	BackgroundImage 	string		`json:"backgroundImage"`
	Icon 				string		`json:"icon"`
	Description 		string		`json:"description"`
	Order 				int 		`json:"order"`
}


