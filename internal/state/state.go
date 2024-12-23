package state

import (
	"fmt"
	"math/rand"

	"github.com/migopp/gocards/internal/types"
)

// This is some representation of the state
type State struct {
	LoadedDeck types.LRepDeck
	Index      int
	Correct    int
	Attempts   int
}

// Update the current deck for a state
func (s *State) UpdateDeck(ld types.LRepDeck) {
	GlobalState.LoadedDeck = ld
}

// Reset the state of the deck
func (s *State) Reset() {
	// Shuffle the deck
	rand.Shuffle(len(s.LoadedDeck.Cards), func(i, j int) {
		s.LoadedDeck.Cards[i], s.LoadedDeck.Cards[j] = s.LoadedDeck.Cards[j], s.LoadedDeck.Cards[i]
	})

	// Reset the stats
	s.Index = 0
	s.Correct = 0
	s.Attempts = 0
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

// Say that we made an attempt, but it was wrong
func (s *State) AddWrong() {
	s.Attempts++
}

// Say that we got one right
func (s *State) AddRight() {
	s.Correct++
	s.Attempts++
}

// Returns the correct : attempt ratio
func (s *State) Ratio() float64 {
	if s.Attempts == 0 {
		return float64(0)
	}
	return (float64(s.Correct) / float64(s.Attempts)) * 100
}

// Likely, this should actually be stored per-user,
// but I am not dealing with that complexity right now
//
// Should be simple enough to refactor later
var GlobalState State
