package Personal

import (
	"../../Integrations/BlizzardOpenAPI"
	"../../Integrations/Raider.io"
	"../../Integrations/WarcraftLogs"
	"../../Integrations/Wowprogress"
	Postgres "../../Postgres/PreparedProcedures"
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

	name, realm, region, e := Postgres.GetMain(id)
	if e != nil {
		log.WithLocation().WithError(e).WithField("User", id).Error("There is no main registered to the user!")
		return e
	}

	blizzard, raiderio, logs, wowprog, e := fetchAll(id, name, realm, region)
	fillUpPersonal(blizzard, raiderio, logs, wowprog, &Profile)

	copier.Copy(profile, Profile)

	return e
}

func FetchFullInspect(name string, realm string, region string, profile *interface{}) error {
	var Profile Overview
	blizzard, raiderio, logs, wowprog, e := fetchAll(0, name, realm, region)
	fillUpPersonal(blizzard, raiderio, logs, wowprog, &Profile)

	copier.Copy(profile, Profile)
	return e
}

func fillUpPersonal(blizzard BlizzardOpenAPI.FullCharInfo, raider Raider_io.CharacterProfile, encounters WarcraftLogs.Encounters, rank Wowprogress.GuildRank, profile *Overview) {
	profile.Character.Realm = blizzard.Realm
	profile.Character.SluggedRealm = slugify.Slugify(blizzard.Realm)
	profile.Character.Name = blizzard.Name
	profile.Character.Class = blizzard.Class
	profile.Character.Race = blizzard.Race
	profile.Character.Gender = blizzard.Gender
	profile.Character.Level = blizzard.Level
	profile.Character.Avatar = blizzard.Thumbnail
	profile.Character.Main = strings.Replace(blizzard.Thumbnail, "-avatar.", "-main.", 1)
	profile.Character.Faction = blizzard.Faction
	profile.Character.Spec = findAndFormatActiveSpec(blizzard)
	profile.Character.Spec.MasterySpellID = getMasteryID(profile.Character.Class, profile.Character.Spec.Name)
	profile.Character.Items = formatItems(blizzard.Items)
	profile.Character.Stats = formatStats(blizzard.Stats)
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
func getMasteryID(class int, spec string) int {

	result := -1

	switch class {
	case 1: // Warrior
		switch spec {
		case "Arms":
			result = 76838
		case "Fury":
			return 76856
		case "Protection":
			return 76857
		}

	case 2: // Paladin

		switch spec {
		case "Retribution":
			result = 267316
		case "Protection":
			result = 76671
		case "Holy":
			result = 183997
		}

	case 3: // Hunter
		switch spec {
		case "Beast Mastery":
			return 76657
		case "Marksmanship":
			return 193468
		case "Survival":
			return 191334
		}
	case 4: // Rogue
		switch spec {
		case "Assassination":
			result = 76803
		case "Outlaw":
			result = 76806
		case "Subtlety":
			result = 76808
		}
	case 5: // Priest
		switch spec {
		case "Discipline":
			result = 271534
		case "Holy":
			result = 77485
		case "Shadow":
			result = 77486
		}
	case 6: // Death Knight
		switch spec {
		case "Frost":
			result = 77514
		case "Blood":
			return 77513
		case "Unholy":
			return 77515
		}
	case 7: // Shaman
		switch spec {
		case "Elemental":
			result = 168534
		case "Enhancement":
			result = 77223
		case "Restoration":
			result = 77226
		}
	case 8: // Mage
		switch spec {
		case "Arcane":
			result = 190740
		case "Fire":
			result = 12846
		case "Frost":
			result = 76613
		}
	case 9: // Warlock
		switch spec {
		case "Affliction":
			result = 77215
		case "Demonology":
			result = 77219
		case "Destruction":
			result = 77220
		}
	case 10: // Monk
		switch spec {
		case "Brewmaster":
			result = 117906
		case "Mistweaver":
			result = 117907
		case "Windwalker":
			result = 115636
		}
	case 11: // Druid
		switch spec {
		case "Balance":
			return 77492
		case "Feral":
			return 77493
		case "Guardian":
			return 155783
		case "Restoration":
			return 77495
		}

	case 12: // Demon Hunter
		switch spec {
		case "Havoc":
			result = 185164
		case "Vengeance":
			return 203747
		}

	default:
		return 1

	}

	return result
}
func formatStats(stats BlizzardOpenAPI.Stats) Stats {

	return Stats{
		stats.Health,
		stats.PowerType,
		stats.Power,
		stats.Str,
		stats.Agi,
		stats.Int,
		stats.Sta,
		stats.SpeedRating,
		stats.SpeedRatingBonus,
		stats.Crit,
		stats.CritRating,
		stats.Haste,
		stats.HasteRating,
		stats.HasteRatingPercent,
		stats.Mastery,
		stats.MasteryRating,
		stats.Leech,
		stats.LeechRating,
		stats.LeechRatingBonus,
		stats.Versatility,
		stats.VersatilityDamageDoneBonus,
		stats.VersatilityHealingDoneBonus,
		stats.VersatilityDamageTakenBonus,
		stats.AvoidanceRating,
		stats.AvoidanceRatingBonus,
		stats.SpellPen,
		stats.SpellCrit,
		stats.SpellCritRating,
		stats.Armor,
		stats.Dodge,
		stats.DodgeRating,
		stats.MainHandDmgMin,
		stats.MainHandDmgMax,
		stats.MainHandSpeed,
		stats.MainHandDps,
		stats.OffHandDmgMin,
		stats.OffHandDmgMax,
		stats.OffHandSpeed,
		stats.OffHandDps,
		stats.RangedDmgMin,
		stats.RangedDmgMax,
		stats.RangedSpeed,
		stats.RangedDps,
	}

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
	return rank{input.World, input.Region, input.Realm}
}
func formatEncounters(encounters []WarcraftLogs.Encounter) []dificultyParse {
	var mythic dificultyParse
	mythic.Difficulty = "Mythic"
	var heroic dificultyParse
	heroic.Difficulty = "Heroic"
	var normal dificultyParse
	normal.Difficulty = "Normal"
	mythicSpecs := map[string]bool{}
	heroicSpecs := map[string]bool{}
	normalSpecs := map[string]bool{}

	for _, v := range encounters {
		if v.Difficulty == 5 {
			mythicSpecs[v.Spec] = true
			mythic.Logs = append(mythic.Logs, encounter{
				v.EncounterID,
				v.EncounterName,
				v.Spec,
				v.Rank,
				v.OutOf,
				v.Duration,
				"https://www.warcraftlogs.com/reports/" + v.ReportID+"#fight="+strconv.Itoa(v.FightID),
				v.CharacterName,
				v.Percentile,
				v.ItemLevelKeyOrPath,
				v.Total,
			})
		} else if v.Difficulty == 4 {
			heroicSpecs[v.Spec] = true
			heroic.Logs = append(heroic.Logs, encounter{
				v.EncounterID,
				v.EncounterName,
				v.Spec,
				v.Rank,
				v.OutOf,
				v.Duration,
				"https://www.warcraftlogs.com/reports/" + v.ReportID+"#fight="+strconv.Itoa(v.FightID),
				v.CharacterName,
				v.Percentile,
				v.ItemLevelKeyOrPath,
				v.Total,
			})
		} else if v.Difficulty == 3 {
			normalSpecs[v.Spec] = true
			normal.Logs = append(normal.Logs, encounter{
				v.EncounterID,
				v.EncounterName,
				v.Spec,
				v.Rank,
				v.OutOf,
				v.Duration,
				"https://www.warcraftlogs.com/reports/" + v.ReportID+"#fight="+strconv.Itoa(v.FightID),
				v.CharacterName,
				v.Percentile,
				v.ItemLevelKeyOrPath,
				v.Total,
			})
		}
	}

	for i := range mythicSpecs {
		mythic.Specs = append(mythic.Specs, i)
	}
	for i := range heroicSpecs {
		heroic.Specs = append(heroic.Specs, i)
	}
	for i := range normalSpecs {
		normal.Specs = append(normal.Specs, i)
	}

	var result []dificultyParse

	if len(mythic.Logs) != 0 {
		result = append(result, mythic)
	}
	if len(heroic.Logs) != 0 {
		result = append(result, heroic)
	}
	if len(normal.Logs) != 0 {
		result = append(result, normal)
	}

	return result
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
func formatItems(input BlizzardOpenAPI.Items) items {
	head := formatItem(input.Head)
	neck := formatItem(input.Neck)
	shoulder := formatItem(input.Shoulder)
	back := formatItem(input.Back)
	check := formatItem(input.Chest)
	wrist := formatItem(input.Wrist)
	hands := formatItem(input.Hands)
	waist := formatItem(input.Waist)
	legs := formatItem(input.Legs)
	feet := formatItem(input.Feet)
	finger1 := formatItem(input.Finger1)
	finger2 := formatItem(input.Finger2)
	trinket1 := formatItem(input.Trinket1)
	trinket2 := formatItem(input.Trinket2)
	main := formatItem(input.MainHand)
	off := formatItem(input.OffHand)
	i := items{input.AverageItemLevel, input.AverItemLevelEquipped, head, neck, shoulder, back, check, wrist, hands, waist, legs, feet, finger1, finger2, trinket1, trinket2, main, off}
	return i

}

func formatItem(input BlizzardOpenAPI.Item) item {

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

	return result
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
func fetchAll(id int, name string, realm string, region string) (BlizzardOpenAPI.FullCharInfo, Raider_io.CharacterProfile, WarcraftLogs.Encounters, Wowprogress.GuildRank, error) {

	var wg sync.WaitGroup
	var blizzwait sync.WaitGroup
	blizzwait.Add(1)
	wg.Add(4)
	var e error
	var blizzChar BlizzardOpenAPI.FullCharInfo

	go func() {
		blizzChar, e = BlizzardOpenAPI.GetBlizzardChar(realm, name, region)
		if id != 0 {
			go Redis.Set("GUILD:"+strconv.Itoa(id), blizzChar.Guild.Name+":"+slugify.Slugify(blizzChar.Guild.Realm)+":"+region)
		}
		wg.Done()
		blizzwait.Done()
	}()

	var raiderio Raider_io.CharacterProfile
	go func() {
		raiderio, e = Raider_io.GetRaiderIORank(name, realm, region)
		wg.Done()
	}()

	var logs WarcraftLogs.Encounters
	go func() {
		logs, e = WarcraftLogs.GetWarcraftLogsRanks(WarcraftLogs.CharInput{Name: name, Realm: realm, Region: region})
		wg.Done()
	}()

	var wowprog Wowprogress.GuildRank
	go func() {
		blizzwait.Wait()
		wowprog, e = Wowprogress.GetGuildRank(Wowprogress.Input{Region: region, Realm: slugify.Slugify(realm), Guild: blizzChar.Guild.Name})
		wg.Done()
	}()

	wg.Wait()

	return blizzChar, raiderio, logs, wowprog, e

}

func FetchRaiderioPersonal(id int, Profile *interface{}) error {

	name, realm, region, e := Postgres.GetMain(id)
	if e != nil {
		log.WithLocation().WithError(e).WithField("User", id).Error("There is no main registered to the user!")
		return e
	}
	prof, e := Raider_io.GetRaiderIORank(name, realm, region)
	copier.Copy(Profile, &prof)
	return e
}

func FetchWarcraftlogsPersonal(id int, Logs *interface{}) error {
	name, realm, region, e := Postgres.GetMain(id)
	if e != nil {
		log.WithLocation().WithError(e).WithField("User", id).Error("There is no main registered to the user!")
		return e
	}

	logs, e := WarcraftLogs.GetWarcraftLogsRanks(WarcraftLogs.CharInput{Name: name, Realm: realm, Region: region})
	copier.Copy(Logs, &logs)
	return e
}

func FetchBlizzardPersonal(id int, Profile *interface{}) error {

	name, realm, region, e := Postgres.GetMain(id)
	if e != nil {
		log.WithLocation().WithError(e).WithField("User", id).Error("There is no main registered to the user!")
		return e
	}

	blizzChar, e := BlizzardOpenAPI.GetBlizzardChar(realm, name, region)
	go Redis.Set("GUILD:"+strconv.Itoa(id), blizzChar.Guild.Name+":"+slugify.Slugify(blizzChar.Guild.Realm)+":"+region)
	copier.Copy(Profile, &blizzChar)
	return e
}
