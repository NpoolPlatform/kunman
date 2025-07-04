// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/extrainfo"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ExtraInfoUpdate is the builder for updating ExtraInfo entities.
type ExtraInfoUpdate struct {
	config
	hooks     []Hook
	mutation  *ExtraInfoMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ExtraInfoUpdate builder.
func (eiu *ExtraInfoUpdate) Where(ps ...predicate.ExtraInfo) *ExtraInfoUpdate {
	eiu.mutation.Where(ps...)
	return eiu
}

// SetEntID sets the "ent_id" field.
func (eiu *ExtraInfoUpdate) SetEntID(u uuid.UUID) *ExtraInfoUpdate {
	eiu.mutation.SetEntID(u)
	return eiu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (eiu *ExtraInfoUpdate) SetNillableEntID(u *uuid.UUID) *ExtraInfoUpdate {
	if u != nil {
		eiu.SetEntID(*u)
	}
	return eiu
}

// SetCreatedAt sets the "created_at" field.
func (eiu *ExtraInfoUpdate) SetCreatedAt(u uint32) *ExtraInfoUpdate {
	eiu.mutation.ResetCreatedAt()
	eiu.mutation.SetCreatedAt(u)
	return eiu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (eiu *ExtraInfoUpdate) SetNillableCreatedAt(u *uint32) *ExtraInfoUpdate {
	if u != nil {
		eiu.SetCreatedAt(*u)
	}
	return eiu
}

// AddCreatedAt adds u to the "created_at" field.
func (eiu *ExtraInfoUpdate) AddCreatedAt(u int32) *ExtraInfoUpdate {
	eiu.mutation.AddCreatedAt(u)
	return eiu
}

// SetUpdatedAt sets the "updated_at" field.
func (eiu *ExtraInfoUpdate) SetUpdatedAt(u uint32) *ExtraInfoUpdate {
	eiu.mutation.ResetUpdatedAt()
	eiu.mutation.SetUpdatedAt(u)
	return eiu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (eiu *ExtraInfoUpdate) AddUpdatedAt(u int32) *ExtraInfoUpdate {
	eiu.mutation.AddUpdatedAt(u)
	return eiu
}

// SetDeletedAt sets the "deleted_at" field.
func (eiu *ExtraInfoUpdate) SetDeletedAt(u uint32) *ExtraInfoUpdate {
	eiu.mutation.ResetDeletedAt()
	eiu.mutation.SetDeletedAt(u)
	return eiu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (eiu *ExtraInfoUpdate) SetNillableDeletedAt(u *uint32) *ExtraInfoUpdate {
	if u != nil {
		eiu.SetDeletedAt(*u)
	}
	return eiu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (eiu *ExtraInfoUpdate) AddDeletedAt(u int32) *ExtraInfoUpdate {
	eiu.mutation.AddDeletedAt(u)
	return eiu
}

// SetAppGoodID sets the "app_good_id" field.
func (eiu *ExtraInfoUpdate) SetAppGoodID(u uuid.UUID) *ExtraInfoUpdate {
	eiu.mutation.SetAppGoodID(u)
	return eiu
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (eiu *ExtraInfoUpdate) SetNillableAppGoodID(u *uuid.UUID) *ExtraInfoUpdate {
	if u != nil {
		eiu.SetAppGoodID(*u)
	}
	return eiu
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (eiu *ExtraInfoUpdate) ClearAppGoodID() *ExtraInfoUpdate {
	eiu.mutation.ClearAppGoodID()
	return eiu
}

// SetLikes sets the "likes" field.
func (eiu *ExtraInfoUpdate) SetLikes(u uint32) *ExtraInfoUpdate {
	eiu.mutation.ResetLikes()
	eiu.mutation.SetLikes(u)
	return eiu
}

// SetNillableLikes sets the "likes" field if the given value is not nil.
func (eiu *ExtraInfoUpdate) SetNillableLikes(u *uint32) *ExtraInfoUpdate {
	if u != nil {
		eiu.SetLikes(*u)
	}
	return eiu
}

// AddLikes adds u to the "likes" field.
func (eiu *ExtraInfoUpdate) AddLikes(u int32) *ExtraInfoUpdate {
	eiu.mutation.AddLikes(u)
	return eiu
}

// ClearLikes clears the value of the "likes" field.
func (eiu *ExtraInfoUpdate) ClearLikes() *ExtraInfoUpdate {
	eiu.mutation.ClearLikes()
	return eiu
}

// SetDislikes sets the "dislikes" field.
func (eiu *ExtraInfoUpdate) SetDislikes(u uint32) *ExtraInfoUpdate {
	eiu.mutation.ResetDislikes()
	eiu.mutation.SetDislikes(u)
	return eiu
}

// SetNillableDislikes sets the "dislikes" field if the given value is not nil.
func (eiu *ExtraInfoUpdate) SetNillableDislikes(u *uint32) *ExtraInfoUpdate {
	if u != nil {
		eiu.SetDislikes(*u)
	}
	return eiu
}

// AddDislikes adds u to the "dislikes" field.
func (eiu *ExtraInfoUpdate) AddDislikes(u int32) *ExtraInfoUpdate {
	eiu.mutation.AddDislikes(u)
	return eiu
}

// ClearDislikes clears the value of the "dislikes" field.
func (eiu *ExtraInfoUpdate) ClearDislikes() *ExtraInfoUpdate {
	eiu.mutation.ClearDislikes()
	return eiu
}

// SetRecommendCount sets the "recommend_count" field.
func (eiu *ExtraInfoUpdate) SetRecommendCount(u uint32) *ExtraInfoUpdate {
	eiu.mutation.ResetRecommendCount()
	eiu.mutation.SetRecommendCount(u)
	return eiu
}

// SetNillableRecommendCount sets the "recommend_count" field if the given value is not nil.
func (eiu *ExtraInfoUpdate) SetNillableRecommendCount(u *uint32) *ExtraInfoUpdate {
	if u != nil {
		eiu.SetRecommendCount(*u)
	}
	return eiu
}

// AddRecommendCount adds u to the "recommend_count" field.
func (eiu *ExtraInfoUpdate) AddRecommendCount(u int32) *ExtraInfoUpdate {
	eiu.mutation.AddRecommendCount(u)
	return eiu
}

// ClearRecommendCount clears the value of the "recommend_count" field.
func (eiu *ExtraInfoUpdate) ClearRecommendCount() *ExtraInfoUpdate {
	eiu.mutation.ClearRecommendCount()
	return eiu
}

// SetCommentCount sets the "comment_count" field.
func (eiu *ExtraInfoUpdate) SetCommentCount(u uint32) *ExtraInfoUpdate {
	eiu.mutation.ResetCommentCount()
	eiu.mutation.SetCommentCount(u)
	return eiu
}

// SetNillableCommentCount sets the "comment_count" field if the given value is not nil.
func (eiu *ExtraInfoUpdate) SetNillableCommentCount(u *uint32) *ExtraInfoUpdate {
	if u != nil {
		eiu.SetCommentCount(*u)
	}
	return eiu
}

// AddCommentCount adds u to the "comment_count" field.
func (eiu *ExtraInfoUpdate) AddCommentCount(u int32) *ExtraInfoUpdate {
	eiu.mutation.AddCommentCount(u)
	return eiu
}

// ClearCommentCount clears the value of the "comment_count" field.
func (eiu *ExtraInfoUpdate) ClearCommentCount() *ExtraInfoUpdate {
	eiu.mutation.ClearCommentCount()
	return eiu
}

// SetScoreCount sets the "score_count" field.
func (eiu *ExtraInfoUpdate) SetScoreCount(u uint32) *ExtraInfoUpdate {
	eiu.mutation.ResetScoreCount()
	eiu.mutation.SetScoreCount(u)
	return eiu
}

// SetNillableScoreCount sets the "score_count" field if the given value is not nil.
func (eiu *ExtraInfoUpdate) SetNillableScoreCount(u *uint32) *ExtraInfoUpdate {
	if u != nil {
		eiu.SetScoreCount(*u)
	}
	return eiu
}

// AddScoreCount adds u to the "score_count" field.
func (eiu *ExtraInfoUpdate) AddScoreCount(u int32) *ExtraInfoUpdate {
	eiu.mutation.AddScoreCount(u)
	return eiu
}

// ClearScoreCount clears the value of the "score_count" field.
func (eiu *ExtraInfoUpdate) ClearScoreCount() *ExtraInfoUpdate {
	eiu.mutation.ClearScoreCount()
	return eiu
}

// SetScore sets the "score" field.
func (eiu *ExtraInfoUpdate) SetScore(d decimal.Decimal) *ExtraInfoUpdate {
	eiu.mutation.SetScore(d)
	return eiu
}

// SetNillableScore sets the "score" field if the given value is not nil.
func (eiu *ExtraInfoUpdate) SetNillableScore(d *decimal.Decimal) *ExtraInfoUpdate {
	if d != nil {
		eiu.SetScore(*d)
	}
	return eiu
}

// ClearScore clears the value of the "score" field.
func (eiu *ExtraInfoUpdate) ClearScore() *ExtraInfoUpdate {
	eiu.mutation.ClearScore()
	return eiu
}

// Mutation returns the ExtraInfoMutation object of the builder.
func (eiu *ExtraInfoUpdate) Mutation() *ExtraInfoMutation {
	return eiu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eiu *ExtraInfoUpdate) Save(ctx context.Context) (int, error) {
	eiu.defaults()
	return withHooks(ctx, eiu.sqlSave, eiu.mutation, eiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eiu *ExtraInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := eiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eiu *ExtraInfoUpdate) Exec(ctx context.Context) error {
	_, err := eiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eiu *ExtraInfoUpdate) ExecX(ctx context.Context) {
	if err := eiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eiu *ExtraInfoUpdate) defaults() {
	if _, ok := eiu.mutation.UpdatedAt(); !ok {
		v := extrainfo.UpdateDefaultUpdatedAt()
		eiu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (eiu *ExtraInfoUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ExtraInfoUpdate {
	eiu.modifiers = append(eiu.modifiers, modifiers...)
	return eiu
}

func (eiu *ExtraInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(extrainfo.Table, extrainfo.Columns, sqlgraph.NewFieldSpec(extrainfo.FieldID, field.TypeUint32))
	if ps := eiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eiu.mutation.EntID(); ok {
		_spec.SetField(extrainfo.FieldEntID, field.TypeUUID, value)
	}
	if value, ok := eiu.mutation.CreatedAt(); ok {
		_spec.SetField(extrainfo.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := eiu.mutation.AddedCreatedAt(); ok {
		_spec.AddField(extrainfo.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := eiu.mutation.UpdatedAt(); ok {
		_spec.SetField(extrainfo.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := eiu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(extrainfo.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := eiu.mutation.DeletedAt(); ok {
		_spec.SetField(extrainfo.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := eiu.mutation.AddedDeletedAt(); ok {
		_spec.AddField(extrainfo.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := eiu.mutation.AppGoodID(); ok {
		_spec.SetField(extrainfo.FieldAppGoodID, field.TypeUUID, value)
	}
	if eiu.mutation.AppGoodIDCleared() {
		_spec.ClearField(extrainfo.FieldAppGoodID, field.TypeUUID)
	}
	if value, ok := eiu.mutation.Likes(); ok {
		_spec.SetField(extrainfo.FieldLikes, field.TypeUint32, value)
	}
	if value, ok := eiu.mutation.AddedLikes(); ok {
		_spec.AddField(extrainfo.FieldLikes, field.TypeUint32, value)
	}
	if eiu.mutation.LikesCleared() {
		_spec.ClearField(extrainfo.FieldLikes, field.TypeUint32)
	}
	if value, ok := eiu.mutation.Dislikes(); ok {
		_spec.SetField(extrainfo.FieldDislikes, field.TypeUint32, value)
	}
	if value, ok := eiu.mutation.AddedDislikes(); ok {
		_spec.AddField(extrainfo.FieldDislikes, field.TypeUint32, value)
	}
	if eiu.mutation.DislikesCleared() {
		_spec.ClearField(extrainfo.FieldDislikes, field.TypeUint32)
	}
	if value, ok := eiu.mutation.RecommendCount(); ok {
		_spec.SetField(extrainfo.FieldRecommendCount, field.TypeUint32, value)
	}
	if value, ok := eiu.mutation.AddedRecommendCount(); ok {
		_spec.AddField(extrainfo.FieldRecommendCount, field.TypeUint32, value)
	}
	if eiu.mutation.RecommendCountCleared() {
		_spec.ClearField(extrainfo.FieldRecommendCount, field.TypeUint32)
	}
	if value, ok := eiu.mutation.CommentCount(); ok {
		_spec.SetField(extrainfo.FieldCommentCount, field.TypeUint32, value)
	}
	if value, ok := eiu.mutation.AddedCommentCount(); ok {
		_spec.AddField(extrainfo.FieldCommentCount, field.TypeUint32, value)
	}
	if eiu.mutation.CommentCountCleared() {
		_spec.ClearField(extrainfo.FieldCommentCount, field.TypeUint32)
	}
	if value, ok := eiu.mutation.ScoreCount(); ok {
		_spec.SetField(extrainfo.FieldScoreCount, field.TypeUint32, value)
	}
	if value, ok := eiu.mutation.AddedScoreCount(); ok {
		_spec.AddField(extrainfo.FieldScoreCount, field.TypeUint32, value)
	}
	if eiu.mutation.ScoreCountCleared() {
		_spec.ClearField(extrainfo.FieldScoreCount, field.TypeUint32)
	}
	if value, ok := eiu.mutation.Score(); ok {
		_spec.SetField(extrainfo.FieldScore, field.TypeOther, value)
	}
	if eiu.mutation.ScoreCleared() {
		_spec.ClearField(extrainfo.FieldScore, field.TypeOther)
	}
	_spec.AddModifiers(eiu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, eiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{extrainfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eiu.mutation.done = true
	return n, nil
}

// ExtraInfoUpdateOne is the builder for updating a single ExtraInfo entity.
type ExtraInfoUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ExtraInfoMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetEntID sets the "ent_id" field.
func (eiuo *ExtraInfoUpdateOne) SetEntID(u uuid.UUID) *ExtraInfoUpdateOne {
	eiuo.mutation.SetEntID(u)
	return eiuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (eiuo *ExtraInfoUpdateOne) SetNillableEntID(u *uuid.UUID) *ExtraInfoUpdateOne {
	if u != nil {
		eiuo.SetEntID(*u)
	}
	return eiuo
}

// SetCreatedAt sets the "created_at" field.
func (eiuo *ExtraInfoUpdateOne) SetCreatedAt(u uint32) *ExtraInfoUpdateOne {
	eiuo.mutation.ResetCreatedAt()
	eiuo.mutation.SetCreatedAt(u)
	return eiuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (eiuo *ExtraInfoUpdateOne) SetNillableCreatedAt(u *uint32) *ExtraInfoUpdateOne {
	if u != nil {
		eiuo.SetCreatedAt(*u)
	}
	return eiuo
}

// AddCreatedAt adds u to the "created_at" field.
func (eiuo *ExtraInfoUpdateOne) AddCreatedAt(u int32) *ExtraInfoUpdateOne {
	eiuo.mutation.AddCreatedAt(u)
	return eiuo
}

// SetUpdatedAt sets the "updated_at" field.
func (eiuo *ExtraInfoUpdateOne) SetUpdatedAt(u uint32) *ExtraInfoUpdateOne {
	eiuo.mutation.ResetUpdatedAt()
	eiuo.mutation.SetUpdatedAt(u)
	return eiuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (eiuo *ExtraInfoUpdateOne) AddUpdatedAt(u int32) *ExtraInfoUpdateOne {
	eiuo.mutation.AddUpdatedAt(u)
	return eiuo
}

// SetDeletedAt sets the "deleted_at" field.
func (eiuo *ExtraInfoUpdateOne) SetDeletedAt(u uint32) *ExtraInfoUpdateOne {
	eiuo.mutation.ResetDeletedAt()
	eiuo.mutation.SetDeletedAt(u)
	return eiuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (eiuo *ExtraInfoUpdateOne) SetNillableDeletedAt(u *uint32) *ExtraInfoUpdateOne {
	if u != nil {
		eiuo.SetDeletedAt(*u)
	}
	return eiuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (eiuo *ExtraInfoUpdateOne) AddDeletedAt(u int32) *ExtraInfoUpdateOne {
	eiuo.mutation.AddDeletedAt(u)
	return eiuo
}

// SetAppGoodID sets the "app_good_id" field.
func (eiuo *ExtraInfoUpdateOne) SetAppGoodID(u uuid.UUID) *ExtraInfoUpdateOne {
	eiuo.mutation.SetAppGoodID(u)
	return eiuo
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (eiuo *ExtraInfoUpdateOne) SetNillableAppGoodID(u *uuid.UUID) *ExtraInfoUpdateOne {
	if u != nil {
		eiuo.SetAppGoodID(*u)
	}
	return eiuo
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (eiuo *ExtraInfoUpdateOne) ClearAppGoodID() *ExtraInfoUpdateOne {
	eiuo.mutation.ClearAppGoodID()
	return eiuo
}

// SetLikes sets the "likes" field.
func (eiuo *ExtraInfoUpdateOne) SetLikes(u uint32) *ExtraInfoUpdateOne {
	eiuo.mutation.ResetLikes()
	eiuo.mutation.SetLikes(u)
	return eiuo
}

// SetNillableLikes sets the "likes" field if the given value is not nil.
func (eiuo *ExtraInfoUpdateOne) SetNillableLikes(u *uint32) *ExtraInfoUpdateOne {
	if u != nil {
		eiuo.SetLikes(*u)
	}
	return eiuo
}

// AddLikes adds u to the "likes" field.
func (eiuo *ExtraInfoUpdateOne) AddLikes(u int32) *ExtraInfoUpdateOne {
	eiuo.mutation.AddLikes(u)
	return eiuo
}

// ClearLikes clears the value of the "likes" field.
func (eiuo *ExtraInfoUpdateOne) ClearLikes() *ExtraInfoUpdateOne {
	eiuo.mutation.ClearLikes()
	return eiuo
}

// SetDislikes sets the "dislikes" field.
func (eiuo *ExtraInfoUpdateOne) SetDislikes(u uint32) *ExtraInfoUpdateOne {
	eiuo.mutation.ResetDislikes()
	eiuo.mutation.SetDislikes(u)
	return eiuo
}

// SetNillableDislikes sets the "dislikes" field if the given value is not nil.
func (eiuo *ExtraInfoUpdateOne) SetNillableDislikes(u *uint32) *ExtraInfoUpdateOne {
	if u != nil {
		eiuo.SetDislikes(*u)
	}
	return eiuo
}

// AddDislikes adds u to the "dislikes" field.
func (eiuo *ExtraInfoUpdateOne) AddDislikes(u int32) *ExtraInfoUpdateOne {
	eiuo.mutation.AddDislikes(u)
	return eiuo
}

// ClearDislikes clears the value of the "dislikes" field.
func (eiuo *ExtraInfoUpdateOne) ClearDislikes() *ExtraInfoUpdateOne {
	eiuo.mutation.ClearDislikes()
	return eiuo
}

// SetRecommendCount sets the "recommend_count" field.
func (eiuo *ExtraInfoUpdateOne) SetRecommendCount(u uint32) *ExtraInfoUpdateOne {
	eiuo.mutation.ResetRecommendCount()
	eiuo.mutation.SetRecommendCount(u)
	return eiuo
}

// SetNillableRecommendCount sets the "recommend_count" field if the given value is not nil.
func (eiuo *ExtraInfoUpdateOne) SetNillableRecommendCount(u *uint32) *ExtraInfoUpdateOne {
	if u != nil {
		eiuo.SetRecommendCount(*u)
	}
	return eiuo
}

// AddRecommendCount adds u to the "recommend_count" field.
func (eiuo *ExtraInfoUpdateOne) AddRecommendCount(u int32) *ExtraInfoUpdateOne {
	eiuo.mutation.AddRecommendCount(u)
	return eiuo
}

// ClearRecommendCount clears the value of the "recommend_count" field.
func (eiuo *ExtraInfoUpdateOne) ClearRecommendCount() *ExtraInfoUpdateOne {
	eiuo.mutation.ClearRecommendCount()
	return eiuo
}

// SetCommentCount sets the "comment_count" field.
func (eiuo *ExtraInfoUpdateOne) SetCommentCount(u uint32) *ExtraInfoUpdateOne {
	eiuo.mutation.ResetCommentCount()
	eiuo.mutation.SetCommentCount(u)
	return eiuo
}

// SetNillableCommentCount sets the "comment_count" field if the given value is not nil.
func (eiuo *ExtraInfoUpdateOne) SetNillableCommentCount(u *uint32) *ExtraInfoUpdateOne {
	if u != nil {
		eiuo.SetCommentCount(*u)
	}
	return eiuo
}

// AddCommentCount adds u to the "comment_count" field.
func (eiuo *ExtraInfoUpdateOne) AddCommentCount(u int32) *ExtraInfoUpdateOne {
	eiuo.mutation.AddCommentCount(u)
	return eiuo
}

// ClearCommentCount clears the value of the "comment_count" field.
func (eiuo *ExtraInfoUpdateOne) ClearCommentCount() *ExtraInfoUpdateOne {
	eiuo.mutation.ClearCommentCount()
	return eiuo
}

// SetScoreCount sets the "score_count" field.
func (eiuo *ExtraInfoUpdateOne) SetScoreCount(u uint32) *ExtraInfoUpdateOne {
	eiuo.mutation.ResetScoreCount()
	eiuo.mutation.SetScoreCount(u)
	return eiuo
}

// SetNillableScoreCount sets the "score_count" field if the given value is not nil.
func (eiuo *ExtraInfoUpdateOne) SetNillableScoreCount(u *uint32) *ExtraInfoUpdateOne {
	if u != nil {
		eiuo.SetScoreCount(*u)
	}
	return eiuo
}

// AddScoreCount adds u to the "score_count" field.
func (eiuo *ExtraInfoUpdateOne) AddScoreCount(u int32) *ExtraInfoUpdateOne {
	eiuo.mutation.AddScoreCount(u)
	return eiuo
}

// ClearScoreCount clears the value of the "score_count" field.
func (eiuo *ExtraInfoUpdateOne) ClearScoreCount() *ExtraInfoUpdateOne {
	eiuo.mutation.ClearScoreCount()
	return eiuo
}

// SetScore sets the "score" field.
func (eiuo *ExtraInfoUpdateOne) SetScore(d decimal.Decimal) *ExtraInfoUpdateOne {
	eiuo.mutation.SetScore(d)
	return eiuo
}

// SetNillableScore sets the "score" field if the given value is not nil.
func (eiuo *ExtraInfoUpdateOne) SetNillableScore(d *decimal.Decimal) *ExtraInfoUpdateOne {
	if d != nil {
		eiuo.SetScore(*d)
	}
	return eiuo
}

// ClearScore clears the value of the "score" field.
func (eiuo *ExtraInfoUpdateOne) ClearScore() *ExtraInfoUpdateOne {
	eiuo.mutation.ClearScore()
	return eiuo
}

// Mutation returns the ExtraInfoMutation object of the builder.
func (eiuo *ExtraInfoUpdateOne) Mutation() *ExtraInfoMutation {
	return eiuo.mutation
}

// Where appends a list predicates to the ExtraInfoUpdate builder.
func (eiuo *ExtraInfoUpdateOne) Where(ps ...predicate.ExtraInfo) *ExtraInfoUpdateOne {
	eiuo.mutation.Where(ps...)
	return eiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (eiuo *ExtraInfoUpdateOne) Select(field string, fields ...string) *ExtraInfoUpdateOne {
	eiuo.fields = append([]string{field}, fields...)
	return eiuo
}

// Save executes the query and returns the updated ExtraInfo entity.
func (eiuo *ExtraInfoUpdateOne) Save(ctx context.Context) (*ExtraInfo, error) {
	eiuo.defaults()
	return withHooks(ctx, eiuo.sqlSave, eiuo.mutation, eiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eiuo *ExtraInfoUpdateOne) SaveX(ctx context.Context) *ExtraInfo {
	node, err := eiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (eiuo *ExtraInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := eiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eiuo *ExtraInfoUpdateOne) ExecX(ctx context.Context) {
	if err := eiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eiuo *ExtraInfoUpdateOne) defaults() {
	if _, ok := eiuo.mutation.UpdatedAt(); !ok {
		v := extrainfo.UpdateDefaultUpdatedAt()
		eiuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (eiuo *ExtraInfoUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ExtraInfoUpdateOne {
	eiuo.modifiers = append(eiuo.modifiers, modifiers...)
	return eiuo
}

func (eiuo *ExtraInfoUpdateOne) sqlSave(ctx context.Context) (_node *ExtraInfo, err error) {
	_spec := sqlgraph.NewUpdateSpec(extrainfo.Table, extrainfo.Columns, sqlgraph.NewFieldSpec(extrainfo.FieldID, field.TypeUint32))
	id, ok := eiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "ExtraInfo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := eiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, extrainfo.FieldID)
		for _, f := range fields {
			if !extrainfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != extrainfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := eiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eiuo.mutation.EntID(); ok {
		_spec.SetField(extrainfo.FieldEntID, field.TypeUUID, value)
	}
	if value, ok := eiuo.mutation.CreatedAt(); ok {
		_spec.SetField(extrainfo.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := eiuo.mutation.AddedCreatedAt(); ok {
		_spec.AddField(extrainfo.FieldCreatedAt, field.TypeUint32, value)
	}
	if value, ok := eiuo.mutation.UpdatedAt(); ok {
		_spec.SetField(extrainfo.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := eiuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(extrainfo.FieldUpdatedAt, field.TypeUint32, value)
	}
	if value, ok := eiuo.mutation.DeletedAt(); ok {
		_spec.SetField(extrainfo.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := eiuo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(extrainfo.FieldDeletedAt, field.TypeUint32, value)
	}
	if value, ok := eiuo.mutation.AppGoodID(); ok {
		_spec.SetField(extrainfo.FieldAppGoodID, field.TypeUUID, value)
	}
	if eiuo.mutation.AppGoodIDCleared() {
		_spec.ClearField(extrainfo.FieldAppGoodID, field.TypeUUID)
	}
	if value, ok := eiuo.mutation.Likes(); ok {
		_spec.SetField(extrainfo.FieldLikes, field.TypeUint32, value)
	}
	if value, ok := eiuo.mutation.AddedLikes(); ok {
		_spec.AddField(extrainfo.FieldLikes, field.TypeUint32, value)
	}
	if eiuo.mutation.LikesCleared() {
		_spec.ClearField(extrainfo.FieldLikes, field.TypeUint32)
	}
	if value, ok := eiuo.mutation.Dislikes(); ok {
		_spec.SetField(extrainfo.FieldDislikes, field.TypeUint32, value)
	}
	if value, ok := eiuo.mutation.AddedDislikes(); ok {
		_spec.AddField(extrainfo.FieldDislikes, field.TypeUint32, value)
	}
	if eiuo.mutation.DislikesCleared() {
		_spec.ClearField(extrainfo.FieldDislikes, field.TypeUint32)
	}
	if value, ok := eiuo.mutation.RecommendCount(); ok {
		_spec.SetField(extrainfo.FieldRecommendCount, field.TypeUint32, value)
	}
	if value, ok := eiuo.mutation.AddedRecommendCount(); ok {
		_spec.AddField(extrainfo.FieldRecommendCount, field.TypeUint32, value)
	}
	if eiuo.mutation.RecommendCountCleared() {
		_spec.ClearField(extrainfo.FieldRecommendCount, field.TypeUint32)
	}
	if value, ok := eiuo.mutation.CommentCount(); ok {
		_spec.SetField(extrainfo.FieldCommentCount, field.TypeUint32, value)
	}
	if value, ok := eiuo.mutation.AddedCommentCount(); ok {
		_spec.AddField(extrainfo.FieldCommentCount, field.TypeUint32, value)
	}
	if eiuo.mutation.CommentCountCleared() {
		_spec.ClearField(extrainfo.FieldCommentCount, field.TypeUint32)
	}
	if value, ok := eiuo.mutation.ScoreCount(); ok {
		_spec.SetField(extrainfo.FieldScoreCount, field.TypeUint32, value)
	}
	if value, ok := eiuo.mutation.AddedScoreCount(); ok {
		_spec.AddField(extrainfo.FieldScoreCount, field.TypeUint32, value)
	}
	if eiuo.mutation.ScoreCountCleared() {
		_spec.ClearField(extrainfo.FieldScoreCount, field.TypeUint32)
	}
	if value, ok := eiuo.mutation.Score(); ok {
		_spec.SetField(extrainfo.FieldScore, field.TypeOther, value)
	}
	if eiuo.mutation.ScoreCleared() {
		_spec.ClearField(extrainfo.FieldScore, field.TypeOther)
	}
	_spec.AddModifiers(eiuo.modifiers...)
	_node = &ExtraInfo{config: eiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, eiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{extrainfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	eiuo.mutation.done = true
	return _node, nil
}
