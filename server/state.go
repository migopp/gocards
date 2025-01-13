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
