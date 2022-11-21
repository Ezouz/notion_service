package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/dstotijn/go-notion"
)

func NotionClient() (client *notion.Client) {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	client = notion.NewClient(os.Getenv("NOTION_TOKEN"), notion.WithHTTPClient(httpClient))

	return client
}
