// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderbase"
	"github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/predicate"
	"github.com/google/uuid"
)

// OrderBaseUpdate is the builder for updating OrderBase entities.
type OrderBaseUpdate struct {
	config
	hooks     []Hook
	mutation  *OrderBaseMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the OrderBaseUpdate builder.
func (obu *OrderBaseUpdate) Where(ps ...predicate.OrderBase) *OrderBaseUpdate {
	obu.mutation.Where(ps...)
	return obu
}

// SetEntID sets the "ent_id" field.
func (obu *OrderBaseUpdate) SetEntID(u uuid.UUID) *OrderBaseUpdate {
	obu.mutation.SetEntID(u)
	return obu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (obu *OrderBaseUpdate) SetNillableEntID(u *uuid.UUID) *OrderBaseUpdate {
	if u != nil {
		obu.SetEntID(*u)
	}
	return obu
}

// SetCreatedAt sets the "created_at" field.
func (obu *OrderBaseUpdate) SetCreatedAt(u uint32) *OrderBaseUpdate {
	obu.mutation.ResetCreatedAt()
	obu.mutation.SetCreatedAt(u)
	return obu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (obu *OrderBaseUpdate) SetNillableCreatedAt(u *uint32) *OrderBaseUpdate {
	if u != nil {
		obu.SetCreatedAt(*u)
	}
	return obu
}

// AddCreatedAt adds u to the "created_at" field.
func (obu *OrderBaseUpdate) AddCreatedAt(u int32) *OrderBaseUpdate {
	obu.mutation.AddCreatedAt(u)
	return obu
}

// SetUpdatedAt sets the "updated_at" field.
func (obu *OrderBaseUpdate) SetUpdatedAt(u uint32) *OrderBaseUpdate {
	obu.mutation.ResetUpdatedAt()
	obu.mutation.SetUpdatedAt(u)
	return obu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (obu *OrderBaseUpdate) AddUpdatedAt(u int32) *OrderBaseUpdate {
	obu.mutation.AddUpdatedAt(u)
	return obu
}

// SetDeletedAt sets the "deleted_at" field.
func (obu *OrderBaseUpdate) SetDeletedAt(u uint32) *OrderBaseUpdate {
	obu.mutation.ResetDeletedAt()
	obu.mutation.SetDeletedAt(u)
	return obu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (obu *OrderBaseUpdate) SetNillableDeletedAt(u *uint32) *OrderBaseUpdate {
	if u != nil {
		obu.SetDeletedAt(*u)
	}
	return obu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (obu *OrderBaseUpdate) AddDeletedAt(u int32) *OrderBaseUpdate {
	obu.mutation.AddDeletedAt(u)
	return obu
}

// SetAppID sets the "app_id" field.
func (obu *OrderBaseUpdate) SetAppID(u uuid.UUID) *OrderBaseUpdate {
	obu.mutation.SetAppID(u)
	return obu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (obu *OrderBaseUpdate) SetNillableAppID(u *uuid.UUID) *OrderBaseUpdate {
	if u != nil {
		obu.SetAppID(*u)
	}
	return obu
}

// ClearAppID clears the value of the "app_id" field.
func (obu *OrderBaseUpdate) ClearAppID() *OrderBaseUpdate {
	obu.mutation.ClearAppID()
	return obu
}

// SetUserID sets the "user_id" field.
func (obu *OrderBaseUpdate) SetUserID(u uuid.UUID) *OrderBaseUpdate {
	obu.mutation.SetUserID(u)
	return obu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (obu *OrderBaseUpdate) SetNillableUserID(u *uuid.UUID) *OrderBaseUpdate {
	if u != nil {
		obu.SetUserID(*u)
	}
	return obu
}

// ClearUserID clears the value of the "user_id" field.
func (obu *OrderBaseUpdate) ClearUserID() *OrderBaseUpdate {
	obu.mutation.ClearUserID()
	return obu
}

// SetGoodID sets the "good_id" field.
func (obu *OrderBaseUpdate) SetGoodID(u uuid.UUID) *OrderBaseUpdate {
	obu.mutation.SetGoodID(u)
	return obu
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (obu *OrderBaseUpdate) SetNillableGoodID(u *uuid.UUID) *OrderBaseUpdate {
	if u != nil {
		obu.SetGoodID(*u)
	}
	return obu
}

// ClearGoodID clears the value of the "good_id" field.
func (obu *OrderBaseUpdate) ClearGoodID() *OrderBaseUpdate {
	obu.mutation.ClearGoodID()
	return obu
}

// SetAppGoodID sets the "app_good_id" field.
func (obu *OrderBaseUpdate) SetAppGoodID(u uuid.UUID) *OrderBaseUpdate {
	obu.mutation.SetAppGoodID(u)
	return obu
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (obu *OrderBaseUpdate) SetNillableAppGoodID(u *uuid.UUID) *OrderBaseUpdate {
	if u != nil {
		obu.SetAppGoodID(*u)
	}
	return obu
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (obu *OrderBaseUpdate) ClearAppGoodID() *OrderBaseUpdate {
	obu.mutation.ClearAppGoodID()
	return obu
}

// SetGoodType sets the "good_type" field.
func (obu *OrderBaseUpdate) SetGoodType(s string) *OrderBaseUpdate {
	obu.mutation.SetGoodType(s)
	return obu
}

// SetNillableGoodType sets the "good_type" field if the given value is not nil.
func (obu *OrderBaseUpdate) SetNillableGoodType(s *string) *OrderBaseUpdate {
	if s != nil {
		obu.SetGoodType(*s)
	}
	return obu
}

// ClearGoodType clears the value of the "good_type" field.
func (obu *OrderBaseUpdate) ClearGoodType() *OrderBaseUpdate {
	obu.mutation.ClearGoodType()
	return obu
}

// SetParentOrderID sets the "parent_order_id" field.
func (obu *OrderBaseUpdate) SetParentOrderID(u uuid.UUID) *OrderBaseUpdate {
	obu.mutation.SetParentOrderID(u)
	return obu
}

// SetNillableParentOrderID sets the "parent_order_id" field if the given value is not nil.
func (obu *OrderBaseUpdate) SetNillableParentOrderID(u *uuid.UUID) *OrderBaseUpdate {
	if u != nil {
		obu.SetParentOrderID(*u)
	}
	return obu
}

// ClearParentOrderID clears the value of the "parent_order_id" field.
func (obu *OrderBaseUpdate) ClearParentOrderID() *OrderBaseUpdate {
	obu.mutation.ClearParentOrderID()
	return obu
}

// SetOrderType sets the "order_type" field.
func (obu *OrderBaseUpdate) SetOrderType(s string) *OrderBaseUpdate {
	obu.mutation.SetOrderType(s)
	return obu
}

// SetNillableOrderType sets the "order_type" field if the given value is not nil.
func (obu *OrderBaseUpdate) SetNillableOrderType(s *string) *OrderBaseUpdate {
	if s != nil {
		obu.SetOrderType(*s)
	}
	return obu
}

// ClearOrderType clears the value of the "order_type" field.
func (obu *OrderBaseUpdate) ClearOrderType() *OrderBaseUpdate {
	obu.mutation.ClearOrderType()
	return obu
}

// SetCreateMethod sets the "create_method" field.
func (obu *OrderBaseUpdate) SetCreateMethod(s string) *OrderBaseUpdate {
	obu.mutation.SetCreateMethod(s)
	return obu
}

// SetNillableCreateMethod sets the "create_method" field if the given value is not nil.
func (obu *OrderBaseUpdate) SetNillableCreateMethod(s *string) *OrderBaseUpdate {
	if s != nil {
		obu.SetCreateMethod(*s)
	}
	return obu
}

// ClearCreateMethod clears the value of the "create_method" field.
func (obu *OrderBaseUpdate) ClearCreateMethod() *OrderBaseUpdate {
	obu.mutation.ClearCreateMethod()
	return obu
}

// SetSimulate sets the "simulate" field.
func (obu *OrderBaseUpdate) SetSimulate(b bool) *OrderBaseUpdate {
	obu.mutation.SetSimulate(b)
	return obu
}

// SetNillableSimulate sets the "simulate" field if the given value is not nil.
func (obu *OrderBaseUpdate) SetNillableSimulate(b *bool) *OrderBaseUpdate {
	if b != nil {
		obu.SetSimulate(*b)
	}
	return obu
}

// ClearSimulate clears the value of the "simulate" field.
func (obu *OrderBaseUpdate) ClearSimulate() *OrderBaseUpdate {
	obu.mutation.ClearSimulate()
	return obu
}

// Mutation returns the OrderBaseMutation object of the builder.
func (obu *OrderBaseUpdate) Mutation() *OrderBaseMutation {
	return obu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (obu *OrderBaseUpdate) Save(ctx context.Context) (int, error) {
	obu.defaults()
	return withHooks(ctx, obu.sqlSave, obu.mutation, obu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (obu *OrderBaseUpdate) SaveX(ctx context.Context) int {
	affected, err := obu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (obu *OrderBaseUpdate) Exec(ctx context.Context) error {
	_, err := obu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (obu *OrderBaseUpdate) ExecX(ctx context.Context) {
	if err := obu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (obu *OrderBaseUpdate) defaults() {
	if _, ok := obu.mutation.UpdatedAt(); !ok {
		v := orderbase.UpdateDefaultUpdatedAt()
		obu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (obu *OrderBaseUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderBaseUpdate {
	obu.modifiers = append(obu.modifiers, modifiers...)
	return obu
}

func (obu *OrderBaseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(orderbase.Table, orderbase.Columns, sqlgraph.NewFieldSpec(orderbase.FieldID, field.TypeUint32))
	if ps := obu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := obu.mutation.EntID(); ok {
		_spec.SetField(orderbase.FieldEntID, field.TypeUUID, value)
	}
	if value, ok := obu.mutation.CreatedAt(); ok {
		_spec.SetField(orderbase.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := obu.mutation.AddedCreatedAt(); ok {
		_spec.AddField(orderbase.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := obu.mutation.UpdatedAt(); ok {
		_spec.SetField(orderbase.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := obu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(orderbase.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := obu.mutation.DeletedAt(); ok {
		_spec.SetField(orderbase.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := obu.mutation.AddedDeletedAt(); ok {
		_spec.AddField(orderbase.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := obu.mutation.AppID(); ok {
		_spec.SetField(orderbase.FieldAppID, field.TypeUUID, value)
	}
	if obu.mutation.AppIDCleared() {
		_spec.ClearField(orderbase.FieldAppID, field.TypeUUID)
	}
	if value, ok := obu.mutation.UserID(); ok {
		_spec.SetField(orderbase.FieldUserID, field.TypeUUID, value)
	}
	if obu.mutation.UserIDCleared() {
		_spec.ClearField(orderbase.FieldUserID, field.TypeUUID)
	}
	if value, ok := obu.mutation.GoodID(); ok {
		_spec.SetField(orderbase.FieldGoodID, field.TypeUUID, value)
	}
	if obu.mutation.GoodIDCleared() {
		_spec.ClearField(orderbase.FieldGoodID, field.TypeUUID)
	}
	if value, ok := obu.mutation.AppGoodID(); ok {
		_spec.SetField(orderbase.FieldAppGoodID, field.TypeUUID, value)
	}
	if obu.mutation.AppGoodIDCleared() {
		_spec.ClearField(orderbase.FieldAppGoodID, field.TypeUUID)
	}
	if value, ok := obu.mutation.GoodType(); ok {
		_spec.SetField(orderbase.FieldGoodType, field.TypeString, value)
	}
	if obu.mutation.GoodTypeCleared() {
		_spec.ClearField(orderbase.FieldGoodType, field.TypeString)
	}
	if value, ok := obu.mutation.ParentOrderID(); ok {
		_spec.SetField(orderbase.FieldParentOrderID, field.TypeUUID, value)
	}
	if obu.mutation.ParentOrderIDCleared() {
		_spec.ClearField(orderbase.FieldParentOrderID, field.TypeUUID)
	}
	if value, ok := obu.mutation.OrderType(); ok {
		_spec.SetField(orderbase.FieldOrderType, field.TypeString, value)
	}
	if obu.mutation.OrderTypeCleared() {
		_spec.ClearField(orderbase.FieldOrderType, field.TypeString)
	}
	if value, ok := obu.mutation.CreateMethod(); ok {
		_spec.SetField(orderbase.FieldCreateMethod, field.TypeString, value)
	}
	if obu.mutation.CreateMethodCleared() {
		_spec.ClearField(orderbase.FieldCreateMethod, field.TypeString)
	}
	if value, ok := obu.mutation.Simulate(); ok {
		_spec.SetField(orderbase.FieldSimulate, field.TypeBool, value)
	}
	if obu.mutation.SimulateCleared() {
		_spec.ClearField(orderbase.FieldSimulate, field.TypeBool)
	}
	_spec.AddModifiers(obu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, obu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderbase.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	obu.mutation.done = true
	return n, nil
}

// OrderBaseUpdateOne is the builder for updating a single OrderBase entity.
type OrderBaseUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *OrderBaseMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetEntID sets the "ent_id" field.
func (obuo *OrderBaseUpdateOne) SetEntID(u uuid.UUID) *OrderBaseUpdateOne {
	obuo.mutation.SetEntID(u)
	return obuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (obuo *OrderBaseUpdateOne) SetNillableEntID(u *uuid.UUID) *OrderBaseUpdateOne {
	if u != nil {
		obuo.SetEntID(*u)
	}
	return obuo
}

// SetCreatedAt sets the "created_at" field.
func (obuo *OrderBaseUpdateOne) SetCreatedAt(u uint32) *OrderBaseUpdateOne {
	obuo.mutation.ResetCreatedAt()
	obuo.mutation.SetCreatedAt(u)
	return obuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (obuo *OrderBaseUpdateOne) SetNillableCreatedAt(u *uint32) *OrderBaseUpdateOne {
	if u != nil {
		obuo.SetCreatedAt(*u)
	}
	return obuo
}

// AddCreatedAt adds u to the "created_at" field.
func (obuo *OrderBaseUpdateOne) AddCreatedAt(u int32) *OrderBaseUpdateOne {
	obuo.mutation.AddCreatedAt(u)
	return obuo
}

// SetUpdatedAt sets the "updated_at" field.
func (obuo *OrderBaseUpdateOne) SetUpdatedAt(u uint32) *OrderBaseUpdateOne {
	obuo.mutation.ResetUpdatedAt()
	obuo.mutation.SetUpdatedAt(u)
	return obuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (obuo *OrderBaseUpdateOne) AddUpdatedAt(u int32) *OrderBaseUpdateOne {
	obuo.mutation.AddUpdatedAt(u)
	return obuo
}

// SetDeletedAt sets the "deleted_at" field.
func (obuo *OrderBaseUpdateOne) SetDeletedAt(u uint32) *OrderBaseUpdateOne {
	obuo.mutation.ResetDeletedAt()
	obuo.mutation.SetDeletedAt(u)
	return obuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (obuo *OrderBaseUpdateOne) SetNillableDeletedAt(u *uint32) *OrderBaseUpdateOne {
	if u != nil {
		obuo.SetDeletedAt(*u)
	}
	return obuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (obuo *OrderBaseUpdateOne) AddDeletedAt(u int32) *OrderBaseUpdateOne {
	obuo.mutation.AddDeletedAt(u)
	return obuo
}

// SetAppID sets the "app_id" field.
func (obuo *OrderBaseUpdateOne) SetAppID(u uuid.UUID) *OrderBaseUpdateOne {
	obuo.mutation.SetAppID(u)
	return obuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (obuo *OrderBaseUpdateOne) SetNillableAppID(u *uuid.UUID) *OrderBaseUpdateOne {
	if u != nil {
		obuo.SetAppID(*u)
	}
	return obuo
}

// ClearAppID clears the value of the "app_id" field.
func (obuo *OrderBaseUpdateOne) ClearAppID() *OrderBaseUpdateOne {
	obuo.mutation.ClearAppID()
	return obuo
}

// SetUserID sets the "user_id" field.
func (obuo *OrderBaseUpdateOne) SetUserID(u uuid.UUID) *OrderBaseUpdateOne {
	obuo.mutation.SetUserID(u)
	return obuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (obuo *OrderBaseUpdateOne) SetNillableUserID(u *uuid.UUID) *OrderBaseUpdateOne {
	if u != nil {
		obuo.SetUserID(*u)
	}
	return obuo
}

// ClearUserID clears the value of the "user_id" field.
func (obuo *OrderBaseUpdateOne) ClearUserID() *OrderBaseUpdateOne {
	obuo.mutation.ClearUserID()
	return obuo
}

// SetGoodID sets the "good_id" field.
func (obuo *OrderBaseUpdateOne) SetGoodID(u uuid.UUID) *OrderBaseUpdateOne {
	obuo.mutation.SetGoodID(u)
	return obuo
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (obuo *OrderBaseUpdateOne) SetNillableGoodID(u *uuid.UUID) *OrderBaseUpdateOne {
	if u != nil {
		obuo.SetGoodID(*u)
	}
	return obuo
}

// ClearGoodID clears the value of the "good_id" field.
func (obuo *OrderBaseUpdateOne) ClearGoodID() *OrderBaseUpdateOne {
	obuo.mutation.ClearGoodID()
	return obuo
}

// SetAppGoodID sets the "app_good_id" field.
func (obuo *OrderBaseUpdateOne) SetAppGoodID(u uuid.UUID) *OrderBaseUpdateOne {
	obuo.mutation.SetAppGoodID(u)
	return obuo
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (obuo *OrderBaseUpdateOne) SetNillableAppGoodID(u *uuid.UUID) *OrderBaseUpdateOne {
	if u != nil {
		obuo.SetAppGoodID(*u)
	}
	return obuo
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (obuo *OrderBaseUpdateOne) ClearAppGoodID() *OrderBaseUpdateOne {
	obuo.mutation.ClearAppGoodID()
	return obuo
}

// SetGoodType sets the "good_type" field.
func (obuo *OrderBaseUpdateOne) SetGoodType(s string) *OrderBaseUpdateOne {
	obuo.mutation.SetGoodType(s)
	return obuo
}

// SetNillableGoodType sets the "good_type" field if the given value is not nil.
func (obuo *OrderBaseUpdateOne) SetNillableGoodType(s *string) *OrderBaseUpdateOne {
	if s != nil {
		obuo.SetGoodType(*s)
	}
	return obuo
}

// ClearGoodType clears the value of the "good_type" field.
func (obuo *OrderBaseUpdateOne) ClearGoodType() *OrderBaseUpdateOne {
	obuo.mutation.ClearGoodType()
	return obuo
}

// SetParentOrderID sets the "parent_order_id" field.
func (obuo *OrderBaseUpdateOne) SetParentOrderID(u uuid.UUID) *OrderBaseUpdateOne {
	obuo.mutation.SetParentOrderID(u)
	return obuo
}

// SetNillableParentOrderID sets the "parent_order_id" field if the given value is not nil.
func (obuo *OrderBaseUpdateOne) SetNillableParentOrderID(u *uuid.UUID) *OrderBaseUpdateOne {
	if u != nil {
		obuo.SetParentOrderID(*u)
	}
	return obuo
}

// ClearParentOrderID clears the value of the "parent_order_id" field.
func (obuo *OrderBaseUpdateOne) ClearParentOrderID() *OrderBaseUpdateOne {
	obuo.mutation.ClearParentOrderID()
	return obuo
}

// SetOrderType sets the "order_type" field.
func (obuo *OrderBaseUpdateOne) SetOrderType(s string) *OrderBaseUpdateOne {
	obuo.mutation.SetOrderType(s)
	return obuo
}

// SetNillableOrderType sets the "order_type" field if the given value is not nil.
func (obuo *OrderBaseUpdateOne) SetNillableOrderType(s *string) *OrderBaseUpdateOne {
	if s != nil {
		obuo.SetOrderType(*s)
	}
	return obuo
}

// ClearOrderType clears the value of the "order_type" field.
func (obuo *OrderBaseUpdateOne) ClearOrderType() *OrderBaseUpdateOne {
	obuo.mutation.ClearOrderType()
	return obuo
}

// SetCreateMethod sets the "create_method" field.
func (obuo *OrderBaseUpdateOne) SetCreateMethod(s string) *OrderBaseUpdateOne {
	obuo.mutation.SetCreateMethod(s)
	return obuo
}

// SetNillableCreateMethod sets the "create_method" field if the given value is not nil.
func (obuo *OrderBaseUpdateOne) SetNillableCreateMethod(s *string) *OrderBaseUpdateOne {
	if s != nil {
		obuo.SetCreateMethod(*s)
	}
	return obuo
}

// ClearCreateMethod clears the value of the "create_method" field.
func (obuo *OrderBaseUpdateOne) ClearCreateMethod() *OrderBaseUpdateOne {
	obuo.mutation.ClearCreateMethod()
	return obuo
}

// SetSimulate sets the "simulate" field.
func (obuo *OrderBaseUpdateOne) SetSimulate(b bool) *OrderBaseUpdateOne {
	obuo.mutation.SetSimulate(b)
	return obuo
}

// SetNillableSimulate sets the "simulate" field if the given value is not nil.
func (obuo *OrderBaseUpdateOne) SetNillableSimulate(b *bool) *OrderBaseUpdateOne {
	if b != nil {
		obuo.SetSimulate(*b)
	}
	return obuo
}

// ClearSimulate clears the value of the "simulate" field.
func (obuo *OrderBaseUpdateOne) ClearSimulate() *OrderBaseUpdateOne {
	obuo.mutation.ClearSimulate()
	return obuo
}

// Mutation returns the OrderBaseMutation object of the builder.
func (obuo *OrderBaseUpdateOne) Mutation() *OrderBaseMutation {
	return obuo.mutation
}

// Where appends a list predicates to the OrderBaseUpdate builder.
func (obuo *OrderBaseUpdateOne) Where(ps ...predicate.OrderBase) *OrderBaseUpdateOne {
	obuo.mutation.Where(ps...)
	return obuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (obuo *OrderBaseUpdateOne) Select(field string, fields ...string) *OrderBaseUpdateOne {
	obuo.fields = append([]string{field}, fields...)
	return obuo
}

// Save executes the query and returns the updated OrderBase entity.
func (obuo *OrderBaseUpdateOne) Save(ctx context.Context) (*OrderBase, error) {
	obuo.defaults()
	return withHooks(ctx, obuo.sqlSave, obuo.mutation, obuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (obuo *OrderBaseUpdateOne) SaveX(ctx context.Context) *OrderBase {
	node, err := obuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (obuo *OrderBaseUpdateOne) Exec(ctx context.Context) error {
	_, err := obuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (obuo *OrderBaseUpdateOne) ExecX(ctx context.Context) {
	if err := obuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (obuo *OrderBaseUpdateOne) defaults() {
	if _, ok := obuo.mutation.UpdatedAt(); !ok {
		v := orderbase.UpdateDefaultUpdatedAt()
		obuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (obuo *OrderBaseUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderBaseUpdateOne {
	obuo.modifiers = append(obuo.modifiers, modifiers...)
	return obuo
}

func (obuo *OrderBaseUpdateOne) sqlSave(ctx context.Context) (_node *OrderBase, err error) {
	_spec := sqlgraph.NewUpdateSpec(orderbase.Table, orderbase.Columns, sqlgraph.NewFieldSpec(orderbase.FieldID, field.TypeUint32))
	id, ok := obuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "OrderBase.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := obuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderbase.FieldID)
		for _, f := range fields {
			if !orderbase.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != orderbase.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := obuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := obuo.mutation.EntID(); ok {
		_spec.SetField(orderbase.FieldEntID, field.TypeUUID, value)
	}
	if value, ok := obuo.mutation.CreatedAt(); ok {
		_spec.SetField(orderbase.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := obuo.mutation.AddedCreatedAt(); ok {
		_spec.AddField(orderbase.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := obuo.mutation.UpdatedAt(); ok {
		_spec.SetField(orderbase.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := obuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(orderbase.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := obuo.mutation.DeletedAt(); ok {
		_spec.SetField(orderbase.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := obuo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(orderbase.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := obuo.mutation.AppID(); ok {
		_spec.SetField(orderbase.FieldAppID, field.TypeUUID, value)
	}
	if obuo.mutation.AppIDCleared() {
		_spec.ClearField(orderbase.FieldAppID, field.TypeUUID)
	}
	if value, ok := obuo.mutation.UserID(); ok {
		_spec.SetField(orderbase.FieldUserID, field.TypeUUID, value)
	}
	if obuo.mutation.UserIDCleared() {
		_spec.ClearField(orderbase.FieldUserID, field.TypeUUID)
	}
	if value, ok := obuo.mutation.GoodID(); ok {
		_spec.SetField(orderbase.FieldGoodID, field.TypeUUID, value)
	}
	if obuo.mutation.GoodIDCleared() {
		_spec.ClearField(orderbase.FieldGoodID, field.TypeUUID)
	}
	if value, ok := obuo.mutation.AppGoodID(); ok {
		_spec.SetField(orderbase.FieldAppGoodID, field.TypeUUID, value)
	}
	if obuo.mutation.AppGoodIDCleared() {
		_spec.ClearField(orderbase.FieldAppGoodID, field.TypeUUID)
	}
	if value, ok := obuo.mutation.GoodType(); ok {
		_spec.SetField(orderbase.FieldGoodType, field.TypeString, value)
	}
	if obuo.mutation.GoodTypeCleared() {
		_spec.ClearField(orderbase.FieldGoodType, field.TypeString)
	}
	if value, ok := obuo.mutation.ParentOrderID(); ok {
		_spec.SetField(orderbase.FieldParentOrderID, field.TypeUUID, value)
	}
	if obuo.mutation.ParentOrderIDCleared() {
		_spec.ClearField(orderbase.FieldParentOrderID, field.TypeUUID)
	}
	if value, ok := obuo.mutation.OrderType(); ok {
		_spec.SetField(orderbase.FieldOrderType, field.TypeString, value)
	}
	if obuo.mutation.OrderTypeCleared() {
		_spec.ClearField(orderbase.FieldOrderType, field.TypeString)
	}
	if value, ok := obuo.mutation.CreateMethod(); ok {
		_spec.SetField(orderbase.FieldCreateMethod, field.TypeString, value)
	}
	if obuo.mutation.CreateMethodCleared() {
		_spec.ClearField(orderbase.FieldCreateMethod, field.TypeString)
	}
	if value, ok := obuo.mutation.Simulate(); ok {
		_spec.SetField(orderbase.FieldSimulate, field.TypeBool, value)
	}
	if obuo.mutation.SimulateCleared() {
		_spec.ClearField(orderbase.FieldSimulate, field.TypeBool)
	}
	_spec.AddModifiers(obuo.modifiers...)
	_node = &OrderBase{config: obuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, obuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderbase.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	obuo.mutation.done = true
	return _node, nil
}
