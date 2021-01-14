package handlers

import (
	"encoding/json"
	"net/http"
)

// RespondWithError return error message
func RespondWithError(w http.ResponseWriter, code int, msg interface{}) {

	RespondwithJSON(w, code, map[string]interface{}{"message": msg})

}

// RespondwithJSON write json response format
func RespondwithJSON(w http.ResponseWriter, code int, payload interface{}) {

	t := "data"

	if code > 300 {

		t = "errors"

	}

	response, _ := json.Marshal(map[string]interface{}{
		t: payload,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
