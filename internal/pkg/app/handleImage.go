package app

import (
	"fmt"
	"image"
	"image/draw"
	"log"
	"os"
	"strconv"

	"github.com/fogleman/gg"
	"github.com/tatsster/albion_killboard/config"
	"github.com/tatsster/albion_killboard/internal/pkg/api"
	"github.com/tatsster/albion_killboard/internal/pkg/util"
	"golang.org/x/image/font"
)

var background image.Image

var bounds []image.Rectangle
var fontBold font.Face
var fontNormal font.Face

func init() {
	imgFile1, err := os.Open("assets/background-1.png")

	if err != nil {
		fmt.Println(err)
	}

	background, _, err = image.Decode(imgFile1)
	if err != nil {
		fmt.Println(err)
	}

	// font, err := truetype.Parse(goregular.TTF)
	fontBold, err = gg.LoadFontFace("assets/fonts/Helvetica-Bold.ttf", 28)
	if err != nil {
		log.Fatal(err)
	}
	fontNormal, err = gg.LoadFontFace("assets/fonts/Helvetica.ttf", 28)
	if err != nil {
		log.Fatal(err)
	}

	// face := truetype.NewFace(font, &truetype.Options{Size: 48})

	bounds = []image.Rectangle{
		background.Bounds(),
		background.Bounds().Add(image.Point{65, 260}),
		background.Bounds().Add(image.Point{325, 260}),
		background.Bounds().Add(image.Point{40, 135}),
		background.Bounds().Add(image.Point{350, 135}),
		background.Bounds().Add(image.Point{40, 390}),
		background.Bounds().Add(image.Point{350, 390}),
		background.Bounds().Add(image.Point{195, 145}),
		background.Bounds().Add(image.Point{195, 260}),
		background.Bounds().Add(image.Point{195, 375}),
		background.Bounds().Add(image.Point{195, 488}),
		background.Bounds().Add(image.Point{800, 260}),
		background.Bounds().Add(image.Point{1060, 260}),
		background.Bounds().Add(image.Point{770, 135}),
		background.Bounds().Add(image.Point{1085, 135}),
		background.Bounds().Add(image.Point{770, 390}),
		background.Bounds().Add(image.Point{1085, 390}),
		background.Bounds().Add(image.Point{925, 145}),
		background.Bounds().Add(image.Point{925, 260}),
		background.Bounds().Add(image.Point{925, 375}),
		background.Bounds().Add(image.Point{925, 488}),
	}
}

func readImage(item struct {
	Type          string `json:"Type,omitempty"`
	Count         int    `json:"Count,omitempty"`
	Quality       int    `json:"Quality,omitempty"`
	ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
	PassiveSpells []any  `json:"PassiveSpells,omitempty"`
}) (image.Image, error) {
	file, err := os.Open(fmt.Sprintf("assets/items/%s_%d.png", item.Type, item.Quality-1))
	if err != nil {
		return nil, err
	}
	image, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return image, nil
}

func InsertItemImage(item struct {
	Type          string `json:"Type,omitempty"`
	Count         int    `json:"Count,omitempty"`
	Quality       int    `json:"Quality,omitempty"`
	ActiveSpells  []any  `json:"ActiveSpells,omitempty"`
	PassiveSpells []any  `json:"PassiveSpells,omitempty"`
}, bound image.Rectangle, rgba *image.RGBA) error {
	if item.Count == 0 {
		return nil
	}
	file, err := os.Open(fmt.Sprintf("assets/items/%s_%d.png", item.Type, item.Quality))
	// trying download img if missing
	if err != nil {
		_, err2 := api.SaveImage(item.Type, item.Quality)
		if err2 != nil {
			return err
		}
		file, err = os.Open(fmt.Sprintf("assets/items/%s_%d.png", item.Type, item.Quality))
		if err != nil {
			return err
		}
	}
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}
	draw.Draw(rgba, bound, img, image.Point{0, 0}, draw.Over)
	//  draw.Draw(rgba, bounds[3], bag, image.Point{0, 0}, draw.Over)

	return nil
}

func HandleImage(event config.Event) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			// Handle the panic gracefully
			fmt.Println("Panic occurred: ", r)
			fmt.Println(event.TotalVictimKillFame)
			fmt.Println(event.Killer.Name)
			fmt.Println(event.Victim.Name)
			fmt.Println(event.Killer.GuildName)
			fmt.Println(event.Victim.GuildName)
		}
	}()

	r := image.Rectangle{image.Point{0, 0}, bounds[0].Max}
	rgba := image.NewRGBA(r)

	draw.Draw(rgba, bounds[0], background, image.Point{0, 0}, draw.Src)
	// Killer equipment
	InsertItemImage(event.Killer.Equipment.MainHand, bounds[1], rgba)
	InsertItemImage(event.Killer.Equipment.OffHand, bounds[2], rgba)
	InsertItemImage(event.Killer.Equipment.Bag, bounds[3], rgba)
	InsertItemImage(event.Killer.Equipment.Cape, bounds[4], rgba)
	InsertItemImage(event.Killer.Equipment.Potion, bounds[5], rgba)
	InsertItemImage(event.Killer.Equipment.Food, bounds[6], rgba)
	InsertItemImage(event.Killer.Equipment.Head, bounds[7], rgba)
	InsertItemImage(event.Killer.Equipment.Armor, bounds[8], rgba)
	InsertItemImage(event.Killer.Equipment.Shoes, bounds[9], rgba)
	InsertItemImage(event.Killer.Equipment.Mount, bounds[10], rgba)
	// Victim equipment
	InsertItemImage(event.Victim.Equipment.MainHand, bounds[11], rgba)
	InsertItemImage(event.Victim.Equipment.OffHand, bounds[12], rgba)
	InsertItemImage(event.Victim.Equipment.Bag, bounds[13], rgba)
	InsertItemImage(event.Victim.Equipment.Cape, bounds[14], rgba)
	InsertItemImage(event.Victim.Equipment.Potion, bounds[15], rgba)
	InsertItemImage(event.Victim.Equipment.Food, bounds[16], rgba)
	InsertItemImage(event.Victim.Equipment.Head, bounds[17], rgba)
	InsertItemImage(event.Victim.Equipment.Armor, bounds[18], rgba)
	InsertItemImage(event.Victim.Equipment.Shoes, bounds[19], rgba)
	InsertItemImage(event.Victim.Equipment.Mount, bounds[20], rgba)

	dc := gg.NewContextForRGBA(rgba)
	// dc.SetRGB(1, 1, 1)
	// dc.Clear()
	dc.SetRGB(0, 0, 0)

	dc.SetFontFace(fontBold)
	dc.DrawStringAnchored(util.FormatInt(event.TotalVictimKillFame), 1250/2, 290, 0.5, 0.5)

	dc.DrawStringAnchored(event.Killer.Name, 260, 60, 0.5, 0.5)
	dc.DrawStringAnchored(event.Victim.Name, 985, 60, 0.5, 0.5)

	dc.SetFontFace(fontNormal)

	dc.DrawStringAnchored(util.GetGuildAndTag(event.Killer.GuildName, event.Killer.AllianceName), 260, 95, 0.5, 0.5)
	dc.DrawStringAnchored(util.GetGuildAndTag(event.Victim.GuildName, event.Victim.AllianceName), 985, 95, 0.5, 0.5)

	// dc.SetFontFace(fontNormal)
	// dc.SetRGB(0.75, 0.75, 0.75)
	if event.NumberOfParticipants == 0 {
		event.NumberOfParticipants = 1
	}
	splitKillFame := fmt.Sprintf("%d x %s", event.NumberOfParticipants, util.FormatInt(int(event.TotalVictimKillFame/event.NumberOfParticipants)))
	dc.DrawStringAnchored(splitKillFame, 1250/2, 320, 0.5, 0.5)

	dc.SetFontFace(fontBold)

	dc.SetRGB(1, 1, 1)
	dc.DrawStringAnchored(strconv.Itoa(int(event.Killer.AverageItemPower)), 1250/2-56, 402, 0.5, 0.5)
	dc.DrawStringAnchored(strconv.Itoa(int(event.Victim.AverageItemPower)), 1250/2+54, 402, 0.5, 0.5)

	dc.SetRGB(0.25, 0.25, 0.25)
	timeString := event.TimeStamp.Format("15:04 2006-01-02")
	dc.DrawStringAnchored(timeString, 1250/2, 500, 0.5, 0.5)
	// fmt.Println(event.Victim.Inventory)

	imgPath := fmt.Sprintf("assets/image/%d.png", event.EventID)
	err := dc.SavePNG(imgPath)
	if err != nil {
		return "", err
	}
	return imgPath, nil
}
