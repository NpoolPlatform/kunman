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
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/topmostgoodposter"
)

// TopMostGoodPosterQuery is the builder for querying TopMostGoodPoster entities.
type TopMostGoodPosterQuery struct {
	config
	ctx        *QueryContext
	order      []topmostgoodposter.OrderOption
	inters     []Interceptor
	predicates []predicate.TopMostGoodPoster
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TopMostGoodPosterQuery builder.
func (tmgpq *TopMostGoodPosterQuery) Where(ps ...predicate.TopMostGoodPoster) *TopMostGoodPosterQuery {
	tmgpq.predicates = append(tmgpq.predicates, ps...)
	return tmgpq
}

// Limit the number of records to be returned by this query.
func (tmgpq *TopMostGoodPosterQuery) Limit(limit int) *TopMostGoodPosterQuery {
	tmgpq.ctx.Limit = &limit
	return tmgpq
}

// Offset to start from.
func (tmgpq *TopMostGoodPosterQuery) Offset(offset int) *TopMostGoodPosterQuery {
	tmgpq.ctx.Offset = &offset
	return tmgpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tmgpq *TopMostGoodPosterQuery) Unique(unique bool) *TopMostGoodPosterQuery {
	tmgpq.ctx.Unique = &unique
	return tmgpq
}

// Order specifies how the records should be ordered.
func (tmgpq *TopMostGoodPosterQuery) Order(o ...topmostgoodposter.OrderOption) *TopMostGoodPosterQuery {
	tmgpq.order = append(tmgpq.order, o...)
	return tmgpq
}

// First returns the first TopMostGoodPoster entity from the query.
// Returns a *NotFoundError when no TopMostGoodPoster was found.
func (tmgpq *TopMostGoodPosterQuery) First(ctx context.Context) (*TopMostGoodPoster, error) {
	nodes, err := tmgpq.Limit(1).All(setContextOp(ctx, tmgpq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{topmostgoodposter.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tmgpq *TopMostGoodPosterQuery) FirstX(ctx context.Context) *TopMostGoodPoster {
	node, err := tmgpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TopMostGoodPoster ID from the query.
// Returns a *NotFoundError when no TopMostGoodPoster ID was found.
func (tmgpq *TopMostGoodPosterQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = tmgpq.Limit(1).IDs(setContextOp(ctx, tmgpq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{topmostgoodposter.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tmgpq *TopMostGoodPosterQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := tmgpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TopMostGoodPoster entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TopMostGoodPoster entity is found.
// Returns a *NotFoundError when no TopMostGoodPoster entities are found.
func (tmgpq *TopMostGoodPosterQuery) Only(ctx context.Context) (*TopMostGoodPoster, error) {
	nodes, err := tmgpq.Limit(2).All(setContextOp(ctx, tmgpq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{topmostgoodposter.Label}
	default:
		return nil, &NotSingularError{topmostgoodposter.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tmgpq *TopMostGoodPosterQuery) OnlyX(ctx context.Context) *TopMostGoodPoster {
	node, err := tmgpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TopMostGoodPoster ID in the query.
// Returns a *NotSingularError when more than one TopMostGoodPoster ID is found.
// Returns a *NotFoundError when no entities are found.
func (tmgpq *TopMostGoodPosterQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = tmgpq.Limit(2).IDs(setContextOp(ctx, tmgpq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{topmostgoodposter.Label}
	default:
		err = &NotSingularError{topmostgoodposter.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tmgpq *TopMostGoodPosterQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := tmgpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TopMostGoodPosters.
func (tmgpq *TopMostGoodPosterQuery) All(ctx context.Context) ([]*TopMostGoodPoster, error) {
	ctx = setContextOp(ctx, tmgpq.ctx, ent.OpQueryAll)
	if err := tmgpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TopMostGoodPoster, *TopMostGoodPosterQuery]()
	return withInterceptors[[]*TopMostGoodPoster](ctx, tmgpq, qr, tmgpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tmgpq *TopMostGoodPosterQuery) AllX(ctx context.Context) []*TopMostGoodPoster {
	nodes, err := tmgpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TopMostGoodPoster IDs.
func (tmgpq *TopMostGoodPosterQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if tmgpq.ctx.Unique == nil && tmgpq.path != nil {
		tmgpq.Unique(true)
	}
	ctx = setContextOp(ctx, tmgpq.ctx, ent.OpQueryIDs)
	if err = tmgpq.Select(topmostgoodposter.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tmgpq *TopMostGoodPosterQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := tmgpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tmgpq *TopMostGoodPosterQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tmgpq.ctx, ent.OpQueryCount)
	if err := tmgpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tmgpq, querierCount[*TopMostGoodPosterQuery](), tmgpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tmgpq *TopMostGoodPosterQuery) CountX(ctx context.Context) int {
	count, err := tmgpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tmgpq *TopMostGoodPosterQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tmgpq.ctx, ent.OpQueryExist)
	switch _, err := tmgpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tmgpq *TopMostGoodPosterQuery) ExistX(ctx context.Context) bool {
	exist, err := tmgpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TopMostGoodPosterQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tmgpq *TopMostGoodPosterQuery) Clone() *TopMostGoodPosterQuery {
	if tmgpq == nil {
		return nil
	}
	return &TopMostGoodPosterQuery{
		config:     tmgpq.config,
		ctx:        tmgpq.ctx.Clone(),
		order:      append([]topmostgoodposter.OrderOption{}, tmgpq.order...),
		inters:     append([]Interceptor{}, tmgpq.inters...),
		predicates: append([]predicate.TopMostGoodPoster{}, tmgpq.predicates...),
		// clone intermediate query.
		sql:       tmgpq.sql.Clone(),
		path:      tmgpq.path,
		modifiers: append([]func(*sql.Selector){}, tmgpq.modifiers...),
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
//	client.TopMostGoodPoster.Query().
//		GroupBy(topmostgoodposter.FieldEntID).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (tmgpq *TopMostGoodPosterQuery) GroupBy(field string, fields ...string) *TopMostGoodPosterGroupBy {
	tmgpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TopMostGoodPosterGroupBy{build: tmgpq}
	grbuild.flds = &tmgpq.ctx.Fields
	grbuild.label = topmostgoodposter.Label
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
//	client.TopMostGoodPoster.Query().
//		Select(topmostgoodposter.FieldEntID).
//		Scan(ctx, &v)
func (tmgpq *TopMostGoodPosterQuery) Select(fields ...string) *TopMostGoodPosterSelect {
	tmgpq.ctx.Fields = append(tmgpq.ctx.Fields, fields...)
	sbuild := &TopMostGoodPosterSelect{TopMostGoodPosterQuery: tmgpq}
	sbuild.label = topmostgoodposter.Label
	sbuild.flds, sbuild.scan = &tmgpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TopMostGoodPosterSelect configured with the given aggregations.
func (tmgpq *TopMostGoodPosterQuery) Aggregate(fns ...AggregateFunc) *TopMostGoodPosterSelect {
	return tmgpq.Select().Aggregate(fns...)
}

func (tmgpq *TopMostGoodPosterQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tmgpq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tmgpq); err != nil {
				return err
			}
		}
	}
	for _, f := range tmgpq.ctx.Fields {
		if !topmostgoodposter.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if tmgpq.path != nil {
		prev, err := tmgpq.path(ctx)
		if err != nil {
			return err
		}
		tmgpq.sql = prev
	}
	return nil
}

func (tmgpq *TopMostGoodPosterQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TopMostGoodPoster, error) {
	var (
		nodes = []*TopMostGoodPoster{}
		_spec = tmgpq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TopMostGoodPoster).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TopMostGoodPoster{config: tmgpq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(tmgpq.modifiers) > 0 {
		_spec.Modifiers = tmgpq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tmgpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tmgpq *TopMostGoodPosterQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tmgpq.querySpec()
	if len(tmgpq.modifiers) > 0 {
		_spec.Modifiers = tmgpq.modifiers
	}
	_spec.Node.Columns = tmgpq.ctx.Fields
	if len(tmgpq.ctx.Fields) > 0 {
		_spec.Unique = tmgpq.ctx.Unique != nil && *tmgpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tmgpq.driver, _spec)
}

func (tmgpq *TopMostGoodPosterQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(topmostgoodposter.Table, topmostgoodposter.Columns, sqlgraph.NewFieldSpec(topmostgoodposter.FieldID, field.TypeUint32))
	_spec.From = tmgpq.sql
	if unique := tmgpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tmgpq.path != nil {
		_spec.Unique = true
	}
	if fields := tmgpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, topmostgoodposter.FieldID)
		for i := range fields {
			if fields[i] != topmostgoodposter.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tmgpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tmgpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tmgpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tmgpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tmgpq *TopMostGoodPosterQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tmgpq.driver.Dialect())
	t1 := builder.Table(topmostgoodposter.Table)
	columns := tmgpq.ctx.Fields
	if len(columns) == 0 {
		columns = topmostgoodposter.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tmgpq.sql != nil {
		selector = tmgpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tmgpq.ctx.Unique != nil && *tmgpq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range tmgpq.modifiers {
		m(selector)
	}
	for _, p := range tmgpq.predicates {
		p(selector)
	}
	for _, p := range tmgpq.order {
		p(selector)
	}
	if offset := tmgpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tmgpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (tmgpq *TopMostGoodPosterQuery) ForUpdate(opts ...sql.LockOption) *TopMostGoodPosterQuery {
	if tmgpq.driver.Dialect() == dialect.Postgres {
		tmgpq.Unique(false)
	}
	tmgpq.modifiers = append(tmgpq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return tmgpq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (tmgpq *TopMostGoodPosterQuery) ForShare(opts ...sql.LockOption) *TopMostGoodPosterQuery {
	if tmgpq.driver.Dialect() == dialect.Postgres {
		tmgpq.Unique(false)
	}
	tmgpq.modifiers = append(tmgpq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return tmgpq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tmgpq *TopMostGoodPosterQuery) Modify(modifiers ...func(s *sql.Selector)) *TopMostGoodPosterSelect {
	tmgpq.modifiers = append(tmgpq.modifiers, modifiers...)
	return tmgpq.Select()
}

// TopMostGoodPosterGroupBy is the group-by builder for TopMostGoodPoster entities.
type TopMostGoodPosterGroupBy struct {
	selector
	build *TopMostGoodPosterQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tmgpgb *TopMostGoodPosterGroupBy) Aggregate(fns ...AggregateFunc) *TopMostGoodPosterGroupBy {
	tmgpgb.fns = append(tmgpgb.fns, fns...)
	return tmgpgb
}

// Scan applies the selector query and scans the result into the given value.
func (tmgpgb *TopMostGoodPosterGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tmgpgb.build.ctx, ent.OpQueryGroupBy)
	if err := tmgpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TopMostGoodPosterQuery, *TopMostGoodPosterGroupBy](ctx, tmgpgb.build, tmgpgb, tmgpgb.build.inters, v)
}

func (tmgpgb *TopMostGoodPosterGroupBy) sqlScan(ctx context.Context, root *TopMostGoodPosterQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tmgpgb.fns))
	for _, fn := range tmgpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tmgpgb.flds)+len(tmgpgb.fns))
		for _, f := range *tmgpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tmgpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tmgpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TopMostGoodPosterSelect is the builder for selecting fields of TopMostGoodPoster entities.
type TopMostGoodPosterSelect struct {
	*TopMostGoodPosterQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tmgps *TopMostGoodPosterSelect) Aggregate(fns ...AggregateFunc) *TopMostGoodPosterSelect {
	tmgps.fns = append(tmgps.fns, fns...)
	return tmgps
}

// Scan applies the selector query and scans the result into the given value.
func (tmgps *TopMostGoodPosterSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tmgps.ctx, ent.OpQuerySelect)
	if err := tmgps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TopMostGoodPosterQuery, *TopMostGoodPosterSelect](ctx, tmgps.TopMostGoodPosterQuery, tmgps, tmgps.inters, v)
}

func (tmgps *TopMostGoodPosterSelect) sqlScan(ctx context.Context, root *TopMostGoodPosterQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tmgps.fns))
	for _, fn := range tmgps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tmgps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tmgps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tmgps *TopMostGoodPosterSelect) Modify(modifiers ...func(s *sql.Selector)) *TopMostGoodPosterSelect {
	tmgps.modifiers = append(tmgps.modifiers, modifiers...)
	return tmgps
}
