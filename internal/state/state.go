package state

import (
	"fmt"

	"github.com/migopp/gocards/internal/types"
)

// This is some representation of the state
type State struct {
	LoadedDeck types.LRepDeck
	Index      int
}

// Update the current deck for a state
func (s *State) UpdateDeck(ld types.LRepDeck) {
	GlobalState.LoadedDeck = ld
	GlobalState.Index = 0
}

// Get `front` of the current card
func (s *State) GetFront() (string, error) {
	if s.Index >= len(s.LoadedDeck.Cards) {
		return "", fmt.Errorf("OUT OF BOUNDS CARD [%d]", s.Index)
	}
	return s.LoadedDeck.Cards[s.Index].Front, nil
}

// Get `back` of the current card
func (s *State) GetBack() (string, error) {
	if s.Index >= len(s.LoadedDeck.Cards) {
		return "", fmt.Errorf("OUT OF BOUNDS CARD [%d]", s.Index)
	}
	return s.LoadedDeck.Cards[s.Index].Back, nil
}

// Iterates to the next card
//
// `true` if the iteration was successful,
// `false` otherwise
func (s *State) NextCard() bool {
	if s.Index >= (len(s.LoadedDeck.Cards) - 1) {
		return false
	}
	s.Index++
	return true
}

// Likely, this should actually be stored per-user,
// but I am not dealing with that complexity right now
//
// Should be simple enough to refactor later
var GlobalState State
