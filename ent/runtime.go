// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/vector-ops/lifefolio/ent/schema"
	"github.com/vector-ops/lifefolio/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescFirstName is the schema descriptor for firstName field.
	userDescFirstName := userFields[1].Descriptor()
	// user.FirstNameValidator is a validator for the "firstName" field. It is called by the builders before save.
	user.FirstNameValidator = userDescFirstName.Validators[0].(func(string) error)
	// userDescLastName is the schema descriptor for lastName field.
	userDescLastName := userFields[2].Descriptor()
	// user.LastNameValidator is a validator for the "lastName" field. It is called by the builders before save.
	user.LastNameValidator = userDescLastName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[3].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescDOB is the schema descriptor for DOB field.
	userDescDOB := userFields[5].Descriptor()
	// user.DefaultDOB holds the default value on creation for the DOB field.
	user.DefaultDOB = userDescDOB.Default.(time.Time)
	// userDescIsArchived is the schema descriptor for isArchived field.
	userDescIsArchived := userFields[9].Descriptor()
	// user.DefaultIsArchived holds the default value on creation for the isArchived field.
	user.DefaultIsArchived = userDescIsArchived.Default.(bool)
	// userDescIsVerified is the schema descriptor for isVerified field.
	userDescIsVerified := userFields[10].Descriptor()
	// user.DefaultIsVerified holds the default value on creation for the isVerified field.
	user.DefaultIsVerified = userDescIsVerified.Default.(bool)
	// userDescCreatedAt is the schema descriptor for createdAt field.
	userDescCreatedAt := userFields[11].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the createdAt field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(time.Time)
	// userDescUpdatedAt is the schema descriptor for updatedAt field.
	userDescUpdatedAt := userFields[12].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(time.Time)
}
