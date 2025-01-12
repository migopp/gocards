package db

import "mime/multipart"

type LDeck struct {
	DBDeck  Deck
	DBCards []Card
}

// `YMLToDeck` parses a `.yml` -> deck to load into DB
//
// This means that `Deck` in `LDeck.DBDeck` is populated with:
//   - `Name`
//   - `UserID`
//   - `User`
//
// Likewise, each `Card` in `LDeck.DBCards` is populated with:
//   - `Front`
//   - `Back`
//   - `DeckID`
//   - `Deck`
func YMLToDeck(f multipart.File, h *multipart.FileHeader) (LDeck, error) {
	// TODO:
	var ld LDeck
	var err error
	return ld, err
}
