// Code generated by ent, DO NOT EDIT.

package database

const (
	// Label holds the string label denoting the database type in the database.
	Label = "database"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStatusID holds the string denoting the status_id field in the database.
	FieldStatusID = "status_id"
	// FieldTagsID holds the string denoting the tags_id field in the database.
	FieldTagsID = "tags_id"
	// FieldCheckboxID holds the string denoting the checkbox_id field in the database.
	FieldCheckboxID = "checkbox_id"
	// FieldDbID holds the string denoting the db_id field in the database.
	FieldDbID = "db_id"
	// EdgeStatus holds the string denoting the status edge name in mutations.
	EdgeStatus = "status"
	// Table holds the table name of the database in the database.
	Table = "databases"
	// StatusTable is the table that holds the status relation/edge.
	StatusTable = "status"
	// StatusInverseTable is the table name for the Status entity.
	// It exists in this package in order to avoid circular dependency with the "status" package.
	StatusInverseTable = "status"
	// StatusColumn is the table column denoting the status relation/edge.
	StatusColumn = "database_status"
)

// Columns holds all SQL columns for database fields.
var Columns = []string{
	FieldID,
	FieldStatusID,
	FieldTagsID,
	FieldCheckboxID,
	FieldDbID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DbIDValidator is a validator for the "db_id" field. It is called by the builders before save.
	DbIDValidator func(string) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int) error
)
