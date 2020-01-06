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

// SubmitPuzzleHandler processes the POST request from /puzzle/submit, submit a puzzle by its pid for an uid
func SubmitPuzzleHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	pid, _ := strconv.ParseUint(r.Form.Get("pid"), 10, 32)
	uid, _ := strconv.ParseUint(r.Form.Get("uid"), 10, 32)
	passed, _ := strconv.ParseUint(r.Form.Get("passed"), 10, 32)
	authentication := r.Form.Get("authentication")
	puzzle := &models.Puzzle{
		PID: uint32(pid),
	}
	log.Printf("[SubmitPuzzleHandler] Finding pid=%d\n", pid)
	err = puzzle.GetPuzzleByPID(puzzle.PID)
	if err != nil {
		log.Printf("[SubmitPuzzleHandler] Get Puzzle Error: %s\n", err.Error())
	}

	user := &models.User{
		UID: uint32(uid),
	}
	log.Printf("[SubmitPuzzleHandler] Finding uid=%d\n", uid)
	err = user.GetUserByUID(user.UID)
	if err != nil {
		log.Printf("[SubmitPuzzleHandler] Get User Error: %s\n", err.Error())
	}
	if user.Authentication == authentication {
		log.Printf("[SubmitPuzzleHandler] Submitting pid=%d, uid=%d\n", pid, uid)
		puzzle.Submited++
		user.Submited++
		if passed == 1 {
			puzzle.Passed++
			user.Passed++
		}
		user.Score += 81 - uint32(puzzle.Level)
		err = puzzle.SavePuzzleByPID(puzzle.PID)
		if err != nil {
			log.Printf("[SubmitPuzzleHandler] Save Puzzle Error: %s\n", err.Error())
		} else {
			err = user.SaveUserByUID(user.UID)
			if err != nil {
				log.Printf("[SubmitPuzzleHandler] Save User Error: %s\n", err.Error())
			}
		}
	} else {
		err = errors.New("Wrong authentication")
	}

	SendResponse(w, user, err)
}
