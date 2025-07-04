// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/predicate"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/topmostposter"
	"github.com/google/uuid"
)

// TopMostPosterUpdate is the builder for updating TopMostPoster entities.
type TopMostPosterUpdate struct {
	config
	hooks     []Hook
	mutation  *TopMostPosterMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TopMostPosterUpdate builder.
func (tmpu *TopMostPosterUpdate) Where(ps ...predicate.TopMostPoster) *TopMostPosterUpdate {
	tmpu.mutation.Where(ps...)
	return tmpu
}

// SetEntID sets the "ent_id" field.
func (tmpu *TopMostPosterUpdate) SetEntID(u uuid.UUID) *TopMostPosterUpdate {
	tmpu.mutation.SetEntID(u)
	return tmpu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (tmpu *TopMostPosterUpdate) SetNillableEntID(u *uuid.UUID) *TopMostPosterUpdate {
	if u != nil {
		tmpu.SetEntID(*u)
	}
	return tmpu
}

// SetCreatedAt sets the "created_at" field.
func (tmpu *TopMostPosterUpdate) SetCreatedAt(u uint32) *TopMostPosterUpdate {
	tmpu.mutation.ResetCreatedAt()
	tmpu.mutation.SetCreatedAt(u)
	return tmpu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tmpu *TopMostPosterUpdate) SetNillableCreatedAt(u *uint32) *TopMostPosterUpdate {
	if u != nil {
		tmpu.SetCreatedAt(*u)
	}
	return tmpu
}

// AddCreatedAt adds u to the "created_at" field.
func (tmpu *TopMostPosterUpdate) AddCreatedAt(u int32) *TopMostPosterUpdate {
	tmpu.mutation.AddCreatedAt(u)
	return tmpu
}

// SetUpdatedAt sets the "updated_at" field.
func (tmpu *TopMostPosterUpdate) SetUpdatedAt(u uint32) *TopMostPosterUpdate {
	tmpu.mutation.ResetUpdatedAt()
	tmpu.mutation.SetUpdatedAt(u)
	return tmpu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (tmpu *TopMostPosterUpdate) AddUpdatedAt(u int32) *TopMostPosterUpdate {
	tmpu.mutation.AddUpdatedAt(u)
	return tmpu
}

// SetDeletedAt sets the "deleted_at" field.
func (tmpu *TopMostPosterUpdate) SetDeletedAt(u uint32) *TopMostPosterUpdate {
	tmpu.mutation.ResetDeletedAt()
	tmpu.mutation.SetDeletedAt(u)
	return tmpu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tmpu *TopMostPosterUpdate) SetNillableDeletedAt(u *uint32) *TopMostPosterUpdate {
	if u != nil {
		tmpu.SetDeletedAt(*u)
	}
	return tmpu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (tmpu *TopMostPosterUpdate) AddDeletedAt(u int32) *TopMostPosterUpdate {
	tmpu.mutation.AddDeletedAt(u)
	return tmpu
}

// SetTopMostID sets the "top_most_id" field.
func (tmpu *TopMostPosterUpdate) SetTopMostID(u uuid.UUID) *TopMostPosterUpdate {
	tmpu.mutation.SetTopMostID(u)
	return tmpu
}

// SetNillableTopMostID sets the "top_most_id" field if the given value is not nil.
func (tmpu *TopMostPosterUpdate) SetNillableTopMostID(u *uuid.UUID) *TopMostPosterUpdate {
	if u != nil {
		tmpu.SetTopMostID(*u)
	}
	return tmpu
}

// ClearTopMostID clears the value of the "top_most_id" field.
func (tmpu *TopMostPosterUpdate) ClearTopMostID() *TopMostPosterUpdate {
	tmpu.mutation.ClearTopMostID()
	return tmpu
}

// SetPoster sets the "poster" field.
func (tmpu *TopMostPosterUpdate) SetPoster(s string) *TopMostPosterUpdate {
	tmpu.mutation.SetPoster(s)
	return tmpu
}

// SetNillablePoster sets the "poster" field if the given value is not nil.
func (tmpu *TopMostPosterUpdate) SetNillablePoster(s *string) *TopMostPosterUpdate {
	if s != nil {
		tmpu.SetPoster(*s)
	}
	return tmpu
}

// ClearPoster clears the value of the "poster" field.
func (tmpu *TopMostPosterUpdate) ClearPoster() *TopMostPosterUpdate {
	tmpu.mutation.ClearPoster()
	return tmpu
}

// SetIndex sets the "index" field.
func (tmpu *TopMostPosterUpdate) SetIndex(u uint8) *TopMostPosterUpdate {
	tmpu.mutation.ResetIndex()
	tmpu.mutation.SetIndex(u)
	return tmpu
}

// SetNillableIndex sets the "index" field if the given value is not nil.
func (tmpu *TopMostPosterUpdate) SetNillableIndex(u *uint8) *TopMostPosterUpdate {
	if u != nil {
		tmpu.SetIndex(*u)
	}
	return tmpu
}

// AddIndex adds u to the "index" field.
func (tmpu *TopMostPosterUpdate) AddIndex(u int8) *TopMostPosterUpdate {
	tmpu.mutation.AddIndex(u)
	return tmpu
}

// ClearIndex clears the value of the "index" field.
func (tmpu *TopMostPosterUpdate) ClearIndex() *TopMostPosterUpdate {
	tmpu.mutation.ClearIndex()
	return tmpu
}

// Mutation returns the TopMostPosterMutation object of the builder.
func (tmpu *TopMostPosterUpdate) Mutation() *TopMostPosterMutation {
	return tmpu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tmpu *TopMostPosterUpdate) Save(ctx context.Context) (int, error) {
	tmpu.defaults()
	return withHooks(ctx, tmpu.sqlSave, tmpu.mutation, tmpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tmpu *TopMostPosterUpdate) SaveX(ctx context.Context) int {
	affected, err := tmpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tmpu *TopMostPosterUpdate) Exec(ctx context.Context) error {
	_, err := tmpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tmpu *TopMostPosterUpdate) ExecX(ctx context.Context) {
	if err := tmpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tmpu *TopMostPosterUpdate) defaults() {
	if _, ok := tmpu.mutation.UpdatedAt(); !ok {
		v := topmostposter.UpdateDefaultUpdatedAt()
		tmpu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tmpu *TopMostPosterUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TopMostPosterUpdate {
	tmpu.modifiers = append(tmpu.modifiers, modifiers...)
	return tmpu
}

func (tmpu *TopMostPosterUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(topmostposter.Table, topmostposter.Columns, sqlgraph.NewFieldSpec(topmostposter.FieldID, field.TypeUint32))
	if ps := tmpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tmpu.mutation.EntID(); ok {
		_spec.SetField(topmostposter.FieldEntID, field.TypeUUID, value)
	}
	if value, ok := tmpu.mutation.CreatedAt(); ok {
		_spec.SetField(topmostposter.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := tmpu.mutation.AddedCreatedAt(); ok {
		_spec.AddField(topmostposter.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := tmpu.mutation.UpdatedAt(); ok {
		_spec.SetField(topmostposter.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := tmpu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(topmostposter.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := tmpu.mutation.DeletedAt(); ok {
		_spec.SetField(topmostposter.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := tmpu.mutation.AddedDeletedAt(); ok {
		_spec.AddField(topmostposter.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := tmpu.mutation.TopMostID(); ok {
		_spec.SetField(topmostposter.FieldTopMostID, field.TypeUUID, value)
	}
	if tmpu.mutation.TopMostIDCleared() {
		_spec.ClearField(topmostposter.FieldTopMostID, field.TypeUUID)
	}
	if value, ok := tmpu.mutation.Poster(); ok {
		_spec.SetField(topmostposter.FieldPoster, field.TypeString, value)
	}
	if tmpu.mutation.PosterCleared() {
		_spec.ClearField(topmostposter.FieldPoster, field.TypeString)
	}
	if value, ok := tmpu.mutation.Index(); ok {
		_spec.SetField(topmostposter.FieldIndex, field.TypeUint8, value)
	}
	if value, ok := tmpu.mutation.AddedIndex(); ok {
		_spec.AddField(topmostposter.FieldIndex, field.TypeUint8, value)
	}
	if tmpu.mutation.IndexCleared() {
		_spec.ClearField(topmostposter.FieldIndex, field.TypeUint8)
	}
	_spec.AddModifiers(tmpu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, tmpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{topmostposter.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tmpu.mutation.done = true
	return n, nil
}

// TopMostPosterUpdateOne is the builder for updating a single TopMostPoster entity.
type TopMostPosterUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TopMostPosterMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetEntID sets the "ent_id" field.
func (tmpuo *TopMostPosterUpdateOne) SetEntID(u uuid.UUID) *TopMostPosterUpdateOne {
	tmpuo.mutation.SetEntID(u)
	return tmpuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (tmpuo *TopMostPosterUpdateOne) SetNillableEntID(u *uuid.UUID) *TopMostPosterUpdateOne {
	if u != nil {
		tmpuo.SetEntID(*u)
	}
	return tmpuo
}

// SetCreatedAt sets the "created_at" field.
func (tmpuo *TopMostPosterUpdateOne) SetCreatedAt(u uint32) *TopMostPosterUpdateOne {
	tmpuo.mutation.ResetCreatedAt()
	tmpuo.mutation.SetCreatedAt(u)
	return tmpuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tmpuo *TopMostPosterUpdateOne) SetNillableCreatedAt(u *uint32) *TopMostPosterUpdateOne {
	if u != nil {
		tmpuo.SetCreatedAt(*u)
	}
	return tmpuo
}

// AddCreatedAt adds u to the "created_at" field.
func (tmpuo *TopMostPosterUpdateOne) AddCreatedAt(u int32) *TopMostPosterUpdateOne {
	tmpuo.mutation.AddCreatedAt(u)
	return tmpuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tmpuo *TopMostPosterUpdateOne) SetUpdatedAt(u uint32) *TopMostPosterUpdateOne {
	tmpuo.mutation.ResetUpdatedAt()
	tmpuo.mutation.SetUpdatedAt(u)
	return tmpuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (tmpuo *TopMostPosterUpdateOne) AddUpdatedAt(u int32) *TopMostPosterUpdateOne {
	tmpuo.mutation.AddUpdatedAt(u)
	return tmpuo
}

// SetDeletedAt sets the "deleted_at" field.
func (tmpuo *TopMostPosterUpdateOne) SetDeletedAt(u uint32) *TopMostPosterUpdateOne {
	tmpuo.mutation.ResetDeletedAt()
	tmpuo.mutation.SetDeletedAt(u)
	return tmpuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tmpuo *TopMostPosterUpdateOne) SetNillableDeletedAt(u *uint32) *TopMostPosterUpdateOne {
	if u != nil {
		tmpuo.SetDeletedAt(*u)
	}
	return tmpuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (tmpuo *TopMostPosterUpdateOne) AddDeletedAt(u int32) *TopMostPosterUpdateOne {
	tmpuo.mutation.AddDeletedAt(u)
	return tmpuo
}

// SetTopMostID sets the "top_most_id" field.
func (tmpuo *TopMostPosterUpdateOne) SetTopMostID(u uuid.UUID) *TopMostPosterUpdateOne {
	tmpuo.mutation.SetTopMostID(u)
	return tmpuo
}

// SetNillableTopMostID sets the "top_most_id" field if the given value is not nil.
func (tmpuo *TopMostPosterUpdateOne) SetNillableTopMostID(u *uuid.UUID) *TopMostPosterUpdateOne {
	if u != nil {
		tmpuo.SetTopMostID(*u)
	}
	return tmpuo
}

// ClearTopMostID clears the value of the "top_most_id" field.
func (tmpuo *TopMostPosterUpdateOne) ClearTopMostID() *TopMostPosterUpdateOne {
	tmpuo.mutation.ClearTopMostID()
	return tmpuo
}

// SetPoster sets the "poster" field.
func (tmpuo *TopMostPosterUpdateOne) SetPoster(s string) *TopMostPosterUpdateOne {
	tmpuo.mutation.SetPoster(s)
	return tmpuo
}

// SetNillablePoster sets the "poster" field if the given value is not nil.
func (tmpuo *TopMostPosterUpdateOne) SetNillablePoster(s *string) *TopMostPosterUpdateOne {
	if s != nil {
		tmpuo.SetPoster(*s)
	}
	return tmpuo
}

// ClearPoster clears the value of the "poster" field.
func (tmpuo *TopMostPosterUpdateOne) ClearPoster() *TopMostPosterUpdateOne {
	tmpuo.mutation.ClearPoster()
	return tmpuo
}

// SetIndex sets the "index" field.
func (tmpuo *TopMostPosterUpdateOne) SetIndex(u uint8) *TopMostPosterUpdateOne {
	tmpuo.mutation.ResetIndex()
	tmpuo.mutation.SetIndex(u)
	return tmpuo
}

// SetNillableIndex sets the "index" field if the given value is not nil.
func (tmpuo *TopMostPosterUpdateOne) SetNillableIndex(u *uint8) *TopMostPosterUpdateOne {
	if u != nil {
		tmpuo.SetIndex(*u)
	}
	return tmpuo
}

// AddIndex adds u to the "index" field.
func (tmpuo *TopMostPosterUpdateOne) AddIndex(u int8) *TopMostPosterUpdateOne {
	tmpuo.mutation.AddIndex(u)
	return tmpuo
}

// ClearIndex clears the value of the "index" field.
func (tmpuo *TopMostPosterUpdateOne) ClearIndex() *TopMostPosterUpdateOne {
	tmpuo.mutation.ClearIndex()
	return tmpuo
}

// Mutation returns the TopMostPosterMutation object of the builder.
func (tmpuo *TopMostPosterUpdateOne) Mutation() *TopMostPosterMutation {
	return tmpuo.mutation
}

// Where appends a list predicates to the TopMostPosterUpdate builder.
func (tmpuo *TopMostPosterUpdateOne) Where(ps ...predicate.TopMostPoster) *TopMostPosterUpdateOne {
	tmpuo.mutation.Where(ps...)
	return tmpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tmpuo *TopMostPosterUpdateOne) Select(field string, fields ...string) *TopMostPosterUpdateOne {
	tmpuo.fields = append([]string{field}, fields...)
	return tmpuo
}

// Save executes the query and returns the updated TopMostPoster entity.
func (tmpuo *TopMostPosterUpdateOne) Save(ctx context.Context) (*TopMostPoster, error) {
	tmpuo.defaults()
	return withHooks(ctx, tmpuo.sqlSave, tmpuo.mutation, tmpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tmpuo *TopMostPosterUpdateOne) SaveX(ctx context.Context) *TopMostPoster {
	node, err := tmpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tmpuo *TopMostPosterUpdateOne) Exec(ctx context.Context) error {
	_, err := tmpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tmpuo *TopMostPosterUpdateOne) ExecX(ctx context.Context) {
	if err := tmpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tmpuo *TopMostPosterUpdateOne) defaults() {
	if _, ok := tmpuo.mutation.UpdatedAt(); !ok {
		v := topmostposter.UpdateDefaultUpdatedAt()
		tmpuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tmpuo *TopMostPosterUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TopMostPosterUpdateOne {
	tmpuo.modifiers = append(tmpuo.modifiers, modifiers...)
	return tmpuo
}

func (tmpuo *TopMostPosterUpdateOne) sqlSave(ctx context.Context) (_node *TopMostPoster, err error) {
	_spec := sqlgraph.NewUpdateSpec(topmostposter.Table, topmostposter.Columns, sqlgraph.NewFieldSpec(topmostposter.FieldID, field.TypeUint32))
	id, ok := tmpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "TopMostPoster.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tmpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, topmostposter.FieldID)
		for _, f := range fields {
			if !topmostposter.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != topmostposter.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tmpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tmpuo.mutation.EntID(); ok {
		_spec.SetField(topmostposter.FieldEntID, field.TypeUUID, value)
	}
	if value, ok := tmpuo.mutation.CreatedAt(); ok {
		_spec.SetField(topmostposter.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := tmpuo.mutation.AddedCreatedAt(); ok {
		_spec.AddField(topmostposter.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := tmpuo.mutation.UpdatedAt(); ok {
		_spec.SetField(topmostposter.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := tmpuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(topmostposter.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := tmpuo.mutation.DeletedAt(); ok {
		_spec.SetField(topmostposter.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := tmpuo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(topmostposter.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := tmpuo.mutation.TopMostID(); ok {
		_spec.SetField(topmostposter.FieldTopMostID, field.TypeUUID, value)
	}
	if tmpuo.mutation.TopMostIDCleared() {
		_spec.ClearField(topmostposter.FieldTopMostID, field.TypeUUID)
	}
	if value, ok := tmpuo.mutation.Poster(); ok {
		_spec.SetField(topmostposter.FieldPoster, field.TypeString, value)
	}
	if tmpuo.mutation.PosterCleared() {
		_spec.ClearField(topmostposter.FieldPoster, field.TypeString)
	}
	if value, ok := tmpuo.mutation.Index(); ok {
		_spec.SetField(topmostposter.FieldIndex, field.TypeUint8, value)
	}
	if value, ok := tmpuo.mutation.AddedIndex(); ok {
		_spec.AddField(topmostposter.FieldIndex, field.TypeUint8, value)
	}
	if tmpuo.mutation.IndexCleared() {
		_spec.ClearField(topmostposter.FieldIndex, field.TypeUint8)
	}
	_spec.AddModifiers(tmpuo.modifiers...)
	_node = &TopMostPoster{config: tmpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tmpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{topmostposter.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tmpuo.mutation.done = true
	return _node, nil
}
