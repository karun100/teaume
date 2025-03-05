package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("full_name").NotEmpty(),
		field.String("email").Unique().NotEmpty(),
		field.String("username").Unique().NotEmpty(),
		field.String("password").Sensitive().NotEmpty(),
		field.Enum("role").
			Values("Seller", "User").
			Default("User"),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
