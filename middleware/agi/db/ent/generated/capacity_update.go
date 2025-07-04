// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated/capacity"
	"github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated/predicate"
	"github.com/google/uuid"
)

// CapacityUpdate is the builder for updating Capacity entities.
type CapacityUpdate struct {
	config
	hooks     []Hook
	mutation  *CapacityMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CapacityUpdate builder.
func (cu *CapacityUpdate) Where(ps ...predicate.Capacity) *CapacityUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetEntID sets the "ent_id" field.
func (cu *CapacityUpdate) SetEntID(u uuid.UUID) *CapacityUpdate {
	cu.mutation.SetEntID(u)
	return cu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cu *CapacityUpdate) SetNillableEntID(u *uuid.UUID) *CapacityUpdate {
	if u != nil {
		cu.SetEntID(*u)
	}
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *CapacityUpdate) SetCreatedAt(u uint32) *CapacityUpdate {
	cu.mutation.ResetCreatedAt()
	cu.mutation.SetCreatedAt(u)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *CapacityUpdate) SetNillableCreatedAt(u *uint32) *CapacityUpdate {
	if u != nil {
		cu.SetCreatedAt(*u)
	}
	return cu
}

// AddCreatedAt adds u to the "created_at" field.
func (cu *CapacityUpdate) AddCreatedAt(u int32) *CapacityUpdate {
	cu.mutation.AddCreatedAt(u)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CapacityUpdate) SetUpdatedAt(u uint32) *CapacityUpdate {
	cu.mutation.ResetUpdatedAt()
	cu.mutation.SetUpdatedAt(u)
	return cu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cu *CapacityUpdate) AddUpdatedAt(u int32) *CapacityUpdate {
	cu.mutation.AddUpdatedAt(u)
	return cu
}

// SetDeletedAt sets the "deleted_at" field.
func (cu *CapacityUpdate) SetDeletedAt(u uint32) *CapacityUpdate {
	cu.mutation.ResetDeletedAt()
	cu.mutation.SetDeletedAt(u)
	return cu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cu *CapacityUpdate) SetNillableDeletedAt(u *uint32) *CapacityUpdate {
	if u != nil {
		cu.SetDeletedAt(*u)
	}
	return cu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cu *CapacityUpdate) AddDeletedAt(u int32) *CapacityUpdate {
	cu.mutation.AddDeletedAt(u)
	return cu
}

// SetAppGoodID sets the "app_good_id" field.
func (cu *CapacityUpdate) SetAppGoodID(u uuid.UUID) *CapacityUpdate {
	cu.mutation.SetAppGoodID(u)
	return cu
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (cu *CapacityUpdate) SetNillableAppGoodID(u *uuid.UUID) *CapacityUpdate {
	if u != nil {
		cu.SetAppGoodID(*u)
	}
	return cu
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (cu *CapacityUpdate) ClearAppGoodID() *CapacityUpdate {
	cu.mutation.ClearAppGoodID()
	return cu
}

// SetCapacityKey sets the "capacity_key" field.
func (cu *CapacityUpdate) SetCapacityKey(s string) *CapacityUpdate {
	cu.mutation.SetCapacityKey(s)
	return cu
}

// SetNillableCapacityKey sets the "capacity_key" field if the given value is not nil.
func (cu *CapacityUpdate) SetNillableCapacityKey(s *string) *CapacityUpdate {
	if s != nil {
		cu.SetCapacityKey(*s)
	}
	return cu
}

// ClearCapacityKey clears the value of the "capacity_key" field.
func (cu *CapacityUpdate) ClearCapacityKey() *CapacityUpdate {
	cu.mutation.ClearCapacityKey()
	return cu
}

// SetCapacityValue sets the "capacity_value" field.
func (cu *CapacityUpdate) SetCapacityValue(s string) *CapacityUpdate {
	cu.mutation.SetCapacityValue(s)
	return cu
}

// SetNillableCapacityValue sets the "capacity_value" field if the given value is not nil.
func (cu *CapacityUpdate) SetNillableCapacityValue(s *string) *CapacityUpdate {
	if s != nil {
		cu.SetCapacityValue(*s)
	}
	return cu
}

// ClearCapacityValue clears the value of the "capacity_value" field.
func (cu *CapacityUpdate) ClearCapacityValue() *CapacityUpdate {
	cu.mutation.ClearCapacityValue()
	return cu
}

// SetDescription sets the "description" field.
func (cu *CapacityUpdate) SetDescription(s string) *CapacityUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cu *CapacityUpdate) SetNillableDescription(s *string) *CapacityUpdate {
	if s != nil {
		cu.SetDescription(*s)
	}
	return cu
}

// ClearDescription clears the value of the "description" field.
func (cu *CapacityUpdate) ClearDescription() *CapacityUpdate {
	cu.mutation.ClearDescription()
	return cu
}

// Mutation returns the CapacityMutation object of the builder.
func (cu *CapacityUpdate) Mutation() *CapacityMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CapacityUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CapacityUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CapacityUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CapacityUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CapacityUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := capacity.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cu *CapacityUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CapacityUpdate {
	cu.modifiers = append(cu.modifiers, modifiers...)
	return cu
}

func (cu *CapacityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(capacity.Table, capacity.Columns, sqlgraph.NewFieldSpec(capacity.FieldID, field.TypeUint32))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.EntID(); ok {
		_spec.SetField(capacity.FieldEntID, field.TypeUUID, value)
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.SetField(capacity.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := cu.mutation.AddedCreatedAt(); ok {
		_spec.AddField(capacity.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(capacity.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := cu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(capacity.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := cu.mutation.DeletedAt(); ok {
		_spec.SetField(capacity.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := cu.mutation.AddedDeletedAt(); ok {
		_spec.AddField(capacity.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := cu.mutation.AppGoodID(); ok {
		_spec.SetField(capacity.FieldAppGoodID, field.TypeUUID, value)
	}
	if cu.mutation.AppGoodIDCleared() {
		_spec.ClearField(capacity.FieldAppGoodID, field.TypeUUID)
	}
	if value, ok := cu.mutation.CapacityKey(); ok {
		_spec.SetField(capacity.FieldCapacityKey, field.TypeString, value)
	}
	if cu.mutation.CapacityKeyCleared() {
		_spec.ClearField(capacity.FieldCapacityKey, field.TypeString)
	}
	if value, ok := cu.mutation.CapacityValue(); ok {
		_spec.SetField(capacity.FieldCapacityValue, field.TypeString, value)
	}
	if cu.mutation.CapacityValueCleared() {
		_spec.ClearField(capacity.FieldCapacityValue, field.TypeString)
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.SetField(capacity.FieldDescription, field.TypeString, value)
	}
	if cu.mutation.DescriptionCleared() {
		_spec.ClearField(capacity.FieldDescription, field.TypeString)
	}
	_spec.AddModifiers(cu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{capacity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CapacityUpdateOne is the builder for updating a single Capacity entity.
type CapacityUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CapacityMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetEntID sets the "ent_id" field.
func (cuo *CapacityUpdateOne) SetEntID(u uuid.UUID) *CapacityUpdateOne {
	cuo.mutation.SetEntID(u)
	return cuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cuo *CapacityUpdateOne) SetNillableEntID(u *uuid.UUID) *CapacityUpdateOne {
	if u != nil {
		cuo.SetEntID(*u)
	}
	return cuo
}

// SetCreatedAt sets the "created_at" field.
func (cuo *CapacityUpdateOne) SetCreatedAt(u uint32) *CapacityUpdateOne {
	cuo.mutation.ResetCreatedAt()
	cuo.mutation.SetCreatedAt(u)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *CapacityUpdateOne) SetNillableCreatedAt(u *uint32) *CapacityUpdateOne {
	if u != nil {
		cuo.SetCreatedAt(*u)
	}
	return cuo
}

// AddCreatedAt adds u to the "created_at" field.
func (cuo *CapacityUpdateOne) AddCreatedAt(u int32) *CapacityUpdateOne {
	cuo.mutation.AddCreatedAt(u)
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CapacityUpdateOne) SetUpdatedAt(u uint32) *CapacityUpdateOne {
	cuo.mutation.ResetUpdatedAt()
	cuo.mutation.SetUpdatedAt(u)
	return cuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cuo *CapacityUpdateOne) AddUpdatedAt(u int32) *CapacityUpdateOne {
	cuo.mutation.AddUpdatedAt(u)
	return cuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cuo *CapacityUpdateOne) SetDeletedAt(u uint32) *CapacityUpdateOne {
	cuo.mutation.ResetDeletedAt()
	cuo.mutation.SetDeletedAt(u)
	return cuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cuo *CapacityUpdateOne) SetNillableDeletedAt(u *uint32) *CapacityUpdateOne {
	if u != nil {
		cuo.SetDeletedAt(*u)
	}
	return cuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cuo *CapacityUpdateOne) AddDeletedAt(u int32) *CapacityUpdateOne {
	cuo.mutation.AddDeletedAt(u)
	return cuo
}

// SetAppGoodID sets the "app_good_id" field.
func (cuo *CapacityUpdateOne) SetAppGoodID(u uuid.UUID) *CapacityUpdateOne {
	cuo.mutation.SetAppGoodID(u)
	return cuo
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (cuo *CapacityUpdateOne) SetNillableAppGoodID(u *uuid.UUID) *CapacityUpdateOne {
	if u != nil {
		cuo.SetAppGoodID(*u)
	}
	return cuo
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (cuo *CapacityUpdateOne) ClearAppGoodID() *CapacityUpdateOne {
	cuo.mutation.ClearAppGoodID()
	return cuo
}

// SetCapacityKey sets the "capacity_key" field.
func (cuo *CapacityUpdateOne) SetCapacityKey(s string) *CapacityUpdateOne {
	cuo.mutation.SetCapacityKey(s)
	return cuo
}

// SetNillableCapacityKey sets the "capacity_key" field if the given value is not nil.
func (cuo *CapacityUpdateOne) SetNillableCapacityKey(s *string) *CapacityUpdateOne {
	if s != nil {
		cuo.SetCapacityKey(*s)
	}
	return cuo
}

// ClearCapacityKey clears the value of the "capacity_key" field.
func (cuo *CapacityUpdateOne) ClearCapacityKey() *CapacityUpdateOne {
	cuo.mutation.ClearCapacityKey()
	return cuo
}

// SetCapacityValue sets the "capacity_value" field.
func (cuo *CapacityUpdateOne) SetCapacityValue(s string) *CapacityUpdateOne {
	cuo.mutation.SetCapacityValue(s)
	return cuo
}

// SetNillableCapacityValue sets the "capacity_value" field if the given value is not nil.
func (cuo *CapacityUpdateOne) SetNillableCapacityValue(s *string) *CapacityUpdateOne {
	if s != nil {
		cuo.SetCapacityValue(*s)
	}
	return cuo
}

// ClearCapacityValue clears the value of the "capacity_value" field.
func (cuo *CapacityUpdateOne) ClearCapacityValue() *CapacityUpdateOne {
	cuo.mutation.ClearCapacityValue()
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *CapacityUpdateOne) SetDescription(s string) *CapacityUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cuo *CapacityUpdateOne) SetNillableDescription(s *string) *CapacityUpdateOne {
	if s != nil {
		cuo.SetDescription(*s)
	}
	return cuo
}

// ClearDescription clears the value of the "description" field.
func (cuo *CapacityUpdateOne) ClearDescription() *CapacityUpdateOne {
	cuo.mutation.ClearDescription()
	return cuo
}

// Mutation returns the CapacityMutation object of the builder.
func (cuo *CapacityUpdateOne) Mutation() *CapacityMutation {
	return cuo.mutation
}

// Where appends a list predicates to the CapacityUpdate builder.
func (cuo *CapacityUpdateOne) Where(ps ...predicate.Capacity) *CapacityUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CapacityUpdateOne) Select(field string, fields ...string) *CapacityUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Capacity entity.
func (cuo *CapacityUpdateOne) Save(ctx context.Context) (*Capacity, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CapacityUpdateOne) SaveX(ctx context.Context) *Capacity {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CapacityUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CapacityUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CapacityUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := capacity.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuo *CapacityUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CapacityUpdateOne {
	cuo.modifiers = append(cuo.modifiers, modifiers...)
	return cuo
}

func (cuo *CapacityUpdateOne) sqlSave(ctx context.Context) (_node *Capacity, err error) {
	_spec := sqlgraph.NewUpdateSpec(capacity.Table, capacity.Columns, sqlgraph.NewFieldSpec(capacity.FieldID, field.TypeUint32))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Capacity.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, capacity.FieldID)
		for _, f := range fields {
			if !capacity.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != capacity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.EntID(); ok {
		_spec.SetField(capacity.FieldEntID, field.TypeUUID, value)
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.SetField(capacity.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := cuo.mutation.AddedCreatedAt(); ok {
		_spec.AddField(capacity.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(capacity.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := cuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(capacity.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := cuo.mutation.DeletedAt(); ok {
		_spec.SetField(capacity.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := cuo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(capacity.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := cuo.mutation.AppGoodID(); ok {
		_spec.SetField(capacity.FieldAppGoodID, field.TypeUUID, value)
	}
	if cuo.mutation.AppGoodIDCleared() {
		_spec.ClearField(capacity.FieldAppGoodID, field.TypeUUID)
	}
	if value, ok := cuo.mutation.CapacityKey(); ok {
		_spec.SetField(capacity.FieldCapacityKey, field.TypeString, value)
	}
	if cuo.mutation.CapacityKeyCleared() {
		_spec.ClearField(capacity.FieldCapacityKey, field.TypeString)
	}
	if value, ok := cuo.mutation.CapacityValue(); ok {
		_spec.SetField(capacity.FieldCapacityValue, field.TypeString, value)
	}
	if cuo.mutation.CapacityValueCleared() {
		_spec.ClearField(capacity.FieldCapacityValue, field.TypeString)
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.SetField(capacity.FieldDescription, field.TypeString, value)
	}
	if cuo.mutation.DescriptionCleared() {
		_spec.ClearField(capacity.FieldDescription, field.TypeString)
	}
	_spec.AddModifiers(cuo.modifiers...)
	_node = &Capacity{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{capacity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
