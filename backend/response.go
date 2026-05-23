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
	Players    []*Player `json:"player"`
	GameStatus bool      `json:"gameStatus"`
	Type       string    `json:"type"`
}

type WordsResponse struct {
	Words []string `json:"words"`
	Type  string   `json:"type"`
}

type Drawing struct {
	StartingX string `json:"startingX"`
	StartingY string `json:"startingY"`
	X         string `json:"x"`
	Y         string `json:"y"`
	Pen       string `json:"pen"`
	Color     string `json:"color"`
	LineWidth string `json:"lineWidth"`
	Type      string `json:"type"`
}
