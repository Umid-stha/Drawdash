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

var wordsCollection []string = []string{
	// animals
	"cat", "dog", "fish", "bird", "lion", "bear", "frog", "duck", "wolf", "crab",
	"shark", "horse", "snake", "eagle", "panda", "koala", "tiger", "rabbit", "turtle", "penguin",
	// food
	"pizza", "burger", "apple", "bread", "cake", "taco", "sushi", "donut", "grape", "lemon",
	"carrot", "cookie", "banana", "cheese", "hotdog", "mango", "onion", "waffle", "pretzel", "popcorn",
	// nature
	"tree", "moon", "star", "sun", "rain", "snow", "cloud", "river", "ocean", "island",
	"desert", "forest", "flower", "volcano", "rainbow", "mountain", "cactus", "waterfall", "tornado", "glacier",
	// objects
	"book", "chair", "clock", "phone", "piano", "sword", "crown", "ladder", "candle", "mirror",
	"camera", "rocket", "anchor", "compass", "lantern", "umbrella", "telescope", "parachute", "boomerang", "magnifier",
	// places
	"castle", "bridge", "museum", "lighthouse", "pyramid", "igloo", "windmill", "skyscraper", "cathedral", "stadium",
	// actions / concepts
	"sleep", "dance", "swim", "climb", "laugh", "dream", "fall", "jump", "ghost", "magic",
	// hard concepts
	"entropy", "paradox", "eclipse", "labyrinth", "quicksand", "avalanche", "illusion", "conspiracy", "blackhole", "dimension",
	// science / tech
	"telescope", "microscope", "chromosome", "algorithm", "submarine", "satellite", "radiation", "magnetism", "photosynthesis", "evolution",
	// history / mythology
	"gladiator", "pharaoh", "mythology", "colosseum", "spartacus", "odyssey", "minotaur", "excalibur", "apocalypse", "renaissance",
	// sports / activities
	"skateboard", "snorkeling", "archery", "fencing", "wrestling", "gymnastics", "polo", "bobsled", "paragliding", "sumo",
	// musicians / artists
	"beethoven", "mozart", "picasso", "davinci", "michaelangelo", "rembrandt", "shakespeare", "chopsticks", "bach", "caravaggio",
	// rappers
	"drake", "eminem", "kendrick", "tupac", "biggie", "snoop", "jayz", "travis", "kanye", "nicki",
	// celebrities / pop culture
	"einstein", "napoleon", "cleopatra", "tesla", "darwin", "freud", "aristotle", "socrates", "galileo", "newton",
	// movies / shows
	"inception", "titanic", "matrix", "godfather", "interstellar", "joker", "avengers", "batman", "sherlock", "naruto",
	// hard objects
	"guillotine", "catapult", "periscope", "sarcophagus", "sundial", "abacus", "trebuchet", "hieroglyph", "kaleidoscope", "hourglass",
	// misc hard
	"vertigo", "insomnia", "solitude", "nostalgia", "epiphany", "déjà vu", "serendipity", "schadenfreude", "wanderlust", "euphoria",
}

// Generate a list of random 4 words
func generateWords() []string {
	words := make([]string, 4)

	for i := range words {
		words[i] = wordsCollection[rand.Int32N(int32(len(wordsCollection)))]
	}
	return words
}
