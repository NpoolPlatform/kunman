// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/creditallocated"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/predicate"
)

// CreditAllocatedDelete is the builder for deleting a CreditAllocated entity.
type CreditAllocatedDelete struct {
	config
	hooks    []Hook
	mutation *CreditAllocatedMutation
}

// Where appends a list predicates to the CreditAllocatedDelete builder.
func (cad *CreditAllocatedDelete) Where(ps ...predicate.CreditAllocated) *CreditAllocatedDelete {
	cad.mutation.Where(ps...)
	return cad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cad *CreditAllocatedDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cad.sqlExec, cad.mutation, cad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cad *CreditAllocatedDelete) ExecX(ctx context.Context) int {
	n, err := cad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cad *CreditAllocatedDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(creditallocated.Table, sqlgraph.NewFieldSpec(creditallocated.FieldID, field.TypeUint32))
	if ps := cad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cad.mutation.done = true
	return affected, err
}

// CreditAllocatedDeleteOne is the builder for deleting a single CreditAllocated entity.
type CreditAllocatedDeleteOne struct {
	cad *CreditAllocatedDelete
}

// Where appends a list predicates to the CreditAllocatedDelete builder.
func (cado *CreditAllocatedDeleteOne) Where(ps ...predicate.CreditAllocated) *CreditAllocatedDeleteOne {
	cado.cad.mutation.Where(ps...)
	return cado
}

// Exec executes the deletion query.
func (cado *CreditAllocatedDeleteOne) Exec(ctx context.Context) error {
	n, err := cado.cad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{creditallocated.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cado *CreditAllocatedDeleteOne) ExecX(ctx context.Context) {
	if err := cado.Exec(ctx); err != nil {
		panic(err)
	}
}
