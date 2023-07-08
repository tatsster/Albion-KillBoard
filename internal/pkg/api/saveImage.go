package api

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/tatsster/albion_killboard/config"
	"github.com/tatsster/albion_killboard/internal/pkg/util"
)

func SaveImage(itemId string, quality int) (bool, error) {
	// https://render.albiononline.com/v1/item/T8_2H_DUALSCIMITAR_UNDEAD@4.png?count=1&quality=4&size=128
	querries := url.Values{}
	querries.Add("quality", strconv.Itoa(quality))
	querries.Add("size", "128")

	res, err := http.Get(util.BaseURl(config.ALBION_ITEMS_LINK).BuildURL(
		"/:item_id.png",
		map[string]string{
			"item_id": itemId,
		},
		querries),
	)

	if err != nil {
		fmt.Println("Error in download image from url ", err)
		return false, err
	}

	defer res.Body.Close()

	// Some item dont have image
	if res.StatusCode == http.StatusNotFound {
		return false, nil
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println("Status is not OK")
		fmt.Print(res.Request.URL)
		return false, errors.New("status not OK")
	}

	file, err := os.Create(fmt.Sprintf("assets/items/%s_%d.png", itemId, quality))
	if err != nil {
		log.Println(err)
		return false, errors.New("can't create file")
	}
	defer file.Close()
	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return true, nil
}
