package app

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/tatsster/albion_killboard/config"
	"github.com/tatsster/albion_killboard/internal/pkg/api"
	"github.com/tatsster/albion_killboard/internal/pkg/db"
	"github.com/tatsster/albion_killboard/internal/pkg/util"
	"golang.org/x/sync/errgroup"
)

// Because in multithread and async, get db data can be old => Conflict in updating timestamp
// Therefore first run to update the latest data before can go cron
func FirstUpdate() {
	var (
		sqlite = config.SingletonModel.GetDatabase()
	)

	members, err := db.GetMembers(sqlite)
	if err != nil || len(members) == 0 {
		LogError(fmt.Errorf("Error get members from SQLite: %v", err))
		return
	}

	for _, member := range members {
		fmt.Println("Updating player: ", member.Name)

		kills, err := api.GetKills(member.ID)
		if err != nil {
			LogError(fmt.Errorf("Error in get kill: %v", err))
			return
		}
		if len(kills) > 0 {
			kill := kills[0]
			util.TruncateTime(&kill.TimeStamp)
			if kill.TimeStamp.After(member.LastKill.Time) {
				ProcessKillDeathEvent(kill)
				db.UpdateKillTime(sqlite, kill)
				// Update again time
				member.LastKill = sql.NullTime{Time: kill.TimeStamp, Valid: true}
			}
		}

		deaths, err := api.GetDeaths(member.ID)
		if err != nil {
			LogError(fmt.Errorf("Error in get death: %v", err))
			return
		}
		if len(deaths) > 0 {
			death := deaths[0]
			util.TruncateTime(&death.TimeStamp)
			if death.TimeStamp.After(member.LastDeath.Time) {
				ProcessKillDeathEvent(death)
				db.UpdatDeathTime(sqlite, death)
				// Update again time
				member.LastDeath = sql.NullTime{Time: death.TimeStamp, Valid: true}
			}
		}

		fmt.Println("Done player: ", member.Name)
	}
}

func UpdateKillDeath() {
	var (
		sqlite   = config.SingletonModel.GetDatabase()
		ctx      = context.Background()
		memberCh = make(chan config.Member)
	)

	fmt.Println("Updating kill board - ", time.Now().String())

	var wg sync.WaitGroup
	for i := 0; i < config.NumWorker; i++ {
		workerId := i
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer func() {
				wg.Done()
				fmt.Printf("[UpdateKillDeath][worker-%d] stopped\n", workerId)
			}()

			for {
				select {
				case <-ctx.Done():
					return
				case member, ok := <-memberCh:
					if !ok {
						return
					}
					// Fetch kill data of player
					// Process each kills reverse order & Save timestamp
					kills, err := api.GetKills(member.ID)
					if err != nil {
						fmt.Println("Error in get kill: ", err)
						return
					}
					for i := len(kills) - 1; i >= 0; i-- {
						kill := kills[i]
						util.TruncateTime(&kill.TimeStamp)
						// Null time or newer kill
						if kill.TimeStamp.After(member.LastKill.Time) {
							ProcessKillDeathEvent(kill)
							db.UpdateKillTime(sqlite, kill)
							// Update again time
							member.LastKill = sql.NullTime{Time: kill.TimeStamp, Valid: true}
						}
					}

					deaths, err := api.GetDeaths(member.ID)
					if err != nil {
						fmt.Println("Error in get death: ", err)
						return
					}
					for i := len(deaths) - 1; i >= 0; i-- {
						death := deaths[i]
						util.TruncateTime(&death.TimeStamp)
						if death.TimeStamp.After(member.LastDeath.Time) {
							ProcessKillDeathEvent(deaths[i])
							db.UpdatDeathTime(sqlite, death)
							// Update again time
							member.LastDeath = sql.NullTime{Time: death.TimeStamp, Valid: true}
						}
					}
				}
			}
		}(&wg)
	}

	// Get members
	members, err := db.GetMembers(sqlite)
	if err != nil || len(members) == 0 {
		fmt.Println("Error get members from SQLite:", err)
		return
	}
	eg, _ := errgroup.WithContext(ctx)
	for idx := range members {
		member := members[idx]
		eg.Go(func() error {
			memberCh <- member
			return nil
		})
		time.Sleep(100 * time.Millisecond)
	}

	wg.Wait()
}

// ProcessKillDeathEvent returns timeStamp event string
func ProcessKillDeathEvent(event config.Event) {
	var (
		discordBot         = config.SingletonModel.GetDiscord()
		participantStrings = make([]string, 0)
		description        = ""
	)

	imgPath, err := HandleImage(event)
	if err != nil {
		discordBot.ChannelMessageSend(config.ChannelID, "Error handle image.")
		return
	}

	// Send image
	file, err := os.Open(imgPath)
	if err != nil {
		discordBot.ChannelMessageSend(config.ChannelID, "Error opening image.")
		return
	}
	defer func() {
		file.Close()
		os.Remove(imgPath)
	}()

	totalParticipants := len(event.Participants)
	if totalParticipants > 0 {
		for i := 0; i < totalParticipants; i++ {
			if event.Participants[i].Name == event.Killer.Name {
				continue
			}
			participantStrings = append(participantStrings,
				fmt.Sprintf("[**%s**](https://albiononline.com/killboard/player/%s?server=live_sgp)", event.Participants[i].Name, event.Participants[i].ID))
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
		Title: fmt.Sprintf("%s (%d)  :crossed_swords:  %s (%d)", event.Killer.Name, int(event.Killer.AverageItemPower),
			event.Victim.Name, int(event.Victim.AverageItemPower)),
		Description: description,
		URL:         fmt.Sprintf("https://albiononline.com/en/killboard/kill/%d", event.EventID),
		Image: &discordgo.MessageEmbedImage{
			URL: "attachment://image.png",
		},
		Color: util.If(event.Killer.GuildID == config.GuildID, 0x00ff00, 0xff0000), // Set the color of the embed (optional)
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
}

func LogError(err error) {
	fmt.Println(err.Error())
	config.SingletonModel.GetDiscord().ChannelMessageSend(config.ChannelID, err.Error())
}
