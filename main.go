package main

import (
	"log"
	"net/http"

	"github.com/JmPotato/Sudoku-backend/handlers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", handlers.HomeHandler)
	router.GET("/user/:username", handlers.UserHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
