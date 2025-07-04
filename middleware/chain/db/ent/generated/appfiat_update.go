// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/appfiat"
	"github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/predicate"
	"github.com/google/uuid"
)

// AppFiatUpdate is the builder for updating AppFiat entities.
type AppFiatUpdate struct {
	config
	hooks     []Hook
	mutation  *AppFiatMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AppFiatUpdate builder.
func (afu *AppFiatUpdate) Where(ps ...predicate.AppFiat) *AppFiatUpdate {
	afu.mutation.Where(ps...)
	return afu
}

// SetCreatedAt sets the "created_at" field.
func (afu *AppFiatUpdate) SetCreatedAt(u uint32) *AppFiatUpdate {
	afu.mutation.ResetCreatedAt()
	afu.mutation.SetCreatedAt(u)
	return afu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (afu *AppFiatUpdate) SetNillableCreatedAt(u *uint32) *AppFiatUpdate {
	if u != nil {
		afu.SetCreatedAt(*u)
	}
	return afu
}

// AddCreatedAt adds u to the "created_at" field.
func (afu *AppFiatUpdate) AddCreatedAt(u int32) *AppFiatUpdate {
	afu.mutation.AddCreatedAt(u)
	return afu
}

// SetUpdatedAt sets the "updated_at" field.
func (afu *AppFiatUpdate) SetUpdatedAt(u uint32) *AppFiatUpdate {
	afu.mutation.ResetUpdatedAt()
	afu.mutation.SetUpdatedAt(u)
	return afu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (afu *AppFiatUpdate) AddUpdatedAt(u int32) *AppFiatUpdate {
	afu.mutation.AddUpdatedAt(u)
	return afu
}

// SetDeletedAt sets the "deleted_at" field.
func (afu *AppFiatUpdate) SetDeletedAt(u uint32) *AppFiatUpdate {
	afu.mutation.ResetDeletedAt()
	afu.mutation.SetDeletedAt(u)
	return afu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (afu *AppFiatUpdate) SetNillableDeletedAt(u *uint32) *AppFiatUpdate {
	if u != nil {
		afu.SetDeletedAt(*u)
	}
	return afu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (afu *AppFiatUpdate) AddDeletedAt(u int32) *AppFiatUpdate {
	afu.mutation.AddDeletedAt(u)
	return afu
}

// SetEntID sets the "ent_id" field.
func (afu *AppFiatUpdate) SetEntID(u uuid.UUID) *AppFiatUpdate {
	afu.mutation.SetEntID(u)
	return afu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (afu *AppFiatUpdate) SetNillableEntID(u *uuid.UUID) *AppFiatUpdate {
	if u != nil {
		afu.SetEntID(*u)
	}
	return afu
}

// SetAppID sets the "app_id" field.
func (afu *AppFiatUpdate) SetAppID(u uuid.UUID) *AppFiatUpdate {
	afu.mutation.SetAppID(u)
	return afu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (afu *AppFiatUpdate) SetNillableAppID(u *uuid.UUID) *AppFiatUpdate {
	if u != nil {
		afu.SetAppID(*u)
	}
	return afu
}

// ClearAppID clears the value of the "app_id" field.
func (afu *AppFiatUpdate) ClearAppID() *AppFiatUpdate {
	afu.mutation.ClearAppID()
	return afu
}

// SetFiatID sets the "fiat_id" field.
func (afu *AppFiatUpdate) SetFiatID(u uuid.UUID) *AppFiatUpdate {
	afu.mutation.SetFiatID(u)
	return afu
}

// SetNillableFiatID sets the "fiat_id" field if the given value is not nil.
func (afu *AppFiatUpdate) SetNillableFiatID(u *uuid.UUID) *AppFiatUpdate {
	if u != nil {
		afu.SetFiatID(*u)
	}
	return afu
}

// ClearFiatID clears the value of the "fiat_id" field.
func (afu *AppFiatUpdate) ClearFiatID() *AppFiatUpdate {
	afu.mutation.ClearFiatID()
	return afu
}

// SetName sets the "name" field.
func (afu *AppFiatUpdate) SetName(s string) *AppFiatUpdate {
	afu.mutation.SetName(s)
	return afu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (afu *AppFiatUpdate) SetNillableName(s *string) *AppFiatUpdate {
	if s != nil {
		afu.SetName(*s)
	}
	return afu
}

// ClearName clears the value of the "name" field.
func (afu *AppFiatUpdate) ClearName() *AppFiatUpdate {
	afu.mutation.ClearName()
	return afu
}

// SetDisplayNames sets the "display_names" field.
func (afu *AppFiatUpdate) SetDisplayNames(s []string) *AppFiatUpdate {
	afu.mutation.SetDisplayNames(s)
	return afu
}

// AppendDisplayNames appends s to the "display_names" field.
func (afu *AppFiatUpdate) AppendDisplayNames(s []string) *AppFiatUpdate {
	afu.mutation.AppendDisplayNames(s)
	return afu
}

// ClearDisplayNames clears the value of the "display_names" field.
func (afu *AppFiatUpdate) ClearDisplayNames() *AppFiatUpdate {
	afu.mutation.ClearDisplayNames()
	return afu
}

// SetLogo sets the "logo" field.
func (afu *AppFiatUpdate) SetLogo(s string) *AppFiatUpdate {
	afu.mutation.SetLogo(s)
	return afu
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (afu *AppFiatUpdate) SetNillableLogo(s *string) *AppFiatUpdate {
	if s != nil {
		afu.SetLogo(*s)
	}
	return afu
}

// ClearLogo clears the value of the "logo" field.
func (afu *AppFiatUpdate) ClearLogo() *AppFiatUpdate {
	afu.mutation.ClearLogo()
	return afu
}

// SetDisabled sets the "disabled" field.
func (afu *AppFiatUpdate) SetDisabled(b bool) *AppFiatUpdate {
	afu.mutation.SetDisabled(b)
	return afu
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (afu *AppFiatUpdate) SetNillableDisabled(b *bool) *AppFiatUpdate {
	if b != nil {
		afu.SetDisabled(*b)
	}
	return afu
}

// ClearDisabled clears the value of the "disabled" field.
func (afu *AppFiatUpdate) ClearDisabled() *AppFiatUpdate {
	afu.mutation.ClearDisabled()
	return afu
}

// SetDisplay sets the "display" field.
func (afu *AppFiatUpdate) SetDisplay(b bool) *AppFiatUpdate {
	afu.mutation.SetDisplay(b)
	return afu
}

// SetNillableDisplay sets the "display" field if the given value is not nil.
func (afu *AppFiatUpdate) SetNillableDisplay(b *bool) *AppFiatUpdate {
	if b != nil {
		afu.SetDisplay(*b)
	}
	return afu
}

// ClearDisplay clears the value of the "display" field.
func (afu *AppFiatUpdate) ClearDisplay() *AppFiatUpdate {
	afu.mutation.ClearDisplay()
	return afu
}

// SetDisplayIndex sets the "display_index" field.
func (afu *AppFiatUpdate) SetDisplayIndex(u uint32) *AppFiatUpdate {
	afu.mutation.ResetDisplayIndex()
	afu.mutation.SetDisplayIndex(u)
	return afu
}

// SetNillableDisplayIndex sets the "display_index" field if the given value is not nil.
func (afu *AppFiatUpdate) SetNillableDisplayIndex(u *uint32) *AppFiatUpdate {
	if u != nil {
		afu.SetDisplayIndex(*u)
	}
	return afu
}

// AddDisplayIndex adds u to the "display_index" field.
func (afu *AppFiatUpdate) AddDisplayIndex(u int32) *AppFiatUpdate {
	afu.mutation.AddDisplayIndex(u)
	return afu
}

// ClearDisplayIndex clears the value of the "display_index" field.
func (afu *AppFiatUpdate) ClearDisplayIndex() *AppFiatUpdate {
	afu.mutation.ClearDisplayIndex()
	return afu
}

// Mutation returns the AppFiatMutation object of the builder.
func (afu *AppFiatUpdate) Mutation() *AppFiatMutation {
	return afu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (afu *AppFiatUpdate) Save(ctx context.Context) (int, error) {
	afu.defaults()
	return withHooks(ctx, afu.sqlSave, afu.mutation, afu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (afu *AppFiatUpdate) SaveX(ctx context.Context) int {
	affected, err := afu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (afu *AppFiatUpdate) Exec(ctx context.Context) error {
	_, err := afu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (afu *AppFiatUpdate) ExecX(ctx context.Context) {
	if err := afu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (afu *AppFiatUpdate) defaults() {
	if _, ok := afu.mutation.UpdatedAt(); !ok {
		v := appfiat.UpdateDefaultUpdatedAt()
		afu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (afu *AppFiatUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppFiatUpdate {
	afu.modifiers = append(afu.modifiers, modifiers...)
	return afu
}

func (afu *AppFiatUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(appfiat.Table, appfiat.Columns, sqlgraph.NewFieldSpec(appfiat.FieldID, field.TypeUint32))
	if ps := afu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := afu.mutation.CreatedAt(); ok {
		_spec.SetField(appfiat.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := afu.mutation.AddedCreatedAt(); ok {
		_spec.AddField(appfiat.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := afu.mutation.UpdatedAt(); ok {
		_spec.SetField(appfiat.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := afu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(appfiat.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := afu.mutation.DeletedAt(); ok {
		_spec.SetField(appfiat.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := afu.mutation.AddedDeletedAt(); ok {
		_spec.AddField(appfiat.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := afu.mutation.EntID(); ok {
		_spec.SetField(appfiat.FieldEntID, field.TypeUUID, value)
	}
	if value, ok := afu.mutation.AppID(); ok {
		_spec.SetField(appfiat.FieldAppID, field.TypeUUID, value)
	}
	if afu.mutation.AppIDCleared() {
		_spec.ClearField(appfiat.FieldAppID, field.TypeUUID)
	}
	if value, ok := afu.mutation.FiatID(); ok {
		_spec.SetField(appfiat.FieldFiatID, field.TypeUUID, value)
	}
	if afu.mutation.FiatIDCleared() {
		_spec.ClearField(appfiat.FieldFiatID, field.TypeUUID)
	}
	if value, ok := afu.mutation.Name(); ok {
		_spec.SetField(appfiat.FieldName, field.TypeString, value)
	}
	if afu.mutation.NameCleared() {
		_spec.ClearField(appfiat.FieldName, field.TypeString)
	}
	if value, ok := afu.mutation.DisplayNames(); ok {
		_spec.SetField(appfiat.FieldDisplayNames, field.TypeJSON, value)
	}
	if value, ok := afu.mutation.AppendedDisplayNames(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, appfiat.FieldDisplayNames, value)
		})
	}
	if afu.mutation.DisplayNamesCleared() {
		_spec.ClearField(appfiat.FieldDisplayNames, field.TypeJSON)
	}
	if value, ok := afu.mutation.Logo(); ok {
		_spec.SetField(appfiat.FieldLogo, field.TypeString, value)
	}
	if afu.mutation.LogoCleared() {
		_spec.ClearField(appfiat.FieldLogo, field.TypeString)
	}
	if value, ok := afu.mutation.Disabled(); ok {
		_spec.SetField(appfiat.FieldDisabled, field.TypeBool, value)
	}
	if afu.mutation.DisabledCleared() {
		_spec.ClearField(appfiat.FieldDisabled, field.TypeBool)
	}
	if value, ok := afu.mutation.Display(); ok {
		_spec.SetField(appfiat.FieldDisplay, field.TypeBool, value)
	}
	if afu.mutation.DisplayCleared() {
		_spec.ClearField(appfiat.FieldDisplay, field.TypeBool)
	}
	if value, ok := afu.mutation.DisplayIndex(); ok {
		_spec.SetField(appfiat.FieldDisplayIndex, field.TypeUint32, value)
	}
	if value, ok := afu.mutation.AddedDisplayIndex(); ok {
		_spec.AddField(appfiat.FieldDisplayIndex, field.TypeUint32, value)
	}
	if afu.mutation.DisplayIndexCleared() {
		_spec.ClearField(appfiat.FieldDisplayIndex, field.TypeUint32)
	}
	_spec.AddModifiers(afu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, afu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appfiat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	afu.mutation.done = true
	return n, nil
}

// AppFiatUpdateOne is the builder for updating a single AppFiat entity.
type AppFiatUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AppFiatMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (afuo *AppFiatUpdateOne) SetCreatedAt(u uint32) *AppFiatUpdateOne {
	afuo.mutation.ResetCreatedAt()
	afuo.mutation.SetCreatedAt(u)
	return afuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (afuo *AppFiatUpdateOne) SetNillableCreatedAt(u *uint32) *AppFiatUpdateOne {
	if u != nil {
		afuo.SetCreatedAt(*u)
	}
	return afuo
}

// AddCreatedAt adds u to the "created_at" field.
func (afuo *AppFiatUpdateOne) AddCreatedAt(u int32) *AppFiatUpdateOne {
	afuo.mutation.AddCreatedAt(u)
	return afuo
}

// SetUpdatedAt sets the "updated_at" field.
func (afuo *AppFiatUpdateOne) SetUpdatedAt(u uint32) *AppFiatUpdateOne {
	afuo.mutation.ResetUpdatedAt()
	afuo.mutation.SetUpdatedAt(u)
	return afuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (afuo *AppFiatUpdateOne) AddUpdatedAt(u int32) *AppFiatUpdateOne {
	afuo.mutation.AddUpdatedAt(u)
	return afuo
}

// SetDeletedAt sets the "deleted_at" field.
func (afuo *AppFiatUpdateOne) SetDeletedAt(u uint32) *AppFiatUpdateOne {
	afuo.mutation.ResetDeletedAt()
	afuo.mutation.SetDeletedAt(u)
	return afuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (afuo *AppFiatUpdateOne) SetNillableDeletedAt(u *uint32) *AppFiatUpdateOne {
	if u != nil {
		afuo.SetDeletedAt(*u)
	}
	return afuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (afuo *AppFiatUpdateOne) AddDeletedAt(u int32) *AppFiatUpdateOne {
	afuo.mutation.AddDeletedAt(u)
	return afuo
}

// SetEntID sets the "ent_id" field.
func (afuo *AppFiatUpdateOne) SetEntID(u uuid.UUID) *AppFiatUpdateOne {
	afuo.mutation.SetEntID(u)
	return afuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (afuo *AppFiatUpdateOne) SetNillableEntID(u *uuid.UUID) *AppFiatUpdateOne {
	if u != nil {
		afuo.SetEntID(*u)
	}
	return afuo
}

// SetAppID sets the "app_id" field.
func (afuo *AppFiatUpdateOne) SetAppID(u uuid.UUID) *AppFiatUpdateOne {
	afuo.mutation.SetAppID(u)
	return afuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (afuo *AppFiatUpdateOne) SetNillableAppID(u *uuid.UUID) *AppFiatUpdateOne {
	if u != nil {
		afuo.SetAppID(*u)
	}
	return afuo
}

// ClearAppID clears the value of the "app_id" field.
func (afuo *AppFiatUpdateOne) ClearAppID() *AppFiatUpdateOne {
	afuo.mutation.ClearAppID()
	return afuo
}

// SetFiatID sets the "fiat_id" field.
func (afuo *AppFiatUpdateOne) SetFiatID(u uuid.UUID) *AppFiatUpdateOne {
	afuo.mutation.SetFiatID(u)
	return afuo
}

// SetNillableFiatID sets the "fiat_id" field if the given value is not nil.
func (afuo *AppFiatUpdateOne) SetNillableFiatID(u *uuid.UUID) *AppFiatUpdateOne {
	if u != nil {
		afuo.SetFiatID(*u)
	}
	return afuo
}

// ClearFiatID clears the value of the "fiat_id" field.
func (afuo *AppFiatUpdateOne) ClearFiatID() *AppFiatUpdateOne {
	afuo.mutation.ClearFiatID()
	return afuo
}

// SetName sets the "name" field.
func (afuo *AppFiatUpdateOne) SetName(s string) *AppFiatUpdateOne {
	afuo.mutation.SetName(s)
	return afuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (afuo *AppFiatUpdateOne) SetNillableName(s *string) *AppFiatUpdateOne {
	if s != nil {
		afuo.SetName(*s)
	}
	return afuo
}

// ClearName clears the value of the "name" field.
func (afuo *AppFiatUpdateOne) ClearName() *AppFiatUpdateOne {
	afuo.mutation.ClearName()
	return afuo
}

// SetDisplayNames sets the "display_names" field.
func (afuo *AppFiatUpdateOne) SetDisplayNames(s []string) *AppFiatUpdateOne {
	afuo.mutation.SetDisplayNames(s)
	return afuo
}

// AppendDisplayNames appends s to the "display_names" field.
func (afuo *AppFiatUpdateOne) AppendDisplayNames(s []string) *AppFiatUpdateOne {
	afuo.mutation.AppendDisplayNames(s)
	return afuo
}

// ClearDisplayNames clears the value of the "display_names" field.
func (afuo *AppFiatUpdateOne) ClearDisplayNames() *AppFiatUpdateOne {
	afuo.mutation.ClearDisplayNames()
	return afuo
}

// SetLogo sets the "logo" field.
func (afuo *AppFiatUpdateOne) SetLogo(s string) *AppFiatUpdateOne {
	afuo.mutation.SetLogo(s)
	return afuo
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (afuo *AppFiatUpdateOne) SetNillableLogo(s *string) *AppFiatUpdateOne {
	if s != nil {
		afuo.SetLogo(*s)
	}
	return afuo
}

// ClearLogo clears the value of the "logo" field.
func (afuo *AppFiatUpdateOne) ClearLogo() *AppFiatUpdateOne {
	afuo.mutation.ClearLogo()
	return afuo
}

// SetDisabled sets the "disabled" field.
func (afuo *AppFiatUpdateOne) SetDisabled(b bool) *AppFiatUpdateOne {
	afuo.mutation.SetDisabled(b)
	return afuo
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (afuo *AppFiatUpdateOne) SetNillableDisabled(b *bool) *AppFiatUpdateOne {
	if b != nil {
		afuo.SetDisabled(*b)
	}
	return afuo
}

// ClearDisabled clears the value of the "disabled" field.
func (afuo *AppFiatUpdateOne) ClearDisabled() *AppFiatUpdateOne {
	afuo.mutation.ClearDisabled()
	return afuo
}

// SetDisplay sets the "display" field.
func (afuo *AppFiatUpdateOne) SetDisplay(b bool) *AppFiatUpdateOne {
	afuo.mutation.SetDisplay(b)
	return afuo
}

// SetNillableDisplay sets the "display" field if the given value is not nil.
func (afuo *AppFiatUpdateOne) SetNillableDisplay(b *bool) *AppFiatUpdateOne {
	if b != nil {
		afuo.SetDisplay(*b)
	}
	return afuo
}

// ClearDisplay clears the value of the "display" field.
func (afuo *AppFiatUpdateOne) ClearDisplay() *AppFiatUpdateOne {
	afuo.mutation.ClearDisplay()
	return afuo
}

// SetDisplayIndex sets the "display_index" field.
func (afuo *AppFiatUpdateOne) SetDisplayIndex(u uint32) *AppFiatUpdateOne {
	afuo.mutation.ResetDisplayIndex()
	afuo.mutation.SetDisplayIndex(u)
	return afuo
}

// SetNillableDisplayIndex sets the "display_index" field if the given value is not nil.
func (afuo *AppFiatUpdateOne) SetNillableDisplayIndex(u *uint32) *AppFiatUpdateOne {
	if u != nil {
		afuo.SetDisplayIndex(*u)
	}
	return afuo
}

// AddDisplayIndex adds u to the "display_index" field.
func (afuo *AppFiatUpdateOne) AddDisplayIndex(u int32) *AppFiatUpdateOne {
	afuo.mutation.AddDisplayIndex(u)
	return afuo
}

// ClearDisplayIndex clears the value of the "display_index" field.
func (afuo *AppFiatUpdateOne) ClearDisplayIndex() *AppFiatUpdateOne {
	afuo.mutation.ClearDisplayIndex()
	return afuo
}

// Mutation returns the AppFiatMutation object of the builder.
func (afuo *AppFiatUpdateOne) Mutation() *AppFiatMutation {
	return afuo.mutation
}

// Where appends a list predicates to the AppFiatUpdate builder.
func (afuo *AppFiatUpdateOne) Where(ps ...predicate.AppFiat) *AppFiatUpdateOne {
	afuo.mutation.Where(ps...)
	return afuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (afuo *AppFiatUpdateOne) Select(field string, fields ...string) *AppFiatUpdateOne {
	afuo.fields = append([]string{field}, fields...)
	return afuo
}

// Save executes the query and returns the updated AppFiat entity.
func (afuo *AppFiatUpdateOne) Save(ctx context.Context) (*AppFiat, error) {
	afuo.defaults()
	return withHooks(ctx, afuo.sqlSave, afuo.mutation, afuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (afuo *AppFiatUpdateOne) SaveX(ctx context.Context) *AppFiat {
	node, err := afuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (afuo *AppFiatUpdateOne) Exec(ctx context.Context) error {
	_, err := afuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (afuo *AppFiatUpdateOne) ExecX(ctx context.Context) {
	if err := afuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (afuo *AppFiatUpdateOne) defaults() {
	if _, ok := afuo.mutation.UpdatedAt(); !ok {
		v := appfiat.UpdateDefaultUpdatedAt()
		afuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (afuo *AppFiatUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppFiatUpdateOne {
	afuo.modifiers = append(afuo.modifiers, modifiers...)
	return afuo
}

func (afuo *AppFiatUpdateOne) sqlSave(ctx context.Context) (_node *AppFiat, err error) {
	_spec := sqlgraph.NewUpdateSpec(appfiat.Table, appfiat.Columns, sqlgraph.NewFieldSpec(appfiat.FieldID, field.TypeUint32))
	id, ok := afuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "AppFiat.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := afuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appfiat.FieldID)
		for _, f := range fields {
			if !appfiat.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != appfiat.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := afuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := afuo.mutation.CreatedAt(); ok {
		_spec.SetField(appfiat.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := afuo.mutation.AddedCreatedAt(); ok {
		_spec.AddField(appfiat.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := afuo.mutation.UpdatedAt(); ok {
		_spec.SetField(appfiat.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := afuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(appfiat.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := afuo.mutation.DeletedAt(); ok {
		_spec.SetField(appfiat.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := afuo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(appfiat.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := afuo.mutation.EntID(); ok {
		_spec.SetField(appfiat.FieldEntID, field.TypeUUID, value)
	}
	if value, ok := afuo.mutation.AppID(); ok {
		_spec.SetField(appfiat.FieldAppID, field.TypeUUID, value)
	}
	if afuo.mutation.AppIDCleared() {
		_spec.ClearField(appfiat.FieldAppID, field.TypeUUID)
	}
	if value, ok := afuo.mutation.FiatID(); ok {
		_spec.SetField(appfiat.FieldFiatID, field.TypeUUID, value)
	}
	if afuo.mutation.FiatIDCleared() {
		_spec.ClearField(appfiat.FieldFiatID, field.TypeUUID)
	}
	if value, ok := afuo.mutation.Name(); ok {
		_spec.SetField(appfiat.FieldName, field.TypeString, value)
	}
	if afuo.mutation.NameCleared() {
		_spec.ClearField(appfiat.FieldName, field.TypeString)
	}
	if value, ok := afuo.mutation.DisplayNames(); ok {
		_spec.SetField(appfiat.FieldDisplayNames, field.TypeJSON, value)
	}
	if value, ok := afuo.mutation.AppendedDisplayNames(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, appfiat.FieldDisplayNames, value)
		})
	}
	if afuo.mutation.DisplayNamesCleared() {
		_spec.ClearField(appfiat.FieldDisplayNames, field.TypeJSON)
	}
	if value, ok := afuo.mutation.Logo(); ok {
		_spec.SetField(appfiat.FieldLogo, field.TypeString, value)
	}
	if afuo.mutation.LogoCleared() {
		_spec.ClearField(appfiat.FieldLogo, field.TypeString)
	}
	if value, ok := afuo.mutation.Disabled(); ok {
		_spec.SetField(appfiat.FieldDisabled, field.TypeBool, value)
	}
	if afuo.mutation.DisabledCleared() {
		_spec.ClearField(appfiat.FieldDisabled, field.TypeBool)
	}
	if value, ok := afuo.mutation.Display(); ok {
		_spec.SetField(appfiat.FieldDisplay, field.TypeBool, value)
	}
	if afuo.mutation.DisplayCleared() {
		_spec.ClearField(appfiat.FieldDisplay, field.TypeBool)
	}
	if value, ok := afuo.mutation.DisplayIndex(); ok {
		_spec.SetField(appfiat.FieldDisplayIndex, field.TypeUint32, value)
	}
	if value, ok := afuo.mutation.AddedDisplayIndex(); ok {
		_spec.AddField(appfiat.FieldDisplayIndex, field.TypeUint32, value)
	}
	if afuo.mutation.DisplayIndexCleared() {
		_spec.ClearField(appfiat.FieldDisplayIndex, field.TypeUint32)
	}
	_spec.AddModifiers(afuo.modifiers...)
	_node = &AppFiat{config: afuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, afuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appfiat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	afuo.mutation.done = true
	return _node, nil
}
