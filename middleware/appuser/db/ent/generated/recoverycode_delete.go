// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/predicate"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/recoverycode"
)

// RecoveryCodeDelete is the builder for deleting a RecoveryCode entity.
type RecoveryCodeDelete struct {
	config
	hooks    []Hook
	mutation *RecoveryCodeMutation
}

// Where appends a list predicates to the RecoveryCodeDelete builder.
func (rcd *RecoveryCodeDelete) Where(ps ...predicate.RecoveryCode) *RecoveryCodeDelete {
	rcd.mutation.Where(ps...)
	return rcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rcd *RecoveryCodeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rcd.sqlExec, rcd.mutation, rcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rcd *RecoveryCodeDelete) ExecX(ctx context.Context) int {
	n, err := rcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rcd *RecoveryCodeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(recoverycode.Table, sqlgraph.NewFieldSpec(recoverycode.FieldID, field.TypeUint32))
	if ps := rcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rcd.mutation.done = true
	return affected, err
}

// RecoveryCodeDeleteOne is the builder for deleting a single RecoveryCode entity.
type RecoveryCodeDeleteOne struct {
	rcd *RecoveryCodeDelete
}

// Where appends a list predicates to the RecoveryCodeDelete builder.
func (rcdo *RecoveryCodeDeleteOne) Where(ps ...predicate.RecoveryCode) *RecoveryCodeDeleteOne {
	rcdo.rcd.mutation.Where(ps...)
	return rcdo
}

// Exec executes the deletion query.
func (rcdo *RecoveryCodeDeleteOne) Exec(ctx context.Context) error {
	n, err := rcdo.rcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{recoverycode.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rcdo *RecoveryCodeDeleteOne) ExecX(ctx context.Context) {
	if err := rcdo.Exec(ctx); err != nil {
		panic(err)
	}
}
