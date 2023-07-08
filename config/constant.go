package config

const (
	NumWorker = 10

	ChannelID = "1123855668336672880"
	// ChannelID         = "1120599137512079361"
	MemberSchedule    = "@hourly"
	KillDeathSchedule = "@every 20s"

	ALBION_API        = "https://gameinfo-sgp.albiononline.com/api/gameinfo"
	ALBION_ITEMS_LINK = "https://render.albiononline.com/v1/item"
	SQLITE_PATH       = "./sql/killboard.db"
	GuildID           = "MgQXspsCQmKh_m402mxJbw"

	ASSET_DIR  = "./assets/items"
	RESULT_DIR = "./assets/image"
)

var (
	SingletonModel = &Singleton{}
)
