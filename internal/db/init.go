package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/migopp/gocards/internal/debug"
)

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
