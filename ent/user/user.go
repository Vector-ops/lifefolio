// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPatientID holds the string denoting the patient_id field in the database.
	FieldPatientID = "patient_id"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldDOB holds the string denoting the dob field in the database.
	FieldDOB = "dob"
	// FieldUserType holds the string denoting the user_type field in the database.
	FieldUserType = "user_type"
	// FieldBloodGroup holds the string denoting the blood_group field in the database.
	FieldBloodGroup = "blood_group"
	// FieldWeight holds the string denoting the weight field in the database.
	FieldWeight = "weight"
	// FieldHeight holds the string denoting the height field in the database.
	FieldHeight = "height"
	// FieldIsArchived holds the string denoting the is_archived field in the database.
	FieldIsArchived = "is_archived"
	// FieldIsVerified holds the string denoting the is_verified field in the database.
	FieldIsVerified = "is_verified"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldArchivedAt holds the string denoting the archived_at field in the database.
	FieldArchivedAt = "archived_at"
	// FieldVerifiedAt holds the string denoting the verified_at field in the database.
	FieldVerifiedAt = "verified_at"
	// FieldOtp holds the string denoting the otp field in the database.
	FieldOtp = "otp"
	// EdgeMedicalrecord holds the string denoting the medicalrecord edge name in mutations.
	EdgeMedicalrecord = "medicalrecord"
	// EdgeInstitution holds the string denoting the institution edge name in mutations.
	EdgeInstitution = "institution"
	// Table holds the table name of the user in the database.
	Table = "users"
	// MedicalrecordTable is the table that holds the medicalrecord relation/edge.
	MedicalrecordTable = "medical_records"
	// MedicalrecordInverseTable is the table name for the MedicalRecord entity.
	// It exists in this package in order to avoid circular dependency with the "medicalrecord" package.
	MedicalrecordInverseTable = "medical_records"
	// MedicalrecordColumn is the table column denoting the medicalrecord relation/edge.
	MedicalrecordColumn = "user_id"
	// InstitutionTable is the table that holds the institution relation/edge.
	InstitutionTable = "users"
	// InstitutionInverseTable is the table name for the Institution entity.
	// It exists in this package in order to avoid circular dependency with the "institution" package.
	InstitutionInverseTable = "institutions"
	// InstitutionColumn is the table column denoting the institution relation/edge.
	InstitutionColumn = "institution_doctor"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldPatientID,
	FieldFirstName,
	FieldLastName,
	FieldEmail,
	FieldPassword,
	FieldDOB,
	FieldUserType,
	FieldBloodGroup,
	FieldWeight,
	FieldHeight,
	FieldIsArchived,
	FieldIsVerified,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldArchivedAt,
	FieldVerifiedAt,
	FieldOtp,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"institution_doctor",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// PatientIDValidator is a validator for the "patient_id" field. It is called by the builders before save.
	PatientIDValidator func(string) error
	// FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	FirstNameValidator func(string) error
	// LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	LastNameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// DefaultIsArchived holds the default value on creation for the "is_archived" field.
	DefaultIsArchived bool
	// DefaultIsVerified holds the default value on creation for the "is_verified" field.
	DefaultIsVerified bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)

// UserType defines the type for the "user_type" enum field.
type UserType string

// UserTypePATIENT is the default value of the UserType enum.
const DefaultUserType = UserTypePATIENT

// UserType values.
const (
	UserTypePATIENT UserType = "PATIENT"
	UserTypeDOCTOR  UserType = "DOCTOR"
)

func (ut UserType) String() string {
	return string(ut)
}

// UserTypeValidator is a validator for the "user_type" field enum values. It is called by the builders before save.
func UserTypeValidator(ut UserType) error {
	switch ut {
	case UserTypePATIENT, UserTypeDOCTOR:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for user_type field: %q", ut)
	}
}

// BloodGroup defines the type for the "blood_group" enum field.
type BloodGroup string

// BloodGroup values.
const (
	BloodGroupAPOS  BloodGroup = "APOS"
	BloodGroupANEG  BloodGroup = "ANEG"
	BloodGroupBPOS  BloodGroup = "BPOS"
	BloodGroupBNEG  BloodGroup = "BNEG"
	BloodGroupABPOS BloodGroup = "ABPOS"
	BloodGroupABNEG BloodGroup = "ABNEG"
	BloodGroupOPOS  BloodGroup = "OPOS"
	BloodGroupONEG  BloodGroup = "ONEG"
)

func (bg BloodGroup) String() string {
	return string(bg)
}

// BloodGroupValidator is a validator for the "blood_group" field enum values. It is called by the builders before save.
func BloodGroupValidator(bg BloodGroup) error {
	switch bg {
	case BloodGroupAPOS, BloodGroupANEG, BloodGroupBPOS, BloodGroupBNEG, BloodGroupABPOS, BloodGroupABNEG, BloodGroupOPOS, BloodGroupONEG:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for blood_group field: %q", bg)
	}
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPatientID orders the results by the patient_id field.
func ByPatientID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPatientID, opts...).ToFunc()
}

// ByFirstName orders the results by the first_name field.
func ByFirstName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstName, opts...).ToFunc()
}

// ByLastName orders the results by the last_name field.
func ByLastName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastName, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByDOB orders the results by the DOB field.
func ByDOB(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDOB, opts...).ToFunc()
}

// ByUserType orders the results by the user_type field.
func ByUserType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserType, opts...).ToFunc()
}

// ByBloodGroup orders the results by the blood_group field.
func ByBloodGroup(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBloodGroup, opts...).ToFunc()
}

// ByWeight orders the results by the weight field.
func ByWeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWeight, opts...).ToFunc()
}

// ByHeight orders the results by the height field.
func ByHeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHeight, opts...).ToFunc()
}

// ByIsArchived orders the results by the is_archived field.
func ByIsArchived(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsArchived, opts...).ToFunc()
}

// ByIsVerified orders the results by the is_verified field.
func ByIsVerified(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsVerified, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByArchivedAt orders the results by the archived_at field.
func ByArchivedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldArchivedAt, opts...).ToFunc()
}

// ByVerifiedAt orders the results by the verified_at field.
func ByVerifiedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVerifiedAt, opts...).ToFunc()
}

// ByOtp orders the results by the otp field.
func ByOtp(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOtp, opts...).ToFunc()
}

// ByMedicalrecordCount orders the results by medicalrecord count.
func ByMedicalrecordCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMedicalrecordStep(), opts...)
	}
}

// ByMedicalrecord orders the results by medicalrecord terms.
func ByMedicalrecord(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMedicalrecordStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByInstitutionField orders the results by institution field.
func ByInstitutionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInstitutionStep(), sql.OrderByField(field, opts...))
	}
}
func newMedicalrecordStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MedicalrecordInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MedicalrecordTable, MedicalrecordColumn),
	)
}
func newInstitutionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InstitutionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, InstitutionTable, InstitutionColumn),
	)
}
