package controllers

import (
	"context"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gitlab.42paris.fr/notion_service/ent"
)

type ConnString struct {
	Host     string
	DB       string
	Port     string
	Password string
	User     string
}

// This function will make a connection to the database only once.
func PostgresClient(connstring ConnString) *ent.Client {
	var err error

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", connstring.Host, connstring.Port, connstring.User, connstring.Password, connstring.DB)
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
