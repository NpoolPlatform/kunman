// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/goodledger"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// GoodLedgerUpdate is the builder for updating GoodLedger entities.
type GoodLedgerUpdate struct {
	config
	hooks     []Hook
	mutation  *GoodLedgerMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the GoodLedgerUpdate builder.
func (glu *GoodLedgerUpdate) Where(ps ...predicate.GoodLedger) *GoodLedgerUpdate {
	glu.mutation.Where(ps...)
	return glu
}

// SetCreatedAt sets the "created_at" field.
func (glu *GoodLedgerUpdate) SetCreatedAt(u uint32) *GoodLedgerUpdate {
	glu.mutation.ResetCreatedAt()
	glu.mutation.SetCreatedAt(u)
	return glu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (glu *GoodLedgerUpdate) SetNillableCreatedAt(u *uint32) *GoodLedgerUpdate {
	if u != nil {
		glu.SetCreatedAt(*u)
	}
	return glu
}

// AddCreatedAt adds u to the "created_at" field.
func (glu *GoodLedgerUpdate) AddCreatedAt(u int32) *GoodLedgerUpdate {
	glu.mutation.AddCreatedAt(u)
	return glu
}

// SetUpdatedAt sets the "updated_at" field.
func (glu *GoodLedgerUpdate) SetUpdatedAt(u uint32) *GoodLedgerUpdate {
	glu.mutation.ResetUpdatedAt()
	glu.mutation.SetUpdatedAt(u)
	return glu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (glu *GoodLedgerUpdate) AddUpdatedAt(u int32) *GoodLedgerUpdate {
	glu.mutation.AddUpdatedAt(u)
	return glu
}

// SetDeletedAt sets the "deleted_at" field.
func (glu *GoodLedgerUpdate) SetDeletedAt(u uint32) *GoodLedgerUpdate {
	glu.mutation.ResetDeletedAt()
	glu.mutation.SetDeletedAt(u)
	return glu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (glu *GoodLedgerUpdate) SetNillableDeletedAt(u *uint32) *GoodLedgerUpdate {
	if u != nil {
		glu.SetDeletedAt(*u)
	}
	return glu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (glu *GoodLedgerUpdate) AddDeletedAt(u int32) *GoodLedgerUpdate {
	glu.mutation.AddDeletedAt(u)
	return glu
}

// SetEntID sets the "ent_id" field.
func (glu *GoodLedgerUpdate) SetEntID(u uuid.UUID) *GoodLedgerUpdate {
	glu.mutation.SetEntID(u)
	return glu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (glu *GoodLedgerUpdate) SetNillableEntID(u *uuid.UUID) *GoodLedgerUpdate {
	if u != nil {
		glu.SetEntID(*u)
	}
	return glu
}

// SetGoodID sets the "good_id" field.
func (glu *GoodLedgerUpdate) SetGoodID(u uuid.UUID) *GoodLedgerUpdate {
	glu.mutation.SetGoodID(u)
	return glu
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (glu *GoodLedgerUpdate) SetNillableGoodID(u *uuid.UUID) *GoodLedgerUpdate {
	if u != nil {
		glu.SetGoodID(*u)
	}
	return glu
}

// ClearGoodID clears the value of the "good_id" field.
func (glu *GoodLedgerUpdate) ClearGoodID() *GoodLedgerUpdate {
	glu.mutation.ClearGoodID()
	return glu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (glu *GoodLedgerUpdate) SetCoinTypeID(u uuid.UUID) *GoodLedgerUpdate {
	glu.mutation.SetCoinTypeID(u)
	return glu
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (glu *GoodLedgerUpdate) SetNillableCoinTypeID(u *uuid.UUID) *GoodLedgerUpdate {
	if u != nil {
		glu.SetCoinTypeID(*u)
	}
	return glu
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (glu *GoodLedgerUpdate) ClearCoinTypeID() *GoodLedgerUpdate {
	glu.mutation.ClearCoinTypeID()
	return glu
}

// SetAmount sets the "amount" field.
func (glu *GoodLedgerUpdate) SetAmount(d decimal.Decimal) *GoodLedgerUpdate {
	glu.mutation.ResetAmount()
	glu.mutation.SetAmount(d)
	return glu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (glu *GoodLedgerUpdate) SetNillableAmount(d *decimal.Decimal) *GoodLedgerUpdate {
	if d != nil {
		glu.SetAmount(*d)
	}
	return glu
}

// AddAmount adds d to the "amount" field.
func (glu *GoodLedgerUpdate) AddAmount(d decimal.Decimal) *GoodLedgerUpdate {
	glu.mutation.AddAmount(d)
	return glu
}

// ClearAmount clears the value of the "amount" field.
func (glu *GoodLedgerUpdate) ClearAmount() *GoodLedgerUpdate {
	glu.mutation.ClearAmount()
	return glu
}

// SetToPlatform sets the "to_platform" field.
func (glu *GoodLedgerUpdate) SetToPlatform(d decimal.Decimal) *GoodLedgerUpdate {
	glu.mutation.ResetToPlatform()
	glu.mutation.SetToPlatform(d)
	return glu
}

// SetNillableToPlatform sets the "to_platform" field if the given value is not nil.
func (glu *GoodLedgerUpdate) SetNillableToPlatform(d *decimal.Decimal) *GoodLedgerUpdate {
	if d != nil {
		glu.SetToPlatform(*d)
	}
	return glu
}

// AddToPlatform adds d to the "to_platform" field.
func (glu *GoodLedgerUpdate) AddToPlatform(d decimal.Decimal) *GoodLedgerUpdate {
	glu.mutation.AddToPlatform(d)
	return glu
}

// ClearToPlatform clears the value of the "to_platform" field.
func (glu *GoodLedgerUpdate) ClearToPlatform() *GoodLedgerUpdate {
	glu.mutation.ClearToPlatform()
	return glu
}

// SetToUser sets the "to_user" field.
func (glu *GoodLedgerUpdate) SetToUser(d decimal.Decimal) *GoodLedgerUpdate {
	glu.mutation.ResetToUser()
	glu.mutation.SetToUser(d)
	return glu
}

// SetNillableToUser sets the "to_user" field if the given value is not nil.
func (glu *GoodLedgerUpdate) SetNillableToUser(d *decimal.Decimal) *GoodLedgerUpdate {
	if d != nil {
		glu.SetToUser(*d)
	}
	return glu
}

// AddToUser adds d to the "to_user" field.
func (glu *GoodLedgerUpdate) AddToUser(d decimal.Decimal) *GoodLedgerUpdate {
	glu.mutation.AddToUser(d)
	return glu
}

// ClearToUser clears the value of the "to_user" field.
func (glu *GoodLedgerUpdate) ClearToUser() *GoodLedgerUpdate {
	glu.mutation.ClearToUser()
	return glu
}

// Mutation returns the GoodLedgerMutation object of the builder.
func (glu *GoodLedgerUpdate) Mutation() *GoodLedgerMutation {
	return glu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (glu *GoodLedgerUpdate) Save(ctx context.Context) (int, error) {
	glu.defaults()
	return withHooks(ctx, glu.sqlSave, glu.mutation, glu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (glu *GoodLedgerUpdate) SaveX(ctx context.Context) int {
	affected, err := glu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (glu *GoodLedgerUpdate) Exec(ctx context.Context) error {
	_, err := glu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (glu *GoodLedgerUpdate) ExecX(ctx context.Context) {
	if err := glu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (glu *GoodLedgerUpdate) defaults() {
	if _, ok := glu.mutation.UpdatedAt(); !ok {
		v := goodledger.UpdateDefaultUpdatedAt()
		glu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (glu *GoodLedgerUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GoodLedgerUpdate {
	glu.modifiers = append(glu.modifiers, modifiers...)
	return glu
}

func (glu *GoodLedgerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(goodledger.Table, goodledger.Columns, sqlgraph.NewFieldSpec(goodledger.FieldID, field.TypeUint32))
	if ps := glu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := glu.mutation.CreatedAt(); ok {
		_spec.SetField(goodledger.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := glu.mutation.AddedCreatedAt(); ok {
		_spec.AddField(goodledger.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := glu.mutation.UpdatedAt(); ok {
		_spec.SetField(goodledger.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := glu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(goodledger.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := glu.mutation.DeletedAt(); ok {
		_spec.SetField(goodledger.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := glu.mutation.AddedDeletedAt(); ok {
		_spec.AddField(goodledger.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := glu.mutation.EntID(); ok {
		_spec.SetField(goodledger.FieldEntID, field.TypeUUID, value)
	}
	if value, ok := glu.mutation.GoodID(); ok {
		_spec.SetField(goodledger.FieldGoodID, field.TypeUUID, value)
	}
	if glu.mutation.GoodIDCleared() {
		_spec.ClearField(goodledger.FieldGoodID, field.TypeUUID)
	}
	if value, ok := glu.mutation.CoinTypeID(); ok {
		_spec.SetField(goodledger.FieldCoinTypeID, field.TypeUUID, value)
	}
	if glu.mutation.CoinTypeIDCleared() {
		_spec.ClearField(goodledger.FieldCoinTypeID, field.TypeUUID)
	}
	if value, ok := glu.mutation.Amount(); ok {
		_spec.SetField(goodledger.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := glu.mutation.AddedAmount(); ok {
		_spec.AddField(goodledger.FieldAmount, field.TypeFloat64, value)
	}
	if glu.mutation.AmountCleared() {
		_spec.ClearField(goodledger.FieldAmount, field.TypeFloat64)
	}
	if value, ok := glu.mutation.ToPlatform(); ok {
		_spec.SetField(goodledger.FieldToPlatform, field.TypeFloat64, value)
	}
	if value, ok := glu.mutation.AddedToPlatform(); ok {
		_spec.AddField(goodledger.FieldToPlatform, field.TypeFloat64, value)
	}
	if glu.mutation.ToPlatformCleared() {
		_spec.ClearField(goodledger.FieldToPlatform, field.TypeFloat64)
	}
	if value, ok := glu.mutation.ToUser(); ok {
		_spec.SetField(goodledger.FieldToUser, field.TypeFloat64, value)
	}
	if value, ok := glu.mutation.AddedToUser(); ok {
		_spec.AddField(goodledger.FieldToUser, field.TypeFloat64, value)
	}
	if glu.mutation.ToUserCleared() {
		_spec.ClearField(goodledger.FieldToUser, field.TypeFloat64)
	}
	_spec.AddModifiers(glu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, glu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goodledger.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	glu.mutation.done = true
	return n, nil
}

// GoodLedgerUpdateOne is the builder for updating a single GoodLedger entity.
type GoodLedgerUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *GoodLedgerMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (gluo *GoodLedgerUpdateOne) SetCreatedAt(u uint32) *GoodLedgerUpdateOne {
	gluo.mutation.ResetCreatedAt()
	gluo.mutation.SetCreatedAt(u)
	return gluo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gluo *GoodLedgerUpdateOne) SetNillableCreatedAt(u *uint32) *GoodLedgerUpdateOne {
	if u != nil {
		gluo.SetCreatedAt(*u)
	}
	return gluo
}

// AddCreatedAt adds u to the "created_at" field.
func (gluo *GoodLedgerUpdateOne) AddCreatedAt(u int32) *GoodLedgerUpdateOne {
	gluo.mutation.AddCreatedAt(u)
	return gluo
}

// SetUpdatedAt sets the "updated_at" field.
func (gluo *GoodLedgerUpdateOne) SetUpdatedAt(u uint32) *GoodLedgerUpdateOne {
	gluo.mutation.ResetUpdatedAt()
	gluo.mutation.SetUpdatedAt(u)
	return gluo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (gluo *GoodLedgerUpdateOne) AddUpdatedAt(u int32) *GoodLedgerUpdateOne {
	gluo.mutation.AddUpdatedAt(u)
	return gluo
}

// SetDeletedAt sets the "deleted_at" field.
func (gluo *GoodLedgerUpdateOne) SetDeletedAt(u uint32) *GoodLedgerUpdateOne {
	gluo.mutation.ResetDeletedAt()
	gluo.mutation.SetDeletedAt(u)
	return gluo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gluo *GoodLedgerUpdateOne) SetNillableDeletedAt(u *uint32) *GoodLedgerUpdateOne {
	if u != nil {
		gluo.SetDeletedAt(*u)
	}
	return gluo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (gluo *GoodLedgerUpdateOne) AddDeletedAt(u int32) *GoodLedgerUpdateOne {
	gluo.mutation.AddDeletedAt(u)
	return gluo
}

// SetEntID sets the "ent_id" field.
func (gluo *GoodLedgerUpdateOne) SetEntID(u uuid.UUID) *GoodLedgerUpdateOne {
	gluo.mutation.SetEntID(u)
	return gluo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (gluo *GoodLedgerUpdateOne) SetNillableEntID(u *uuid.UUID) *GoodLedgerUpdateOne {
	if u != nil {
		gluo.SetEntID(*u)
	}
	return gluo
}

// SetGoodID sets the "good_id" field.
func (gluo *GoodLedgerUpdateOne) SetGoodID(u uuid.UUID) *GoodLedgerUpdateOne {
	gluo.mutation.SetGoodID(u)
	return gluo
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (gluo *GoodLedgerUpdateOne) SetNillableGoodID(u *uuid.UUID) *GoodLedgerUpdateOne {
	if u != nil {
		gluo.SetGoodID(*u)
	}
	return gluo
}

// ClearGoodID clears the value of the "good_id" field.
func (gluo *GoodLedgerUpdateOne) ClearGoodID() *GoodLedgerUpdateOne {
	gluo.mutation.ClearGoodID()
	return gluo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (gluo *GoodLedgerUpdateOne) SetCoinTypeID(u uuid.UUID) *GoodLedgerUpdateOne {
	gluo.mutation.SetCoinTypeID(u)
	return gluo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (gluo *GoodLedgerUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *GoodLedgerUpdateOne {
	if u != nil {
		gluo.SetCoinTypeID(*u)
	}
	return gluo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (gluo *GoodLedgerUpdateOne) ClearCoinTypeID() *GoodLedgerUpdateOne {
	gluo.mutation.ClearCoinTypeID()
	return gluo
}

// SetAmount sets the "amount" field.
func (gluo *GoodLedgerUpdateOne) SetAmount(d decimal.Decimal) *GoodLedgerUpdateOne {
	gluo.mutation.ResetAmount()
	gluo.mutation.SetAmount(d)
	return gluo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (gluo *GoodLedgerUpdateOne) SetNillableAmount(d *decimal.Decimal) *GoodLedgerUpdateOne {
	if d != nil {
		gluo.SetAmount(*d)
	}
	return gluo
}

// AddAmount adds d to the "amount" field.
func (gluo *GoodLedgerUpdateOne) AddAmount(d decimal.Decimal) *GoodLedgerUpdateOne {
	gluo.mutation.AddAmount(d)
	return gluo
}

// ClearAmount clears the value of the "amount" field.
func (gluo *GoodLedgerUpdateOne) ClearAmount() *GoodLedgerUpdateOne {
	gluo.mutation.ClearAmount()
	return gluo
}

// SetToPlatform sets the "to_platform" field.
func (gluo *GoodLedgerUpdateOne) SetToPlatform(d decimal.Decimal) *GoodLedgerUpdateOne {
	gluo.mutation.ResetToPlatform()
	gluo.mutation.SetToPlatform(d)
	return gluo
}

// SetNillableToPlatform sets the "to_platform" field if the given value is not nil.
func (gluo *GoodLedgerUpdateOne) SetNillableToPlatform(d *decimal.Decimal) *GoodLedgerUpdateOne {
	if d != nil {
		gluo.SetToPlatform(*d)
	}
	return gluo
}

// AddToPlatform adds d to the "to_platform" field.
func (gluo *GoodLedgerUpdateOne) AddToPlatform(d decimal.Decimal) *GoodLedgerUpdateOne {
	gluo.mutation.AddToPlatform(d)
	return gluo
}

// ClearToPlatform clears the value of the "to_platform" field.
func (gluo *GoodLedgerUpdateOne) ClearToPlatform() *GoodLedgerUpdateOne {
	gluo.mutation.ClearToPlatform()
	return gluo
}

// SetToUser sets the "to_user" field.
func (gluo *GoodLedgerUpdateOne) SetToUser(d decimal.Decimal) *GoodLedgerUpdateOne {
	gluo.mutation.ResetToUser()
	gluo.mutation.SetToUser(d)
	return gluo
}

// SetNillableToUser sets the "to_user" field if the given value is not nil.
func (gluo *GoodLedgerUpdateOne) SetNillableToUser(d *decimal.Decimal) *GoodLedgerUpdateOne {
	if d != nil {
		gluo.SetToUser(*d)
	}
	return gluo
}

// AddToUser adds d to the "to_user" field.
func (gluo *GoodLedgerUpdateOne) AddToUser(d decimal.Decimal) *GoodLedgerUpdateOne {
	gluo.mutation.AddToUser(d)
	return gluo
}

// ClearToUser clears the value of the "to_user" field.
func (gluo *GoodLedgerUpdateOne) ClearToUser() *GoodLedgerUpdateOne {
	gluo.mutation.ClearToUser()
	return gluo
}

// Mutation returns the GoodLedgerMutation object of the builder.
func (gluo *GoodLedgerUpdateOne) Mutation() *GoodLedgerMutation {
	return gluo.mutation
}

// Where appends a list predicates to the GoodLedgerUpdate builder.
func (gluo *GoodLedgerUpdateOne) Where(ps ...predicate.GoodLedger) *GoodLedgerUpdateOne {
	gluo.mutation.Where(ps...)
	return gluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gluo *GoodLedgerUpdateOne) Select(field string, fields ...string) *GoodLedgerUpdateOne {
	gluo.fields = append([]string{field}, fields...)
	return gluo
}

// Save executes the query and returns the updated GoodLedger entity.
func (gluo *GoodLedgerUpdateOne) Save(ctx context.Context) (*GoodLedger, error) {
	gluo.defaults()
	return withHooks(ctx, gluo.sqlSave, gluo.mutation, gluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gluo *GoodLedgerUpdateOne) SaveX(ctx context.Context) *GoodLedger {
	node, err := gluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gluo *GoodLedgerUpdateOne) Exec(ctx context.Context) error {
	_, err := gluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gluo *GoodLedgerUpdateOne) ExecX(ctx context.Context) {
	if err := gluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gluo *GoodLedgerUpdateOne) defaults() {
	if _, ok := gluo.mutation.UpdatedAt(); !ok {
		v := goodledger.UpdateDefaultUpdatedAt()
		gluo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (gluo *GoodLedgerUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GoodLedgerUpdateOne {
	gluo.modifiers = append(gluo.modifiers, modifiers...)
	return gluo
}

func (gluo *GoodLedgerUpdateOne) sqlSave(ctx context.Context) (_node *GoodLedger, err error) {
	_spec := sqlgraph.NewUpdateSpec(goodledger.Table, goodledger.Columns, sqlgraph.NewFieldSpec(goodledger.FieldID, field.TypeUint32))
	id, ok := gluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "GoodLedger.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodledger.FieldID)
		for _, f := range fields {
			if !goodledger.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != goodledger.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gluo.mutation.CreatedAt(); ok {
		_spec.SetField(goodledger.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := gluo.mutation.AddedCreatedAt(); ok {
		_spec.AddField(goodledger.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := gluo.mutation.UpdatedAt(); ok {
		_spec.SetField(goodledger.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := gluo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(goodledger.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := gluo.mutation.DeletedAt(); ok {
		_spec.SetField(goodledger.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := gluo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(goodledger.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := gluo.mutation.EntID(); ok {
		_spec.SetField(goodledger.FieldEntID, field.TypeUUID, value)
	}
	if value, ok := gluo.mutation.GoodID(); ok {
		_spec.SetField(goodledger.FieldGoodID, field.TypeUUID, value)
	}
	if gluo.mutation.GoodIDCleared() {
		_spec.ClearField(goodledger.FieldGoodID, field.TypeUUID)
	}
	if value, ok := gluo.mutation.CoinTypeID(); ok {
		_spec.SetField(goodledger.FieldCoinTypeID, field.TypeUUID, value)
	}
	if gluo.mutation.CoinTypeIDCleared() {
		_spec.ClearField(goodledger.FieldCoinTypeID, field.TypeUUID)
	}
	if value, ok := gluo.mutation.Amount(); ok {
		_spec.SetField(goodledger.FieldAmount, field.TypeFloat64, value)
	}
	if value, ok := gluo.mutation.AddedAmount(); ok {
		_spec.AddField(goodledger.FieldAmount, field.TypeFloat64, value)
	}
	if gluo.mutation.AmountCleared() {
		_spec.ClearField(goodledger.FieldAmount, field.TypeFloat64)
	}
	if value, ok := gluo.mutation.ToPlatform(); ok {
		_spec.SetField(goodledger.FieldToPlatform, field.TypeFloat64, value)
	}
	if value, ok := gluo.mutation.AddedToPlatform(); ok {
		_spec.AddField(goodledger.FieldToPlatform, field.TypeFloat64, value)
	}
	if gluo.mutation.ToPlatformCleared() {
		_spec.ClearField(goodledger.FieldToPlatform, field.TypeFloat64)
	}
	if value, ok := gluo.mutation.ToUser(); ok {
		_spec.SetField(goodledger.FieldToUser, field.TypeFloat64, value)
	}
	if value, ok := gluo.mutation.AddedToUser(); ok {
		_spec.AddField(goodledger.FieldToUser, field.TypeFloat64, value)
	}
	if gluo.mutation.ToUserCleared() {
		_spec.ClearField(goodledger.FieldToUser, field.TypeFloat64)
	}
	_spec.AddModifiers(gluo.modifiers...)
	_node = &GoodLedger{config: gluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goodledger.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	gluo.mutation.done = true
	return _node, nil
}
