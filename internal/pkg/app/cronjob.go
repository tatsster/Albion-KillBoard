package app

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"github.com/tatsster/albion_killboard/config"
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
