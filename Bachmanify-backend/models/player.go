package models

type PlayerState struct {
	ID       string   `json:"id"`
	Crypto   int      `json:"crypto"`
	Level    int      `json:"level"`
	Missions []string `json:"missions"`
}



