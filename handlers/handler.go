package handlers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendResponse sends the result and some other info to requester.
func SendResponse(w http.ResponseWriter, data interface{}, err error) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var res Response
	if err != nil {
		res = Response{Message: err.Error(), Data: nil}
	} else {
		res = Response{Message: "Success", Data: data}
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		return err
	}
	return nil
}
