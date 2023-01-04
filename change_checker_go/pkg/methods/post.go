package methods

import (
	"context"
	"log"

	"github.com/dstotijn/go-notion"
	"gitlab.42paris.fr/utilities/notion_service/pkg/controllers"
)

// post value to notion db
func PostToNotionDB(DBid string, params notion.CreatePageParams) (*notion.Page, error) {
	clientNotion := controllers.NotionClient()
	ctx := context.Background()

	page, err := clientNotion.CreatePage(ctx, params)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return nil, err
	}

	return &page, nil
}
