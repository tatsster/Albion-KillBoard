package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"github.com/tatsster/albion_killboard/config"
	"github.com/tatsster/albion_killboard/internal/pkg/api"
	"github.com/tatsster/albion_killboard/internal/pkg/discord"

	"github.com/bwmarrin/discordgo"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	BotToken := os.Getenv("TOKEN")

	discordBot, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
	}

	// Register a message handler
	discordBot.AddHandler(discord.MessangeHandler)

	// Open connection to discord
	err = discordBot.Open()
	if err != nil {
		log.Fatal("Error opening connection: ", err)
	}

	c := cron.New()
	_, err = c.AddFunc(config.CronSchedule, func() {
		// Fetch data
		data, err := api.GetKillDeath()
		if err != nil {
			fmt.Println("Error in get kill death: ", err)
			return
		}

		// Pre process data - Image handling

		// Send result as embed to discord
		// _, err := discordBot.ChannelMessageSendEmbed(config.ChannelID, data)
		_, err = discordBot.ChannelMessageSend(config.ChannelID, data)
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	})

	if err != nil {
		log.Fatal("Error scheduling cronjob: ", err)
	}

	c.Start()

	// Wait until the bot is stopped
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Stop the cron job
	c.Stop()

	// Close the Discord session
	discordBot.Close()
}
