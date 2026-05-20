package main

import (
	"math/rand/v2"
)

// Generates a random sequence of 6 characters string for room code
func generateCode() string {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	code := make([]byte, 6)

	for i := range code {
		code[i] = chars[rand.Int32N(int32(len(chars)))]
	}

	return string(code)
}

// Returns the current leaderboard status to all clients in room
func returnLeaderboard(game *Game) {
	var players []*Player
	for _, player := range game.players {
		players = append(players, player)
	}
	leaderboard := Leaderboard{
		Players: players,
		Type:    "leaderboard",
	}
	for client := range game.players {
		client.WriteJSON(leaderboard)
	}
}
