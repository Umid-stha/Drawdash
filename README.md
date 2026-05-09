# DrawDash

A real-time multiplayer drawing and guessing game built with a Go backend and vanilla JavaScript frontend. Players create rooms, take turns drawing, and guess each other's drawings in a fast-paced competition.

## Features

### Core Gameplay
- **Create/Join Rooms** - Players can create new game rooms or join existing ones using room codes
- **Turn-Based Gameplay** - Players take turns drawing while others guess
- **Real-time Drawing Sync** - Live drawing updates streamed to all players in the room
- **Real-time Guessing/Chat** - Players can submit guesses and chat messages that appear instantly
- **Word Selection System** - Drawers choose from multiple word options each round
- **Round Timer** - Configurable timer for each drawing round with countdown display
- **Score Tracking** - Points awarded for correct guesses and drawing performance
- **Player List + Leaderboards** - View active players and final game statistics

## Tech Stack

### Frontend
- **HTML** - Semantic markup for game interface
- **CSS** - Responsive styling for desktop and tablet displays
- **JavaScript (Vanilla)** - Game logic, canvas drawing, and WebSocket communication
- **Canvas API** - Drawing board functionality

### Backend
- **Go** - High-performance server with concurrent connection handling
- **WebSocket** - Real-time bidirectional communication between clients and server
- **JSON** - Data serialization for game state and messages

## Project Structure

```
DrawDash/
├── README.md                 # Project documentation
├── frontend/
│   ├── index.html           # Main game interface
│   ├── app.js               # Game logic and client-side state management
│   └── style.css            # Styling for game UI
└── backend/
    ├── main.go              # Server entry point
    ├── game/                # Game logic and room management
    ├── websocket/           # WebSocket connection handling
    └── models/              # Data structures for game state
```

## Getting Started

### Prerequisites
- Go 1.16 or higher
- A modern web browser supporting WebSocket and Canvas API

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd DrawDash
   ```

2. **Backend Setup**
   ```bash
   cd backend
   go mod download
   ```

3. **Frontend Setup**
   - Frontend files are served directly by the Go backend
   - No additional build step required

### Running the Application

1. **Start the backend server**
   ```bash
   cd backend
   go run main.go
   ```
   - Server will start on `http://localhost:8080`

2. **Open in browser**
   - Navigate to `http://localhost:8080`
   - Create a new room or join an existing one

## Game Rules

1. **Joining a Game**
   - Create a new room or join with room code
   - Wait for minimum players to start

2. **Drawing Phase**
   - One player draws based on the selected word
   - 60-second timer per round
   - Drawer cannot see who has guessed correctly

3. **Guessing Phase**
   - Other players submit guesses
   - First correct guess earns bonus points
   - Correct guesses are hidden from other players

4. **Scoring**
   - Correct guess: Base points + time bonus
   - Successful drawing: Points based on number of correct guesses

5. **Round Progression**
   - Turns rotate to the next player
   - Game continues for configured number of rounds
   - Final leaderboard displayed at game end

## WebSocket Events

### Client → Server
- `create_room` - Create a new game room
- `join_room` - Join an existing room
- `select_word` - Drawer selects word to draw
- `draw` - Send drawing stroke data
- `guess` - Submit a guess
- `chat` - Send chat message
- `start_game` - Host starts the game

### Server → Client
- `room_created` - New room created with code
- `player_joined` - Player joined the room
- `room_state` - Current game state update
- `drawing_stroke` - Incoming stroke from drawer
- `guess_received` - Guess submitted by player
- `round_start` - New round begins with word options
- `round_end` - Round finished, scores calculated
- `game_end` - Game completed, final results

## Configuration

- `PORT` - Server port (default: 8080)
- `ROUND_TIME` - Seconds per drawing round (default: 60)
- `MIN_PLAYERS` - Minimum players to start (default: 2)
- `ROUNDS_TO_PLAY` - Total rounds per game (default: 3)

## Browser Support

- Chrome 90+
- Firefox 88+
- Safari 14+
- Edge 90+

## License

MIT License

## Contributing

Contributions are welcome! Please fork and submit pull requests.

---

**Last Updated:** May 2026
