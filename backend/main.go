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
	Host       bool   `json:"host"`
}

type Game struct {
	players map[*websocket.Conn]*Player
	Host    string
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
	host := r.URL.Query().Get("host")
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
		Points:     0,
		ActiveTurn: false,
		Host:       host == "true",
	}

	// return a list of players whenever a new connection appears
	returnLeaderboard(game)

	for {
		_, raw, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("An error occured: ", err)
			mutex.Lock()
			delete(game.players, conn)
			mutex.Unlock()
			break
		}
		base := &BaseMessage{}
		json.Unmarshal(raw, base)
		switch base.Type {
		case "chat":
			chat := &ChatMessage{}
			json.Unmarshal(raw, chat)
			for client := range game.players {
				err := client.WriteJSON(chat)
				if err != nil {
					fmt.Println("An error occured: ", err)
					delete(game.players, conn)
					break
				}
			}
		case "start":
			for client, player := range game.players {
				if player.Host {
					player.ActiveTurn = true
				}
				err := client.WriteJSON(base)
				if err != nil {
					fmt.Println("An error occured: ", err)
					delete(game.players, conn)
					break
				}
			}
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
