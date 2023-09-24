package api

import (
	"encoding/json"
	"gowebserver/data"
	"net/http"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query()["id"]

	if id != nil {
		intId, err := strconv.Atoi(id[0])

		if err != nil {
			http.Error(w, "Invalid ID", 404)
			return
		}

		if intId >= len(data.GetAll()) {
			http.Error(w, "Out of bounds", 404)
			return
		}

		json.NewEncoder(w).Encode(data.GetAll()[intId])
		return
	}

	json.NewEncoder(w).Encode(data.GetAll())	
}
