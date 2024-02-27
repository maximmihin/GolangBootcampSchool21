package config

import (
	"errors"
	"flag"
)

var (
	ErrAbsenceFile = errors.New("нужно передать файл через флаг -f")
)

func GetDBFileName() (DBFileName string, err error) {
	flag.StringVar(&DBFileName, "f", "", "usage")
	flag.Parse()

	if DBFileName == "" {
		return "", ErrAbsenceFile
	}

	return
}
