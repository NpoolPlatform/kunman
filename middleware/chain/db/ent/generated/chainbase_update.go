// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/chainbase"
	"github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/predicate"
	"github.com/google/uuid"
)

// ChainBaseUpdate is the builder for updating ChainBase entities.
type ChainBaseUpdate struct {
	config
	hooks     []Hook
	mutation  *ChainBaseMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ChainBaseUpdate builder.
func (cbu *ChainBaseUpdate) Where(ps ...predicate.ChainBase) *ChainBaseUpdate {
	cbu.mutation.Where(ps...)
	return cbu
}

// SetCreatedAt sets the "created_at" field.
func (cbu *ChainBaseUpdate) SetCreatedAt(u uint32) *ChainBaseUpdate {
	cbu.mutation.ResetCreatedAt()
	cbu.mutation.SetCreatedAt(u)
	return cbu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cbu *ChainBaseUpdate) SetNillableCreatedAt(u *uint32) *ChainBaseUpdate {
	if u != nil {
		cbu.SetCreatedAt(*u)
	}
	return cbu
}

// AddCreatedAt adds u to the "created_at" field.
func (cbu *ChainBaseUpdate) AddCreatedAt(u int32) *ChainBaseUpdate {
	cbu.mutation.AddCreatedAt(u)
	return cbu
}

// SetUpdatedAt sets the "updated_at" field.
func (cbu *ChainBaseUpdate) SetUpdatedAt(u uint32) *ChainBaseUpdate {
	cbu.mutation.ResetUpdatedAt()
	cbu.mutation.SetUpdatedAt(u)
	return cbu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cbu *ChainBaseUpdate) AddUpdatedAt(u int32) *ChainBaseUpdate {
	cbu.mutation.AddUpdatedAt(u)
	return cbu
}

// SetDeletedAt sets the "deleted_at" field.
func (cbu *ChainBaseUpdate) SetDeletedAt(u uint32) *ChainBaseUpdate {
	cbu.mutation.ResetDeletedAt()
	cbu.mutation.SetDeletedAt(u)
	return cbu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cbu *ChainBaseUpdate) SetNillableDeletedAt(u *uint32) *ChainBaseUpdate {
	if u != nil {
		cbu.SetDeletedAt(*u)
	}
	return cbu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cbu *ChainBaseUpdate) AddDeletedAt(u int32) *ChainBaseUpdate {
	cbu.mutation.AddDeletedAt(u)
	return cbu
}

// SetEntID sets the "ent_id" field.
func (cbu *ChainBaseUpdate) SetEntID(u uuid.UUID) *ChainBaseUpdate {
	cbu.mutation.SetEntID(u)
	return cbu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cbu *ChainBaseUpdate) SetNillableEntID(u *uuid.UUID) *ChainBaseUpdate {
	if u != nil {
		cbu.SetEntID(*u)
	}
	return cbu
}

// SetName sets the "name" field.
func (cbu *ChainBaseUpdate) SetName(s string) *ChainBaseUpdate {
	cbu.mutation.SetName(s)
	return cbu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cbu *ChainBaseUpdate) SetNillableName(s *string) *ChainBaseUpdate {
	if s != nil {
		cbu.SetName(*s)
	}
	return cbu
}

// ClearName clears the value of the "name" field.
func (cbu *ChainBaseUpdate) ClearName() *ChainBaseUpdate {
	cbu.mutation.ClearName()
	return cbu
}

// SetLogo sets the "logo" field.
func (cbu *ChainBaseUpdate) SetLogo(s string) *ChainBaseUpdate {
	cbu.mutation.SetLogo(s)
	return cbu
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (cbu *ChainBaseUpdate) SetNillableLogo(s *string) *ChainBaseUpdate {
	if s != nil {
		cbu.SetLogo(*s)
	}
	return cbu
}

// ClearLogo clears the value of the "logo" field.
func (cbu *ChainBaseUpdate) ClearLogo() *ChainBaseUpdate {
	cbu.mutation.ClearLogo()
	return cbu
}

// SetNativeUnit sets the "native_unit" field.
func (cbu *ChainBaseUpdate) SetNativeUnit(s string) *ChainBaseUpdate {
	cbu.mutation.SetNativeUnit(s)
	return cbu
}

// SetNillableNativeUnit sets the "native_unit" field if the given value is not nil.
func (cbu *ChainBaseUpdate) SetNillableNativeUnit(s *string) *ChainBaseUpdate {
	if s != nil {
		cbu.SetNativeUnit(*s)
	}
	return cbu
}

// ClearNativeUnit clears the value of the "native_unit" field.
func (cbu *ChainBaseUpdate) ClearNativeUnit() *ChainBaseUpdate {
	cbu.mutation.ClearNativeUnit()
	return cbu
}

// SetAtomicUnit sets the "atomic_unit" field.
func (cbu *ChainBaseUpdate) SetAtomicUnit(s string) *ChainBaseUpdate {
	cbu.mutation.SetAtomicUnit(s)
	return cbu
}

// SetNillableAtomicUnit sets the "atomic_unit" field if the given value is not nil.
func (cbu *ChainBaseUpdate) SetNillableAtomicUnit(s *string) *ChainBaseUpdate {
	if s != nil {
		cbu.SetAtomicUnit(*s)
	}
	return cbu
}

// ClearAtomicUnit clears the value of the "atomic_unit" field.
func (cbu *ChainBaseUpdate) ClearAtomicUnit() *ChainBaseUpdate {
	cbu.mutation.ClearAtomicUnit()
	return cbu
}

// SetUnitExp sets the "unit_exp" field.
func (cbu *ChainBaseUpdate) SetUnitExp(u uint32) *ChainBaseUpdate {
	cbu.mutation.ResetUnitExp()
	cbu.mutation.SetUnitExp(u)
	return cbu
}

// SetNillableUnitExp sets the "unit_exp" field if the given value is not nil.
func (cbu *ChainBaseUpdate) SetNillableUnitExp(u *uint32) *ChainBaseUpdate {
	if u != nil {
		cbu.SetUnitExp(*u)
	}
	return cbu
}

// AddUnitExp adds u to the "unit_exp" field.
func (cbu *ChainBaseUpdate) AddUnitExp(u int32) *ChainBaseUpdate {
	cbu.mutation.AddUnitExp(u)
	return cbu
}

// ClearUnitExp clears the value of the "unit_exp" field.
func (cbu *ChainBaseUpdate) ClearUnitExp() *ChainBaseUpdate {
	cbu.mutation.ClearUnitExp()
	return cbu
}

// SetEnv sets the "env" field.
func (cbu *ChainBaseUpdate) SetEnv(s string) *ChainBaseUpdate {
	cbu.mutation.SetEnv(s)
	return cbu
}

// SetNillableEnv sets the "env" field if the given value is not nil.
func (cbu *ChainBaseUpdate) SetNillableEnv(s *string) *ChainBaseUpdate {
	if s != nil {
		cbu.SetEnv(*s)
	}
	return cbu
}

// ClearEnv clears the value of the "env" field.
func (cbu *ChainBaseUpdate) ClearEnv() *ChainBaseUpdate {
	cbu.mutation.ClearEnv()
	return cbu
}

// SetChainID sets the "chain_id" field.
func (cbu *ChainBaseUpdate) SetChainID(s string) *ChainBaseUpdate {
	cbu.mutation.SetChainID(s)
	return cbu
}

// SetNillableChainID sets the "chain_id" field if the given value is not nil.
func (cbu *ChainBaseUpdate) SetNillableChainID(s *string) *ChainBaseUpdate {
	if s != nil {
		cbu.SetChainID(*s)
	}
	return cbu
}

// ClearChainID clears the value of the "chain_id" field.
func (cbu *ChainBaseUpdate) ClearChainID() *ChainBaseUpdate {
	cbu.mutation.ClearChainID()
	return cbu
}

// SetNickname sets the "nickname" field.
func (cbu *ChainBaseUpdate) SetNickname(s string) *ChainBaseUpdate {
	cbu.mutation.SetNickname(s)
	return cbu
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (cbu *ChainBaseUpdate) SetNillableNickname(s *string) *ChainBaseUpdate {
	if s != nil {
		cbu.SetNickname(*s)
	}
	return cbu
}

// ClearNickname clears the value of the "nickname" field.
func (cbu *ChainBaseUpdate) ClearNickname() *ChainBaseUpdate {
	cbu.mutation.ClearNickname()
	return cbu
}

// SetGasType sets the "gas_type" field.
func (cbu *ChainBaseUpdate) SetGasType(s string) *ChainBaseUpdate {
	cbu.mutation.SetGasType(s)
	return cbu
}

// SetNillableGasType sets the "gas_type" field if the given value is not nil.
func (cbu *ChainBaseUpdate) SetNillableGasType(s *string) *ChainBaseUpdate {
	if s != nil {
		cbu.SetGasType(*s)
	}
	return cbu
}

// ClearGasType clears the value of the "gas_type" field.
func (cbu *ChainBaseUpdate) ClearGasType() *ChainBaseUpdate {
	cbu.mutation.ClearGasType()
	return cbu
}

// Mutation returns the ChainBaseMutation object of the builder.
func (cbu *ChainBaseUpdate) Mutation() *ChainBaseMutation {
	return cbu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cbu *ChainBaseUpdate) Save(ctx context.Context) (int, error) {
	cbu.defaults()
	return withHooks(ctx, cbu.sqlSave, cbu.mutation, cbu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cbu *ChainBaseUpdate) SaveX(ctx context.Context) int {
	affected, err := cbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cbu *ChainBaseUpdate) Exec(ctx context.Context) error {
	_, err := cbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cbu *ChainBaseUpdate) ExecX(ctx context.Context) {
	if err := cbu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cbu *ChainBaseUpdate) defaults() {
	if _, ok := cbu.mutation.UpdatedAt(); !ok {
		v := chainbase.UpdateDefaultUpdatedAt()
		cbu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cbu *ChainBaseUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ChainBaseUpdate {
	cbu.modifiers = append(cbu.modifiers, modifiers...)
	return cbu
}

func (cbu *ChainBaseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(chainbase.Table, chainbase.Columns, sqlgraph.NewFieldSpec(chainbase.FieldID, field.TypeUint32))
	if ps := cbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cbu.mutation.CreatedAt(); ok {
		_spec.SetField(chainbase.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := cbu.mutation.AddedCreatedAt(); ok {
		_spec.AddField(chainbase.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := cbu.mutation.UpdatedAt(); ok {
		_spec.SetField(chainbase.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := cbu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(chainbase.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := cbu.mutation.DeletedAt(); ok {
		_spec.SetField(chainbase.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := cbu.mutation.AddedDeletedAt(); ok {
		_spec.AddField(chainbase.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := cbu.mutation.EntID(); ok {
		_spec.SetField(chainbase.FieldEntID, field.TypeUUID, value)
	}
	if value, ok := cbu.mutation.Name(); ok {
		_spec.SetField(chainbase.FieldName, field.TypeString, value)
	}
	if cbu.mutation.NameCleared() {
		_spec.ClearField(chainbase.FieldName, field.TypeString)
	}
	if value, ok := cbu.mutation.Logo(); ok {
		_spec.SetField(chainbase.FieldLogo, field.TypeString, value)
	}
	if cbu.mutation.LogoCleared() {
		_spec.ClearField(chainbase.FieldLogo, field.TypeString)
	}
	if value, ok := cbu.mutation.NativeUnit(); ok {
		_spec.SetField(chainbase.FieldNativeUnit, field.TypeString, value)
	}
	if cbu.mutation.NativeUnitCleared() {
		_spec.ClearField(chainbase.FieldNativeUnit, field.TypeString)
	}
	if value, ok := cbu.mutation.AtomicUnit(); ok {
		_spec.SetField(chainbase.FieldAtomicUnit, field.TypeString, value)
	}
	if cbu.mutation.AtomicUnitCleared() {
		_spec.ClearField(chainbase.FieldAtomicUnit, field.TypeString)
	}
	if value, ok := cbu.mutation.UnitExp(); ok {
		_spec.SetField(chainbase.FieldUnitExp, field.TypeUint32, value)
	}
	if value, ok := cbu.mutation.AddedUnitExp(); ok {
		_spec.AddField(chainbase.FieldUnitExp, field.TypeUint32, value)
	}
	if cbu.mutation.UnitExpCleared() {
		_spec.ClearField(chainbase.FieldUnitExp, field.TypeUint32)
	}
	if value, ok := cbu.mutation.Env(); ok {
		_spec.SetField(chainbase.FieldEnv, field.TypeString, value)
	}
	if cbu.mutation.EnvCleared() {
		_spec.ClearField(chainbase.FieldEnv, field.TypeString)
	}
	if value, ok := cbu.mutation.ChainID(); ok {
		_spec.SetField(chainbase.FieldChainID, field.TypeString, value)
	}
	if cbu.mutation.ChainIDCleared() {
		_spec.ClearField(chainbase.FieldChainID, field.TypeString)
	}
	if value, ok := cbu.mutation.Nickname(); ok {
		_spec.SetField(chainbase.FieldNickname, field.TypeString, value)
	}
	if cbu.mutation.NicknameCleared() {
		_spec.ClearField(chainbase.FieldNickname, field.TypeString)
	}
	if value, ok := cbu.mutation.GasType(); ok {
		_spec.SetField(chainbase.FieldGasType, field.TypeString, value)
	}
	if cbu.mutation.GasTypeCleared() {
		_spec.ClearField(chainbase.FieldGasType, field.TypeString)
	}
	_spec.AddModifiers(cbu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chainbase.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cbu.mutation.done = true
	return n, nil
}

// ChainBaseUpdateOne is the builder for updating a single ChainBase entity.
type ChainBaseUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ChainBaseMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (cbuo *ChainBaseUpdateOne) SetCreatedAt(u uint32) *ChainBaseUpdateOne {
	cbuo.mutation.ResetCreatedAt()
	cbuo.mutation.SetCreatedAt(u)
	return cbuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cbuo *ChainBaseUpdateOne) SetNillableCreatedAt(u *uint32) *ChainBaseUpdateOne {
	if u != nil {
		cbuo.SetCreatedAt(*u)
	}
	return cbuo
}

// AddCreatedAt adds u to the "created_at" field.
func (cbuo *ChainBaseUpdateOne) AddCreatedAt(u int32) *ChainBaseUpdateOne {
	cbuo.mutation.AddCreatedAt(u)
	return cbuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cbuo *ChainBaseUpdateOne) SetUpdatedAt(u uint32) *ChainBaseUpdateOne {
	cbuo.mutation.ResetUpdatedAt()
	cbuo.mutation.SetUpdatedAt(u)
	return cbuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cbuo *ChainBaseUpdateOne) AddUpdatedAt(u int32) *ChainBaseUpdateOne {
	cbuo.mutation.AddUpdatedAt(u)
	return cbuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cbuo *ChainBaseUpdateOne) SetDeletedAt(u uint32) *ChainBaseUpdateOne {
	cbuo.mutation.ResetDeletedAt()
	cbuo.mutation.SetDeletedAt(u)
	return cbuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cbuo *ChainBaseUpdateOne) SetNillableDeletedAt(u *uint32) *ChainBaseUpdateOne {
	if u != nil {
		cbuo.SetDeletedAt(*u)
	}
	return cbuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cbuo *ChainBaseUpdateOne) AddDeletedAt(u int32) *ChainBaseUpdateOne {
	cbuo.mutation.AddDeletedAt(u)
	return cbuo
}

// SetEntID sets the "ent_id" field.
func (cbuo *ChainBaseUpdateOne) SetEntID(u uuid.UUID) *ChainBaseUpdateOne {
	cbuo.mutation.SetEntID(u)
	return cbuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cbuo *ChainBaseUpdateOne) SetNillableEntID(u *uuid.UUID) *ChainBaseUpdateOne {
	if u != nil {
		cbuo.SetEntID(*u)
	}
	return cbuo
}

// SetName sets the "name" field.
func (cbuo *ChainBaseUpdateOne) SetName(s string) *ChainBaseUpdateOne {
	cbuo.mutation.SetName(s)
	return cbuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cbuo *ChainBaseUpdateOne) SetNillableName(s *string) *ChainBaseUpdateOne {
	if s != nil {
		cbuo.SetName(*s)
	}
	return cbuo
}

// ClearName clears the value of the "name" field.
func (cbuo *ChainBaseUpdateOne) ClearName() *ChainBaseUpdateOne {
	cbuo.mutation.ClearName()
	return cbuo
}

// SetLogo sets the "logo" field.
func (cbuo *ChainBaseUpdateOne) SetLogo(s string) *ChainBaseUpdateOne {
	cbuo.mutation.SetLogo(s)
	return cbuo
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (cbuo *ChainBaseUpdateOne) SetNillableLogo(s *string) *ChainBaseUpdateOne {
	if s != nil {
		cbuo.SetLogo(*s)
	}
	return cbuo
}

// ClearLogo clears the value of the "logo" field.
func (cbuo *ChainBaseUpdateOne) ClearLogo() *ChainBaseUpdateOne {
	cbuo.mutation.ClearLogo()
	return cbuo
}

// SetNativeUnit sets the "native_unit" field.
func (cbuo *ChainBaseUpdateOne) SetNativeUnit(s string) *ChainBaseUpdateOne {
	cbuo.mutation.SetNativeUnit(s)
	return cbuo
}

// SetNillableNativeUnit sets the "native_unit" field if the given value is not nil.
func (cbuo *ChainBaseUpdateOne) SetNillableNativeUnit(s *string) *ChainBaseUpdateOne {
	if s != nil {
		cbuo.SetNativeUnit(*s)
	}
	return cbuo
}

// ClearNativeUnit clears the value of the "native_unit" field.
func (cbuo *ChainBaseUpdateOne) ClearNativeUnit() *ChainBaseUpdateOne {
	cbuo.mutation.ClearNativeUnit()
	return cbuo
}

// SetAtomicUnit sets the "atomic_unit" field.
func (cbuo *ChainBaseUpdateOne) SetAtomicUnit(s string) *ChainBaseUpdateOne {
	cbuo.mutation.SetAtomicUnit(s)
	return cbuo
}

// SetNillableAtomicUnit sets the "atomic_unit" field if the given value is not nil.
func (cbuo *ChainBaseUpdateOne) SetNillableAtomicUnit(s *string) *ChainBaseUpdateOne {
	if s != nil {
		cbuo.SetAtomicUnit(*s)
	}
	return cbuo
}

// ClearAtomicUnit clears the value of the "atomic_unit" field.
func (cbuo *ChainBaseUpdateOne) ClearAtomicUnit() *ChainBaseUpdateOne {
	cbuo.mutation.ClearAtomicUnit()
	return cbuo
}

// SetUnitExp sets the "unit_exp" field.
func (cbuo *ChainBaseUpdateOne) SetUnitExp(u uint32) *ChainBaseUpdateOne {
	cbuo.mutation.ResetUnitExp()
	cbuo.mutation.SetUnitExp(u)
	return cbuo
}

// SetNillableUnitExp sets the "unit_exp" field if the given value is not nil.
func (cbuo *ChainBaseUpdateOne) SetNillableUnitExp(u *uint32) *ChainBaseUpdateOne {
	if u != nil {
		cbuo.SetUnitExp(*u)
	}
	return cbuo
}

// AddUnitExp adds u to the "unit_exp" field.
func (cbuo *ChainBaseUpdateOne) AddUnitExp(u int32) *ChainBaseUpdateOne {
	cbuo.mutation.AddUnitExp(u)
	return cbuo
}

// ClearUnitExp clears the value of the "unit_exp" field.
func (cbuo *ChainBaseUpdateOne) ClearUnitExp() *ChainBaseUpdateOne {
	cbuo.mutation.ClearUnitExp()
	return cbuo
}

// SetEnv sets the "env" field.
func (cbuo *ChainBaseUpdateOne) SetEnv(s string) *ChainBaseUpdateOne {
	cbuo.mutation.SetEnv(s)
	return cbuo
}

// SetNillableEnv sets the "env" field if the given value is not nil.
func (cbuo *ChainBaseUpdateOne) SetNillableEnv(s *string) *ChainBaseUpdateOne {
	if s != nil {
		cbuo.SetEnv(*s)
	}
	return cbuo
}

// ClearEnv clears the value of the "env" field.
func (cbuo *ChainBaseUpdateOne) ClearEnv() *ChainBaseUpdateOne {
	cbuo.mutation.ClearEnv()
	return cbuo
}

// SetChainID sets the "chain_id" field.
func (cbuo *ChainBaseUpdateOne) SetChainID(s string) *ChainBaseUpdateOne {
	cbuo.mutation.SetChainID(s)
	return cbuo
}

// SetNillableChainID sets the "chain_id" field if the given value is not nil.
func (cbuo *ChainBaseUpdateOne) SetNillableChainID(s *string) *ChainBaseUpdateOne {
	if s != nil {
		cbuo.SetChainID(*s)
	}
	return cbuo
}

// ClearChainID clears the value of the "chain_id" field.
func (cbuo *ChainBaseUpdateOne) ClearChainID() *ChainBaseUpdateOne {
	cbuo.mutation.ClearChainID()
	return cbuo
}

// SetNickname sets the "nickname" field.
func (cbuo *ChainBaseUpdateOne) SetNickname(s string) *ChainBaseUpdateOne {
	cbuo.mutation.SetNickname(s)
	return cbuo
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (cbuo *ChainBaseUpdateOne) SetNillableNickname(s *string) *ChainBaseUpdateOne {
	if s != nil {
		cbuo.SetNickname(*s)
	}
	return cbuo
}

// ClearNickname clears the value of the "nickname" field.
func (cbuo *ChainBaseUpdateOne) ClearNickname() *ChainBaseUpdateOne {
	cbuo.mutation.ClearNickname()
	return cbuo
}

// SetGasType sets the "gas_type" field.
func (cbuo *ChainBaseUpdateOne) SetGasType(s string) *ChainBaseUpdateOne {
	cbuo.mutation.SetGasType(s)
	return cbuo
}

// SetNillableGasType sets the "gas_type" field if the given value is not nil.
func (cbuo *ChainBaseUpdateOne) SetNillableGasType(s *string) *ChainBaseUpdateOne {
	if s != nil {
		cbuo.SetGasType(*s)
	}
	return cbuo
}

// ClearGasType clears the value of the "gas_type" field.
func (cbuo *ChainBaseUpdateOne) ClearGasType() *ChainBaseUpdateOne {
	cbuo.mutation.ClearGasType()
	return cbuo
}

// Mutation returns the ChainBaseMutation object of the builder.
func (cbuo *ChainBaseUpdateOne) Mutation() *ChainBaseMutation {
	return cbuo.mutation
}

// Where appends a list predicates to the ChainBaseUpdate builder.
func (cbuo *ChainBaseUpdateOne) Where(ps ...predicate.ChainBase) *ChainBaseUpdateOne {
	cbuo.mutation.Where(ps...)
	return cbuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cbuo *ChainBaseUpdateOne) Select(field string, fields ...string) *ChainBaseUpdateOne {
	cbuo.fields = append([]string{field}, fields...)
	return cbuo
}

// Save executes the query and returns the updated ChainBase entity.
func (cbuo *ChainBaseUpdateOne) Save(ctx context.Context) (*ChainBase, error) {
	cbuo.defaults()
	return withHooks(ctx, cbuo.sqlSave, cbuo.mutation, cbuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cbuo *ChainBaseUpdateOne) SaveX(ctx context.Context) *ChainBase {
	node, err := cbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cbuo *ChainBaseUpdateOne) Exec(ctx context.Context) error {
	_, err := cbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cbuo *ChainBaseUpdateOne) ExecX(ctx context.Context) {
	if err := cbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cbuo *ChainBaseUpdateOne) defaults() {
	if _, ok := cbuo.mutation.UpdatedAt(); !ok {
		v := chainbase.UpdateDefaultUpdatedAt()
		cbuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cbuo *ChainBaseUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ChainBaseUpdateOne {
	cbuo.modifiers = append(cbuo.modifiers, modifiers...)
	return cbuo
}

func (cbuo *ChainBaseUpdateOne) sqlSave(ctx context.Context) (_node *ChainBase, err error) {
	_spec := sqlgraph.NewUpdateSpec(chainbase.Table, chainbase.Columns, sqlgraph.NewFieldSpec(chainbase.FieldID, field.TypeUint32))
	id, ok := cbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "ChainBase.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chainbase.FieldID)
		for _, f := range fields {
			if !chainbase.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != chainbase.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cbuo.mutation.CreatedAt(); ok {
		_spec.SetField(chainbase.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := cbuo.mutation.AddedCreatedAt(); ok {
		_spec.AddField(chainbase.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := cbuo.mutation.UpdatedAt(); ok {
		_spec.SetField(chainbase.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := cbuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(chainbase.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := cbuo.mutation.DeletedAt(); ok {
		_spec.SetField(chainbase.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := cbuo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(chainbase.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := cbuo.mutation.EntID(); ok {
		_spec.SetField(chainbase.FieldEntID, field.TypeUUID, value)
	}
	if value, ok := cbuo.mutation.Name(); ok {
		_spec.SetField(chainbase.FieldName, field.TypeString, value)
	}
	if cbuo.mutation.NameCleared() {
		_spec.ClearField(chainbase.FieldName, field.TypeString)
	}
	if value, ok := cbuo.mutation.Logo(); ok {
		_spec.SetField(chainbase.FieldLogo, field.TypeString, value)
	}
	if cbuo.mutation.LogoCleared() {
		_spec.ClearField(chainbase.FieldLogo, field.TypeString)
	}
	if value, ok := cbuo.mutation.NativeUnit(); ok {
		_spec.SetField(chainbase.FieldNativeUnit, field.TypeString, value)
	}
	if cbuo.mutation.NativeUnitCleared() {
		_spec.ClearField(chainbase.FieldNativeUnit, field.TypeString)
	}
	if value, ok := cbuo.mutation.AtomicUnit(); ok {
		_spec.SetField(chainbase.FieldAtomicUnit, field.TypeString, value)
	}
	if cbuo.mutation.AtomicUnitCleared() {
		_spec.ClearField(chainbase.FieldAtomicUnit, field.TypeString)
	}
	if value, ok := cbuo.mutation.UnitExp(); ok {
		_spec.SetField(chainbase.FieldUnitExp, field.TypeUint32, value)
	}
	if value, ok := cbuo.mutation.AddedUnitExp(); ok {
		_spec.AddField(chainbase.FieldUnitExp, field.TypeUint32, value)
	}
	if cbuo.mutation.UnitExpCleared() {
		_spec.ClearField(chainbase.FieldUnitExp, field.TypeUint32)
	}
	if value, ok := cbuo.mutation.Env(); ok {
		_spec.SetField(chainbase.FieldEnv, field.TypeString, value)
	}
	if cbuo.mutation.EnvCleared() {
		_spec.ClearField(chainbase.FieldEnv, field.TypeString)
	}
	if value, ok := cbuo.mutation.ChainID(); ok {
		_spec.SetField(chainbase.FieldChainID, field.TypeString, value)
	}
	if cbuo.mutation.ChainIDCleared() {
		_spec.ClearField(chainbase.FieldChainID, field.TypeString)
	}
	if value, ok := cbuo.mutation.Nickname(); ok {
		_spec.SetField(chainbase.FieldNickname, field.TypeString, value)
	}
	if cbuo.mutation.NicknameCleared() {
		_spec.ClearField(chainbase.FieldNickname, field.TypeString)
	}
	if value, ok := cbuo.mutation.GasType(); ok {
		_spec.SetField(chainbase.FieldGasType, field.TypeString, value)
	}
	if cbuo.mutation.GasTypeCleared() {
		_spec.ClearField(chainbase.FieldGasType, field.TypeString)
	}
	_spec.AddModifiers(cbuo.modifiers...)
	_node = &ChainBase{config: cbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chainbase.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cbuo.mutation.done = true
	return _node, nil
}
