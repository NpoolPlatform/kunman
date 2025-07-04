// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/authhistory"
	"github.com/google/uuid"
)

// AuthHistoryCreate is the builder for creating a AuthHistory entity.
type AuthHistoryCreate struct {
	config
	mutation *AuthHistoryMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (ahc *AuthHistoryCreate) SetCreatedAt(u uint32) *AuthHistoryCreate {
	ahc.mutation.SetCreatedAt(u)
	return ahc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ahc *AuthHistoryCreate) SetNillableCreatedAt(u *uint32) *AuthHistoryCreate {
	if u != nil {
		ahc.SetCreatedAt(*u)
	}
	return ahc
}

// SetUpdatedAt sets the "updated_at" field.
func (ahc *AuthHistoryCreate) SetUpdatedAt(u uint32) *AuthHistoryCreate {
	ahc.mutation.SetUpdatedAt(u)
	return ahc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ahc *AuthHistoryCreate) SetNillableUpdatedAt(u *uint32) *AuthHistoryCreate {
	if u != nil {
		ahc.SetUpdatedAt(*u)
	}
	return ahc
}

// SetDeletedAt sets the "deleted_at" field.
func (ahc *AuthHistoryCreate) SetDeletedAt(u uint32) *AuthHistoryCreate {
	ahc.mutation.SetDeletedAt(u)
	return ahc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ahc *AuthHistoryCreate) SetNillableDeletedAt(u *uint32) *AuthHistoryCreate {
	if u != nil {
		ahc.SetDeletedAt(*u)
	}
	return ahc
}

// SetEntID sets the "ent_id" field.
func (ahc *AuthHistoryCreate) SetEntID(u uuid.UUID) *AuthHistoryCreate {
	ahc.mutation.SetEntID(u)
	return ahc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (ahc *AuthHistoryCreate) SetNillableEntID(u *uuid.UUID) *AuthHistoryCreate {
	if u != nil {
		ahc.SetEntID(*u)
	}
	return ahc
}

// SetAppID sets the "app_id" field.
func (ahc *AuthHistoryCreate) SetAppID(u uuid.UUID) *AuthHistoryCreate {
	ahc.mutation.SetAppID(u)
	return ahc
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (ahc *AuthHistoryCreate) SetNillableAppID(u *uuid.UUID) *AuthHistoryCreate {
	if u != nil {
		ahc.SetAppID(*u)
	}
	return ahc
}

// SetUserID sets the "user_id" field.
func (ahc *AuthHistoryCreate) SetUserID(u uuid.UUID) *AuthHistoryCreate {
	ahc.mutation.SetUserID(u)
	return ahc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ahc *AuthHistoryCreate) SetNillableUserID(u *uuid.UUID) *AuthHistoryCreate {
	if u != nil {
		ahc.SetUserID(*u)
	}
	return ahc
}

// SetResource sets the "resource" field.
func (ahc *AuthHistoryCreate) SetResource(s string) *AuthHistoryCreate {
	ahc.mutation.SetResource(s)
	return ahc
}

// SetNillableResource sets the "resource" field if the given value is not nil.
func (ahc *AuthHistoryCreate) SetNillableResource(s *string) *AuthHistoryCreate {
	if s != nil {
		ahc.SetResource(*s)
	}
	return ahc
}

// SetMethod sets the "method" field.
func (ahc *AuthHistoryCreate) SetMethod(s string) *AuthHistoryCreate {
	ahc.mutation.SetMethod(s)
	return ahc
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (ahc *AuthHistoryCreate) SetNillableMethod(s *string) *AuthHistoryCreate {
	if s != nil {
		ahc.SetMethod(*s)
	}
	return ahc
}

// SetAllowed sets the "allowed" field.
func (ahc *AuthHistoryCreate) SetAllowed(b bool) *AuthHistoryCreate {
	ahc.mutation.SetAllowed(b)
	return ahc
}

// SetNillableAllowed sets the "allowed" field if the given value is not nil.
func (ahc *AuthHistoryCreate) SetNillableAllowed(b *bool) *AuthHistoryCreate {
	if b != nil {
		ahc.SetAllowed(*b)
	}
	return ahc
}

// SetID sets the "id" field.
func (ahc *AuthHistoryCreate) SetID(u uint32) *AuthHistoryCreate {
	ahc.mutation.SetID(u)
	return ahc
}

// Mutation returns the AuthHistoryMutation object of the builder.
func (ahc *AuthHistoryCreate) Mutation() *AuthHistoryMutation {
	return ahc.mutation
}

// Save creates the AuthHistory in the database.
func (ahc *AuthHistoryCreate) Save(ctx context.Context) (*AuthHistory, error) {
	ahc.defaults()
	return withHooks(ctx, ahc.sqlSave, ahc.mutation, ahc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ahc *AuthHistoryCreate) SaveX(ctx context.Context) *AuthHistory {
	v, err := ahc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ahc *AuthHistoryCreate) Exec(ctx context.Context) error {
	_, err := ahc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ahc *AuthHistoryCreate) ExecX(ctx context.Context) {
	if err := ahc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ahc *AuthHistoryCreate) defaults() {
	if _, ok := ahc.mutation.CreatedAt(); !ok {
		v := authhistory.DefaultCreatedAt()
		ahc.mutation.SetCreatedAt(v)
	}
	if _, ok := ahc.mutation.UpdatedAt(); !ok {
		v := authhistory.DefaultUpdatedAt()
		ahc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ahc.mutation.DeletedAt(); !ok {
		v := authhistory.DefaultDeletedAt()
		ahc.mutation.SetDeletedAt(v)
	}
	if _, ok := ahc.mutation.EntID(); !ok {
		v := authhistory.DefaultEntID()
		ahc.mutation.SetEntID(v)
	}
	if _, ok := ahc.mutation.AppID(); !ok {
		v := authhistory.DefaultAppID()
		ahc.mutation.SetAppID(v)
	}
	if _, ok := ahc.mutation.UserID(); !ok {
		v := authhistory.DefaultUserID()
		ahc.mutation.SetUserID(v)
	}
	if _, ok := ahc.mutation.Resource(); !ok {
		v := authhistory.DefaultResource
		ahc.mutation.SetResource(v)
	}
	if _, ok := ahc.mutation.Method(); !ok {
		v := authhistory.DefaultMethod
		ahc.mutation.SetMethod(v)
	}
	if _, ok := ahc.mutation.Allowed(); !ok {
		v := authhistory.DefaultAllowed
		ahc.mutation.SetAllowed(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ahc *AuthHistoryCreate) check() error {
	if _, ok := ahc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "AuthHistory.created_at"`)}
	}
	if _, ok := ahc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "AuthHistory.updated_at"`)}
	}
	if _, ok := ahc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`generated: missing required field "AuthHistory.deleted_at"`)}
	}
	if _, ok := ahc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`generated: missing required field "AuthHistory.ent_id"`)}
	}
	return nil
}

func (ahc *AuthHistoryCreate) sqlSave(ctx context.Context) (*AuthHistory, error) {
	if err := ahc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ahc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ahc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	ahc.mutation.id = &_node.ID
	ahc.mutation.done = true
	return _node, nil
}

func (ahc *AuthHistoryCreate) createSpec() (*AuthHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &AuthHistory{config: ahc.config}
		_spec = sqlgraph.NewCreateSpec(authhistory.Table, sqlgraph.NewFieldSpec(authhistory.FieldID, field.TypeUint32))
	)
	_spec.OnConflict = ahc.conflict
	if id, ok := ahc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ahc.mutation.CreatedAt(); ok {
		_spec.SetField(authhistory.FieldCreatedAt, field.TypeUint32, value)
		_node.CreatedAt = value
	}
	if value, ok := ahc.mutation.UpdatedAt(); ok {
		_spec.SetField(authhistory.FieldUpdatedAt, field.TypeUint32, value)
		_node.UpdatedAt = value
	}
	if value, ok := ahc.mutation.DeletedAt(); ok {
		_spec.SetField(authhistory.FieldDeletedAt, field.TypeUint32, value)
		_node.DeletedAt = value
	}
	if value, ok := ahc.mutation.EntID(); ok {
		_spec.SetField(authhistory.FieldEntID, field.TypeUUID, value)
		_node.EntID = value
	}
	if value, ok := ahc.mutation.AppID(); ok {
		_spec.SetField(authhistory.FieldAppID, field.TypeUUID, value)
		_node.AppID = value
	}
	if value, ok := ahc.mutation.UserID(); ok {
		_spec.SetField(authhistory.FieldUserID, field.TypeUUID, value)
		_node.UserID = value
	}
	if value, ok := ahc.mutation.Resource(); ok {
		_spec.SetField(authhistory.FieldResource, field.TypeString, value)
		_node.Resource = value
	}
	if value, ok := ahc.mutation.Method(); ok {
		_spec.SetField(authhistory.FieldMethod, field.TypeString, value)
		_node.Method = value
	}
	if value, ok := ahc.mutation.Allowed(); ok {
		_spec.SetField(authhistory.FieldAllowed, field.TypeBool, value)
		_node.Allowed = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AuthHistory.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AuthHistoryUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ahc *AuthHistoryCreate) OnConflict(opts ...sql.ConflictOption) *AuthHistoryUpsertOne {
	ahc.conflict = opts
	return &AuthHistoryUpsertOne{
		create: ahc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AuthHistory.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ahc *AuthHistoryCreate) OnConflictColumns(columns ...string) *AuthHistoryUpsertOne {
	ahc.conflict = append(ahc.conflict, sql.ConflictColumns(columns...))
	return &AuthHistoryUpsertOne{
		create: ahc,
	}
}

type (
	// AuthHistoryUpsertOne is the builder for "upsert"-ing
	//  one AuthHistory node.
	AuthHistoryUpsertOne struct {
		create *AuthHistoryCreate
	}

	// AuthHistoryUpsert is the "OnConflict" setter.
	AuthHistoryUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *AuthHistoryUpsert) SetCreatedAt(v uint32) *AuthHistoryUpsert {
	u.Set(authhistory.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AuthHistoryUpsert) UpdateCreatedAt() *AuthHistoryUpsert {
	u.SetExcluded(authhistory.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AuthHistoryUpsert) AddCreatedAt(v uint32) *AuthHistoryUpsert {
	u.Add(authhistory.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AuthHistoryUpsert) SetUpdatedAt(v uint32) *AuthHistoryUpsert {
	u.Set(authhistory.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AuthHistoryUpsert) UpdateUpdatedAt() *AuthHistoryUpsert {
	u.SetExcluded(authhistory.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AuthHistoryUpsert) AddUpdatedAt(v uint32) *AuthHistoryUpsert {
	u.Add(authhistory.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AuthHistoryUpsert) SetDeletedAt(v uint32) *AuthHistoryUpsert {
	u.Set(authhistory.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AuthHistoryUpsert) UpdateDeletedAt() *AuthHistoryUpsert {
	u.SetExcluded(authhistory.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AuthHistoryUpsert) AddDeletedAt(v uint32) *AuthHistoryUpsert {
	u.Add(authhistory.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *AuthHistoryUpsert) SetEntID(v uuid.UUID) *AuthHistoryUpsert {
	u.Set(authhistory.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AuthHistoryUpsert) UpdateEntID() *AuthHistoryUpsert {
	u.SetExcluded(authhistory.FieldEntID)
	return u
}

// SetAppID sets the "app_id" field.
func (u *AuthHistoryUpsert) SetAppID(v uuid.UUID) *AuthHistoryUpsert {
	u.Set(authhistory.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AuthHistoryUpsert) UpdateAppID() *AuthHistoryUpsert {
	u.SetExcluded(authhistory.FieldAppID)
	return u
}

// ClearAppID clears the value of the "app_id" field.
func (u *AuthHistoryUpsert) ClearAppID() *AuthHistoryUpsert {
	u.SetNull(authhistory.FieldAppID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *AuthHistoryUpsert) SetUserID(v uuid.UUID) *AuthHistoryUpsert {
	u.Set(authhistory.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AuthHistoryUpsert) UpdateUserID() *AuthHistoryUpsert {
	u.SetExcluded(authhistory.FieldUserID)
	return u
}

// ClearUserID clears the value of the "user_id" field.
func (u *AuthHistoryUpsert) ClearUserID() *AuthHistoryUpsert {
	u.SetNull(authhistory.FieldUserID)
	return u
}

// SetResource sets the "resource" field.
func (u *AuthHistoryUpsert) SetResource(v string) *AuthHistoryUpsert {
	u.Set(authhistory.FieldResource, v)
	return u
}

// UpdateResource sets the "resource" field to the value that was provided on create.
func (u *AuthHistoryUpsert) UpdateResource() *AuthHistoryUpsert {
	u.SetExcluded(authhistory.FieldResource)
	return u
}

// ClearResource clears the value of the "resource" field.
func (u *AuthHistoryUpsert) ClearResource() *AuthHistoryUpsert {
	u.SetNull(authhistory.FieldResource)
	return u
}

// SetMethod sets the "method" field.
func (u *AuthHistoryUpsert) SetMethod(v string) *AuthHistoryUpsert {
	u.Set(authhistory.FieldMethod, v)
	return u
}

// UpdateMethod sets the "method" field to the value that was provided on create.
func (u *AuthHistoryUpsert) UpdateMethod() *AuthHistoryUpsert {
	u.SetExcluded(authhistory.FieldMethod)
	return u
}

// ClearMethod clears the value of the "method" field.
func (u *AuthHistoryUpsert) ClearMethod() *AuthHistoryUpsert {
	u.SetNull(authhistory.FieldMethod)
	return u
}

// SetAllowed sets the "allowed" field.
func (u *AuthHistoryUpsert) SetAllowed(v bool) *AuthHistoryUpsert {
	u.Set(authhistory.FieldAllowed, v)
	return u
}

// UpdateAllowed sets the "allowed" field to the value that was provided on create.
func (u *AuthHistoryUpsert) UpdateAllowed() *AuthHistoryUpsert {
	u.SetExcluded(authhistory.FieldAllowed)
	return u
}

// ClearAllowed clears the value of the "allowed" field.
func (u *AuthHistoryUpsert) ClearAllowed() *AuthHistoryUpsert {
	u.SetNull(authhistory.FieldAllowed)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AuthHistory.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(authhistory.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AuthHistoryUpsertOne) UpdateNewValues() *AuthHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(authhistory.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AuthHistory.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AuthHistoryUpsertOne) Ignore() *AuthHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AuthHistoryUpsertOne) DoNothing() *AuthHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AuthHistoryCreate.OnConflict
// documentation for more info.
func (u *AuthHistoryUpsertOne) Update(set func(*AuthHistoryUpsert)) *AuthHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AuthHistoryUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AuthHistoryUpsertOne) SetCreatedAt(v uint32) *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AuthHistoryUpsertOne) AddCreatedAt(v uint32) *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AuthHistoryUpsertOne) UpdateCreatedAt() *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AuthHistoryUpsertOne) SetUpdatedAt(v uint32) *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AuthHistoryUpsertOne) AddUpdatedAt(v uint32) *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AuthHistoryUpsertOne) UpdateUpdatedAt() *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AuthHistoryUpsertOne) SetDeletedAt(v uint32) *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AuthHistoryUpsertOne) AddDeletedAt(v uint32) *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AuthHistoryUpsertOne) UpdateDeletedAt() *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *AuthHistoryUpsertOne) SetEntID(v uuid.UUID) *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AuthHistoryUpsertOne) UpdateEntID() *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateEntID()
	})
}

// SetAppID sets the "app_id" field.
func (u *AuthHistoryUpsertOne) SetAppID(v uuid.UUID) *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AuthHistoryUpsertOne) UpdateAppID() *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *AuthHistoryUpsertOne) ClearAppID() *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.ClearAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *AuthHistoryUpsertOne) SetUserID(v uuid.UUID) *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AuthHistoryUpsertOne) UpdateUserID() *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *AuthHistoryUpsertOne) ClearUserID() *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.ClearUserID()
	})
}

// SetResource sets the "resource" field.
func (u *AuthHistoryUpsertOne) SetResource(v string) *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetResource(v)
	})
}

// UpdateResource sets the "resource" field to the value that was provided on create.
func (u *AuthHistoryUpsertOne) UpdateResource() *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateResource()
	})
}

// ClearResource clears the value of the "resource" field.
func (u *AuthHistoryUpsertOne) ClearResource() *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.ClearResource()
	})
}

// SetMethod sets the "method" field.
func (u *AuthHistoryUpsertOne) SetMethod(v string) *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetMethod(v)
	})
}

// UpdateMethod sets the "method" field to the value that was provided on create.
func (u *AuthHistoryUpsertOne) UpdateMethod() *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateMethod()
	})
}

// ClearMethod clears the value of the "method" field.
func (u *AuthHistoryUpsertOne) ClearMethod() *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.ClearMethod()
	})
}

// SetAllowed sets the "allowed" field.
func (u *AuthHistoryUpsertOne) SetAllowed(v bool) *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetAllowed(v)
	})
}

// UpdateAllowed sets the "allowed" field to the value that was provided on create.
func (u *AuthHistoryUpsertOne) UpdateAllowed() *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateAllowed()
	})
}

// ClearAllowed clears the value of the "allowed" field.
func (u *AuthHistoryUpsertOne) ClearAllowed() *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.ClearAllowed()
	})
}

// Exec executes the query.
func (u *AuthHistoryUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for AuthHistoryCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AuthHistoryUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AuthHistoryUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AuthHistoryUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AuthHistoryCreateBulk is the builder for creating many AuthHistory entities in bulk.
type AuthHistoryCreateBulk struct {
	config
	err      error
	builders []*AuthHistoryCreate
	conflict []sql.ConflictOption
}

// Save creates the AuthHistory entities in the database.
func (ahcb *AuthHistoryCreateBulk) Save(ctx context.Context) ([]*AuthHistory, error) {
	if ahcb.err != nil {
		return nil, ahcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ahcb.builders))
	nodes := make([]*AuthHistory, len(ahcb.builders))
	mutators := make([]Mutator, len(ahcb.builders))
	for i := range ahcb.builders {
		func(i int, root context.Context) {
			builder := ahcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AuthHistoryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ahcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ahcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ahcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint32(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ahcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ahcb *AuthHistoryCreateBulk) SaveX(ctx context.Context) []*AuthHistory {
	v, err := ahcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ahcb *AuthHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := ahcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ahcb *AuthHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := ahcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AuthHistory.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AuthHistoryUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ahcb *AuthHistoryCreateBulk) OnConflict(opts ...sql.ConflictOption) *AuthHistoryUpsertBulk {
	ahcb.conflict = opts
	return &AuthHistoryUpsertBulk{
		create: ahcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AuthHistory.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ahcb *AuthHistoryCreateBulk) OnConflictColumns(columns ...string) *AuthHistoryUpsertBulk {
	ahcb.conflict = append(ahcb.conflict, sql.ConflictColumns(columns...))
	return &AuthHistoryUpsertBulk{
		create: ahcb,
	}
}

// AuthHistoryUpsertBulk is the builder for "upsert"-ing
// a bulk of AuthHistory nodes.
type AuthHistoryUpsertBulk struct {
	create *AuthHistoryCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AuthHistory.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(authhistory.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AuthHistoryUpsertBulk) UpdateNewValues() *AuthHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(authhistory.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AuthHistory.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AuthHistoryUpsertBulk) Ignore() *AuthHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AuthHistoryUpsertBulk) DoNothing() *AuthHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AuthHistoryCreateBulk.OnConflict
// documentation for more info.
func (u *AuthHistoryUpsertBulk) Update(set func(*AuthHistoryUpsert)) *AuthHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AuthHistoryUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AuthHistoryUpsertBulk) SetCreatedAt(v uint32) *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AuthHistoryUpsertBulk) AddCreatedAt(v uint32) *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AuthHistoryUpsertBulk) UpdateCreatedAt() *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AuthHistoryUpsertBulk) SetUpdatedAt(v uint32) *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AuthHistoryUpsertBulk) AddUpdatedAt(v uint32) *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AuthHistoryUpsertBulk) UpdateUpdatedAt() *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AuthHistoryUpsertBulk) SetDeletedAt(v uint32) *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AuthHistoryUpsertBulk) AddDeletedAt(v uint32) *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AuthHistoryUpsertBulk) UpdateDeletedAt() *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *AuthHistoryUpsertBulk) SetEntID(v uuid.UUID) *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AuthHistoryUpsertBulk) UpdateEntID() *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateEntID()
	})
}

// SetAppID sets the "app_id" field.
func (u *AuthHistoryUpsertBulk) SetAppID(v uuid.UUID) *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AuthHistoryUpsertBulk) UpdateAppID() *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *AuthHistoryUpsertBulk) ClearAppID() *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.ClearAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *AuthHistoryUpsertBulk) SetUserID(v uuid.UUID) *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AuthHistoryUpsertBulk) UpdateUserID() *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *AuthHistoryUpsertBulk) ClearUserID() *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.ClearUserID()
	})
}

// SetResource sets the "resource" field.
func (u *AuthHistoryUpsertBulk) SetResource(v string) *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetResource(v)
	})
}

// UpdateResource sets the "resource" field to the value that was provided on create.
func (u *AuthHistoryUpsertBulk) UpdateResource() *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateResource()
	})
}

// ClearResource clears the value of the "resource" field.
func (u *AuthHistoryUpsertBulk) ClearResource() *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.ClearResource()
	})
}

// SetMethod sets the "method" field.
func (u *AuthHistoryUpsertBulk) SetMethod(v string) *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetMethod(v)
	})
}

// UpdateMethod sets the "method" field to the value that was provided on create.
func (u *AuthHistoryUpsertBulk) UpdateMethod() *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateMethod()
	})
}

// ClearMethod clears the value of the "method" field.
func (u *AuthHistoryUpsertBulk) ClearMethod() *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.ClearMethod()
	})
}

// SetAllowed sets the "allowed" field.
func (u *AuthHistoryUpsertBulk) SetAllowed(v bool) *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetAllowed(v)
	})
}

// UpdateAllowed sets the "allowed" field to the value that was provided on create.
func (u *AuthHistoryUpsertBulk) UpdateAllowed() *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateAllowed()
	})
}

// ClearAllowed clears the value of the "allowed" field.
func (u *AuthHistoryUpsertBulk) ClearAllowed() *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.ClearAllowed()
	})
}

// Exec executes the query.
func (u *AuthHistoryUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("generated: OnConflict was set for builder %d. Set it on the AuthHistoryCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for AuthHistoryCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AuthHistoryUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
