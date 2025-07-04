// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/predicate"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/vendorbrand"
)

// VendorBrandDelete is the builder for deleting a VendorBrand entity.
type VendorBrandDelete struct {
	config
	hooks    []Hook
	mutation *VendorBrandMutation
}

// Where appends a list predicates to the VendorBrandDelete builder.
func (vbd *VendorBrandDelete) Where(ps ...predicate.VendorBrand) *VendorBrandDelete {
	vbd.mutation.Where(ps...)
	return vbd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vbd *VendorBrandDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, vbd.sqlExec, vbd.mutation, vbd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (vbd *VendorBrandDelete) ExecX(ctx context.Context) int {
	n, err := vbd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vbd *VendorBrandDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(vendorbrand.Table, sqlgraph.NewFieldSpec(vendorbrand.FieldID, field.TypeUint32))
	if ps := vbd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, vbd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	vbd.mutation.done = true
	return affected, err
}

// VendorBrandDeleteOne is the builder for deleting a single VendorBrand entity.
type VendorBrandDeleteOne struct {
	vbd *VendorBrandDelete
}

// Where appends a list predicates to the VendorBrandDelete builder.
func (vbdo *VendorBrandDeleteOne) Where(ps ...predicate.VendorBrand) *VendorBrandDeleteOne {
	vbdo.vbd.mutation.Where(ps...)
	return vbdo
}

// Exec executes the deletion query.
func (vbdo *VendorBrandDeleteOne) Exec(ctx context.Context) error {
	n, err := vbdo.vbd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{vendorbrand.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vbdo *VendorBrandDeleteOne) ExecX(ctx context.Context) {
	if err := vbdo.Exec(ctx); err != nil {
		panic(err)
	}
}
