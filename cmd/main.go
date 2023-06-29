package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/tatsster/albion_killboard/config"
	"github.com/tatsster/albion_killboard/internal/pkg/app"
	"github.com/tatsster/albion_killboard/internal/pkg/db"
	"github.com/tatsster/albion_killboard/internal/pkg/discord"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	discordBot, err := discord.NewDiscordBot()
	if err != nil {
		return
	}

	cron, err := app.NewCronScheduler()
	if err != nil {
		return
	}

	db, err := db.NewSqliteHandler()
	if err != nil {
		return
	}

	config.SingletonModel.WithDiscord(discordBot).WithScheduler(cron).WithDB(db)

	app.UpdateMember()
	config.SingletonModel.GetScheduler().Start()

	// Wait until the bot is stopped
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	config.SingletonModel.Shutdown()
}
