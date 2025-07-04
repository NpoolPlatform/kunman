// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/paymentcontract"
	"github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// PaymentContractUpdate is the builder for updating PaymentContract entities.
type PaymentContractUpdate struct {
	config
	hooks     []Hook
	mutation  *PaymentContractMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PaymentContractUpdate builder.
func (pcu *PaymentContractUpdate) Where(ps ...predicate.PaymentContract) *PaymentContractUpdate {
	pcu.mutation.Where(ps...)
	return pcu
}

// SetEntID sets the "ent_id" field.
func (pcu *PaymentContractUpdate) SetEntID(u uuid.UUID) *PaymentContractUpdate {
	pcu.mutation.SetEntID(u)
	return pcu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (pcu *PaymentContractUpdate) SetNillableEntID(u *uuid.UUID) *PaymentContractUpdate {
	if u != nil {
		pcu.SetEntID(*u)
	}
	return pcu
}

// SetCreatedAt sets the "created_at" field.
func (pcu *PaymentContractUpdate) SetCreatedAt(u uint32) *PaymentContractUpdate {
	pcu.mutation.ResetCreatedAt()
	pcu.mutation.SetCreatedAt(u)
	return pcu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pcu *PaymentContractUpdate) SetNillableCreatedAt(u *uint32) *PaymentContractUpdate {
	if u != nil {
		pcu.SetCreatedAt(*u)
	}
	return pcu
}

// AddCreatedAt adds u to the "created_at" field.
func (pcu *PaymentContractUpdate) AddCreatedAt(u int32) *PaymentContractUpdate {
	pcu.mutation.AddCreatedAt(u)
	return pcu
}

// SetUpdatedAt sets the "updated_at" field.
func (pcu *PaymentContractUpdate) SetUpdatedAt(u uint32) *PaymentContractUpdate {
	pcu.mutation.ResetUpdatedAt()
	pcu.mutation.SetUpdatedAt(u)
	return pcu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (pcu *PaymentContractUpdate) AddUpdatedAt(u int32) *PaymentContractUpdate {
	pcu.mutation.AddUpdatedAt(u)
	return pcu
}

// SetDeletedAt sets the "deleted_at" field.
func (pcu *PaymentContractUpdate) SetDeletedAt(u uint32) *PaymentContractUpdate {
	pcu.mutation.ResetDeletedAt()
	pcu.mutation.SetDeletedAt(u)
	return pcu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pcu *PaymentContractUpdate) SetNillableDeletedAt(u *uint32) *PaymentContractUpdate {
	if u != nil {
		pcu.SetDeletedAt(*u)
	}
	return pcu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (pcu *PaymentContractUpdate) AddDeletedAt(u int32) *PaymentContractUpdate {
	pcu.mutation.AddDeletedAt(u)
	return pcu
}

// SetOrderID sets the "order_id" field.
func (pcu *PaymentContractUpdate) SetOrderID(u uuid.UUID) *PaymentContractUpdate {
	pcu.mutation.SetOrderID(u)
	return pcu
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (pcu *PaymentContractUpdate) SetNillableOrderID(u *uuid.UUID) *PaymentContractUpdate {
	if u != nil {
		pcu.SetOrderID(*u)
	}
	return pcu
}

// ClearOrderID clears the value of the "order_id" field.
func (pcu *PaymentContractUpdate) ClearOrderID() *PaymentContractUpdate {
	pcu.mutation.ClearOrderID()
	return pcu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (pcu *PaymentContractUpdate) SetCoinTypeID(u uuid.UUID) *PaymentContractUpdate {
	pcu.mutation.SetCoinTypeID(u)
	return pcu
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (pcu *PaymentContractUpdate) SetNillableCoinTypeID(u *uuid.UUID) *PaymentContractUpdate {
	if u != nil {
		pcu.SetCoinTypeID(*u)
	}
	return pcu
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (pcu *PaymentContractUpdate) ClearCoinTypeID() *PaymentContractUpdate {
	pcu.mutation.ClearCoinTypeID()
	return pcu
}

// SetAmount sets the "amount" field.
func (pcu *PaymentContractUpdate) SetAmount(d decimal.Decimal) *PaymentContractUpdate {
	pcu.mutation.SetAmount(d)
	return pcu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (pcu *PaymentContractUpdate) SetNillableAmount(d *decimal.Decimal) *PaymentContractUpdate {
	if d != nil {
		pcu.SetAmount(*d)
	}
	return pcu
}

// ClearAmount clears the value of the "amount" field.
func (pcu *PaymentContractUpdate) ClearAmount() *PaymentContractUpdate {
	pcu.mutation.ClearAmount()
	return pcu
}

// Mutation returns the PaymentContractMutation object of the builder.
func (pcu *PaymentContractUpdate) Mutation() *PaymentContractMutation {
	return pcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pcu *PaymentContractUpdate) Save(ctx context.Context) (int, error) {
	pcu.defaults()
	return withHooks(ctx, pcu.sqlSave, pcu.mutation, pcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pcu *PaymentContractUpdate) SaveX(ctx context.Context) int {
	affected, err := pcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pcu *PaymentContractUpdate) Exec(ctx context.Context) error {
	_, err := pcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcu *PaymentContractUpdate) ExecX(ctx context.Context) {
	if err := pcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pcu *PaymentContractUpdate) defaults() {
	if _, ok := pcu.mutation.UpdatedAt(); !ok {
		v := paymentcontract.UpdateDefaultUpdatedAt()
		pcu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pcu *PaymentContractUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PaymentContractUpdate {
	pcu.modifiers = append(pcu.modifiers, modifiers...)
	return pcu
}

func (pcu *PaymentContractUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(paymentcontract.Table, paymentcontract.Columns, sqlgraph.NewFieldSpec(paymentcontract.FieldID, field.TypeUint32))
	if ps := pcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pcu.mutation.EntID(); ok {
		_spec.SetField(paymentcontract.FieldEntID, field.TypeUUID, value)
	}
	if value, ok := pcu.mutation.CreatedAt(); ok {
		_spec.SetField(paymentcontract.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := pcu.mutation.AddedCreatedAt(); ok {
		_spec.AddField(paymentcontract.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := pcu.mutation.UpdatedAt(); ok {
		_spec.SetField(paymentcontract.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := pcu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(paymentcontract.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := pcu.mutation.DeletedAt(); ok {
		_spec.SetField(paymentcontract.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := pcu.mutation.AddedDeletedAt(); ok {
		_spec.AddField(paymentcontract.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := pcu.mutation.OrderID(); ok {
		_spec.SetField(paymentcontract.FieldOrderID, field.TypeUUID, value)
	}
	if pcu.mutation.OrderIDCleared() {
		_spec.ClearField(paymentcontract.FieldOrderID, field.TypeUUID)
	}
	if value, ok := pcu.mutation.CoinTypeID(); ok {
		_spec.SetField(paymentcontract.FieldCoinTypeID, field.TypeUUID, value)
	}
	if pcu.mutation.CoinTypeIDCleared() {
		_spec.ClearField(paymentcontract.FieldCoinTypeID, field.TypeUUID)
	}
	if value, ok := pcu.mutation.Amount(); ok {
		_spec.SetField(paymentcontract.FieldAmount, field.TypeOther, value)
	}
	if pcu.mutation.AmountCleared() {
		_spec.ClearField(paymentcontract.FieldAmount, field.TypeOther)
	}
	_spec.AddModifiers(pcu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, pcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{paymentcontract.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pcu.mutation.done = true
	return n, nil
}

// PaymentContractUpdateOne is the builder for updating a single PaymentContract entity.
type PaymentContractUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PaymentContractMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetEntID sets the "ent_id" field.
func (pcuo *PaymentContractUpdateOne) SetEntID(u uuid.UUID) *PaymentContractUpdateOne {
	pcuo.mutation.SetEntID(u)
	return pcuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (pcuo *PaymentContractUpdateOne) SetNillableEntID(u *uuid.UUID) *PaymentContractUpdateOne {
	if u != nil {
		pcuo.SetEntID(*u)
	}
	return pcuo
}

// SetCreatedAt sets the "created_at" field.
func (pcuo *PaymentContractUpdateOne) SetCreatedAt(u uint32) *PaymentContractUpdateOne {
	pcuo.mutation.ResetCreatedAt()
	pcuo.mutation.SetCreatedAt(u)
	return pcuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pcuo *PaymentContractUpdateOne) SetNillableCreatedAt(u *uint32) *PaymentContractUpdateOne {
	if u != nil {
		pcuo.SetCreatedAt(*u)
	}
	return pcuo
}

// AddCreatedAt adds u to the "created_at" field.
func (pcuo *PaymentContractUpdateOne) AddCreatedAt(u int32) *PaymentContractUpdateOne {
	pcuo.mutation.AddCreatedAt(u)
	return pcuo
}

// SetUpdatedAt sets the "updated_at" field.
func (pcuo *PaymentContractUpdateOne) SetUpdatedAt(u uint32) *PaymentContractUpdateOne {
	pcuo.mutation.ResetUpdatedAt()
	pcuo.mutation.SetUpdatedAt(u)
	return pcuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (pcuo *PaymentContractUpdateOne) AddUpdatedAt(u int32) *PaymentContractUpdateOne {
	pcuo.mutation.AddUpdatedAt(u)
	return pcuo
}

// SetDeletedAt sets the "deleted_at" field.
func (pcuo *PaymentContractUpdateOne) SetDeletedAt(u uint32) *PaymentContractUpdateOne {
	pcuo.mutation.ResetDeletedAt()
	pcuo.mutation.SetDeletedAt(u)
	return pcuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pcuo *PaymentContractUpdateOne) SetNillableDeletedAt(u *uint32) *PaymentContractUpdateOne {
	if u != nil {
		pcuo.SetDeletedAt(*u)
	}
	return pcuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (pcuo *PaymentContractUpdateOne) AddDeletedAt(u int32) *PaymentContractUpdateOne {
	pcuo.mutation.AddDeletedAt(u)
	return pcuo
}

// SetOrderID sets the "order_id" field.
func (pcuo *PaymentContractUpdateOne) SetOrderID(u uuid.UUID) *PaymentContractUpdateOne {
	pcuo.mutation.SetOrderID(u)
	return pcuo
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (pcuo *PaymentContractUpdateOne) SetNillableOrderID(u *uuid.UUID) *PaymentContractUpdateOne {
	if u != nil {
		pcuo.SetOrderID(*u)
	}
	return pcuo
}

// ClearOrderID clears the value of the "order_id" field.
func (pcuo *PaymentContractUpdateOne) ClearOrderID() *PaymentContractUpdateOne {
	pcuo.mutation.ClearOrderID()
	return pcuo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (pcuo *PaymentContractUpdateOne) SetCoinTypeID(u uuid.UUID) *PaymentContractUpdateOne {
	pcuo.mutation.SetCoinTypeID(u)
	return pcuo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (pcuo *PaymentContractUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *PaymentContractUpdateOne {
	if u != nil {
		pcuo.SetCoinTypeID(*u)
	}
	return pcuo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (pcuo *PaymentContractUpdateOne) ClearCoinTypeID() *PaymentContractUpdateOne {
	pcuo.mutation.ClearCoinTypeID()
	return pcuo
}

// SetAmount sets the "amount" field.
func (pcuo *PaymentContractUpdateOne) SetAmount(d decimal.Decimal) *PaymentContractUpdateOne {
	pcuo.mutation.SetAmount(d)
	return pcuo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (pcuo *PaymentContractUpdateOne) SetNillableAmount(d *decimal.Decimal) *PaymentContractUpdateOne {
	if d != nil {
		pcuo.SetAmount(*d)
	}
	return pcuo
}

// ClearAmount clears the value of the "amount" field.
func (pcuo *PaymentContractUpdateOne) ClearAmount() *PaymentContractUpdateOne {
	pcuo.mutation.ClearAmount()
	return pcuo
}

// Mutation returns the PaymentContractMutation object of the builder.
func (pcuo *PaymentContractUpdateOne) Mutation() *PaymentContractMutation {
	return pcuo.mutation
}

// Where appends a list predicates to the PaymentContractUpdate builder.
func (pcuo *PaymentContractUpdateOne) Where(ps ...predicate.PaymentContract) *PaymentContractUpdateOne {
	pcuo.mutation.Where(ps...)
	return pcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pcuo *PaymentContractUpdateOne) Select(field string, fields ...string) *PaymentContractUpdateOne {
	pcuo.fields = append([]string{field}, fields...)
	return pcuo
}

// Save executes the query and returns the updated PaymentContract entity.
func (pcuo *PaymentContractUpdateOne) Save(ctx context.Context) (*PaymentContract, error) {
	pcuo.defaults()
	return withHooks(ctx, pcuo.sqlSave, pcuo.mutation, pcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pcuo *PaymentContractUpdateOne) SaveX(ctx context.Context) *PaymentContract {
	node, err := pcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pcuo *PaymentContractUpdateOne) Exec(ctx context.Context) error {
	_, err := pcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcuo *PaymentContractUpdateOne) ExecX(ctx context.Context) {
	if err := pcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pcuo *PaymentContractUpdateOne) defaults() {
	if _, ok := pcuo.mutation.UpdatedAt(); !ok {
		v := paymentcontract.UpdateDefaultUpdatedAt()
		pcuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pcuo *PaymentContractUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PaymentContractUpdateOne {
	pcuo.modifiers = append(pcuo.modifiers, modifiers...)
	return pcuo
}

func (pcuo *PaymentContractUpdateOne) sqlSave(ctx context.Context) (_node *PaymentContract, err error) {
	_spec := sqlgraph.NewUpdateSpec(paymentcontract.Table, paymentcontract.Columns, sqlgraph.NewFieldSpec(paymentcontract.FieldID, field.TypeUint32))
	id, ok := pcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "PaymentContract.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, paymentcontract.FieldID)
		for _, f := range fields {
			if !paymentcontract.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != paymentcontract.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pcuo.mutation.EntID(); ok {
		_spec.SetField(paymentcontract.FieldEntID, field.TypeUUID, value)
	}
	if value, ok := pcuo.mutation.CreatedAt(); ok {
		_spec.SetField(paymentcontract.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := pcuo.mutation.AddedCreatedAt(); ok {
		_spec.AddField(paymentcontract.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := pcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(paymentcontract.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := pcuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(paymentcontract.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := pcuo.mutation.DeletedAt(); ok {
		_spec.SetField(paymentcontract.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := pcuo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(paymentcontract.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := pcuo.mutation.OrderID(); ok {
		_spec.SetField(paymentcontract.FieldOrderID, field.TypeUUID, value)
	}
	if pcuo.mutation.OrderIDCleared() {
		_spec.ClearField(paymentcontract.FieldOrderID, field.TypeUUID)
	}
	if value, ok := pcuo.mutation.CoinTypeID(); ok {
		_spec.SetField(paymentcontract.FieldCoinTypeID, field.TypeUUID, value)
	}
	if pcuo.mutation.CoinTypeIDCleared() {
		_spec.ClearField(paymentcontract.FieldCoinTypeID, field.TypeUUID)
	}
	if value, ok := pcuo.mutation.Amount(); ok {
		_spec.SetField(paymentcontract.FieldAmount, field.TypeOther, value)
	}
	if pcuo.mutation.AmountCleared() {
		_spec.ClearField(paymentcontract.FieldAmount, field.TypeOther)
	}
	_spec.AddModifiers(pcuo.modifiers...)
	_node = &PaymentContract{config: pcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{paymentcontract.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pcuo.mutation.done = true
	return _node, nil
}
