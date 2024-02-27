package config

import (
	"errors"
	"flag"
)

func GetOldAndNewSnap() (oldSnapshotFileName string, newSnapshotFileName string, err error) {
	flag.StringVar(&oldSnapshotFileName, "old", "", "")
	flag.StringVar(&newSnapshotFileName, "new", "", "")
	flag.Parse()

	if oldSnapshotFileName == "" || newSnapshotFileName == "" {
		return "", "", errors.New("нужно указать две базы для сравнения")
	}

	return
}
