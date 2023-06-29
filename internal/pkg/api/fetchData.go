package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/tatsster/albion_killboard/config"
)

type RequestAPI struct {
	BaseURL string
	query   url.Values
}

func CreateAPI(request RequestAPI) string {
	var queryString string
	if len(request.query) > 0 {
		queryString = fmt.Sprintf("?%s", request.query.Encode())
	}

	url := strings.TrimLeft(request.BaseURL, "/")
	return fmt.Sprintf("%s%s", url, queryString)
}

// GetKillDeath should return list of struct (kill death info) to construct image
func GetKillDeath() (string, error) {
	var (
		request = RequestAPI{
			BaseURL: config.ALBION_API,
			query: url.Values{
				"limit":   []string{"51"},
				"offset":  []string{"0"},
				"guildId": []string{config.GuildID},
			},
		}
	)

	res, err := http.Get(CreateAPI(request))
	if err != nil {
		fmt.Println("Error in get kill death from API: ", err)
		return "", err
	}

	// After return, clean up
	defer func() {
		_, _ = io.Copy(io.Discard, res.Body)
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Status is not OK")
		return "", errors.New("status not OK")
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error in reading response body: ", err)
		return "", err
	}

	var response KillDeathResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		fmt.Println("Error in JSON unmarshal: ", err)
		return "", err
	}

	for _, kill := range response {
		fmt.Printf("Killer: %s - Guild: %s\n", kill.Killer.Name, kill.Killer.GuildName)
	}
	return response[0].Killer.Name, nil
}
