package methods

import (
	"context"
	"log"

	"github.com/dstotijn/go-notion"
	"gitlab.42paris.fr/utilities/notion_service/internal/models"
	"gitlab.42paris.fr/utilities/notion_service/pkg/controllers"
)

func FetchNotionPage(pageID string) (*notion.Page, error) {
	clientNotion := controllers.NotionClient()

	ctx := context.Background()

	page, err := clientNotion.FindPageByID(ctx, pageID)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return nil, err
	}
	return &page, err
}

func FetchNotionDB(DBid string) (*models.CompleteDB, error) {

	clientNotion := controllers.NotionClient()

	ctx := context.Background()

	dbInfo, err := clientNotion.FindDatabaseByID(ctx, DBid)
	// dbInfo, err := clientNotion.FindDatabaseByID(ctx, DBid)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return nil, err
	}

	page, err := clientNotion.QueryDatabase(ctx, DBid, nil)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return nil, err
	}

	CDB := new(models.CompleteDB)
	CDB.DB = page
	CDB.Infos = dbInfo

	// log.Printf("Page : %v", dbInfo.Title[0].Text.Content)
	// log.Printf("CDB : %v", CDB)
	return CDB, err
}
