// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/orderpaymentstatement"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// OrderPaymentStatementCreate is the builder for creating a OrderPaymentStatement entity.
type OrderPaymentStatementCreate struct {
	config
	mutation *OrderPaymentStatementMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (opsc *OrderPaymentStatementCreate) SetCreatedAt(u uint32) *OrderPaymentStatementCreate {
	opsc.mutation.SetCreatedAt(u)
	return opsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (opsc *OrderPaymentStatementCreate) SetNillableCreatedAt(u *uint32) *OrderPaymentStatementCreate {
	if u != nil {
		opsc.SetCreatedAt(*u)
	}
	return opsc
}

// SetUpdatedAt sets the "updated_at" field.
func (opsc *OrderPaymentStatementCreate) SetUpdatedAt(u uint32) *OrderPaymentStatementCreate {
	opsc.mutation.SetUpdatedAt(u)
	return opsc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (opsc *OrderPaymentStatementCreate) SetNillableUpdatedAt(u *uint32) *OrderPaymentStatementCreate {
	if u != nil {
		opsc.SetUpdatedAt(*u)
	}
	return opsc
}

// SetDeletedAt sets the "deleted_at" field.
func (opsc *OrderPaymentStatementCreate) SetDeletedAt(u uint32) *OrderPaymentStatementCreate {
	opsc.mutation.SetDeletedAt(u)
	return opsc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (opsc *OrderPaymentStatementCreate) SetNillableDeletedAt(u *uint32) *OrderPaymentStatementCreate {
	if u != nil {
		opsc.SetDeletedAt(*u)
	}
	return opsc
}

// SetEntID sets the "ent_id" field.
func (opsc *OrderPaymentStatementCreate) SetEntID(u uuid.UUID) *OrderPaymentStatementCreate {
	opsc.mutation.SetEntID(u)
	return opsc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (opsc *OrderPaymentStatementCreate) SetNillableEntID(u *uuid.UUID) *OrderPaymentStatementCreate {
	if u != nil {
		opsc.SetEntID(*u)
	}
	return opsc
}

// SetStatementID sets the "statement_id" field.
func (opsc *OrderPaymentStatementCreate) SetStatementID(u uuid.UUID) *OrderPaymentStatementCreate {
	opsc.mutation.SetStatementID(u)
	return opsc
}

// SetNillableStatementID sets the "statement_id" field if the given value is not nil.
func (opsc *OrderPaymentStatementCreate) SetNillableStatementID(u *uuid.UUID) *OrderPaymentStatementCreate {
	if u != nil {
		opsc.SetStatementID(*u)
	}
	return opsc
}

// SetPaymentCoinTypeID sets the "payment_coin_type_id" field.
func (opsc *OrderPaymentStatementCreate) SetPaymentCoinTypeID(u uuid.UUID) *OrderPaymentStatementCreate {
	opsc.mutation.SetPaymentCoinTypeID(u)
	return opsc
}

// SetNillablePaymentCoinTypeID sets the "payment_coin_type_id" field if the given value is not nil.
func (opsc *OrderPaymentStatementCreate) SetNillablePaymentCoinTypeID(u *uuid.UUID) *OrderPaymentStatementCreate {
	if u != nil {
		opsc.SetPaymentCoinTypeID(*u)
	}
	return opsc
}

// SetAmount sets the "amount" field.
func (opsc *OrderPaymentStatementCreate) SetAmount(d decimal.Decimal) *OrderPaymentStatementCreate {
	opsc.mutation.SetAmount(d)
	return opsc
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (opsc *OrderPaymentStatementCreate) SetNillableAmount(d *decimal.Decimal) *OrderPaymentStatementCreate {
	if d != nil {
		opsc.SetAmount(*d)
	}
	return opsc
}

// SetCommissionAmount sets the "commission_amount" field.
func (opsc *OrderPaymentStatementCreate) SetCommissionAmount(d decimal.Decimal) *OrderPaymentStatementCreate {
	opsc.mutation.SetCommissionAmount(d)
	return opsc
}

// SetNillableCommissionAmount sets the "commission_amount" field if the given value is not nil.
func (opsc *OrderPaymentStatementCreate) SetNillableCommissionAmount(d *decimal.Decimal) *OrderPaymentStatementCreate {
	if d != nil {
		opsc.SetCommissionAmount(*d)
	}
	return opsc
}

// SetID sets the "id" field.
func (opsc *OrderPaymentStatementCreate) SetID(u uint32) *OrderPaymentStatementCreate {
	opsc.mutation.SetID(u)
	return opsc
}

// Mutation returns the OrderPaymentStatementMutation object of the builder.
func (opsc *OrderPaymentStatementCreate) Mutation() *OrderPaymentStatementMutation {
	return opsc.mutation
}

// Save creates the OrderPaymentStatement in the database.
func (opsc *OrderPaymentStatementCreate) Save(ctx context.Context) (*OrderPaymentStatement, error) {
	opsc.defaults()
	return withHooks(ctx, opsc.sqlSave, opsc.mutation, opsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (opsc *OrderPaymentStatementCreate) SaveX(ctx context.Context) *OrderPaymentStatement {
	v, err := opsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (opsc *OrderPaymentStatementCreate) Exec(ctx context.Context) error {
	_, err := opsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opsc *OrderPaymentStatementCreate) ExecX(ctx context.Context) {
	if err := opsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (opsc *OrderPaymentStatementCreate) defaults() {
	if _, ok := opsc.mutation.CreatedAt(); !ok {
		v := orderpaymentstatement.DefaultCreatedAt()
		opsc.mutation.SetCreatedAt(v)
	}
	if _, ok := opsc.mutation.UpdatedAt(); !ok {
		v := orderpaymentstatement.DefaultUpdatedAt()
		opsc.mutation.SetUpdatedAt(v)
	}
	if _, ok := opsc.mutation.DeletedAt(); !ok {
		v := orderpaymentstatement.DefaultDeletedAt()
		opsc.mutation.SetDeletedAt(v)
	}
	if _, ok := opsc.mutation.EntID(); !ok {
		v := orderpaymentstatement.DefaultEntID()
		opsc.mutation.SetEntID(v)
	}
	if _, ok := opsc.mutation.StatementID(); !ok {
		v := orderpaymentstatement.DefaultStatementID()
		opsc.mutation.SetStatementID(v)
	}
	if _, ok := opsc.mutation.PaymentCoinTypeID(); !ok {
		v := orderpaymentstatement.DefaultPaymentCoinTypeID()
		opsc.mutation.SetPaymentCoinTypeID(v)
	}
	if _, ok := opsc.mutation.Amount(); !ok {
		v := orderpaymentstatement.DefaultAmount
		opsc.mutation.SetAmount(v)
	}
	if _, ok := opsc.mutation.CommissionAmount(); !ok {
		v := orderpaymentstatement.DefaultCommissionAmount
		opsc.mutation.SetCommissionAmount(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (opsc *OrderPaymentStatementCreate) check() error {
	if _, ok := opsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "OrderPaymentStatement.created_at"`)}
	}
	if _, ok := opsc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "OrderPaymentStatement.updated_at"`)}
	}
	if _, ok := opsc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`generated: missing required field "OrderPaymentStatement.deleted_at"`)}
	}
	if _, ok := opsc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`generated: missing required field "OrderPaymentStatement.ent_id"`)}
	}
	return nil
}

func (opsc *OrderPaymentStatementCreate) sqlSave(ctx context.Context) (*OrderPaymentStatement, error) {
	if err := opsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := opsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, opsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	opsc.mutation.id = &_node.ID
	opsc.mutation.done = true
	return _node, nil
}

func (opsc *OrderPaymentStatementCreate) createSpec() (*OrderPaymentStatement, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderPaymentStatement{config: opsc.config}
		_spec = sqlgraph.NewCreateSpec(orderpaymentstatement.Table, sqlgraph.NewFieldSpec(orderpaymentstatement.FieldID, field.TypeUint32))
	)
	_spec.OnConflict = opsc.conflict
	if id, ok := opsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := opsc.mutation.CreatedAt(); ok {
		_spec.SetField(orderpaymentstatement.FieldCreatedAt, field.TypeUint32, value)
		_node.CreatedAt = value
	}
	if value, ok := opsc.mutation.UpdatedAt(); ok {
		_spec.SetField(orderpaymentstatement.FieldUpdatedAt, field.TypeUint32, value)
		_node.UpdatedAt = value
	}
	if value, ok := opsc.mutation.DeletedAt(); ok {
		_spec.SetField(orderpaymentstatement.FieldDeletedAt, field.TypeUint32, value)
		_node.DeletedAt = value
	}
	if value, ok := opsc.mutation.EntID(); ok {
		_spec.SetField(orderpaymentstatement.FieldEntID, field.TypeUUID, value)
		_node.EntID = value
	}
	if value, ok := opsc.mutation.StatementID(); ok {
		_spec.SetField(orderpaymentstatement.FieldStatementID, field.TypeUUID, value)
		_node.StatementID = value
	}
	if value, ok := opsc.mutation.PaymentCoinTypeID(); ok {
		_spec.SetField(orderpaymentstatement.FieldPaymentCoinTypeID, field.TypeUUID, value)
		_node.PaymentCoinTypeID = value
	}
	if value, ok := opsc.mutation.Amount(); ok {
		_spec.SetField(orderpaymentstatement.FieldAmount, field.TypeOther, value)
		_node.Amount = value
	}
	if value, ok := opsc.mutation.CommissionAmount(); ok {
		_spec.SetField(orderpaymentstatement.FieldCommissionAmount, field.TypeOther, value)
		_node.CommissionAmount = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OrderPaymentStatement.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrderPaymentStatementUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (opsc *OrderPaymentStatementCreate) OnConflict(opts ...sql.ConflictOption) *OrderPaymentStatementUpsertOne {
	opsc.conflict = opts
	return &OrderPaymentStatementUpsertOne{
		create: opsc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OrderPaymentStatement.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (opsc *OrderPaymentStatementCreate) OnConflictColumns(columns ...string) *OrderPaymentStatementUpsertOne {
	opsc.conflict = append(opsc.conflict, sql.ConflictColumns(columns...))
	return &OrderPaymentStatementUpsertOne{
		create: opsc,
	}
}

type (
	// OrderPaymentStatementUpsertOne is the builder for "upsert"-ing
	//  one OrderPaymentStatement node.
	OrderPaymentStatementUpsertOne struct {
		create *OrderPaymentStatementCreate
	}

	// OrderPaymentStatementUpsert is the "OnConflict" setter.
	OrderPaymentStatementUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *OrderPaymentStatementUpsert) SetCreatedAt(v uint32) *OrderPaymentStatementUpsert {
	u.Set(orderpaymentstatement.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *OrderPaymentStatementUpsert) UpdateCreatedAt() *OrderPaymentStatementUpsert {
	u.SetExcluded(orderpaymentstatement.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *OrderPaymentStatementUpsert) AddCreatedAt(v uint32) *OrderPaymentStatementUpsert {
	u.Add(orderpaymentstatement.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrderPaymentStatementUpsert) SetUpdatedAt(v uint32) *OrderPaymentStatementUpsert {
	u.Set(orderpaymentstatement.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrderPaymentStatementUpsert) UpdateUpdatedAt() *OrderPaymentStatementUpsert {
	u.SetExcluded(orderpaymentstatement.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *OrderPaymentStatementUpsert) AddUpdatedAt(v uint32) *OrderPaymentStatementUpsert {
	u.Add(orderpaymentstatement.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *OrderPaymentStatementUpsert) SetDeletedAt(v uint32) *OrderPaymentStatementUpsert {
	u.Set(orderpaymentstatement.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *OrderPaymentStatementUpsert) UpdateDeletedAt() *OrderPaymentStatementUpsert {
	u.SetExcluded(orderpaymentstatement.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *OrderPaymentStatementUpsert) AddDeletedAt(v uint32) *OrderPaymentStatementUpsert {
	u.Add(orderpaymentstatement.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *OrderPaymentStatementUpsert) SetEntID(v uuid.UUID) *OrderPaymentStatementUpsert {
	u.Set(orderpaymentstatement.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *OrderPaymentStatementUpsert) UpdateEntID() *OrderPaymentStatementUpsert {
	u.SetExcluded(orderpaymentstatement.FieldEntID)
	return u
}

// SetStatementID sets the "statement_id" field.
func (u *OrderPaymentStatementUpsert) SetStatementID(v uuid.UUID) *OrderPaymentStatementUpsert {
	u.Set(orderpaymentstatement.FieldStatementID, v)
	return u
}

// UpdateStatementID sets the "statement_id" field to the value that was provided on create.
func (u *OrderPaymentStatementUpsert) UpdateStatementID() *OrderPaymentStatementUpsert {
	u.SetExcluded(orderpaymentstatement.FieldStatementID)
	return u
}

// ClearStatementID clears the value of the "statement_id" field.
func (u *OrderPaymentStatementUpsert) ClearStatementID() *OrderPaymentStatementUpsert {
	u.SetNull(orderpaymentstatement.FieldStatementID)
	return u
}

// SetPaymentCoinTypeID sets the "payment_coin_type_id" field.
func (u *OrderPaymentStatementUpsert) SetPaymentCoinTypeID(v uuid.UUID) *OrderPaymentStatementUpsert {
	u.Set(orderpaymentstatement.FieldPaymentCoinTypeID, v)
	return u
}

// UpdatePaymentCoinTypeID sets the "payment_coin_type_id" field to the value that was provided on create.
func (u *OrderPaymentStatementUpsert) UpdatePaymentCoinTypeID() *OrderPaymentStatementUpsert {
	u.SetExcluded(orderpaymentstatement.FieldPaymentCoinTypeID)
	return u
}

// ClearPaymentCoinTypeID clears the value of the "payment_coin_type_id" field.
func (u *OrderPaymentStatementUpsert) ClearPaymentCoinTypeID() *OrderPaymentStatementUpsert {
	u.SetNull(orderpaymentstatement.FieldPaymentCoinTypeID)
	return u
}

// SetAmount sets the "amount" field.
func (u *OrderPaymentStatementUpsert) SetAmount(v decimal.Decimal) *OrderPaymentStatementUpsert {
	u.Set(orderpaymentstatement.FieldAmount, v)
	return u
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *OrderPaymentStatementUpsert) UpdateAmount() *OrderPaymentStatementUpsert {
	u.SetExcluded(orderpaymentstatement.FieldAmount)
	return u
}

// ClearAmount clears the value of the "amount" field.
func (u *OrderPaymentStatementUpsert) ClearAmount() *OrderPaymentStatementUpsert {
	u.SetNull(orderpaymentstatement.FieldAmount)
	return u
}

// SetCommissionAmount sets the "commission_amount" field.
func (u *OrderPaymentStatementUpsert) SetCommissionAmount(v decimal.Decimal) *OrderPaymentStatementUpsert {
	u.Set(orderpaymentstatement.FieldCommissionAmount, v)
	return u
}

// UpdateCommissionAmount sets the "commission_amount" field to the value that was provided on create.
func (u *OrderPaymentStatementUpsert) UpdateCommissionAmount() *OrderPaymentStatementUpsert {
	u.SetExcluded(orderpaymentstatement.FieldCommissionAmount)
	return u
}

// ClearCommissionAmount clears the value of the "commission_amount" field.
func (u *OrderPaymentStatementUpsert) ClearCommissionAmount() *OrderPaymentStatementUpsert {
	u.SetNull(orderpaymentstatement.FieldCommissionAmount)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.OrderPaymentStatement.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(orderpaymentstatement.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *OrderPaymentStatementUpsertOne) UpdateNewValues() *OrderPaymentStatementUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(orderpaymentstatement.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OrderPaymentStatement.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *OrderPaymentStatementUpsertOne) Ignore() *OrderPaymentStatementUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrderPaymentStatementUpsertOne) DoNothing() *OrderPaymentStatementUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrderPaymentStatementCreate.OnConflict
// documentation for more info.
func (u *OrderPaymentStatementUpsertOne) Update(set func(*OrderPaymentStatementUpsert)) *OrderPaymentStatementUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrderPaymentStatementUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *OrderPaymentStatementUpsertOne) SetCreatedAt(v uint32) *OrderPaymentStatementUpsertOne {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *OrderPaymentStatementUpsertOne) AddCreatedAt(v uint32) *OrderPaymentStatementUpsertOne {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *OrderPaymentStatementUpsertOne) UpdateCreatedAt() *OrderPaymentStatementUpsertOne {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrderPaymentStatementUpsertOne) SetUpdatedAt(v uint32) *OrderPaymentStatementUpsertOne {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *OrderPaymentStatementUpsertOne) AddUpdatedAt(v uint32) *OrderPaymentStatementUpsertOne {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrderPaymentStatementUpsertOne) UpdateUpdatedAt() *OrderPaymentStatementUpsertOne {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *OrderPaymentStatementUpsertOne) SetDeletedAt(v uint32) *OrderPaymentStatementUpsertOne {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *OrderPaymentStatementUpsertOne) AddDeletedAt(v uint32) *OrderPaymentStatementUpsertOne {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *OrderPaymentStatementUpsertOne) UpdateDeletedAt() *OrderPaymentStatementUpsertOne {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *OrderPaymentStatementUpsertOne) SetEntID(v uuid.UUID) *OrderPaymentStatementUpsertOne {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *OrderPaymentStatementUpsertOne) UpdateEntID() *OrderPaymentStatementUpsertOne {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.UpdateEntID()
	})
}

// SetStatementID sets the "statement_id" field.
func (u *OrderPaymentStatementUpsertOne) SetStatementID(v uuid.UUID) *OrderPaymentStatementUpsertOne {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.SetStatementID(v)
	})
}

// UpdateStatementID sets the "statement_id" field to the value that was provided on create.
func (u *OrderPaymentStatementUpsertOne) UpdateStatementID() *OrderPaymentStatementUpsertOne {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.UpdateStatementID()
	})
}

// ClearStatementID clears the value of the "statement_id" field.
func (u *OrderPaymentStatementUpsertOne) ClearStatementID() *OrderPaymentStatementUpsertOne {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.ClearStatementID()
	})
}

// SetPaymentCoinTypeID sets the "payment_coin_type_id" field.
func (u *OrderPaymentStatementUpsertOne) SetPaymentCoinTypeID(v uuid.UUID) *OrderPaymentStatementUpsertOne {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.SetPaymentCoinTypeID(v)
	})
}

// UpdatePaymentCoinTypeID sets the "payment_coin_type_id" field to the value that was provided on create.
func (u *OrderPaymentStatementUpsertOne) UpdatePaymentCoinTypeID() *OrderPaymentStatementUpsertOne {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.UpdatePaymentCoinTypeID()
	})
}

// ClearPaymentCoinTypeID clears the value of the "payment_coin_type_id" field.
func (u *OrderPaymentStatementUpsertOne) ClearPaymentCoinTypeID() *OrderPaymentStatementUpsertOne {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.ClearPaymentCoinTypeID()
	})
}

// SetAmount sets the "amount" field.
func (u *OrderPaymentStatementUpsertOne) SetAmount(v decimal.Decimal) *OrderPaymentStatementUpsertOne {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.SetAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *OrderPaymentStatementUpsertOne) UpdateAmount() *OrderPaymentStatementUpsertOne {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.UpdateAmount()
	})
}

// ClearAmount clears the value of the "amount" field.
func (u *OrderPaymentStatementUpsertOne) ClearAmount() *OrderPaymentStatementUpsertOne {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.ClearAmount()
	})
}

// SetCommissionAmount sets the "commission_amount" field.
func (u *OrderPaymentStatementUpsertOne) SetCommissionAmount(v decimal.Decimal) *OrderPaymentStatementUpsertOne {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.SetCommissionAmount(v)
	})
}

// UpdateCommissionAmount sets the "commission_amount" field to the value that was provided on create.
func (u *OrderPaymentStatementUpsertOne) UpdateCommissionAmount() *OrderPaymentStatementUpsertOne {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.UpdateCommissionAmount()
	})
}

// ClearCommissionAmount clears the value of the "commission_amount" field.
func (u *OrderPaymentStatementUpsertOne) ClearCommissionAmount() *OrderPaymentStatementUpsertOne {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.ClearCommissionAmount()
	})
}

// Exec executes the query.
func (u *OrderPaymentStatementUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for OrderPaymentStatementCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrderPaymentStatementUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *OrderPaymentStatementUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *OrderPaymentStatementUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// OrderPaymentStatementCreateBulk is the builder for creating many OrderPaymentStatement entities in bulk.
type OrderPaymentStatementCreateBulk struct {
	config
	err      error
	builders []*OrderPaymentStatementCreate
	conflict []sql.ConflictOption
}

// Save creates the OrderPaymentStatement entities in the database.
func (opscb *OrderPaymentStatementCreateBulk) Save(ctx context.Context) ([]*OrderPaymentStatement, error) {
	if opscb.err != nil {
		return nil, opscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(opscb.builders))
	nodes := make([]*OrderPaymentStatement, len(opscb.builders))
	mutators := make([]Mutator, len(opscb.builders))
	for i := range opscb.builders {
		func(i int, root context.Context) {
			builder := opscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderPaymentStatementMutation)
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
					_, err = mutators[i+1].Mutate(root, opscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = opscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, opscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, opscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (opscb *OrderPaymentStatementCreateBulk) SaveX(ctx context.Context) []*OrderPaymentStatement {
	v, err := opscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (opscb *OrderPaymentStatementCreateBulk) Exec(ctx context.Context) error {
	_, err := opscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opscb *OrderPaymentStatementCreateBulk) ExecX(ctx context.Context) {
	if err := opscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OrderPaymentStatement.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrderPaymentStatementUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (opscb *OrderPaymentStatementCreateBulk) OnConflict(opts ...sql.ConflictOption) *OrderPaymentStatementUpsertBulk {
	opscb.conflict = opts
	return &OrderPaymentStatementUpsertBulk{
		create: opscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OrderPaymentStatement.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (opscb *OrderPaymentStatementCreateBulk) OnConflictColumns(columns ...string) *OrderPaymentStatementUpsertBulk {
	opscb.conflict = append(opscb.conflict, sql.ConflictColumns(columns...))
	return &OrderPaymentStatementUpsertBulk{
		create: opscb,
	}
}

// OrderPaymentStatementUpsertBulk is the builder for "upsert"-ing
// a bulk of OrderPaymentStatement nodes.
type OrderPaymentStatementUpsertBulk struct {
	create *OrderPaymentStatementCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.OrderPaymentStatement.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(orderpaymentstatement.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *OrderPaymentStatementUpsertBulk) UpdateNewValues() *OrderPaymentStatementUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(orderpaymentstatement.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OrderPaymentStatement.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *OrderPaymentStatementUpsertBulk) Ignore() *OrderPaymentStatementUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrderPaymentStatementUpsertBulk) DoNothing() *OrderPaymentStatementUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrderPaymentStatementCreateBulk.OnConflict
// documentation for more info.
func (u *OrderPaymentStatementUpsertBulk) Update(set func(*OrderPaymentStatementUpsert)) *OrderPaymentStatementUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrderPaymentStatementUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *OrderPaymentStatementUpsertBulk) SetCreatedAt(v uint32) *OrderPaymentStatementUpsertBulk {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *OrderPaymentStatementUpsertBulk) AddCreatedAt(v uint32) *OrderPaymentStatementUpsertBulk {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *OrderPaymentStatementUpsertBulk) UpdateCreatedAt() *OrderPaymentStatementUpsertBulk {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrderPaymentStatementUpsertBulk) SetUpdatedAt(v uint32) *OrderPaymentStatementUpsertBulk {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *OrderPaymentStatementUpsertBulk) AddUpdatedAt(v uint32) *OrderPaymentStatementUpsertBulk {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrderPaymentStatementUpsertBulk) UpdateUpdatedAt() *OrderPaymentStatementUpsertBulk {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *OrderPaymentStatementUpsertBulk) SetDeletedAt(v uint32) *OrderPaymentStatementUpsertBulk {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *OrderPaymentStatementUpsertBulk) AddDeletedAt(v uint32) *OrderPaymentStatementUpsertBulk {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *OrderPaymentStatementUpsertBulk) UpdateDeletedAt() *OrderPaymentStatementUpsertBulk {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *OrderPaymentStatementUpsertBulk) SetEntID(v uuid.UUID) *OrderPaymentStatementUpsertBulk {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *OrderPaymentStatementUpsertBulk) UpdateEntID() *OrderPaymentStatementUpsertBulk {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.UpdateEntID()
	})
}

// SetStatementID sets the "statement_id" field.
func (u *OrderPaymentStatementUpsertBulk) SetStatementID(v uuid.UUID) *OrderPaymentStatementUpsertBulk {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.SetStatementID(v)
	})
}

// UpdateStatementID sets the "statement_id" field to the value that was provided on create.
func (u *OrderPaymentStatementUpsertBulk) UpdateStatementID() *OrderPaymentStatementUpsertBulk {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.UpdateStatementID()
	})
}

// ClearStatementID clears the value of the "statement_id" field.
func (u *OrderPaymentStatementUpsertBulk) ClearStatementID() *OrderPaymentStatementUpsertBulk {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.ClearStatementID()
	})
}

// SetPaymentCoinTypeID sets the "payment_coin_type_id" field.
func (u *OrderPaymentStatementUpsertBulk) SetPaymentCoinTypeID(v uuid.UUID) *OrderPaymentStatementUpsertBulk {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.SetPaymentCoinTypeID(v)
	})
}

// UpdatePaymentCoinTypeID sets the "payment_coin_type_id" field to the value that was provided on create.
func (u *OrderPaymentStatementUpsertBulk) UpdatePaymentCoinTypeID() *OrderPaymentStatementUpsertBulk {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.UpdatePaymentCoinTypeID()
	})
}

// ClearPaymentCoinTypeID clears the value of the "payment_coin_type_id" field.
func (u *OrderPaymentStatementUpsertBulk) ClearPaymentCoinTypeID() *OrderPaymentStatementUpsertBulk {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.ClearPaymentCoinTypeID()
	})
}

// SetAmount sets the "amount" field.
func (u *OrderPaymentStatementUpsertBulk) SetAmount(v decimal.Decimal) *OrderPaymentStatementUpsertBulk {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.SetAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *OrderPaymentStatementUpsertBulk) UpdateAmount() *OrderPaymentStatementUpsertBulk {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.UpdateAmount()
	})
}

// ClearAmount clears the value of the "amount" field.
func (u *OrderPaymentStatementUpsertBulk) ClearAmount() *OrderPaymentStatementUpsertBulk {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.ClearAmount()
	})
}

// SetCommissionAmount sets the "commission_amount" field.
func (u *OrderPaymentStatementUpsertBulk) SetCommissionAmount(v decimal.Decimal) *OrderPaymentStatementUpsertBulk {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.SetCommissionAmount(v)
	})
}

// UpdateCommissionAmount sets the "commission_amount" field to the value that was provided on create.
func (u *OrderPaymentStatementUpsertBulk) UpdateCommissionAmount() *OrderPaymentStatementUpsertBulk {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.UpdateCommissionAmount()
	})
}

// ClearCommissionAmount clears the value of the "commission_amount" field.
func (u *OrderPaymentStatementUpsertBulk) ClearCommissionAmount() *OrderPaymentStatementUpsertBulk {
	return u.Update(func(s *OrderPaymentStatementUpsert) {
		s.ClearCommissionAmount()
	})
}

// Exec executes the query.
func (u *OrderPaymentStatementUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("generated: OnConflict was set for builder %d. Set it on the OrderPaymentStatementCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for OrderPaymentStatementCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrderPaymentStatementUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
