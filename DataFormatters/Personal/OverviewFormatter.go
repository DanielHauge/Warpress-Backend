package Personal

import (
	"../../Integrations/BlizzardOauthAPI"
	"../../Integrations/BlizzardOpenAPI"
	"../../Integrations/Raider.io"
	"../../Integrations/WarcraftLogs"
	"../../Integrations/Wowprogress"
	"../../Redis"
	log "../../Utility/Logrus"
	"github.com/avelino/slugify"
	"github.com/jinzhu/copier"
	"strconv"
	"strings"
	"sync"
)

func FetchFullPersonal(id int, profile *interface{}) error {

	var Profile Overview

	charMap, e := Redis.GetStruct("MAIN:" + strconv.Itoa(id))
	if e != nil {
		log.WithLocation().WithError(e).WithField("User", id).Error("There is no main registered to the user!")
		return e
	}
	char := BlizzardOauthAPI.CharacterMinimalFromMap(charMap)
	blizzard, raiderio, logs, wowprog, e := fetchAll(id, char)
	fillUpPersonal(blizzard, raiderio, logs, wowprog, &Profile)

	copier.Copy(profile, Profile)

	return e
}

func fillUpPersonal(blizzard BlizzardOpenAPI.FullCharInfo, raider Raider_io.CharacterProfile, encounters WarcraftLogs.Encounters, rank Wowprogress.GuildRank, profile *Overview) {
	profile.Character.Realm = blizzard.Realm
	profile.Character.Name = blizzard.Name
	profile.Character.Class = blizzard.Class
	profile.Character.Race = blizzard.Race
	profile.Character.Gender = blizzard.Gender
	profile.Character.Level = blizzard.Level
	profile.Character.Avatar = blizzard.Thumbnail
	profile.Character.Main = strings.Replace(blizzard.Thumbnail, "-avatar.", "-main.", 1)
	profile.Character.Faction = blizzard.Faction
	profile.Character.Spec = findAndFormatActiveSpec(blizzard)
	profile.Character.Items, profile.Character.Stats = formatItems(blizzard.Items)
	profile.Guild.Name = blizzard.Guild.Name
	profile.Guild.Realm = blizzard.Guild.Realm
	profile.Guild.Members = blizzard.Guild.Members
	profile.Guild.Emblem = formatEmblemFromGuild(blizzard.Guild.Emblem)
	profile.Guild.GuildRank = guildRank{rank.Score, rank.WorldRank, rank.AreaRank, rank.RealmRank}
	profile.BestParses = formatEncounters(encounters.Encounters)
	profile.RaiderIOProfile.ProfileUrl = raider.ProfileUrl
	profile.RaiderIOProfile.MythicPlusRanks = formatMythicRanks(raider.MythicPlusRanks)
	profile.RaiderIOProfile.MythicPlusHighestLevelRuns = formatRuns(raider.MythicPlusHighestLevelRuns)
	profile.RaiderIOProfile.MythicPlusRecentRuns = formatRuns(raider.MythicPlusRecentRuns)
	profile.RaiderIOProfile.MythicPlusWeeklyHighestLevelRuns = formatRuns(raider.MythicPlusWeeklyHighestLevelRuns)
}
func formatRuns(runs []Raider_io.DungeonRun) []dungeonRun {
	var result []dungeonRun
	for _, v := range runs {
		affixes := formatAffixes(v.Affixes)
		result = append(result, dungeonRun{
			v.Dungeon,
			v.ShortName,
			v.MythicLevel,
			v.CompletedAt,
			v.ClearTimeMs,
			v.NumKeystoneUpgrade,
			v.Score,
			affixes,
			v.Url,
		})
	}
	return result
}
func formatAffixes(affixes []Raider_io.Affix) []affix {
	var afix []affix
	for _, v := range affixes {
		afix = append(afix, affix{v.Id, v.Name, v.Description, v.WowheadUrl})
	}
	return afix
}
func formatMythicRanks(ranks Raider_io.MythicRanks) mythicRanks {
	return mythicRanks{
		formatRank(ranks.Overall),
		formatRank(ranks.Dps),
		formatRank(ranks.Healer),
		formatRank(ranks.Tank),
		formatRank(ranks.Class),
		formatRank(ranks.ClassDps),
		formatRank(ranks.ClassHealer),
		formatRank(ranks.ClassTank),
	}
}
func formatRank(input Raider_io.Rank) rank {
	return rank{input.World, input.Region, input.World}
}
func formatEncounters(encounters []WarcraftLogs.Encounter) []encounter {
	var enc []encounter
	for _, v := range encounters {
		enc = append(enc, encounter{
			v.EncounterID,
			v.EncounterName,
			v.Class,
			v.Spec,
			v.Rank,
			v.OutOf,
			v.Duration,
			"https://www.warcraftlogs.com/reports/" + v.ReportID,
			v.Difficulty,
			v.CharacterName,
			v.Server,
			v.Percentile,
			v.ItemLevelKeyOrPath,
			v.Total,
		})
	}
	return enc
}
func formatEmblemFromGuild(input BlizzardOpenAPI.Emblem) emblem {
	return emblem{
		Icon:              input.Icon,
		IconColor:         input.IconColor,
		IconColorId:       input.IconColorId,
		Border:            input.Border,
		BorderColor:       input.BorderColor,
		BorderColorId:     input.BorderColorId,
		BackgroundColor:   input.BackgroundColor,
		BackgroundColorId: input.BorderColorId,
	}
}
func formatItems(input BlizzardOpenAPI.Items) (items, []stat) {
	var AllStats [][]stat
	head, stat := formatItem(input.Head)
	AllStats = append(AllStats, stat)
	neck, stat := formatItem(input.Neck)
	AllStats = append(AllStats, stat)
	shoulder, stat := formatItem(input.Shoulder)
	AllStats = append(AllStats, stat)
	back, stat := formatItem(input.Back)
	AllStats = append(AllStats, stat)
	check, stat := formatItem(input.Chest)
	AllStats = append(AllStats, stat)
	wrist, stat := formatItem(input.Wrist)
	AllStats = append(AllStats, stat)
	hands, stat := formatItem(input.Hands)
	AllStats = append(AllStats, stat)
	waist, stat := formatItem(input.Waist)
	AllStats = append(AllStats, stat)
	legs, stat := formatItem(input.Legs)
	AllStats = append(AllStats, stat)
	feet, stat := formatItem(input.Feet)
	AllStats = append(AllStats, stat)
	finger1, stat := formatItem(input.Finger1)
	AllStats = append(AllStats, stat)
	finger2, stat := formatItem(input.Finger2)
	AllStats = append(AllStats, stat)
	trinket1, stat := formatItem(input.Trinket1)
	AllStats = append(AllStats, stat)
	trinket2, stat := formatItem(input.Trinket2)
	AllStats = append(AllStats, stat)
	main, stat := formatItem(input.MainHand)
	AllStats = append(AllStats, stat)
	off, stat := formatItem(input.OffHand)
	AllStats = append(AllStats, stat)
	i := items{input.AverageItemLevel, input.AverItemLevelEquipped, head, neck, shoulder, back, check, wrist, hands, waist, legs, feet, finger1, finger2, trinket1, trinket2, main, off}
	return i, aggregateStats(AllStats)

}
func aggregateStats(stats [][]stat) []stat {
	st := map[int]int{}
	for _, e := range stats {
		for _, d := range e {
			st[d.Stat] = st[d.Stat] + d.Amount
		}
	}
	var result []stat
	for statid, amount := range st {
		result = append(result, stat{statid, amount})
	}
	return result
}
func formatItem(input BlizzardOpenAPI.Item) (item, []stat) {

	result := item{
		Id:          input.Id,
		Name:        input.Name,
		Icon:        input.Icon,
		Quality:     input.Quality,
		ItemLevel:   input.ItemLevel,
		BonusLists:  input.BonusLists,
		Appearance:  input.Appearance.ItemAppearanceModId,
		AzeriteItem: azeriteItem{AzeriteLevel: input.AzeriteItem.AzeriteLevel, AzeriteExperience: input.AzeriteItem.AzeriteExperience, AzeriteExperienceRemaining: input.AzeriteItem.AzeriteExperienceRemaining},
		Gem0:        input.ToolTipParams.Gem0,
		Enchant:     input.ToolTipParams.Enchant,
	}
	for _, e := range input.AzeriteEmpoweredItem.AzeritePowers {
		result.AzeriteEmpoweredItem = append(result.AzeriteEmpoweredItem, azeritePower{e.Id, e.Tier, e.SpellId, e.BonusListId})
	}
	var stats []stat
	for _, s := range input.Stats {
		stats = append(stats, stat{s.Stat, s.Amount})
	}
	stats = append(stats, stat{-1, input.Armor})
	return result, stats
}
func findAndFormatActiveSpec(info BlizzardOpenAPI.FullCharInfo) specialization {
	for _, e := range info.Talents {
		if e.Selected {
			result := specialization{
				Name:            e.Spec.Name,
				Role:            e.Spec.Role,
				BackgroundImage: e.Spec.BackgroundImage,
				Icon:            e.Spec.Icon,
				Description:     e.Spec.Description,
				Order:           e.Spec.Order,
			}
			for _, e := range e.Talents {
				result.Talents = append(result.Talents, talentTier{e.Tier, e.Column, spell{e.Spell.Id, e.Spell.Name, e.Spell.Icon, e.Spell.Description, e.Spell.CastTime, e.Spell.Cooldown}})
			}
			return result
		}
	}
	return specialization{}
}
func fetchAll(id int, char BlizzardOauthAPI.CharacterMinimal) (BlizzardOpenAPI.FullCharInfo, Raider_io.CharacterProfile, WarcraftLogs.Encounters, Wowprogress.GuildRank, error) {

	var wg sync.WaitGroup
	var blizzwait sync.WaitGroup
	blizzwait.Add(1)
	wg.Add(4)
	var e error
	var blizzChar BlizzardOpenAPI.FullCharInfo
	go func() {
		blizzChar, e = BlizzardOpenAPI.GetBlizzardChar(char.Realm, char.Name, char.Region)
		go Redis.Set("GUILD:"+strconv.Itoa(id), blizzChar.Guild.Name+":"+blizzChar.Guild.Realm+":"+char.Region)
		wg.Done()
		blizzwait.Done()
	}()

	var raiderio Raider_io.CharacterProfile
	go func() {
		raiderio, e = Raider_io.GetRaiderIORank(Raider_io.CharInput{Name: char.Name, Realm: char.Realm, Region: char.Region})
		wg.Done()
	}()

	var logs WarcraftLogs.Encounters
	go func() {
		logs, e = WarcraftLogs.GetWarcraftLogsRanks(WarcraftLogs.CharInput{Name: char.Name, Realm: char.Realm, Region: char.Region})
		wg.Done()
	}()

	var wowprog Wowprogress.GuildRank
	go func() {
		blizzwait.Wait()
		wowprog, e = Wowprogress.GetGuildRank(Wowprogress.Input{Region: char.Region, Realm: slugify.Slugify(char.Realm), Guild: blizzChar.Guild.Name})
		wg.Done()
	}()

	wg.Wait()

	return blizzChar, raiderio, logs, wowprog, e

}


func FetchRaiderioPersonal(id int, Profile *interface{}) error {

	charMap, e := Redis.GetStruct("MAIN:" + strconv.Itoa(id))
	if e != nil {
		log.WithLocation().WithError(e).WithField("User", id).Error("There is no main registered to the user!")
		return e
	}
	char := BlizzardOauthAPI.CharacterMinimalFromMap(charMap)
	prof, e := Raider_io.GetRaiderIORank(Raider_io.CharInput{Name: char.Name, Realm: char.Realm, Region: char.Region})
	copier.Copy(Profile, &prof)
	return e
}

func FetchWarcraftlogsPersonal(id int, Logs *interface{}) error {
	charMap, e := Redis.GetStruct("MAIN:" + strconv.Itoa(id))
	if e != nil {
		log.WithLocation().WithError(e).WithField("User", id).Error("There is no main registered to the user!")
		return e
	}
	char := BlizzardOauthAPI.CharacterMinimalFromMap(charMap)

	logs, e := WarcraftLogs.GetWarcraftLogsRanks(WarcraftLogs.CharInput{Name: char.Name, Realm: char.Realm, Region: char.Region})
	copier.Copy(Logs, &logs)
	return e
}

func FetchBlizzardPersonal(id int, Profile *interface{}) error {

	charMap, e := Redis.GetStruct("MAIN:" + strconv.Itoa(id))
	if e != nil {
		log.WithLocation().WithError(e).WithField("User", id).Error("There is no main registered to the user!")
		return e
	}
	char := BlizzardOauthAPI.CharacterMinimalFromMap(charMap)
	blizzChar, e := BlizzardOpenAPI.GetBlizzardChar(char.Realm, char.Name, char.Region)
	go Redis.Set("GUILD:"+strconv.Itoa(id), blizzChar.Guild.Name+":"+blizzChar.Guild.Realm+":"+char.Region)
	copier.Copy(Profile, &blizzChar)
	return e
}
