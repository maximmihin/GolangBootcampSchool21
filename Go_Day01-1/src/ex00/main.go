package main

import (
	"log"

	"readDB/config"
	"readDB/internal/app"
)

func main() {
	dbFileNameWithExtension, err := config.GetDBFileName()
	if err != nil {
		log.Fatal(err)
	}

	app.Run(dbFileNameWithExtension)
}
