// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/apppool"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/predicate"
)

// AppPoolDelete is the builder for deleting a AppPool entity.
type AppPoolDelete struct {
	config
	hooks    []Hook
	mutation *AppPoolMutation
}

// Where appends a list predicates to the AppPoolDelete builder.
func (apd *AppPoolDelete) Where(ps ...predicate.AppPool) *AppPoolDelete {
	apd.mutation.Where(ps...)
	return apd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (apd *AppPoolDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, apd.sqlExec, apd.mutation, apd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (apd *AppPoolDelete) ExecX(ctx context.Context) int {
	n, err := apd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (apd *AppPoolDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(apppool.Table, sqlgraph.NewFieldSpec(apppool.FieldID, field.TypeUint32))
	if ps := apd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, apd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	apd.mutation.done = true
	return affected, err
}

// AppPoolDeleteOne is the builder for deleting a single AppPool entity.
type AppPoolDeleteOne struct {
	apd *AppPoolDelete
}

// Where appends a list predicates to the AppPoolDelete builder.
func (apdo *AppPoolDeleteOne) Where(ps ...predicate.AppPool) *AppPoolDeleteOne {
	apdo.apd.mutation.Where(ps...)
	return apdo
}

// Exec executes the deletion query.
func (apdo *AppPoolDeleteOne) Exec(ctx context.Context) error {
	n, err := apdo.apd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{apppool.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (apdo *AppPoolDeleteOne) ExecX(ctx context.Context) {
	if err := apdo.Exec(ctx); err != nil {
		panic(err)
	}
}
