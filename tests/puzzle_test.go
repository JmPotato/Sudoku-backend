package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/JmPotato/Sudoku-backend/handlers"
	"github.com/julienschmidt/httprouter"
)

func TestGetPuzzle(t *testing.T) {
	router := httprouter.New()
	rr := httptest.NewRecorder()

	router.GET("/puzzle", handlers.GetPuzzleHandler)

	req, _ := http.NewRequest("GET", "/puzzle", nil)
	req.Header.Add("pid", "1")

	router.ServeHTTP(rr, req)
	response := new(handlers.Response)
	json.Unmarshal(rr.Body.Bytes(), response)
	if response.Message == "record not found" {
		t.Errorf("Wrong GET puzzle result.\nGot: %s", strings.TrimSpace(rr.Body.String()))
	}
}
