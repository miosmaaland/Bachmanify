package handlers

import (
	"ErlichBachmanify-backend/models"
	"ErlichBachmanify-backend/state"
	"encoding/json"
	"net/http"
)

func GetProgress(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if state, ok := state.GetState(id); ok {
		json.NewEncoder(w).Encode(state)
	} else {
		http.Error(w, "Player not found", http.StatusNotFound)
	}
}

func SaveProgress(w http.ResponseWriter, r *http.Request) {
	var ps models.PlayerState
	if err := json.NewDecoder(r.Body).Decode(&ps); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	state.SaveState(ps.ID, ps)
	w.WriteHeader(http.StatusOK)
}
