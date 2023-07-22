package config

import "time"

type KillDeathResponse []Event

type Event struct {
	GroupMemberCount     int       `json:"groupMemberCount,omitempty"`
	NumberOfParticipants int       `json:"numberOfParticipants,omitempty"`
	EventID              int       `json:"EventId,omitempty"`
	TimeStamp            time.Time `json:"TimeStamp,omitempty"`
	Version              int       `json:"Version,omitempty"`
	Killer               struct {
		AverageItemPower float64 `json:"AverageItemPower,omitempty"`
		Equipment        struct {
			MainHand struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"MainHand,omitempty"`
			OffHand struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"OffHand,omitempty"`
			Head struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"Head,omitempty"`
			Armor struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"Armor,omitempty"`
			Shoes struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"Shoes,omitempty"`
			Bag struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"Bag,omitempty"`
			Cape struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"Cape,omitempty"`
			Mount struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"Mount,omitempty"`
			Potion struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"Potion,omitempty"`
			Food struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"Food,omitempty"`
		} `json:"Equipment,omitempty"`
		Inventory          []any   `json:"Inventory,omitempty"`
		Name               string  `json:"Name,omitempty"`
		ID                 string  `json:"Id,omitempty"`
		GuildName          string  `json:"GuildName,omitempty"`
		GuildID            string  `json:"GuildId,omitempty"`
		AllianceName       string  `json:"AllianceName,omitempty"`
		AllianceID         string  `json:"AllianceId,omitempty"`
		AllianceTag        string  `json:"AllianceTag,omitempty"`
		Avatar             string  `json:"Avatar,omitempty"`
		AvatarRing         string  `json:"AvatarRing,omitempty"`
		DeathFame          int     `json:"DeathFame,omitempty"`
		KillFame           int     `json:"KillFame,omitempty"`
		FameRatio          float64 `json:"FameRatio,omitempty"`
		LifetimeStatistics struct {
			PvE struct {
				Total            int `json:"Total,omitempty"`
				Royal            int `json:"Royal,omitempty"`
				Outlands         int `json:"Outlands,omitempty"`
				Avalon           int `json:"Avalon,omitempty"`
				Hellgate         int `json:"Hellgate,omitempty"`
				CorruptedDungeon int `json:"CorruptedDungeon,omitempty"`
				Mists            int `json:"Mists,omitempty"`
			} `json:"PvE,omitempty"`
			Gathering struct {
				Fiber struct {
					Total    int `json:"Total,omitempty"`
					Royal    int `json:"Royal,omitempty"`
					Outlands int `json:"Outlands,omitempty"`
					Avalon   int `json:"Avalon,omitempty"`
				} `json:"Fiber,omitempty"`
				Hide struct {
					Total    int `json:"Total,omitempty"`
					Royal    int `json:"Royal,omitempty"`
					Outlands int `json:"Outlands,omitempty"`
					Avalon   int `json:"Avalon,omitempty"`
				} `json:"Hide,omitempty"`
				Ore struct {
					Total    int `json:"Total,omitempty"`
					Royal    int `json:"Royal,omitempty"`
					Outlands int `json:"Outlands,omitempty"`
					Avalon   int `json:"Avalon,omitempty"`
				} `json:"Ore,omitempty"`
				Rock struct {
					Total    int `json:"Total,omitempty"`
					Royal    int `json:"Royal,omitempty"`
					Outlands int `json:"Outlands,omitempty"`
					Avalon   int `json:"Avalon,omitempty"`
				} `json:"Rock,omitempty"`
				Wood struct {
					Total    int `json:"Total,omitempty"`
					Royal    int `json:"Royal,omitempty"`
					Outlands int `json:"Outlands,omitempty"`
					Avalon   int `json:"Avalon,omitempty"`
				} `json:"Wood,omitempty"`
				All struct {
					Total    int `json:"Total,omitempty"`
					Royal    int `json:"Royal,omitempty"`
					Outlands int `json:"Outlands,omitempty"`
					Avalon   int `json:"Avalon,omitempty"`
				} `json:"All,omitempty"`
			} `json:"Gathering,omitempty"`
			Crafting struct {
				Total    int `json:"Total,omitempty"`
				Royal    int `json:"Royal,omitempty"`
				Outlands int `json:"Outlands,omitempty"`
				Avalon   int `json:"Avalon,omitempty"`
			} `json:"Crafting,omitempty"`
			CrystalLeague int `json:"CrystalLeague,omitempty"`
			FishingFame   int `json:"FishingFame,omitempty"`
			FarmingFame   int `json:"FarmingFame,omitempty"`
			Timestamp     any `json:"Timestamp,omitempty"`
		} `json:"LifetimeStatistics,omitempty"`
	} `json:"Killer,omitempty"`
	Victim struct {
		AverageItemPower float64 `json:"AverageItemPower,omitempty"`
		Equipment        struct {
			MainHand struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"MainHand,omitempty"`
			OffHand struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"OffHand,omitempty"`
			Head struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"Head,omitempty"`
			Armor struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"Armor,omitempty"`
			Shoes struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"Shoes,omitempty"`
			Bag struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"Bag,omitempty"`
			Cape struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"Cape,omitempty"`
			Mount struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"Mount,omitempty"`
			Potion struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"Potion,omitempty"`
			Food struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"Food,omitempty"`
		} `json:"Equipment,omitempty"`
		// Inventory          []any   `json:"Inventory,omitempty"`
		Inventory []struct {
			Type          string `json:"Type,omitempty"`
			Count         int    `json:"Count,omitempty"`
			Quality       int    `json:"Quality,omitempty"`
			ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
			PassiveSpells []any  `json:"PassiveSpells,omitempty"`
		} `json:"Inventory,omitempty"`
		Name               string  `json:"Name,omitempty"`
		ID                 string  `json:"Id,omitempty"`
		GuildName          string  `json:"GuildName,omitempty"`
		GuildID            string  `json:"GuildId,omitempty"`
		AllianceName       string  `json:"AllianceName,omitempty"`
		AllianceID         string  `json:"AllianceId,omitempty"`
		AllianceTag        string  `json:"AllianceTag,omitempty"`
		Avatar             string  `json:"Avatar,omitempty"`
		AvatarRing         string  `json:"AvatarRing,omitempty"`
		DeathFame          int     `json:"DeathFame,omitempty"`
		KillFame           int     `json:"KillFame,omitempty"`
		FameRatio          float64 `json:"FameRatio,omitempty"`
		LifetimeStatistics struct {
			PvE struct {
				Total            int `json:"Total,omitempty"`
				Royal            int `json:"Royal,omitempty"`
				Outlands         int `json:"Outlands,omitempty"`
				Avalon           int `json:"Avalon,omitempty"`
				Hellgate         int `json:"Hellgate,omitempty"`
				CorruptedDungeon int `json:"CorruptedDungeon,omitempty"`
				Mists            int `json:"Mists,omitempty"`
			} `json:"PvE,omitempty"`
			Gathering struct {
				Fiber struct {
					Total    int `json:"Total,omitempty"`
					Royal    int `json:"Royal,omitempty"`
					Outlands int `json:"Outlands,omitempty"`
					Avalon   int `json:"Avalon,omitempty"`
				} `json:"Fiber,omitempty"`
				Hide struct {
					Total    int `json:"Total,omitempty"`
					Royal    int `json:"Royal,omitempty"`
					Outlands int `json:"Outlands,omitempty"`
					Avalon   int `json:"Avalon,omitempty"`
				} `json:"Hide,omitempty"`
				Ore struct {
					Total    int `json:"Total,omitempty"`
					Royal    int `json:"Royal,omitempty"`
					Outlands int `json:"Outlands,omitempty"`
					Avalon   int `json:"Avalon,omitempty"`
				} `json:"Ore,omitempty"`
				Rock struct {
					Total    int `json:"Total,omitempty"`
					Royal    int `json:"Royal,omitempty"`
					Outlands int `json:"Outlands,omitempty"`
					Avalon   int `json:"Avalon,omitempty"`
				} `json:"Rock,omitempty"`
				Wood struct {
					Total    int `json:"Total,omitempty"`
					Royal    int `json:"Royal,omitempty"`
					Outlands int `json:"Outlands,omitempty"`
					Avalon   int `json:"Avalon,omitempty"`
				} `json:"Wood,omitempty"`
				All struct {
					Total    int `json:"Total,omitempty"`
					Royal    int `json:"Royal,omitempty"`
					Outlands int `json:"Outlands,omitempty"`
					Avalon   int `json:"Avalon,omitempty"`
				} `json:"All,omitempty"`
			} `json:"Gathering,omitempty"`
			Crafting struct {
				Total    int `json:"Total,omitempty"`
				Royal    int `json:"Royal,omitempty"`
				Outlands int `json:"Outlands,omitempty"`
				Avalon   int `json:"Avalon,omitempty"`
			} `json:"Crafting,omitempty"`
			CrystalLeague int `json:"CrystalLeague,omitempty"`
			FishingFame   int `json:"FishingFame,omitempty"`
			FarmingFame   int `json:"FarmingFame,omitempty"`
			Timestamp     any `json:"Timestamp,omitempty"`
		} `json:"LifetimeStatistics,omitempty"`
	} `json:"Victim,omitempty"`
	TotalVictimKillFame int `json:"TotalVictimKillFame,omitempty"`
	Location            any `json:"Location,omitempty"`
	Participants        []struct {
		AverageItemPower float64 `json:"AverageItemPower,omitempty"`
		Equipment        struct {
			MainHand struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"MainHand,omitempty"`
			OffHand any `json:"OffHand,omitempty"`
			Head    struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"Head,omitempty"`
			Armor struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"Armor,omitempty"`
			Shoes struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"Shoes,omitempty"`
			Bag struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"Bag,omitempty"`
			Cape struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"Cape,omitempty"`
			Mount struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"Mount,omitempty"`
			Potion struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"Potion,omitempty"`
			Food struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"Food,omitempty"`
		} `json:"Equipment,omitempty"`
		Inventory          []any   `json:"Inventory,omitempty"`
		Name               string  `json:"Name,omitempty"`
		ID                 string  `json:"Id,omitempty"`
		GuildName          string  `json:"GuildName,omitempty"`
		GuildID            string  `json:"GuildId,omitempty"`
		AllianceName       string  `json:"AllianceName,omitempty"`
		AllianceID         string  `json:"AllianceId,omitempty"`
		AllianceTag        string  `json:"AllianceTag,omitempty"`
		Avatar             string  `json:"Avatar,omitempty"`
		AvatarRing         string  `json:"AvatarRing,omitempty"`
		DeathFame          int     `json:"DeathFame,omitempty"`
		KillFame           int     `json:"KillFame,omitempty"`
		FameRatio          float64 `json:"FameRatio,omitempty"`
		LifetimeStatistics struct {
			PvE struct {
				Total            int `json:"Total,omitempty"`
				Royal            int `json:"Royal,omitempty"`
				Outlands         int `json:"Outlands,omitempty"`
				Avalon           int `json:"Avalon,omitempty"`
				Hellgate         int `json:"Hellgate,omitempty"`
				CorruptedDungeon int `json:"CorruptedDungeon,omitempty"`
				Mists            int `json:"Mists,omitempty"`
			} `json:"PvE,omitempty"`
			Gathering struct {
				Fiber struct {
					Total    int `json:"Total,omitempty"`
					Royal    int `json:"Royal,omitempty"`
					Outlands int `json:"Outlands,omitempty"`
					Avalon   int `json:"Avalon,omitempty"`
				} `json:"Fiber,omitempty"`
				Hide struct {
					Total    int `json:"Total,omitempty"`
					Royal    int `json:"Royal,omitempty"`
					Outlands int `json:"Outlands,omitempty"`
					Avalon   int `json:"Avalon,omitempty"`
				} `json:"Hide,omitempty"`
				Ore struct {
					Total    int `json:"Total,omitempty"`
					Royal    int `json:"Royal,omitempty"`
					Outlands int `json:"Outlands,omitempty"`
					Avalon   int `json:"Avalon,omitempty"`
				} `json:"Ore,omitempty"`
				Rock struct {
					Total    int `json:"Total,omitempty"`
					Royal    int `json:"Royal,omitempty"`
					Outlands int `json:"Outlands,omitempty"`
					Avalon   int `json:"Avalon,omitempty"`
				} `json:"Rock,omitempty"`
				Wood struct {
					Total    int `json:"Total,omitempty"`
					Royal    int `json:"Royal,omitempty"`
					Outlands int `json:"Outlands,omitempty"`
					Avalon   int `json:"Avalon,omitempty"`
				} `json:"Wood,omitempty"`
				All struct {
					Total    int `json:"Total,omitempty"`
					Royal    int `json:"Royal,omitempty"`
					Outlands int `json:"Outlands,omitempty"`
					Avalon   int `json:"Avalon,omitempty"`
				} `json:"All,omitempty"`
			} `json:"Gathering,omitempty"`
			Crafting struct {
				Total    int `json:"Total,omitempty"`
				Royal    int `json:"Royal,omitempty"`
				Outlands int `json:"Outlands,omitempty"`
				Avalon   int `json:"Avalon,omitempty"`
			} `json:"Crafting,omitempty"`
			CrystalLeague int `json:"CrystalLeague,omitempty"`
			FishingFame   int `json:"FishingFame,omitempty"`
			FarmingFame   int `json:"FarmingFame,omitempty"`
			Timestamp     any `json:"Timestamp,omitempty"`
		} `json:"LifetimeStatistics,omitempty"`
		DamageDone         float64 `json:"DamageDone,omitempty"`
		SupportHealingDone float64 `json:"SupportHealingDone,omitempty"`
	} `json:"Participants,omitempty"`
	GroupMembers []struct {
		AverageItemPower float64 `json:"AverageItemPower,omitempty"`
		Equipment        struct {
			MainHand struct {
				Type          string `json:"Type,omitempty"`
				Count         int    `json:"Count,omitempty"`
				Quality       int    `json:"Quality,omitempty"`
				ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
				PassiveSpells []any  `json:"PassiveSpells,omitempty"`
			} `json:"MainHand,omitempty"`
			OffHand any `json:"OffHand,omitempty"`
			Head    any `json:"Head,omitempty"`
			Armor   any `json:"Armor,omitempty"`
			Shoes   any `json:"Shoes,omitempty"`
			Bag     any `json:"Bag,omitempty"`
			Cape    any `json:"Cape,omitempty"`
			Mount   any `json:"Mount,omitempty"`
			Potion  any `json:"Potion,omitempty"`
			Food    any `json:"Food,omitempty"`
		} `json:"Equipment,omitempty"`
		Inventory          []any   `json:"Inventory,omitempty"`
		Name               string  `json:"Name,omitempty"`
		ID                 string  `json:"Id,omitempty"`
		GuildName          string  `json:"GuildName,omitempty"`
		GuildID            string  `json:"GuildId,omitempty"`
		AllianceName       string  `json:"AllianceName,omitempty"`
		AllianceID         string  `json:"AllianceId,omitempty"`
		AllianceTag        string  `json:"AllianceTag,omitempty"`
		Avatar             string  `json:"Avatar,omitempty"`
		AvatarRing         string  `json:"AvatarRing,omitempty"`
		DeathFame          int     `json:"DeathFame,omitempty"`
		KillFame           int     `json:"KillFame,omitempty"`
		FameRatio          float64 `json:"FameRatio,omitempty"`
		LifetimeStatistics struct {
			PvE struct {
				Total            int `json:"Total,omitempty"`
				Royal            int `json:"Royal,omitempty"`
				Outlands         int `json:"Outlands,omitempty"`
				Avalon           int `json:"Avalon,omitempty"`
				Hellgate         int `json:"Hellgate,omitempty"`
				CorruptedDungeon int `json:"CorruptedDungeon,omitempty"`
				Mists            int `json:"Mists,omitempty"`
			} `json:"PvE,omitempty"`
			Gathering struct {
				Fiber struct {
					Total    int `json:"Total,omitempty"`
					Royal    int `json:"Royal,omitempty"`
					Outlands int `json:"Outlands,omitempty"`
					Avalon   int `json:"Avalon,omitempty"`
				} `json:"Fiber,omitempty"`
				Hide struct {
					Total    int `json:"Total,omitempty"`
					Royal    int `json:"Royal,omitempty"`
					Outlands int `json:"Outlands,omitempty"`
					Avalon   int `json:"Avalon,omitempty"`
				} `json:"Hide,omitempty"`
				Ore struct {
					Total    int `json:"Total,omitempty"`
					Royal    int `json:"Royal,omitempty"`
					Outlands int `json:"Outlands,omitempty"`
					Avalon   int `json:"Avalon,omitempty"`
				} `json:"Ore,omitempty"`
				Rock struct {
					Total    int `json:"Total,omitempty"`
					Royal    int `json:"Royal,omitempty"`
					Outlands int `json:"Outlands,omitempty"`
					Avalon   int `json:"Avalon,omitempty"`
				} `json:"Rock,omitempty"`
				Wood struct {
					Total    int `json:"Total,omitempty"`
					Royal    int `json:"Royal,omitempty"`
					Outlands int `json:"Outlands,omitempty"`
					Avalon   int `json:"Avalon,omitempty"`
				} `json:"Wood,omitempty"`
				All struct {
					Total    int `json:"Total,omitempty"`
					Royal    int `json:"Royal,omitempty"`
					Outlands int `json:"Outlands,omitempty"`
					Avalon   int `json:"Avalon,omitempty"`
				} `json:"All,omitempty"`
			} `json:"Gathering,omitempty"`
			Crafting struct {
				Total    int `json:"Total,omitempty"`
				Royal    int `json:"Royal,omitempty"`
				Outlands int `json:"Outlands,omitempty"`
				Avalon   int `json:"Avalon,omitempty"`
			} `json:"Crafting,omitempty"`
			CrystalLeague int `json:"CrystalLeague,omitempty"`
			FishingFame   int `json:"FishingFame,omitempty"`
			FarmingFame   int `json:"FarmingFame,omitempty"`
			Timestamp     any `json:"Timestamp,omitempty"`
		} `json:"LifetimeStatistics,omitempty"`
	} `json:"GroupMembers,omitempty"`
	GvGMatch any    `json:"GvGMatch,omitempty"`
	BattleID int    `json:"BattleId,omitempty"`
	KillArea string `json:"KillArea,omitempty"`
	Category any    `json:"Category,omitempty"`
	Type     string `json:"Type,omitempty"`
}

type MemberInfo []struct {
	AverageItemPower float64 `json:"AverageItemPower,omitempty"`
	Equipment        struct {
		MainHand any `json:"MainHand,omitempty"`
		OffHand  any `json:"OffHand,omitempty"`
		Head     any `json:"Head,omitempty"`
		Armor    any `json:"Armor,omitempty"`
		Shoes    any `json:"Shoes,omitempty"`
		Bag      any `json:"Bag,omitempty"`
		Cape     any `json:"Cape,omitempty"`
		Mount    any `json:"Mount,omitempty"`
		Potion   any `json:"Potion,omitempty"`
		Food     any `json:"Food,omitempty"`
	} `json:"Equipment,omitempty"`
	Inventory          []any   `json:"Inventory,omitempty"`
	Name               string  `json:"Name,omitempty"`
	ID                 string  `json:"Id,omitempty"`
	GuildName          string  `json:"GuildName,omitempty"`
	GuildID            string  `json:"GuildId,omitempty"`
	AllianceName       any     `json:"AllianceName,omitempty"`
	AllianceID         string  `json:"AllianceId,omitempty"`
	AllianceTag        any     `json:"AllianceTag,omitempty"`
	Avatar             string  `json:"Avatar,omitempty"`
	AvatarRing         string  `json:"AvatarRing,omitempty"`
	DeathFame          int     `json:"DeathFame,omitempty"`
	KillFame           int     `json:"KillFame,omitempty"`
	FameRatio          float64 `json:"FameRatio,omitempty"`
	LifetimeStatistics struct {
		PvE struct {
			Total            int `json:"Total,omitempty"`
			Royal            int `json:"Royal,omitempty"`
			Outlands         int `json:"Outlands,omitempty"`
			Avalon           int `json:"Avalon,omitempty"`
			Hellgate         int `json:"Hellgate,omitempty"`
			CorruptedDungeon int `json:"CorruptedDungeon,omitempty"`
			Mists            int `json:"Mists,omitempty"`
		} `json:"PvE,omitempty"`
		Gathering struct {
			Fiber struct {
				Total    int `json:"Total,omitempty"`
				Royal    int `json:"Royal,omitempty"`
				Outlands int `json:"Outlands,omitempty"`
				Avalon   int `json:"Avalon,omitempty"`
			} `json:"Fiber,omitempty"`
			Hide struct {
				Total    int `json:"Total,omitempty"`
				Royal    int `json:"Royal,omitempty"`
				Outlands int `json:"Outlands,omitempty"`
				Avalon   int `json:"Avalon,omitempty"`
			} `json:"Hide,omitempty"`
			Ore struct {
				Total    int `json:"Total,omitempty"`
				Royal    int `json:"Royal,omitempty"`
				Outlands int `json:"Outlands,omitempty"`
				Avalon   int `json:"Avalon,omitempty"`
			} `json:"Ore,omitempty"`
			Rock struct {
				Total    int `json:"Total,omitempty"`
				Royal    int `json:"Royal,omitempty"`
				Outlands int `json:"Outlands,omitempty"`
				Avalon   int `json:"Avalon,omitempty"`
			} `json:"Rock,omitempty"`
			Wood struct {
				Total    int `json:"Total,omitempty"`
				Royal    int `json:"Royal,omitempty"`
				Outlands int `json:"Outlands,omitempty"`
				Avalon   int `json:"Avalon,omitempty"`
			} `json:"Wood,omitempty"`
			All struct {
				Total    int `json:"Total,omitempty"`
				Royal    int `json:"Royal,omitempty"`
				Outlands int `json:"Outlands,omitempty"`
				Avalon   int `json:"Avalon,omitempty"`
			} `json:"All,omitempty"`
		} `json:"Gathering,omitempty"`
		Crafting struct {
			Total    int `json:"Total,omitempty"`
			Royal    int `json:"Royal,omitempty"`
			Outlands int `json:"Outlands,omitempty"`
			Avalon   int `json:"Avalon,omitempty"`
		} `json:"Crafting,omitempty"`
		CrystalLeague int       `json:"CrystalLeague,omitempty"`
		FishingFame   int       `json:"FishingFame,omitempty"`
		FarmingFame   int       `json:"FarmingFame,omitempty"`
		Timestamp     time.Time `json:"Timestamp,omitempty"`
	} `json:"LifetimeStatistics,omitempty"`
}
