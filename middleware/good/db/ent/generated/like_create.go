// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/like"
	"github.com/google/uuid"
)

// LikeCreate is the builder for creating a Like entity.
type LikeCreate struct {
	config
	mutation *LikeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetEntID sets the "ent_id" field.
func (lc *LikeCreate) SetEntID(u uuid.UUID) *LikeCreate {
	lc.mutation.SetEntID(u)
	return lc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (lc *LikeCreate) SetNillableEntID(u *uuid.UUID) *LikeCreate {
	if u != nil {
		lc.SetEntID(*u)
	}
	return lc
}

// SetCreatedAt sets the "created_at" field.
func (lc *LikeCreate) SetCreatedAt(u uint32) *LikeCreate {
	lc.mutation.SetCreatedAt(u)
	return lc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lc *LikeCreate) SetNillableCreatedAt(u *uint32) *LikeCreate {
	if u != nil {
		lc.SetCreatedAt(*u)
	}
	return lc
}

// SetUpdatedAt sets the "updated_at" field.
func (lc *LikeCreate) SetUpdatedAt(u uint32) *LikeCreate {
	lc.mutation.SetUpdatedAt(u)
	return lc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lc *LikeCreate) SetNillableUpdatedAt(u *uint32) *LikeCreate {
	if u != nil {
		lc.SetUpdatedAt(*u)
	}
	return lc
}

// SetDeletedAt sets the "deleted_at" field.
func (lc *LikeCreate) SetDeletedAt(u uint32) *LikeCreate {
	lc.mutation.SetDeletedAt(u)
	return lc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (lc *LikeCreate) SetNillableDeletedAt(u *uint32) *LikeCreate {
	if u != nil {
		lc.SetDeletedAt(*u)
	}
	return lc
}

// SetUserID sets the "user_id" field.
func (lc *LikeCreate) SetUserID(u uuid.UUID) *LikeCreate {
	lc.mutation.SetUserID(u)
	return lc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (lc *LikeCreate) SetNillableUserID(u *uuid.UUID) *LikeCreate {
	if u != nil {
		lc.SetUserID(*u)
	}
	return lc
}

// SetAppGoodID sets the "app_good_id" field.
func (lc *LikeCreate) SetAppGoodID(u uuid.UUID) *LikeCreate {
	lc.mutation.SetAppGoodID(u)
	return lc
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (lc *LikeCreate) SetNillableAppGoodID(u *uuid.UUID) *LikeCreate {
	if u != nil {
		lc.SetAppGoodID(*u)
	}
	return lc
}

// SetLike sets the "like" field.
func (lc *LikeCreate) SetLike(b bool) *LikeCreate {
	lc.mutation.SetLike(b)
	return lc
}

// SetID sets the "id" field.
func (lc *LikeCreate) SetID(u uint32) *LikeCreate {
	lc.mutation.SetID(u)
	return lc
}

// Mutation returns the LikeMutation object of the builder.
func (lc *LikeCreate) Mutation() *LikeMutation {
	return lc.mutation
}

// Save creates the Like in the database.
func (lc *LikeCreate) Save(ctx context.Context) (*Like, error) {
	lc.defaults()
	return withHooks(ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LikeCreate) SaveX(ctx context.Context) *Like {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LikeCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LikeCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LikeCreate) defaults() {
	if _, ok := lc.mutation.EntID(); !ok {
		v := like.DefaultEntID()
		lc.mutation.SetEntID(v)
	}
	if _, ok := lc.mutation.CreatedAt(); !ok {
		v := like.DefaultCreatedAt()
		lc.mutation.SetCreatedAt(v)
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		v := like.DefaultUpdatedAt()
		lc.mutation.SetUpdatedAt(v)
	}
	if _, ok := lc.mutation.DeletedAt(); !ok {
		v := like.DefaultDeletedAt()
		lc.mutation.SetDeletedAt(v)
	}
	if _, ok := lc.mutation.UserID(); !ok {
		v := like.DefaultUserID()
		lc.mutation.SetUserID(v)
	}
	if _, ok := lc.mutation.AppGoodID(); !ok {
		v := like.DefaultAppGoodID()
		lc.mutation.SetAppGoodID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LikeCreate) check() error {
	if _, ok := lc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`generated: missing required field "Like.ent_id"`)}
	}
	if _, ok := lc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "Like.created_at"`)}
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "Like.updated_at"`)}
	}
	if _, ok := lc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`generated: missing required field "Like.deleted_at"`)}
	}
	if _, ok := lc.mutation.Like(); !ok {
		return &ValidationError{Name: "like", err: errors.New(`generated: missing required field "Like.like"`)}
	}
	return nil
}

func (lc *LikeCreate) sqlSave(ctx context.Context) (*Like, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LikeCreate) createSpec() (*Like, *sqlgraph.CreateSpec) {
	var (
		_node = &Like{config: lc.config}
		_spec = sqlgraph.NewCreateSpec(like.Table, sqlgraph.NewFieldSpec(like.FieldID, field.TypeUint32))
	)
	_spec.OnConflict = lc.conflict
	if id, ok := lc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := lc.mutation.EntID(); ok {
		_spec.SetField(like.FieldEntID, field.TypeUUID, value)
		_node.EntID = value
	}
	if value, ok := lc.mutation.CreatedAt(); ok {
		_spec.SetField(like.FieldCreatedAt, field.TypeUint32, value)
		_node.CreatedAt = value
	}
	if value, ok := lc.mutation.UpdatedAt(); ok {
		_spec.SetField(like.FieldUpdatedAt, field.TypeUint32, value)
		_node.UpdatedAt = value
	}
	if value, ok := lc.mutation.DeletedAt(); ok {
		_spec.SetField(like.FieldDeletedAt, field.TypeUint32, value)
		_node.DeletedAt = value
	}
	if value, ok := lc.mutation.UserID(); ok {
		_spec.SetField(like.FieldUserID, field.TypeUUID, value)
		_node.UserID = value
	}
	if value, ok := lc.mutation.AppGoodID(); ok {
		_spec.SetField(like.FieldAppGoodID, field.TypeUUID, value)
		_node.AppGoodID = value
	}
	if value, ok := lc.mutation.Like(); ok {
		_spec.SetField(like.FieldLike, field.TypeBool, value)
		_node.Like = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Like.Create().
//		SetEntID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LikeUpsert) {
//			SetEntID(v+v).
//		}).
//		Exec(ctx)
func (lc *LikeCreate) OnConflict(opts ...sql.ConflictOption) *LikeUpsertOne {
	lc.conflict = opts
	return &LikeUpsertOne{
		create: lc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Like.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (lc *LikeCreate) OnConflictColumns(columns ...string) *LikeUpsertOne {
	lc.conflict = append(lc.conflict, sql.ConflictColumns(columns...))
	return &LikeUpsertOne{
		create: lc,
	}
}

type (
	// LikeUpsertOne is the builder for "upsert"-ing
	//  one Like node.
	LikeUpsertOne struct {
		create *LikeCreate
	}

	// LikeUpsert is the "OnConflict" setter.
	LikeUpsert struct {
		*sql.UpdateSet
	}
)

// SetEntID sets the "ent_id" field.
func (u *LikeUpsert) SetEntID(v uuid.UUID) *LikeUpsert {
	u.Set(like.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *LikeUpsert) UpdateEntID() *LikeUpsert {
	u.SetExcluded(like.FieldEntID)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *LikeUpsert) SetCreatedAt(v uint32) *LikeUpsert {
	u.Set(like.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *LikeUpsert) UpdateCreatedAt() *LikeUpsert {
	u.SetExcluded(like.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *LikeUpsert) AddCreatedAt(v uint32) *LikeUpsert {
	u.Add(like.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *LikeUpsert) SetUpdatedAt(v uint32) *LikeUpsert {
	u.Set(like.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *LikeUpsert) UpdateUpdatedAt() *LikeUpsert {
	u.SetExcluded(like.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *LikeUpsert) AddUpdatedAt(v uint32) *LikeUpsert {
	u.Add(like.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *LikeUpsert) SetDeletedAt(v uint32) *LikeUpsert {
	u.Set(like.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *LikeUpsert) UpdateDeletedAt() *LikeUpsert {
	u.SetExcluded(like.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *LikeUpsert) AddDeletedAt(v uint32) *LikeUpsert {
	u.Add(like.FieldDeletedAt, v)
	return u
}

// SetUserID sets the "user_id" field.
func (u *LikeUpsert) SetUserID(v uuid.UUID) *LikeUpsert {
	u.Set(like.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *LikeUpsert) UpdateUserID() *LikeUpsert {
	u.SetExcluded(like.FieldUserID)
	return u
}

// ClearUserID clears the value of the "user_id" field.
func (u *LikeUpsert) ClearUserID() *LikeUpsert {
	u.SetNull(like.FieldUserID)
	return u
}

// SetAppGoodID sets the "app_good_id" field.
func (u *LikeUpsert) SetAppGoodID(v uuid.UUID) *LikeUpsert {
	u.Set(like.FieldAppGoodID, v)
	return u
}

// UpdateAppGoodID sets the "app_good_id" field to the value that was provided on create.
func (u *LikeUpsert) UpdateAppGoodID() *LikeUpsert {
	u.SetExcluded(like.FieldAppGoodID)
	return u
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (u *LikeUpsert) ClearAppGoodID() *LikeUpsert {
	u.SetNull(like.FieldAppGoodID)
	return u
}

// SetLike sets the "like" field.
func (u *LikeUpsert) SetLike(v bool) *LikeUpsert {
	u.Set(like.FieldLike, v)
	return u
}

// UpdateLike sets the "like" field to the value that was provided on create.
func (u *LikeUpsert) UpdateLike() *LikeUpsert {
	u.SetExcluded(like.FieldLike)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Like.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(like.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *LikeUpsertOne) UpdateNewValues() *LikeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(like.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Like.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *LikeUpsertOne) Ignore() *LikeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LikeUpsertOne) DoNothing() *LikeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LikeCreate.OnConflict
// documentation for more info.
func (u *LikeUpsertOne) Update(set func(*LikeUpsert)) *LikeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LikeUpsert{UpdateSet: update})
	}))
	return u
}

// SetEntID sets the "ent_id" field.
func (u *LikeUpsertOne) SetEntID(v uuid.UUID) *LikeUpsertOne {
	return u.Update(func(s *LikeUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *LikeUpsertOne) UpdateEntID() *LikeUpsertOne {
	return u.Update(func(s *LikeUpsert) {
		s.UpdateEntID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *LikeUpsertOne) SetCreatedAt(v uint32) *LikeUpsertOne {
	return u.Update(func(s *LikeUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *LikeUpsertOne) AddCreatedAt(v uint32) *LikeUpsertOne {
	return u.Update(func(s *LikeUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *LikeUpsertOne) UpdateCreatedAt() *LikeUpsertOne {
	return u.Update(func(s *LikeUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *LikeUpsertOne) SetUpdatedAt(v uint32) *LikeUpsertOne {
	return u.Update(func(s *LikeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *LikeUpsertOne) AddUpdatedAt(v uint32) *LikeUpsertOne {
	return u.Update(func(s *LikeUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *LikeUpsertOne) UpdateUpdatedAt() *LikeUpsertOne {
	return u.Update(func(s *LikeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *LikeUpsertOne) SetDeletedAt(v uint32) *LikeUpsertOne {
	return u.Update(func(s *LikeUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *LikeUpsertOne) AddDeletedAt(v uint32) *LikeUpsertOne {
	return u.Update(func(s *LikeUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *LikeUpsertOne) UpdateDeletedAt() *LikeUpsertOne {
	return u.Update(func(s *LikeUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetUserID sets the "user_id" field.
func (u *LikeUpsertOne) SetUserID(v uuid.UUID) *LikeUpsertOne {
	return u.Update(func(s *LikeUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *LikeUpsertOne) UpdateUserID() *LikeUpsertOne {
	return u.Update(func(s *LikeUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *LikeUpsertOne) ClearUserID() *LikeUpsertOne {
	return u.Update(func(s *LikeUpsert) {
		s.ClearUserID()
	})
}

// SetAppGoodID sets the "app_good_id" field.
func (u *LikeUpsertOne) SetAppGoodID(v uuid.UUID) *LikeUpsertOne {
	return u.Update(func(s *LikeUpsert) {
		s.SetAppGoodID(v)
	})
}

// UpdateAppGoodID sets the "app_good_id" field to the value that was provided on create.
func (u *LikeUpsertOne) UpdateAppGoodID() *LikeUpsertOne {
	return u.Update(func(s *LikeUpsert) {
		s.UpdateAppGoodID()
	})
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (u *LikeUpsertOne) ClearAppGoodID() *LikeUpsertOne {
	return u.Update(func(s *LikeUpsert) {
		s.ClearAppGoodID()
	})
}

// SetLike sets the "like" field.
func (u *LikeUpsertOne) SetLike(v bool) *LikeUpsertOne {
	return u.Update(func(s *LikeUpsert) {
		s.SetLike(v)
	})
}

// UpdateLike sets the "like" field to the value that was provided on create.
func (u *LikeUpsertOne) UpdateLike() *LikeUpsertOne {
	return u.Update(func(s *LikeUpsert) {
		s.UpdateLike()
	})
}

// Exec executes the query.
func (u *LikeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for LikeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LikeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *LikeUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *LikeUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// LikeCreateBulk is the builder for creating many Like entities in bulk.
type LikeCreateBulk struct {
	config
	err      error
	builders []*LikeCreate
	conflict []sql.ConflictOption
}

// Save creates the Like entities in the database.
func (lcb *LikeCreateBulk) Save(ctx context.Context) ([]*Like, error) {
	if lcb.err != nil {
		return nil, lcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Like, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LikeMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = lcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LikeCreateBulk) SaveX(ctx context.Context) []*Like {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LikeCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LikeCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Like.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LikeUpsert) {
//			SetEntID(v+v).
//		}).
//		Exec(ctx)
func (lcb *LikeCreateBulk) OnConflict(opts ...sql.ConflictOption) *LikeUpsertBulk {
	lcb.conflict = opts
	return &LikeUpsertBulk{
		create: lcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Like.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (lcb *LikeCreateBulk) OnConflictColumns(columns ...string) *LikeUpsertBulk {
	lcb.conflict = append(lcb.conflict, sql.ConflictColumns(columns...))
	return &LikeUpsertBulk{
		create: lcb,
	}
}

// LikeUpsertBulk is the builder for "upsert"-ing
// a bulk of Like nodes.
type LikeUpsertBulk struct {
	create *LikeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Like.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(like.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *LikeUpsertBulk) UpdateNewValues() *LikeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(like.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Like.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *LikeUpsertBulk) Ignore() *LikeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LikeUpsertBulk) DoNothing() *LikeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LikeCreateBulk.OnConflict
// documentation for more info.
func (u *LikeUpsertBulk) Update(set func(*LikeUpsert)) *LikeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LikeUpsert{UpdateSet: update})
	}))
	return u
}

// SetEntID sets the "ent_id" field.
func (u *LikeUpsertBulk) SetEntID(v uuid.UUID) *LikeUpsertBulk {
	return u.Update(func(s *LikeUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *LikeUpsertBulk) UpdateEntID() *LikeUpsertBulk {
	return u.Update(func(s *LikeUpsert) {
		s.UpdateEntID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *LikeUpsertBulk) SetCreatedAt(v uint32) *LikeUpsertBulk {
	return u.Update(func(s *LikeUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *LikeUpsertBulk) AddCreatedAt(v uint32) *LikeUpsertBulk {
	return u.Update(func(s *LikeUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *LikeUpsertBulk) UpdateCreatedAt() *LikeUpsertBulk {
	return u.Update(func(s *LikeUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *LikeUpsertBulk) SetUpdatedAt(v uint32) *LikeUpsertBulk {
	return u.Update(func(s *LikeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *LikeUpsertBulk) AddUpdatedAt(v uint32) *LikeUpsertBulk {
	return u.Update(func(s *LikeUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *LikeUpsertBulk) UpdateUpdatedAt() *LikeUpsertBulk {
	return u.Update(func(s *LikeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *LikeUpsertBulk) SetDeletedAt(v uint32) *LikeUpsertBulk {
	return u.Update(func(s *LikeUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *LikeUpsertBulk) AddDeletedAt(v uint32) *LikeUpsertBulk {
	return u.Update(func(s *LikeUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *LikeUpsertBulk) UpdateDeletedAt() *LikeUpsertBulk {
	return u.Update(func(s *LikeUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetUserID sets the "user_id" field.
func (u *LikeUpsertBulk) SetUserID(v uuid.UUID) *LikeUpsertBulk {
	return u.Update(func(s *LikeUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *LikeUpsertBulk) UpdateUserID() *LikeUpsertBulk {
	return u.Update(func(s *LikeUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *LikeUpsertBulk) ClearUserID() *LikeUpsertBulk {
	return u.Update(func(s *LikeUpsert) {
		s.ClearUserID()
	})
}

// SetAppGoodID sets the "app_good_id" field.
func (u *LikeUpsertBulk) SetAppGoodID(v uuid.UUID) *LikeUpsertBulk {
	return u.Update(func(s *LikeUpsert) {
		s.SetAppGoodID(v)
	})
}

// UpdateAppGoodID sets the "app_good_id" field to the value that was provided on create.
func (u *LikeUpsertBulk) UpdateAppGoodID() *LikeUpsertBulk {
	return u.Update(func(s *LikeUpsert) {
		s.UpdateAppGoodID()
	})
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (u *LikeUpsertBulk) ClearAppGoodID() *LikeUpsertBulk {
	return u.Update(func(s *LikeUpsert) {
		s.ClearAppGoodID()
	})
}

// SetLike sets the "like" field.
func (u *LikeUpsertBulk) SetLike(v bool) *LikeUpsertBulk {
	return u.Update(func(s *LikeUpsert) {
		s.SetLike(v)
	})
}

// UpdateLike sets the "like" field to the value that was provided on create.
func (u *LikeUpsertBulk) UpdateLike() *LikeUpsertBulk {
	return u.Update(func(s *LikeUpsert) {
		s.UpdateLike()
	})
}

// Exec executes the query.
func (u *LikeUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("generated: OnConflict was set for builder %d. Set it on the LikeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for LikeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LikeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
