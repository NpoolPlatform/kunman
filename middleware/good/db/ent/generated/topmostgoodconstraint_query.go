// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/predicate"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/topmostgoodconstraint"
)

// TopMostGoodConstraintQuery is the builder for querying TopMostGoodConstraint entities.
type TopMostGoodConstraintQuery struct {
	config
	ctx        *QueryContext
	order      []topmostgoodconstraint.OrderOption
	inters     []Interceptor
	predicates []predicate.TopMostGoodConstraint
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TopMostGoodConstraintQuery builder.
func (tmgcq *TopMostGoodConstraintQuery) Where(ps ...predicate.TopMostGoodConstraint) *TopMostGoodConstraintQuery {
	tmgcq.predicates = append(tmgcq.predicates, ps...)
	return tmgcq
}

// Limit the number of records to be returned by this query.
func (tmgcq *TopMostGoodConstraintQuery) Limit(limit int) *TopMostGoodConstraintQuery {
	tmgcq.ctx.Limit = &limit
	return tmgcq
}

// Offset to start from.
func (tmgcq *TopMostGoodConstraintQuery) Offset(offset int) *TopMostGoodConstraintQuery {
	tmgcq.ctx.Offset = &offset
	return tmgcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tmgcq *TopMostGoodConstraintQuery) Unique(unique bool) *TopMostGoodConstraintQuery {
	tmgcq.ctx.Unique = &unique
	return tmgcq
}

// Order specifies how the records should be ordered.
func (tmgcq *TopMostGoodConstraintQuery) Order(o ...topmostgoodconstraint.OrderOption) *TopMostGoodConstraintQuery {
	tmgcq.order = append(tmgcq.order, o...)
	return tmgcq
}

// First returns the first TopMostGoodConstraint entity from the query.
// Returns a *NotFoundError when no TopMostGoodConstraint was found.
func (tmgcq *TopMostGoodConstraintQuery) First(ctx context.Context) (*TopMostGoodConstraint, error) {
	nodes, err := tmgcq.Limit(1).All(setContextOp(ctx, tmgcq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{topmostgoodconstraint.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tmgcq *TopMostGoodConstraintQuery) FirstX(ctx context.Context) *TopMostGoodConstraint {
	node, err := tmgcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TopMostGoodConstraint ID from the query.
// Returns a *NotFoundError when no TopMostGoodConstraint ID was found.
func (tmgcq *TopMostGoodConstraintQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = tmgcq.Limit(1).IDs(setContextOp(ctx, tmgcq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{topmostgoodconstraint.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tmgcq *TopMostGoodConstraintQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := tmgcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TopMostGoodConstraint entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TopMostGoodConstraint entity is found.
// Returns a *NotFoundError when no TopMostGoodConstraint entities are found.
func (tmgcq *TopMostGoodConstraintQuery) Only(ctx context.Context) (*TopMostGoodConstraint, error) {
	nodes, err := tmgcq.Limit(2).All(setContextOp(ctx, tmgcq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{topmostgoodconstraint.Label}
	default:
		return nil, &NotSingularError{topmostgoodconstraint.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tmgcq *TopMostGoodConstraintQuery) OnlyX(ctx context.Context) *TopMostGoodConstraint {
	node, err := tmgcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TopMostGoodConstraint ID in the query.
// Returns a *NotSingularError when more than one TopMostGoodConstraint ID is found.
// Returns a *NotFoundError when no entities are found.
func (tmgcq *TopMostGoodConstraintQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = tmgcq.Limit(2).IDs(setContextOp(ctx, tmgcq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{topmostgoodconstraint.Label}
	default:
		err = &NotSingularError{topmostgoodconstraint.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tmgcq *TopMostGoodConstraintQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := tmgcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TopMostGoodConstraints.
func (tmgcq *TopMostGoodConstraintQuery) All(ctx context.Context) ([]*TopMostGoodConstraint, error) {
	ctx = setContextOp(ctx, tmgcq.ctx, ent.OpQueryAll)
	if err := tmgcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TopMostGoodConstraint, *TopMostGoodConstraintQuery]()
	return withInterceptors[[]*TopMostGoodConstraint](ctx, tmgcq, qr, tmgcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tmgcq *TopMostGoodConstraintQuery) AllX(ctx context.Context) []*TopMostGoodConstraint {
	nodes, err := tmgcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TopMostGoodConstraint IDs.
func (tmgcq *TopMostGoodConstraintQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if tmgcq.ctx.Unique == nil && tmgcq.path != nil {
		tmgcq.Unique(true)
	}
	ctx = setContextOp(ctx, tmgcq.ctx, ent.OpQueryIDs)
	if err = tmgcq.Select(topmostgoodconstraint.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tmgcq *TopMostGoodConstraintQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := tmgcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tmgcq *TopMostGoodConstraintQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tmgcq.ctx, ent.OpQueryCount)
	if err := tmgcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tmgcq, querierCount[*TopMostGoodConstraintQuery](), tmgcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tmgcq *TopMostGoodConstraintQuery) CountX(ctx context.Context) int {
	count, err := tmgcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tmgcq *TopMostGoodConstraintQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tmgcq.ctx, ent.OpQueryExist)
	switch _, err := tmgcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tmgcq *TopMostGoodConstraintQuery) ExistX(ctx context.Context) bool {
	exist, err := tmgcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TopMostGoodConstraintQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tmgcq *TopMostGoodConstraintQuery) Clone() *TopMostGoodConstraintQuery {
	if tmgcq == nil {
		return nil
	}
	return &TopMostGoodConstraintQuery{
		config:     tmgcq.config,
		ctx:        tmgcq.ctx.Clone(),
		order:      append([]topmostgoodconstraint.OrderOption{}, tmgcq.order...),
		inters:     append([]Interceptor{}, tmgcq.inters...),
		predicates: append([]predicate.TopMostGoodConstraint{}, tmgcq.predicates...),
		// clone intermediate query.
		sql:       tmgcq.sql.Clone(),
		path:      tmgcq.path,
		modifiers: append([]func(*sql.Selector){}, tmgcq.modifiers...),
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		EntID uuid.UUID `json:"ent_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TopMostGoodConstraint.Query().
//		GroupBy(topmostgoodconstraint.FieldEntID).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (tmgcq *TopMostGoodConstraintQuery) GroupBy(field string, fields ...string) *TopMostGoodConstraintGroupBy {
	tmgcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TopMostGoodConstraintGroupBy{build: tmgcq}
	grbuild.flds = &tmgcq.ctx.Fields
	grbuild.label = topmostgoodconstraint.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		EntID uuid.UUID `json:"ent_id,omitempty"`
//	}
//
//	client.TopMostGoodConstraint.Query().
//		Select(topmostgoodconstraint.FieldEntID).
//		Scan(ctx, &v)
func (tmgcq *TopMostGoodConstraintQuery) Select(fields ...string) *TopMostGoodConstraintSelect {
	tmgcq.ctx.Fields = append(tmgcq.ctx.Fields, fields...)
	sbuild := &TopMostGoodConstraintSelect{TopMostGoodConstraintQuery: tmgcq}
	sbuild.label = topmostgoodconstraint.Label
	sbuild.flds, sbuild.scan = &tmgcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TopMostGoodConstraintSelect configured with the given aggregations.
func (tmgcq *TopMostGoodConstraintQuery) Aggregate(fns ...AggregateFunc) *TopMostGoodConstraintSelect {
	return tmgcq.Select().Aggregate(fns...)
}

func (tmgcq *TopMostGoodConstraintQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tmgcq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tmgcq); err != nil {
				return err
			}
		}
	}
	for _, f := range tmgcq.ctx.Fields {
		if !topmostgoodconstraint.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if tmgcq.path != nil {
		prev, err := tmgcq.path(ctx)
		if err != nil {
			return err
		}
		tmgcq.sql = prev
	}
	return nil
}

func (tmgcq *TopMostGoodConstraintQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TopMostGoodConstraint, error) {
	var (
		nodes = []*TopMostGoodConstraint{}
		_spec = tmgcq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TopMostGoodConstraint).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TopMostGoodConstraint{config: tmgcq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(tmgcq.modifiers) > 0 {
		_spec.Modifiers = tmgcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tmgcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tmgcq *TopMostGoodConstraintQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tmgcq.querySpec()
	if len(tmgcq.modifiers) > 0 {
		_spec.Modifiers = tmgcq.modifiers
	}
	_spec.Node.Columns = tmgcq.ctx.Fields
	if len(tmgcq.ctx.Fields) > 0 {
		_spec.Unique = tmgcq.ctx.Unique != nil && *tmgcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tmgcq.driver, _spec)
}

func (tmgcq *TopMostGoodConstraintQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(topmostgoodconstraint.Table, topmostgoodconstraint.Columns, sqlgraph.NewFieldSpec(topmostgoodconstraint.FieldID, field.TypeUint32))
	_spec.From = tmgcq.sql
	if unique := tmgcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tmgcq.path != nil {
		_spec.Unique = true
	}
	if fields := tmgcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, topmostgoodconstraint.FieldID)
		for i := range fields {
			if fields[i] != topmostgoodconstraint.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tmgcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tmgcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tmgcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tmgcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tmgcq *TopMostGoodConstraintQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tmgcq.driver.Dialect())
	t1 := builder.Table(topmostgoodconstraint.Table)
	columns := tmgcq.ctx.Fields
	if len(columns) == 0 {
		columns = topmostgoodconstraint.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tmgcq.sql != nil {
		selector = tmgcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tmgcq.ctx.Unique != nil && *tmgcq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range tmgcq.modifiers {
		m(selector)
	}
	for _, p := range tmgcq.predicates {
		p(selector)
	}
	for _, p := range tmgcq.order {
		p(selector)
	}
	if offset := tmgcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tmgcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (tmgcq *TopMostGoodConstraintQuery) ForUpdate(opts ...sql.LockOption) *TopMostGoodConstraintQuery {
	if tmgcq.driver.Dialect() == dialect.Postgres {
		tmgcq.Unique(false)
	}
	tmgcq.modifiers = append(tmgcq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return tmgcq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (tmgcq *TopMostGoodConstraintQuery) ForShare(opts ...sql.LockOption) *TopMostGoodConstraintQuery {
	if tmgcq.driver.Dialect() == dialect.Postgres {
		tmgcq.Unique(false)
	}
	tmgcq.modifiers = append(tmgcq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return tmgcq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tmgcq *TopMostGoodConstraintQuery) Modify(modifiers ...func(s *sql.Selector)) *TopMostGoodConstraintSelect {
	tmgcq.modifiers = append(tmgcq.modifiers, modifiers...)
	return tmgcq.Select()
}

// TopMostGoodConstraintGroupBy is the group-by builder for TopMostGoodConstraint entities.
type TopMostGoodConstraintGroupBy struct {
	selector
	build *TopMostGoodConstraintQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tmgcgb *TopMostGoodConstraintGroupBy) Aggregate(fns ...AggregateFunc) *TopMostGoodConstraintGroupBy {
	tmgcgb.fns = append(tmgcgb.fns, fns...)
	return tmgcgb
}

// Scan applies the selector query and scans the result into the given value.
func (tmgcgb *TopMostGoodConstraintGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tmgcgb.build.ctx, ent.OpQueryGroupBy)
	if err := tmgcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TopMostGoodConstraintQuery, *TopMostGoodConstraintGroupBy](ctx, tmgcgb.build, tmgcgb, tmgcgb.build.inters, v)
}

func (tmgcgb *TopMostGoodConstraintGroupBy) sqlScan(ctx context.Context, root *TopMostGoodConstraintQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tmgcgb.fns))
	for _, fn := range tmgcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tmgcgb.flds)+len(tmgcgb.fns))
		for _, f := range *tmgcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tmgcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tmgcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TopMostGoodConstraintSelect is the builder for selecting fields of TopMostGoodConstraint entities.
type TopMostGoodConstraintSelect struct {
	*TopMostGoodConstraintQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tmgcs *TopMostGoodConstraintSelect) Aggregate(fns ...AggregateFunc) *TopMostGoodConstraintSelect {
	tmgcs.fns = append(tmgcs.fns, fns...)
	return tmgcs
}

// Scan applies the selector query and scans the result into the given value.
func (tmgcs *TopMostGoodConstraintSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tmgcs.ctx, ent.OpQuerySelect)
	if err := tmgcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TopMostGoodConstraintQuery, *TopMostGoodConstraintSelect](ctx, tmgcs.TopMostGoodConstraintQuery, tmgcs, tmgcs.inters, v)
}

func (tmgcs *TopMostGoodConstraintSelect) sqlScan(ctx context.Context, root *TopMostGoodConstraintQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tmgcs.fns))
	for _, fn := range tmgcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tmgcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tmgcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tmgcs *TopMostGoodConstraintSelect) Modify(modifiers ...func(s *sql.Selector)) *TopMostGoodConstraintSelect {
	tmgcs.modifiers = append(tmgcs.modifiers, modifiers...)
	return tmgcs
}
