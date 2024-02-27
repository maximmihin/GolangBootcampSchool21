package main

import (
	"ex03/db"
	"ex03/handlers"
	"log"
	"net/http"
)

func main() {
	var err error
	store, err := db.New()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected db")

	http.HandleFunc("/api/recommend", handlers.HandleApiRecommend(store))

	err = http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ошибка сервера", err)
	}
}
