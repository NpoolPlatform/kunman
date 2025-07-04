// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated/lang"
	"github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated/predicate"
	"github.com/google/uuid"
)

// LangUpdate is the builder for updating Lang entities.
type LangUpdate struct {
	config
	hooks     []Hook
	mutation  *LangMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the LangUpdate builder.
func (lu *LangUpdate) Where(ps ...predicate.Lang) *LangUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetCreatedAt sets the "created_at" field.
func (lu *LangUpdate) SetCreatedAt(u uint32) *LangUpdate {
	lu.mutation.ResetCreatedAt()
	lu.mutation.SetCreatedAt(u)
	return lu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lu *LangUpdate) SetNillableCreatedAt(u *uint32) *LangUpdate {
	if u != nil {
		lu.SetCreatedAt(*u)
	}
	return lu
}

// AddCreatedAt adds u to the "created_at" field.
func (lu *LangUpdate) AddCreatedAt(u int32) *LangUpdate {
	lu.mutation.AddCreatedAt(u)
	return lu
}

// SetUpdatedAt sets the "updated_at" field.
func (lu *LangUpdate) SetUpdatedAt(u uint32) *LangUpdate {
	lu.mutation.ResetUpdatedAt()
	lu.mutation.SetUpdatedAt(u)
	return lu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (lu *LangUpdate) AddUpdatedAt(u int32) *LangUpdate {
	lu.mutation.AddUpdatedAt(u)
	return lu
}

// SetDeletedAt sets the "deleted_at" field.
func (lu *LangUpdate) SetDeletedAt(u uint32) *LangUpdate {
	lu.mutation.ResetDeletedAt()
	lu.mutation.SetDeletedAt(u)
	return lu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (lu *LangUpdate) SetNillableDeletedAt(u *uint32) *LangUpdate {
	if u != nil {
		lu.SetDeletedAt(*u)
	}
	return lu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (lu *LangUpdate) AddDeletedAt(u int32) *LangUpdate {
	lu.mutation.AddDeletedAt(u)
	return lu
}

// SetEntID sets the "ent_id" field.
func (lu *LangUpdate) SetEntID(u uuid.UUID) *LangUpdate {
	lu.mutation.SetEntID(u)
	return lu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (lu *LangUpdate) SetNillableEntID(u *uuid.UUID) *LangUpdate {
	if u != nil {
		lu.SetEntID(*u)
	}
	return lu
}

// SetLang sets the "lang" field.
func (lu *LangUpdate) SetLang(s string) *LangUpdate {
	lu.mutation.SetLang(s)
	return lu
}

// SetNillableLang sets the "lang" field if the given value is not nil.
func (lu *LangUpdate) SetNillableLang(s *string) *LangUpdate {
	if s != nil {
		lu.SetLang(*s)
	}
	return lu
}

// ClearLang clears the value of the "lang" field.
func (lu *LangUpdate) ClearLang() *LangUpdate {
	lu.mutation.ClearLang()
	return lu
}

// SetLogo sets the "logo" field.
func (lu *LangUpdate) SetLogo(s string) *LangUpdate {
	lu.mutation.SetLogo(s)
	return lu
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (lu *LangUpdate) SetNillableLogo(s *string) *LangUpdate {
	if s != nil {
		lu.SetLogo(*s)
	}
	return lu
}

// ClearLogo clears the value of the "logo" field.
func (lu *LangUpdate) ClearLogo() *LangUpdate {
	lu.mutation.ClearLogo()
	return lu
}

// SetName sets the "name" field.
func (lu *LangUpdate) SetName(s string) *LangUpdate {
	lu.mutation.SetName(s)
	return lu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (lu *LangUpdate) SetNillableName(s *string) *LangUpdate {
	if s != nil {
		lu.SetName(*s)
	}
	return lu
}

// ClearName clears the value of the "name" field.
func (lu *LangUpdate) ClearName() *LangUpdate {
	lu.mutation.ClearName()
	return lu
}

// SetShort sets the "short" field.
func (lu *LangUpdate) SetShort(s string) *LangUpdate {
	lu.mutation.SetShort(s)
	return lu
}

// SetNillableShort sets the "short" field if the given value is not nil.
func (lu *LangUpdate) SetNillableShort(s *string) *LangUpdate {
	if s != nil {
		lu.SetShort(*s)
	}
	return lu
}

// ClearShort clears the value of the "short" field.
func (lu *LangUpdate) ClearShort() *LangUpdate {
	lu.mutation.ClearShort()
	return lu
}

// Mutation returns the LangMutation object of the builder.
func (lu *LangUpdate) Mutation() *LangMutation {
	return lu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LangUpdate) Save(ctx context.Context) (int, error) {
	lu.defaults()
	return withHooks(ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LangUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LangUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LangUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lu *LangUpdate) defaults() {
	if _, ok := lu.mutation.UpdatedAt(); !ok {
		v := lang.UpdateDefaultUpdatedAt()
		lu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (lu *LangUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *LangUpdate {
	lu.modifiers = append(lu.modifiers, modifiers...)
	return lu
}

func (lu *LangUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(lang.Table, lang.Columns, sqlgraph.NewFieldSpec(lang.FieldID, field.TypeUint32))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.CreatedAt(); ok {
		_spec.SetField(lang.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := lu.mutation.AddedCreatedAt(); ok {
		_spec.AddField(lang.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := lu.mutation.UpdatedAt(); ok {
		_spec.SetField(lang.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := lu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(lang.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := lu.mutation.DeletedAt(); ok {
		_spec.SetField(lang.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := lu.mutation.AddedDeletedAt(); ok {
		_spec.AddField(lang.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := lu.mutation.EntID(); ok {
		_spec.SetField(lang.FieldEntID, field.TypeUUID, value)
	}
	if value, ok := lu.mutation.Lang(); ok {
		_spec.SetField(lang.FieldLang, field.TypeString, value)
	}
	if lu.mutation.LangCleared() {
		_spec.ClearField(lang.FieldLang, field.TypeString)
	}
	if value, ok := lu.mutation.Logo(); ok {
		_spec.SetField(lang.FieldLogo, field.TypeString, value)
	}
	if lu.mutation.LogoCleared() {
		_spec.ClearField(lang.FieldLogo, field.TypeString)
	}
	if value, ok := lu.mutation.Name(); ok {
		_spec.SetField(lang.FieldName, field.TypeString, value)
	}
	if lu.mutation.NameCleared() {
		_spec.ClearField(lang.FieldName, field.TypeString)
	}
	if value, ok := lu.mutation.Short(); ok {
		_spec.SetField(lang.FieldShort, field.TypeString, value)
	}
	if lu.mutation.ShortCleared() {
		_spec.ClearField(lang.FieldShort, field.TypeString)
	}
	_spec.AddModifiers(lu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lang.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LangUpdateOne is the builder for updating a single Lang entity.
type LangUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *LangMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (luo *LangUpdateOne) SetCreatedAt(u uint32) *LangUpdateOne {
	luo.mutation.ResetCreatedAt()
	luo.mutation.SetCreatedAt(u)
	return luo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (luo *LangUpdateOne) SetNillableCreatedAt(u *uint32) *LangUpdateOne {
	if u != nil {
		luo.SetCreatedAt(*u)
	}
	return luo
}

// AddCreatedAt adds u to the "created_at" field.
func (luo *LangUpdateOne) AddCreatedAt(u int32) *LangUpdateOne {
	luo.mutation.AddCreatedAt(u)
	return luo
}

// SetUpdatedAt sets the "updated_at" field.
func (luo *LangUpdateOne) SetUpdatedAt(u uint32) *LangUpdateOne {
	luo.mutation.ResetUpdatedAt()
	luo.mutation.SetUpdatedAt(u)
	return luo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (luo *LangUpdateOne) AddUpdatedAt(u int32) *LangUpdateOne {
	luo.mutation.AddUpdatedAt(u)
	return luo
}

// SetDeletedAt sets the "deleted_at" field.
func (luo *LangUpdateOne) SetDeletedAt(u uint32) *LangUpdateOne {
	luo.mutation.ResetDeletedAt()
	luo.mutation.SetDeletedAt(u)
	return luo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (luo *LangUpdateOne) SetNillableDeletedAt(u *uint32) *LangUpdateOne {
	if u != nil {
		luo.SetDeletedAt(*u)
	}
	return luo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (luo *LangUpdateOne) AddDeletedAt(u int32) *LangUpdateOne {
	luo.mutation.AddDeletedAt(u)
	return luo
}

// SetEntID sets the "ent_id" field.
func (luo *LangUpdateOne) SetEntID(u uuid.UUID) *LangUpdateOne {
	luo.mutation.SetEntID(u)
	return luo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (luo *LangUpdateOne) SetNillableEntID(u *uuid.UUID) *LangUpdateOne {
	if u != nil {
		luo.SetEntID(*u)
	}
	return luo
}

// SetLang sets the "lang" field.
func (luo *LangUpdateOne) SetLang(s string) *LangUpdateOne {
	luo.mutation.SetLang(s)
	return luo
}

// SetNillableLang sets the "lang" field if the given value is not nil.
func (luo *LangUpdateOne) SetNillableLang(s *string) *LangUpdateOne {
	if s != nil {
		luo.SetLang(*s)
	}
	return luo
}

// ClearLang clears the value of the "lang" field.
func (luo *LangUpdateOne) ClearLang() *LangUpdateOne {
	luo.mutation.ClearLang()
	return luo
}

// SetLogo sets the "logo" field.
func (luo *LangUpdateOne) SetLogo(s string) *LangUpdateOne {
	luo.mutation.SetLogo(s)
	return luo
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (luo *LangUpdateOne) SetNillableLogo(s *string) *LangUpdateOne {
	if s != nil {
		luo.SetLogo(*s)
	}
	return luo
}

// ClearLogo clears the value of the "logo" field.
func (luo *LangUpdateOne) ClearLogo() *LangUpdateOne {
	luo.mutation.ClearLogo()
	return luo
}

// SetName sets the "name" field.
func (luo *LangUpdateOne) SetName(s string) *LangUpdateOne {
	luo.mutation.SetName(s)
	return luo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (luo *LangUpdateOne) SetNillableName(s *string) *LangUpdateOne {
	if s != nil {
		luo.SetName(*s)
	}
	return luo
}

// ClearName clears the value of the "name" field.
func (luo *LangUpdateOne) ClearName() *LangUpdateOne {
	luo.mutation.ClearName()
	return luo
}

// SetShort sets the "short" field.
func (luo *LangUpdateOne) SetShort(s string) *LangUpdateOne {
	luo.mutation.SetShort(s)
	return luo
}

// SetNillableShort sets the "short" field if the given value is not nil.
func (luo *LangUpdateOne) SetNillableShort(s *string) *LangUpdateOne {
	if s != nil {
		luo.SetShort(*s)
	}
	return luo
}

// ClearShort clears the value of the "short" field.
func (luo *LangUpdateOne) ClearShort() *LangUpdateOne {
	luo.mutation.ClearShort()
	return luo
}

// Mutation returns the LangMutation object of the builder.
func (luo *LangUpdateOne) Mutation() *LangMutation {
	return luo.mutation
}

// Where appends a list predicates to the LangUpdate builder.
func (luo *LangUpdateOne) Where(ps ...predicate.Lang) *LangUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LangUpdateOne) Select(field string, fields ...string) *LangUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Lang entity.
func (luo *LangUpdateOne) Save(ctx context.Context) (*Lang, error) {
	luo.defaults()
	return withHooks(ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LangUpdateOne) SaveX(ctx context.Context) *Lang {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LangUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LangUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (luo *LangUpdateOne) defaults() {
	if _, ok := luo.mutation.UpdatedAt(); !ok {
		v := lang.UpdateDefaultUpdatedAt()
		luo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (luo *LangUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *LangUpdateOne {
	luo.modifiers = append(luo.modifiers, modifiers...)
	return luo
}

func (luo *LangUpdateOne) sqlSave(ctx context.Context) (_node *Lang, err error) {
	_spec := sqlgraph.NewUpdateSpec(lang.Table, lang.Columns, sqlgraph.NewFieldSpec(lang.FieldID, field.TypeUint32))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Lang.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lang.FieldID)
		for _, f := range fields {
			if !lang.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != lang.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.CreatedAt(); ok {
		_spec.SetField(lang.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := luo.mutation.AddedCreatedAt(); ok {
		_spec.AddField(lang.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := luo.mutation.UpdatedAt(); ok {
		_spec.SetField(lang.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := luo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(lang.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := luo.mutation.DeletedAt(); ok {
		_spec.SetField(lang.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := luo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(lang.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := luo.mutation.EntID(); ok {
		_spec.SetField(lang.FieldEntID, field.TypeUUID, value)
	}
	if value, ok := luo.mutation.Lang(); ok {
		_spec.SetField(lang.FieldLang, field.TypeString, value)
	}
	if luo.mutation.LangCleared() {
		_spec.ClearField(lang.FieldLang, field.TypeString)
	}
	if value, ok := luo.mutation.Logo(); ok {
		_spec.SetField(lang.FieldLogo, field.TypeString, value)
	}
	if luo.mutation.LogoCleared() {
		_spec.ClearField(lang.FieldLogo, field.TypeString)
	}
	if value, ok := luo.mutation.Name(); ok {
		_spec.SetField(lang.FieldName, field.TypeString, value)
	}
	if luo.mutation.NameCleared() {
		_spec.ClearField(lang.FieldName, field.TypeString)
	}
	if value, ok := luo.mutation.Short(); ok {
		_spec.SetField(lang.FieldShort, field.TypeString, value)
	}
	if luo.mutation.ShortCleared() {
		_spec.ClearField(lang.FieldShort, field.TypeString)
	}
	_spec.AddModifiers(luo.modifiers...)
	_node = &Lang{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lang.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
