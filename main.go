package main

import (
	"log"
	"net/http"

	"github.com/JmPotato/Sudoku-backend/handlers"
	"github.com/julienschmidt/httprouter"
)

type Logger struct {
	handler http.Handler
}

func (l Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)
	l.handler.ServeHTTP(w, r)
}

func main() {
	router := httprouter.New()

	// Other handlers
	router.GET("/", handlers.HomeHandler)

	// User handlers
	router.GET("/user", handlers.GetUserHandler)
	router.POST("/user", handlers.CreatUserHandler)
	router.DELETE("/user", handlers.DeleteUserHandler)

	// Puzzle handlers
	router.GET("/puzzle", handlers.GetPuzzleHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", Logger{router}))
}
