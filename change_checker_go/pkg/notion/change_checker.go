package notion

import (
	"sync"

	"gitlab.42paris.fr/notion_service/pkg/methods"
	"gitlab.42paris.fr/notion_service/pkg/models"
)

type Databases struct {
	data    *models.CompleteDB
	actions *models.ProcessDB
	err     error
}

func DBChangeChecker(DBList *[]models.ProcessDB) {
	// log.Printf("toCheck : %v", DBList)

	var wg sync.WaitGroup
	content := make(chan Databases, 300) // Declare a unbuffered channel
	var databases []Databases

	// with goroutines
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

	// log.Printf("databases: %v", databases)
	wg.Wait()
	// finished writing on channel
	defer close(content)

	// get db in postgres and values to compare (database_id, property type)
	// check content change if one of the two sets of data is ["down", "deprecated", "maintenance", "done"]
	// save new values if changed (database_id, property type)
	// post email or discord message with change
}
