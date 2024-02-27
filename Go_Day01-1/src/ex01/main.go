package main

import (
	"compareDB/config"
	"compareDB/internal/app"
	"log"
)

func main() {
	oldDBFileName, newDBFileName, err := config.GetOldAndNewDB()
	if err != nil {
		log.Fatal(err)
	}

	app.Run(oldDBFileName, newDBFileName)
}
