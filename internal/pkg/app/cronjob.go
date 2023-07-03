package app

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
	"github.com/tatsster/albion_killboard/config"
	"github.com/tatsster/albion_killboard/internal/pkg/api"
	"github.com/tatsster/albion_killboard/internal/pkg/db"
	"github.com/tatsster/albion_killboard/internal/pkg/util"
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
		discordBot = config.SingletonModel.GetDiscord()
		sqlite     = config.SingletonModel.GetDatabase()
		wg         sync.WaitGroup
	)

	// Get members
	members, err := db.GetMembers(sqlite)
	if err != nil || len(members) == 0 {
		fmt.Println("Error get members from SQLite:", err)
	}

	members = members[:4]

	for _, mem := range members {
		// Each member is 1 goroutine
		wg.Add(1)
		go func(mem config.Member, wg *sync.WaitGroup) {
			defer wg.Done()

			// Fetch data
			kills, err := api.GetKills(mem.ID)
			if err != nil {
				fmt.Println("Error in get kill: ", err)
				return
			}
			// deaths, err := api.GetDeaths(mem.ID)
			if err != nil {
				fmt.Println("Error in get death: ", err)
				return
			}

			for _, kill := range kills {
				imgPath, err := HandleImage(kill)
				if err != nil {
					discordBot.ChannelMessageSend(config.ChannelID, "Error handle image.")
					continue
				}
				fmt.Println(imgPath)
				file, err := os.Open(imgPath)
				if err != nil {
					discordBot.ChannelMessageSend(config.ChannelID, "Error opening image.")
					continue
				}
				defer file.Close()
				participantStrings := []string{}
				totalParticipants := len(kill.Participants)
				description := ""
				if totalParticipants > 0 {
					for i := 0; i < totalParticipants; i++ {
						if kill.Participants[i].Name == kill.Killer.Name {
							continue
						}
						participantStrings = append(participantStrings,
							fmt.Sprintf("[**%s**](https://albiononline.com/killboard/player/%s?server=live_sgp)", kill.Participants[i].Name, kill.Participants[i].ID))
						if i == 3 {
							break
						}
					}
					if len(participantStrings) > 0 {
						description = "Participants: " + strings.Join(participantStrings, ", ")
					}
				}

				// Create a new embed
				embed := &discordgo.MessageEmbed{
					Title: fmt.Sprintf("%s (%d)  :crossed_swords:  %s (%d)", kill.Killer.Name, int(kill.Killer.AverageItemPower),
						kill.Victim.Name, int(kill.Victim.AverageItemPower)),
					Description: description,
					URL:         fmt.Sprintf("https://albiononline.com/en/killboard/kill/%d", kill.EventID),
					Image: &discordgo.MessageEmbedImage{
						URL: "attachment://image.png",
					},
					Color: util.If(kill.Killer.GuildID == config.GuildID, 0x00ff00, 0xff0000), // Set the color of the embed (optional)
				}

				// Send the embedded message with the attached image
				discordBot.ChannelMessageSendComplex(config.ChannelID, &discordgo.MessageSend{
					Embed: embed,
					Files: []*discordgo.File{
						{
							Name:   "image.png",
							Reader: file,
						},
					},
				})

				// embed := &discordgo.MessageEmbed{
				// 	Title: "d33r (2222)  :crossed_swords:  dubo (1111)",
				// 	URL:   "https://albiononline.com/en/killboard/kill/81324154",

				// 	Color: 0x00ff00, // Green
				// 	Image: &discordgo.MessageEmbedImage{
				// 		URL: "https://media.discordapp.net/attachments/1016676828569669672/1121088400280277112/image.png",
				// 	},
				// }
				// discordBot.ChannelMessageSendEmbed(config.ChannelID, embed)
				// fmt.Println(kill.Killer.name)
			}

			// Send result as embed to discord
			// _, err := discordBot.ChannelMessageSendEmbed(config.ChannelID, data)
			// for _, kill := range kills {
			// 	// Pre process data - Image handling
			// 	// _, _ = discordBot.ChannelMessageSend(config.ChannelID, kill.Killer.Name)
			// 	fmt.Println(kill.Killer.Name)
			// }

			// for _, death := range deaths {
			// 	// Pre process data - Image handling
			// 	// _, _ = discordBot.ChannelMessageSend(config.ChannelID, death.Killer.Name)
			// 	fmt.Println(death.Killer.Name)
			// }

			// Update last kill/death time
		}(mem, &wg)
		// return
	}

	wg.Wait()
}
