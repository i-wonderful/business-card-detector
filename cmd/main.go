package main

import (
	. "card_detector/internal/app"
	"log"
)

func main() {
	config, err := NewConfigFromYml()
	if err != nil {
		log.Fatal(err)
		return
	}
	app := NewApp2(config)
	log.Fatal(app.Run())
}
