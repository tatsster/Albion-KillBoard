package config

const (
	ChannelID = "1123855668336672880"
	// ChannelID         = "1120599137512079361"
	MemberSchedule    = "@hourly"
	KillDeathSchedule = "@every 30s"
	ALBION_API        = "https://gameinfo-sgp.albiononline.com/api/gameinfo"
	ALBION_ITEMS_LINK = "https://render.albiononline.com/v1/item"
	SQLITE_PATH       = "./sql/killboard.db"
	GuildID           = "MgQXspsCQmKh_m402mxJbw"
)

var (
	SingletonModel = &Singleton{}
)
