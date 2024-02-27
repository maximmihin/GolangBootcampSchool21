package main

import (
	"ex04/db"
	"ex04/handlers"
	"log"
	"net/http"
)

func main() {

	store, err := db.New()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected db")

	http.HandleFunc("/api/get_token", handlers.HandleApiGetToken("extra_secret"))
	http.HandleFunc("/api/recommend", handlers.MiddlewareJwtCheck("extra_secret", handlers.HandleApiRecommend(store)))

	err = http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ошибка сервера", err)
	}
}
