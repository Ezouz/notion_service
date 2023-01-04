// go run -mod=mod entgo.io/ent/cmd/ent generate --target ./ent ./ent/schema/notion
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Database holds the schema definition for the Database entity.
type Database struct {
	ent.Schema
}

// Fields of the Database.
func (Database) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Unique().Immutable(),
		field.Int("status_id").Nillable(),
		field.Int("tags_id").Nillable(),
		field.Int("checkbox_id").Nillable(),
		field.String("db_id").MaxLen(200),
	}
}

// Edges of the Database.
func (Database) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("status", Status.Type),
	}
}
