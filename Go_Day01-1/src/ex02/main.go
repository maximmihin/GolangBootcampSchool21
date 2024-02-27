package main

import (
	"compareFS/config"
	"compareFS/intrernal/app"
	"log"
)

func main() {
	oldSnapshotFileName, newSnapshotFileName, err := config.GetOldAndNewSnap()
	if err != nil {
		log.Fatal(err)
	}
	app.Run(oldSnapshotFileName, newSnapshotFileName)
}
