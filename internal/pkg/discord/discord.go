package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tatsster/albion_killboard/config"
)

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
