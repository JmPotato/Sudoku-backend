package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/JmPotato/Sudoku-backend/models"
	"github.com/julienschmidt/httprouter"
)

//GetPuzzleHandler processes the GET request from /puzzle, gets a puzzle by its pid
func GetPuzzleHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	pid, _ := strconv.ParseUint(r.Header.Get("pid"), 10, 32)

	puzzle := new(models.Puzzle)
	var err error
	if pid != 0 {
		log.Printf("Finding pid=%d\n", pid)
		err = puzzle.GetPuzzleByPID(uint32(pid))
	} else {
		err = errors.New("No parameters")
	}

	if err != nil {
		log.Printf("Error: %s\n", err.Error())
	}

	SendResponse(w, puzzle, err)
}
