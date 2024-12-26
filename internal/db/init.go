package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
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

func Init() error {
	debug.Printf("| DB Init\n")

	// Get env variables to spin up DB
	user := os.Getenv("DB_USER")
	passwd := os.Getenv("DB_PASSWD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	// Spin up DB
	debug.Printf("| Opening DB connection\n")
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, passwd, host, port, name)
	debug.Printf("| Connection string: %s\n", connStr)
	tdb, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("Error opening database connection: %v", err)
	}
	db = tdb

	// Test the connection
	debug.Printf("| Testing DB connection\n")
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("Could not ping database: %v", err)
	}

	// Success!
	return nil
}

func Close() error {
	return db.Close()
}
