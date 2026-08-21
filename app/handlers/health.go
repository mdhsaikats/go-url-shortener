package handler

import (
	"encoding/json"
	"net/http"
)

func ServerHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Server is healthy",
	})
}