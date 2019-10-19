package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/JmPotato/Sudoku-backend/handlers"
	"github.com/JmPotato/Sudoku-backend/utils"
	"github.com/julienschmidt/httprouter"
)

func TestGetUser(t *testing.T) {
	router := httprouter.New()
	rr := httptest.NewRecorder()

	router.GET("/user", handlers.GetUserHandler)

	req, _ := http.NewRequest("GET", "/user", nil)
	req.Header.Add("username", "test")

	router.ServeHTTP(rr, req)
	response := new(handlers.Response)
	json.Unmarshal(rr.Body.Bytes(), response)
	if response.Message != "record not found" {
		t.Errorf("Wrong GET user result.\nGot: %s", strings.TrimSpace(rr.Body.String()))
	}
}

func TestPostUser(t *testing.T) {
	router := httprouter.New()
	rr := httptest.NewRecorder()

	router.POST("/user", handlers.CreatUserHandler)

	req, _ := http.NewRequest("POST", "/user", nil)
	req.ParseForm()
	req.Form.Add("type", "1")
	req.Form.Add("username", "test")
	req.Form.Add("authentication", utils.HashToMD5("test"))

	router.ServeHTTP(rr, req)
	response := new(handlers.Response)
	json.Unmarshal(rr.Body.Bytes(), response)
	if response.Message != "Success" {
		t.Errorf("Wrong POST user result.\nGot: %s", strings.TrimSpace(rr.Body.String()))
	}
}

func TestUser(t *testing.T) {
	router := httprouter.New()
	rr := httptest.NewRecorder()

	router.DELETE("/user", handlers.DeleteUserHandler)

	req, _ := http.NewRequest("DELETE", "/user", nil)
	req.Header.Add("username", "test")

	router.ServeHTTP(rr, req)
	response := new(handlers.Response)
	json.Unmarshal(rr.Body.Bytes(), response)
	if response.Message != "Success" {
		t.Errorf("Wrong DELETE user result.\nGot: %s", strings.TrimSpace(rr.Body.String()))
	}
}
