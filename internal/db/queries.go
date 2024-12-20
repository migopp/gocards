package db

import (
	"fmt"

	"github.com/migopp/gocards/internal/debug"
)

// We have a few tables:
// I'll briefly layout their purposes and fields.
//
// I'm a complete noob, so whether this architecture is good or not,
// I have no idea. Hopefully it doesn't cause problems in the future.
//
// `users`:
//      Describes all registered users.
//      Fields:
//          - `user_id`: ID (PRIMARY)
//          - `user_name`: Username
//          - `created_at`: Timestamp
//
// `decks`:
//      Describes all created decks.
//      Fields:
//          - `deck_id`: Deck ID (PRIMARY)
//          - `user_id`: User ID (FOREIGN)
//          - `deck_name`: Deck Name
//          - `created_at`: Timestamp
//
// `cards`:
//      Describes all created flashcards.
//      Fields:
//          - `card_id`: Flashcard ID (PRIMARY)
//          - `deck_id`: Deck ID (FOREIGN)
//          - `front`: Front
//          - `back`: Back
//          - `created_at`: Timestamp

// Init tables
func tablesInit() error {
	var query string
	var err error

	// `users`
	debug.Printf("| Creating table `users`\n")
	query = "CREATE TABLE IF NOT EXISTS users(" +
		"user_id SERIAL PRIMARY KEY," +
		"user_name VARCHAR(50) NOT NULL UNIQUE," +
		"created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP" +
		");"
	if _, err = db.Exec(query); err != nil {
		return fmt.Errorf("Error creating table `users`: %v", err)
	}
	debug.Printf("| Table `users` creation successful\n")

	// `decks`
	debug.Printf("| Creating table `decks`\n")
	query = "CREATE TABLE IF NOT EXISTS decks(" +
		"deck_id SERIAL PRIMARY KEY," +
		"user_id INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE," +
		"deck_name VARCHAR(100) NOT NULL," +
		"created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP," +
		"UNIQUE (user_id, deck_name)" +
		");"
	if _, err = db.Exec(query); err != nil {
		return fmt.Errorf("Error creating table `decks`: %v", err)
	}
	debug.Printf("| Table `decks` creation successful\n")

	// `cards`
	debug.Printf("| Creating table `cards`\n")
	query = "CREATE TABLE IF NOT EXISTS cards(" +
		"card_id SERIAL PRIMARY KEY," +
		"deck_id INT NOT NULL REFERENCES decks(deck_id) ON DELETE CASCADE," +
		"front TEXT NOT NULL," +
		"back TEXT NOT NULL," +
		"created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP" +
		");"
	if _, err = db.Exec(query); err != nil {
		return fmt.Errorf("Error creating table `cards`: %v", err)
	}
	debug.Printf("| Table `cards` creation successful\n")

	// OK
	return nil
}

// Insert a user into the `users` table
func CreateUser(u IRepUser) (DBRepUser, error) {
	var dbu DBRepUser
	var err error

	query := "INSERT INTO users (user_name)" +
		"VALUES ($1)" +
		"RETURNING *;"
	debug.Printf("| Inserting `%s` into `users`\n", u.UserName)
	err = db.QueryRow(query, u.UserName).Scan(
		&dbu.UserID,
		&dbu.UserName,
		&dbu.Timestamp)
	if err != nil {
		return dbu, fmt.Errorf("Error executing `users` insertion query: %v", err)
	}
	debug.Printf("| Successful insertion into `users`\n")
	return dbu, nil
}

// Insert a deck into the `decks` table
func CreateDeck(u DBRepUser, d IRepDeck) (DBRepDeck, error) {
	var dbd DBRepDeck
	var err error

	query := "INSERT INTO decks (user_id, deck_name)" +
		"VALUES ($1, $2)" +
		"RETURNING *;"
	debug.Printf("| Inserting into `decks` with details:\n"+
		"|\t`user_id`: %d\n"+
		"|\t`deck_name`: %s\n",
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
	debug.Printf("| Successful insertion into `decks`\n")
	return dbd, err
}

// Insert a card into the `cards` table
func CreateCard(d DBRepDeck, c IRepCard) (DBRepCard, error) {
	var dbc DBRepCard
	var err error

	query := "INSERT INTO cards (deck_id, front, back)" +
		"VALUES ($1, $2, $3)" +
		"RETURNING *;"
	debug.Printf("| Inserting into `cards` with details:\n"+
		"|\t`front`: %s\n"+
		"|\t`back`: %s\n",
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
	debug.Printf("| Successful insertion into `cards`\n")
	return dbc, err
}
