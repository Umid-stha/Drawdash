package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"sync"

	"github.com/gorilla/websocket"
)

type Player struct {
	Username   string `json:"username"`
	Points     int    `json:"points"`
	ActiveTurn bool   `json:"turn"`
}

type Game struct {
	players map[*websocket.Conn]*Player
}

var games = make(map[string]*Game)
var mutex = &sync.Mutex{} // Protect clients map

// Upgrader is used to upgrade HTTP connections to websocket connection
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	user := r.URL.Query().Get("user")
	game, ok := games[code]
	if !ok {
		fmt.Println("Error Game doesn't exist: ")
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading: ", err)
		return
	}

	defer func() {
		delete(game.players, conn)
		conn.Close()
		returnLeaderboard(game)
	}()

	game.players[conn] = &Player{
		Username:   user,
		Points:     100,
		ActiveTurn: false,
	}

	// return a list of players whenever a new connection appears
	returnLeaderboard(game)

	for {
		chat := &ChatMessage{}
		err := conn.ReadJSON(chat)
		if err != nil {
			fmt.Println("An error occured: ", err)
			mutex.Lock()
			delete(game.players, conn)
			mutex.Unlock()
			break
		}
		fmt.Println(chat)
		for client := range game.players {
			client.WriteJSON(chat)
		}
	}

}

type RoomCode struct {
	Code string `json:"code"`
}

func createRoom(w http.ResponseWriter, r *http.Request) {
	// Allow requests from frontend
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Allow request methods
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

	// Allow headers
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	roomCode := generateCode()

	//add new game instance
	games[roomCode] = &Game{
		players: make(map[*websocket.Conn]*Player),
	}

	response := RoomCode{roomCode}

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/create", createRoom)
	http.HandleFunc("/join/{code}", wsHandler)
	fmt.Println("Websocket server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
