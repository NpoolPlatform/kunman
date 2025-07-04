// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/predicate"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/topmostconstraint"
)

// TopMostConstraintDelete is the builder for deleting a TopMostConstraint entity.
type TopMostConstraintDelete struct {
	config
	hooks    []Hook
	mutation *TopMostConstraintMutation
}

// Where appends a list predicates to the TopMostConstraintDelete builder.
func (tmcd *TopMostConstraintDelete) Where(ps ...predicate.TopMostConstraint) *TopMostConstraintDelete {
	tmcd.mutation.Where(ps...)
	return tmcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tmcd *TopMostConstraintDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, tmcd.sqlExec, tmcd.mutation, tmcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tmcd *TopMostConstraintDelete) ExecX(ctx context.Context) int {
	n, err := tmcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tmcd *TopMostConstraintDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(topmostconstraint.Table, sqlgraph.NewFieldSpec(topmostconstraint.FieldID, field.TypeUint32))
	if ps := tmcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tmcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tmcd.mutation.done = true
	return affected, err
}

// TopMostConstraintDeleteOne is the builder for deleting a single TopMostConstraint entity.
type TopMostConstraintDeleteOne struct {
	tmcd *TopMostConstraintDelete
}

// Where appends a list predicates to the TopMostConstraintDelete builder.
func (tmcdo *TopMostConstraintDeleteOne) Where(ps ...predicate.TopMostConstraint) *TopMostConstraintDeleteOne {
	tmcdo.tmcd.mutation.Where(ps...)
	return tmcdo
}

// Exec executes the deletion query.
func (tmcdo *TopMostConstraintDeleteOne) Exec(ctx context.Context) error {
	n, err := tmcdo.tmcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{topmostconstraint.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tmcdo *TopMostConstraintDeleteOne) ExecX(ctx context.Context) {
	if err := tmcdo.Exec(ctx); err != nil {
		panic(err)
	}
}
