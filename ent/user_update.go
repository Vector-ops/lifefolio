// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vector-ops/lifefolio/ent/predicate"
	"github.com/vector-ops/lifefolio/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetFirstName sets the "firstName" field.
func (uu *UserUpdate) SetFirstName(s string) *UserUpdate {
	uu.mutation.SetFirstName(s)
	return uu
}

// SetNillableFirstName sets the "firstName" field if the given value is not nil.
func (uu *UserUpdate) SetNillableFirstName(s *string) *UserUpdate {
	if s != nil {
		uu.SetFirstName(*s)
	}
	return uu
}

// SetLastName sets the "lastName" field.
func (uu *UserUpdate) SetLastName(s string) *UserUpdate {
	uu.mutation.SetLastName(s)
	return uu
}

// SetNillableLastName sets the "lastName" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLastName(s *string) *UserUpdate {
	if s != nil {
		uu.SetLastName(*s)
	}
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePassword(s *string) *UserUpdate {
	if s != nil {
		uu.SetPassword(*s)
	}
	return uu
}

// SetDOB sets the "DOB" field.
func (uu *UserUpdate) SetDOB(t time.Time) *UserUpdate {
	uu.mutation.SetDOB(t)
	return uu
}

// SetNillableDOB sets the "DOB" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDOB(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetDOB(*t)
	}
	return uu
}

// SetBloodGroup sets the "bloodGroup" field.
func (uu *UserUpdate) SetBloodGroup(ug user.BloodGroup) *UserUpdate {
	uu.mutation.SetBloodGroup(ug)
	return uu
}

// SetNillableBloodGroup sets the "bloodGroup" field if the given value is not nil.
func (uu *UserUpdate) SetNillableBloodGroup(ug *user.BloodGroup) *UserUpdate {
	if ug != nil {
		uu.SetBloodGroup(*ug)
	}
	return uu
}

// ClearBloodGroup clears the value of the "bloodGroup" field.
func (uu *UserUpdate) ClearBloodGroup() *UserUpdate {
	uu.mutation.ClearBloodGroup()
	return uu
}

// SetWeight sets the "weight" field.
func (uu *UserUpdate) SetWeight(f float32) *UserUpdate {
	uu.mutation.ResetWeight()
	uu.mutation.SetWeight(f)
	return uu
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (uu *UserUpdate) SetNillableWeight(f *float32) *UserUpdate {
	if f != nil {
		uu.SetWeight(*f)
	}
	return uu
}

// AddWeight adds f to the "weight" field.
func (uu *UserUpdate) AddWeight(f float32) *UserUpdate {
	uu.mutation.AddWeight(f)
	return uu
}

// ClearWeight clears the value of the "weight" field.
func (uu *UserUpdate) ClearWeight() *UserUpdate {
	uu.mutation.ClearWeight()
	return uu
}

// SetHeight sets the "height" field.
func (uu *UserUpdate) SetHeight(f float32) *UserUpdate {
	uu.mutation.ResetHeight()
	uu.mutation.SetHeight(f)
	return uu
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (uu *UserUpdate) SetNillableHeight(f *float32) *UserUpdate {
	if f != nil {
		uu.SetHeight(*f)
	}
	return uu
}

// AddHeight adds f to the "height" field.
func (uu *UserUpdate) AddHeight(f float32) *UserUpdate {
	uu.mutation.AddHeight(f)
	return uu
}

// ClearHeight clears the value of the "height" field.
func (uu *UserUpdate) ClearHeight() *UserUpdate {
	uu.mutation.ClearHeight()
	return uu
}

// SetIsArchived sets the "isArchived" field.
func (uu *UserUpdate) SetIsArchived(b bool) *UserUpdate {
	uu.mutation.SetIsArchived(b)
	return uu
}

// SetNillableIsArchived sets the "isArchived" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsArchived(b *bool) *UserUpdate {
	if b != nil {
		uu.SetIsArchived(*b)
	}
	return uu
}

// SetIsVerified sets the "isVerified" field.
func (uu *UserUpdate) SetIsVerified(b bool) *UserUpdate {
	uu.mutation.SetIsVerified(b)
	return uu
}

// SetNillableIsVerified sets the "isVerified" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsVerified(b *bool) *UserUpdate {
	if b != nil {
		uu.SetIsVerified(*b)
	}
	return uu
}

// SetCreatedAt sets the "createdAt" field.
func (uu *UserUpdate) SetCreatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetCreatedAt(t)
	return uu
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetCreatedAt(*t)
	}
	return uu
}

// SetUpdatedAt sets the "updatedAt" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUpdatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetUpdatedAt(*t)
	}
	return uu
}

// SetArchivedAt sets the "archivedAt" field.
func (uu *UserUpdate) SetArchivedAt(t time.Time) *UserUpdate {
	uu.mutation.SetArchivedAt(t)
	return uu
}

// SetNillableArchivedAt sets the "archivedAt" field if the given value is not nil.
func (uu *UserUpdate) SetNillableArchivedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetArchivedAt(*t)
	}
	return uu
}

// SetVerifiedAt sets the "verifiedAt" field.
func (uu *UserUpdate) SetVerifiedAt(t time.Time) *UserUpdate {
	uu.mutation.SetVerifiedAt(t)
	return uu
}

// SetNillableVerifiedAt sets the "verifiedAt" field if the given value is not nil.
func (uu *UserUpdate) SetNillableVerifiedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetVerifiedAt(*t)
	}
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.FirstName(); ok {
		if err := user.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "firstName", err: fmt.Errorf(`ent: validator failed for field "User.firstName": %w`, err)}
		}
	}
	if v, ok := uu.mutation.LastName(); ok {
		if err := user.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "lastName", err: fmt.Errorf(`ent: validator failed for field "User.lastName": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.BloodGroup(); ok {
		if err := user.BloodGroupValidator(v); err != nil {
			return &ValidationError{Name: "bloodGroup", err: fmt.Errorf(`ent: validator failed for field "User.bloodGroup": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if value, ok := uu.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.DOB(); ok {
		_spec.SetField(user.FieldDOB, field.TypeTime, value)
	}
	if value, ok := uu.mutation.BloodGroup(); ok {
		_spec.SetField(user.FieldBloodGroup, field.TypeEnum, value)
	}
	if uu.mutation.BloodGroupCleared() {
		_spec.ClearField(user.FieldBloodGroup, field.TypeEnum)
	}
	if value, ok := uu.mutation.Weight(); ok {
		_spec.SetField(user.FieldWeight, field.TypeFloat32, value)
	}
	if value, ok := uu.mutation.AddedWeight(); ok {
		_spec.AddField(user.FieldWeight, field.TypeFloat32, value)
	}
	if uu.mutation.WeightCleared() {
		_spec.ClearField(user.FieldWeight, field.TypeFloat32)
	}
	if value, ok := uu.mutation.Height(); ok {
		_spec.SetField(user.FieldHeight, field.TypeFloat32, value)
	}
	if value, ok := uu.mutation.AddedHeight(); ok {
		_spec.AddField(user.FieldHeight, field.TypeFloat32, value)
	}
	if uu.mutation.HeightCleared() {
		_spec.ClearField(user.FieldHeight, field.TypeFloat32)
	}
	if value, ok := uu.mutation.IsArchived(); ok {
		_spec.SetField(user.FieldIsArchived, field.TypeBool, value)
	}
	if value, ok := uu.mutation.IsVerified(); ok {
		_spec.SetField(user.FieldIsVerified, field.TypeBool, value)
	}
	if value, ok := uu.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.ArchivedAt(); ok {
		_spec.SetField(user.FieldArchivedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.VerifiedAt(); ok {
		_spec.SetField(user.FieldVerifiedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetFirstName sets the "firstName" field.
func (uuo *UserUpdateOne) SetFirstName(s string) *UserUpdateOne {
	uuo.mutation.SetFirstName(s)
	return uuo
}

// SetNillableFirstName sets the "firstName" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableFirstName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetFirstName(*s)
	}
	return uuo
}

// SetLastName sets the "lastName" field.
func (uuo *UserUpdateOne) SetLastName(s string) *UserUpdateOne {
	uuo.mutation.SetLastName(s)
	return uuo
}

// SetNillableLastName sets the "lastName" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLastName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetLastName(*s)
	}
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePassword(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPassword(*s)
	}
	return uuo
}

// SetDOB sets the "DOB" field.
func (uuo *UserUpdateOne) SetDOB(t time.Time) *UserUpdateOne {
	uuo.mutation.SetDOB(t)
	return uuo
}

// SetNillableDOB sets the "DOB" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDOB(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetDOB(*t)
	}
	return uuo
}

// SetBloodGroup sets the "bloodGroup" field.
func (uuo *UserUpdateOne) SetBloodGroup(ug user.BloodGroup) *UserUpdateOne {
	uuo.mutation.SetBloodGroup(ug)
	return uuo
}

// SetNillableBloodGroup sets the "bloodGroup" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableBloodGroup(ug *user.BloodGroup) *UserUpdateOne {
	if ug != nil {
		uuo.SetBloodGroup(*ug)
	}
	return uuo
}

// ClearBloodGroup clears the value of the "bloodGroup" field.
func (uuo *UserUpdateOne) ClearBloodGroup() *UserUpdateOne {
	uuo.mutation.ClearBloodGroup()
	return uuo
}

// SetWeight sets the "weight" field.
func (uuo *UserUpdateOne) SetWeight(f float32) *UserUpdateOne {
	uuo.mutation.ResetWeight()
	uuo.mutation.SetWeight(f)
	return uuo
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableWeight(f *float32) *UserUpdateOne {
	if f != nil {
		uuo.SetWeight(*f)
	}
	return uuo
}

// AddWeight adds f to the "weight" field.
func (uuo *UserUpdateOne) AddWeight(f float32) *UserUpdateOne {
	uuo.mutation.AddWeight(f)
	return uuo
}

// ClearWeight clears the value of the "weight" field.
func (uuo *UserUpdateOne) ClearWeight() *UserUpdateOne {
	uuo.mutation.ClearWeight()
	return uuo
}

// SetHeight sets the "height" field.
func (uuo *UserUpdateOne) SetHeight(f float32) *UserUpdateOne {
	uuo.mutation.ResetHeight()
	uuo.mutation.SetHeight(f)
	return uuo
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableHeight(f *float32) *UserUpdateOne {
	if f != nil {
		uuo.SetHeight(*f)
	}
	return uuo
}

// AddHeight adds f to the "height" field.
func (uuo *UserUpdateOne) AddHeight(f float32) *UserUpdateOne {
	uuo.mutation.AddHeight(f)
	return uuo
}

// ClearHeight clears the value of the "height" field.
func (uuo *UserUpdateOne) ClearHeight() *UserUpdateOne {
	uuo.mutation.ClearHeight()
	return uuo
}

// SetIsArchived sets the "isArchived" field.
func (uuo *UserUpdateOne) SetIsArchived(b bool) *UserUpdateOne {
	uuo.mutation.SetIsArchived(b)
	return uuo
}

// SetNillableIsArchived sets the "isArchived" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsArchived(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetIsArchived(*b)
	}
	return uuo
}

// SetIsVerified sets the "isVerified" field.
func (uuo *UserUpdateOne) SetIsVerified(b bool) *UserUpdateOne {
	uuo.mutation.SetIsVerified(b)
	return uuo
}

// SetNillableIsVerified sets the "isVerified" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsVerified(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetIsVerified(*b)
	}
	return uuo
}

// SetCreatedAt sets the "createdAt" field.
func (uuo *UserUpdateOne) SetCreatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreatedAt(t)
	return uuo
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetCreatedAt(*t)
	}
	return uuo
}

// SetUpdatedAt sets the "updatedAt" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUpdatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetUpdatedAt(*t)
	}
	return uuo
}

// SetArchivedAt sets the "archivedAt" field.
func (uuo *UserUpdateOne) SetArchivedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetArchivedAt(t)
	return uuo
}

// SetNillableArchivedAt sets the "archivedAt" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableArchivedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetArchivedAt(*t)
	}
	return uuo
}

// SetVerifiedAt sets the "verifiedAt" field.
func (uuo *UserUpdateOne) SetVerifiedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetVerifiedAt(t)
	return uuo
}

// SetNillableVerifiedAt sets the "verifiedAt" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableVerifiedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetVerifiedAt(*t)
	}
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.FirstName(); ok {
		if err := user.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "firstName", err: fmt.Errorf(`ent: validator failed for field "User.firstName": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.LastName(); ok {
		if err := user.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "lastName", err: fmt.Errorf(`ent: validator failed for field "User.lastName": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.BloodGroup(); ok {
		if err := user.BloodGroupValidator(v); err != nil {
			return &ValidationError{Name: "bloodGroup", err: fmt.Errorf(`ent: validator failed for field "User.bloodGroup": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.DOB(); ok {
		_spec.SetField(user.FieldDOB, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.BloodGroup(); ok {
		_spec.SetField(user.FieldBloodGroup, field.TypeEnum, value)
	}
	if uuo.mutation.BloodGroupCleared() {
		_spec.ClearField(user.FieldBloodGroup, field.TypeEnum)
	}
	if value, ok := uuo.mutation.Weight(); ok {
		_spec.SetField(user.FieldWeight, field.TypeFloat32, value)
	}
	if value, ok := uuo.mutation.AddedWeight(); ok {
		_spec.AddField(user.FieldWeight, field.TypeFloat32, value)
	}
	if uuo.mutation.WeightCleared() {
		_spec.ClearField(user.FieldWeight, field.TypeFloat32)
	}
	if value, ok := uuo.mutation.Height(); ok {
		_spec.SetField(user.FieldHeight, field.TypeFloat32, value)
	}
	if value, ok := uuo.mutation.AddedHeight(); ok {
		_spec.AddField(user.FieldHeight, field.TypeFloat32, value)
	}
	if uuo.mutation.HeightCleared() {
		_spec.ClearField(user.FieldHeight, field.TypeFloat32)
	}
	if value, ok := uuo.mutation.IsArchived(); ok {
		_spec.SetField(user.FieldIsArchived, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.IsVerified(); ok {
		_spec.SetField(user.FieldIsVerified, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.ArchivedAt(); ok {
		_spec.SetField(user.FieldArchivedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.VerifiedAt(); ok {
		_spec.SetField(user.FieldVerifiedAt, field.TypeTime, value)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
