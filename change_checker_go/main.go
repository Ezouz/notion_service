package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"

	"gitlab.42paris.fr/notion_service/pkg/controllers"
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
			ID: os.Getenv("DB_TEST"),
			Check: []models.ProcessProperty{
				{
					Type:   "status",
					Values: []string{"complete", "Done"},
				},
			},
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
	// open connection with db
	client := controllers.PostgresClient()

	redisClientOpt := asynq.RedisClientOpt{Addr: fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))}
	srv := asynq.NewServer(redisClientOpt, asynq.Config{
		// Specify how many concurrent workers to use
		Concurrency: 10,
		// Optionally specify multiple queues with different priority.
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
		// See the godoc for other configuration options
	})

	notion.DBChangeChecker(&toCheck, client)

	mux := asynq.NewServeMux()
	runServ(srv, mux)
}

func runServ(a *asynq.Server, mux *asynq.ServeMux) {
	fmt.Printf("running asynq server...\n")
	_ = a.Run(mux)
}
