package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.String("firstName").NotEmpty(),
		field.String("lastName").NotEmpty(),
		field.String("email").NotEmpty(),
		field.String("password").Sensitive(),
		field.Time("DOB").Default(time.Now()),
		field.Enum("bloodGroup").Values("APOS", "ANEG", "BPOS", "BNEG", "ABPOS", "ABNEG", "OPOS", "ONEG").Optional(),
		field.Float32("weight").Optional(),
		field.Float32("height").Optional(),
		field.Bool("isArchived").Default(false),
		field.Bool("isVerified").Default(false),
		field.Time("createdAt").Default(time.Now()),
		field.Time("updatedAt").Default(time.Now()),
		field.Time("archivedAt").Nillable(),
		field.Time("verifiedAt").Nillable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
