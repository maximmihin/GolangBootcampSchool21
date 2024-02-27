package main

import (
	"encoding/csv"
	"ex00/db"
	"ex00/entities"
	"github.com/gocarina/gocsv"
	"io"
	"log"
	"os"
)

const pathToCsv = "/opt/goinfre/gradagas/Go_Day03-1/materials/data.csv"

func main() {

	el, err := db.New()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected to db")

	placesToAdd, err := getAllPlacesFromCsv(pathToCsv)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("parsed csv")

	_, err = el.TakePlaces(placesToAdd)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("success upload files")

}

func getAllPlacesFromCsv(csvPath string) ([]entities.Place, error) {

	csvFile, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	var places []entities.Place

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = '\t'
		return r
	})

	err = gocsv.Unmarshal(csvFile, &places)
	if err != nil {
		return nil, err
	}

	return places, err
}
