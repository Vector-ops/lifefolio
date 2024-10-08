// Code generated by ent, DO NOT EDIT.

package medicalrecord

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/vector-ops/lifefolio/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldEQ(FieldUserID, v))
}

// InstitutionID applies equality check predicate on the "institution_id" field. It's identical to InstitutionIDEQ.
func InstitutionID(v uuid.UUID) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldEQ(FieldInstitutionID, v))
}

// File applies equality check predicate on the "file" field. It's identical to FileEQ.
func File(v string) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldEQ(FieldFile, v))
}

// IsArchived applies equality check predicate on the "is_archived" field. It's identical to IsArchivedEQ.
func IsArchived(v bool) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldEQ(FieldIsArchived, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldEQ(FieldUpdatedAt, v))
}

// ArchivedAt applies equality check predicate on the "archived_at" field. It's identical to ArchivedAtEQ.
func ArchivedAt(v time.Time) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldEQ(FieldArchivedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldNotIn(FieldUserID, vs...))
}

// InstitutionIDEQ applies the EQ predicate on the "institution_id" field.
func InstitutionIDEQ(v uuid.UUID) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldEQ(FieldInstitutionID, v))
}

// InstitutionIDNEQ applies the NEQ predicate on the "institution_id" field.
func InstitutionIDNEQ(v uuid.UUID) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldNEQ(FieldInstitutionID, v))
}

// InstitutionIDIn applies the In predicate on the "institution_id" field.
func InstitutionIDIn(vs ...uuid.UUID) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldIn(FieldInstitutionID, vs...))
}

// InstitutionIDNotIn applies the NotIn predicate on the "institution_id" field.
func InstitutionIDNotIn(vs ...uuid.UUID) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldNotIn(FieldInstitutionID, vs...))
}

// InstitutionIDIsNil applies the IsNil predicate on the "institution_id" field.
func InstitutionIDIsNil() predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldIsNull(FieldInstitutionID))
}

// InstitutionIDNotNil applies the NotNil predicate on the "institution_id" field.
func InstitutionIDNotNil() predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldNotNull(FieldInstitutionID))
}

// FileEQ applies the EQ predicate on the "file" field.
func FileEQ(v string) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldEQ(FieldFile, v))
}

// FileNEQ applies the NEQ predicate on the "file" field.
func FileNEQ(v string) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldNEQ(FieldFile, v))
}

// FileIn applies the In predicate on the "file" field.
func FileIn(vs ...string) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldIn(FieldFile, vs...))
}

// FileNotIn applies the NotIn predicate on the "file" field.
func FileNotIn(vs ...string) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldNotIn(FieldFile, vs...))
}

// FileGT applies the GT predicate on the "file" field.
func FileGT(v string) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldGT(FieldFile, v))
}

// FileGTE applies the GTE predicate on the "file" field.
func FileGTE(v string) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldGTE(FieldFile, v))
}

// FileLT applies the LT predicate on the "file" field.
func FileLT(v string) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldLT(FieldFile, v))
}

// FileLTE applies the LTE predicate on the "file" field.
func FileLTE(v string) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldLTE(FieldFile, v))
}

// FileContains applies the Contains predicate on the "file" field.
func FileContains(v string) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldContains(FieldFile, v))
}

// FileHasPrefix applies the HasPrefix predicate on the "file" field.
func FileHasPrefix(v string) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldHasPrefix(FieldFile, v))
}

// FileHasSuffix applies the HasSuffix predicate on the "file" field.
func FileHasSuffix(v string) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldHasSuffix(FieldFile, v))
}

// FileEqualFold applies the EqualFold predicate on the "file" field.
func FileEqualFold(v string) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldEqualFold(FieldFile, v))
}

// FileContainsFold applies the ContainsFold predicate on the "file" field.
func FileContainsFold(v string) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldContainsFold(FieldFile, v))
}

// IsArchivedEQ applies the EQ predicate on the "is_archived" field.
func IsArchivedEQ(v bool) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldEQ(FieldIsArchived, v))
}

// IsArchivedNEQ applies the NEQ predicate on the "is_archived" field.
func IsArchivedNEQ(v bool) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldNEQ(FieldIsArchived, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldLTE(FieldUpdatedAt, v))
}

// ArchivedAtEQ applies the EQ predicate on the "archived_at" field.
func ArchivedAtEQ(v time.Time) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldEQ(FieldArchivedAt, v))
}

// ArchivedAtNEQ applies the NEQ predicate on the "archived_at" field.
func ArchivedAtNEQ(v time.Time) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldNEQ(FieldArchivedAt, v))
}

// ArchivedAtIn applies the In predicate on the "archived_at" field.
func ArchivedAtIn(vs ...time.Time) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldIn(FieldArchivedAt, vs...))
}

// ArchivedAtNotIn applies the NotIn predicate on the "archived_at" field.
func ArchivedAtNotIn(vs ...time.Time) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldNotIn(FieldArchivedAt, vs...))
}

// ArchivedAtGT applies the GT predicate on the "archived_at" field.
func ArchivedAtGT(v time.Time) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldGT(FieldArchivedAt, v))
}

// ArchivedAtGTE applies the GTE predicate on the "archived_at" field.
func ArchivedAtGTE(v time.Time) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldGTE(FieldArchivedAt, v))
}

// ArchivedAtLT applies the LT predicate on the "archived_at" field.
func ArchivedAtLT(v time.Time) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldLT(FieldArchivedAt, v))
}

// ArchivedAtLTE applies the LTE predicate on the "archived_at" field.
func ArchivedAtLTE(v time.Time) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.FieldLTE(FieldArchivedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.MedicalRecord {
	return predicate.MedicalRecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.MedicalRecord {
	return predicate.MedicalRecord(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInstitution applies the HasEdge predicate on the "institution" edge.
func HasInstitution() predicate.MedicalRecord {
	return predicate.MedicalRecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InstitutionTable, InstitutionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInstitutionWith applies the HasEdge predicate on the "institution" edge with a given conditions (other predicates).
func HasInstitutionWith(preds ...predicate.Institution) predicate.MedicalRecord {
	return predicate.MedicalRecord(func(s *sql.Selector) {
		step := newInstitutionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRecordaccess applies the HasEdge predicate on the "recordaccess" edge.
func HasRecordaccess() predicate.MedicalRecord {
	return predicate.MedicalRecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RecordaccessTable, RecordaccessColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRecordaccessWith applies the HasEdge predicate on the "recordaccess" edge with a given conditions (other predicates).
func HasRecordaccessWith(preds ...predicate.RecordAccess) predicate.MedicalRecord {
	return predicate.MedicalRecord(func(s *sql.Selector) {
		step := newRecordaccessStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MedicalRecord) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MedicalRecord) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MedicalRecord) predicate.MedicalRecord {
	return predicate.MedicalRecord(sql.NotPredicates(p))
}
