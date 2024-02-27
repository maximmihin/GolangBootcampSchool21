package config

import (
	"errors"
	"flag"
)

func GetOldAndNewDB() (oldDBFileName string, newDBFileName string, err error) {
	flag.StringVar(&oldDBFileName, "old", "", "")
	flag.StringVar(&newDBFileName, "new", "", "")
	flag.Parse()

	if oldDBFileName == "" || newDBFileName == "" {
		return "", "", errors.New("нужно указать две базы для сравнения")
	}

	return
}
