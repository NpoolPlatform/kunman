// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/coinfiatcurrencyhistory"
	"github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// CoinFiatCurrencyHistoryUpdate is the builder for updating CoinFiatCurrencyHistory entities.
type CoinFiatCurrencyHistoryUpdate struct {
	config
	hooks     []Hook
	mutation  *CoinFiatCurrencyHistoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CoinFiatCurrencyHistoryUpdate builder.
func (cfchu *CoinFiatCurrencyHistoryUpdate) Where(ps ...predicate.CoinFiatCurrencyHistory) *CoinFiatCurrencyHistoryUpdate {
	cfchu.mutation.Where(ps...)
	return cfchu
}

// SetCreatedAt sets the "created_at" field.
func (cfchu *CoinFiatCurrencyHistoryUpdate) SetCreatedAt(u uint32) *CoinFiatCurrencyHistoryUpdate {
	cfchu.mutation.ResetCreatedAt()
	cfchu.mutation.SetCreatedAt(u)
	return cfchu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cfchu *CoinFiatCurrencyHistoryUpdate) SetNillableCreatedAt(u *uint32) *CoinFiatCurrencyHistoryUpdate {
	if u != nil {
		cfchu.SetCreatedAt(*u)
	}
	return cfchu
}

// AddCreatedAt adds u to the "created_at" field.
func (cfchu *CoinFiatCurrencyHistoryUpdate) AddCreatedAt(u int32) *CoinFiatCurrencyHistoryUpdate {
	cfchu.mutation.AddCreatedAt(u)
	return cfchu
}

// SetUpdatedAt sets the "updated_at" field.
func (cfchu *CoinFiatCurrencyHistoryUpdate) SetUpdatedAt(u uint32) *CoinFiatCurrencyHistoryUpdate {
	cfchu.mutation.ResetUpdatedAt()
	cfchu.mutation.SetUpdatedAt(u)
	return cfchu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cfchu *CoinFiatCurrencyHistoryUpdate) AddUpdatedAt(u int32) *CoinFiatCurrencyHistoryUpdate {
	cfchu.mutation.AddUpdatedAt(u)
	return cfchu
}

// SetDeletedAt sets the "deleted_at" field.
func (cfchu *CoinFiatCurrencyHistoryUpdate) SetDeletedAt(u uint32) *CoinFiatCurrencyHistoryUpdate {
	cfchu.mutation.ResetDeletedAt()
	cfchu.mutation.SetDeletedAt(u)
	return cfchu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cfchu *CoinFiatCurrencyHistoryUpdate) SetNillableDeletedAt(u *uint32) *CoinFiatCurrencyHistoryUpdate {
	if u != nil {
		cfchu.SetDeletedAt(*u)
	}
	return cfchu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cfchu *CoinFiatCurrencyHistoryUpdate) AddDeletedAt(u int32) *CoinFiatCurrencyHistoryUpdate {
	cfchu.mutation.AddDeletedAt(u)
	return cfchu
}

// SetEntID sets the "ent_id" field.
func (cfchu *CoinFiatCurrencyHistoryUpdate) SetEntID(u uuid.UUID) *CoinFiatCurrencyHistoryUpdate {
	cfchu.mutation.SetEntID(u)
	return cfchu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cfchu *CoinFiatCurrencyHistoryUpdate) SetNillableEntID(u *uuid.UUID) *CoinFiatCurrencyHistoryUpdate {
	if u != nil {
		cfchu.SetEntID(*u)
	}
	return cfchu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (cfchu *CoinFiatCurrencyHistoryUpdate) SetCoinTypeID(u uuid.UUID) *CoinFiatCurrencyHistoryUpdate {
	cfchu.mutation.SetCoinTypeID(u)
	return cfchu
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (cfchu *CoinFiatCurrencyHistoryUpdate) SetNillableCoinTypeID(u *uuid.UUID) *CoinFiatCurrencyHistoryUpdate {
	if u != nil {
		cfchu.SetCoinTypeID(*u)
	}
	return cfchu
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (cfchu *CoinFiatCurrencyHistoryUpdate) ClearCoinTypeID() *CoinFiatCurrencyHistoryUpdate {
	cfchu.mutation.ClearCoinTypeID()
	return cfchu
}

// SetFiatID sets the "fiat_id" field.
func (cfchu *CoinFiatCurrencyHistoryUpdate) SetFiatID(u uuid.UUID) *CoinFiatCurrencyHistoryUpdate {
	cfchu.mutation.SetFiatID(u)
	return cfchu
}

// SetNillableFiatID sets the "fiat_id" field if the given value is not nil.
func (cfchu *CoinFiatCurrencyHistoryUpdate) SetNillableFiatID(u *uuid.UUID) *CoinFiatCurrencyHistoryUpdate {
	if u != nil {
		cfchu.SetFiatID(*u)
	}
	return cfchu
}

// ClearFiatID clears the value of the "fiat_id" field.
func (cfchu *CoinFiatCurrencyHistoryUpdate) ClearFiatID() *CoinFiatCurrencyHistoryUpdate {
	cfchu.mutation.ClearFiatID()
	return cfchu
}

// SetFeedType sets the "feed_type" field.
func (cfchu *CoinFiatCurrencyHistoryUpdate) SetFeedType(s string) *CoinFiatCurrencyHistoryUpdate {
	cfchu.mutation.SetFeedType(s)
	return cfchu
}

// SetNillableFeedType sets the "feed_type" field if the given value is not nil.
func (cfchu *CoinFiatCurrencyHistoryUpdate) SetNillableFeedType(s *string) *CoinFiatCurrencyHistoryUpdate {
	if s != nil {
		cfchu.SetFeedType(*s)
	}
	return cfchu
}

// ClearFeedType clears the value of the "feed_type" field.
func (cfchu *CoinFiatCurrencyHistoryUpdate) ClearFeedType() *CoinFiatCurrencyHistoryUpdate {
	cfchu.mutation.ClearFeedType()
	return cfchu
}

// SetMarketValueLow sets the "market_value_low" field.
func (cfchu *CoinFiatCurrencyHistoryUpdate) SetMarketValueLow(d decimal.Decimal) *CoinFiatCurrencyHistoryUpdate {
	cfchu.mutation.SetMarketValueLow(d)
	return cfchu
}

// SetNillableMarketValueLow sets the "market_value_low" field if the given value is not nil.
func (cfchu *CoinFiatCurrencyHistoryUpdate) SetNillableMarketValueLow(d *decimal.Decimal) *CoinFiatCurrencyHistoryUpdate {
	if d != nil {
		cfchu.SetMarketValueLow(*d)
	}
	return cfchu
}

// ClearMarketValueLow clears the value of the "market_value_low" field.
func (cfchu *CoinFiatCurrencyHistoryUpdate) ClearMarketValueLow() *CoinFiatCurrencyHistoryUpdate {
	cfchu.mutation.ClearMarketValueLow()
	return cfchu
}

// SetMarketValueHigh sets the "market_value_high" field.
func (cfchu *CoinFiatCurrencyHistoryUpdate) SetMarketValueHigh(d decimal.Decimal) *CoinFiatCurrencyHistoryUpdate {
	cfchu.mutation.SetMarketValueHigh(d)
	return cfchu
}

// SetNillableMarketValueHigh sets the "market_value_high" field if the given value is not nil.
func (cfchu *CoinFiatCurrencyHistoryUpdate) SetNillableMarketValueHigh(d *decimal.Decimal) *CoinFiatCurrencyHistoryUpdate {
	if d != nil {
		cfchu.SetMarketValueHigh(*d)
	}
	return cfchu
}

// ClearMarketValueHigh clears the value of the "market_value_high" field.
func (cfchu *CoinFiatCurrencyHistoryUpdate) ClearMarketValueHigh() *CoinFiatCurrencyHistoryUpdate {
	cfchu.mutation.ClearMarketValueHigh()
	return cfchu
}

// Mutation returns the CoinFiatCurrencyHistoryMutation object of the builder.
func (cfchu *CoinFiatCurrencyHistoryUpdate) Mutation() *CoinFiatCurrencyHistoryMutation {
	return cfchu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cfchu *CoinFiatCurrencyHistoryUpdate) Save(ctx context.Context) (int, error) {
	cfchu.defaults()
	return withHooks(ctx, cfchu.sqlSave, cfchu.mutation, cfchu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cfchu *CoinFiatCurrencyHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := cfchu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cfchu *CoinFiatCurrencyHistoryUpdate) Exec(ctx context.Context) error {
	_, err := cfchu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cfchu *CoinFiatCurrencyHistoryUpdate) ExecX(ctx context.Context) {
	if err := cfchu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cfchu *CoinFiatCurrencyHistoryUpdate) defaults() {
	if _, ok := cfchu.mutation.UpdatedAt(); !ok {
		v := coinfiatcurrencyhistory.UpdateDefaultUpdatedAt()
		cfchu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cfchu *CoinFiatCurrencyHistoryUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CoinFiatCurrencyHistoryUpdate {
	cfchu.modifiers = append(cfchu.modifiers, modifiers...)
	return cfchu
}

func (cfchu *CoinFiatCurrencyHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(coinfiatcurrencyhistory.Table, coinfiatcurrencyhistory.Columns, sqlgraph.NewFieldSpec(coinfiatcurrencyhistory.FieldID, field.TypeUint32))
	if ps := cfchu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cfchu.mutation.CreatedAt(); ok {
		_spec.SetField(coinfiatcurrencyhistory.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := cfchu.mutation.AddedCreatedAt(); ok {
		_spec.AddField(coinfiatcurrencyhistory.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := cfchu.mutation.UpdatedAt(); ok {
		_spec.SetField(coinfiatcurrencyhistory.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := cfchu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(coinfiatcurrencyhistory.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := cfchu.mutation.DeletedAt(); ok {
		_spec.SetField(coinfiatcurrencyhistory.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := cfchu.mutation.AddedDeletedAt(); ok {
		_spec.AddField(coinfiatcurrencyhistory.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := cfchu.mutation.EntID(); ok {
		_spec.SetField(coinfiatcurrencyhistory.FieldEntID, field.TypeUUID, value)
	}
	if value, ok := cfchu.mutation.CoinTypeID(); ok {
		_spec.SetField(coinfiatcurrencyhistory.FieldCoinTypeID, field.TypeUUID, value)
	}
	if cfchu.mutation.CoinTypeIDCleared() {
		_spec.ClearField(coinfiatcurrencyhistory.FieldCoinTypeID, field.TypeUUID)
	}
	if value, ok := cfchu.mutation.FiatID(); ok {
		_spec.SetField(coinfiatcurrencyhistory.FieldFiatID, field.TypeUUID, value)
	}
	if cfchu.mutation.FiatIDCleared() {
		_spec.ClearField(coinfiatcurrencyhistory.FieldFiatID, field.TypeUUID)
	}
	if value, ok := cfchu.mutation.FeedType(); ok {
		_spec.SetField(coinfiatcurrencyhistory.FieldFeedType, field.TypeString, value)
	}
	if cfchu.mutation.FeedTypeCleared() {
		_spec.ClearField(coinfiatcurrencyhistory.FieldFeedType, field.TypeString)
	}
	if value, ok := cfchu.mutation.MarketValueLow(); ok {
		_spec.SetField(coinfiatcurrencyhistory.FieldMarketValueLow, field.TypeOther, value)
	}
	if cfchu.mutation.MarketValueLowCleared() {
		_spec.ClearField(coinfiatcurrencyhistory.FieldMarketValueLow, field.TypeOther)
	}
	if value, ok := cfchu.mutation.MarketValueHigh(); ok {
		_spec.SetField(coinfiatcurrencyhistory.FieldMarketValueHigh, field.TypeOther, value)
	}
	if cfchu.mutation.MarketValueHighCleared() {
		_spec.ClearField(coinfiatcurrencyhistory.FieldMarketValueHigh, field.TypeOther)
	}
	_spec.AddModifiers(cfchu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cfchu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coinfiatcurrencyhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cfchu.mutation.done = true
	return n, nil
}

// CoinFiatCurrencyHistoryUpdateOne is the builder for updating a single CoinFiatCurrencyHistory entity.
type CoinFiatCurrencyHistoryUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CoinFiatCurrencyHistoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) SetCreatedAt(u uint32) *CoinFiatCurrencyHistoryUpdateOne {
	cfchuo.mutation.ResetCreatedAt()
	cfchuo.mutation.SetCreatedAt(u)
	return cfchuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) SetNillableCreatedAt(u *uint32) *CoinFiatCurrencyHistoryUpdateOne {
	if u != nil {
		cfchuo.SetCreatedAt(*u)
	}
	return cfchuo
}

// AddCreatedAt adds u to the "created_at" field.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) AddCreatedAt(u int32) *CoinFiatCurrencyHistoryUpdateOne {
	cfchuo.mutation.AddCreatedAt(u)
	return cfchuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) SetUpdatedAt(u uint32) *CoinFiatCurrencyHistoryUpdateOne {
	cfchuo.mutation.ResetUpdatedAt()
	cfchuo.mutation.SetUpdatedAt(u)
	return cfchuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) AddUpdatedAt(u int32) *CoinFiatCurrencyHistoryUpdateOne {
	cfchuo.mutation.AddUpdatedAt(u)
	return cfchuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) SetDeletedAt(u uint32) *CoinFiatCurrencyHistoryUpdateOne {
	cfchuo.mutation.ResetDeletedAt()
	cfchuo.mutation.SetDeletedAt(u)
	return cfchuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) SetNillableDeletedAt(u *uint32) *CoinFiatCurrencyHistoryUpdateOne {
	if u != nil {
		cfchuo.SetDeletedAt(*u)
	}
	return cfchuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) AddDeletedAt(u int32) *CoinFiatCurrencyHistoryUpdateOne {
	cfchuo.mutation.AddDeletedAt(u)
	return cfchuo
}

// SetEntID sets the "ent_id" field.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) SetEntID(u uuid.UUID) *CoinFiatCurrencyHistoryUpdateOne {
	cfchuo.mutation.SetEntID(u)
	return cfchuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) SetNillableEntID(u *uuid.UUID) *CoinFiatCurrencyHistoryUpdateOne {
	if u != nil {
		cfchuo.SetEntID(*u)
	}
	return cfchuo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) SetCoinTypeID(u uuid.UUID) *CoinFiatCurrencyHistoryUpdateOne {
	cfchuo.mutation.SetCoinTypeID(u)
	return cfchuo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *CoinFiatCurrencyHistoryUpdateOne {
	if u != nil {
		cfchuo.SetCoinTypeID(*u)
	}
	return cfchuo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) ClearCoinTypeID() *CoinFiatCurrencyHistoryUpdateOne {
	cfchuo.mutation.ClearCoinTypeID()
	return cfchuo
}

// SetFiatID sets the "fiat_id" field.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) SetFiatID(u uuid.UUID) *CoinFiatCurrencyHistoryUpdateOne {
	cfchuo.mutation.SetFiatID(u)
	return cfchuo
}

// SetNillableFiatID sets the "fiat_id" field if the given value is not nil.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) SetNillableFiatID(u *uuid.UUID) *CoinFiatCurrencyHistoryUpdateOne {
	if u != nil {
		cfchuo.SetFiatID(*u)
	}
	return cfchuo
}

// ClearFiatID clears the value of the "fiat_id" field.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) ClearFiatID() *CoinFiatCurrencyHistoryUpdateOne {
	cfchuo.mutation.ClearFiatID()
	return cfchuo
}

// SetFeedType sets the "feed_type" field.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) SetFeedType(s string) *CoinFiatCurrencyHistoryUpdateOne {
	cfchuo.mutation.SetFeedType(s)
	return cfchuo
}

// SetNillableFeedType sets the "feed_type" field if the given value is not nil.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) SetNillableFeedType(s *string) *CoinFiatCurrencyHistoryUpdateOne {
	if s != nil {
		cfchuo.SetFeedType(*s)
	}
	return cfchuo
}

// ClearFeedType clears the value of the "feed_type" field.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) ClearFeedType() *CoinFiatCurrencyHistoryUpdateOne {
	cfchuo.mutation.ClearFeedType()
	return cfchuo
}

// SetMarketValueLow sets the "market_value_low" field.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) SetMarketValueLow(d decimal.Decimal) *CoinFiatCurrencyHistoryUpdateOne {
	cfchuo.mutation.SetMarketValueLow(d)
	return cfchuo
}

// SetNillableMarketValueLow sets the "market_value_low" field if the given value is not nil.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) SetNillableMarketValueLow(d *decimal.Decimal) *CoinFiatCurrencyHistoryUpdateOne {
	if d != nil {
		cfchuo.SetMarketValueLow(*d)
	}
	return cfchuo
}

// ClearMarketValueLow clears the value of the "market_value_low" field.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) ClearMarketValueLow() *CoinFiatCurrencyHistoryUpdateOne {
	cfchuo.mutation.ClearMarketValueLow()
	return cfchuo
}

// SetMarketValueHigh sets the "market_value_high" field.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) SetMarketValueHigh(d decimal.Decimal) *CoinFiatCurrencyHistoryUpdateOne {
	cfchuo.mutation.SetMarketValueHigh(d)
	return cfchuo
}

// SetNillableMarketValueHigh sets the "market_value_high" field if the given value is not nil.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) SetNillableMarketValueHigh(d *decimal.Decimal) *CoinFiatCurrencyHistoryUpdateOne {
	if d != nil {
		cfchuo.SetMarketValueHigh(*d)
	}
	return cfchuo
}

// ClearMarketValueHigh clears the value of the "market_value_high" field.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) ClearMarketValueHigh() *CoinFiatCurrencyHistoryUpdateOne {
	cfchuo.mutation.ClearMarketValueHigh()
	return cfchuo
}

// Mutation returns the CoinFiatCurrencyHistoryMutation object of the builder.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) Mutation() *CoinFiatCurrencyHistoryMutation {
	return cfchuo.mutation
}

// Where appends a list predicates to the CoinFiatCurrencyHistoryUpdate builder.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) Where(ps ...predicate.CoinFiatCurrencyHistory) *CoinFiatCurrencyHistoryUpdateOne {
	cfchuo.mutation.Where(ps...)
	return cfchuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) Select(field string, fields ...string) *CoinFiatCurrencyHistoryUpdateOne {
	cfchuo.fields = append([]string{field}, fields...)
	return cfchuo
}

// Save executes the query and returns the updated CoinFiatCurrencyHistory entity.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) Save(ctx context.Context) (*CoinFiatCurrencyHistory, error) {
	cfchuo.defaults()
	return withHooks(ctx, cfchuo.sqlSave, cfchuo.mutation, cfchuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) SaveX(ctx context.Context) *CoinFiatCurrencyHistory {
	node, err := cfchuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := cfchuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := cfchuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) defaults() {
	if _, ok := cfchuo.mutation.UpdatedAt(); !ok {
		v := coinfiatcurrencyhistory.UpdateDefaultUpdatedAt()
		cfchuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CoinFiatCurrencyHistoryUpdateOne {
	cfchuo.modifiers = append(cfchuo.modifiers, modifiers...)
	return cfchuo
}

func (cfchuo *CoinFiatCurrencyHistoryUpdateOne) sqlSave(ctx context.Context) (_node *CoinFiatCurrencyHistory, err error) {
	_spec := sqlgraph.NewUpdateSpec(coinfiatcurrencyhistory.Table, coinfiatcurrencyhistory.Columns, sqlgraph.NewFieldSpec(coinfiatcurrencyhistory.FieldID, field.TypeUint32))
	id, ok := cfchuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "CoinFiatCurrencyHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cfchuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coinfiatcurrencyhistory.FieldID)
		for _, f := range fields {
			if !coinfiatcurrencyhistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != coinfiatcurrencyhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cfchuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cfchuo.mutation.CreatedAt(); ok {
		_spec.SetField(coinfiatcurrencyhistory.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := cfchuo.mutation.AddedCreatedAt(); ok {
		_spec.AddField(coinfiatcurrencyhistory.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := cfchuo.mutation.UpdatedAt(); ok {
		_spec.SetField(coinfiatcurrencyhistory.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := cfchuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(coinfiatcurrencyhistory.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := cfchuo.mutation.DeletedAt(); ok {
		_spec.SetField(coinfiatcurrencyhistory.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := cfchuo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(coinfiatcurrencyhistory.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := cfchuo.mutation.EntID(); ok {
		_spec.SetField(coinfiatcurrencyhistory.FieldEntID, field.TypeUUID, value)
	}
	if value, ok := cfchuo.mutation.CoinTypeID(); ok {
		_spec.SetField(coinfiatcurrencyhistory.FieldCoinTypeID, field.TypeUUID, value)
	}
	if cfchuo.mutation.CoinTypeIDCleared() {
		_spec.ClearField(coinfiatcurrencyhistory.FieldCoinTypeID, field.TypeUUID)
	}
	if value, ok := cfchuo.mutation.FiatID(); ok {
		_spec.SetField(coinfiatcurrencyhistory.FieldFiatID, field.TypeUUID, value)
	}
	if cfchuo.mutation.FiatIDCleared() {
		_spec.ClearField(coinfiatcurrencyhistory.FieldFiatID, field.TypeUUID)
	}
	if value, ok := cfchuo.mutation.FeedType(); ok {
		_spec.SetField(coinfiatcurrencyhistory.FieldFeedType, field.TypeString, value)
	}
	if cfchuo.mutation.FeedTypeCleared() {
		_spec.ClearField(coinfiatcurrencyhistory.FieldFeedType, field.TypeString)
	}
	if value, ok := cfchuo.mutation.MarketValueLow(); ok {
		_spec.SetField(coinfiatcurrencyhistory.FieldMarketValueLow, field.TypeOther, value)
	}
	if cfchuo.mutation.MarketValueLowCleared() {
		_spec.ClearField(coinfiatcurrencyhistory.FieldMarketValueLow, field.TypeOther)
	}
	if value, ok := cfchuo.mutation.MarketValueHigh(); ok {
		_spec.SetField(coinfiatcurrencyhistory.FieldMarketValueHigh, field.TypeOther, value)
	}
	if cfchuo.mutation.MarketValueHighCleared() {
		_spec.ClearField(coinfiatcurrencyhistory.FieldMarketValueHigh, field.TypeOther)
	}
	_spec.AddModifiers(cfchuo.modifiers...)
	_node = &CoinFiatCurrencyHistory{config: cfchuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cfchuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coinfiatcurrencyhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cfchuo.mutation.done = true
	return _node, nil
}
