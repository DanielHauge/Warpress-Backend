package BlizzardOauthAPI


type CharacterMinimal struct {
	Name string `json:"name"`
	Realm string `json:"realm"`
	Locale string `json:"locale"`
}

func (char *CharacterMinimal) ToMap() map[string]interface{}{
	return map[string]interface{}{
		"name": char.Name,
		"realm": char.Realm,
		"locale": char.Locale,
	}
}

func CharacterMinimalFromMap(m map[string]string) CharacterMinimal{
	return CharacterMinimal{Name:m["name"], Realm:m["realm"], Locale:m["locale"]}
}
