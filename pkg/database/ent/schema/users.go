package schema

import (
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
		field.String("account_id"),
		field.String("display_name"),

		field.String("access_token"),
		field.Time("access_token_expires_at"),

		field.String("refresh_token"),
		field.Time("refresh_token_expires_at"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
