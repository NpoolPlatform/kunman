// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/achievement"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/predicate"
)

// AchievementDelete is the builder for deleting a Achievement entity.
type AchievementDelete struct {
	config
	hooks    []Hook
	mutation *AchievementMutation
}

// Where appends a list predicates to the AchievementDelete builder.
func (ad *AchievementDelete) Where(ps ...predicate.Achievement) *AchievementDelete {
	ad.mutation.Where(ps...)
	return ad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ad *AchievementDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ad.sqlExec, ad.mutation, ad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ad *AchievementDelete) ExecX(ctx context.Context) int {
	n, err := ad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ad *AchievementDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(achievement.Table, sqlgraph.NewFieldSpec(achievement.FieldID, field.TypeUint32))
	if ps := ad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ad.mutation.done = true
	return affected, err
}

// AchievementDeleteOne is the builder for deleting a single Achievement entity.
type AchievementDeleteOne struct {
	ad *AchievementDelete
}

// Where appends a list predicates to the AchievementDelete builder.
func (ado *AchievementDeleteOne) Where(ps ...predicate.Achievement) *AchievementDeleteOne {
	ado.ad.mutation.Where(ps...)
	return ado
}

// Exec executes the deletion query.
func (ado *AchievementDeleteOne) Exec(ctx context.Context) error {
	n, err := ado.ad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{achievement.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ado *AchievementDeleteOne) ExecX(ctx context.Context) {
	if err := ado.Exec(ctx); err != nil {
		panic(err)
	}
}
