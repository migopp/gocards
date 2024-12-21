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
func (s *State) NextCard() error {
	if s.Index >= (len(s.LoadedDeck.Cards) - 1) {
		return fmt.Errorf("COULD NOT GET NEXT CARD")
	}
	s.Index++
	return nil
}

// Likely, this should actually be stored per-user,
// but I am not dealing with that complexity right now
//
// Should be simple enough to refactor later
var GlobalState State
