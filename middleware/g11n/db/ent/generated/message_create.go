// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated/message"
	"github.com/google/uuid"
)

// MessageCreate is the builder for creating a Message entity.
type MessageCreate struct {
	config
	mutation *MessageMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (mc *MessageCreate) SetCreatedAt(u uint32) *MessageCreate {
	mc.mutation.SetCreatedAt(u)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MessageCreate) SetNillableCreatedAt(u *uint32) *MessageCreate {
	if u != nil {
		mc.SetCreatedAt(*u)
	}
	return mc
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *MessageCreate) SetUpdatedAt(u uint32) *MessageCreate {
	mc.mutation.SetUpdatedAt(u)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *MessageCreate) SetNillableUpdatedAt(u *uint32) *MessageCreate {
	if u != nil {
		mc.SetUpdatedAt(*u)
	}
	return mc
}

// SetDeletedAt sets the "deleted_at" field.
func (mc *MessageCreate) SetDeletedAt(u uint32) *MessageCreate {
	mc.mutation.SetDeletedAt(u)
	return mc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mc *MessageCreate) SetNillableDeletedAt(u *uint32) *MessageCreate {
	if u != nil {
		mc.SetDeletedAt(*u)
	}
	return mc
}

// SetEntID sets the "ent_id" field.
func (mc *MessageCreate) SetEntID(u uuid.UUID) *MessageCreate {
	mc.mutation.SetEntID(u)
	return mc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (mc *MessageCreate) SetNillableEntID(u *uuid.UUID) *MessageCreate {
	if u != nil {
		mc.SetEntID(*u)
	}
	return mc
}

// SetAppID sets the "app_id" field.
func (mc *MessageCreate) SetAppID(u uuid.UUID) *MessageCreate {
	mc.mutation.SetAppID(u)
	return mc
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (mc *MessageCreate) SetNillableAppID(u *uuid.UUID) *MessageCreate {
	if u != nil {
		mc.SetAppID(*u)
	}
	return mc
}

// SetLangID sets the "lang_id" field.
func (mc *MessageCreate) SetLangID(u uuid.UUID) *MessageCreate {
	mc.mutation.SetLangID(u)
	return mc
}

// SetNillableLangID sets the "lang_id" field if the given value is not nil.
func (mc *MessageCreate) SetNillableLangID(u *uuid.UUID) *MessageCreate {
	if u != nil {
		mc.SetLangID(*u)
	}
	return mc
}

// SetMessageID sets the "message_id" field.
func (mc *MessageCreate) SetMessageID(s string) *MessageCreate {
	mc.mutation.SetMessageID(s)
	return mc
}

// SetNillableMessageID sets the "message_id" field if the given value is not nil.
func (mc *MessageCreate) SetNillableMessageID(s *string) *MessageCreate {
	if s != nil {
		mc.SetMessageID(*s)
	}
	return mc
}

// SetMessage sets the "message" field.
func (mc *MessageCreate) SetMessage(s string) *MessageCreate {
	mc.mutation.SetMessage(s)
	return mc
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (mc *MessageCreate) SetNillableMessage(s *string) *MessageCreate {
	if s != nil {
		mc.SetMessage(*s)
	}
	return mc
}

// SetGetIndex sets the "get_index" field.
func (mc *MessageCreate) SetGetIndex(u uint32) *MessageCreate {
	mc.mutation.SetGetIndex(u)
	return mc
}

// SetNillableGetIndex sets the "get_index" field if the given value is not nil.
func (mc *MessageCreate) SetNillableGetIndex(u *uint32) *MessageCreate {
	if u != nil {
		mc.SetGetIndex(*u)
	}
	return mc
}

// SetDisabled sets the "disabled" field.
func (mc *MessageCreate) SetDisabled(b bool) *MessageCreate {
	mc.mutation.SetDisabled(b)
	return mc
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (mc *MessageCreate) SetNillableDisabled(b *bool) *MessageCreate {
	if b != nil {
		mc.SetDisabled(*b)
	}
	return mc
}

// SetShort sets the "short" field.
func (mc *MessageCreate) SetShort(s string) *MessageCreate {
	mc.mutation.SetShort(s)
	return mc
}

// SetNillableShort sets the "short" field if the given value is not nil.
func (mc *MessageCreate) SetNillableShort(s *string) *MessageCreate {
	if s != nil {
		mc.SetShort(*s)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MessageCreate) SetID(u uint32) *MessageCreate {
	mc.mutation.SetID(u)
	return mc
}

// Mutation returns the MessageMutation object of the builder.
func (mc *MessageCreate) Mutation() *MessageMutation {
	return mc.mutation
}

// Save creates the Message in the database.
func (mc *MessageCreate) Save(ctx context.Context) (*Message, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MessageCreate) SaveX(ctx context.Context) *Message {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MessageCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MessageCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MessageCreate) defaults() {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := message.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		v := message.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mc.mutation.DeletedAt(); !ok {
		v := message.DefaultDeletedAt()
		mc.mutation.SetDeletedAt(v)
	}
	if _, ok := mc.mutation.EntID(); !ok {
		v := message.DefaultEntID()
		mc.mutation.SetEntID(v)
	}
	if _, ok := mc.mutation.AppID(); !ok {
		v := message.DefaultAppID()
		mc.mutation.SetAppID(v)
	}
	if _, ok := mc.mutation.LangID(); !ok {
		v := message.DefaultLangID()
		mc.mutation.SetLangID(v)
	}
	if _, ok := mc.mutation.MessageID(); !ok {
		v := message.DefaultMessageID
		mc.mutation.SetMessageID(v)
	}
	if _, ok := mc.mutation.Message(); !ok {
		v := message.DefaultMessage
		mc.mutation.SetMessage(v)
	}
	if _, ok := mc.mutation.GetIndex(); !ok {
		v := message.DefaultGetIndex
		mc.mutation.SetGetIndex(v)
	}
	if _, ok := mc.mutation.Disabled(); !ok {
		v := message.DefaultDisabled
		mc.mutation.SetDisabled(v)
	}
	if _, ok := mc.mutation.Short(); !ok {
		v := message.DefaultShort
		mc.mutation.SetShort(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MessageCreate) check() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "Message.created_at"`)}
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "Message.updated_at"`)}
	}
	if _, ok := mc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`generated: missing required field "Message.deleted_at"`)}
	}
	if _, ok := mc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`generated: missing required field "Message.ent_id"`)}
	}
	return nil
}

func (mc *MessageCreate) sqlSave(ctx context.Context) (*Message, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MessageCreate) createSpec() (*Message, *sqlgraph.CreateSpec) {
	var (
		_node = &Message{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(message.Table, sqlgraph.NewFieldSpec(message.FieldID, field.TypeUint32))
	)
	_spec.OnConflict = mc.conflict
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.SetField(message.FieldCreatedAt, field.TypeUint32, value)
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.SetField(message.FieldUpdatedAt, field.TypeUint32, value)
		_node.UpdatedAt = value
	}
	if value, ok := mc.mutation.DeletedAt(); ok {
		_spec.SetField(message.FieldDeletedAt, field.TypeUint32, value)
		_node.DeletedAt = value
	}
	if value, ok := mc.mutation.EntID(); ok {
		_spec.SetField(message.FieldEntID, field.TypeUUID, value)
		_node.EntID = value
	}
	if value, ok := mc.mutation.AppID(); ok {
		_spec.SetField(message.FieldAppID, field.TypeUUID, value)
		_node.AppID = value
	}
	if value, ok := mc.mutation.LangID(); ok {
		_spec.SetField(message.FieldLangID, field.TypeUUID, value)
		_node.LangID = value
	}
	if value, ok := mc.mutation.MessageID(); ok {
		_spec.SetField(message.FieldMessageID, field.TypeString, value)
		_node.MessageID = value
	}
	if value, ok := mc.mutation.Message(); ok {
		_spec.SetField(message.FieldMessage, field.TypeString, value)
		_node.Message = value
	}
	if value, ok := mc.mutation.GetIndex(); ok {
		_spec.SetField(message.FieldGetIndex, field.TypeUint32, value)
		_node.GetIndex = value
	}
	if value, ok := mc.mutation.Disabled(); ok {
		_spec.SetField(message.FieldDisabled, field.TypeBool, value)
		_node.Disabled = value
	}
	if value, ok := mc.mutation.Short(); ok {
		_spec.SetField(message.FieldShort, field.TypeString, value)
		_node.Short = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Message.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MessageUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (mc *MessageCreate) OnConflict(opts ...sql.ConflictOption) *MessageUpsertOne {
	mc.conflict = opts
	return &MessageUpsertOne{
		create: mc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mc *MessageCreate) OnConflictColumns(columns ...string) *MessageUpsertOne {
	mc.conflict = append(mc.conflict, sql.ConflictColumns(columns...))
	return &MessageUpsertOne{
		create: mc,
	}
}

type (
	// MessageUpsertOne is the builder for "upsert"-ing
	//  one Message node.
	MessageUpsertOne struct {
		create *MessageCreate
	}

	// MessageUpsert is the "OnConflict" setter.
	MessageUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *MessageUpsert) SetCreatedAt(v uint32) *MessageUpsert {
	u.Set(message.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *MessageUpsert) UpdateCreatedAt() *MessageUpsert {
	u.SetExcluded(message.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *MessageUpsert) AddCreatedAt(v uint32) *MessageUpsert {
	u.Add(message.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MessageUpsert) SetUpdatedAt(v uint32) *MessageUpsert {
	u.Set(message.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MessageUpsert) UpdateUpdatedAt() *MessageUpsert {
	u.SetExcluded(message.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *MessageUpsert) AddUpdatedAt(v uint32) *MessageUpsert {
	u.Add(message.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *MessageUpsert) SetDeletedAt(v uint32) *MessageUpsert {
	u.Set(message.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *MessageUpsert) UpdateDeletedAt() *MessageUpsert {
	u.SetExcluded(message.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *MessageUpsert) AddDeletedAt(v uint32) *MessageUpsert {
	u.Add(message.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *MessageUpsert) SetEntID(v uuid.UUID) *MessageUpsert {
	u.Set(message.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *MessageUpsert) UpdateEntID() *MessageUpsert {
	u.SetExcluded(message.FieldEntID)
	return u
}

// SetAppID sets the "app_id" field.
func (u *MessageUpsert) SetAppID(v uuid.UUID) *MessageUpsert {
	u.Set(message.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *MessageUpsert) UpdateAppID() *MessageUpsert {
	u.SetExcluded(message.FieldAppID)
	return u
}

// ClearAppID clears the value of the "app_id" field.
func (u *MessageUpsert) ClearAppID() *MessageUpsert {
	u.SetNull(message.FieldAppID)
	return u
}

// SetLangID sets the "lang_id" field.
func (u *MessageUpsert) SetLangID(v uuid.UUID) *MessageUpsert {
	u.Set(message.FieldLangID, v)
	return u
}

// UpdateLangID sets the "lang_id" field to the value that was provided on create.
func (u *MessageUpsert) UpdateLangID() *MessageUpsert {
	u.SetExcluded(message.FieldLangID)
	return u
}

// ClearLangID clears the value of the "lang_id" field.
func (u *MessageUpsert) ClearLangID() *MessageUpsert {
	u.SetNull(message.FieldLangID)
	return u
}

// SetMessageID sets the "message_id" field.
func (u *MessageUpsert) SetMessageID(v string) *MessageUpsert {
	u.Set(message.FieldMessageID, v)
	return u
}

// UpdateMessageID sets the "message_id" field to the value that was provided on create.
func (u *MessageUpsert) UpdateMessageID() *MessageUpsert {
	u.SetExcluded(message.FieldMessageID)
	return u
}

// ClearMessageID clears the value of the "message_id" field.
func (u *MessageUpsert) ClearMessageID() *MessageUpsert {
	u.SetNull(message.FieldMessageID)
	return u
}

// SetMessage sets the "message" field.
func (u *MessageUpsert) SetMessage(v string) *MessageUpsert {
	u.Set(message.FieldMessage, v)
	return u
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *MessageUpsert) UpdateMessage() *MessageUpsert {
	u.SetExcluded(message.FieldMessage)
	return u
}

// ClearMessage clears the value of the "message" field.
func (u *MessageUpsert) ClearMessage() *MessageUpsert {
	u.SetNull(message.FieldMessage)
	return u
}

// SetGetIndex sets the "get_index" field.
func (u *MessageUpsert) SetGetIndex(v uint32) *MessageUpsert {
	u.Set(message.FieldGetIndex, v)
	return u
}

// UpdateGetIndex sets the "get_index" field to the value that was provided on create.
func (u *MessageUpsert) UpdateGetIndex() *MessageUpsert {
	u.SetExcluded(message.FieldGetIndex)
	return u
}

// AddGetIndex adds v to the "get_index" field.
func (u *MessageUpsert) AddGetIndex(v uint32) *MessageUpsert {
	u.Add(message.FieldGetIndex, v)
	return u
}

// ClearGetIndex clears the value of the "get_index" field.
func (u *MessageUpsert) ClearGetIndex() *MessageUpsert {
	u.SetNull(message.FieldGetIndex)
	return u
}

// SetDisabled sets the "disabled" field.
func (u *MessageUpsert) SetDisabled(v bool) *MessageUpsert {
	u.Set(message.FieldDisabled, v)
	return u
}

// UpdateDisabled sets the "disabled" field to the value that was provided on create.
func (u *MessageUpsert) UpdateDisabled() *MessageUpsert {
	u.SetExcluded(message.FieldDisabled)
	return u
}

// ClearDisabled clears the value of the "disabled" field.
func (u *MessageUpsert) ClearDisabled() *MessageUpsert {
	u.SetNull(message.FieldDisabled)
	return u
}

// SetShort sets the "short" field.
func (u *MessageUpsert) SetShort(v string) *MessageUpsert {
	u.Set(message.FieldShort, v)
	return u
}

// UpdateShort sets the "short" field to the value that was provided on create.
func (u *MessageUpsert) UpdateShort() *MessageUpsert {
	u.SetExcluded(message.FieldShort)
	return u
}

// ClearShort clears the value of the "short" field.
func (u *MessageUpsert) ClearShort() *MessageUpsert {
	u.SetNull(message.FieldShort)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(message.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MessageUpsertOne) UpdateNewValues() *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(message.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Message.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MessageUpsertOne) Ignore() *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MessageUpsertOne) DoNothing() *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MessageCreate.OnConflict
// documentation for more info.
func (u *MessageUpsertOne) Update(set func(*MessageUpsert)) *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MessageUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *MessageUpsertOne) SetCreatedAt(v uint32) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *MessageUpsertOne) AddCreatedAt(v uint32) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateCreatedAt() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MessageUpsertOne) SetUpdatedAt(v uint32) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *MessageUpsertOne) AddUpdatedAt(v uint32) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateUpdatedAt() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *MessageUpsertOne) SetDeletedAt(v uint32) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *MessageUpsertOne) AddDeletedAt(v uint32) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateDeletedAt() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *MessageUpsertOne) SetEntID(v uuid.UUID) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateEntID() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateEntID()
	})
}

// SetAppID sets the "app_id" field.
func (u *MessageUpsertOne) SetAppID(v uuid.UUID) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateAppID() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *MessageUpsertOne) ClearAppID() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.ClearAppID()
	})
}

// SetLangID sets the "lang_id" field.
func (u *MessageUpsertOne) SetLangID(v uuid.UUID) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetLangID(v)
	})
}

// UpdateLangID sets the "lang_id" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateLangID() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateLangID()
	})
}

// ClearLangID clears the value of the "lang_id" field.
func (u *MessageUpsertOne) ClearLangID() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.ClearLangID()
	})
}

// SetMessageID sets the "message_id" field.
func (u *MessageUpsertOne) SetMessageID(v string) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetMessageID(v)
	})
}

// UpdateMessageID sets the "message_id" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateMessageID() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateMessageID()
	})
}

// ClearMessageID clears the value of the "message_id" field.
func (u *MessageUpsertOne) ClearMessageID() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.ClearMessageID()
	})
}

// SetMessage sets the "message" field.
func (u *MessageUpsertOne) SetMessage(v string) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateMessage() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateMessage()
	})
}

// ClearMessage clears the value of the "message" field.
func (u *MessageUpsertOne) ClearMessage() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.ClearMessage()
	})
}

// SetGetIndex sets the "get_index" field.
func (u *MessageUpsertOne) SetGetIndex(v uint32) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetGetIndex(v)
	})
}

// AddGetIndex adds v to the "get_index" field.
func (u *MessageUpsertOne) AddGetIndex(v uint32) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.AddGetIndex(v)
	})
}

// UpdateGetIndex sets the "get_index" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateGetIndex() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateGetIndex()
	})
}

// ClearGetIndex clears the value of the "get_index" field.
func (u *MessageUpsertOne) ClearGetIndex() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.ClearGetIndex()
	})
}

// SetDisabled sets the "disabled" field.
func (u *MessageUpsertOne) SetDisabled(v bool) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetDisabled(v)
	})
}

// UpdateDisabled sets the "disabled" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateDisabled() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateDisabled()
	})
}

// ClearDisabled clears the value of the "disabled" field.
func (u *MessageUpsertOne) ClearDisabled() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.ClearDisabled()
	})
}

// SetShort sets the "short" field.
func (u *MessageUpsertOne) SetShort(v string) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetShort(v)
	})
}

// UpdateShort sets the "short" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateShort() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateShort()
	})
}

// ClearShort clears the value of the "short" field.
func (u *MessageUpsertOne) ClearShort() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.ClearShort()
	})
}

// Exec executes the query.
func (u *MessageUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for MessageCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MessageUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MessageUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MessageUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MessageCreateBulk is the builder for creating many Message entities in bulk.
type MessageCreateBulk struct {
	config
	err      error
	builders []*MessageCreate
	conflict []sql.ConflictOption
}

// Save creates the Message entities in the database.
func (mcb *MessageCreateBulk) Save(ctx context.Context) ([]*Message, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Message, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MessageMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MessageCreateBulk) SaveX(ctx context.Context) []*Message {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MessageCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MessageCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Message.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MessageUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (mcb *MessageCreateBulk) OnConflict(opts ...sql.ConflictOption) *MessageUpsertBulk {
	mcb.conflict = opts
	return &MessageUpsertBulk{
		create: mcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mcb *MessageCreateBulk) OnConflictColumns(columns ...string) *MessageUpsertBulk {
	mcb.conflict = append(mcb.conflict, sql.ConflictColumns(columns...))
	return &MessageUpsertBulk{
		create: mcb,
	}
}

// MessageUpsertBulk is the builder for "upsert"-ing
// a bulk of Message nodes.
type MessageUpsertBulk struct {
	create *MessageCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(message.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MessageUpsertBulk) UpdateNewValues() *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(message.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MessageUpsertBulk) Ignore() *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MessageUpsertBulk) DoNothing() *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MessageCreateBulk.OnConflict
// documentation for more info.
func (u *MessageUpsertBulk) Update(set func(*MessageUpsert)) *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MessageUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *MessageUpsertBulk) SetCreatedAt(v uint32) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *MessageUpsertBulk) AddCreatedAt(v uint32) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateCreatedAt() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *MessageUpsertBulk) SetUpdatedAt(v uint32) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *MessageUpsertBulk) AddUpdatedAt(v uint32) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateUpdatedAt() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *MessageUpsertBulk) SetDeletedAt(v uint32) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *MessageUpsertBulk) AddDeletedAt(v uint32) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateDeletedAt() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *MessageUpsertBulk) SetEntID(v uuid.UUID) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateEntID() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateEntID()
	})
}

// SetAppID sets the "app_id" field.
func (u *MessageUpsertBulk) SetAppID(v uuid.UUID) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateAppID() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *MessageUpsertBulk) ClearAppID() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.ClearAppID()
	})
}

// SetLangID sets the "lang_id" field.
func (u *MessageUpsertBulk) SetLangID(v uuid.UUID) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetLangID(v)
	})
}

// UpdateLangID sets the "lang_id" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateLangID() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateLangID()
	})
}

// ClearLangID clears the value of the "lang_id" field.
func (u *MessageUpsertBulk) ClearLangID() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.ClearLangID()
	})
}

// SetMessageID sets the "message_id" field.
func (u *MessageUpsertBulk) SetMessageID(v string) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetMessageID(v)
	})
}

// UpdateMessageID sets the "message_id" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateMessageID() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateMessageID()
	})
}

// ClearMessageID clears the value of the "message_id" field.
func (u *MessageUpsertBulk) ClearMessageID() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.ClearMessageID()
	})
}

// SetMessage sets the "message" field.
func (u *MessageUpsertBulk) SetMessage(v string) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateMessage() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateMessage()
	})
}

// ClearMessage clears the value of the "message" field.
func (u *MessageUpsertBulk) ClearMessage() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.ClearMessage()
	})
}

// SetGetIndex sets the "get_index" field.
func (u *MessageUpsertBulk) SetGetIndex(v uint32) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetGetIndex(v)
	})
}

// AddGetIndex adds v to the "get_index" field.
func (u *MessageUpsertBulk) AddGetIndex(v uint32) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.AddGetIndex(v)
	})
}

// UpdateGetIndex sets the "get_index" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateGetIndex() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateGetIndex()
	})
}

// ClearGetIndex clears the value of the "get_index" field.
func (u *MessageUpsertBulk) ClearGetIndex() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.ClearGetIndex()
	})
}

// SetDisabled sets the "disabled" field.
func (u *MessageUpsertBulk) SetDisabled(v bool) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetDisabled(v)
	})
}

// UpdateDisabled sets the "disabled" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateDisabled() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateDisabled()
	})
}

// ClearDisabled clears the value of the "disabled" field.
func (u *MessageUpsertBulk) ClearDisabled() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.ClearDisabled()
	})
}

// SetShort sets the "short" field.
func (u *MessageUpsertBulk) SetShort(v string) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetShort(v)
	})
}

// UpdateShort sets the "short" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateShort() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateShort()
	})
}

// ClearShort clears the value of the "short" field.
func (u *MessageUpsertBulk) ClearShort() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.ClearShort()
	})
}

// Exec executes the query.
func (u *MessageUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("generated: OnConflict was set for builder %d. Set it on the MessageCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for MessageCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MessageUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
