package models

const (
	waiting = iota
	full
	finished
)

type Room struct {
	Player1Id string `json:"player1Id"`
	Player2Id string `json:"player2Id"`
	Status    int    `json:"status"`
}
