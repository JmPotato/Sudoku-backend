package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JmPotato/Sudoku-backend/models"
	"github.com/julienschmidt/httprouter"
)

// UserHandler processes the request from /user/:username, check whether the username exists.
func UserHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Printf("finding username=%q", p.ByName("username"))
	var user models.User
	result := user.CheckUserName(p.ByName("username"))
	fmt.Fprintf(w, "The username %q is %t\n", p.ByName("username"), result)
}
