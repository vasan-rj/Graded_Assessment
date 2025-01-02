package main

import (
	"ecommerce/config"
	"log"
)

func main() {
	// Initialize the database
	db, err := config.InitializeDatabase()
	if err != nil {
		log.Fatal("Failed to initialize database: ", err)
	}
	defer db.Close()

	log.Println("Database initialized successfully!")
}
