// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodrewardhistory"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/predicate"
)

// GoodRewardHistoryDelete is the builder for deleting a GoodRewardHistory entity.
type GoodRewardHistoryDelete struct {
	config
	hooks    []Hook
	mutation *GoodRewardHistoryMutation
}

// Where appends a list predicates to the GoodRewardHistoryDelete builder.
func (grhd *GoodRewardHistoryDelete) Where(ps ...predicate.GoodRewardHistory) *GoodRewardHistoryDelete {
	grhd.mutation.Where(ps...)
	return grhd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (grhd *GoodRewardHistoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, grhd.sqlExec, grhd.mutation, grhd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (grhd *GoodRewardHistoryDelete) ExecX(ctx context.Context) int {
	n, err := grhd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (grhd *GoodRewardHistoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(goodrewardhistory.Table, sqlgraph.NewFieldSpec(goodrewardhistory.FieldID, field.TypeUint32))
	if ps := grhd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, grhd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	grhd.mutation.done = true
	return affected, err
}

// GoodRewardHistoryDeleteOne is the builder for deleting a single GoodRewardHistory entity.
type GoodRewardHistoryDeleteOne struct {
	grhd *GoodRewardHistoryDelete
}

// Where appends a list predicates to the GoodRewardHistoryDelete builder.
func (grhdo *GoodRewardHistoryDeleteOne) Where(ps ...predicate.GoodRewardHistory) *GoodRewardHistoryDeleteOne {
	grhdo.grhd.mutation.Where(ps...)
	return grhdo
}

// Exec executes the deletion query.
func (grhdo *GoodRewardHistoryDeleteOne) Exec(ctx context.Context) error {
	n, err := grhdo.grhd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{goodrewardhistory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (grhdo *GoodRewardHistoryDeleteOne) ExecX(ctx context.Context) {
	if err := grhdo.Exec(ctx); err != nil {
		panic(err)
	}
}
