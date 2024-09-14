// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccessRequestsColumns holds the columns for the "access_requests" table.
	AccessRequestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// AccessRequestsTable holds the schema information for the "access_requests" table.
	AccessRequestsTable = &schema.Table{
		Name:       "access_requests",
		Columns:    AccessRequestsColumns,
		PrimaryKey: []*schema.Column{AccessRequestsColumns[0]},
	}
	// InstitutionsColumns holds the columns for the "institutions" table.
	InstitutionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Size: 2147483647},
		{Name: "location", Type: field.TypeString, Size: 2147483647},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "is_archived", Type: field.TypeBool, Default: false},
		{Name: "is_verified", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "archived_at", Type: field.TypeTime},
		{Name: "verified_at", Type: field.TypeTime},
		{Name: "otp", Type: field.TypeInt64},
	}
	// InstitutionsTable holds the schema information for the "institutions" table.
	InstitutionsTable = &schema.Table{
		Name:       "institutions",
		Columns:    InstitutionsColumns,
		PrimaryKey: []*schema.Column{InstitutionsColumns[0]},
	}
	// MedicalRecordsColumns holds the columns for the "medical_records" table.
	MedicalRecordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "file", Type: field.TypeString, Size: 2147483647},
		{Name: "is_archived", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "archived_at", Type: field.TypeTime},
		{Name: "institution_medicalrecord", Type: field.TypeUUID, Nullable: true},
		{Name: "user_medicalrecord", Type: field.TypeUUID, Nullable: true},
	}
	// MedicalRecordsTable holds the schema information for the "medical_records" table.
	MedicalRecordsTable = &schema.Table{
		Name:       "medical_records",
		Columns:    MedicalRecordsColumns,
		PrimaryKey: []*schema.Column{MedicalRecordsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "medical_records_institutions_medicalrecord",
				Columns:    []*schema.Column{MedicalRecordsColumns[6]},
				RefColumns: []*schema.Column{InstitutionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "medical_records_users_medicalrecord",
				Columns:    []*schema.Column{MedicalRecordsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RecordAccessesColumns holds the columns for the "record_accesses" table.
	RecordAccessesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "approved", Type: field.TypeBool, Default: false},
		{Name: "approved_at", Type: field.TypeTime},
		{Name: "institution_recordaccess", Type: field.TypeUUID, Nullable: true},
		{Name: "medical_record_recordaccess", Type: field.TypeUUID, Nullable: true},
	}
	// RecordAccessesTable holds the schema information for the "record_accesses" table.
	RecordAccessesTable = &schema.Table{
		Name:       "record_accesses",
		Columns:    RecordAccessesColumns,
		PrimaryKey: []*schema.Column{RecordAccessesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "record_accesses_institutions_recordaccess",
				Columns:    []*schema.Column{RecordAccessesColumns[3]},
				RefColumns: []*schema.Column{InstitutionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "record_accesses_medical_records_recordaccess",
				Columns:    []*schema.Column{RecordAccessesColumns[4]},
				RefColumns: []*schema.Column{MedicalRecordsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "patient_id", Type: field.TypeString},
		{Name: "first_name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "dob", Type: field.TypeTime, Nullable: true},
		{Name: "user_type", Type: field.TypeEnum, Enums: []string{"PATIENT", "DOCTOR"}, Default: "PATIENT"},
		{Name: "blood_group", Type: field.TypeEnum, Nullable: true, Enums: []string{"APOS", "ANEG", "BPOS", "BNEG", "ABPOS", "ABNEG", "OPOS", "ONEG"}},
		{Name: "weight", Type: field.TypeFloat32, Nullable: true},
		{Name: "height", Type: field.TypeFloat32, Nullable: true},
		{Name: "is_archived", Type: field.TypeBool, Default: false},
		{Name: "is_verified", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "archived_at", Type: field.TypeTime},
		{Name: "verified_at", Type: field.TypeTime},
		{Name: "otp", Type: field.TypeUint64},
		{Name: "institution_doctor", Type: field.TypeUUID, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_institutions_doctor",
				Columns:    []*schema.Column{UsersColumns[18]},
				RefColumns: []*schema.Column{InstitutionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccessRequestsTable,
		InstitutionsTable,
		MedicalRecordsTable,
		RecordAccessesTable,
		UsersTable,
	}
)

func init() {
	MedicalRecordsTable.ForeignKeys[0].RefTable = InstitutionsTable
	MedicalRecordsTable.ForeignKeys[1].RefTable = UsersTable
	RecordAccessesTable.ForeignKeys[0].RefTable = InstitutionsTable
	RecordAccessesTable.ForeignKeys[1].RefTable = MedicalRecordsTable
	UsersTable.ForeignKeys[0].RefTable = InstitutionsTable
}
