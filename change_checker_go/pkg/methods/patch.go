package methods

import (
	"context"
	"log"

	"github.com/dstotijn/go-notion"
	"gitlab.42paris.fr/utilities/notion_service/pkg/controllers"
)

// post value to notion db
func UpdateNotionPage(PageID string, params notion.UpdatePageParams) (*notion.Page, error) {
	clientNotion := controllers.NotionClient()
	ctx := context.Background()

	page, err := clientNotion.UpdatePage(ctx, PageID, params)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return nil, err
	}

	return &page, nil
}
