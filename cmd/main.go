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
	app := NewApp(config)
	log.Fatal(app.Run())
}
