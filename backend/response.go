package main

type ChatMessage struct {
	User    string `json:"user"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

type Leaderboard struct {
	Players []*Player `json:"player"`
	Type    string    `json:"type"`
}
