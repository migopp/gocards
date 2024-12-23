package server

import (
	"github.com/migopp/gocards/internal/types"
)

// An empty interface to satisfy compiler
type DynContent interface{}

// Dynamic HTML content for use during the game
type GameDynContent struct {
	Word  string
	Ratio float64
}

// Dynamic HTML content for use in the home screen
type HomeDynContent struct {
	Decks []types.LRepDeck
}
