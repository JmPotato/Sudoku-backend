package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/JmPotato/Sudoku-backend/models"
	"github.com/julienschmidt/httprouter"
)

// GetUserHandler processes the GET request from /user/get, gets a user by its uid or username.
func GetUserHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	uid, _ := strconv.ParseUint(r.Header.Get("uid"), 10, 32)
	username := r.Header.Get("username")

	user := new(models.User)
	var err error
	if uid != 0 {
		log.Printf("finding uid=%d\n", uid)
		err = user.GetUserByUID(uint32(uid))
	} else if username != "" {
		log.Printf("finding username=%s\n", username)
		err = user.GetUserByUsername(username)
	} else {
		err = errors.New("no parameters")
	}

	if err != nil {
		log.Printf("error: %s\n", err.Error())
	}

	SendResponse(w, user, err)
}

// CheckUserHandler process the POST request from /user/check, checks whether a user and its authentication are right.
func CheckUserHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	username := r.Form.Get("username")
	authentication := r.Form.Get("authentication")

	user := new(models.User)
	if username == "" || authentication == "" {
		err = errors.New("Empty username or authentication")
	} else if len(username) <= 20 && len(authentication) == 32 {
		user := models.User{Username: username, Authentication: authentication, Type: uint8(1)}
		log.Printf("Checking username=%s, type=%d\n", username, 1)
		err = user.GetUserByUsername(username)
		if user.Authentication != authentication {
			err = errors.New("Authentication doesn't match the username")
		}
	} else {
		err = errors.New("Illegal username or authentication")
	}

	if err != nil {
		log.Printf("Error: %s\n", err.Error())
	}

	SendResponse(w, user, err)
}

// CreatUserHandler processes the POST request from /user/create, creates a new user if the user doesn't exist.
func CreatUserHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	username := r.Form.Get("username")
	authentication := r.Form.Get("authentication")
	userType, _ := strconv.ParseUint(r.Form.Get("type"), 10, 8)

	user := new(models.User)
	if username == "" || authentication == "" || userType == 0 {
		err = errors.New("Empty username or authentication")
	} else if len(username) <= 20 && len(authentication) == 32 {
		user := models.User{Username: username, Authentication: authentication, Type: uint8(userType)}
		log.Printf("Creating username=%s, type=%d\n", username, userType)
		err = user.CreateUser()
	} else {
		err = errors.New("Illegal username or authentication")
	}

	if err != nil {
		log.Printf("Error: %s\n", err.Error())
	}
	user.GetUserByUsername(username)

	SendResponse(w, user, err)
}

// DeleteUserHandler processes the DELETC request from /user/delete, deletes the user by its uid or username if the user exists.
func DeleteUserHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	uid, _ := strconv.ParseUint(r.Header.Get("uid"), 10, 32)
	username := r.Header.Get("username")

	user := new(models.User)
	var err error
	if uid != 0 {
		log.Printf("Deleting uid=%d\n", uid)
		err = user.DeleteUserByUID(uint32(uid))
	} else if username != "" {
		log.Printf("Deleting username=%s\n", username)
		err = user.DeleteUserByUsername(username)
	} else {
		err = errors.New("No parameters")
	}

	if err != nil {
		log.Printf("Error: %s\n", err.Error())
	}

	SendResponse(w, user, err)
}
