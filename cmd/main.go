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

	// Check if the folder assets
	if _, err := os.Stat(config.ASSET_DIR); os.IsNotExist(err) {
		// Folder does not exist, create it
		err := os.MkdirAll(config.ASSET_DIR, 0755) // 0755 sets the folder permissions
		if err != nil {
			fmt.Printf("Error creating folder: %s\n", err)
			return
		}
		fmt.Printf("Folder created: %s\n", config.ASSET_DIR)
	}

	// Check if the folder result
	if _, err := os.Stat(config.RESULT_DIR); os.IsNotExist(err) {
		// Folder does not exist, create it
		err := os.MkdirAll(config.RESULT_DIR, 0755) // 0755 sets the folder permissions
		if err != nil {
			fmt.Printf("Error creating folder: %s\n", err)
			return
		}
		fmt.Printf("Folder created: %s\n", config.RESULT_DIR)
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
	app.FirstUpdate()
	config.SingletonModel.GetScheduler().Start()

	// Wait until the bot is stopped
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	config.SingletonModel.Shutdown()
}
