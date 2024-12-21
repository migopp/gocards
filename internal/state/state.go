package state

import (
	"github.com/migopp/gocards/internal/types"
)

// This is some representation of the state
type State struct {
	LoadedDeck types.LRepDeck
	Index      uint16
}

// Update the current deck for a state
func (s *State) UpdateDeck(ld types.LRepDeck) {
	GlobalState.LoadedDeck = ld
	GlobalState.Index = 0
}

// Get `front` of the current card
// TODO: Error handling
func (s *State) GetFront() string {
	return s.LoadedDeck.Cards[s.Index].Front
}

// Get `back` of the current card
// TODO: Error handling
func (s *State) GetBack() string {
	return s.LoadedDeck.Cards[s.Index].Back
}

// Iterates to the next card
// TODO: Error handling
func (s *State) NextCard() {
	s.Index++
}

// Likely, this should actually be stored per-user,
// but I am not dealing with that complexity right now
//
// Should be simple enough to refactor later
var GlobalState State
