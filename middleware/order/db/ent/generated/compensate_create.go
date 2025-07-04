// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/compensate"
	"github.com/google/uuid"
)

// CompensateCreate is the builder for creating a Compensate entity.
type CompensateCreate struct {
	config
	mutation *CompensateMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetEntID sets the "ent_id" field.
func (cc *CompensateCreate) SetEntID(u uuid.UUID) *CompensateCreate {
	cc.mutation.SetEntID(u)
	return cc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cc *CompensateCreate) SetNillableEntID(u *uuid.UUID) *CompensateCreate {
	if u != nil {
		cc.SetEntID(*u)
	}
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *CompensateCreate) SetCreatedAt(u uint32) *CompensateCreate {
	cc.mutation.SetCreatedAt(u)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CompensateCreate) SetNillableCreatedAt(u *uint32) *CompensateCreate {
	if u != nil {
		cc.SetCreatedAt(*u)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CompensateCreate) SetUpdatedAt(u uint32) *CompensateCreate {
	cc.mutation.SetUpdatedAt(u)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CompensateCreate) SetNillableUpdatedAt(u *uint32) *CompensateCreate {
	if u != nil {
		cc.SetUpdatedAt(*u)
	}
	return cc
}

// SetDeletedAt sets the "deleted_at" field.
func (cc *CompensateCreate) SetDeletedAt(u uint32) *CompensateCreate {
	cc.mutation.SetDeletedAt(u)
	return cc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cc *CompensateCreate) SetNillableDeletedAt(u *uint32) *CompensateCreate {
	if u != nil {
		cc.SetDeletedAt(*u)
	}
	return cc
}

// SetOrderID sets the "order_id" field.
func (cc *CompensateCreate) SetOrderID(u uuid.UUID) *CompensateCreate {
	cc.mutation.SetOrderID(u)
	return cc
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (cc *CompensateCreate) SetNillableOrderID(u *uuid.UUID) *CompensateCreate {
	if u != nil {
		cc.SetOrderID(*u)
	}
	return cc
}

// SetCompensateFromID sets the "compensate_from_id" field.
func (cc *CompensateCreate) SetCompensateFromID(u uuid.UUID) *CompensateCreate {
	cc.mutation.SetCompensateFromID(u)
	return cc
}

// SetNillableCompensateFromID sets the "compensate_from_id" field if the given value is not nil.
func (cc *CompensateCreate) SetNillableCompensateFromID(u *uuid.UUID) *CompensateCreate {
	if u != nil {
		cc.SetCompensateFromID(*u)
	}
	return cc
}

// SetCompensateType sets the "compensate_type" field.
func (cc *CompensateCreate) SetCompensateType(s string) *CompensateCreate {
	cc.mutation.SetCompensateType(s)
	return cc
}

// SetNillableCompensateType sets the "compensate_type" field if the given value is not nil.
func (cc *CompensateCreate) SetNillableCompensateType(s *string) *CompensateCreate {
	if s != nil {
		cc.SetCompensateType(*s)
	}
	return cc
}

// SetCompensateSeconds sets the "compensate_seconds" field.
func (cc *CompensateCreate) SetCompensateSeconds(u uint32) *CompensateCreate {
	cc.mutation.SetCompensateSeconds(u)
	return cc
}

// SetNillableCompensateSeconds sets the "compensate_seconds" field if the given value is not nil.
func (cc *CompensateCreate) SetNillableCompensateSeconds(u *uint32) *CompensateCreate {
	if u != nil {
		cc.SetCompensateSeconds(*u)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CompensateCreate) SetID(u uint32) *CompensateCreate {
	cc.mutation.SetID(u)
	return cc
}

// Mutation returns the CompensateMutation object of the builder.
func (cc *CompensateCreate) Mutation() *CompensateMutation {
	return cc.mutation
}

// Save creates the Compensate in the database.
func (cc *CompensateCreate) Save(ctx context.Context) (*Compensate, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CompensateCreate) SaveX(ctx context.Context) *Compensate {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CompensateCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CompensateCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CompensateCreate) defaults() {
	if _, ok := cc.mutation.EntID(); !ok {
		v := compensate.DefaultEntID()
		cc.mutation.SetEntID(v)
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := compensate.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := compensate.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.DeletedAt(); !ok {
		v := compensate.DefaultDeletedAt()
		cc.mutation.SetDeletedAt(v)
	}
	if _, ok := cc.mutation.OrderID(); !ok {
		v := compensate.DefaultOrderID()
		cc.mutation.SetOrderID(v)
	}
	if _, ok := cc.mutation.CompensateFromID(); !ok {
		v := compensate.DefaultCompensateFromID()
		cc.mutation.SetCompensateFromID(v)
	}
	if _, ok := cc.mutation.CompensateType(); !ok {
		v := compensate.DefaultCompensateType
		cc.mutation.SetCompensateType(v)
	}
	if _, ok := cc.mutation.CompensateSeconds(); !ok {
		v := compensate.DefaultCompensateSeconds
		cc.mutation.SetCompensateSeconds(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CompensateCreate) check() error {
	if _, ok := cc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`generated: missing required field "Compensate.ent_id"`)}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "Compensate.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "Compensate.updated_at"`)}
	}
	if _, ok := cc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`generated: missing required field "Compensate.deleted_at"`)}
	}
	return nil
}

func (cc *CompensateCreate) sqlSave(ctx context.Context) (*Compensate, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CompensateCreate) createSpec() (*Compensate, *sqlgraph.CreateSpec) {
	var (
		_node = &Compensate{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(compensate.Table, sqlgraph.NewFieldSpec(compensate.FieldID, field.TypeUint32))
	)
	_spec.OnConflict = cc.conflict
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.EntID(); ok {
		_spec.SetField(compensate.FieldEntID, field.TypeUUID, value)
		_node.EntID = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(compensate.FieldCreatedAt, field.TypeUint32, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(compensate.FieldUpdatedAt, field.TypeUint32, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.DeletedAt(); ok {
		_spec.SetField(compensate.FieldDeletedAt, field.TypeUint32, value)
		_node.DeletedAt = value
	}
	if value, ok := cc.mutation.OrderID(); ok {
		_spec.SetField(compensate.FieldOrderID, field.TypeUUID, value)
		_node.OrderID = value
	}
	if value, ok := cc.mutation.CompensateFromID(); ok {
		_spec.SetField(compensate.FieldCompensateFromID, field.TypeUUID, value)
		_node.CompensateFromID = value
	}
	if value, ok := cc.mutation.CompensateType(); ok {
		_spec.SetField(compensate.FieldCompensateType, field.TypeString, value)
		_node.CompensateType = value
	}
	if value, ok := cc.mutation.CompensateSeconds(); ok {
		_spec.SetField(compensate.FieldCompensateSeconds, field.TypeUint32, value)
		_node.CompensateSeconds = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Compensate.Create().
//		SetEntID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CompensateUpsert) {
//			SetEntID(v+v).
//		}).
//		Exec(ctx)
func (cc *CompensateCreate) OnConflict(opts ...sql.ConflictOption) *CompensateUpsertOne {
	cc.conflict = opts
	return &CompensateUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Compensate.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cc *CompensateCreate) OnConflictColumns(columns ...string) *CompensateUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &CompensateUpsertOne{
		create: cc,
	}
}

type (
	// CompensateUpsertOne is the builder for "upsert"-ing
	//  one Compensate node.
	CompensateUpsertOne struct {
		create *CompensateCreate
	}

	// CompensateUpsert is the "OnConflict" setter.
	CompensateUpsert struct {
		*sql.UpdateSet
	}
)

// SetEntID sets the "ent_id" field.
func (u *CompensateUpsert) SetEntID(v uuid.UUID) *CompensateUpsert {
	u.Set(compensate.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *CompensateUpsert) UpdateEntID() *CompensateUpsert {
	u.SetExcluded(compensate.FieldEntID)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CompensateUpsert) SetCreatedAt(v uint32) *CompensateUpsert {
	u.Set(compensate.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CompensateUpsert) UpdateCreatedAt() *CompensateUpsert {
	u.SetExcluded(compensate.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CompensateUpsert) AddCreatedAt(v uint32) *CompensateUpsert {
	u.Add(compensate.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CompensateUpsert) SetUpdatedAt(v uint32) *CompensateUpsert {
	u.Set(compensate.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CompensateUpsert) UpdateUpdatedAt() *CompensateUpsert {
	u.SetExcluded(compensate.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CompensateUpsert) AddUpdatedAt(v uint32) *CompensateUpsert {
	u.Add(compensate.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CompensateUpsert) SetDeletedAt(v uint32) *CompensateUpsert {
	u.Set(compensate.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CompensateUpsert) UpdateDeletedAt() *CompensateUpsert {
	u.SetExcluded(compensate.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CompensateUpsert) AddDeletedAt(v uint32) *CompensateUpsert {
	u.Add(compensate.FieldDeletedAt, v)
	return u
}

// SetOrderID sets the "order_id" field.
func (u *CompensateUpsert) SetOrderID(v uuid.UUID) *CompensateUpsert {
	u.Set(compensate.FieldOrderID, v)
	return u
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *CompensateUpsert) UpdateOrderID() *CompensateUpsert {
	u.SetExcluded(compensate.FieldOrderID)
	return u
}

// ClearOrderID clears the value of the "order_id" field.
func (u *CompensateUpsert) ClearOrderID() *CompensateUpsert {
	u.SetNull(compensate.FieldOrderID)
	return u
}

// SetCompensateFromID sets the "compensate_from_id" field.
func (u *CompensateUpsert) SetCompensateFromID(v uuid.UUID) *CompensateUpsert {
	u.Set(compensate.FieldCompensateFromID, v)
	return u
}

// UpdateCompensateFromID sets the "compensate_from_id" field to the value that was provided on create.
func (u *CompensateUpsert) UpdateCompensateFromID() *CompensateUpsert {
	u.SetExcluded(compensate.FieldCompensateFromID)
	return u
}

// ClearCompensateFromID clears the value of the "compensate_from_id" field.
func (u *CompensateUpsert) ClearCompensateFromID() *CompensateUpsert {
	u.SetNull(compensate.FieldCompensateFromID)
	return u
}

// SetCompensateType sets the "compensate_type" field.
func (u *CompensateUpsert) SetCompensateType(v string) *CompensateUpsert {
	u.Set(compensate.FieldCompensateType, v)
	return u
}

// UpdateCompensateType sets the "compensate_type" field to the value that was provided on create.
func (u *CompensateUpsert) UpdateCompensateType() *CompensateUpsert {
	u.SetExcluded(compensate.FieldCompensateType)
	return u
}

// ClearCompensateType clears the value of the "compensate_type" field.
func (u *CompensateUpsert) ClearCompensateType() *CompensateUpsert {
	u.SetNull(compensate.FieldCompensateType)
	return u
}

// SetCompensateSeconds sets the "compensate_seconds" field.
func (u *CompensateUpsert) SetCompensateSeconds(v uint32) *CompensateUpsert {
	u.Set(compensate.FieldCompensateSeconds, v)
	return u
}

// UpdateCompensateSeconds sets the "compensate_seconds" field to the value that was provided on create.
func (u *CompensateUpsert) UpdateCompensateSeconds() *CompensateUpsert {
	u.SetExcluded(compensate.FieldCompensateSeconds)
	return u
}

// AddCompensateSeconds adds v to the "compensate_seconds" field.
func (u *CompensateUpsert) AddCompensateSeconds(v uint32) *CompensateUpsert {
	u.Add(compensate.FieldCompensateSeconds, v)
	return u
}

// ClearCompensateSeconds clears the value of the "compensate_seconds" field.
func (u *CompensateUpsert) ClearCompensateSeconds() *CompensateUpsert {
	u.SetNull(compensate.FieldCompensateSeconds)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Compensate.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(compensate.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CompensateUpsertOne) UpdateNewValues() *CompensateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(compensate.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Compensate.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *CompensateUpsertOne) Ignore() *CompensateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CompensateUpsertOne) DoNothing() *CompensateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CompensateCreate.OnConflict
// documentation for more info.
func (u *CompensateUpsertOne) Update(set func(*CompensateUpsert)) *CompensateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CompensateUpsert{UpdateSet: update})
	}))
	return u
}

// SetEntID sets the "ent_id" field.
func (u *CompensateUpsertOne) SetEntID(v uuid.UUID) *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *CompensateUpsertOne) UpdateEntID() *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateEntID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *CompensateUpsertOne) SetCreatedAt(v uint32) *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CompensateUpsertOne) AddCreatedAt(v uint32) *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CompensateUpsertOne) UpdateCreatedAt() *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CompensateUpsertOne) SetUpdatedAt(v uint32) *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CompensateUpsertOne) AddUpdatedAt(v uint32) *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CompensateUpsertOne) UpdateUpdatedAt() *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CompensateUpsertOne) SetDeletedAt(v uint32) *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CompensateUpsertOne) AddDeletedAt(v uint32) *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CompensateUpsertOne) UpdateDeletedAt() *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetOrderID sets the "order_id" field.
func (u *CompensateUpsertOne) SetOrderID(v uuid.UUID) *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.SetOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *CompensateUpsertOne) UpdateOrderID() *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateOrderID()
	})
}

// ClearOrderID clears the value of the "order_id" field.
func (u *CompensateUpsertOne) ClearOrderID() *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.ClearOrderID()
	})
}

// SetCompensateFromID sets the "compensate_from_id" field.
func (u *CompensateUpsertOne) SetCompensateFromID(v uuid.UUID) *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.SetCompensateFromID(v)
	})
}

// UpdateCompensateFromID sets the "compensate_from_id" field to the value that was provided on create.
func (u *CompensateUpsertOne) UpdateCompensateFromID() *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateCompensateFromID()
	})
}

// ClearCompensateFromID clears the value of the "compensate_from_id" field.
func (u *CompensateUpsertOne) ClearCompensateFromID() *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.ClearCompensateFromID()
	})
}

// SetCompensateType sets the "compensate_type" field.
func (u *CompensateUpsertOne) SetCompensateType(v string) *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.SetCompensateType(v)
	})
}

// UpdateCompensateType sets the "compensate_type" field to the value that was provided on create.
func (u *CompensateUpsertOne) UpdateCompensateType() *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateCompensateType()
	})
}

// ClearCompensateType clears the value of the "compensate_type" field.
func (u *CompensateUpsertOne) ClearCompensateType() *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.ClearCompensateType()
	})
}

// SetCompensateSeconds sets the "compensate_seconds" field.
func (u *CompensateUpsertOne) SetCompensateSeconds(v uint32) *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.SetCompensateSeconds(v)
	})
}

// AddCompensateSeconds adds v to the "compensate_seconds" field.
func (u *CompensateUpsertOne) AddCompensateSeconds(v uint32) *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.AddCompensateSeconds(v)
	})
}

// UpdateCompensateSeconds sets the "compensate_seconds" field to the value that was provided on create.
func (u *CompensateUpsertOne) UpdateCompensateSeconds() *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateCompensateSeconds()
	})
}

// ClearCompensateSeconds clears the value of the "compensate_seconds" field.
func (u *CompensateUpsertOne) ClearCompensateSeconds() *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.ClearCompensateSeconds()
	})
}

// Exec executes the query.
func (u *CompensateUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for CompensateCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CompensateUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CompensateUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CompensateUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CompensateCreateBulk is the builder for creating many Compensate entities in bulk.
type CompensateCreateBulk struct {
	config
	err      error
	builders []*CompensateCreate
	conflict []sql.ConflictOption
}

// Save creates the Compensate entities in the database.
func (ccb *CompensateCreateBulk) Save(ctx context.Context) ([]*Compensate, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Compensate, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CompensateMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CompensateCreateBulk) SaveX(ctx context.Context) []*Compensate {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CompensateCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CompensateCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Compensate.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CompensateUpsert) {
//			SetEntID(v+v).
//		}).
//		Exec(ctx)
func (ccb *CompensateCreateBulk) OnConflict(opts ...sql.ConflictOption) *CompensateUpsertBulk {
	ccb.conflict = opts
	return &CompensateUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Compensate.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ccb *CompensateCreateBulk) OnConflictColumns(columns ...string) *CompensateUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &CompensateUpsertBulk{
		create: ccb,
	}
}

// CompensateUpsertBulk is the builder for "upsert"-ing
// a bulk of Compensate nodes.
type CompensateUpsertBulk struct {
	create *CompensateCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Compensate.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(compensate.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CompensateUpsertBulk) UpdateNewValues() *CompensateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(compensate.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Compensate.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *CompensateUpsertBulk) Ignore() *CompensateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CompensateUpsertBulk) DoNothing() *CompensateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CompensateCreateBulk.OnConflict
// documentation for more info.
func (u *CompensateUpsertBulk) Update(set func(*CompensateUpsert)) *CompensateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CompensateUpsert{UpdateSet: update})
	}))
	return u
}

// SetEntID sets the "ent_id" field.
func (u *CompensateUpsertBulk) SetEntID(v uuid.UUID) *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *CompensateUpsertBulk) UpdateEntID() *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateEntID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *CompensateUpsertBulk) SetCreatedAt(v uint32) *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CompensateUpsertBulk) AddCreatedAt(v uint32) *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CompensateUpsertBulk) UpdateCreatedAt() *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CompensateUpsertBulk) SetUpdatedAt(v uint32) *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CompensateUpsertBulk) AddUpdatedAt(v uint32) *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CompensateUpsertBulk) UpdateUpdatedAt() *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CompensateUpsertBulk) SetDeletedAt(v uint32) *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CompensateUpsertBulk) AddDeletedAt(v uint32) *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CompensateUpsertBulk) UpdateDeletedAt() *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetOrderID sets the "order_id" field.
func (u *CompensateUpsertBulk) SetOrderID(v uuid.UUID) *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.SetOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *CompensateUpsertBulk) UpdateOrderID() *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateOrderID()
	})
}

// ClearOrderID clears the value of the "order_id" field.
func (u *CompensateUpsertBulk) ClearOrderID() *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.ClearOrderID()
	})
}

// SetCompensateFromID sets the "compensate_from_id" field.
func (u *CompensateUpsertBulk) SetCompensateFromID(v uuid.UUID) *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.SetCompensateFromID(v)
	})
}

// UpdateCompensateFromID sets the "compensate_from_id" field to the value that was provided on create.
func (u *CompensateUpsertBulk) UpdateCompensateFromID() *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateCompensateFromID()
	})
}

// ClearCompensateFromID clears the value of the "compensate_from_id" field.
func (u *CompensateUpsertBulk) ClearCompensateFromID() *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.ClearCompensateFromID()
	})
}

// SetCompensateType sets the "compensate_type" field.
func (u *CompensateUpsertBulk) SetCompensateType(v string) *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.SetCompensateType(v)
	})
}

// UpdateCompensateType sets the "compensate_type" field to the value that was provided on create.
func (u *CompensateUpsertBulk) UpdateCompensateType() *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateCompensateType()
	})
}

// ClearCompensateType clears the value of the "compensate_type" field.
func (u *CompensateUpsertBulk) ClearCompensateType() *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.ClearCompensateType()
	})
}

// SetCompensateSeconds sets the "compensate_seconds" field.
func (u *CompensateUpsertBulk) SetCompensateSeconds(v uint32) *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.SetCompensateSeconds(v)
	})
}

// AddCompensateSeconds adds v to the "compensate_seconds" field.
func (u *CompensateUpsertBulk) AddCompensateSeconds(v uint32) *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.AddCompensateSeconds(v)
	})
}

// UpdateCompensateSeconds sets the "compensate_seconds" field to the value that was provided on create.
func (u *CompensateUpsertBulk) UpdateCompensateSeconds() *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateCompensateSeconds()
	})
}

// ClearCompensateSeconds clears the value of the "compensate_seconds" field.
func (u *CompensateUpsertBulk) ClearCompensateSeconds() *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.ClearCompensateSeconds()
	})
}

// Exec executes the query.
func (u *CompensateUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("generated: OnConflict was set for builder %d. Set it on the CompensateCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for CompensateCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CompensateUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
