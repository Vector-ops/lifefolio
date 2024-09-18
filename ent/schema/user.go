package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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
		field.String("patient_id").NotEmpty().Unique(),
		field.String("first_name").NotEmpty(),
		field.String("last_name").NotEmpty(),
		field.String("email").NotEmpty().Unique(),
		field.String("password").Sensitive().NotEmpty(),
		field.Time("DOB").Optional(),
		field.Enum("user_type").Values("PATIENT", "DOCTOR").Default("PATIENT"),
		field.Enum("blood_group").Values("APOS", "ANEG", "BPOS", "BNEG", "ABPOS", "ABNEG", "OPOS", "ONEG").Optional(),
		field.Float32("weight").Optional(),
		field.Float32("height").Optional(),
		field.Bool("is_archived").Default(false),
		field.Bool("is_verified").Default(false),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
		field.Time("archived_at").Nillable(),
		field.Time("verified_at").Nillable(),
		field.Uint64("otp").Nillable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("medicalrecord", MedicalRecord.Type),
		edge.From("institution", Institution.Type).Ref("doctor").Unique(),
	}
}
