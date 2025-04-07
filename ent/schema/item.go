package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type Item struct {
	ent.Schema
}

// Fields of the User.
func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("name").NotEmpty(),
		field.Int("price"),
		field.String("type").NotEmpty(),
		field.String("size").Match(regexp.MustCompile(`^[0-9]+\s?(ml|l|g|kg)$`)),
		field.String("description").
			NotEmpty().
			MinLen(10).
			MaxLen(500),
		field.String("image").NotEmpty(),
		field.Time("updated_at").
			Default(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the User.
func (Item) Edges() []ent.Edge {
	return nil
}
