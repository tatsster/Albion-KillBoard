package app

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"github.com/tatsster/albion_killboard/config"
	"github.com/tatsster/albion_killboard/internal/pkg/api"
	"github.com/tatsster/albion_killboard/internal/pkg/db"
)

func NewCronScheduler() (*cron.Cron, error) {
	c := cron.New()

	id1, err := c.AddFunc(config.MemberSchedule, UpdateMember)
	if err != nil {
		fmt.Println("fail to schedule update members task: ", err)
		return nil, err
	}
	fmt.Printf("registered an entry for update members task: %v\n", id1)

	id2, err := c.AddFunc(config.KillDeathSchedule, UpdateKillDeath)
	if err != nil {
		fmt.Println("fail to schedule update killboard task: ", err)
		return nil, err
	}
	fmt.Printf("registered an entry for update killboard task: %v\n", id2)
	return c, nil
}

func UpdateMember() {
	var (
		sqlite = config.SingletonModel.GetDatabase()
	)

	members, err := api.GetMembers()
	if err != nil {
		fmt.Println("Error in get members: ", err)
		return
	}

	err = db.UpdateMembers(sqlite, members)
	if err != nil {
		fmt.Println("Error in update members in SQLite: ", err)
		return
	}
	fmt.Println("Update members done!")
}

func UpdateKillDeath() {
	var (
		// discordBot = config.SingletonModel.GetDiscord()
		sqlite     = config.SingletonModel.GetDatabase()
	)

	// Get members
	members, err := db.GetMembers(sqlite)
	if err != nil || len(members) == 0 {
		fmt.Println("Error get members from SQLite:", err)
	}

	for _, mem := range members {
		// Fetch data
		kills, err := api.GetKills(mem.ID)
		if err != nil {
			fmt.Println("Error in get kill: ", err)
			return
		}
		deaths, err := api.GetDeaths(mem.ID)
		if err != nil {
			fmt.Println("Error in get death: ", err)
			return
		}

		// Send result as embed to discord
		// _, err := discordBot.ChannelMessageSendEmbed(config.ChannelID, data)
		for _, kill := range kills {
			// Pre process data - Image handling
			// _, _ = discordBot.ChannelMessageSend(config.ChannelID, kill.Killer.Name)
			fmt.Println(kill.Killer.Name)
		}

		for _, death := range deaths {
			// Pre process data - Image handling
			// _, _ = discordBot.ChannelMessageSend(config.ChannelID, death.Killer.Name)
			fmt.Println(death.Killer.Name)
		}

		// Update last kill/death time
	}

}
