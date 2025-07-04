// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/frontendtemplate"
	"github.com/NpoolPlatform/kunman/middleware/notif/db/ent/generated/predicate"
)

// FrontendTemplateDelete is the builder for deleting a FrontendTemplate entity.
type FrontendTemplateDelete struct {
	config
	hooks    []Hook
	mutation *FrontendTemplateMutation
}

// Where appends a list predicates to the FrontendTemplateDelete builder.
func (ftd *FrontendTemplateDelete) Where(ps ...predicate.FrontendTemplate) *FrontendTemplateDelete {
	ftd.mutation.Where(ps...)
	return ftd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ftd *FrontendTemplateDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ftd.sqlExec, ftd.mutation, ftd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ftd *FrontendTemplateDelete) ExecX(ctx context.Context) int {
	n, err := ftd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ftd *FrontendTemplateDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(frontendtemplate.Table, sqlgraph.NewFieldSpec(frontendtemplate.FieldID, field.TypeUint32))
	if ps := ftd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ftd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ftd.mutation.done = true
	return affected, err
}

// FrontendTemplateDeleteOne is the builder for deleting a single FrontendTemplate entity.
type FrontendTemplateDeleteOne struct {
	ftd *FrontendTemplateDelete
}

// Where appends a list predicates to the FrontendTemplateDelete builder.
func (ftdo *FrontendTemplateDeleteOne) Where(ps ...predicate.FrontendTemplate) *FrontendTemplateDeleteOne {
	ftdo.ftd.mutation.Where(ps...)
	return ftdo
}

// Exec executes the deletion query.
func (ftdo *FrontendTemplateDeleteOne) Exec(ctx context.Context) error {
	n, err := ftdo.ftd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{frontendtemplate.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ftdo *FrontendTemplateDeleteOne) ExecX(ctx context.Context) {
	if err := ftdo.Exec(ctx); err != nil {
		panic(err)
	}
}
