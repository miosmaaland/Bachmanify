package state

import "ErlichBachmanify-backend/models"

var PlayerStore = map[string]models.PlayerState{}

func SaveState(id string, state models.PlayerState) {
	PlayerStore[id] = state
}

func GetState(id string) (models.PlayerState, bool) {
	state, exists := PlayerStore[id]
	return state, exists
}
