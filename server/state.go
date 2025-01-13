package server

import (
	"fmt"

	"github.com/migopp/gocards/db"
)

type deckState struct {
	LoadedDeck db.LDeck
	Index      int
	Correct    int
	Attempts   int
}

func (ds *deckState) attach(u db.User) {
	// Reset metadata
	ds.Index = 0
	ds.Correct = 0
	ds.Attempts = 0

	// Associate on server
	GCS.deckStates[u.ID] = *ds
}

func (ds *deckState) detach(u db.User) {
	delete(GCS.deckStates, u.ID)
}

func (ds *deckState) curr() (db.Card, error) {
	if ds.Index >= len(ds.LoadedDeck.DBCards) {
		var c db.Card
		return c, fmt.Errorf("")
	}
	return ds.LoadedDeck.DBCards[ds.Index], nil
}

func (ds *deckState) next() bool {
	if ds.Index >= len(ds.LoadedDeck.DBCards) {
		return false
	}
	ds.Index++
	return true
}

func (ds *deckState) correct() {
	ds.Correct++
	ds.Attempts++
}

func (ds *deckState) incorrect() {
	ds.Attempts++
}

func (ds *deckState) ratio() float64 {
	if ds.Attempts == 0 {
		return float64(0)
	}
	return float64(ds.Correct) / float64(ds.Attempts)
}
