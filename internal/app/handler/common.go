package handler

import (
	"encoding/json"
	"net/http"
)

// respose error
func respondError(w http.ResponseWriter, r *http.Request, code int, err error) {
	respondJson(w, r, code, map[string]string{"error": err.Error()})
}

// response json
func respondJson(w http.ResponseWriter, r *http.Request, code int, data interface{}) {

	// set header status
	w.WriteHeader(code)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
