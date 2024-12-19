package db

import (
	"fmt"
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

func tablesInit() error {
	var query string
	var err error

	// `users`
	query = `CREATE TABLE IF NOT EXISTS users(
        user_id SERIAL PRIMARY KEY,
        user_name VARCHAR(50) NOT NULL UNIQUE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    `
	if _, err = db.Exec(query); err != nil {
		return fmt.Errorf("Error creating `users`: %v", err)
	}

	// `decks`
	query = `CREATE TABLE IF NOT EXISTS decks(
        deck_id SERIAL PRIMARY KEY,
        user_id INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
        deck_name VARCHAR(100) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        UNIQUE (user_id, deck_name)
    );
    `
	if _, err = db.Exec(query); err != nil {
		return fmt.Errorf("Error creating `decks`: %v", err)
	}

	// `cards`
	query = `CREATE TABLE IF NOT EXISTS cards(
        card_id SERIAL PRIMARY KEY,
        deck_id INT NOT NULL REFERENCES decks(deck_id) ON DELETE CASCADE,
        FRONT TEXT NOT NULL,
        BACK TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    `
	if _, err = db.Exec(query); err != nil {
		return fmt.Errorf("Error creating `cards`: %v", err)
	}

	return err
}
