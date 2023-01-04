package notion

import (
	"context"
	"log"
	"sync"

	"gitlab.42paris.fr/utilities/notion_service/ent"
	"gitlab.42paris.fr/utilities/notion_service/internal/models"
	"gitlab.42paris.fr/utilities/notion_service/pkg/methods"
)

type Databases struct {
	data    *models.CompleteDB
	actions *models.ProcessDB
	err     error
}

func DBChangeChecker(DBList *[]models.ProcessDB, client *ent.Client) {

	// with goroutines
	var wg sync.WaitGroup
	content := make(chan Databases, 300) // Declare a unbuffered channel
	var databases []Databases

	// get db in notion
	for _, db := range *DBList {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine
		go func(DB *models.ProcessDB, channel chan Databases) {
			// Decrement the counter when the goroutine completes.
			// TODO iterate with cursor if page.DB.HasMore
			page, err := methods.FetchNotionDB(DB.ID)
			if err != nil {
				return
			}
			channel <- Databases{data: page, actions: DB, err: err}
			wg.Done()
		}(&db, content)

		databases = append(databases, <-content)
	}

	wg.Wait()
	// finished writing on channel
	defer close(content)

	// get db in postgres and values to compare (database_id, property type)
	var statusList []*ent.Status

	for _, db := range *DBList {
		log.Printf("db: %v", db)

		// get existing Status data
		for _, toCheck := range db.Check {
			if (toCheck.Type == "status") && (len(statusList) == 0) {
				DB, err := client.Status.Query().All(context.Background())
				if err != nil {
					log.Fatalf("failed fetching Status: %v", err)
				}
				statusList = DB
			}
		}
	}

	// compare 2 values per row_id
	log.Printf("status: %v", statusList)

	for _, db := range *DBList {
		log.Printf("db: %v", db)

		// get existing Status data
		for _, toCheck := range db.Check {
			if (toCheck.Type == "status") && (len(statusList) == 0) {

			}
		}
	}

	// newDB, err := client.Database.
	// Create().

	// if object in databases is archived false else delete

	// if one of the two sets of data is referenced in the ProcessProperty.Values Array

	// save new values if changed (database_id, property type)
	// post email or discord message with change
	defer client.Close()
}
