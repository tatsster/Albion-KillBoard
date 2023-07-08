package app

import (
	"fmt"

	"github.com/tatsster/albion_killboard/config"
	"github.com/tatsster/albion_killboard/internal/pkg/api"
	"github.com/tatsster/albion_killboard/internal/pkg/db"
)

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
