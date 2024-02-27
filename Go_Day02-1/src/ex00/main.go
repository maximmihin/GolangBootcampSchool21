package main

import (
	"log"
	"myFind/app"
	"myFind/config"
)

func main() {
	conf, err := config.ParseInput()
	if err != nil {
		log.Fatal(err)
	}
	app.Run(conf)
}
