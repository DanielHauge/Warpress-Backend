package bnet


type ByLevel []WOWCharacter

func (c ByLevel) Len() int 				{ return len(c) }
func (c ByLevel) Less(i, j int) bool 	{ return c[i].Level > c[j].Level }
func (c ByLevel) Swap(i, j int) 		{ c[i], c[j] = c[j], c[i] }

type WOWCharacter struct {
	Name				string		`json:"name"`
	Realm 				string		`json:"realm"`
	Class 				int			`json:"class"`
	Race 				int			`json:"race"`
	Gender 				int			`json:"gender"`
	Level				int			`json:"level"`
	Thumbnail			string		`json:"thumbnail"`
	Spec 				WowSpec		`json:"spec"`
	Guild 				string		`json:"guild"`
}

type WowSpec struct {
	Name 				string		`json:"name"`
	Role 				string		`json:"role"`
	Icon 				string		`json:"icon"`
}


