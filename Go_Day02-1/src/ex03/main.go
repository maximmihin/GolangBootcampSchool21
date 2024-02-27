package main

import (
	"log"
	"myRotate/app"
	"myRotate/config"
)

func main() {
	archiveDir, logs, err := config.ParseInput()
	if err != nil {
		log.Fatal(err)
	}
	app.Run(archiveDir, logs)
}
