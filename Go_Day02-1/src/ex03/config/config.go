package config

import (
	"errors"
	"flag"
)

func ParseInput() (pathDir string, logs []string, err error) {
	flag.StringVar(&pathDir, "a", "", "")
	flag.Parse()

	if pathDir != "" {
		pathDir = pathDir + "/"
	}

	logs = flag.Args()
	if len(logs) == 0 {
		err = errors.New("нужно указать, хотя бы один файл")
	}

	return
}
