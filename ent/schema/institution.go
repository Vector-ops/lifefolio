package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Institution holds the schema definition for the Institution entity.
type Institution struct {
	ent.Schema
}

// Fields of the Institution.
func (Institution) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.String("name"),
		field.Text("description").Nillable(),
		field.Text("location"),
		field.String("email"),
		field.String("password").Sensitive(),
		field.String("phone").Optional(),
		field.Bool("is_archived").Default(false),
		field.Bool("is_verified").Default(false),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
		field.Time("archived_at").Nillable(),
		field.Time("verified_at").Nillable(),
		field.Int64("otp").Nillable(),
	}
}

// Edges of the Institution.
func (Institution) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("recordaccess", RecordAccess.Type),
		edge.To("doctor", User.Type),
		edge.To("medicalrecord", MedicalRecord.Type),
	}
}
