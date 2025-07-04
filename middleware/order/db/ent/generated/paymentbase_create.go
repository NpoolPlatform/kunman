// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/paymentbase"
	"github.com/google/uuid"
)

// PaymentBaseCreate is the builder for creating a PaymentBase entity.
type PaymentBaseCreate struct {
	config
	mutation *PaymentBaseMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetEntID sets the "ent_id" field.
func (pbc *PaymentBaseCreate) SetEntID(u uuid.UUID) *PaymentBaseCreate {
	pbc.mutation.SetEntID(u)
	return pbc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (pbc *PaymentBaseCreate) SetNillableEntID(u *uuid.UUID) *PaymentBaseCreate {
	if u != nil {
		pbc.SetEntID(*u)
	}
	return pbc
}

// SetCreatedAt sets the "created_at" field.
func (pbc *PaymentBaseCreate) SetCreatedAt(u uint32) *PaymentBaseCreate {
	pbc.mutation.SetCreatedAt(u)
	return pbc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pbc *PaymentBaseCreate) SetNillableCreatedAt(u *uint32) *PaymentBaseCreate {
	if u != nil {
		pbc.SetCreatedAt(*u)
	}
	return pbc
}

// SetUpdatedAt sets the "updated_at" field.
func (pbc *PaymentBaseCreate) SetUpdatedAt(u uint32) *PaymentBaseCreate {
	pbc.mutation.SetUpdatedAt(u)
	return pbc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pbc *PaymentBaseCreate) SetNillableUpdatedAt(u *uint32) *PaymentBaseCreate {
	if u != nil {
		pbc.SetUpdatedAt(*u)
	}
	return pbc
}

// SetDeletedAt sets the "deleted_at" field.
func (pbc *PaymentBaseCreate) SetDeletedAt(u uint32) *PaymentBaseCreate {
	pbc.mutation.SetDeletedAt(u)
	return pbc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pbc *PaymentBaseCreate) SetNillableDeletedAt(u *uint32) *PaymentBaseCreate {
	if u != nil {
		pbc.SetDeletedAt(*u)
	}
	return pbc
}

// SetOrderID sets the "order_id" field.
func (pbc *PaymentBaseCreate) SetOrderID(u uuid.UUID) *PaymentBaseCreate {
	pbc.mutation.SetOrderID(u)
	return pbc
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (pbc *PaymentBaseCreate) SetNillableOrderID(u *uuid.UUID) *PaymentBaseCreate {
	if u != nil {
		pbc.SetOrderID(*u)
	}
	return pbc
}

// SetObseleteState sets the "obselete_state" field.
func (pbc *PaymentBaseCreate) SetObseleteState(s string) *PaymentBaseCreate {
	pbc.mutation.SetObseleteState(s)
	return pbc
}

// SetNillableObseleteState sets the "obselete_state" field if the given value is not nil.
func (pbc *PaymentBaseCreate) SetNillableObseleteState(s *string) *PaymentBaseCreate {
	if s != nil {
		pbc.SetObseleteState(*s)
	}
	return pbc
}

// SetID sets the "id" field.
func (pbc *PaymentBaseCreate) SetID(u uint32) *PaymentBaseCreate {
	pbc.mutation.SetID(u)
	return pbc
}

// Mutation returns the PaymentBaseMutation object of the builder.
func (pbc *PaymentBaseCreate) Mutation() *PaymentBaseMutation {
	return pbc.mutation
}

// Save creates the PaymentBase in the database.
func (pbc *PaymentBaseCreate) Save(ctx context.Context) (*PaymentBase, error) {
	pbc.defaults()
	return withHooks(ctx, pbc.sqlSave, pbc.mutation, pbc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pbc *PaymentBaseCreate) SaveX(ctx context.Context) *PaymentBase {
	v, err := pbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pbc *PaymentBaseCreate) Exec(ctx context.Context) error {
	_, err := pbc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pbc *PaymentBaseCreate) ExecX(ctx context.Context) {
	if err := pbc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pbc *PaymentBaseCreate) defaults() {
	if _, ok := pbc.mutation.EntID(); !ok {
		v := paymentbase.DefaultEntID()
		pbc.mutation.SetEntID(v)
	}
	if _, ok := pbc.mutation.CreatedAt(); !ok {
		v := paymentbase.DefaultCreatedAt()
		pbc.mutation.SetCreatedAt(v)
	}
	if _, ok := pbc.mutation.UpdatedAt(); !ok {
		v := paymentbase.DefaultUpdatedAt()
		pbc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pbc.mutation.DeletedAt(); !ok {
		v := paymentbase.DefaultDeletedAt()
		pbc.mutation.SetDeletedAt(v)
	}
	if _, ok := pbc.mutation.OrderID(); !ok {
		v := paymentbase.DefaultOrderID()
		pbc.mutation.SetOrderID(v)
	}
	if _, ok := pbc.mutation.ObseleteState(); !ok {
		v := paymentbase.DefaultObseleteState
		pbc.mutation.SetObseleteState(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pbc *PaymentBaseCreate) check() error {
	if _, ok := pbc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`generated: missing required field "PaymentBase.ent_id"`)}
	}
	if _, ok := pbc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "PaymentBase.created_at"`)}
	}
	if _, ok := pbc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "PaymentBase.updated_at"`)}
	}
	if _, ok := pbc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`generated: missing required field "PaymentBase.deleted_at"`)}
	}
	return nil
}

func (pbc *PaymentBaseCreate) sqlSave(ctx context.Context) (*PaymentBase, error) {
	if err := pbc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pbc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	pbc.mutation.id = &_node.ID
	pbc.mutation.done = true
	return _node, nil
}

func (pbc *PaymentBaseCreate) createSpec() (*PaymentBase, *sqlgraph.CreateSpec) {
	var (
		_node = &PaymentBase{config: pbc.config}
		_spec = sqlgraph.NewCreateSpec(paymentbase.Table, sqlgraph.NewFieldSpec(paymentbase.FieldID, field.TypeUint32))
	)
	_spec.OnConflict = pbc.conflict
	if id, ok := pbc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pbc.mutation.EntID(); ok {
		_spec.SetField(paymentbase.FieldEntID, field.TypeUUID, value)
		_node.EntID = value
	}
	if value, ok := pbc.mutation.CreatedAt(); ok {
		_spec.SetField(paymentbase.FieldCreatedAt, field.TypeUint32, value)
		_node.CreatedAt = value
	}
	if value, ok := pbc.mutation.UpdatedAt(); ok {
		_spec.SetField(paymentbase.FieldUpdatedAt, field.TypeUint32, value)
		_node.UpdatedAt = value
	}
	if value, ok := pbc.mutation.DeletedAt(); ok {
		_spec.SetField(paymentbase.FieldDeletedAt, field.TypeUint32, value)
		_node.DeletedAt = value
	}
	if value, ok := pbc.mutation.OrderID(); ok {
		_spec.SetField(paymentbase.FieldOrderID, field.TypeUUID, value)
		_node.OrderID = value
	}
	if value, ok := pbc.mutation.ObseleteState(); ok {
		_spec.SetField(paymentbase.FieldObseleteState, field.TypeString, value)
		_node.ObseleteState = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.PaymentBase.Create().
//		SetEntID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PaymentBaseUpsert) {
//			SetEntID(v+v).
//		}).
//		Exec(ctx)
func (pbc *PaymentBaseCreate) OnConflict(opts ...sql.ConflictOption) *PaymentBaseUpsertOne {
	pbc.conflict = opts
	return &PaymentBaseUpsertOne{
		create: pbc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.PaymentBase.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pbc *PaymentBaseCreate) OnConflictColumns(columns ...string) *PaymentBaseUpsertOne {
	pbc.conflict = append(pbc.conflict, sql.ConflictColumns(columns...))
	return &PaymentBaseUpsertOne{
		create: pbc,
	}
}

type (
	// PaymentBaseUpsertOne is the builder for "upsert"-ing
	//  one PaymentBase node.
	PaymentBaseUpsertOne struct {
		create *PaymentBaseCreate
	}

	// PaymentBaseUpsert is the "OnConflict" setter.
	PaymentBaseUpsert struct {
		*sql.UpdateSet
	}
)

// SetEntID sets the "ent_id" field.
func (u *PaymentBaseUpsert) SetEntID(v uuid.UUID) *PaymentBaseUpsert {
	u.Set(paymentbase.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *PaymentBaseUpsert) UpdateEntID() *PaymentBaseUpsert {
	u.SetExcluded(paymentbase.FieldEntID)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *PaymentBaseUpsert) SetCreatedAt(v uint32) *PaymentBaseUpsert {
	u.Set(paymentbase.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PaymentBaseUpsert) UpdateCreatedAt() *PaymentBaseUpsert {
	u.SetExcluded(paymentbase.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *PaymentBaseUpsert) AddCreatedAt(v uint32) *PaymentBaseUpsert {
	u.Add(paymentbase.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PaymentBaseUpsert) SetUpdatedAt(v uint32) *PaymentBaseUpsert {
	u.Set(paymentbase.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PaymentBaseUpsert) UpdateUpdatedAt() *PaymentBaseUpsert {
	u.SetExcluded(paymentbase.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *PaymentBaseUpsert) AddUpdatedAt(v uint32) *PaymentBaseUpsert {
	u.Add(paymentbase.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *PaymentBaseUpsert) SetDeletedAt(v uint32) *PaymentBaseUpsert {
	u.Set(paymentbase.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *PaymentBaseUpsert) UpdateDeletedAt() *PaymentBaseUpsert {
	u.SetExcluded(paymentbase.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *PaymentBaseUpsert) AddDeletedAt(v uint32) *PaymentBaseUpsert {
	u.Add(paymentbase.FieldDeletedAt, v)
	return u
}

// SetOrderID sets the "order_id" field.
func (u *PaymentBaseUpsert) SetOrderID(v uuid.UUID) *PaymentBaseUpsert {
	u.Set(paymentbase.FieldOrderID, v)
	return u
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *PaymentBaseUpsert) UpdateOrderID() *PaymentBaseUpsert {
	u.SetExcluded(paymentbase.FieldOrderID)
	return u
}

// ClearOrderID clears the value of the "order_id" field.
func (u *PaymentBaseUpsert) ClearOrderID() *PaymentBaseUpsert {
	u.SetNull(paymentbase.FieldOrderID)
	return u
}

// SetObseleteState sets the "obselete_state" field.
func (u *PaymentBaseUpsert) SetObseleteState(v string) *PaymentBaseUpsert {
	u.Set(paymentbase.FieldObseleteState, v)
	return u
}

// UpdateObseleteState sets the "obselete_state" field to the value that was provided on create.
func (u *PaymentBaseUpsert) UpdateObseleteState() *PaymentBaseUpsert {
	u.SetExcluded(paymentbase.FieldObseleteState)
	return u
}

// ClearObseleteState clears the value of the "obselete_state" field.
func (u *PaymentBaseUpsert) ClearObseleteState() *PaymentBaseUpsert {
	u.SetNull(paymentbase.FieldObseleteState)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.PaymentBase.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(paymentbase.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PaymentBaseUpsertOne) UpdateNewValues() *PaymentBaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(paymentbase.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.PaymentBase.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PaymentBaseUpsertOne) Ignore() *PaymentBaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PaymentBaseUpsertOne) DoNothing() *PaymentBaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PaymentBaseCreate.OnConflict
// documentation for more info.
func (u *PaymentBaseUpsertOne) Update(set func(*PaymentBaseUpsert)) *PaymentBaseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PaymentBaseUpsert{UpdateSet: update})
	}))
	return u
}

// SetEntID sets the "ent_id" field.
func (u *PaymentBaseUpsertOne) SetEntID(v uuid.UUID) *PaymentBaseUpsertOne {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *PaymentBaseUpsertOne) UpdateEntID() *PaymentBaseUpsertOne {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.UpdateEntID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *PaymentBaseUpsertOne) SetCreatedAt(v uint32) *PaymentBaseUpsertOne {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *PaymentBaseUpsertOne) AddCreatedAt(v uint32) *PaymentBaseUpsertOne {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PaymentBaseUpsertOne) UpdateCreatedAt() *PaymentBaseUpsertOne {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PaymentBaseUpsertOne) SetUpdatedAt(v uint32) *PaymentBaseUpsertOne {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *PaymentBaseUpsertOne) AddUpdatedAt(v uint32) *PaymentBaseUpsertOne {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PaymentBaseUpsertOne) UpdateUpdatedAt() *PaymentBaseUpsertOne {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *PaymentBaseUpsertOne) SetDeletedAt(v uint32) *PaymentBaseUpsertOne {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *PaymentBaseUpsertOne) AddDeletedAt(v uint32) *PaymentBaseUpsertOne {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *PaymentBaseUpsertOne) UpdateDeletedAt() *PaymentBaseUpsertOne {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetOrderID sets the "order_id" field.
func (u *PaymentBaseUpsertOne) SetOrderID(v uuid.UUID) *PaymentBaseUpsertOne {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.SetOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *PaymentBaseUpsertOne) UpdateOrderID() *PaymentBaseUpsertOne {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.UpdateOrderID()
	})
}

// ClearOrderID clears the value of the "order_id" field.
func (u *PaymentBaseUpsertOne) ClearOrderID() *PaymentBaseUpsertOne {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.ClearOrderID()
	})
}

// SetObseleteState sets the "obselete_state" field.
func (u *PaymentBaseUpsertOne) SetObseleteState(v string) *PaymentBaseUpsertOne {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.SetObseleteState(v)
	})
}

// UpdateObseleteState sets the "obselete_state" field to the value that was provided on create.
func (u *PaymentBaseUpsertOne) UpdateObseleteState() *PaymentBaseUpsertOne {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.UpdateObseleteState()
	})
}

// ClearObseleteState clears the value of the "obselete_state" field.
func (u *PaymentBaseUpsertOne) ClearObseleteState() *PaymentBaseUpsertOne {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.ClearObseleteState()
	})
}

// Exec executes the query.
func (u *PaymentBaseUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for PaymentBaseCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PaymentBaseUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PaymentBaseUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PaymentBaseUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PaymentBaseCreateBulk is the builder for creating many PaymentBase entities in bulk.
type PaymentBaseCreateBulk struct {
	config
	err      error
	builders []*PaymentBaseCreate
	conflict []sql.ConflictOption
}

// Save creates the PaymentBase entities in the database.
func (pbcb *PaymentBaseCreateBulk) Save(ctx context.Context) ([]*PaymentBase, error) {
	if pbcb.err != nil {
		return nil, pbcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pbcb.builders))
	nodes := make([]*PaymentBase, len(pbcb.builders))
	mutators := make([]Mutator, len(pbcb.builders))
	for i := range pbcb.builders {
		func(i int, root context.Context) {
			builder := pbcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PaymentBaseMutation)
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
					_, err = mutators[i+1].Mutate(root, pbcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pbcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pbcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pbcb *PaymentBaseCreateBulk) SaveX(ctx context.Context) []*PaymentBase {
	v, err := pbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pbcb *PaymentBaseCreateBulk) Exec(ctx context.Context) error {
	_, err := pbcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pbcb *PaymentBaseCreateBulk) ExecX(ctx context.Context) {
	if err := pbcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.PaymentBase.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PaymentBaseUpsert) {
//			SetEntID(v+v).
//		}).
//		Exec(ctx)
func (pbcb *PaymentBaseCreateBulk) OnConflict(opts ...sql.ConflictOption) *PaymentBaseUpsertBulk {
	pbcb.conflict = opts
	return &PaymentBaseUpsertBulk{
		create: pbcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.PaymentBase.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pbcb *PaymentBaseCreateBulk) OnConflictColumns(columns ...string) *PaymentBaseUpsertBulk {
	pbcb.conflict = append(pbcb.conflict, sql.ConflictColumns(columns...))
	return &PaymentBaseUpsertBulk{
		create: pbcb,
	}
}

// PaymentBaseUpsertBulk is the builder for "upsert"-ing
// a bulk of PaymentBase nodes.
type PaymentBaseUpsertBulk struct {
	create *PaymentBaseCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.PaymentBase.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(paymentbase.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PaymentBaseUpsertBulk) UpdateNewValues() *PaymentBaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(paymentbase.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.PaymentBase.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PaymentBaseUpsertBulk) Ignore() *PaymentBaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PaymentBaseUpsertBulk) DoNothing() *PaymentBaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PaymentBaseCreateBulk.OnConflict
// documentation for more info.
func (u *PaymentBaseUpsertBulk) Update(set func(*PaymentBaseUpsert)) *PaymentBaseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PaymentBaseUpsert{UpdateSet: update})
	}))
	return u
}

// SetEntID sets the "ent_id" field.
func (u *PaymentBaseUpsertBulk) SetEntID(v uuid.UUID) *PaymentBaseUpsertBulk {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *PaymentBaseUpsertBulk) UpdateEntID() *PaymentBaseUpsertBulk {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.UpdateEntID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *PaymentBaseUpsertBulk) SetCreatedAt(v uint32) *PaymentBaseUpsertBulk {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *PaymentBaseUpsertBulk) AddCreatedAt(v uint32) *PaymentBaseUpsertBulk {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PaymentBaseUpsertBulk) UpdateCreatedAt() *PaymentBaseUpsertBulk {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PaymentBaseUpsertBulk) SetUpdatedAt(v uint32) *PaymentBaseUpsertBulk {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *PaymentBaseUpsertBulk) AddUpdatedAt(v uint32) *PaymentBaseUpsertBulk {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PaymentBaseUpsertBulk) UpdateUpdatedAt() *PaymentBaseUpsertBulk {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *PaymentBaseUpsertBulk) SetDeletedAt(v uint32) *PaymentBaseUpsertBulk {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *PaymentBaseUpsertBulk) AddDeletedAt(v uint32) *PaymentBaseUpsertBulk {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *PaymentBaseUpsertBulk) UpdateDeletedAt() *PaymentBaseUpsertBulk {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetOrderID sets the "order_id" field.
func (u *PaymentBaseUpsertBulk) SetOrderID(v uuid.UUID) *PaymentBaseUpsertBulk {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.SetOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *PaymentBaseUpsertBulk) UpdateOrderID() *PaymentBaseUpsertBulk {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.UpdateOrderID()
	})
}

// ClearOrderID clears the value of the "order_id" field.
func (u *PaymentBaseUpsertBulk) ClearOrderID() *PaymentBaseUpsertBulk {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.ClearOrderID()
	})
}

// SetObseleteState sets the "obselete_state" field.
func (u *PaymentBaseUpsertBulk) SetObseleteState(v string) *PaymentBaseUpsertBulk {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.SetObseleteState(v)
	})
}

// UpdateObseleteState sets the "obselete_state" field to the value that was provided on create.
func (u *PaymentBaseUpsertBulk) UpdateObseleteState() *PaymentBaseUpsertBulk {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.UpdateObseleteState()
	})
}

// ClearObseleteState clears the value of the "obselete_state" field.
func (u *PaymentBaseUpsertBulk) ClearObseleteState() *PaymentBaseUpsertBulk {
	return u.Update(func(s *PaymentBaseUpsert) {
		s.ClearObseleteState()
	})
}

// Exec executes the query.
func (u *PaymentBaseUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("generated: OnConflict was set for builder %d. Set it on the PaymentBaseCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for PaymentBaseCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PaymentBaseUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
