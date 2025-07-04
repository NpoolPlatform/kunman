// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated/applang"
	"github.com/google/uuid"
)

// AppLangCreate is the builder for creating a AppLang entity.
type AppLangCreate struct {
	config
	mutation *AppLangMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (alc *AppLangCreate) SetCreatedAt(u uint32) *AppLangCreate {
	alc.mutation.SetCreatedAt(u)
	return alc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (alc *AppLangCreate) SetNillableCreatedAt(u *uint32) *AppLangCreate {
	if u != nil {
		alc.SetCreatedAt(*u)
	}
	return alc
}

// SetUpdatedAt sets the "updated_at" field.
func (alc *AppLangCreate) SetUpdatedAt(u uint32) *AppLangCreate {
	alc.mutation.SetUpdatedAt(u)
	return alc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (alc *AppLangCreate) SetNillableUpdatedAt(u *uint32) *AppLangCreate {
	if u != nil {
		alc.SetUpdatedAt(*u)
	}
	return alc
}

// SetDeletedAt sets the "deleted_at" field.
func (alc *AppLangCreate) SetDeletedAt(u uint32) *AppLangCreate {
	alc.mutation.SetDeletedAt(u)
	return alc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (alc *AppLangCreate) SetNillableDeletedAt(u *uint32) *AppLangCreate {
	if u != nil {
		alc.SetDeletedAt(*u)
	}
	return alc
}

// SetEntID sets the "ent_id" field.
func (alc *AppLangCreate) SetEntID(u uuid.UUID) *AppLangCreate {
	alc.mutation.SetEntID(u)
	return alc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (alc *AppLangCreate) SetNillableEntID(u *uuid.UUID) *AppLangCreate {
	if u != nil {
		alc.SetEntID(*u)
	}
	return alc
}

// SetAppID sets the "app_id" field.
func (alc *AppLangCreate) SetAppID(u uuid.UUID) *AppLangCreate {
	alc.mutation.SetAppID(u)
	return alc
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (alc *AppLangCreate) SetNillableAppID(u *uuid.UUID) *AppLangCreate {
	if u != nil {
		alc.SetAppID(*u)
	}
	return alc
}

// SetLangID sets the "lang_id" field.
func (alc *AppLangCreate) SetLangID(u uuid.UUID) *AppLangCreate {
	alc.mutation.SetLangID(u)
	return alc
}

// SetNillableLangID sets the "lang_id" field if the given value is not nil.
func (alc *AppLangCreate) SetNillableLangID(u *uuid.UUID) *AppLangCreate {
	if u != nil {
		alc.SetLangID(*u)
	}
	return alc
}

// SetMain sets the "main" field.
func (alc *AppLangCreate) SetMain(b bool) *AppLangCreate {
	alc.mutation.SetMain(b)
	return alc
}

// SetNillableMain sets the "main" field if the given value is not nil.
func (alc *AppLangCreate) SetNillableMain(b *bool) *AppLangCreate {
	if b != nil {
		alc.SetMain(*b)
	}
	return alc
}

// SetID sets the "id" field.
func (alc *AppLangCreate) SetID(u uint32) *AppLangCreate {
	alc.mutation.SetID(u)
	return alc
}

// Mutation returns the AppLangMutation object of the builder.
func (alc *AppLangCreate) Mutation() *AppLangMutation {
	return alc.mutation
}

// Save creates the AppLang in the database.
func (alc *AppLangCreate) Save(ctx context.Context) (*AppLang, error) {
	alc.defaults()
	return withHooks(ctx, alc.sqlSave, alc.mutation, alc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (alc *AppLangCreate) SaveX(ctx context.Context) *AppLang {
	v, err := alc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (alc *AppLangCreate) Exec(ctx context.Context) error {
	_, err := alc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alc *AppLangCreate) ExecX(ctx context.Context) {
	if err := alc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (alc *AppLangCreate) defaults() {
	if _, ok := alc.mutation.CreatedAt(); !ok {
		v := applang.DefaultCreatedAt()
		alc.mutation.SetCreatedAt(v)
	}
	if _, ok := alc.mutation.UpdatedAt(); !ok {
		v := applang.DefaultUpdatedAt()
		alc.mutation.SetUpdatedAt(v)
	}
	if _, ok := alc.mutation.DeletedAt(); !ok {
		v := applang.DefaultDeletedAt()
		alc.mutation.SetDeletedAt(v)
	}
	if _, ok := alc.mutation.EntID(); !ok {
		v := applang.DefaultEntID()
		alc.mutation.SetEntID(v)
	}
	if _, ok := alc.mutation.AppID(); !ok {
		v := applang.DefaultAppID()
		alc.mutation.SetAppID(v)
	}
	if _, ok := alc.mutation.LangID(); !ok {
		v := applang.DefaultLangID()
		alc.mutation.SetLangID(v)
	}
	if _, ok := alc.mutation.Main(); !ok {
		v := applang.DefaultMain
		alc.mutation.SetMain(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (alc *AppLangCreate) check() error {
	if _, ok := alc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "AppLang.created_at"`)}
	}
	if _, ok := alc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "AppLang.updated_at"`)}
	}
	if _, ok := alc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`generated: missing required field "AppLang.deleted_at"`)}
	}
	if _, ok := alc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`generated: missing required field "AppLang.ent_id"`)}
	}
	return nil
}

func (alc *AppLangCreate) sqlSave(ctx context.Context) (*AppLang, error) {
	if err := alc.check(); err != nil {
		return nil, err
	}
	_node, _spec := alc.createSpec()
	if err := sqlgraph.CreateNode(ctx, alc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	alc.mutation.id = &_node.ID
	alc.mutation.done = true
	return _node, nil
}

func (alc *AppLangCreate) createSpec() (*AppLang, *sqlgraph.CreateSpec) {
	var (
		_node = &AppLang{config: alc.config}
		_spec = sqlgraph.NewCreateSpec(applang.Table, sqlgraph.NewFieldSpec(applang.FieldID, field.TypeUint32))
	)
	_spec.OnConflict = alc.conflict
	if id, ok := alc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := alc.mutation.CreatedAt(); ok {
		_spec.SetField(applang.FieldCreatedAt, field.TypeUint32, value)
		_node.CreatedAt = value
	}
	if value, ok := alc.mutation.UpdatedAt(); ok {
		_spec.SetField(applang.FieldUpdatedAt, field.TypeUint32, value)
		_node.UpdatedAt = value
	}
	if value, ok := alc.mutation.DeletedAt(); ok {
		_spec.SetField(applang.FieldDeletedAt, field.TypeUint32, value)
		_node.DeletedAt = value
	}
	if value, ok := alc.mutation.EntID(); ok {
		_spec.SetField(applang.FieldEntID, field.TypeUUID, value)
		_node.EntID = value
	}
	if value, ok := alc.mutation.AppID(); ok {
		_spec.SetField(applang.FieldAppID, field.TypeUUID, value)
		_node.AppID = value
	}
	if value, ok := alc.mutation.LangID(); ok {
		_spec.SetField(applang.FieldLangID, field.TypeUUID, value)
		_node.LangID = value
	}
	if value, ok := alc.mutation.Main(); ok {
		_spec.SetField(applang.FieldMain, field.TypeBool, value)
		_node.Main = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppLang.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppLangUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (alc *AppLangCreate) OnConflict(opts ...sql.ConflictOption) *AppLangUpsertOne {
	alc.conflict = opts
	return &AppLangUpsertOne{
		create: alc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppLang.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (alc *AppLangCreate) OnConflictColumns(columns ...string) *AppLangUpsertOne {
	alc.conflict = append(alc.conflict, sql.ConflictColumns(columns...))
	return &AppLangUpsertOne{
		create: alc,
	}
}

type (
	// AppLangUpsertOne is the builder for "upsert"-ing
	//  one AppLang node.
	AppLangUpsertOne struct {
		create *AppLangCreate
	}

	// AppLangUpsert is the "OnConflict" setter.
	AppLangUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *AppLangUpsert) SetCreatedAt(v uint32) *AppLangUpsert {
	u.Set(applang.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppLangUpsert) UpdateCreatedAt() *AppLangUpsert {
	u.SetExcluded(applang.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppLangUpsert) AddCreatedAt(v uint32) *AppLangUpsert {
	u.Add(applang.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppLangUpsert) SetUpdatedAt(v uint32) *AppLangUpsert {
	u.Set(applang.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppLangUpsert) UpdateUpdatedAt() *AppLangUpsert {
	u.SetExcluded(applang.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppLangUpsert) AddUpdatedAt(v uint32) *AppLangUpsert {
	u.Add(applang.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppLangUpsert) SetDeletedAt(v uint32) *AppLangUpsert {
	u.Set(applang.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppLangUpsert) UpdateDeletedAt() *AppLangUpsert {
	u.SetExcluded(applang.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppLangUpsert) AddDeletedAt(v uint32) *AppLangUpsert {
	u.Add(applang.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *AppLangUpsert) SetEntID(v uuid.UUID) *AppLangUpsert {
	u.Set(applang.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AppLangUpsert) UpdateEntID() *AppLangUpsert {
	u.SetExcluded(applang.FieldEntID)
	return u
}

// SetAppID sets the "app_id" field.
func (u *AppLangUpsert) SetAppID(v uuid.UUID) *AppLangUpsert {
	u.Set(applang.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppLangUpsert) UpdateAppID() *AppLangUpsert {
	u.SetExcluded(applang.FieldAppID)
	return u
}

// ClearAppID clears the value of the "app_id" field.
func (u *AppLangUpsert) ClearAppID() *AppLangUpsert {
	u.SetNull(applang.FieldAppID)
	return u
}

// SetLangID sets the "lang_id" field.
func (u *AppLangUpsert) SetLangID(v uuid.UUID) *AppLangUpsert {
	u.Set(applang.FieldLangID, v)
	return u
}

// UpdateLangID sets the "lang_id" field to the value that was provided on create.
func (u *AppLangUpsert) UpdateLangID() *AppLangUpsert {
	u.SetExcluded(applang.FieldLangID)
	return u
}

// ClearLangID clears the value of the "lang_id" field.
func (u *AppLangUpsert) ClearLangID() *AppLangUpsert {
	u.SetNull(applang.FieldLangID)
	return u
}

// SetMain sets the "main" field.
func (u *AppLangUpsert) SetMain(v bool) *AppLangUpsert {
	u.Set(applang.FieldMain, v)
	return u
}

// UpdateMain sets the "main" field to the value that was provided on create.
func (u *AppLangUpsert) UpdateMain() *AppLangUpsert {
	u.SetExcluded(applang.FieldMain)
	return u
}

// ClearMain clears the value of the "main" field.
func (u *AppLangUpsert) ClearMain() *AppLangUpsert {
	u.SetNull(applang.FieldMain)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppLang.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(applang.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AppLangUpsertOne) UpdateNewValues() *AppLangUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(applang.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppLang.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AppLangUpsertOne) Ignore() *AppLangUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppLangUpsertOne) DoNothing() *AppLangUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppLangCreate.OnConflict
// documentation for more info.
func (u *AppLangUpsertOne) Update(set func(*AppLangUpsert)) *AppLangUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppLangUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppLangUpsertOne) SetCreatedAt(v uint32) *AppLangUpsertOne {
	return u.Update(func(s *AppLangUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppLangUpsertOne) AddCreatedAt(v uint32) *AppLangUpsertOne {
	return u.Update(func(s *AppLangUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppLangUpsertOne) UpdateCreatedAt() *AppLangUpsertOne {
	return u.Update(func(s *AppLangUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppLangUpsertOne) SetUpdatedAt(v uint32) *AppLangUpsertOne {
	return u.Update(func(s *AppLangUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppLangUpsertOne) AddUpdatedAt(v uint32) *AppLangUpsertOne {
	return u.Update(func(s *AppLangUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppLangUpsertOne) UpdateUpdatedAt() *AppLangUpsertOne {
	return u.Update(func(s *AppLangUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppLangUpsertOne) SetDeletedAt(v uint32) *AppLangUpsertOne {
	return u.Update(func(s *AppLangUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppLangUpsertOne) AddDeletedAt(v uint32) *AppLangUpsertOne {
	return u.Update(func(s *AppLangUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppLangUpsertOne) UpdateDeletedAt() *AppLangUpsertOne {
	return u.Update(func(s *AppLangUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *AppLangUpsertOne) SetEntID(v uuid.UUID) *AppLangUpsertOne {
	return u.Update(func(s *AppLangUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AppLangUpsertOne) UpdateEntID() *AppLangUpsertOne {
	return u.Update(func(s *AppLangUpsert) {
		s.UpdateEntID()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppLangUpsertOne) SetAppID(v uuid.UUID) *AppLangUpsertOne {
	return u.Update(func(s *AppLangUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppLangUpsertOne) UpdateAppID() *AppLangUpsertOne {
	return u.Update(func(s *AppLangUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *AppLangUpsertOne) ClearAppID() *AppLangUpsertOne {
	return u.Update(func(s *AppLangUpsert) {
		s.ClearAppID()
	})
}

// SetLangID sets the "lang_id" field.
func (u *AppLangUpsertOne) SetLangID(v uuid.UUID) *AppLangUpsertOne {
	return u.Update(func(s *AppLangUpsert) {
		s.SetLangID(v)
	})
}

// UpdateLangID sets the "lang_id" field to the value that was provided on create.
func (u *AppLangUpsertOne) UpdateLangID() *AppLangUpsertOne {
	return u.Update(func(s *AppLangUpsert) {
		s.UpdateLangID()
	})
}

// ClearLangID clears the value of the "lang_id" field.
func (u *AppLangUpsertOne) ClearLangID() *AppLangUpsertOne {
	return u.Update(func(s *AppLangUpsert) {
		s.ClearLangID()
	})
}

// SetMain sets the "main" field.
func (u *AppLangUpsertOne) SetMain(v bool) *AppLangUpsertOne {
	return u.Update(func(s *AppLangUpsert) {
		s.SetMain(v)
	})
}

// UpdateMain sets the "main" field to the value that was provided on create.
func (u *AppLangUpsertOne) UpdateMain() *AppLangUpsertOne {
	return u.Update(func(s *AppLangUpsert) {
		s.UpdateMain()
	})
}

// ClearMain clears the value of the "main" field.
func (u *AppLangUpsertOne) ClearMain() *AppLangUpsertOne {
	return u.Update(func(s *AppLangUpsert) {
		s.ClearMain()
	})
}

// Exec executes the query.
func (u *AppLangUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for AppLangCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppLangUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppLangUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppLangUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppLangCreateBulk is the builder for creating many AppLang entities in bulk.
type AppLangCreateBulk struct {
	config
	err      error
	builders []*AppLangCreate
	conflict []sql.ConflictOption
}

// Save creates the AppLang entities in the database.
func (alcb *AppLangCreateBulk) Save(ctx context.Context) ([]*AppLang, error) {
	if alcb.err != nil {
		return nil, alcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(alcb.builders))
	nodes := make([]*AppLang, len(alcb.builders))
	mutators := make([]Mutator, len(alcb.builders))
	for i := range alcb.builders {
		func(i int, root context.Context) {
			builder := alcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppLangMutation)
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
					_, err = mutators[i+1].Mutate(root, alcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = alcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, alcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, alcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (alcb *AppLangCreateBulk) SaveX(ctx context.Context) []*AppLang {
	v, err := alcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (alcb *AppLangCreateBulk) Exec(ctx context.Context) error {
	_, err := alcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alcb *AppLangCreateBulk) ExecX(ctx context.Context) {
	if err := alcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppLang.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppLangUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (alcb *AppLangCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppLangUpsertBulk {
	alcb.conflict = opts
	return &AppLangUpsertBulk{
		create: alcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppLang.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (alcb *AppLangCreateBulk) OnConflictColumns(columns ...string) *AppLangUpsertBulk {
	alcb.conflict = append(alcb.conflict, sql.ConflictColumns(columns...))
	return &AppLangUpsertBulk{
		create: alcb,
	}
}

// AppLangUpsertBulk is the builder for "upsert"-ing
// a bulk of AppLang nodes.
type AppLangUpsertBulk struct {
	create *AppLangCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppLang.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(applang.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AppLangUpsertBulk) UpdateNewValues() *AppLangUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(applang.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppLang.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AppLangUpsertBulk) Ignore() *AppLangUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppLangUpsertBulk) DoNothing() *AppLangUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppLangCreateBulk.OnConflict
// documentation for more info.
func (u *AppLangUpsertBulk) Update(set func(*AppLangUpsert)) *AppLangUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppLangUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppLangUpsertBulk) SetCreatedAt(v uint32) *AppLangUpsertBulk {
	return u.Update(func(s *AppLangUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppLangUpsertBulk) AddCreatedAt(v uint32) *AppLangUpsertBulk {
	return u.Update(func(s *AppLangUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppLangUpsertBulk) UpdateCreatedAt() *AppLangUpsertBulk {
	return u.Update(func(s *AppLangUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppLangUpsertBulk) SetUpdatedAt(v uint32) *AppLangUpsertBulk {
	return u.Update(func(s *AppLangUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppLangUpsertBulk) AddUpdatedAt(v uint32) *AppLangUpsertBulk {
	return u.Update(func(s *AppLangUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppLangUpsertBulk) UpdateUpdatedAt() *AppLangUpsertBulk {
	return u.Update(func(s *AppLangUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppLangUpsertBulk) SetDeletedAt(v uint32) *AppLangUpsertBulk {
	return u.Update(func(s *AppLangUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppLangUpsertBulk) AddDeletedAt(v uint32) *AppLangUpsertBulk {
	return u.Update(func(s *AppLangUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppLangUpsertBulk) UpdateDeletedAt() *AppLangUpsertBulk {
	return u.Update(func(s *AppLangUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *AppLangUpsertBulk) SetEntID(v uuid.UUID) *AppLangUpsertBulk {
	return u.Update(func(s *AppLangUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AppLangUpsertBulk) UpdateEntID() *AppLangUpsertBulk {
	return u.Update(func(s *AppLangUpsert) {
		s.UpdateEntID()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppLangUpsertBulk) SetAppID(v uuid.UUID) *AppLangUpsertBulk {
	return u.Update(func(s *AppLangUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppLangUpsertBulk) UpdateAppID() *AppLangUpsertBulk {
	return u.Update(func(s *AppLangUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *AppLangUpsertBulk) ClearAppID() *AppLangUpsertBulk {
	return u.Update(func(s *AppLangUpsert) {
		s.ClearAppID()
	})
}

// SetLangID sets the "lang_id" field.
func (u *AppLangUpsertBulk) SetLangID(v uuid.UUID) *AppLangUpsertBulk {
	return u.Update(func(s *AppLangUpsert) {
		s.SetLangID(v)
	})
}

// UpdateLangID sets the "lang_id" field to the value that was provided on create.
func (u *AppLangUpsertBulk) UpdateLangID() *AppLangUpsertBulk {
	return u.Update(func(s *AppLangUpsert) {
		s.UpdateLangID()
	})
}

// ClearLangID clears the value of the "lang_id" field.
func (u *AppLangUpsertBulk) ClearLangID() *AppLangUpsertBulk {
	return u.Update(func(s *AppLangUpsert) {
		s.ClearLangID()
	})
}

// SetMain sets the "main" field.
func (u *AppLangUpsertBulk) SetMain(v bool) *AppLangUpsertBulk {
	return u.Update(func(s *AppLangUpsert) {
		s.SetMain(v)
	})
}

// UpdateMain sets the "main" field to the value that was provided on create.
func (u *AppLangUpsertBulk) UpdateMain() *AppLangUpsertBulk {
	return u.Update(func(s *AppLangUpsert) {
		s.UpdateMain()
	})
}

// ClearMain clears the value of the "main" field.
func (u *AppLangUpsertBulk) ClearMain() *AppLangUpsertBulk {
	return u.Update(func(s *AppLangUpsert) {
		s.ClearMain()
	})
}

// Exec executes the query.
func (u *AppLangUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("generated: OnConflict was set for builder %d. Set it on the AppLangCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for AppLangCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppLangUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
