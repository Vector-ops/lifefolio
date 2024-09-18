package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		field.UUID("record_id", uuid.UUID{}),
		field.UUID("institution_id", uuid.UUID{}),
		field.Bool("approved").Default(false),
		field.Time("approved_at").Nillable(),
	}
}

// Indexes of the RecordAccess
func (RecordAccess) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("record_id"),
		index.Fields("institution_id"),
	}
}

// Edges of the RecordAccess.
func (RecordAccess) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("medicalrecord", MedicalRecord.Type).Ref("recordaccess").Field("record_id").Unique().Required(),
		edge.From("institution", Institution.Type).Ref("recordaccess").Field("institution_id").Unique().Required(),
	}
}
