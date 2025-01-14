package db

import (
	"fmt"
	"mime/multipart"

	"gopkg.in/yaml.v3"
)

type LDeck struct {
	DBDeck  Deck   `yaml:"deck"`
	DBCards []Card `yaml:"cards"`
}

// `YMLToLDeck` parses a `.yml` -> deck to load into DB
//
// This means that `Deck` in `LDeck.DBDeck` is populated with:
//   - `Name`
//
// Likewise, each `Card` in `LDeck.DBCards` is populated with:
//   - `Front`
//   - `Back`
//
// Some of the data, such as `Deck.UserID` or `Card.DeckID` will
// be loaded into the load representation at a later, more convenient
// time.
func YMLToLDeck(f multipart.File, h *multipart.FileHeader) (LDeck, error) {
	var ld LDeck
	var err error

	// Parse the `.yml` -> `LDeck`
	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(&ld); err != nil {
		return ld, fmt.Errorf("Cannot decode file: %s", h.Filename)
	}

	// Later, we load DB metadata fields...

	return ld, nil
}

// `DeckToLDeck` makes database queries to take a `Deck` DB reference
// and turn it into a fully-loaded form in memory that the server
// can use and manipulate
//
// This way we can store them in some per-user state on the server
// side rather than bouncing it around in a JWT or something
func DeckToLDeck(d Deck) (LDeck, error) {
	var ld LDeck
	// NOTE: Consider upgrading this to be a method later...
	//
	// This is maybe tricky because we are assuming
	// we want this operation done on the global DB, instead
	// of making this a function of an individual DB
	cards, err := GCDB.FetchCardsForDeck(d)
	if err != nil {
		return ld, err
	}
	ld.DBDeck = d
	ld.DBCards = cards
	return ld, nil
}

// Loads a given `LDeck` into the DB
//
// Basically a wrapper around a call to `CreateDeck` and
// a bunch of `CreateCard`
func (db *DB) LoadDeck(ld *LDeck, u User) error {
	// Fill `DBDeck` metadata
	ld.DBDeck.UserID = u.ID
	ld.DBDeck.User = u

	// Put `ld.DBDeck` into DB
	if err := db.CreateDeck(&ld.DBDeck); err != nil {
		return err
	}

	// Populate cards
	for i := range ld.DBCards {
		// Get card
		card := &ld.DBCards[i]

		// Fill `card` metadata
		card.DeckID = ld.DBDeck.ID
		card.Deck = ld.DBDeck

		// Put `card` into DB
		if err := db.CreateCard(card); err != nil {
			return err
		}
	}

	return nil
}
