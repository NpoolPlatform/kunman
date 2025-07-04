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
	"github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/emailtemplate"
	"github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/predicate"
	"github.com/google/uuid"
)

// EmailTemplateUpdate is the builder for updating EmailTemplate entities.
type EmailTemplateUpdate struct {
	config
	hooks     []Hook
	mutation  *EmailTemplateMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the EmailTemplateUpdate builder.
func (etu *EmailTemplateUpdate) Where(ps ...predicate.EmailTemplate) *EmailTemplateUpdate {
	etu.mutation.Where(ps...)
	return etu
}

// SetCreatedAt sets the "created_at" field.
func (etu *EmailTemplateUpdate) SetCreatedAt(u uint32) *EmailTemplateUpdate {
	etu.mutation.ResetCreatedAt()
	etu.mutation.SetCreatedAt(u)
	return etu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (etu *EmailTemplateUpdate) SetNillableCreatedAt(u *uint32) *EmailTemplateUpdate {
	if u != nil {
		etu.SetCreatedAt(*u)
	}
	return etu
}

// AddCreatedAt adds u to the "created_at" field.
func (etu *EmailTemplateUpdate) AddCreatedAt(u int32) *EmailTemplateUpdate {
	etu.mutation.AddCreatedAt(u)
	return etu
}

// SetUpdatedAt sets the "updated_at" field.
func (etu *EmailTemplateUpdate) SetUpdatedAt(u uint32) *EmailTemplateUpdate {
	etu.mutation.ResetUpdatedAt()
	etu.mutation.SetUpdatedAt(u)
	return etu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (etu *EmailTemplateUpdate) AddUpdatedAt(u int32) *EmailTemplateUpdate {
	etu.mutation.AddUpdatedAt(u)
	return etu
}

// SetDeletedAt sets the "deleted_at" field.
func (etu *EmailTemplateUpdate) SetDeletedAt(u uint32) *EmailTemplateUpdate {
	etu.mutation.ResetDeletedAt()
	etu.mutation.SetDeletedAt(u)
	return etu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (etu *EmailTemplateUpdate) SetNillableDeletedAt(u *uint32) *EmailTemplateUpdate {
	if u != nil {
		etu.SetDeletedAt(*u)
	}
	return etu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (etu *EmailTemplateUpdate) AddDeletedAt(u int32) *EmailTemplateUpdate {
	etu.mutation.AddDeletedAt(u)
	return etu
}

// SetEntID sets the "ent_id" field.
func (etu *EmailTemplateUpdate) SetEntID(u uuid.UUID) *EmailTemplateUpdate {
	etu.mutation.SetEntID(u)
	return etu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (etu *EmailTemplateUpdate) SetNillableEntID(u *uuid.UUID) *EmailTemplateUpdate {
	if u != nil {
		etu.SetEntID(*u)
	}
	return etu
}

// SetAppID sets the "app_id" field.
func (etu *EmailTemplateUpdate) SetAppID(u uuid.UUID) *EmailTemplateUpdate {
	etu.mutation.SetAppID(u)
	return etu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (etu *EmailTemplateUpdate) SetNillableAppID(u *uuid.UUID) *EmailTemplateUpdate {
	if u != nil {
		etu.SetAppID(*u)
	}
	return etu
}

// SetLangID sets the "lang_id" field.
func (etu *EmailTemplateUpdate) SetLangID(u uuid.UUID) *EmailTemplateUpdate {
	etu.mutation.SetLangID(u)
	return etu
}

// SetNillableLangID sets the "lang_id" field if the given value is not nil.
func (etu *EmailTemplateUpdate) SetNillableLangID(u *uuid.UUID) *EmailTemplateUpdate {
	if u != nil {
		etu.SetLangID(*u)
	}
	return etu
}

// SetDefaultToUsername sets the "default_to_username" field.
func (etu *EmailTemplateUpdate) SetDefaultToUsername(s string) *EmailTemplateUpdate {
	etu.mutation.SetDefaultToUsername(s)
	return etu
}

// SetNillableDefaultToUsername sets the "default_to_username" field if the given value is not nil.
func (etu *EmailTemplateUpdate) SetNillableDefaultToUsername(s *string) *EmailTemplateUpdate {
	if s != nil {
		etu.SetDefaultToUsername(*s)
	}
	return etu
}

// SetUsedFor sets the "used_for" field.
func (etu *EmailTemplateUpdate) SetUsedFor(s string) *EmailTemplateUpdate {
	etu.mutation.SetUsedFor(s)
	return etu
}

// SetNillableUsedFor sets the "used_for" field if the given value is not nil.
func (etu *EmailTemplateUpdate) SetNillableUsedFor(s *string) *EmailTemplateUpdate {
	if s != nil {
		etu.SetUsedFor(*s)
	}
	return etu
}

// ClearUsedFor clears the value of the "used_for" field.
func (etu *EmailTemplateUpdate) ClearUsedFor() *EmailTemplateUpdate {
	etu.mutation.ClearUsedFor()
	return etu
}

// SetSender sets the "sender" field.
func (etu *EmailTemplateUpdate) SetSender(s string) *EmailTemplateUpdate {
	etu.mutation.SetSender(s)
	return etu
}

// SetNillableSender sets the "sender" field if the given value is not nil.
func (etu *EmailTemplateUpdate) SetNillableSender(s *string) *EmailTemplateUpdate {
	if s != nil {
		etu.SetSender(*s)
	}
	return etu
}

// ClearSender clears the value of the "sender" field.
func (etu *EmailTemplateUpdate) ClearSender() *EmailTemplateUpdate {
	etu.mutation.ClearSender()
	return etu
}

// SetReplyTos sets the "reply_tos" field.
func (etu *EmailTemplateUpdate) SetReplyTos(s []string) *EmailTemplateUpdate {
	etu.mutation.SetReplyTos(s)
	return etu
}

// AppendReplyTos appends s to the "reply_tos" field.
func (etu *EmailTemplateUpdate) AppendReplyTos(s []string) *EmailTemplateUpdate {
	etu.mutation.AppendReplyTos(s)
	return etu
}

// ClearReplyTos clears the value of the "reply_tos" field.
func (etu *EmailTemplateUpdate) ClearReplyTos() *EmailTemplateUpdate {
	etu.mutation.ClearReplyTos()
	return etu
}

// SetCcTos sets the "cc_tos" field.
func (etu *EmailTemplateUpdate) SetCcTos(s []string) *EmailTemplateUpdate {
	etu.mutation.SetCcTos(s)
	return etu
}

// AppendCcTos appends s to the "cc_tos" field.
func (etu *EmailTemplateUpdate) AppendCcTos(s []string) *EmailTemplateUpdate {
	etu.mutation.AppendCcTos(s)
	return etu
}

// ClearCcTos clears the value of the "cc_tos" field.
func (etu *EmailTemplateUpdate) ClearCcTos() *EmailTemplateUpdate {
	etu.mutation.ClearCcTos()
	return etu
}

// SetSubject sets the "subject" field.
func (etu *EmailTemplateUpdate) SetSubject(s string) *EmailTemplateUpdate {
	etu.mutation.SetSubject(s)
	return etu
}

// SetNillableSubject sets the "subject" field if the given value is not nil.
func (etu *EmailTemplateUpdate) SetNillableSubject(s *string) *EmailTemplateUpdate {
	if s != nil {
		etu.SetSubject(*s)
	}
	return etu
}

// ClearSubject clears the value of the "subject" field.
func (etu *EmailTemplateUpdate) ClearSubject() *EmailTemplateUpdate {
	etu.mutation.ClearSubject()
	return etu
}

// SetBody sets the "body" field.
func (etu *EmailTemplateUpdate) SetBody(s string) *EmailTemplateUpdate {
	etu.mutation.SetBody(s)
	return etu
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (etu *EmailTemplateUpdate) SetNillableBody(s *string) *EmailTemplateUpdate {
	if s != nil {
		etu.SetBody(*s)
	}
	return etu
}

// ClearBody clears the value of the "body" field.
func (etu *EmailTemplateUpdate) ClearBody() *EmailTemplateUpdate {
	etu.mutation.ClearBody()
	return etu
}

// Mutation returns the EmailTemplateMutation object of the builder.
func (etu *EmailTemplateUpdate) Mutation() *EmailTemplateMutation {
	return etu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (etu *EmailTemplateUpdate) Save(ctx context.Context) (int, error) {
	etu.defaults()
	return withHooks(ctx, etu.sqlSave, etu.mutation, etu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (etu *EmailTemplateUpdate) SaveX(ctx context.Context) int {
	affected, err := etu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (etu *EmailTemplateUpdate) Exec(ctx context.Context) error {
	_, err := etu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (etu *EmailTemplateUpdate) ExecX(ctx context.Context) {
	if err := etu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (etu *EmailTemplateUpdate) defaults() {
	if _, ok := etu.mutation.UpdatedAt(); !ok {
		v := emailtemplate.UpdateDefaultUpdatedAt()
		etu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (etu *EmailTemplateUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *EmailTemplateUpdate {
	etu.modifiers = append(etu.modifiers, modifiers...)
	return etu
}

func (etu *EmailTemplateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(emailtemplate.Table, emailtemplate.Columns, sqlgraph.NewFieldSpec(emailtemplate.FieldID, field.TypeUint32))
	if ps := etu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := etu.mutation.CreatedAt(); ok {
		_spec.SetField(emailtemplate.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := etu.mutation.AddedCreatedAt(); ok {
		_spec.AddField(emailtemplate.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := etu.mutation.UpdatedAt(); ok {
		_spec.SetField(emailtemplate.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := etu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(emailtemplate.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := etu.mutation.DeletedAt(); ok {
		_spec.SetField(emailtemplate.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := etu.mutation.AddedDeletedAt(); ok {
		_spec.AddField(emailtemplate.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := etu.mutation.EntID(); ok {
		_spec.SetField(emailtemplate.FieldEntID, field.TypeUUID, value)
	}
	if value, ok := etu.mutation.AppID(); ok {
		_spec.SetField(emailtemplate.FieldAppID, field.TypeUUID, value)
	}
	if value, ok := etu.mutation.LangID(); ok {
		_spec.SetField(emailtemplate.FieldLangID, field.TypeUUID, value)
	}
	if value, ok := etu.mutation.DefaultToUsername(); ok {
		_spec.SetField(emailtemplate.FieldDefaultToUsername, field.TypeString, value)
	}
	if value, ok := etu.mutation.UsedFor(); ok {
		_spec.SetField(emailtemplate.FieldUsedFor, field.TypeString, value)
	}
	if etu.mutation.UsedForCleared() {
		_spec.ClearField(emailtemplate.FieldUsedFor, field.TypeString)
	}
	if value, ok := etu.mutation.Sender(); ok {
		_spec.SetField(emailtemplate.FieldSender, field.TypeString, value)
	}
	if etu.mutation.SenderCleared() {
		_spec.ClearField(emailtemplate.FieldSender, field.TypeString)
	}
	if value, ok := etu.mutation.ReplyTos(); ok {
		_spec.SetField(emailtemplate.FieldReplyTos, field.TypeJSON, value)
	}
	if value, ok := etu.mutation.AppendedReplyTos(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, emailtemplate.FieldReplyTos, value)
		})
	}
	if etu.mutation.ReplyTosCleared() {
		_spec.ClearField(emailtemplate.FieldReplyTos, field.TypeJSON)
	}
	if value, ok := etu.mutation.CcTos(); ok {
		_spec.SetField(emailtemplate.FieldCcTos, field.TypeJSON, value)
	}
	if value, ok := etu.mutation.AppendedCcTos(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, emailtemplate.FieldCcTos, value)
		})
	}
	if etu.mutation.CcTosCleared() {
		_spec.ClearField(emailtemplate.FieldCcTos, field.TypeJSON)
	}
	if value, ok := etu.mutation.Subject(); ok {
		_spec.SetField(emailtemplate.FieldSubject, field.TypeString, value)
	}
	if etu.mutation.SubjectCleared() {
		_spec.ClearField(emailtemplate.FieldSubject, field.TypeString)
	}
	if value, ok := etu.mutation.Body(); ok {
		_spec.SetField(emailtemplate.FieldBody, field.TypeString, value)
	}
	if etu.mutation.BodyCleared() {
		_spec.ClearField(emailtemplate.FieldBody, field.TypeString)
	}
	_spec.AddModifiers(etu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, etu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{emailtemplate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	etu.mutation.done = true
	return n, nil
}

// EmailTemplateUpdateOne is the builder for updating a single EmailTemplate entity.
type EmailTemplateUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *EmailTemplateMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (etuo *EmailTemplateUpdateOne) SetCreatedAt(u uint32) *EmailTemplateUpdateOne {
	etuo.mutation.ResetCreatedAt()
	etuo.mutation.SetCreatedAt(u)
	return etuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (etuo *EmailTemplateUpdateOne) SetNillableCreatedAt(u *uint32) *EmailTemplateUpdateOne {
	if u != nil {
		etuo.SetCreatedAt(*u)
	}
	return etuo
}

// AddCreatedAt adds u to the "created_at" field.
func (etuo *EmailTemplateUpdateOne) AddCreatedAt(u int32) *EmailTemplateUpdateOne {
	etuo.mutation.AddCreatedAt(u)
	return etuo
}

// SetUpdatedAt sets the "updated_at" field.
func (etuo *EmailTemplateUpdateOne) SetUpdatedAt(u uint32) *EmailTemplateUpdateOne {
	etuo.mutation.ResetUpdatedAt()
	etuo.mutation.SetUpdatedAt(u)
	return etuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (etuo *EmailTemplateUpdateOne) AddUpdatedAt(u int32) *EmailTemplateUpdateOne {
	etuo.mutation.AddUpdatedAt(u)
	return etuo
}

// SetDeletedAt sets the "deleted_at" field.
func (etuo *EmailTemplateUpdateOne) SetDeletedAt(u uint32) *EmailTemplateUpdateOne {
	etuo.mutation.ResetDeletedAt()
	etuo.mutation.SetDeletedAt(u)
	return etuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (etuo *EmailTemplateUpdateOne) SetNillableDeletedAt(u *uint32) *EmailTemplateUpdateOne {
	if u != nil {
		etuo.SetDeletedAt(*u)
	}
	return etuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (etuo *EmailTemplateUpdateOne) AddDeletedAt(u int32) *EmailTemplateUpdateOne {
	etuo.mutation.AddDeletedAt(u)
	return etuo
}

// SetEntID sets the "ent_id" field.
func (etuo *EmailTemplateUpdateOne) SetEntID(u uuid.UUID) *EmailTemplateUpdateOne {
	etuo.mutation.SetEntID(u)
	return etuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (etuo *EmailTemplateUpdateOne) SetNillableEntID(u *uuid.UUID) *EmailTemplateUpdateOne {
	if u != nil {
		etuo.SetEntID(*u)
	}
	return etuo
}

// SetAppID sets the "app_id" field.
func (etuo *EmailTemplateUpdateOne) SetAppID(u uuid.UUID) *EmailTemplateUpdateOne {
	etuo.mutation.SetAppID(u)
	return etuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (etuo *EmailTemplateUpdateOne) SetNillableAppID(u *uuid.UUID) *EmailTemplateUpdateOne {
	if u != nil {
		etuo.SetAppID(*u)
	}
	return etuo
}

// SetLangID sets the "lang_id" field.
func (etuo *EmailTemplateUpdateOne) SetLangID(u uuid.UUID) *EmailTemplateUpdateOne {
	etuo.mutation.SetLangID(u)
	return etuo
}

// SetNillableLangID sets the "lang_id" field if the given value is not nil.
func (etuo *EmailTemplateUpdateOne) SetNillableLangID(u *uuid.UUID) *EmailTemplateUpdateOne {
	if u != nil {
		etuo.SetLangID(*u)
	}
	return etuo
}

// SetDefaultToUsername sets the "default_to_username" field.
func (etuo *EmailTemplateUpdateOne) SetDefaultToUsername(s string) *EmailTemplateUpdateOne {
	etuo.mutation.SetDefaultToUsername(s)
	return etuo
}

// SetNillableDefaultToUsername sets the "default_to_username" field if the given value is not nil.
func (etuo *EmailTemplateUpdateOne) SetNillableDefaultToUsername(s *string) *EmailTemplateUpdateOne {
	if s != nil {
		etuo.SetDefaultToUsername(*s)
	}
	return etuo
}

// SetUsedFor sets the "used_for" field.
func (etuo *EmailTemplateUpdateOne) SetUsedFor(s string) *EmailTemplateUpdateOne {
	etuo.mutation.SetUsedFor(s)
	return etuo
}

// SetNillableUsedFor sets the "used_for" field if the given value is not nil.
func (etuo *EmailTemplateUpdateOne) SetNillableUsedFor(s *string) *EmailTemplateUpdateOne {
	if s != nil {
		etuo.SetUsedFor(*s)
	}
	return etuo
}

// ClearUsedFor clears the value of the "used_for" field.
func (etuo *EmailTemplateUpdateOne) ClearUsedFor() *EmailTemplateUpdateOne {
	etuo.mutation.ClearUsedFor()
	return etuo
}

// SetSender sets the "sender" field.
func (etuo *EmailTemplateUpdateOne) SetSender(s string) *EmailTemplateUpdateOne {
	etuo.mutation.SetSender(s)
	return etuo
}

// SetNillableSender sets the "sender" field if the given value is not nil.
func (etuo *EmailTemplateUpdateOne) SetNillableSender(s *string) *EmailTemplateUpdateOne {
	if s != nil {
		etuo.SetSender(*s)
	}
	return etuo
}

// ClearSender clears the value of the "sender" field.
func (etuo *EmailTemplateUpdateOne) ClearSender() *EmailTemplateUpdateOne {
	etuo.mutation.ClearSender()
	return etuo
}

// SetReplyTos sets the "reply_tos" field.
func (etuo *EmailTemplateUpdateOne) SetReplyTos(s []string) *EmailTemplateUpdateOne {
	etuo.mutation.SetReplyTos(s)
	return etuo
}

// AppendReplyTos appends s to the "reply_tos" field.
func (etuo *EmailTemplateUpdateOne) AppendReplyTos(s []string) *EmailTemplateUpdateOne {
	etuo.mutation.AppendReplyTos(s)
	return etuo
}

// ClearReplyTos clears the value of the "reply_tos" field.
func (etuo *EmailTemplateUpdateOne) ClearReplyTos() *EmailTemplateUpdateOne {
	etuo.mutation.ClearReplyTos()
	return etuo
}

// SetCcTos sets the "cc_tos" field.
func (etuo *EmailTemplateUpdateOne) SetCcTos(s []string) *EmailTemplateUpdateOne {
	etuo.mutation.SetCcTos(s)
	return etuo
}

// AppendCcTos appends s to the "cc_tos" field.
func (etuo *EmailTemplateUpdateOne) AppendCcTos(s []string) *EmailTemplateUpdateOne {
	etuo.mutation.AppendCcTos(s)
	return etuo
}

// ClearCcTos clears the value of the "cc_tos" field.
func (etuo *EmailTemplateUpdateOne) ClearCcTos() *EmailTemplateUpdateOne {
	etuo.mutation.ClearCcTos()
	return etuo
}

// SetSubject sets the "subject" field.
func (etuo *EmailTemplateUpdateOne) SetSubject(s string) *EmailTemplateUpdateOne {
	etuo.mutation.SetSubject(s)
	return etuo
}

// SetNillableSubject sets the "subject" field if the given value is not nil.
func (etuo *EmailTemplateUpdateOne) SetNillableSubject(s *string) *EmailTemplateUpdateOne {
	if s != nil {
		etuo.SetSubject(*s)
	}
	return etuo
}

// ClearSubject clears the value of the "subject" field.
func (etuo *EmailTemplateUpdateOne) ClearSubject() *EmailTemplateUpdateOne {
	etuo.mutation.ClearSubject()
	return etuo
}

// SetBody sets the "body" field.
func (etuo *EmailTemplateUpdateOne) SetBody(s string) *EmailTemplateUpdateOne {
	etuo.mutation.SetBody(s)
	return etuo
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (etuo *EmailTemplateUpdateOne) SetNillableBody(s *string) *EmailTemplateUpdateOne {
	if s != nil {
		etuo.SetBody(*s)
	}
	return etuo
}

// ClearBody clears the value of the "body" field.
func (etuo *EmailTemplateUpdateOne) ClearBody() *EmailTemplateUpdateOne {
	etuo.mutation.ClearBody()
	return etuo
}

// Mutation returns the EmailTemplateMutation object of the builder.
func (etuo *EmailTemplateUpdateOne) Mutation() *EmailTemplateMutation {
	return etuo.mutation
}

// Where appends a list predicates to the EmailTemplateUpdate builder.
func (etuo *EmailTemplateUpdateOne) Where(ps ...predicate.EmailTemplate) *EmailTemplateUpdateOne {
	etuo.mutation.Where(ps...)
	return etuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (etuo *EmailTemplateUpdateOne) Select(field string, fields ...string) *EmailTemplateUpdateOne {
	etuo.fields = append([]string{field}, fields...)
	return etuo
}

// Save executes the query and returns the updated EmailTemplate entity.
func (etuo *EmailTemplateUpdateOne) Save(ctx context.Context) (*EmailTemplate, error) {
	etuo.defaults()
	return withHooks(ctx, etuo.sqlSave, etuo.mutation, etuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (etuo *EmailTemplateUpdateOne) SaveX(ctx context.Context) *EmailTemplate {
	node, err := etuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (etuo *EmailTemplateUpdateOne) Exec(ctx context.Context) error {
	_, err := etuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (etuo *EmailTemplateUpdateOne) ExecX(ctx context.Context) {
	if err := etuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (etuo *EmailTemplateUpdateOne) defaults() {
	if _, ok := etuo.mutation.UpdatedAt(); !ok {
		v := emailtemplate.UpdateDefaultUpdatedAt()
		etuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (etuo *EmailTemplateUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *EmailTemplateUpdateOne {
	etuo.modifiers = append(etuo.modifiers, modifiers...)
	return etuo
}

func (etuo *EmailTemplateUpdateOne) sqlSave(ctx context.Context) (_node *EmailTemplate, err error) {
	_spec := sqlgraph.NewUpdateSpec(emailtemplate.Table, emailtemplate.Columns, sqlgraph.NewFieldSpec(emailtemplate.FieldID, field.TypeUint32))
	id, ok := etuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "EmailTemplate.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := etuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, emailtemplate.FieldID)
		for _, f := range fields {
			if !emailtemplate.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != emailtemplate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := etuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := etuo.mutation.CreatedAt(); ok {
		_spec.SetField(emailtemplate.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := etuo.mutation.AddedCreatedAt(); ok {
		_spec.AddField(emailtemplate.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := etuo.mutation.UpdatedAt(); ok {
		_spec.SetField(emailtemplate.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := etuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(emailtemplate.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := etuo.mutation.DeletedAt(); ok {
		_spec.SetField(emailtemplate.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := etuo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(emailtemplate.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := etuo.mutation.EntID(); ok {
		_spec.SetField(emailtemplate.FieldEntID, field.TypeUUID, value)
	}
	if value, ok := etuo.mutation.AppID(); ok {
		_spec.SetField(emailtemplate.FieldAppID, field.TypeUUID, value)
	}
	if value, ok := etuo.mutation.LangID(); ok {
		_spec.SetField(emailtemplate.FieldLangID, field.TypeUUID, value)
	}
	if value, ok := etuo.mutation.DefaultToUsername(); ok {
		_spec.SetField(emailtemplate.FieldDefaultToUsername, field.TypeString, value)
	}
	if value, ok := etuo.mutation.UsedFor(); ok {
		_spec.SetField(emailtemplate.FieldUsedFor, field.TypeString, value)
	}
	if etuo.mutation.UsedForCleared() {
		_spec.ClearField(emailtemplate.FieldUsedFor, field.TypeString)
	}
	if value, ok := etuo.mutation.Sender(); ok {
		_spec.SetField(emailtemplate.FieldSender, field.TypeString, value)
	}
	if etuo.mutation.SenderCleared() {
		_spec.ClearField(emailtemplate.FieldSender, field.TypeString)
	}
	if value, ok := etuo.mutation.ReplyTos(); ok {
		_spec.SetField(emailtemplate.FieldReplyTos, field.TypeJSON, value)
	}
	if value, ok := etuo.mutation.AppendedReplyTos(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, emailtemplate.FieldReplyTos, value)
		})
	}
	if etuo.mutation.ReplyTosCleared() {
		_spec.ClearField(emailtemplate.FieldReplyTos, field.TypeJSON)
	}
	if value, ok := etuo.mutation.CcTos(); ok {
		_spec.SetField(emailtemplate.FieldCcTos, field.TypeJSON, value)
	}
	if value, ok := etuo.mutation.AppendedCcTos(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, emailtemplate.FieldCcTos, value)
		})
	}
	if etuo.mutation.CcTosCleared() {
		_spec.ClearField(emailtemplate.FieldCcTos, field.TypeJSON)
	}
	if value, ok := etuo.mutation.Subject(); ok {
		_spec.SetField(emailtemplate.FieldSubject, field.TypeString, value)
	}
	if etuo.mutation.SubjectCleared() {
		_spec.ClearField(emailtemplate.FieldSubject, field.TypeString)
	}
	if value, ok := etuo.mutation.Body(); ok {
		_spec.SetField(emailtemplate.FieldBody, field.TypeString, value)
	}
	if etuo.mutation.BodyCleared() {
		_spec.ClearField(emailtemplate.FieldBody, field.TypeString)
	}
	_spec.AddModifiers(etuo.modifiers...)
	_node = &EmailTemplate{config: etuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, etuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{emailtemplate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	etuo.mutation.done = true
	return _node, nil
}
