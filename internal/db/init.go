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
	debug.Printf("| DB connection opened\n")

	// Test the connection
	debug.Printf("| Testing DB connection\n")
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("Could not ping database: %v", err)
	}
	debug.Printf("| DB connection success\n")

	// Init tables if not present
	debug.Printf("| Initializing DB tables\n")
	if err = tablesInit(); err != nil {
		return fmt.Errorf("Error during `tablesInit`: %v", err)
	}
	debug.Printf("| DB tables initialized\n")

	// Success!
	return nil
}

func Close() error {
	return db.Close()
}
