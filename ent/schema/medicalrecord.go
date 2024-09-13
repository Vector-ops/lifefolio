package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// MedicalRecord holds the schema definition for the MedicalRecord entity.
type MedicalRecord struct {
	ent.Schema
}

// Fields of the MedicalRecord.
func (MedicalRecord) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.Text("file"),
		field.Bool("is_archived").Default(false),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
		field.Time("archived_at").Nillable(),
	}
}

// Edges of the MedicalRecord.
func (MedicalRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("medicalrecord").Unique(),
		edge.From("institution", Institution.Type).Ref("medicalrecord").Unique(),
		edge.To("recordaccess", RecordAccess.Type),
	}
}
