// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"gitlab.42paris.fr/notion_service/ent/database"
	"gitlab.42paris.fr/notion_service/ent/schema"
	"gitlab.42paris.fr/notion_service/ent/status"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	databaseFields := schema.Database{}.Fields()
	_ = databaseFields
	// databaseDescDbID is the schema descriptor for db_id field.
	databaseDescDbID := databaseFields[4].Descriptor()
	// database.DbIDValidator is a validator for the "db_id" field. It is called by the builders before save.
	database.DbIDValidator = databaseDescDbID.Validators[0].(func(string) error)
	// databaseDescID is the schema descriptor for id field.
	databaseDescID := databaseFields[0].Descriptor()
	// database.IDValidator is a validator for the "id" field. It is called by the builders before save.
	database.IDValidator = databaseDescID.Validators[0].(func(int) error)
	statusFields := schema.Status{}.Fields()
	_ = statusFields
	// statusDescDbID is the schema descriptor for db_id field.
	statusDescDbID := statusFields[1].Descriptor()
	// status.DbIDValidator is a validator for the "db_id" field. It is called by the builders before save.
	status.DbIDValidator = statusDescDbID.Validators[0].(func(string) error)
	// statusDescRowID is the schema descriptor for row_id field.
	statusDescRowID := statusFields[2].Descriptor()
	// status.RowIDValidator is a validator for the "row_id" field. It is called by the builders before save.
	status.RowIDValidator = statusDescRowID.Validators[0].(func(string) error)
	// statusDescStatus is the schema descriptor for status field.
	statusDescStatus := statusFields[3].Descriptor()
	// status.StatusValidator is a validator for the "status" field. It is called by the builders before save.
	status.StatusValidator = statusDescStatus.Validators[0].(func(string) error)
	// statusDescSavedAt is the schema descriptor for saved_at field.
	statusDescSavedAt := statusFields[4].Descriptor()
	// status.DefaultSavedAt holds the default value on creation for the saved_at field.
	status.DefaultSavedAt = statusDescSavedAt.Default.(func() time.Time)
	// statusDescID is the schema descriptor for id field.
	statusDescID := statusFields[0].Descriptor()
	// status.IDValidator is a validator for the "id" field. It is called by the builders before save.
	status.IDValidator = statusDescID.Validators[0].(func(int) error)
}
