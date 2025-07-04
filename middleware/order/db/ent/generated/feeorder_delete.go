// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/feeorder"
	"github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/predicate"
)

// FeeOrderDelete is the builder for deleting a FeeOrder entity.
type FeeOrderDelete struct {
	config
	hooks    []Hook
	mutation *FeeOrderMutation
}

// Where appends a list predicates to the FeeOrderDelete builder.
func (fod *FeeOrderDelete) Where(ps ...predicate.FeeOrder) *FeeOrderDelete {
	fod.mutation.Where(ps...)
	return fod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fod *FeeOrderDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, fod.sqlExec, fod.mutation, fod.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fod *FeeOrderDelete) ExecX(ctx context.Context) int {
	n, err := fod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fod *FeeOrderDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(feeorder.Table, sqlgraph.NewFieldSpec(feeorder.FieldID, field.TypeUint32))
	if ps := fod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fod.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	fod.mutation.done = true
	return affected, err
}

// FeeOrderDeleteOne is the builder for deleting a single FeeOrder entity.
type FeeOrderDeleteOne struct {
	fod *FeeOrderDelete
}

// Where appends a list predicates to the FeeOrderDelete builder.
func (fodo *FeeOrderDeleteOne) Where(ps ...predicate.FeeOrder) *FeeOrderDeleteOne {
	fodo.fod.mutation.Where(ps...)
	return fodo
}

// Exec executes the deletion query.
func (fodo *FeeOrderDeleteOne) Exec(ctx context.Context) error {
	n, err := fodo.fod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{feeorder.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fodo *FeeOrderDeleteOne) ExecX(ctx context.Context) {
	if err := fodo.Exec(ctx); err != nil {
		panic(err)
	}
}
