// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/predicate"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/subscriber"
)

// SubscriberDelete is the builder for deleting a Subscriber entity.
type SubscriberDelete struct {
	config
	hooks    []Hook
	mutation *SubscriberMutation
}

// Where appends a list predicates to the SubscriberDelete builder.
func (sd *SubscriberDelete) Where(ps ...predicate.Subscriber) *SubscriberDelete {
	sd.mutation.Where(ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *SubscriberDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sd.sqlExec, sd.mutation, sd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *SubscriberDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *SubscriberDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(subscriber.Table, sqlgraph.NewFieldSpec(subscriber.FieldID, field.TypeUint32))
	if ps := sd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sd.mutation.done = true
	return affected, err
}

// SubscriberDeleteOne is the builder for deleting a single Subscriber entity.
type SubscriberDeleteOne struct {
	sd *SubscriberDelete
}

// Where appends a list predicates to the SubscriberDelete builder.
func (sdo *SubscriberDeleteOne) Where(ps ...predicate.Subscriber) *SubscriberDeleteOne {
	sdo.sd.mutation.Where(ps...)
	return sdo
}

// Exec executes the deletion query.
func (sdo *SubscriberDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{subscriber.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *SubscriberDeleteOne) ExecX(ctx context.Context) {
	if err := sdo.Exec(ctx); err != nil {
		panic(err)
	}
}
