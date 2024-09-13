// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/vector-ops/lifefolio/ent/institution"
	"github.com/vector-ops/lifefolio/ent/medicalrecord"
	"github.com/vector-ops/lifefolio/ent/user"
)

// MedicalRecord is the model entity for the MedicalRecord schema.
type MedicalRecord struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// File holds the value of the "file" field.
	File string `json:"file,omitempty"`
	// IsArchived holds the value of the "is_archived" field.
	IsArchived bool `json:"is_archived,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// ArchivedAt holds the value of the "archived_at" field.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MedicalRecordQuery when eager-loading is set.
	Edges                     MedicalRecordEdges `json:"edges"`
	institution_medicalrecord *uuid.UUID
	user_medicalrecord        *uuid.UUID
	selectValues              sql.SelectValues
}

// MedicalRecordEdges holds the relations/edges for other nodes in the graph.
type MedicalRecordEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Institution holds the value of the institution edge.
	Institution *Institution `json:"institution,omitempty"`
	// Recordaccess holds the value of the recordaccess edge.
	Recordaccess []*RecordAccess `json:"recordaccess,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MedicalRecordEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// InstitutionOrErr returns the Institution value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MedicalRecordEdges) InstitutionOrErr() (*Institution, error) {
	if e.Institution != nil {
		return e.Institution, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: institution.Label}
	}
	return nil, &NotLoadedError{edge: "institution"}
}

// RecordaccessOrErr returns the Recordaccess value or an error if the edge
// was not loaded in eager-loading.
func (e MedicalRecordEdges) RecordaccessOrErr() ([]*RecordAccess, error) {
	if e.loadedTypes[2] {
		return e.Recordaccess, nil
	}
	return nil, &NotLoadedError{edge: "recordaccess"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MedicalRecord) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case medicalrecord.FieldIsArchived:
			values[i] = new(sql.NullBool)
		case medicalrecord.FieldFile:
			values[i] = new(sql.NullString)
		case medicalrecord.FieldCreatedAt, medicalrecord.FieldUpdatedAt, medicalrecord.FieldArchivedAt:
			values[i] = new(sql.NullTime)
		case medicalrecord.FieldID:
			values[i] = new(uuid.UUID)
		case medicalrecord.ForeignKeys[0]: // institution_medicalrecord
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case medicalrecord.ForeignKeys[1]: // user_medicalrecord
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MedicalRecord fields.
func (mr *MedicalRecord) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case medicalrecord.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				mr.ID = *value
			}
		case medicalrecord.FieldFile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field file", values[i])
			} else if value.Valid {
				mr.File = value.String
			}
		case medicalrecord.FieldIsArchived:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_archived", values[i])
			} else if value.Valid {
				mr.IsArchived = value.Bool
			}
		case medicalrecord.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				mr.CreatedAt = value.Time
			}
		case medicalrecord.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				mr.UpdatedAt = value.Time
			}
		case medicalrecord.FieldArchivedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field archived_at", values[i])
			} else if value.Valid {
				mr.ArchivedAt = new(time.Time)
				*mr.ArchivedAt = value.Time
			}
		case medicalrecord.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field institution_medicalrecord", values[i])
			} else if value.Valid {
				mr.institution_medicalrecord = new(uuid.UUID)
				*mr.institution_medicalrecord = *value.S.(*uuid.UUID)
			}
		case medicalrecord.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_medicalrecord", values[i])
			} else if value.Valid {
				mr.user_medicalrecord = new(uuid.UUID)
				*mr.user_medicalrecord = *value.S.(*uuid.UUID)
			}
		default:
			mr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MedicalRecord.
// This includes values selected through modifiers, order, etc.
func (mr *MedicalRecord) Value(name string) (ent.Value, error) {
	return mr.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the MedicalRecord entity.
func (mr *MedicalRecord) QueryUser() *UserQuery {
	return NewMedicalRecordClient(mr.config).QueryUser(mr)
}

// QueryInstitution queries the "institution" edge of the MedicalRecord entity.
func (mr *MedicalRecord) QueryInstitution() *InstitutionQuery {
	return NewMedicalRecordClient(mr.config).QueryInstitution(mr)
}

// QueryRecordaccess queries the "recordaccess" edge of the MedicalRecord entity.
func (mr *MedicalRecord) QueryRecordaccess() *RecordAccessQuery {
	return NewMedicalRecordClient(mr.config).QueryRecordaccess(mr)
}

// Update returns a builder for updating this MedicalRecord.
// Note that you need to call MedicalRecord.Unwrap() before calling this method if this MedicalRecord
// was returned from a transaction, and the transaction was committed or rolled back.
func (mr *MedicalRecord) Update() *MedicalRecordUpdateOne {
	return NewMedicalRecordClient(mr.config).UpdateOne(mr)
}

// Unwrap unwraps the MedicalRecord entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mr *MedicalRecord) Unwrap() *MedicalRecord {
	_tx, ok := mr.config.driver.(*txDriver)
	if !ok {
		panic("ent: MedicalRecord is not a transactional entity")
	}
	mr.config.driver = _tx.drv
	return mr
}

// String implements the fmt.Stringer.
func (mr *MedicalRecord) String() string {
	var builder strings.Builder
	builder.WriteString("MedicalRecord(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mr.ID))
	builder.WriteString("file=")
	builder.WriteString(mr.File)
	builder.WriteString(", ")
	builder.WriteString("is_archived=")
	builder.WriteString(fmt.Sprintf("%v", mr.IsArchived))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(mr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(mr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := mr.ArchivedAt; v != nil {
		builder.WriteString("archived_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// MedicalRecords is a parsable slice of MedicalRecord.
type MedicalRecords []*MedicalRecord
