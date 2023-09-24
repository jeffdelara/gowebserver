package api

import (
	"encoding/json"
	"gowebserver/data"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var exhibition data.Exhibition
	err := json.NewDecoder(r.Body).Decode(&exhibition)
	if err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}

	exhibitions := data.Add(exhibition)

	json.NewEncoder(w).Encode(exhibitions)
}
