// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/goodcoinachievement"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/predicate"
)

// GoodCoinAchievementDelete is the builder for deleting a GoodCoinAchievement entity.
type GoodCoinAchievementDelete struct {
	config
	hooks    []Hook
	mutation *GoodCoinAchievementMutation
}

// Where appends a list predicates to the GoodCoinAchievementDelete builder.
func (gcad *GoodCoinAchievementDelete) Where(ps ...predicate.GoodCoinAchievement) *GoodCoinAchievementDelete {
	gcad.mutation.Where(ps...)
	return gcad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gcad *GoodCoinAchievementDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, gcad.sqlExec, gcad.mutation, gcad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (gcad *GoodCoinAchievementDelete) ExecX(ctx context.Context) int {
	n, err := gcad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gcad *GoodCoinAchievementDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(goodcoinachievement.Table, sqlgraph.NewFieldSpec(goodcoinachievement.FieldID, field.TypeUint32))
	if ps := gcad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gcad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	gcad.mutation.done = true
	return affected, err
}

// GoodCoinAchievementDeleteOne is the builder for deleting a single GoodCoinAchievement entity.
type GoodCoinAchievementDeleteOne struct {
	gcad *GoodCoinAchievementDelete
}

// Where appends a list predicates to the GoodCoinAchievementDelete builder.
func (gcado *GoodCoinAchievementDeleteOne) Where(ps ...predicate.GoodCoinAchievement) *GoodCoinAchievementDeleteOne {
	gcado.gcad.mutation.Where(ps...)
	return gcado
}

// Exec executes the deletion query.
func (gcado *GoodCoinAchievementDeleteOne) Exec(ctx context.Context) error {
	n, err := gcado.gcad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{goodcoinachievement.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gcado *GoodCoinAchievementDeleteOne) ExecX(ctx context.Context) {
	if err := gcado.Exec(ctx); err != nil {
		panic(err)
	}
}
