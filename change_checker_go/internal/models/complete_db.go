package models

import (
	"github.com/dstotijn/go-notion"
)

type CompleteDB struct {
	Infos notion.Database
	DB    notion.DatabaseQueryResponse
}
