package handlers

import (
	"log"
	"net/http"

	"github.com/JmPotato/Sudoku-backend/models"
	"github.com/julienschmidt/httprouter"
)

// UserHandler processes the request from /user/:username, check whether the username exists.
func UserHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var user = models.User{Username: p.ByName("username")}
	log.Printf("finding username=%q", user.Username)
	err := user.CheckUserName()
	if err != nil {
		log.Println(err)
	}
	SendResponse(w, user, err)
}
