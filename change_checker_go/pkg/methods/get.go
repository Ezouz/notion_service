package methods

import (
	"context"
	"log"

	"gitlab.42paris.fr/notion_service/pkg/controllers"
	"gitlab.42paris.fr/notion_service/pkg/models"
)

func FetchNotionDB(DBid string) (*models.CompleteDB, error) {

	clientNotion := controllers.NotionClient()

	ctx := context.Background()

	dbInfo, err := clientNotion.FindDatabaseByID(ctx, DBid)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	page, err := clientNotion.QueryDatabase(ctx, DBid, nil)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	CDB := new(models.CompleteDB)
	CDB.DB = page
	CDB.Infos = dbInfo

	// log.Printf("Page : %v", dbInfo.Title[0].Text.Content)
	// log.Printf("CDB : %v", CDB)
	return CDB, err
}
