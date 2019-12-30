package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/JmPotato/Sudoku-backend/models"
	"github.com/julienschmidt/httprouter"
)

// GetPuzzleHandler processes the GET request from /puzzle/get, gets a puzzle by its pid
func GetPuzzleHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	pid, _ := strconv.ParseUint(r.Header.Get("pid"), 10, 32)
	level, _ := strconv.ParseUint(r.Header.Get("level"), 10, 8)

	puzzle := &models.Puzzle{
		Level: uint8(level),
	}
	var err error
	if pid != 0 {
		log.Printf("[GetPuzzleHandler] Finding pid=%d\n", pid)
		err = puzzle.GetPuzzleByPID(uint32(pid))
		if err != nil && err.Error() == "record not found" {
			err = puzzle.AddPuzzleByPID(uint32(pid))
		}
	} else {
		err = errors.New("No parameters")
	}

	if err != nil {
		log.Printf("[GetPuzzleHandler] Error: %s\n", err.Error())
	}

	SendResponse(w, puzzle, err)
}

// PassPuzzleHandler processes the POST request from /puzzle/pass, pass a puzzle by its pid for an uid
func PassPuzzleHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	pid, _ := strconv.ParseUint(r.Form.Get("pid"), 10, 32)
	uid, _ := strconv.ParseUint(r.Form.Get("uid"), 10, 32)
	authentication := r.Form.Get("authentication")

	puzzle := &models.Puzzle{
		PID: uint32(pid),
	}
	err = puzzle.GetPuzzleByPID(puzzle.PID)
	if err != nil {
		log.Printf("[PassPuzzleHandler] Error: %s\n", err.Error())
	}

	user := &models.User{
		UID: uint32(uid),
	}
	err = user.GetUserByUID(user.UID)
	if err != nil {
		log.Printf("[PassPuzzleHandler] Error: %s\n", err.Error())
	}
	if user.Authentication == authentication {
		log.Printf("[PassPuzzleHandler] Passing pid=%d, uid=%d\n", pid, uid)
		puzzle.Passed++

		user.Score += 81 - uint32(puzzle.Level)
		user.Passed++

		err = puzzle.SavePuzzleByPID(puzzle.PID)
		if err != nil {
			log.Printf("[PassPuzzleHandler] Error: %s\n", err.Error())
		}

		err = user.SaveUserByUID(user.UID)
		if err != nil {
			log.Printf("[PassPuzzleHandler] Error: %s\n", err.Error())
		}
	} else {
		err = errors.New("wrong authentication")
	}

	SendResponse(w, user, err)
}

// SubmitPuzzleHandler processes the POST request from /puzzle/submit, submit a puzzle by its pid for an uid
func SubmitPuzzleHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	pid, _ := strconv.ParseUint(r.Form.Get("pid"), 10, 32)
	uid, _ := strconv.ParseUint(r.Form.Get("uid"), 10, 32)
	authentication := r.Form.Get("authentication")

	puzzle := &models.Puzzle{
		PID: uint32(pid),
	}
	err = puzzle.GetPuzzleByPID(puzzle.PID)
	if err != nil {
		log.Printf("[PassPuzzleHandler] Error: %s\n", err.Error())
	}

	user := &models.User{
		UID: uint32(uid),
	}
	err = user.GetUserByUID(user.UID)
	if err != nil {
		log.Printf("[PassPuzzleHandler] Error: %s\n", err.Error())
	}
	if user.Authentication == authentication {
		log.Printf("[PassPuzzleHandler] Submitting pid=%d, uid=%d\n", pid, uid)
		puzzle.Submited++

		user.Score += 81 - uint32(puzzle.Level)
		user.Submited++

		err = puzzle.SavePuzzleByPID(puzzle.PID)
		if err != nil {
			log.Printf("[PassPuzzleHandler] Error: %s\n", err.Error())
		}

		err = user.SaveUserByUID(user.UID)
		if err != nil {
			log.Printf("[PassPuzzleHandler] Error: %s\n", err.Error())
		}
	} else {
		err = errors.New("wrong authentication")
	}

	SendResponse(w, user, err)
}
