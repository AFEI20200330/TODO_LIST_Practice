package handler

import (
	"encoding/json"
	"net/http"
)

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	response,err := json.Marshal(data)
	if err!= nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func respondErr(w http.ResponseWriter, code int, mag string) {
	respondJSON(w, code, map[string]string{"error":mag})
}

