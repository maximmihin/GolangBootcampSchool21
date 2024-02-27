package app

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"myWc/config"
	"os"
	"sync"
	"unicode/utf8"
)

func Run(workMode int) {
	mapFiles := make(map[string]int, len(flag.Args()))
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	wg.Add(len(flag.Args()))

	switch workMode {
	case config.LinesFlag:
		for _, fileName := range flag.Args() {
			go ftCountLines(mapFiles, fileName, mu, wg)
		}
	case config.WordsFlag:
		for _, fileName := range flag.Args() {
			go ftCountWords(mapFiles, fileName, mu, wg)
		}
	case config.RunesFlag:
		for _, fileName := range flag.Args() {
			go ftCountChar(mapFiles, fileName, mu, wg)
		}
	default:
		panic(errors.New("неизвестный флаг"))
	}

	wg.Wait()
	for _, fileName := range flag.Args() {
		counter, ok := mapFiles[fileName]
		if ok {
			fmt.Println(counter, "\t", fileName)
		} else {
			fmt.Printf("не удалось открыть файл: %s\n", fileName)
		}
	}
}

func ftCountLines(mapFiles map[string]int, fileName string, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	linesCount := 0

	file, err := os.Open(fileName)
	if err != nil {
		return
	}

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		linesCount++
	}

	mu.Lock()
	mapFiles[fileName] = linesCount
	mu.Unlock()
}

func ftCountWords(mapFiles map[string]int, fileName string, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	wordsCount := 0

	file, err := os.Open(fileName)
	if err != nil {
		return
	}

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		wordsCount++
	}

	mu.Lock()
	mapFiles[fileName] = wordsCount
	mu.Unlock()
}

func ftCountChar(mapFiles map[string]int, fileName string, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	runeCount := 0

	file, err := os.Open(fileName)
	if err != nil {
		return
	}

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		runeCount += utf8.RuneCountInString(scan.Text())
	}

	mu.Lock()
	mapFiles[fileName] = runeCount
	mu.Unlock()
}
