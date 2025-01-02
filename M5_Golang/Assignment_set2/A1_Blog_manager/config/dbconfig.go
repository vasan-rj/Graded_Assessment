package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DatabaseConnection *sql.DB

func SetupDatabase() error {
	var setupErr error
	DatabaseConnection, setupErr = sql.Open("sqlite3", "./blogs.db")
	if setupErr != nil {
		return fmt.Errorf("error connecting to the database: %v", setupErr)
	}

	// Verify the database connection
	if pingErr := DatabaseConnection.Ping(); pingErr != nil {
		return fmt.Errorf("database connection verification failed: %v", pingErr)
	}

	// Create Blogs table if it doesn't exist
	_, execErr := DatabaseConnection.Exec(`CREATE TABLE IF NOT EXISTS blogs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		author TEXT NOT NULL,
		timestamp TEXT NOT NULL
	);`)
	if execErr != nil {
		return fmt.Errorf("error creating the blogs table: %v", execErr)
	}

	log.Println("Database setup completed: Connection successful and schema validated.")
	return nil
}

func GetDatabaseConnection() *sql.DB {
	if DatabaseConnection == nil {
		log.Fatal("Database connection has not been initialized.")
	}
	return DatabaseConnection
}
