// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/orderbenefit"
	"github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/predicate"
)

// OrderBenefitDelete is the builder for deleting a OrderBenefit entity.
type OrderBenefitDelete struct {
	config
	hooks    []Hook
	mutation *OrderBenefitMutation
}

// Where appends a list predicates to the OrderBenefitDelete builder.
func (obd *OrderBenefitDelete) Where(ps ...predicate.OrderBenefit) *OrderBenefitDelete {
	obd.mutation.Where(ps...)
	return obd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (obd *OrderBenefitDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, obd.sqlExec, obd.mutation, obd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (obd *OrderBenefitDelete) ExecX(ctx context.Context) int {
	n, err := obd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (obd *OrderBenefitDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(orderbenefit.Table, sqlgraph.NewFieldSpec(orderbenefit.FieldID, field.TypeUint32))
	if ps := obd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, obd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	obd.mutation.done = true
	return affected, err
}

// OrderBenefitDeleteOne is the builder for deleting a single OrderBenefit entity.
type OrderBenefitDeleteOne struct {
	obd *OrderBenefitDelete
}

// Where appends a list predicates to the OrderBenefitDelete builder.
func (obdo *OrderBenefitDeleteOne) Where(ps ...predicate.OrderBenefit) *OrderBenefitDeleteOne {
	obdo.obd.mutation.Where(ps...)
	return obdo
}

// Exec executes the deletion query.
func (obdo *OrderBenefitDeleteOne) Exec(ctx context.Context) error {
	n, err := obdo.obd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{orderbenefit.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (obdo *OrderBenefitDeleteOne) ExecX(ctx context.Context) {
	if err := obdo.Exec(ctx); err != nil {
		panic(err)
	}
}
