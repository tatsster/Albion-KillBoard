package discord

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/tatsster/albion_killboard/config"
)

func NewDiscordBot() (*discordgo.Session, error) {
	var (
		BotToken = os.Getenv("TOKEN")
	)

	discordBot, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return nil, err
	}

	// Register a message handler
	discordBot.AddHandler(MessangeHandler)
	// Open connection to discord
	err = discordBot.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return nil, err
	}
	return discordBot, nil
}

func MessangeHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from other channels or bots
	if m.ChannelID != config.ChannelID || m.Author.Bot {
		return
	}

	// Handle commands or messages in the designated channel if needed
	if m.Content == "!ping" {
		s.ChannelMessageSend(m.ChannelID, "Bot is running!")
	}
}
