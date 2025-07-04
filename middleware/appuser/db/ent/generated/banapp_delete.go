// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/banapp"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/predicate"
)

// BanAppDelete is the builder for deleting a BanApp entity.
type BanAppDelete struct {
	config
	hooks    []Hook
	mutation *BanAppMutation
}

// Where appends a list predicates to the BanAppDelete builder.
func (bad *BanAppDelete) Where(ps ...predicate.BanApp) *BanAppDelete {
	bad.mutation.Where(ps...)
	return bad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bad *BanAppDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bad.sqlExec, bad.mutation, bad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bad *BanAppDelete) ExecX(ctx context.Context) int {
	n, err := bad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bad *BanAppDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(banapp.Table, sqlgraph.NewFieldSpec(banapp.FieldID, field.TypeUint32))
	if ps := bad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bad.mutation.done = true
	return affected, err
}

// BanAppDeleteOne is the builder for deleting a single BanApp entity.
type BanAppDeleteOne struct {
	bad *BanAppDelete
}

// Where appends a list predicates to the BanAppDelete builder.
func (bado *BanAppDeleteOne) Where(ps ...predicate.BanApp) *BanAppDeleteOne {
	bado.bad.mutation.Where(ps...)
	return bado
}

// Exec executes the deletion query.
func (bado *BanAppDeleteOne) Exec(ctx context.Context) error {
	n, err := bado.bad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{banapp.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bado *BanAppDeleteOne) ExecX(ctx context.Context) {
	if err := bado.Exec(ctx); err != nil {
		panic(err)
	}
}
