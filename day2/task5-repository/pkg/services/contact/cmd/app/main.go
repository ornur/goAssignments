package main

import (
    "fmt"
	"log"
    "go/pkg/store/postgresql"
)

func main() {
	db := postgresql.Connect("local", "5432", "postgres", "12345", "goConnection")
	err := db.Ping()
	if err != nil {
		log.Fatal("Could not connect to PostgreSQL:", err)
	}
	defer db.Close()

	fmt.Println("Connected to PostgreSQL!")
}
