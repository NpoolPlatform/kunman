// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/feeorder"
	"github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// FeeOrderUpdate is the builder for updating FeeOrder entities.
type FeeOrderUpdate struct {
	config
	hooks     []Hook
	mutation  *FeeOrderMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the FeeOrderUpdate builder.
func (fou *FeeOrderUpdate) Where(ps ...predicate.FeeOrder) *FeeOrderUpdate {
	fou.mutation.Where(ps...)
	return fou
}

// SetEntID sets the "ent_id" field.
func (fou *FeeOrderUpdate) SetEntID(u uuid.UUID) *FeeOrderUpdate {
	fou.mutation.SetEntID(u)
	return fou
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (fou *FeeOrderUpdate) SetNillableEntID(u *uuid.UUID) *FeeOrderUpdate {
	if u != nil {
		fou.SetEntID(*u)
	}
	return fou
}

// SetCreatedAt sets the "created_at" field.
func (fou *FeeOrderUpdate) SetCreatedAt(u uint32) *FeeOrderUpdate {
	fou.mutation.ResetCreatedAt()
	fou.mutation.SetCreatedAt(u)
	return fou
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fou *FeeOrderUpdate) SetNillableCreatedAt(u *uint32) *FeeOrderUpdate {
	if u != nil {
		fou.SetCreatedAt(*u)
	}
	return fou
}

// AddCreatedAt adds u to the "created_at" field.
func (fou *FeeOrderUpdate) AddCreatedAt(u int32) *FeeOrderUpdate {
	fou.mutation.AddCreatedAt(u)
	return fou
}

// SetUpdatedAt sets the "updated_at" field.
func (fou *FeeOrderUpdate) SetUpdatedAt(u uint32) *FeeOrderUpdate {
	fou.mutation.ResetUpdatedAt()
	fou.mutation.SetUpdatedAt(u)
	return fou
}

// AddUpdatedAt adds u to the "updated_at" field.
func (fou *FeeOrderUpdate) AddUpdatedAt(u int32) *FeeOrderUpdate {
	fou.mutation.AddUpdatedAt(u)
	return fou
}

// SetDeletedAt sets the "deleted_at" field.
func (fou *FeeOrderUpdate) SetDeletedAt(u uint32) *FeeOrderUpdate {
	fou.mutation.ResetDeletedAt()
	fou.mutation.SetDeletedAt(u)
	return fou
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fou *FeeOrderUpdate) SetNillableDeletedAt(u *uint32) *FeeOrderUpdate {
	if u != nil {
		fou.SetDeletedAt(*u)
	}
	return fou
}

// AddDeletedAt adds u to the "deleted_at" field.
func (fou *FeeOrderUpdate) AddDeletedAt(u int32) *FeeOrderUpdate {
	fou.mutation.AddDeletedAt(u)
	return fou
}

// SetOrderID sets the "order_id" field.
func (fou *FeeOrderUpdate) SetOrderID(u uuid.UUID) *FeeOrderUpdate {
	fou.mutation.SetOrderID(u)
	return fou
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (fou *FeeOrderUpdate) SetNillableOrderID(u *uuid.UUID) *FeeOrderUpdate {
	if u != nil {
		fou.SetOrderID(*u)
	}
	return fou
}

// ClearOrderID clears the value of the "order_id" field.
func (fou *FeeOrderUpdate) ClearOrderID() *FeeOrderUpdate {
	fou.mutation.ClearOrderID()
	return fou
}

// SetGoodValueUsd sets the "good_value_usd" field.
func (fou *FeeOrderUpdate) SetGoodValueUsd(d decimal.Decimal) *FeeOrderUpdate {
	fou.mutation.SetGoodValueUsd(d)
	return fou
}

// SetNillableGoodValueUsd sets the "good_value_usd" field if the given value is not nil.
func (fou *FeeOrderUpdate) SetNillableGoodValueUsd(d *decimal.Decimal) *FeeOrderUpdate {
	if d != nil {
		fou.SetGoodValueUsd(*d)
	}
	return fou
}

// ClearGoodValueUsd clears the value of the "good_value_usd" field.
func (fou *FeeOrderUpdate) ClearGoodValueUsd() *FeeOrderUpdate {
	fou.mutation.ClearGoodValueUsd()
	return fou
}

// SetPaymentAmountUsd sets the "payment_amount_usd" field.
func (fou *FeeOrderUpdate) SetPaymentAmountUsd(d decimal.Decimal) *FeeOrderUpdate {
	fou.mutation.SetPaymentAmountUsd(d)
	return fou
}

// SetNillablePaymentAmountUsd sets the "payment_amount_usd" field if the given value is not nil.
func (fou *FeeOrderUpdate) SetNillablePaymentAmountUsd(d *decimal.Decimal) *FeeOrderUpdate {
	if d != nil {
		fou.SetPaymentAmountUsd(*d)
	}
	return fou
}

// ClearPaymentAmountUsd clears the value of the "payment_amount_usd" field.
func (fou *FeeOrderUpdate) ClearPaymentAmountUsd() *FeeOrderUpdate {
	fou.mutation.ClearPaymentAmountUsd()
	return fou
}

// SetDiscountAmountUsd sets the "discount_amount_usd" field.
func (fou *FeeOrderUpdate) SetDiscountAmountUsd(d decimal.Decimal) *FeeOrderUpdate {
	fou.mutation.SetDiscountAmountUsd(d)
	return fou
}

// SetNillableDiscountAmountUsd sets the "discount_amount_usd" field if the given value is not nil.
func (fou *FeeOrderUpdate) SetNillableDiscountAmountUsd(d *decimal.Decimal) *FeeOrderUpdate {
	if d != nil {
		fou.SetDiscountAmountUsd(*d)
	}
	return fou
}

// ClearDiscountAmountUsd clears the value of the "discount_amount_usd" field.
func (fou *FeeOrderUpdate) ClearDiscountAmountUsd() *FeeOrderUpdate {
	fou.mutation.ClearDiscountAmountUsd()
	return fou
}

// SetPromotionID sets the "promotion_id" field.
func (fou *FeeOrderUpdate) SetPromotionID(u uuid.UUID) *FeeOrderUpdate {
	fou.mutation.SetPromotionID(u)
	return fou
}

// SetNillablePromotionID sets the "promotion_id" field if the given value is not nil.
func (fou *FeeOrderUpdate) SetNillablePromotionID(u *uuid.UUID) *FeeOrderUpdate {
	if u != nil {
		fou.SetPromotionID(*u)
	}
	return fou
}

// ClearPromotionID clears the value of the "promotion_id" field.
func (fou *FeeOrderUpdate) ClearPromotionID() *FeeOrderUpdate {
	fou.mutation.ClearPromotionID()
	return fou
}

// SetDurationSeconds sets the "duration_seconds" field.
func (fou *FeeOrderUpdate) SetDurationSeconds(u uint32) *FeeOrderUpdate {
	fou.mutation.ResetDurationSeconds()
	fou.mutation.SetDurationSeconds(u)
	return fou
}

// SetNillableDurationSeconds sets the "duration_seconds" field if the given value is not nil.
func (fou *FeeOrderUpdate) SetNillableDurationSeconds(u *uint32) *FeeOrderUpdate {
	if u != nil {
		fou.SetDurationSeconds(*u)
	}
	return fou
}

// AddDurationSeconds adds u to the "duration_seconds" field.
func (fou *FeeOrderUpdate) AddDurationSeconds(u int32) *FeeOrderUpdate {
	fou.mutation.AddDurationSeconds(u)
	return fou
}

// ClearDurationSeconds clears the value of the "duration_seconds" field.
func (fou *FeeOrderUpdate) ClearDurationSeconds() *FeeOrderUpdate {
	fou.mutation.ClearDurationSeconds()
	return fou
}

// Mutation returns the FeeOrderMutation object of the builder.
func (fou *FeeOrderUpdate) Mutation() *FeeOrderMutation {
	return fou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fou *FeeOrderUpdate) Save(ctx context.Context) (int, error) {
	fou.defaults()
	return withHooks(ctx, fou.sqlSave, fou.mutation, fou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fou *FeeOrderUpdate) SaveX(ctx context.Context) int {
	affected, err := fou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fou *FeeOrderUpdate) Exec(ctx context.Context) error {
	_, err := fou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fou *FeeOrderUpdate) ExecX(ctx context.Context) {
	if err := fou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fou *FeeOrderUpdate) defaults() {
	if _, ok := fou.mutation.UpdatedAt(); !ok {
		v := feeorder.UpdateDefaultUpdatedAt()
		fou.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (fou *FeeOrderUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FeeOrderUpdate {
	fou.modifiers = append(fou.modifiers, modifiers...)
	return fou
}

func (fou *FeeOrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(feeorder.Table, feeorder.Columns, sqlgraph.NewFieldSpec(feeorder.FieldID, field.TypeUint32))
	if ps := fou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fou.mutation.EntID(); ok {
		_spec.SetField(feeorder.FieldEntID, field.TypeUUID, value)
	}
	if value, ok := fou.mutation.CreatedAt(); ok {
		_spec.SetField(feeorder.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := fou.mutation.AddedCreatedAt(); ok {
		_spec.AddField(feeorder.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := fou.mutation.UpdatedAt(); ok {
		_spec.SetField(feeorder.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := fou.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(feeorder.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := fou.mutation.DeletedAt(); ok {
		_spec.SetField(feeorder.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := fou.mutation.AddedDeletedAt(); ok {
		_spec.AddField(feeorder.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := fou.mutation.OrderID(); ok {
		_spec.SetField(feeorder.FieldOrderID, field.TypeUUID, value)
	}
	if fou.mutation.OrderIDCleared() {
		_spec.ClearField(feeorder.FieldOrderID, field.TypeUUID)
	}
	if value, ok := fou.mutation.GoodValueUsd(); ok {
		_spec.SetField(feeorder.FieldGoodValueUsd, field.TypeOther, value)
	}
	if fou.mutation.GoodValueUsdCleared() {
		_spec.ClearField(feeorder.FieldGoodValueUsd, field.TypeOther)
	}
	if value, ok := fou.mutation.PaymentAmountUsd(); ok {
		_spec.SetField(feeorder.FieldPaymentAmountUsd, field.TypeOther, value)
	}
	if fou.mutation.PaymentAmountUsdCleared() {
		_spec.ClearField(feeorder.FieldPaymentAmountUsd, field.TypeOther)
	}
	if value, ok := fou.mutation.DiscountAmountUsd(); ok {
		_spec.SetField(feeorder.FieldDiscountAmountUsd, field.TypeOther, value)
	}
	if fou.mutation.DiscountAmountUsdCleared() {
		_spec.ClearField(feeorder.FieldDiscountAmountUsd, field.TypeOther)
	}
	if value, ok := fou.mutation.PromotionID(); ok {
		_spec.SetField(feeorder.FieldPromotionID, field.TypeUUID, value)
	}
	if fou.mutation.PromotionIDCleared() {
		_spec.ClearField(feeorder.FieldPromotionID, field.TypeUUID)
	}
	if value, ok := fou.mutation.DurationSeconds(); ok {
		_spec.SetField(feeorder.FieldDurationSeconds, field.TypeUint32, value)
	}
	if value, ok := fou.mutation.AddedDurationSeconds(); ok {
		_spec.AddField(feeorder.FieldDurationSeconds, field.TypeUint32, value)
	}
	if fou.mutation.DurationSecondsCleared() {
		_spec.ClearField(feeorder.FieldDurationSeconds, field.TypeUint32)
	}
	_spec.AddModifiers(fou.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, fou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{feeorder.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fou.mutation.done = true
	return n, nil
}

// FeeOrderUpdateOne is the builder for updating a single FeeOrder entity.
type FeeOrderUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *FeeOrderMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetEntID sets the "ent_id" field.
func (fouo *FeeOrderUpdateOne) SetEntID(u uuid.UUID) *FeeOrderUpdateOne {
	fouo.mutation.SetEntID(u)
	return fouo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (fouo *FeeOrderUpdateOne) SetNillableEntID(u *uuid.UUID) *FeeOrderUpdateOne {
	if u != nil {
		fouo.SetEntID(*u)
	}
	return fouo
}

// SetCreatedAt sets the "created_at" field.
func (fouo *FeeOrderUpdateOne) SetCreatedAt(u uint32) *FeeOrderUpdateOne {
	fouo.mutation.ResetCreatedAt()
	fouo.mutation.SetCreatedAt(u)
	return fouo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fouo *FeeOrderUpdateOne) SetNillableCreatedAt(u *uint32) *FeeOrderUpdateOne {
	if u != nil {
		fouo.SetCreatedAt(*u)
	}
	return fouo
}

// AddCreatedAt adds u to the "created_at" field.
func (fouo *FeeOrderUpdateOne) AddCreatedAt(u int32) *FeeOrderUpdateOne {
	fouo.mutation.AddCreatedAt(u)
	return fouo
}

// SetUpdatedAt sets the "updated_at" field.
func (fouo *FeeOrderUpdateOne) SetUpdatedAt(u uint32) *FeeOrderUpdateOne {
	fouo.mutation.ResetUpdatedAt()
	fouo.mutation.SetUpdatedAt(u)
	return fouo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (fouo *FeeOrderUpdateOne) AddUpdatedAt(u int32) *FeeOrderUpdateOne {
	fouo.mutation.AddUpdatedAt(u)
	return fouo
}

// SetDeletedAt sets the "deleted_at" field.
func (fouo *FeeOrderUpdateOne) SetDeletedAt(u uint32) *FeeOrderUpdateOne {
	fouo.mutation.ResetDeletedAt()
	fouo.mutation.SetDeletedAt(u)
	return fouo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fouo *FeeOrderUpdateOne) SetNillableDeletedAt(u *uint32) *FeeOrderUpdateOne {
	if u != nil {
		fouo.SetDeletedAt(*u)
	}
	return fouo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (fouo *FeeOrderUpdateOne) AddDeletedAt(u int32) *FeeOrderUpdateOne {
	fouo.mutation.AddDeletedAt(u)
	return fouo
}

// SetOrderID sets the "order_id" field.
func (fouo *FeeOrderUpdateOne) SetOrderID(u uuid.UUID) *FeeOrderUpdateOne {
	fouo.mutation.SetOrderID(u)
	return fouo
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (fouo *FeeOrderUpdateOne) SetNillableOrderID(u *uuid.UUID) *FeeOrderUpdateOne {
	if u != nil {
		fouo.SetOrderID(*u)
	}
	return fouo
}

// ClearOrderID clears the value of the "order_id" field.
func (fouo *FeeOrderUpdateOne) ClearOrderID() *FeeOrderUpdateOne {
	fouo.mutation.ClearOrderID()
	return fouo
}

// SetGoodValueUsd sets the "good_value_usd" field.
func (fouo *FeeOrderUpdateOne) SetGoodValueUsd(d decimal.Decimal) *FeeOrderUpdateOne {
	fouo.mutation.SetGoodValueUsd(d)
	return fouo
}

// SetNillableGoodValueUsd sets the "good_value_usd" field if the given value is not nil.
func (fouo *FeeOrderUpdateOne) SetNillableGoodValueUsd(d *decimal.Decimal) *FeeOrderUpdateOne {
	if d != nil {
		fouo.SetGoodValueUsd(*d)
	}
	return fouo
}

// ClearGoodValueUsd clears the value of the "good_value_usd" field.
func (fouo *FeeOrderUpdateOne) ClearGoodValueUsd() *FeeOrderUpdateOne {
	fouo.mutation.ClearGoodValueUsd()
	return fouo
}

// SetPaymentAmountUsd sets the "payment_amount_usd" field.
func (fouo *FeeOrderUpdateOne) SetPaymentAmountUsd(d decimal.Decimal) *FeeOrderUpdateOne {
	fouo.mutation.SetPaymentAmountUsd(d)
	return fouo
}

// SetNillablePaymentAmountUsd sets the "payment_amount_usd" field if the given value is not nil.
func (fouo *FeeOrderUpdateOne) SetNillablePaymentAmountUsd(d *decimal.Decimal) *FeeOrderUpdateOne {
	if d != nil {
		fouo.SetPaymentAmountUsd(*d)
	}
	return fouo
}

// ClearPaymentAmountUsd clears the value of the "payment_amount_usd" field.
func (fouo *FeeOrderUpdateOne) ClearPaymentAmountUsd() *FeeOrderUpdateOne {
	fouo.mutation.ClearPaymentAmountUsd()
	return fouo
}

// SetDiscountAmountUsd sets the "discount_amount_usd" field.
func (fouo *FeeOrderUpdateOne) SetDiscountAmountUsd(d decimal.Decimal) *FeeOrderUpdateOne {
	fouo.mutation.SetDiscountAmountUsd(d)
	return fouo
}

// SetNillableDiscountAmountUsd sets the "discount_amount_usd" field if the given value is not nil.
func (fouo *FeeOrderUpdateOne) SetNillableDiscountAmountUsd(d *decimal.Decimal) *FeeOrderUpdateOne {
	if d != nil {
		fouo.SetDiscountAmountUsd(*d)
	}
	return fouo
}

// ClearDiscountAmountUsd clears the value of the "discount_amount_usd" field.
func (fouo *FeeOrderUpdateOne) ClearDiscountAmountUsd() *FeeOrderUpdateOne {
	fouo.mutation.ClearDiscountAmountUsd()
	return fouo
}

// SetPromotionID sets the "promotion_id" field.
func (fouo *FeeOrderUpdateOne) SetPromotionID(u uuid.UUID) *FeeOrderUpdateOne {
	fouo.mutation.SetPromotionID(u)
	return fouo
}

// SetNillablePromotionID sets the "promotion_id" field if the given value is not nil.
func (fouo *FeeOrderUpdateOne) SetNillablePromotionID(u *uuid.UUID) *FeeOrderUpdateOne {
	if u != nil {
		fouo.SetPromotionID(*u)
	}
	return fouo
}

// ClearPromotionID clears the value of the "promotion_id" field.
func (fouo *FeeOrderUpdateOne) ClearPromotionID() *FeeOrderUpdateOne {
	fouo.mutation.ClearPromotionID()
	return fouo
}

// SetDurationSeconds sets the "duration_seconds" field.
func (fouo *FeeOrderUpdateOne) SetDurationSeconds(u uint32) *FeeOrderUpdateOne {
	fouo.mutation.ResetDurationSeconds()
	fouo.mutation.SetDurationSeconds(u)
	return fouo
}

// SetNillableDurationSeconds sets the "duration_seconds" field if the given value is not nil.
func (fouo *FeeOrderUpdateOne) SetNillableDurationSeconds(u *uint32) *FeeOrderUpdateOne {
	if u != nil {
		fouo.SetDurationSeconds(*u)
	}
	return fouo
}

// AddDurationSeconds adds u to the "duration_seconds" field.
func (fouo *FeeOrderUpdateOne) AddDurationSeconds(u int32) *FeeOrderUpdateOne {
	fouo.mutation.AddDurationSeconds(u)
	return fouo
}

// ClearDurationSeconds clears the value of the "duration_seconds" field.
func (fouo *FeeOrderUpdateOne) ClearDurationSeconds() *FeeOrderUpdateOne {
	fouo.mutation.ClearDurationSeconds()
	return fouo
}

// Mutation returns the FeeOrderMutation object of the builder.
func (fouo *FeeOrderUpdateOne) Mutation() *FeeOrderMutation {
	return fouo.mutation
}

// Where appends a list predicates to the FeeOrderUpdate builder.
func (fouo *FeeOrderUpdateOne) Where(ps ...predicate.FeeOrder) *FeeOrderUpdateOne {
	fouo.mutation.Where(ps...)
	return fouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fouo *FeeOrderUpdateOne) Select(field string, fields ...string) *FeeOrderUpdateOne {
	fouo.fields = append([]string{field}, fields...)
	return fouo
}

// Save executes the query and returns the updated FeeOrder entity.
func (fouo *FeeOrderUpdateOne) Save(ctx context.Context) (*FeeOrder, error) {
	fouo.defaults()
	return withHooks(ctx, fouo.sqlSave, fouo.mutation, fouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fouo *FeeOrderUpdateOne) SaveX(ctx context.Context) *FeeOrder {
	node, err := fouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fouo *FeeOrderUpdateOne) Exec(ctx context.Context) error {
	_, err := fouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fouo *FeeOrderUpdateOne) ExecX(ctx context.Context) {
	if err := fouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fouo *FeeOrderUpdateOne) defaults() {
	if _, ok := fouo.mutation.UpdatedAt(); !ok {
		v := feeorder.UpdateDefaultUpdatedAt()
		fouo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (fouo *FeeOrderUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FeeOrderUpdateOne {
	fouo.modifiers = append(fouo.modifiers, modifiers...)
	return fouo
}

func (fouo *FeeOrderUpdateOne) sqlSave(ctx context.Context) (_node *FeeOrder, err error) {
	_spec := sqlgraph.NewUpdateSpec(feeorder.Table, feeorder.Columns, sqlgraph.NewFieldSpec(feeorder.FieldID, field.TypeUint32))
	id, ok := fouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "FeeOrder.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, feeorder.FieldID)
		for _, f := range fields {
			if !feeorder.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != feeorder.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fouo.mutation.EntID(); ok {
		_spec.SetField(feeorder.FieldEntID, field.TypeUUID, value)
	}
	if value, ok := fouo.mutation.CreatedAt(); ok {
		_spec.SetField(feeorder.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := fouo.mutation.AddedCreatedAt(); ok {
		_spec.AddField(feeorder.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := fouo.mutation.UpdatedAt(); ok {
		_spec.SetField(feeorder.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := fouo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(feeorder.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := fouo.mutation.DeletedAt(); ok {
		_spec.SetField(feeorder.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := fouo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(feeorder.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := fouo.mutation.OrderID(); ok {
		_spec.SetField(feeorder.FieldOrderID, field.TypeUUID, value)
	}
	if fouo.mutation.OrderIDCleared() {
		_spec.ClearField(feeorder.FieldOrderID, field.TypeUUID)
	}
	if value, ok := fouo.mutation.GoodValueUsd(); ok {
		_spec.SetField(feeorder.FieldGoodValueUsd, field.TypeOther, value)
	}
	if fouo.mutation.GoodValueUsdCleared() {
		_spec.ClearField(feeorder.FieldGoodValueUsd, field.TypeOther)
	}
	if value, ok := fouo.mutation.PaymentAmountUsd(); ok {
		_spec.SetField(feeorder.FieldPaymentAmountUsd, field.TypeOther, value)
	}
	if fouo.mutation.PaymentAmountUsdCleared() {
		_spec.ClearField(feeorder.FieldPaymentAmountUsd, field.TypeOther)
	}
	if value, ok := fouo.mutation.DiscountAmountUsd(); ok {
		_spec.SetField(feeorder.FieldDiscountAmountUsd, field.TypeOther, value)
	}
	if fouo.mutation.DiscountAmountUsdCleared() {
		_spec.ClearField(feeorder.FieldDiscountAmountUsd, field.TypeOther)
	}
	if value, ok := fouo.mutation.PromotionID(); ok {
		_spec.SetField(feeorder.FieldPromotionID, field.TypeUUID, value)
	}
	if fouo.mutation.PromotionIDCleared() {
		_spec.ClearField(feeorder.FieldPromotionID, field.TypeUUID)
	}
	if value, ok := fouo.mutation.DurationSeconds(); ok {
		_spec.SetField(feeorder.FieldDurationSeconds, field.TypeUint32, value)
	}
	if value, ok := fouo.mutation.AddedDurationSeconds(); ok {
		_spec.AddField(feeorder.FieldDurationSeconds, field.TypeUint32, value)
	}
	if fouo.mutation.DurationSecondsCleared() {
		_spec.ClearField(feeorder.FieldDurationSeconds, field.TypeUint32)
	}
	_spec.AddModifiers(fouo.modifiers...)
	_node = &FeeOrder{config: fouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{feeorder.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fouo.mutation.done = true
	return _node, nil
}
