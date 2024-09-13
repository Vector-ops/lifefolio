package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// RecordAccess holds the schema definition for the RecordAccess entity.
type RecordAccess struct {
	ent.Schema
}

// Fields of the RecordAccess.
func (RecordAccess) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.Bool("approved").Default(false),
		field.Time("approved_at").Nillable(),
	}
}

// Edges of the RecordAccess.
func (RecordAccess) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("medicalrecord", MedicalRecord.Type).Ref("recordaccess").Unique(),
		edge.From("institution", Institution.Type).Ref("recordaccess").Unique(),
	}
}
