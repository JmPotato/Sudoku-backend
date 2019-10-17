package main

import (
	"log"
	"net/http"

	"github.com/JmPotato/Sudoku-backend/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
