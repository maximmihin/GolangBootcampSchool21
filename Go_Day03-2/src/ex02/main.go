package main

import (
	"ex02/db"
	"ex02/handlers"
	"log"
	"net/http"
)

func main() {
	store, err := db.New()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected db")

	http.HandleFunc("/api/places", handlers.HandleApiPlaces(store))

	err = http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ошибка сервера", err)
	}
}
