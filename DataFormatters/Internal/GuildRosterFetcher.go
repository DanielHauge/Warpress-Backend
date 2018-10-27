package Internal

import (
	"../../Integrations/BlizzardOpenAPI"
	"../../Redis"
	"github.com/jinzhu/copier"
	"strconv"
	"strings"
)

func FetchGuildRooster(id int, Guild *interface{}) error {
	guildstring, e := Redis.Get("GUILD:" + strconv.Itoa(id))
	if e != nil {
		return e
	}
	split := strings.Split(guildstring, ":")
	guild := struct {
		Name   string
		Realm  string
		Region string
	}{Name: split[0], Realm: split[1], Region: split[2]}
	guildswithmembers, e := BlizzardOpenAPI.GetBlizzardGuildMembers(guild.Name, guild.Realm, guild.Region)

	copier.Copy(Guild, &guildswithmembers)
	return e
}
