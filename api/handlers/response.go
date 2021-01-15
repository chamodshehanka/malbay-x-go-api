package handlers

import (
	"encoding/json"
	"net/http"
)

// RespondWithError return error message
func RespondWithError(w http.ResponseWriter, code int, msg interface{}) {

	ResponseWithJSON(w, code, map[string]interface{}{"message": msg})

}

// ResponseWithJSON write json response format
func ResponseWithJSON(w http.ResponseWriter, code int, payload interface{}) {

	t := "data"

	if code > 300 {

		t = "errors"

	}

	response, _ := json.Marshal(map[string]interface{}{
		t: payload,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}
