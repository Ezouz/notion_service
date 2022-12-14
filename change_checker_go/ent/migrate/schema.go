// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DatabasesColumns holds the columns for the "databases" table.
	DatabasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "status_id", Type: field.TypeInt},
		{Name: "tags_id", Type: field.TypeInt},
		{Name: "checkbox_id", Type: field.TypeInt},
		{Name: "db_id", Type: field.TypeString, Size: 200},
	}
	// DatabasesTable holds the schema information for the "databases" table.
	DatabasesTable = &schema.Table{
		Name:       "databases",
		Columns:    DatabasesColumns,
		PrimaryKey: []*schema.Column{DatabasesColumns[0]},
	}
	// StatusColumns holds the columns for the "status" table.
	StatusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "db_id", Type: field.TypeString, Size: 200},
		{Name: "row_id", Type: field.TypeString, Size: 200},
		{Name: "status", Type: field.TypeString, Size: 99},
		{Name: "saved_at", Type: field.TypeTime},
		{Name: "database_status", Type: field.TypeInt, Nullable: true},
	}
	// StatusTable holds the schema information for the "status" table.
	StatusTable = &schema.Table{
		Name:       "status",
		Columns:    StatusColumns,
		PrimaryKey: []*schema.Column{StatusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "status_databases_status",
				Columns:    []*schema.Column{StatusColumns[5]},
				RefColumns: []*schema.Column{DatabasesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DatabasesTable,
		StatusTable,
	}
)

func init() {
	StatusTable.ForeignKeys[0].RefTable = DatabasesTable
}
