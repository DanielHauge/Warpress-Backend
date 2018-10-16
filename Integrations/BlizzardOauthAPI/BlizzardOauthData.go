package BlizzardOauthAPI


type CharacterMinimal struct {
	Name   string `json:"name"`
	Realm  string `json:"realm"`
	Region string  `json:"region"`
}

func (char *CharacterMinimal) ToMap() map[string]interface{}{
	return map[string]interface{}{
		"name": char.Name,
		"realm": char.Realm,
		"region": char.Region,
	}
}

func CharacterMinimalFromMap(m map[string]string) CharacterMinimal{
	return CharacterMinimal{Name:m["name"], Realm:m["realm"], Region:m["region"]}
}
