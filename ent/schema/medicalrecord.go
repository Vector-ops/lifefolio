package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("institution_id", uuid.UUID{}).Optional(),
		field.Text("file"),
		field.Bool("is_archived").Default(false),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
		field.Time("archived_at").Nillable(),
	}
}

// Indexes of the MedicalRecord
func (MedicalRecord) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("institution_id"),
	}
}

// Edges of the MedicalRecord.
func (MedicalRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("medicalrecord").Field("user_id").Unique().Required(),
		edge.From("institution", Institution.Type).Ref("medicalrecord").Field("institution_id").Unique(),
		edge.To("recordaccess", RecordAccess.Type),
	}
}
