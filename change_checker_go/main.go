package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"gitlab.42paris.fr/notion_service/pkg/models"
	"gitlab.42paris.fr/notion_service/pkg/notion"
)

func main() {

	// err := godotenv.Load("../.env")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	toCheck := []models.ProcessDB{
		{
			ID:     os.Getenv("DB_TEST"),
			Check:  []string{"status"},
			Notify: os.Getenv("ZOUZ_MAIL"),
		},
		// {
		// 	ID:     os.Getenv("DB_TEST"),
		// 	Check:  []string{"status"},
		// 	Notify: os.Getenv("ZOUZ_MAIL"),
		// },
		// {
		// 	ID:     os.Getenv("DB_TEST"),
		// 	Check:  []string{"status"},
		// 	Notify: os.Getenv("ZOUZ_MAIL"),
		// },
	}

	notion.DBChangeChecker(&toCheck)
}
