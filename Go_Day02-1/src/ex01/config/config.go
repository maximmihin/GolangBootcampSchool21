package config

import (
	"errors"
	"flag"
)

const (
	LinesFlag = iota + 1
	RunesFlag
	WordsFlag
)

func ParseInput() (mode int, err error) {
	var lFlag, mFlag, wFlag bool
	flag.BoolVar(&lFlag, "l", false, "count lines")
	flag.BoolVar(&mFlag, "m", false, "count characters")
	flag.BoolVar(&wFlag, "w", false, "count words")
	flag.Parse()

	flagCounter := 0
	if lFlag {
		flagCounter++
		mode = LinesFlag
	}
	if mFlag {
		flagCounter++
		mode = RunesFlag
	}
	if wFlag {
		flagCounter++
		mode = WordsFlag
	}

	if flagCounter == 0 {
		mode = WordsFlag
	} else if flagCounter > 1 {
		err = errors.New("можно указать только 1 флаг")
	}

	if len(flag.Args()) < 1 {
		err = errors.New("нужно указать хотя бы 1 файл для чтения")
	}

	return
}
