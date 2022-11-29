package controllers

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gitlab.42paris.fr/notion_service/ent"
)

// This function will make a connection to the database only once.
func PostgresClient() *ent.Client {
	var err error

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	// "postgres://postgres:password@localhost/DB_1?sslmode=disable"
	postgres, err := ent.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}
	// this will be printed in the terminal, confirming the connection to the database
	log.Printf("The database is connected: %v", postgres)

	if err := postgres.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return postgres
}
