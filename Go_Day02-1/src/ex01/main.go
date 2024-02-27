package main

import (
	"log"
	"myWc/app"
	"myWc/config"
)

func main() {
	Flag, err := config.ParseInput()
	if err != nil {
		log.Fatal(err)
	}
	app.Run(Flag)
}
