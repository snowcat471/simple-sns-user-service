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
		field.String("email").Unique(),
		field.String("password"),
		field.String("name"),
		field.Enum("gender").Values("MALE", "FEMALE"),
		field.String("intro"),
		field.Time("created_at"),
		field.Time("updated_at"),
		field.Time("last_logined_at"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
