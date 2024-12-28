package db

import (
	"fmt"

	"github.com/migopp/gocards/internal/debug"
	"github.com/migopp/gocards/internal/types"
)

// Insert a user into the `users` table
func CreateUser(u types.IRepUser) (types.DBRepUser, error) {
	var dbu types.DBRepUser
	var err error
	query := "INSERT INTO users (user_name)" +
		"VALUES ($1)" +
		"RETURNING *;"
	debug.Printf("Inserting into `users` with details:\n"+
		"\t`user_name`: %s\n",
		u.UserName)
	err = db.QueryRow(query, u.UserName).Scan(
		&dbu.UserID,
		&dbu.UserName,
		&dbu.Timestamp)
	if err != nil {
		return dbu, fmt.Errorf("Error executing `users` insertion query: %v", err)
	}
	debug.Printf("Successful insertion into `users`\n")
	return dbu, nil
}

// Insert a deck into the `decks` table
func CreateDeck(u types.DBRepUser, d types.IRepDeck) (types.DBRepDeck, error) {
	var dbd types.DBRepDeck
	var err error
	query := "INSERT INTO decks (user_id, deck_name)" +
		"VALUES ($1, $2)" +
		"RETURNING *;"
	debug.Printf("Inserting into `decks` with details:\n"+
		"\t`user_id`: %d\n"+
		"\t`deck_name`: %s\n",
		u.UserID,
		d.DeckName)
	err = db.QueryRow(query, u.UserID, d.DeckName).Scan(
		&dbd.DeckID,
		&dbd.UserID,
		&dbd.DeckName,
		&dbd.Timestamp)
	if err != nil {
		return dbd, fmt.Errorf("Error executing `decks` insertion query: %v", err)
	}
	debug.Printf("Successful insertion into `decks`\n")
	return dbd, err
}

// Insert a card into the `cards` table
func CreateCard(d types.DBRepDeck, c types.IRepCard) (types.DBRepCard, error) {
	var dbc types.DBRepCard
	var err error
	query := "INSERT INTO cards (deck_id, front, back)" +
		"VALUES ($1, $2, $3)" +
		"RETURNING *;"
	debug.Printf("Inserting into `cards` with details:\n"+
		"\t`front`: %s\n"+
		"\t`back`: %s\n",
		c.Front,
		c.Back)
	err = db.QueryRow(query, d.DeckID, c.Front, c.Back).Scan(
		&dbc.CardID,
		&dbc.DeckID,
		&dbc.Front,
		&dbc.Back,
		&dbc.Timestamp)
	if err != nil {
		return dbc, fmt.Errorf("Error executing `cards` insertion query: %v", err)
	}
	debug.Printf("Successful insertion into `cards`\n")
	return dbc, err
}
