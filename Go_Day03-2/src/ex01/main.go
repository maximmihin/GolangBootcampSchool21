package main

import (
	"ex01/db"
	"ex01/handlers"
	"html/template"
	"log"
	"net/http"
)

func main() {

	store, err := db.New()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected db")

	indexTemple, err := template.New("index.gohtml").ParseFiles("templates/index.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handlers.HandleHtmlPlaces(store, indexTemple))

	err = http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ошибка сервера", err)
	}
}
