package handlers

import (
	"fmt"
	"net/http"
)

// HomeHandler processes the request from /
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The very first begining of Sudoku project!\n")
}
