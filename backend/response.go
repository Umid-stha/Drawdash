package main

type BaseMessage struct {
	Type string `json:"type"`
}

type ChatMessage struct {
	User    string `json:"user"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

type Leaderboard struct {
	Players []*Player `json:"player"`
	Type    string    `json:"type"`
}

type WordsResponse struct {
	Words []string `json:"words"`
	Type  string   `json:"type"`
}
