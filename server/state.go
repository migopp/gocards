package server

import "github.com/migopp/gocards/db"

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
	GCS.deckStates[u.ID] = ds
}

func (ds *deckState) detach(u db.User) {
	delete(GCS.deckStates, u.ID)
}

func (ds *deckState) curr() (db.Card, bool) {
	if ds.Index >= len(ds.LoadedDeck.DBCards) {
		var c db.Card
		return c, false
	}
	return ds.LoadedDeck.DBCards[ds.Index], true
}

func (ds *deckState) next() bool {
	if ds.Index+1 >= len(ds.LoadedDeck.DBCards) {
		return false
	}
	ds.Index += 1
	return true
}

func (ds *deckState) correct() {
	ds.Correct += 1
	ds.Attempts += 1
}

func (ds *deckState) incorrect() {
	ds.Attempts += 1
}

func (ds *deckState) ratio() float64 {
	if ds.Attempts == 0 {
		return float64(0)
	}
	return float64(ds.Correct) / float64(ds.Attempts)
}
