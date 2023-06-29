package config

import (
	"database/sql"

	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
)

type Singleton struct {
	discordBot    *discordgo.Session
	cronScheduler *cron.Cron
	database      *sql.DB
}

func (s *Singleton) WithDiscord(discordBot *discordgo.Session) *Singleton {
	s.discordBot = discordBot
	return s
}

func (s *Singleton) WithScheduler(cronScheduler *cron.Cron) *Singleton {
	s.cronScheduler = cronScheduler
	return s
}

func (s *Singleton) WithDB(database *sql.DB) *Singleton {
	s.database = database
	return s
}

func (s *Singleton) GetDiscord() *discordgo.Session {
	return s.discordBot
}

func (s *Singleton) GetScheduler() *cron.Cron {
	return s.cronScheduler
}

func (s *Singleton) GetDatabase() *sql.DB {
	return s.database
}

func (s *Singleton) Shutdown() {
	s.cronScheduler.Stop()
	s.discordBot.Close()
	s.database.Close()
}

type Member struct {
	ID        string
	Name      string
	LastKill  sql.NullTime
	LastDeath sql.NullTime
}
