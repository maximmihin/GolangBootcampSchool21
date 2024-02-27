package app

import (
	"bufio"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"io"
	"log"
	"os"
)

func makeSetFileNames(snapShotFileFD *os.File) mapset.Set[string] {
	setFileNames := mapset.NewSet[string]()
	readLine := bufio.NewScanner(snapShotFileFD)

	for readLine.Scan() {
		setFileNames.Add(readLine.Text())
	}

	return setFileNames
}

func printDiff(firstFD, secondFD *os.File, typeDiff string) {
	snapSet := makeSetFileNames(firstFD)

	SnapScanner := bufio.NewScanner(secondFD)
	for SnapScanner.Scan() {
		line := SnapScanner.Text()
		contains := snapSet.Contains(line)
		if !contains {
			fmt.Println(typeDiff, line)
		}
	}
}

func resetOffsetFD(oldSnapFD, newSnapFD *os.File) error {
	_, err := oldSnapFD.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}
	_, err = newSnapFD.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	return nil
}

func Run(oldSnapshot, newSnapshot string) {

	oldSnapFD, err := os.Open(oldSnapshot)
	if err != nil {
		log.Fatal(err)
	}
	defer oldSnapFD.Close()

	newSnapFD, err := os.Open(newSnapshot)
	if err != nil {
		log.Fatal(err)
	}
	defer newSnapFD.Close()

	printDiff(oldSnapFD, newSnapFD, "ADDED")

	err = resetOffsetFD(oldSnapFD, newSnapFD)
	if err != nil {
		log.Fatal(err)
	}

	printDiff(newSnapFD, oldSnapFD, "REMOVED")
}
