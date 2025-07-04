// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/fiatcurrencyfeed"
	"github.com/google/uuid"
)

// FiatCurrencyFeedCreate is the builder for creating a FiatCurrencyFeed entity.
type FiatCurrencyFeedCreate struct {
	config
	mutation *FiatCurrencyFeedMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (fcfc *FiatCurrencyFeedCreate) SetCreatedAt(u uint32) *FiatCurrencyFeedCreate {
	fcfc.mutation.SetCreatedAt(u)
	return fcfc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fcfc *FiatCurrencyFeedCreate) SetNillableCreatedAt(u *uint32) *FiatCurrencyFeedCreate {
	if u != nil {
		fcfc.SetCreatedAt(*u)
	}
	return fcfc
}

// SetUpdatedAt sets the "updated_at" field.
func (fcfc *FiatCurrencyFeedCreate) SetUpdatedAt(u uint32) *FiatCurrencyFeedCreate {
	fcfc.mutation.SetUpdatedAt(u)
	return fcfc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fcfc *FiatCurrencyFeedCreate) SetNillableUpdatedAt(u *uint32) *FiatCurrencyFeedCreate {
	if u != nil {
		fcfc.SetUpdatedAt(*u)
	}
	return fcfc
}

// SetDeletedAt sets the "deleted_at" field.
func (fcfc *FiatCurrencyFeedCreate) SetDeletedAt(u uint32) *FiatCurrencyFeedCreate {
	fcfc.mutation.SetDeletedAt(u)
	return fcfc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fcfc *FiatCurrencyFeedCreate) SetNillableDeletedAt(u *uint32) *FiatCurrencyFeedCreate {
	if u != nil {
		fcfc.SetDeletedAt(*u)
	}
	return fcfc
}

// SetEntID sets the "ent_id" field.
func (fcfc *FiatCurrencyFeedCreate) SetEntID(u uuid.UUID) *FiatCurrencyFeedCreate {
	fcfc.mutation.SetEntID(u)
	return fcfc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (fcfc *FiatCurrencyFeedCreate) SetNillableEntID(u *uuid.UUID) *FiatCurrencyFeedCreate {
	if u != nil {
		fcfc.SetEntID(*u)
	}
	return fcfc
}

// SetFiatID sets the "fiat_id" field.
func (fcfc *FiatCurrencyFeedCreate) SetFiatID(u uuid.UUID) *FiatCurrencyFeedCreate {
	fcfc.mutation.SetFiatID(u)
	return fcfc
}

// SetNillableFiatID sets the "fiat_id" field if the given value is not nil.
func (fcfc *FiatCurrencyFeedCreate) SetNillableFiatID(u *uuid.UUID) *FiatCurrencyFeedCreate {
	if u != nil {
		fcfc.SetFiatID(*u)
	}
	return fcfc
}

// SetFeedType sets the "feed_type" field.
func (fcfc *FiatCurrencyFeedCreate) SetFeedType(s string) *FiatCurrencyFeedCreate {
	fcfc.mutation.SetFeedType(s)
	return fcfc
}

// SetNillableFeedType sets the "feed_type" field if the given value is not nil.
func (fcfc *FiatCurrencyFeedCreate) SetNillableFeedType(s *string) *FiatCurrencyFeedCreate {
	if s != nil {
		fcfc.SetFeedType(*s)
	}
	return fcfc
}

// SetFeedFiatName sets the "feed_fiat_name" field.
func (fcfc *FiatCurrencyFeedCreate) SetFeedFiatName(s string) *FiatCurrencyFeedCreate {
	fcfc.mutation.SetFeedFiatName(s)
	return fcfc
}

// SetNillableFeedFiatName sets the "feed_fiat_name" field if the given value is not nil.
func (fcfc *FiatCurrencyFeedCreate) SetNillableFeedFiatName(s *string) *FiatCurrencyFeedCreate {
	if s != nil {
		fcfc.SetFeedFiatName(*s)
	}
	return fcfc
}

// SetDisabled sets the "disabled" field.
func (fcfc *FiatCurrencyFeedCreate) SetDisabled(b bool) *FiatCurrencyFeedCreate {
	fcfc.mutation.SetDisabled(b)
	return fcfc
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (fcfc *FiatCurrencyFeedCreate) SetNillableDisabled(b *bool) *FiatCurrencyFeedCreate {
	if b != nil {
		fcfc.SetDisabled(*b)
	}
	return fcfc
}

// SetID sets the "id" field.
func (fcfc *FiatCurrencyFeedCreate) SetID(u uint32) *FiatCurrencyFeedCreate {
	fcfc.mutation.SetID(u)
	return fcfc
}

// Mutation returns the FiatCurrencyFeedMutation object of the builder.
func (fcfc *FiatCurrencyFeedCreate) Mutation() *FiatCurrencyFeedMutation {
	return fcfc.mutation
}

// Save creates the FiatCurrencyFeed in the database.
func (fcfc *FiatCurrencyFeedCreate) Save(ctx context.Context) (*FiatCurrencyFeed, error) {
	fcfc.defaults()
	return withHooks(ctx, fcfc.sqlSave, fcfc.mutation, fcfc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fcfc *FiatCurrencyFeedCreate) SaveX(ctx context.Context) *FiatCurrencyFeed {
	v, err := fcfc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcfc *FiatCurrencyFeedCreate) Exec(ctx context.Context) error {
	_, err := fcfc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcfc *FiatCurrencyFeedCreate) ExecX(ctx context.Context) {
	if err := fcfc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fcfc *FiatCurrencyFeedCreate) defaults() {
	if _, ok := fcfc.mutation.CreatedAt(); !ok {
		v := fiatcurrencyfeed.DefaultCreatedAt()
		fcfc.mutation.SetCreatedAt(v)
	}
	if _, ok := fcfc.mutation.UpdatedAt(); !ok {
		v := fiatcurrencyfeed.DefaultUpdatedAt()
		fcfc.mutation.SetUpdatedAt(v)
	}
	if _, ok := fcfc.mutation.DeletedAt(); !ok {
		v := fiatcurrencyfeed.DefaultDeletedAt()
		fcfc.mutation.SetDeletedAt(v)
	}
	if _, ok := fcfc.mutation.EntID(); !ok {
		v := fiatcurrencyfeed.DefaultEntID()
		fcfc.mutation.SetEntID(v)
	}
	if _, ok := fcfc.mutation.FiatID(); !ok {
		v := fiatcurrencyfeed.DefaultFiatID()
		fcfc.mutation.SetFiatID(v)
	}
	if _, ok := fcfc.mutation.FeedType(); !ok {
		v := fiatcurrencyfeed.DefaultFeedType
		fcfc.mutation.SetFeedType(v)
	}
	if _, ok := fcfc.mutation.FeedFiatName(); !ok {
		v := fiatcurrencyfeed.DefaultFeedFiatName
		fcfc.mutation.SetFeedFiatName(v)
	}
	if _, ok := fcfc.mutation.Disabled(); !ok {
		v := fiatcurrencyfeed.DefaultDisabled
		fcfc.mutation.SetDisabled(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fcfc *FiatCurrencyFeedCreate) check() error {
	if _, ok := fcfc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "FiatCurrencyFeed.created_at"`)}
	}
	if _, ok := fcfc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "FiatCurrencyFeed.updated_at"`)}
	}
	if _, ok := fcfc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`generated: missing required field "FiatCurrencyFeed.deleted_at"`)}
	}
	if _, ok := fcfc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`generated: missing required field "FiatCurrencyFeed.ent_id"`)}
	}
	return nil
}

func (fcfc *FiatCurrencyFeedCreate) sqlSave(ctx context.Context) (*FiatCurrencyFeed, error) {
	if err := fcfc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fcfc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fcfc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	fcfc.mutation.id = &_node.ID
	fcfc.mutation.done = true
	return _node, nil
}

func (fcfc *FiatCurrencyFeedCreate) createSpec() (*FiatCurrencyFeed, *sqlgraph.CreateSpec) {
	var (
		_node = &FiatCurrencyFeed{config: fcfc.config}
		_spec = sqlgraph.NewCreateSpec(fiatcurrencyfeed.Table, sqlgraph.NewFieldSpec(fiatcurrencyfeed.FieldID, field.TypeUint32))
	)
	_spec.OnConflict = fcfc.conflict
	if id, ok := fcfc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fcfc.mutation.CreatedAt(); ok {
		_spec.SetField(fiatcurrencyfeed.FieldCreatedAt, field.TypeUint32, value)
		_node.CreatedAt = value
	}
	if value, ok := fcfc.mutation.UpdatedAt(); ok {
		_spec.SetField(fiatcurrencyfeed.FieldUpdatedAt, field.TypeUint32, value)
		_node.UpdatedAt = value
	}
	if value, ok := fcfc.mutation.DeletedAt(); ok {
		_spec.SetField(fiatcurrencyfeed.FieldDeletedAt, field.TypeUint32, value)
		_node.DeletedAt = value
	}
	if value, ok := fcfc.mutation.EntID(); ok {
		_spec.SetField(fiatcurrencyfeed.FieldEntID, field.TypeUUID, value)
		_node.EntID = value
	}
	if value, ok := fcfc.mutation.FiatID(); ok {
		_spec.SetField(fiatcurrencyfeed.FieldFiatID, field.TypeUUID, value)
		_node.FiatID = value
	}
	if value, ok := fcfc.mutation.FeedType(); ok {
		_spec.SetField(fiatcurrencyfeed.FieldFeedType, field.TypeString, value)
		_node.FeedType = value
	}
	if value, ok := fcfc.mutation.FeedFiatName(); ok {
		_spec.SetField(fiatcurrencyfeed.FieldFeedFiatName, field.TypeString, value)
		_node.FeedFiatName = value
	}
	if value, ok := fcfc.mutation.Disabled(); ok {
		_spec.SetField(fiatcurrencyfeed.FieldDisabled, field.TypeBool, value)
		_node.Disabled = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FiatCurrencyFeed.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FiatCurrencyFeedUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (fcfc *FiatCurrencyFeedCreate) OnConflict(opts ...sql.ConflictOption) *FiatCurrencyFeedUpsertOne {
	fcfc.conflict = opts
	return &FiatCurrencyFeedUpsertOne{
		create: fcfc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FiatCurrencyFeed.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fcfc *FiatCurrencyFeedCreate) OnConflictColumns(columns ...string) *FiatCurrencyFeedUpsertOne {
	fcfc.conflict = append(fcfc.conflict, sql.ConflictColumns(columns...))
	return &FiatCurrencyFeedUpsertOne{
		create: fcfc,
	}
}

type (
	// FiatCurrencyFeedUpsertOne is the builder for "upsert"-ing
	//  one FiatCurrencyFeed node.
	FiatCurrencyFeedUpsertOne struct {
		create *FiatCurrencyFeedCreate
	}

	// FiatCurrencyFeedUpsert is the "OnConflict" setter.
	FiatCurrencyFeedUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *FiatCurrencyFeedUpsert) SetCreatedAt(v uint32) *FiatCurrencyFeedUpsert {
	u.Set(fiatcurrencyfeed.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FiatCurrencyFeedUpsert) UpdateCreatedAt() *FiatCurrencyFeedUpsert {
	u.SetExcluded(fiatcurrencyfeed.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *FiatCurrencyFeedUpsert) AddCreatedAt(v uint32) *FiatCurrencyFeedUpsert {
	u.Add(fiatcurrencyfeed.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FiatCurrencyFeedUpsert) SetUpdatedAt(v uint32) *FiatCurrencyFeedUpsert {
	u.Set(fiatcurrencyfeed.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FiatCurrencyFeedUpsert) UpdateUpdatedAt() *FiatCurrencyFeedUpsert {
	u.SetExcluded(fiatcurrencyfeed.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *FiatCurrencyFeedUpsert) AddUpdatedAt(v uint32) *FiatCurrencyFeedUpsert {
	u.Add(fiatcurrencyfeed.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *FiatCurrencyFeedUpsert) SetDeletedAt(v uint32) *FiatCurrencyFeedUpsert {
	u.Set(fiatcurrencyfeed.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *FiatCurrencyFeedUpsert) UpdateDeletedAt() *FiatCurrencyFeedUpsert {
	u.SetExcluded(fiatcurrencyfeed.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *FiatCurrencyFeedUpsert) AddDeletedAt(v uint32) *FiatCurrencyFeedUpsert {
	u.Add(fiatcurrencyfeed.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *FiatCurrencyFeedUpsert) SetEntID(v uuid.UUID) *FiatCurrencyFeedUpsert {
	u.Set(fiatcurrencyfeed.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *FiatCurrencyFeedUpsert) UpdateEntID() *FiatCurrencyFeedUpsert {
	u.SetExcluded(fiatcurrencyfeed.FieldEntID)
	return u
}

// SetFiatID sets the "fiat_id" field.
func (u *FiatCurrencyFeedUpsert) SetFiatID(v uuid.UUID) *FiatCurrencyFeedUpsert {
	u.Set(fiatcurrencyfeed.FieldFiatID, v)
	return u
}

// UpdateFiatID sets the "fiat_id" field to the value that was provided on create.
func (u *FiatCurrencyFeedUpsert) UpdateFiatID() *FiatCurrencyFeedUpsert {
	u.SetExcluded(fiatcurrencyfeed.FieldFiatID)
	return u
}

// ClearFiatID clears the value of the "fiat_id" field.
func (u *FiatCurrencyFeedUpsert) ClearFiatID() *FiatCurrencyFeedUpsert {
	u.SetNull(fiatcurrencyfeed.FieldFiatID)
	return u
}

// SetFeedType sets the "feed_type" field.
func (u *FiatCurrencyFeedUpsert) SetFeedType(v string) *FiatCurrencyFeedUpsert {
	u.Set(fiatcurrencyfeed.FieldFeedType, v)
	return u
}

// UpdateFeedType sets the "feed_type" field to the value that was provided on create.
func (u *FiatCurrencyFeedUpsert) UpdateFeedType() *FiatCurrencyFeedUpsert {
	u.SetExcluded(fiatcurrencyfeed.FieldFeedType)
	return u
}

// ClearFeedType clears the value of the "feed_type" field.
func (u *FiatCurrencyFeedUpsert) ClearFeedType() *FiatCurrencyFeedUpsert {
	u.SetNull(fiatcurrencyfeed.FieldFeedType)
	return u
}

// SetFeedFiatName sets the "feed_fiat_name" field.
func (u *FiatCurrencyFeedUpsert) SetFeedFiatName(v string) *FiatCurrencyFeedUpsert {
	u.Set(fiatcurrencyfeed.FieldFeedFiatName, v)
	return u
}

// UpdateFeedFiatName sets the "feed_fiat_name" field to the value that was provided on create.
func (u *FiatCurrencyFeedUpsert) UpdateFeedFiatName() *FiatCurrencyFeedUpsert {
	u.SetExcluded(fiatcurrencyfeed.FieldFeedFiatName)
	return u
}

// ClearFeedFiatName clears the value of the "feed_fiat_name" field.
func (u *FiatCurrencyFeedUpsert) ClearFeedFiatName() *FiatCurrencyFeedUpsert {
	u.SetNull(fiatcurrencyfeed.FieldFeedFiatName)
	return u
}

// SetDisabled sets the "disabled" field.
func (u *FiatCurrencyFeedUpsert) SetDisabled(v bool) *FiatCurrencyFeedUpsert {
	u.Set(fiatcurrencyfeed.FieldDisabled, v)
	return u
}

// UpdateDisabled sets the "disabled" field to the value that was provided on create.
func (u *FiatCurrencyFeedUpsert) UpdateDisabled() *FiatCurrencyFeedUpsert {
	u.SetExcluded(fiatcurrencyfeed.FieldDisabled)
	return u
}

// ClearDisabled clears the value of the "disabled" field.
func (u *FiatCurrencyFeedUpsert) ClearDisabled() *FiatCurrencyFeedUpsert {
	u.SetNull(fiatcurrencyfeed.FieldDisabled)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.FiatCurrencyFeed.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(fiatcurrencyfeed.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FiatCurrencyFeedUpsertOne) UpdateNewValues() *FiatCurrencyFeedUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(fiatcurrencyfeed.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.FiatCurrencyFeed.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *FiatCurrencyFeedUpsertOne) Ignore() *FiatCurrencyFeedUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FiatCurrencyFeedUpsertOne) DoNothing() *FiatCurrencyFeedUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FiatCurrencyFeedCreate.OnConflict
// documentation for more info.
func (u *FiatCurrencyFeedUpsertOne) Update(set func(*FiatCurrencyFeedUpsert)) *FiatCurrencyFeedUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FiatCurrencyFeedUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *FiatCurrencyFeedUpsertOne) SetCreatedAt(v uint32) *FiatCurrencyFeedUpsertOne {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *FiatCurrencyFeedUpsertOne) AddCreatedAt(v uint32) *FiatCurrencyFeedUpsertOne {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FiatCurrencyFeedUpsertOne) UpdateCreatedAt() *FiatCurrencyFeedUpsertOne {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FiatCurrencyFeedUpsertOne) SetUpdatedAt(v uint32) *FiatCurrencyFeedUpsertOne {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *FiatCurrencyFeedUpsertOne) AddUpdatedAt(v uint32) *FiatCurrencyFeedUpsertOne {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FiatCurrencyFeedUpsertOne) UpdateUpdatedAt() *FiatCurrencyFeedUpsertOne {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *FiatCurrencyFeedUpsertOne) SetDeletedAt(v uint32) *FiatCurrencyFeedUpsertOne {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *FiatCurrencyFeedUpsertOne) AddDeletedAt(v uint32) *FiatCurrencyFeedUpsertOne {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *FiatCurrencyFeedUpsertOne) UpdateDeletedAt() *FiatCurrencyFeedUpsertOne {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *FiatCurrencyFeedUpsertOne) SetEntID(v uuid.UUID) *FiatCurrencyFeedUpsertOne {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *FiatCurrencyFeedUpsertOne) UpdateEntID() *FiatCurrencyFeedUpsertOne {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.UpdateEntID()
	})
}

// SetFiatID sets the "fiat_id" field.
func (u *FiatCurrencyFeedUpsertOne) SetFiatID(v uuid.UUID) *FiatCurrencyFeedUpsertOne {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.SetFiatID(v)
	})
}

// UpdateFiatID sets the "fiat_id" field to the value that was provided on create.
func (u *FiatCurrencyFeedUpsertOne) UpdateFiatID() *FiatCurrencyFeedUpsertOne {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.UpdateFiatID()
	})
}

// ClearFiatID clears the value of the "fiat_id" field.
func (u *FiatCurrencyFeedUpsertOne) ClearFiatID() *FiatCurrencyFeedUpsertOne {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.ClearFiatID()
	})
}

// SetFeedType sets the "feed_type" field.
func (u *FiatCurrencyFeedUpsertOne) SetFeedType(v string) *FiatCurrencyFeedUpsertOne {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.SetFeedType(v)
	})
}

// UpdateFeedType sets the "feed_type" field to the value that was provided on create.
func (u *FiatCurrencyFeedUpsertOne) UpdateFeedType() *FiatCurrencyFeedUpsertOne {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.UpdateFeedType()
	})
}

// ClearFeedType clears the value of the "feed_type" field.
func (u *FiatCurrencyFeedUpsertOne) ClearFeedType() *FiatCurrencyFeedUpsertOne {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.ClearFeedType()
	})
}

// SetFeedFiatName sets the "feed_fiat_name" field.
func (u *FiatCurrencyFeedUpsertOne) SetFeedFiatName(v string) *FiatCurrencyFeedUpsertOne {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.SetFeedFiatName(v)
	})
}

// UpdateFeedFiatName sets the "feed_fiat_name" field to the value that was provided on create.
func (u *FiatCurrencyFeedUpsertOne) UpdateFeedFiatName() *FiatCurrencyFeedUpsertOne {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.UpdateFeedFiatName()
	})
}

// ClearFeedFiatName clears the value of the "feed_fiat_name" field.
func (u *FiatCurrencyFeedUpsertOne) ClearFeedFiatName() *FiatCurrencyFeedUpsertOne {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.ClearFeedFiatName()
	})
}

// SetDisabled sets the "disabled" field.
func (u *FiatCurrencyFeedUpsertOne) SetDisabled(v bool) *FiatCurrencyFeedUpsertOne {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.SetDisabled(v)
	})
}

// UpdateDisabled sets the "disabled" field to the value that was provided on create.
func (u *FiatCurrencyFeedUpsertOne) UpdateDisabled() *FiatCurrencyFeedUpsertOne {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.UpdateDisabled()
	})
}

// ClearDisabled clears the value of the "disabled" field.
func (u *FiatCurrencyFeedUpsertOne) ClearDisabled() *FiatCurrencyFeedUpsertOne {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.ClearDisabled()
	})
}

// Exec executes the query.
func (u *FiatCurrencyFeedUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for FiatCurrencyFeedCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FiatCurrencyFeedUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FiatCurrencyFeedUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FiatCurrencyFeedUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FiatCurrencyFeedCreateBulk is the builder for creating many FiatCurrencyFeed entities in bulk.
type FiatCurrencyFeedCreateBulk struct {
	config
	err      error
	builders []*FiatCurrencyFeedCreate
	conflict []sql.ConflictOption
}

// Save creates the FiatCurrencyFeed entities in the database.
func (fcfcb *FiatCurrencyFeedCreateBulk) Save(ctx context.Context) ([]*FiatCurrencyFeed, error) {
	if fcfcb.err != nil {
		return nil, fcfcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(fcfcb.builders))
	nodes := make([]*FiatCurrencyFeed, len(fcfcb.builders))
	mutators := make([]Mutator, len(fcfcb.builders))
	for i := range fcfcb.builders {
		func(i int, root context.Context) {
			builder := fcfcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FiatCurrencyFeedMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fcfcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = fcfcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcfcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint32(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, fcfcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcfcb *FiatCurrencyFeedCreateBulk) SaveX(ctx context.Context) []*FiatCurrencyFeed {
	v, err := fcfcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcfcb *FiatCurrencyFeedCreateBulk) Exec(ctx context.Context) error {
	_, err := fcfcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcfcb *FiatCurrencyFeedCreateBulk) ExecX(ctx context.Context) {
	if err := fcfcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FiatCurrencyFeed.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FiatCurrencyFeedUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (fcfcb *FiatCurrencyFeedCreateBulk) OnConflict(opts ...sql.ConflictOption) *FiatCurrencyFeedUpsertBulk {
	fcfcb.conflict = opts
	return &FiatCurrencyFeedUpsertBulk{
		create: fcfcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FiatCurrencyFeed.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fcfcb *FiatCurrencyFeedCreateBulk) OnConflictColumns(columns ...string) *FiatCurrencyFeedUpsertBulk {
	fcfcb.conflict = append(fcfcb.conflict, sql.ConflictColumns(columns...))
	return &FiatCurrencyFeedUpsertBulk{
		create: fcfcb,
	}
}

// FiatCurrencyFeedUpsertBulk is the builder for "upsert"-ing
// a bulk of FiatCurrencyFeed nodes.
type FiatCurrencyFeedUpsertBulk struct {
	create *FiatCurrencyFeedCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.FiatCurrencyFeed.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(fiatcurrencyfeed.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FiatCurrencyFeedUpsertBulk) UpdateNewValues() *FiatCurrencyFeedUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(fiatcurrencyfeed.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.FiatCurrencyFeed.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *FiatCurrencyFeedUpsertBulk) Ignore() *FiatCurrencyFeedUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FiatCurrencyFeedUpsertBulk) DoNothing() *FiatCurrencyFeedUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FiatCurrencyFeedCreateBulk.OnConflict
// documentation for more info.
func (u *FiatCurrencyFeedUpsertBulk) Update(set func(*FiatCurrencyFeedUpsert)) *FiatCurrencyFeedUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FiatCurrencyFeedUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *FiatCurrencyFeedUpsertBulk) SetCreatedAt(v uint32) *FiatCurrencyFeedUpsertBulk {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *FiatCurrencyFeedUpsertBulk) AddCreatedAt(v uint32) *FiatCurrencyFeedUpsertBulk {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FiatCurrencyFeedUpsertBulk) UpdateCreatedAt() *FiatCurrencyFeedUpsertBulk {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FiatCurrencyFeedUpsertBulk) SetUpdatedAt(v uint32) *FiatCurrencyFeedUpsertBulk {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *FiatCurrencyFeedUpsertBulk) AddUpdatedAt(v uint32) *FiatCurrencyFeedUpsertBulk {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FiatCurrencyFeedUpsertBulk) UpdateUpdatedAt() *FiatCurrencyFeedUpsertBulk {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *FiatCurrencyFeedUpsertBulk) SetDeletedAt(v uint32) *FiatCurrencyFeedUpsertBulk {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *FiatCurrencyFeedUpsertBulk) AddDeletedAt(v uint32) *FiatCurrencyFeedUpsertBulk {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *FiatCurrencyFeedUpsertBulk) UpdateDeletedAt() *FiatCurrencyFeedUpsertBulk {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *FiatCurrencyFeedUpsertBulk) SetEntID(v uuid.UUID) *FiatCurrencyFeedUpsertBulk {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *FiatCurrencyFeedUpsertBulk) UpdateEntID() *FiatCurrencyFeedUpsertBulk {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.UpdateEntID()
	})
}

// SetFiatID sets the "fiat_id" field.
func (u *FiatCurrencyFeedUpsertBulk) SetFiatID(v uuid.UUID) *FiatCurrencyFeedUpsertBulk {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.SetFiatID(v)
	})
}

// UpdateFiatID sets the "fiat_id" field to the value that was provided on create.
func (u *FiatCurrencyFeedUpsertBulk) UpdateFiatID() *FiatCurrencyFeedUpsertBulk {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.UpdateFiatID()
	})
}

// ClearFiatID clears the value of the "fiat_id" field.
func (u *FiatCurrencyFeedUpsertBulk) ClearFiatID() *FiatCurrencyFeedUpsertBulk {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.ClearFiatID()
	})
}

// SetFeedType sets the "feed_type" field.
func (u *FiatCurrencyFeedUpsertBulk) SetFeedType(v string) *FiatCurrencyFeedUpsertBulk {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.SetFeedType(v)
	})
}

// UpdateFeedType sets the "feed_type" field to the value that was provided on create.
func (u *FiatCurrencyFeedUpsertBulk) UpdateFeedType() *FiatCurrencyFeedUpsertBulk {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.UpdateFeedType()
	})
}

// ClearFeedType clears the value of the "feed_type" field.
func (u *FiatCurrencyFeedUpsertBulk) ClearFeedType() *FiatCurrencyFeedUpsertBulk {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.ClearFeedType()
	})
}

// SetFeedFiatName sets the "feed_fiat_name" field.
func (u *FiatCurrencyFeedUpsertBulk) SetFeedFiatName(v string) *FiatCurrencyFeedUpsertBulk {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.SetFeedFiatName(v)
	})
}

// UpdateFeedFiatName sets the "feed_fiat_name" field to the value that was provided on create.
func (u *FiatCurrencyFeedUpsertBulk) UpdateFeedFiatName() *FiatCurrencyFeedUpsertBulk {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.UpdateFeedFiatName()
	})
}

// ClearFeedFiatName clears the value of the "feed_fiat_name" field.
func (u *FiatCurrencyFeedUpsertBulk) ClearFeedFiatName() *FiatCurrencyFeedUpsertBulk {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.ClearFeedFiatName()
	})
}

// SetDisabled sets the "disabled" field.
func (u *FiatCurrencyFeedUpsertBulk) SetDisabled(v bool) *FiatCurrencyFeedUpsertBulk {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.SetDisabled(v)
	})
}

// UpdateDisabled sets the "disabled" field to the value that was provided on create.
func (u *FiatCurrencyFeedUpsertBulk) UpdateDisabled() *FiatCurrencyFeedUpsertBulk {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.UpdateDisabled()
	})
}

// ClearDisabled clears the value of the "disabled" field.
func (u *FiatCurrencyFeedUpsertBulk) ClearDisabled() *FiatCurrencyFeedUpsertBulk {
	return u.Update(func(s *FiatCurrencyFeedUpsert) {
		s.ClearDisabled()
	})
}

// Exec executes the query.
func (u *FiatCurrencyFeedUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("generated: OnConflict was set for builder %d. Set it on the FiatCurrencyFeedCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for FiatCurrencyFeedCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FiatCurrencyFeedUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
