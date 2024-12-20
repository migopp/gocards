package state

import (
	"github.com/migopp/gocards/internal/types"
)

// This is some representation of the state
type State struct {
	LoadedDeck types.LRepDeck
}

// Update the current deck for a state
func (s *State) UpdateDeck(ld types.LRepDeck) {
	GlobalState.LoadedDeck = ld
}

// Likely, this should actually be stored per-user,
// but I am not dealing with that complexity right now
//
// Should be simple enough to refactor later
var GlobalState State
