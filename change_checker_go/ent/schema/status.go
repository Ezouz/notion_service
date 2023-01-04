package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Status holds the schema definition for the Status entity.
type Status struct {
	ent.Schema
}

// Fields of the Status.
func (Status) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Unique().Immutable(),
		field.String("db_id").MaxLen(200),
		field.String("row_id").MaxLen(200),
		field.String("status").MaxLen(99),
		field.Time("saved_at").Default(time.Now),
	}
}

// Edges of the Status.
func (Status) Edges() []ent.Edge {
	return nil
}
