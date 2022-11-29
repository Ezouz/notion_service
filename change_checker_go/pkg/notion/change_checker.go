package notion

import (
	"context"
	"log"
	"sync"

	"gitlab.42paris.fr/notion_service/ent"
	"gitlab.42paris.fr/notion_service/pkg/methods"
	"gitlab.42paris.fr/notion_service/pkg/models"
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

	for _, db := range *DBList {
		log.Printf("db: %v", db)

		// get existing Status data
		DB, err := client.Status.Query().All(context.Background())
		if err != nil {
			log.Fatalf("failed fetching Status: %v", err)
		}
		log.Printf("status: %v", DB)
		// 	Query().
		// 	Where()
	}

	// newDB, err := client.Database.
	// Create().

	// if object in databases is archived false else delete

	// compare 2 values per row_id

	// if one of the two sets of data is referenced in the ProcessProperty.Values Array

	// save new values if changed (database_id, property type)
	// post email or discord message with change
	defer client.Close()

}

// Set field value.
// Avoid nil checks.
// Add many edges.
// Set unique edge.
// Create and return.
