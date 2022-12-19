package models

import (
	"github.com/dstotijn/go-notion"
)

type CompleteDB struct {
	DB    notion.DatabaseQueryResponse
	Infos notion.Database
}
