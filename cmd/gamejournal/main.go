package main

import (
	"context"
	"gamejournal/service"
	"gamejournal/store"
	"gamejournal/transport"
	"log"
	"net/http"
)

func main() {
	db, err := store.Open("gamejournal.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	j := service.New(db)
	log.Println(http.ListenAndServe(":8080", transport.New(j).Handler()))
	_ = context.Background()
}
