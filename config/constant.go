package config

const (
	ChannelID         = "1123855668336672880"
	MemberSchedule    = "@hourly"
	KillDeathSchedule = "@every 30s"
	ALBION_API        = "https://gameinfo-sgp.albiononline.com/api/gameinfo"
	SQLITE_PATH       = "./sql/killboard.db"
	GuildID           = "MgQXspsCQmKh_m402mxJbw"
)

var (
	SingletonModel = &Singleton{}
)
