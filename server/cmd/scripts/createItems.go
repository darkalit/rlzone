package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/darkalit/rlzone/server/config"
	"github.com/darkalit/rlzone/server/internal/items"
	"github.com/darkalit/rlzone/server/pkg/db/mysql"
)

type Item struct {
	Name       string      `json:"name"`
	Link       string      `json:"link"`
	Image      string      `json:"src"`
	Type       string      `json:"Type"`
	Hitbox     *string     `json:"Hitbox"`
	Reactive   *bool       `json:"Reactive"`
	Quality    string      `json:"Quality"`
	TradeIn    bool        `json:"Trade In"`
	Paintable  interface{} `json:"Paintable"`
	Blueprints bool        `json:"Blueprints"`
	Released   string      `json:"Released"`
	Platform   string      `json:"Platform"`
	Sideswipe  string      `json:"Sideswipe"`
	Series     string      `json:"Series"`
}

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := mysql.NewMySqlDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	data, err := os.ReadFile("./assets/data.json")
	if err != nil {
		log.Fatal(err)
	}

	var itemsArray []Item
	err = json.Unmarshal([]byte(data), &itemsArray)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.TODO()
	itemRepo := items.NewItemRepository(cfg, db)

	for _, val := range itemsArray {
		switch val.Type {
		case "Body":
			fallthrough
		case "Decal":
			fallthrough
		case "Wheels":
			fallthrough
		case "Rocket Boost":
			fallthrough
		case "Antenna":
			fallthrough
		case "Goal Explosion":
			fallthrough
		case "Trail":
			var paintable bool
			switch t := val.Paintable.(type) {
			case string:
				paintable = true
			case bool:
				paintable = t
			}

			res, err := http.Get(val.Image)
			if err != nil {
				log.Print(err)
				continue
			}
			defer res.Body.Close()

			imageData, err := io.ReadAll(res.Body)
			if err != nil {
				log.Print(err)
				continue
			}

			img64 := base64.StdEncoding.EncodeToString(imageData)
			img64 = "data:" + res.Header.Get("content-type") + ";base64," + img64

			item := items.Item{
				Name:       val.Name,
				Image:      img64,
				Type:       val.Type,
				Quality:    val.Quality,
				Hitbox:     val.Hitbox,
				Reactive:   val.Reactive,
				TradeIn:    val.TradeIn,
				Paintable:  paintable,
				Blueprints: val.Blueprints,
				Released:   val.Released,
				Platform:   val.Platform,
				Sideswipe:  val.Sideswipe,
				Series:     val.Series,
			}

			// fmt.Printf("%+v \n\n", item)
			itemRepo.Create(ctx, &item)
			break
		default:
			break
		}
	}
}
