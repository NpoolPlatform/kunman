// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/comment"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/predicate"
	"github.com/google/uuid"
)

// CommentUpdate is the builder for updating Comment entities.
type CommentUpdate struct {
	config
	hooks     []Hook
	mutation  *CommentMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CommentUpdate builder.
func (cu *CommentUpdate) Where(ps ...predicate.Comment) *CommentUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetEntID sets the "ent_id" field.
func (cu *CommentUpdate) SetEntID(u uuid.UUID) *CommentUpdate {
	cu.mutation.SetEntID(u)
	return cu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableEntID(u *uuid.UUID) *CommentUpdate {
	if u != nil {
		cu.SetEntID(*u)
	}
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *CommentUpdate) SetCreatedAt(u uint32) *CommentUpdate {
	cu.mutation.ResetCreatedAt()
	cu.mutation.SetCreatedAt(u)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableCreatedAt(u *uint32) *CommentUpdate {
	if u != nil {
		cu.SetCreatedAt(*u)
	}
	return cu
}

// AddCreatedAt adds u to the "created_at" field.
func (cu *CommentUpdate) AddCreatedAt(u int32) *CommentUpdate {
	cu.mutation.AddCreatedAt(u)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CommentUpdate) SetUpdatedAt(u uint32) *CommentUpdate {
	cu.mutation.ResetUpdatedAt()
	cu.mutation.SetUpdatedAt(u)
	return cu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cu *CommentUpdate) AddUpdatedAt(u int32) *CommentUpdate {
	cu.mutation.AddUpdatedAt(u)
	return cu
}

// SetDeletedAt sets the "deleted_at" field.
func (cu *CommentUpdate) SetDeletedAt(u uint32) *CommentUpdate {
	cu.mutation.ResetDeletedAt()
	cu.mutation.SetDeletedAt(u)
	return cu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableDeletedAt(u *uint32) *CommentUpdate {
	if u != nil {
		cu.SetDeletedAt(*u)
	}
	return cu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cu *CommentUpdate) AddDeletedAt(u int32) *CommentUpdate {
	cu.mutation.AddDeletedAt(u)
	return cu
}

// SetUserID sets the "user_id" field.
func (cu *CommentUpdate) SetUserID(u uuid.UUID) *CommentUpdate {
	cu.mutation.SetUserID(u)
	return cu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableUserID(u *uuid.UUID) *CommentUpdate {
	if u != nil {
		cu.SetUserID(*u)
	}
	return cu
}

// ClearUserID clears the value of the "user_id" field.
func (cu *CommentUpdate) ClearUserID() *CommentUpdate {
	cu.mutation.ClearUserID()
	return cu
}

// SetAppGoodID sets the "app_good_id" field.
func (cu *CommentUpdate) SetAppGoodID(u uuid.UUID) *CommentUpdate {
	cu.mutation.SetAppGoodID(u)
	return cu
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableAppGoodID(u *uuid.UUID) *CommentUpdate {
	if u != nil {
		cu.SetAppGoodID(*u)
	}
	return cu
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (cu *CommentUpdate) ClearAppGoodID() *CommentUpdate {
	cu.mutation.ClearAppGoodID()
	return cu
}

// SetOrderID sets the "order_id" field.
func (cu *CommentUpdate) SetOrderID(u uuid.UUID) *CommentUpdate {
	cu.mutation.SetOrderID(u)
	return cu
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableOrderID(u *uuid.UUID) *CommentUpdate {
	if u != nil {
		cu.SetOrderID(*u)
	}
	return cu
}

// ClearOrderID clears the value of the "order_id" field.
func (cu *CommentUpdate) ClearOrderID() *CommentUpdate {
	cu.mutation.ClearOrderID()
	return cu
}

// SetContent sets the "content" field.
func (cu *CommentUpdate) SetContent(s string) *CommentUpdate {
	cu.mutation.SetContent(s)
	return cu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableContent(s *string) *CommentUpdate {
	if s != nil {
		cu.SetContent(*s)
	}
	return cu
}

// ClearContent clears the value of the "content" field.
func (cu *CommentUpdate) ClearContent() *CommentUpdate {
	cu.mutation.ClearContent()
	return cu
}

// SetReplyToID sets the "reply_to_id" field.
func (cu *CommentUpdate) SetReplyToID(u uuid.UUID) *CommentUpdate {
	cu.mutation.SetReplyToID(u)
	return cu
}

// SetNillableReplyToID sets the "reply_to_id" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableReplyToID(u *uuid.UUID) *CommentUpdate {
	if u != nil {
		cu.SetReplyToID(*u)
	}
	return cu
}

// ClearReplyToID clears the value of the "reply_to_id" field.
func (cu *CommentUpdate) ClearReplyToID() *CommentUpdate {
	cu.mutation.ClearReplyToID()
	return cu
}

// SetAnonymous sets the "anonymous" field.
func (cu *CommentUpdate) SetAnonymous(b bool) *CommentUpdate {
	cu.mutation.SetAnonymous(b)
	return cu
}

// SetNillableAnonymous sets the "anonymous" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableAnonymous(b *bool) *CommentUpdate {
	if b != nil {
		cu.SetAnonymous(*b)
	}
	return cu
}

// ClearAnonymous clears the value of the "anonymous" field.
func (cu *CommentUpdate) ClearAnonymous() *CommentUpdate {
	cu.mutation.ClearAnonymous()
	return cu
}

// SetTrialUser sets the "trial_user" field.
func (cu *CommentUpdate) SetTrialUser(b bool) *CommentUpdate {
	cu.mutation.SetTrialUser(b)
	return cu
}

// SetNillableTrialUser sets the "trial_user" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableTrialUser(b *bool) *CommentUpdate {
	if b != nil {
		cu.SetTrialUser(*b)
	}
	return cu
}

// ClearTrialUser clears the value of the "trial_user" field.
func (cu *CommentUpdate) ClearTrialUser() *CommentUpdate {
	cu.mutation.ClearTrialUser()
	return cu
}

// SetPurchasedUser sets the "purchased_user" field.
func (cu *CommentUpdate) SetPurchasedUser(b bool) *CommentUpdate {
	cu.mutation.SetPurchasedUser(b)
	return cu
}

// SetNillablePurchasedUser sets the "purchased_user" field if the given value is not nil.
func (cu *CommentUpdate) SetNillablePurchasedUser(b *bool) *CommentUpdate {
	if b != nil {
		cu.SetPurchasedUser(*b)
	}
	return cu
}

// ClearPurchasedUser clears the value of the "purchased_user" field.
func (cu *CommentUpdate) ClearPurchasedUser() *CommentUpdate {
	cu.mutation.ClearPurchasedUser()
	return cu
}

// SetHide sets the "hide" field.
func (cu *CommentUpdate) SetHide(b bool) *CommentUpdate {
	cu.mutation.SetHide(b)
	return cu
}

// SetNillableHide sets the "hide" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableHide(b *bool) *CommentUpdate {
	if b != nil {
		cu.SetHide(*b)
	}
	return cu
}

// ClearHide clears the value of the "hide" field.
func (cu *CommentUpdate) ClearHide() *CommentUpdate {
	cu.mutation.ClearHide()
	return cu
}

// SetHideReason sets the "hide_reason" field.
func (cu *CommentUpdate) SetHideReason(s string) *CommentUpdate {
	cu.mutation.SetHideReason(s)
	return cu
}

// SetNillableHideReason sets the "hide_reason" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableHideReason(s *string) *CommentUpdate {
	if s != nil {
		cu.SetHideReason(*s)
	}
	return cu
}

// ClearHideReason clears the value of the "hide_reason" field.
func (cu *CommentUpdate) ClearHideReason() *CommentUpdate {
	cu.mutation.ClearHideReason()
	return cu
}

// Mutation returns the CommentMutation object of the builder.
func (cu *CommentUpdate) Mutation() *CommentMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CommentUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CommentUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CommentUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CommentUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CommentUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := comment.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cu *CommentUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CommentUpdate {
	cu.modifiers = append(cu.modifiers, modifiers...)
	return cu
}

func (cu *CommentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(comment.Table, comment.Columns, sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUint32))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.EntID(); ok {
		_spec.SetField(comment.FieldEntID, field.TypeUUID, value)
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.SetField(comment.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := cu.mutation.AddedCreatedAt(); ok {
		_spec.AddField(comment.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(comment.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := cu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(comment.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := cu.mutation.DeletedAt(); ok {
		_spec.SetField(comment.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := cu.mutation.AddedDeletedAt(); ok {
		_spec.AddField(comment.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := cu.mutation.UserID(); ok {
		_spec.SetField(comment.FieldUserID, field.TypeUUID, value)
	}
	if cu.mutation.UserIDCleared() {
		_spec.ClearField(comment.FieldUserID, field.TypeUUID)
	}
	if value, ok := cu.mutation.AppGoodID(); ok {
		_spec.SetField(comment.FieldAppGoodID, field.TypeUUID, value)
	}
	if cu.mutation.AppGoodIDCleared() {
		_spec.ClearField(comment.FieldAppGoodID, field.TypeUUID)
	}
	if value, ok := cu.mutation.OrderID(); ok {
		_spec.SetField(comment.FieldOrderID, field.TypeUUID, value)
	}
	if cu.mutation.OrderIDCleared() {
		_spec.ClearField(comment.FieldOrderID, field.TypeUUID)
	}
	if value, ok := cu.mutation.Content(); ok {
		_spec.SetField(comment.FieldContent, field.TypeString, value)
	}
	if cu.mutation.ContentCleared() {
		_spec.ClearField(comment.FieldContent, field.TypeString)
	}
	if value, ok := cu.mutation.ReplyToID(); ok {
		_spec.SetField(comment.FieldReplyToID, field.TypeUUID, value)
	}
	if cu.mutation.ReplyToIDCleared() {
		_spec.ClearField(comment.FieldReplyToID, field.TypeUUID)
	}
	if value, ok := cu.mutation.Anonymous(); ok {
		_spec.SetField(comment.FieldAnonymous, field.TypeBool, value)
	}
	if cu.mutation.AnonymousCleared() {
		_spec.ClearField(comment.FieldAnonymous, field.TypeBool)
	}
	if value, ok := cu.mutation.TrialUser(); ok {
		_spec.SetField(comment.FieldTrialUser, field.TypeBool, value)
	}
	if cu.mutation.TrialUserCleared() {
		_spec.ClearField(comment.FieldTrialUser, field.TypeBool)
	}
	if value, ok := cu.mutation.PurchasedUser(); ok {
		_spec.SetField(comment.FieldPurchasedUser, field.TypeBool, value)
	}
	if cu.mutation.PurchasedUserCleared() {
		_spec.ClearField(comment.FieldPurchasedUser, field.TypeBool)
	}
	if value, ok := cu.mutation.Hide(); ok {
		_spec.SetField(comment.FieldHide, field.TypeBool, value)
	}
	if cu.mutation.HideCleared() {
		_spec.ClearField(comment.FieldHide, field.TypeBool)
	}
	if value, ok := cu.mutation.HideReason(); ok {
		_spec.SetField(comment.FieldHideReason, field.TypeString, value)
	}
	if cu.mutation.HideReasonCleared() {
		_spec.ClearField(comment.FieldHideReason, field.TypeString)
	}
	_spec.AddModifiers(cu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CommentUpdateOne is the builder for updating a single Comment entity.
type CommentUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CommentMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetEntID sets the "ent_id" field.
func (cuo *CommentUpdateOne) SetEntID(u uuid.UUID) *CommentUpdateOne {
	cuo.mutation.SetEntID(u)
	return cuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableEntID(u *uuid.UUID) *CommentUpdateOne {
	if u != nil {
		cuo.SetEntID(*u)
	}
	return cuo
}

// SetCreatedAt sets the "created_at" field.
func (cuo *CommentUpdateOne) SetCreatedAt(u uint32) *CommentUpdateOne {
	cuo.mutation.ResetCreatedAt()
	cuo.mutation.SetCreatedAt(u)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableCreatedAt(u *uint32) *CommentUpdateOne {
	if u != nil {
		cuo.SetCreatedAt(*u)
	}
	return cuo
}

// AddCreatedAt adds u to the "created_at" field.
func (cuo *CommentUpdateOne) AddCreatedAt(u int32) *CommentUpdateOne {
	cuo.mutation.AddCreatedAt(u)
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CommentUpdateOne) SetUpdatedAt(u uint32) *CommentUpdateOne {
	cuo.mutation.ResetUpdatedAt()
	cuo.mutation.SetUpdatedAt(u)
	return cuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cuo *CommentUpdateOne) AddUpdatedAt(u int32) *CommentUpdateOne {
	cuo.mutation.AddUpdatedAt(u)
	return cuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cuo *CommentUpdateOne) SetDeletedAt(u uint32) *CommentUpdateOne {
	cuo.mutation.ResetDeletedAt()
	cuo.mutation.SetDeletedAt(u)
	return cuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableDeletedAt(u *uint32) *CommentUpdateOne {
	if u != nil {
		cuo.SetDeletedAt(*u)
	}
	return cuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cuo *CommentUpdateOne) AddDeletedAt(u int32) *CommentUpdateOne {
	cuo.mutation.AddDeletedAt(u)
	return cuo
}

// SetUserID sets the "user_id" field.
func (cuo *CommentUpdateOne) SetUserID(u uuid.UUID) *CommentUpdateOne {
	cuo.mutation.SetUserID(u)
	return cuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableUserID(u *uuid.UUID) *CommentUpdateOne {
	if u != nil {
		cuo.SetUserID(*u)
	}
	return cuo
}

// ClearUserID clears the value of the "user_id" field.
func (cuo *CommentUpdateOne) ClearUserID() *CommentUpdateOne {
	cuo.mutation.ClearUserID()
	return cuo
}

// SetAppGoodID sets the "app_good_id" field.
func (cuo *CommentUpdateOne) SetAppGoodID(u uuid.UUID) *CommentUpdateOne {
	cuo.mutation.SetAppGoodID(u)
	return cuo
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableAppGoodID(u *uuid.UUID) *CommentUpdateOne {
	if u != nil {
		cuo.SetAppGoodID(*u)
	}
	return cuo
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (cuo *CommentUpdateOne) ClearAppGoodID() *CommentUpdateOne {
	cuo.mutation.ClearAppGoodID()
	return cuo
}

// SetOrderID sets the "order_id" field.
func (cuo *CommentUpdateOne) SetOrderID(u uuid.UUID) *CommentUpdateOne {
	cuo.mutation.SetOrderID(u)
	return cuo
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableOrderID(u *uuid.UUID) *CommentUpdateOne {
	if u != nil {
		cuo.SetOrderID(*u)
	}
	return cuo
}

// ClearOrderID clears the value of the "order_id" field.
func (cuo *CommentUpdateOne) ClearOrderID() *CommentUpdateOne {
	cuo.mutation.ClearOrderID()
	return cuo
}

// SetContent sets the "content" field.
func (cuo *CommentUpdateOne) SetContent(s string) *CommentUpdateOne {
	cuo.mutation.SetContent(s)
	return cuo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableContent(s *string) *CommentUpdateOne {
	if s != nil {
		cuo.SetContent(*s)
	}
	return cuo
}

// ClearContent clears the value of the "content" field.
func (cuo *CommentUpdateOne) ClearContent() *CommentUpdateOne {
	cuo.mutation.ClearContent()
	return cuo
}

// SetReplyToID sets the "reply_to_id" field.
func (cuo *CommentUpdateOne) SetReplyToID(u uuid.UUID) *CommentUpdateOne {
	cuo.mutation.SetReplyToID(u)
	return cuo
}

// SetNillableReplyToID sets the "reply_to_id" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableReplyToID(u *uuid.UUID) *CommentUpdateOne {
	if u != nil {
		cuo.SetReplyToID(*u)
	}
	return cuo
}

// ClearReplyToID clears the value of the "reply_to_id" field.
func (cuo *CommentUpdateOne) ClearReplyToID() *CommentUpdateOne {
	cuo.mutation.ClearReplyToID()
	return cuo
}

// SetAnonymous sets the "anonymous" field.
func (cuo *CommentUpdateOne) SetAnonymous(b bool) *CommentUpdateOne {
	cuo.mutation.SetAnonymous(b)
	return cuo
}

// SetNillableAnonymous sets the "anonymous" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableAnonymous(b *bool) *CommentUpdateOne {
	if b != nil {
		cuo.SetAnonymous(*b)
	}
	return cuo
}

// ClearAnonymous clears the value of the "anonymous" field.
func (cuo *CommentUpdateOne) ClearAnonymous() *CommentUpdateOne {
	cuo.mutation.ClearAnonymous()
	return cuo
}

// SetTrialUser sets the "trial_user" field.
func (cuo *CommentUpdateOne) SetTrialUser(b bool) *CommentUpdateOne {
	cuo.mutation.SetTrialUser(b)
	return cuo
}

// SetNillableTrialUser sets the "trial_user" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableTrialUser(b *bool) *CommentUpdateOne {
	if b != nil {
		cuo.SetTrialUser(*b)
	}
	return cuo
}

// ClearTrialUser clears the value of the "trial_user" field.
func (cuo *CommentUpdateOne) ClearTrialUser() *CommentUpdateOne {
	cuo.mutation.ClearTrialUser()
	return cuo
}

// SetPurchasedUser sets the "purchased_user" field.
func (cuo *CommentUpdateOne) SetPurchasedUser(b bool) *CommentUpdateOne {
	cuo.mutation.SetPurchasedUser(b)
	return cuo
}

// SetNillablePurchasedUser sets the "purchased_user" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillablePurchasedUser(b *bool) *CommentUpdateOne {
	if b != nil {
		cuo.SetPurchasedUser(*b)
	}
	return cuo
}

// ClearPurchasedUser clears the value of the "purchased_user" field.
func (cuo *CommentUpdateOne) ClearPurchasedUser() *CommentUpdateOne {
	cuo.mutation.ClearPurchasedUser()
	return cuo
}

// SetHide sets the "hide" field.
func (cuo *CommentUpdateOne) SetHide(b bool) *CommentUpdateOne {
	cuo.mutation.SetHide(b)
	return cuo
}

// SetNillableHide sets the "hide" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableHide(b *bool) *CommentUpdateOne {
	if b != nil {
		cuo.SetHide(*b)
	}
	return cuo
}

// ClearHide clears the value of the "hide" field.
func (cuo *CommentUpdateOne) ClearHide() *CommentUpdateOne {
	cuo.mutation.ClearHide()
	return cuo
}

// SetHideReason sets the "hide_reason" field.
func (cuo *CommentUpdateOne) SetHideReason(s string) *CommentUpdateOne {
	cuo.mutation.SetHideReason(s)
	return cuo
}

// SetNillableHideReason sets the "hide_reason" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableHideReason(s *string) *CommentUpdateOne {
	if s != nil {
		cuo.SetHideReason(*s)
	}
	return cuo
}

// ClearHideReason clears the value of the "hide_reason" field.
func (cuo *CommentUpdateOne) ClearHideReason() *CommentUpdateOne {
	cuo.mutation.ClearHideReason()
	return cuo
}

// Mutation returns the CommentMutation object of the builder.
func (cuo *CommentUpdateOne) Mutation() *CommentMutation {
	return cuo.mutation
}

// Where appends a list predicates to the CommentUpdate builder.
func (cuo *CommentUpdateOne) Where(ps ...predicate.Comment) *CommentUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CommentUpdateOne) Select(field string, fields ...string) *CommentUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Comment entity.
func (cuo *CommentUpdateOne) Save(ctx context.Context) (*Comment, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CommentUpdateOne) SaveX(ctx context.Context) *Comment {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CommentUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CommentUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CommentUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := comment.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuo *CommentUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CommentUpdateOne {
	cuo.modifiers = append(cuo.modifiers, modifiers...)
	return cuo
}

func (cuo *CommentUpdateOne) sqlSave(ctx context.Context) (_node *Comment, err error) {
	_spec := sqlgraph.NewUpdateSpec(comment.Table, comment.Columns, sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUint32))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Comment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, comment.FieldID)
		for _, f := range fields {
			if !comment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != comment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.EntID(); ok {
		_spec.SetField(comment.FieldEntID, field.TypeUUID, value)
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.SetField(comment.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := cuo.mutation.AddedCreatedAt(); ok {
		_spec.AddField(comment.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(comment.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := cuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(comment.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := cuo.mutation.DeletedAt(); ok {
		_spec.SetField(comment.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := cuo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(comment.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := cuo.mutation.UserID(); ok {
		_spec.SetField(comment.FieldUserID, field.TypeUUID, value)
	}
	if cuo.mutation.UserIDCleared() {
		_spec.ClearField(comment.FieldUserID, field.TypeUUID)
	}
	if value, ok := cuo.mutation.AppGoodID(); ok {
		_spec.SetField(comment.FieldAppGoodID, field.TypeUUID, value)
	}
	if cuo.mutation.AppGoodIDCleared() {
		_spec.ClearField(comment.FieldAppGoodID, field.TypeUUID)
	}
	if value, ok := cuo.mutation.OrderID(); ok {
		_spec.SetField(comment.FieldOrderID, field.TypeUUID, value)
	}
	if cuo.mutation.OrderIDCleared() {
		_spec.ClearField(comment.FieldOrderID, field.TypeUUID)
	}
	if value, ok := cuo.mutation.Content(); ok {
		_spec.SetField(comment.FieldContent, field.TypeString, value)
	}
	if cuo.mutation.ContentCleared() {
		_spec.ClearField(comment.FieldContent, field.TypeString)
	}
	if value, ok := cuo.mutation.ReplyToID(); ok {
		_spec.SetField(comment.FieldReplyToID, field.TypeUUID, value)
	}
	if cuo.mutation.ReplyToIDCleared() {
		_spec.ClearField(comment.FieldReplyToID, field.TypeUUID)
	}
	if value, ok := cuo.mutation.Anonymous(); ok {
		_spec.SetField(comment.FieldAnonymous, field.TypeBool, value)
	}
	if cuo.mutation.AnonymousCleared() {
		_spec.ClearField(comment.FieldAnonymous, field.TypeBool)
	}
	if value, ok := cuo.mutation.TrialUser(); ok {
		_spec.SetField(comment.FieldTrialUser, field.TypeBool, value)
	}
	if cuo.mutation.TrialUserCleared() {
		_spec.ClearField(comment.FieldTrialUser, field.TypeBool)
	}
	if value, ok := cuo.mutation.PurchasedUser(); ok {
		_spec.SetField(comment.FieldPurchasedUser, field.TypeBool, value)
	}
	if cuo.mutation.PurchasedUserCleared() {
		_spec.ClearField(comment.FieldPurchasedUser, field.TypeBool)
	}
	if value, ok := cuo.mutation.Hide(); ok {
		_spec.SetField(comment.FieldHide, field.TypeBool, value)
	}
	if cuo.mutation.HideCleared() {
		_spec.ClearField(comment.FieldHide, field.TypeBool)
	}
	if value, ok := cuo.mutation.HideReason(); ok {
		_spec.SetField(comment.FieldHideReason, field.TypeString, value)
	}
	if cuo.mutation.HideReasonCleared() {
		_spec.ClearField(comment.FieldHideReason, field.TypeString)
	}
	_spec.AddModifiers(cuo.modifiers...)
	_node = &Comment{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
