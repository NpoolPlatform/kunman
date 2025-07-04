// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/ledgerlock"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// LedgerLockCreate is the builder for creating a LedgerLock entity.
type LedgerLockCreate struct {
	config
	mutation *LedgerLockMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (llc *LedgerLockCreate) SetCreatedAt(u uint32) *LedgerLockCreate {
	llc.mutation.SetCreatedAt(u)
	return llc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (llc *LedgerLockCreate) SetNillableCreatedAt(u *uint32) *LedgerLockCreate {
	if u != nil {
		llc.SetCreatedAt(*u)
	}
	return llc
}

// SetUpdatedAt sets the "updated_at" field.
func (llc *LedgerLockCreate) SetUpdatedAt(u uint32) *LedgerLockCreate {
	llc.mutation.SetUpdatedAt(u)
	return llc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (llc *LedgerLockCreate) SetNillableUpdatedAt(u *uint32) *LedgerLockCreate {
	if u != nil {
		llc.SetUpdatedAt(*u)
	}
	return llc
}

// SetDeletedAt sets the "deleted_at" field.
func (llc *LedgerLockCreate) SetDeletedAt(u uint32) *LedgerLockCreate {
	llc.mutation.SetDeletedAt(u)
	return llc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (llc *LedgerLockCreate) SetNillableDeletedAt(u *uint32) *LedgerLockCreate {
	if u != nil {
		llc.SetDeletedAt(*u)
	}
	return llc
}

// SetEntID sets the "ent_id" field.
func (llc *LedgerLockCreate) SetEntID(u uuid.UUID) *LedgerLockCreate {
	llc.mutation.SetEntID(u)
	return llc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (llc *LedgerLockCreate) SetNillableEntID(u *uuid.UUID) *LedgerLockCreate {
	if u != nil {
		llc.SetEntID(*u)
	}
	return llc
}

// SetLedgerID sets the "ledger_id" field.
func (llc *LedgerLockCreate) SetLedgerID(u uuid.UUID) *LedgerLockCreate {
	llc.mutation.SetLedgerID(u)
	return llc
}

// SetNillableLedgerID sets the "ledger_id" field if the given value is not nil.
func (llc *LedgerLockCreate) SetNillableLedgerID(u *uuid.UUID) *LedgerLockCreate {
	if u != nil {
		llc.SetLedgerID(*u)
	}
	return llc
}

// SetStatementID sets the "statement_id" field.
func (llc *LedgerLockCreate) SetStatementID(u uuid.UUID) *LedgerLockCreate {
	llc.mutation.SetStatementID(u)
	return llc
}

// SetNillableStatementID sets the "statement_id" field if the given value is not nil.
func (llc *LedgerLockCreate) SetNillableStatementID(u *uuid.UUID) *LedgerLockCreate {
	if u != nil {
		llc.SetStatementID(*u)
	}
	return llc
}

// SetAmount sets the "amount" field.
func (llc *LedgerLockCreate) SetAmount(d decimal.Decimal) *LedgerLockCreate {
	llc.mutation.SetAmount(d)
	return llc
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (llc *LedgerLockCreate) SetNillableAmount(d *decimal.Decimal) *LedgerLockCreate {
	if d != nil {
		llc.SetAmount(*d)
	}
	return llc
}

// SetLockState sets the "lock_state" field.
func (llc *LedgerLockCreate) SetLockState(s string) *LedgerLockCreate {
	llc.mutation.SetLockState(s)
	return llc
}

// SetNillableLockState sets the "lock_state" field if the given value is not nil.
func (llc *LedgerLockCreate) SetNillableLockState(s *string) *LedgerLockCreate {
	if s != nil {
		llc.SetLockState(*s)
	}
	return llc
}

// SetExLockID sets the "ex_lock_id" field.
func (llc *LedgerLockCreate) SetExLockID(u uuid.UUID) *LedgerLockCreate {
	llc.mutation.SetExLockID(u)
	return llc
}

// SetNillableExLockID sets the "ex_lock_id" field if the given value is not nil.
func (llc *LedgerLockCreate) SetNillableExLockID(u *uuid.UUID) *LedgerLockCreate {
	if u != nil {
		llc.SetExLockID(*u)
	}
	return llc
}

// SetID sets the "id" field.
func (llc *LedgerLockCreate) SetID(u uint32) *LedgerLockCreate {
	llc.mutation.SetID(u)
	return llc
}

// Mutation returns the LedgerLockMutation object of the builder.
func (llc *LedgerLockCreate) Mutation() *LedgerLockMutation {
	return llc.mutation
}

// Save creates the LedgerLock in the database.
func (llc *LedgerLockCreate) Save(ctx context.Context) (*LedgerLock, error) {
	llc.defaults()
	return withHooks(ctx, llc.sqlSave, llc.mutation, llc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (llc *LedgerLockCreate) SaveX(ctx context.Context) *LedgerLock {
	v, err := llc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (llc *LedgerLockCreate) Exec(ctx context.Context) error {
	_, err := llc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (llc *LedgerLockCreate) ExecX(ctx context.Context) {
	if err := llc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (llc *LedgerLockCreate) defaults() {
	if _, ok := llc.mutation.CreatedAt(); !ok {
		v := ledgerlock.DefaultCreatedAt()
		llc.mutation.SetCreatedAt(v)
	}
	if _, ok := llc.mutation.UpdatedAt(); !ok {
		v := ledgerlock.DefaultUpdatedAt()
		llc.mutation.SetUpdatedAt(v)
	}
	if _, ok := llc.mutation.DeletedAt(); !ok {
		v := ledgerlock.DefaultDeletedAt()
		llc.mutation.SetDeletedAt(v)
	}
	if _, ok := llc.mutation.EntID(); !ok {
		v := ledgerlock.DefaultEntID()
		llc.mutation.SetEntID(v)
	}
	if _, ok := llc.mutation.LedgerID(); !ok {
		v := ledgerlock.DefaultLedgerID()
		llc.mutation.SetLedgerID(v)
	}
	if _, ok := llc.mutation.StatementID(); !ok {
		v := ledgerlock.DefaultStatementID()
		llc.mutation.SetStatementID(v)
	}
	if _, ok := llc.mutation.LockState(); !ok {
		v := ledgerlock.DefaultLockState
		llc.mutation.SetLockState(v)
	}
	if _, ok := llc.mutation.ExLockID(); !ok {
		v := ledgerlock.DefaultExLockID()
		llc.mutation.SetExLockID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (llc *LedgerLockCreate) check() error {
	if _, ok := llc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "LedgerLock.created_at"`)}
	}
	if _, ok := llc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "LedgerLock.updated_at"`)}
	}
	if _, ok := llc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`generated: missing required field "LedgerLock.deleted_at"`)}
	}
	if _, ok := llc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`generated: missing required field "LedgerLock.ent_id"`)}
	}
	return nil
}

func (llc *LedgerLockCreate) sqlSave(ctx context.Context) (*LedgerLock, error) {
	if err := llc.check(); err != nil {
		return nil, err
	}
	_node, _spec := llc.createSpec()
	if err := sqlgraph.CreateNode(ctx, llc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	llc.mutation.id = &_node.ID
	llc.mutation.done = true
	return _node, nil
}

func (llc *LedgerLockCreate) createSpec() (*LedgerLock, *sqlgraph.CreateSpec) {
	var (
		_node = &LedgerLock{config: llc.config}
		_spec = sqlgraph.NewCreateSpec(ledgerlock.Table, sqlgraph.NewFieldSpec(ledgerlock.FieldID, field.TypeUint32))
	)
	_spec.OnConflict = llc.conflict
	if id, ok := llc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := llc.mutation.CreatedAt(); ok {
		_spec.SetField(ledgerlock.FieldCreatedAt, field.TypeUint32, value)
		_node.CreatedAt = value
	}
	if value, ok := llc.mutation.UpdatedAt(); ok {
		_spec.SetField(ledgerlock.FieldUpdatedAt, field.TypeUint32, value)
		_node.UpdatedAt = value
	}
	if value, ok := llc.mutation.DeletedAt(); ok {
		_spec.SetField(ledgerlock.FieldDeletedAt, field.TypeUint32, value)
		_node.DeletedAt = value
	}
	if value, ok := llc.mutation.EntID(); ok {
		_spec.SetField(ledgerlock.FieldEntID, field.TypeUUID, value)
		_node.EntID = value
	}
	if value, ok := llc.mutation.LedgerID(); ok {
		_spec.SetField(ledgerlock.FieldLedgerID, field.TypeUUID, value)
		_node.LedgerID = value
	}
	if value, ok := llc.mutation.StatementID(); ok {
		_spec.SetField(ledgerlock.FieldStatementID, field.TypeUUID, value)
		_node.StatementID = value
	}
	if value, ok := llc.mutation.Amount(); ok {
		_spec.SetField(ledgerlock.FieldAmount, field.TypeFloat64, value)
		_node.Amount = value
	}
	if value, ok := llc.mutation.LockState(); ok {
		_spec.SetField(ledgerlock.FieldLockState, field.TypeString, value)
		_node.LockState = value
	}
	if value, ok := llc.mutation.ExLockID(); ok {
		_spec.SetField(ledgerlock.FieldExLockID, field.TypeUUID, value)
		_node.ExLockID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.LedgerLock.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LedgerLockUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (llc *LedgerLockCreate) OnConflict(opts ...sql.ConflictOption) *LedgerLockUpsertOne {
	llc.conflict = opts
	return &LedgerLockUpsertOne{
		create: llc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.LedgerLock.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (llc *LedgerLockCreate) OnConflictColumns(columns ...string) *LedgerLockUpsertOne {
	llc.conflict = append(llc.conflict, sql.ConflictColumns(columns...))
	return &LedgerLockUpsertOne{
		create: llc,
	}
}

type (
	// LedgerLockUpsertOne is the builder for "upsert"-ing
	//  one LedgerLock node.
	LedgerLockUpsertOne struct {
		create *LedgerLockCreate
	}

	// LedgerLockUpsert is the "OnConflict" setter.
	LedgerLockUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *LedgerLockUpsert) SetCreatedAt(v uint32) *LedgerLockUpsert {
	u.Set(ledgerlock.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *LedgerLockUpsert) UpdateCreatedAt() *LedgerLockUpsert {
	u.SetExcluded(ledgerlock.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *LedgerLockUpsert) AddCreatedAt(v uint32) *LedgerLockUpsert {
	u.Add(ledgerlock.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *LedgerLockUpsert) SetUpdatedAt(v uint32) *LedgerLockUpsert {
	u.Set(ledgerlock.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *LedgerLockUpsert) UpdateUpdatedAt() *LedgerLockUpsert {
	u.SetExcluded(ledgerlock.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *LedgerLockUpsert) AddUpdatedAt(v uint32) *LedgerLockUpsert {
	u.Add(ledgerlock.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *LedgerLockUpsert) SetDeletedAt(v uint32) *LedgerLockUpsert {
	u.Set(ledgerlock.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *LedgerLockUpsert) UpdateDeletedAt() *LedgerLockUpsert {
	u.SetExcluded(ledgerlock.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *LedgerLockUpsert) AddDeletedAt(v uint32) *LedgerLockUpsert {
	u.Add(ledgerlock.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *LedgerLockUpsert) SetEntID(v uuid.UUID) *LedgerLockUpsert {
	u.Set(ledgerlock.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *LedgerLockUpsert) UpdateEntID() *LedgerLockUpsert {
	u.SetExcluded(ledgerlock.FieldEntID)
	return u
}

// SetLedgerID sets the "ledger_id" field.
func (u *LedgerLockUpsert) SetLedgerID(v uuid.UUID) *LedgerLockUpsert {
	u.Set(ledgerlock.FieldLedgerID, v)
	return u
}

// UpdateLedgerID sets the "ledger_id" field to the value that was provided on create.
func (u *LedgerLockUpsert) UpdateLedgerID() *LedgerLockUpsert {
	u.SetExcluded(ledgerlock.FieldLedgerID)
	return u
}

// ClearLedgerID clears the value of the "ledger_id" field.
func (u *LedgerLockUpsert) ClearLedgerID() *LedgerLockUpsert {
	u.SetNull(ledgerlock.FieldLedgerID)
	return u
}

// SetStatementID sets the "statement_id" field.
func (u *LedgerLockUpsert) SetStatementID(v uuid.UUID) *LedgerLockUpsert {
	u.Set(ledgerlock.FieldStatementID, v)
	return u
}

// UpdateStatementID sets the "statement_id" field to the value that was provided on create.
func (u *LedgerLockUpsert) UpdateStatementID() *LedgerLockUpsert {
	u.SetExcluded(ledgerlock.FieldStatementID)
	return u
}

// ClearStatementID clears the value of the "statement_id" field.
func (u *LedgerLockUpsert) ClearStatementID() *LedgerLockUpsert {
	u.SetNull(ledgerlock.FieldStatementID)
	return u
}

// SetAmount sets the "amount" field.
func (u *LedgerLockUpsert) SetAmount(v decimal.Decimal) *LedgerLockUpsert {
	u.Set(ledgerlock.FieldAmount, v)
	return u
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *LedgerLockUpsert) UpdateAmount() *LedgerLockUpsert {
	u.SetExcluded(ledgerlock.FieldAmount)
	return u
}

// AddAmount adds v to the "amount" field.
func (u *LedgerLockUpsert) AddAmount(v decimal.Decimal) *LedgerLockUpsert {
	u.Add(ledgerlock.FieldAmount, v)
	return u
}

// ClearAmount clears the value of the "amount" field.
func (u *LedgerLockUpsert) ClearAmount() *LedgerLockUpsert {
	u.SetNull(ledgerlock.FieldAmount)
	return u
}

// SetLockState sets the "lock_state" field.
func (u *LedgerLockUpsert) SetLockState(v string) *LedgerLockUpsert {
	u.Set(ledgerlock.FieldLockState, v)
	return u
}

// UpdateLockState sets the "lock_state" field to the value that was provided on create.
func (u *LedgerLockUpsert) UpdateLockState() *LedgerLockUpsert {
	u.SetExcluded(ledgerlock.FieldLockState)
	return u
}

// ClearLockState clears the value of the "lock_state" field.
func (u *LedgerLockUpsert) ClearLockState() *LedgerLockUpsert {
	u.SetNull(ledgerlock.FieldLockState)
	return u
}

// SetExLockID sets the "ex_lock_id" field.
func (u *LedgerLockUpsert) SetExLockID(v uuid.UUID) *LedgerLockUpsert {
	u.Set(ledgerlock.FieldExLockID, v)
	return u
}

// UpdateExLockID sets the "ex_lock_id" field to the value that was provided on create.
func (u *LedgerLockUpsert) UpdateExLockID() *LedgerLockUpsert {
	u.SetExcluded(ledgerlock.FieldExLockID)
	return u
}

// ClearExLockID clears the value of the "ex_lock_id" field.
func (u *LedgerLockUpsert) ClearExLockID() *LedgerLockUpsert {
	u.SetNull(ledgerlock.FieldExLockID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.LedgerLock.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(ledgerlock.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *LedgerLockUpsertOne) UpdateNewValues() *LedgerLockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(ledgerlock.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.LedgerLock.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *LedgerLockUpsertOne) Ignore() *LedgerLockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LedgerLockUpsertOne) DoNothing() *LedgerLockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LedgerLockCreate.OnConflict
// documentation for more info.
func (u *LedgerLockUpsertOne) Update(set func(*LedgerLockUpsert)) *LedgerLockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LedgerLockUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *LedgerLockUpsertOne) SetCreatedAt(v uint32) *LedgerLockUpsertOne {
	return u.Update(func(s *LedgerLockUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *LedgerLockUpsertOne) AddCreatedAt(v uint32) *LedgerLockUpsertOne {
	return u.Update(func(s *LedgerLockUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *LedgerLockUpsertOne) UpdateCreatedAt() *LedgerLockUpsertOne {
	return u.Update(func(s *LedgerLockUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *LedgerLockUpsertOne) SetUpdatedAt(v uint32) *LedgerLockUpsertOne {
	return u.Update(func(s *LedgerLockUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *LedgerLockUpsertOne) AddUpdatedAt(v uint32) *LedgerLockUpsertOne {
	return u.Update(func(s *LedgerLockUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *LedgerLockUpsertOne) UpdateUpdatedAt() *LedgerLockUpsertOne {
	return u.Update(func(s *LedgerLockUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *LedgerLockUpsertOne) SetDeletedAt(v uint32) *LedgerLockUpsertOne {
	return u.Update(func(s *LedgerLockUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *LedgerLockUpsertOne) AddDeletedAt(v uint32) *LedgerLockUpsertOne {
	return u.Update(func(s *LedgerLockUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *LedgerLockUpsertOne) UpdateDeletedAt() *LedgerLockUpsertOne {
	return u.Update(func(s *LedgerLockUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *LedgerLockUpsertOne) SetEntID(v uuid.UUID) *LedgerLockUpsertOne {
	return u.Update(func(s *LedgerLockUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *LedgerLockUpsertOne) UpdateEntID() *LedgerLockUpsertOne {
	return u.Update(func(s *LedgerLockUpsert) {
		s.UpdateEntID()
	})
}

// SetLedgerID sets the "ledger_id" field.
func (u *LedgerLockUpsertOne) SetLedgerID(v uuid.UUID) *LedgerLockUpsertOne {
	return u.Update(func(s *LedgerLockUpsert) {
		s.SetLedgerID(v)
	})
}

// UpdateLedgerID sets the "ledger_id" field to the value that was provided on create.
func (u *LedgerLockUpsertOne) UpdateLedgerID() *LedgerLockUpsertOne {
	return u.Update(func(s *LedgerLockUpsert) {
		s.UpdateLedgerID()
	})
}

// ClearLedgerID clears the value of the "ledger_id" field.
func (u *LedgerLockUpsertOne) ClearLedgerID() *LedgerLockUpsertOne {
	return u.Update(func(s *LedgerLockUpsert) {
		s.ClearLedgerID()
	})
}

// SetStatementID sets the "statement_id" field.
func (u *LedgerLockUpsertOne) SetStatementID(v uuid.UUID) *LedgerLockUpsertOne {
	return u.Update(func(s *LedgerLockUpsert) {
		s.SetStatementID(v)
	})
}

// UpdateStatementID sets the "statement_id" field to the value that was provided on create.
func (u *LedgerLockUpsertOne) UpdateStatementID() *LedgerLockUpsertOne {
	return u.Update(func(s *LedgerLockUpsert) {
		s.UpdateStatementID()
	})
}

// ClearStatementID clears the value of the "statement_id" field.
func (u *LedgerLockUpsertOne) ClearStatementID() *LedgerLockUpsertOne {
	return u.Update(func(s *LedgerLockUpsert) {
		s.ClearStatementID()
	})
}

// SetAmount sets the "amount" field.
func (u *LedgerLockUpsertOne) SetAmount(v decimal.Decimal) *LedgerLockUpsertOne {
	return u.Update(func(s *LedgerLockUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *LedgerLockUpsertOne) AddAmount(v decimal.Decimal) *LedgerLockUpsertOne {
	return u.Update(func(s *LedgerLockUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *LedgerLockUpsertOne) UpdateAmount() *LedgerLockUpsertOne {
	return u.Update(func(s *LedgerLockUpsert) {
		s.UpdateAmount()
	})
}

// ClearAmount clears the value of the "amount" field.
func (u *LedgerLockUpsertOne) ClearAmount() *LedgerLockUpsertOne {
	return u.Update(func(s *LedgerLockUpsert) {
		s.ClearAmount()
	})
}

// SetLockState sets the "lock_state" field.
func (u *LedgerLockUpsertOne) SetLockState(v string) *LedgerLockUpsertOne {
	return u.Update(func(s *LedgerLockUpsert) {
		s.SetLockState(v)
	})
}

// UpdateLockState sets the "lock_state" field to the value that was provided on create.
func (u *LedgerLockUpsertOne) UpdateLockState() *LedgerLockUpsertOne {
	return u.Update(func(s *LedgerLockUpsert) {
		s.UpdateLockState()
	})
}

// ClearLockState clears the value of the "lock_state" field.
func (u *LedgerLockUpsertOne) ClearLockState() *LedgerLockUpsertOne {
	return u.Update(func(s *LedgerLockUpsert) {
		s.ClearLockState()
	})
}

// SetExLockID sets the "ex_lock_id" field.
func (u *LedgerLockUpsertOne) SetExLockID(v uuid.UUID) *LedgerLockUpsertOne {
	return u.Update(func(s *LedgerLockUpsert) {
		s.SetExLockID(v)
	})
}

// UpdateExLockID sets the "ex_lock_id" field to the value that was provided on create.
func (u *LedgerLockUpsertOne) UpdateExLockID() *LedgerLockUpsertOne {
	return u.Update(func(s *LedgerLockUpsert) {
		s.UpdateExLockID()
	})
}

// ClearExLockID clears the value of the "ex_lock_id" field.
func (u *LedgerLockUpsertOne) ClearExLockID() *LedgerLockUpsertOne {
	return u.Update(func(s *LedgerLockUpsert) {
		s.ClearExLockID()
	})
}

// Exec executes the query.
func (u *LedgerLockUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for LedgerLockCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LedgerLockUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *LedgerLockUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *LedgerLockUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// LedgerLockCreateBulk is the builder for creating many LedgerLock entities in bulk.
type LedgerLockCreateBulk struct {
	config
	err      error
	builders []*LedgerLockCreate
	conflict []sql.ConflictOption
}

// Save creates the LedgerLock entities in the database.
func (llcb *LedgerLockCreateBulk) Save(ctx context.Context) ([]*LedgerLock, error) {
	if llcb.err != nil {
		return nil, llcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(llcb.builders))
	nodes := make([]*LedgerLock, len(llcb.builders))
	mutators := make([]Mutator, len(llcb.builders))
	for i := range llcb.builders {
		func(i int, root context.Context) {
			builder := llcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LedgerLockMutation)
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
					_, err = mutators[i+1].Mutate(root, llcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = llcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, llcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, llcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (llcb *LedgerLockCreateBulk) SaveX(ctx context.Context) []*LedgerLock {
	v, err := llcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (llcb *LedgerLockCreateBulk) Exec(ctx context.Context) error {
	_, err := llcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (llcb *LedgerLockCreateBulk) ExecX(ctx context.Context) {
	if err := llcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.LedgerLock.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LedgerLockUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (llcb *LedgerLockCreateBulk) OnConflict(opts ...sql.ConflictOption) *LedgerLockUpsertBulk {
	llcb.conflict = opts
	return &LedgerLockUpsertBulk{
		create: llcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.LedgerLock.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (llcb *LedgerLockCreateBulk) OnConflictColumns(columns ...string) *LedgerLockUpsertBulk {
	llcb.conflict = append(llcb.conflict, sql.ConflictColumns(columns...))
	return &LedgerLockUpsertBulk{
		create: llcb,
	}
}

// LedgerLockUpsertBulk is the builder for "upsert"-ing
// a bulk of LedgerLock nodes.
type LedgerLockUpsertBulk struct {
	create *LedgerLockCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.LedgerLock.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(ledgerlock.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *LedgerLockUpsertBulk) UpdateNewValues() *LedgerLockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(ledgerlock.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.LedgerLock.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *LedgerLockUpsertBulk) Ignore() *LedgerLockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LedgerLockUpsertBulk) DoNothing() *LedgerLockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LedgerLockCreateBulk.OnConflict
// documentation for more info.
func (u *LedgerLockUpsertBulk) Update(set func(*LedgerLockUpsert)) *LedgerLockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LedgerLockUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *LedgerLockUpsertBulk) SetCreatedAt(v uint32) *LedgerLockUpsertBulk {
	return u.Update(func(s *LedgerLockUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *LedgerLockUpsertBulk) AddCreatedAt(v uint32) *LedgerLockUpsertBulk {
	return u.Update(func(s *LedgerLockUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *LedgerLockUpsertBulk) UpdateCreatedAt() *LedgerLockUpsertBulk {
	return u.Update(func(s *LedgerLockUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *LedgerLockUpsertBulk) SetUpdatedAt(v uint32) *LedgerLockUpsertBulk {
	return u.Update(func(s *LedgerLockUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *LedgerLockUpsertBulk) AddUpdatedAt(v uint32) *LedgerLockUpsertBulk {
	return u.Update(func(s *LedgerLockUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *LedgerLockUpsertBulk) UpdateUpdatedAt() *LedgerLockUpsertBulk {
	return u.Update(func(s *LedgerLockUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *LedgerLockUpsertBulk) SetDeletedAt(v uint32) *LedgerLockUpsertBulk {
	return u.Update(func(s *LedgerLockUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *LedgerLockUpsertBulk) AddDeletedAt(v uint32) *LedgerLockUpsertBulk {
	return u.Update(func(s *LedgerLockUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *LedgerLockUpsertBulk) UpdateDeletedAt() *LedgerLockUpsertBulk {
	return u.Update(func(s *LedgerLockUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *LedgerLockUpsertBulk) SetEntID(v uuid.UUID) *LedgerLockUpsertBulk {
	return u.Update(func(s *LedgerLockUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *LedgerLockUpsertBulk) UpdateEntID() *LedgerLockUpsertBulk {
	return u.Update(func(s *LedgerLockUpsert) {
		s.UpdateEntID()
	})
}

// SetLedgerID sets the "ledger_id" field.
func (u *LedgerLockUpsertBulk) SetLedgerID(v uuid.UUID) *LedgerLockUpsertBulk {
	return u.Update(func(s *LedgerLockUpsert) {
		s.SetLedgerID(v)
	})
}

// UpdateLedgerID sets the "ledger_id" field to the value that was provided on create.
func (u *LedgerLockUpsertBulk) UpdateLedgerID() *LedgerLockUpsertBulk {
	return u.Update(func(s *LedgerLockUpsert) {
		s.UpdateLedgerID()
	})
}

// ClearLedgerID clears the value of the "ledger_id" field.
func (u *LedgerLockUpsertBulk) ClearLedgerID() *LedgerLockUpsertBulk {
	return u.Update(func(s *LedgerLockUpsert) {
		s.ClearLedgerID()
	})
}

// SetStatementID sets the "statement_id" field.
func (u *LedgerLockUpsertBulk) SetStatementID(v uuid.UUID) *LedgerLockUpsertBulk {
	return u.Update(func(s *LedgerLockUpsert) {
		s.SetStatementID(v)
	})
}

// UpdateStatementID sets the "statement_id" field to the value that was provided on create.
func (u *LedgerLockUpsertBulk) UpdateStatementID() *LedgerLockUpsertBulk {
	return u.Update(func(s *LedgerLockUpsert) {
		s.UpdateStatementID()
	})
}

// ClearStatementID clears the value of the "statement_id" field.
func (u *LedgerLockUpsertBulk) ClearStatementID() *LedgerLockUpsertBulk {
	return u.Update(func(s *LedgerLockUpsert) {
		s.ClearStatementID()
	})
}

// SetAmount sets the "amount" field.
func (u *LedgerLockUpsertBulk) SetAmount(v decimal.Decimal) *LedgerLockUpsertBulk {
	return u.Update(func(s *LedgerLockUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *LedgerLockUpsertBulk) AddAmount(v decimal.Decimal) *LedgerLockUpsertBulk {
	return u.Update(func(s *LedgerLockUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *LedgerLockUpsertBulk) UpdateAmount() *LedgerLockUpsertBulk {
	return u.Update(func(s *LedgerLockUpsert) {
		s.UpdateAmount()
	})
}

// ClearAmount clears the value of the "amount" field.
func (u *LedgerLockUpsertBulk) ClearAmount() *LedgerLockUpsertBulk {
	return u.Update(func(s *LedgerLockUpsert) {
		s.ClearAmount()
	})
}

// SetLockState sets the "lock_state" field.
func (u *LedgerLockUpsertBulk) SetLockState(v string) *LedgerLockUpsertBulk {
	return u.Update(func(s *LedgerLockUpsert) {
		s.SetLockState(v)
	})
}

// UpdateLockState sets the "lock_state" field to the value that was provided on create.
func (u *LedgerLockUpsertBulk) UpdateLockState() *LedgerLockUpsertBulk {
	return u.Update(func(s *LedgerLockUpsert) {
		s.UpdateLockState()
	})
}

// ClearLockState clears the value of the "lock_state" field.
func (u *LedgerLockUpsertBulk) ClearLockState() *LedgerLockUpsertBulk {
	return u.Update(func(s *LedgerLockUpsert) {
		s.ClearLockState()
	})
}

// SetExLockID sets the "ex_lock_id" field.
func (u *LedgerLockUpsertBulk) SetExLockID(v uuid.UUID) *LedgerLockUpsertBulk {
	return u.Update(func(s *LedgerLockUpsert) {
		s.SetExLockID(v)
	})
}

// UpdateExLockID sets the "ex_lock_id" field to the value that was provided on create.
func (u *LedgerLockUpsertBulk) UpdateExLockID() *LedgerLockUpsertBulk {
	return u.Update(func(s *LedgerLockUpsert) {
		s.UpdateExLockID()
	})
}

// ClearExLockID clears the value of the "ex_lock_id" field.
func (u *LedgerLockUpsertBulk) ClearExLockID() *LedgerLockUpsertBulk {
	return u.Update(func(s *LedgerLockUpsert) {
		s.ClearExLockID()
	})
}

// Exec executes the query.
func (u *LedgerLockUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("generated: OnConflict was set for builder %d. Set it on the LedgerLockCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for LedgerLockCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LedgerLockUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
