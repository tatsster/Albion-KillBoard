package discord

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/tatsster/albion_killboard/config"
	"github.com/tatsster/albion_killboard/internal/pkg/api"
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
		// https://render.albiononline.com/v1/item/T8_2H_DUALSCIMITAR_UNDEAD@4.png?count=1&quality=4&size=128

		s.ChannelMessageSend(m.ChannelID, "Bot is running!")
	}

	// Handle commands or messages in the designated channel if needed
	// if m.Content == "!init image" {
	if strings.HasPrefix(m.Content, "!init image") {
		args := strings.Split(m.Content, " ")
		if len(args) == 2 {
			args[2] = "1"
		}
		itemId, err := strconv.Atoi(args[2])
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Wrong format: !init image {startSttItem}")
			return
		}
		stt, err := InitImage(itemId)
		if err == nil {
			s.ChannelMessageSend(m.ChannelID, "Done")
		} else {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error at %d", stt))
		}
	}
}

func InitImage(startId int) (int, error) {
	stt := 0
	file, err := os.Open("assets/items.txt")
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// read line by line
	for scanner.Scan() {
		array := strings.Split(strings.TrimSpace(scanner.Text()), ": ")
		stt, err := strconv.Atoi(array[0])
		if stt < startId {
			continue
		}
		array[1] = strings.TrimSpace(array[1])
		if (stt < 1642) || (stt > 8481) && err == nil {
			_, err = api.SaveImage(array[1], 0)
		} else {
			for i := 1; i <= 5; i++ {
				_, err = api.SaveImage(array[1], i)
				if err != nil {
					break
				}
			}
		}
		if err != nil {
			return stt, err
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return stt, err
}
