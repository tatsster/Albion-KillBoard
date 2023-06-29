package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/tatsster/albion_killboard/config"
	"github.com/tatsster/albion_killboard/internal/pkg/util"
)

func GetMembers() (config.MemberInfo, error) {
	res, err := http.Get(util.BaseURl(config.ALBION_API).BuildURL(
		"/guilds/:guild_id/members",
		map[string]string{
			"guild_id": config.GuildID,
		},
		url.Values{}),
	)

	if err != nil {
		fmt.Println("Error in get guild members from API: ", err)
		return nil, err
	}

	// After return, clean up
	defer func() {
		_, _ = io.Copy(io.Discard, res.Body)
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Status is not OK")
		return nil, errors.New("status not OK")
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error in reading response body: ", err)
		return nil, err
	}

	var response config.MemberInfo
	err = json.Unmarshal(data, &response)
	if err != nil {
		fmt.Println("Error in JSON unmarshal: ", err)
		return nil, err
	}
	return response, nil
}

func GetKills(playerID string) (config.KillDeathResponse, error) {
	url := util.BaseURl(config.ALBION_API).BuildURL(
		"/players/:player_id/kills",
		map[string]string{
			"player_id": playerID,
		},
		url.Values{})

	response, err := GetKillDeath(url)
	if err != nil {
		fmt.Println("Error in get kill: ", err)
	}
	return response, err
}

func GetDeaths(playerID string) (config.KillDeathResponse, error) {
	url := util.BaseURl(config.ALBION_API).BuildURL(
		"/players/:player_id/deaths",
		map[string]string{
			"player_id": playerID,
		},
		url.Values{})

	response, err := GetKillDeath(url)
	if err != nil {
		fmt.Println("Error in get death: ", err)
	}
	return response, err
}

func GetKillDeath(url string) (config.KillDeathResponse, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	// After return, clean up
	defer func() {
		_, _ = io.Copy(io.Discard, res.Body)
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Status is not OK")
		return nil, errors.New("status not OK")
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error in reading response body: ", err)
		return nil, err
	}

	var response config.KillDeathResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		fmt.Println("Error in JSON unmarshal: ", err)
		return nil, err
	}
	return response, nil
}
