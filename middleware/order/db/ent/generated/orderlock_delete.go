// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderlock"
	"github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/predicate"
)

// OrderLockDelete is the builder for deleting a OrderLock entity.
type OrderLockDelete struct {
	config
	hooks    []Hook
	mutation *OrderLockMutation
}

// Where appends a list predicates to the OrderLockDelete builder.
func (old *OrderLockDelete) Where(ps ...predicate.OrderLock) *OrderLockDelete {
	old.mutation.Where(ps...)
	return old
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (old *OrderLockDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, old.sqlExec, old.mutation, old.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (old *OrderLockDelete) ExecX(ctx context.Context) int {
	n, err := old.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (old *OrderLockDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(orderlock.Table, sqlgraph.NewFieldSpec(orderlock.FieldID, field.TypeUint32))
	if ps := old.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, old.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	old.mutation.done = true
	return affected, err
}

// OrderLockDeleteOne is the builder for deleting a single OrderLock entity.
type OrderLockDeleteOne struct {
	old *OrderLockDelete
}

// Where appends a list predicates to the OrderLockDelete builder.
func (oldo *OrderLockDeleteOne) Where(ps ...predicate.OrderLock) *OrderLockDeleteOne {
	oldo.old.mutation.Where(ps...)
	return oldo
}

// Exec executes the deletion query.
func (oldo *OrderLockDeleteOne) Exec(ctx context.Context) error {
	n, err := oldo.old.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{orderlock.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (oldo *OrderLockDeleteOne) ExecX(ctx context.Context) {
	if err := oldo.Exec(ctx); err != nil {
		panic(err)
	}
}
