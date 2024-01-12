package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sultan/museum/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query()["id"]
	if id != nil { // We want only 1
		finalId, err := strconv.Atoi(id[0])
		if err == nil && finalId < len(data.GetAll()) {
			json.NewEncoder(w).Encode(data.GetAll()[finalId])
		} else {
			http.Error(w, "Invalid Expedition", http.StatusBadRequest)
		}
	} else { // We want all
		json.NewEncoder(w).Encode(data.GetAll())
	}
}
