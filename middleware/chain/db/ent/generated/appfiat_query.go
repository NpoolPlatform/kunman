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
	"github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/appfiat"
	"github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/predicate"
)

// AppFiatQuery is the builder for querying AppFiat entities.
type AppFiatQuery struct {
	config
	ctx        *QueryContext
	order      []appfiat.OrderOption
	inters     []Interceptor
	predicates []predicate.AppFiat
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppFiatQuery builder.
func (afq *AppFiatQuery) Where(ps ...predicate.AppFiat) *AppFiatQuery {
	afq.predicates = append(afq.predicates, ps...)
	return afq
}

// Limit the number of records to be returned by this query.
func (afq *AppFiatQuery) Limit(limit int) *AppFiatQuery {
	afq.ctx.Limit = &limit
	return afq
}

// Offset to start from.
func (afq *AppFiatQuery) Offset(offset int) *AppFiatQuery {
	afq.ctx.Offset = &offset
	return afq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (afq *AppFiatQuery) Unique(unique bool) *AppFiatQuery {
	afq.ctx.Unique = &unique
	return afq
}

// Order specifies how the records should be ordered.
func (afq *AppFiatQuery) Order(o ...appfiat.OrderOption) *AppFiatQuery {
	afq.order = append(afq.order, o...)
	return afq
}

// First returns the first AppFiat entity from the query.
// Returns a *NotFoundError when no AppFiat was found.
func (afq *AppFiatQuery) First(ctx context.Context) (*AppFiat, error) {
	nodes, err := afq.Limit(1).All(setContextOp(ctx, afq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appfiat.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (afq *AppFiatQuery) FirstX(ctx context.Context) *AppFiat {
	node, err := afq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppFiat ID from the query.
// Returns a *NotFoundError when no AppFiat ID was found.
func (afq *AppFiatQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = afq.Limit(1).IDs(setContextOp(ctx, afq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appfiat.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (afq *AppFiatQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := afq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppFiat entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppFiat entity is found.
// Returns a *NotFoundError when no AppFiat entities are found.
func (afq *AppFiatQuery) Only(ctx context.Context) (*AppFiat, error) {
	nodes, err := afq.Limit(2).All(setContextOp(ctx, afq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appfiat.Label}
	default:
		return nil, &NotSingularError{appfiat.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (afq *AppFiatQuery) OnlyX(ctx context.Context) *AppFiat {
	node, err := afq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppFiat ID in the query.
// Returns a *NotSingularError when more than one AppFiat ID is found.
// Returns a *NotFoundError when no entities are found.
func (afq *AppFiatQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = afq.Limit(2).IDs(setContextOp(ctx, afq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appfiat.Label}
	default:
		err = &NotSingularError{appfiat.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (afq *AppFiatQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := afq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppFiats.
func (afq *AppFiatQuery) All(ctx context.Context) ([]*AppFiat, error) {
	ctx = setContextOp(ctx, afq.ctx, ent.OpQueryAll)
	if err := afq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AppFiat, *AppFiatQuery]()
	return withInterceptors[[]*AppFiat](ctx, afq, qr, afq.inters)
}

// AllX is like All, but panics if an error occurs.
func (afq *AppFiatQuery) AllX(ctx context.Context) []*AppFiat {
	nodes, err := afq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppFiat IDs.
func (afq *AppFiatQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if afq.ctx.Unique == nil && afq.path != nil {
		afq.Unique(true)
	}
	ctx = setContextOp(ctx, afq.ctx, ent.OpQueryIDs)
	if err = afq.Select(appfiat.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (afq *AppFiatQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := afq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (afq *AppFiatQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, afq.ctx, ent.OpQueryCount)
	if err := afq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, afq, querierCount[*AppFiatQuery](), afq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (afq *AppFiatQuery) CountX(ctx context.Context) int {
	count, err := afq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (afq *AppFiatQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, afq.ctx, ent.OpQueryExist)
	switch _, err := afq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (afq *AppFiatQuery) ExistX(ctx context.Context) bool {
	exist, err := afq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppFiatQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (afq *AppFiatQuery) Clone() *AppFiatQuery {
	if afq == nil {
		return nil
	}
	return &AppFiatQuery{
		config:     afq.config,
		ctx:        afq.ctx.Clone(),
		order:      append([]appfiat.OrderOption{}, afq.order...),
		inters:     append([]Interceptor{}, afq.inters...),
		predicates: append([]predicate.AppFiat{}, afq.predicates...),
		// clone intermediate query.
		sql:       afq.sql.Clone(),
		path:      afq.path,
		modifiers: append([]func(*sql.Selector){}, afq.modifiers...),
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AppFiat.Query().
//		GroupBy(appfiat.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (afq *AppFiatQuery) GroupBy(field string, fields ...string) *AppFiatGroupBy {
	afq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AppFiatGroupBy{build: afq}
	grbuild.flds = &afq.ctx.Fields
	grbuild.label = appfiat.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.AppFiat.Query().
//		Select(appfiat.FieldCreatedAt).
//		Scan(ctx, &v)
func (afq *AppFiatQuery) Select(fields ...string) *AppFiatSelect {
	afq.ctx.Fields = append(afq.ctx.Fields, fields...)
	sbuild := &AppFiatSelect{AppFiatQuery: afq}
	sbuild.label = appfiat.Label
	sbuild.flds, sbuild.scan = &afq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AppFiatSelect configured with the given aggregations.
func (afq *AppFiatQuery) Aggregate(fns ...AggregateFunc) *AppFiatSelect {
	return afq.Select().Aggregate(fns...)
}

func (afq *AppFiatQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range afq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, afq); err != nil {
				return err
			}
		}
	}
	for _, f := range afq.ctx.Fields {
		if !appfiat.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if afq.path != nil {
		prev, err := afq.path(ctx)
		if err != nil {
			return err
		}
		afq.sql = prev
	}
	return nil
}

func (afq *AppFiatQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppFiat, error) {
	var (
		nodes = []*AppFiat{}
		_spec = afq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AppFiat).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AppFiat{config: afq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(afq.modifiers) > 0 {
		_spec.Modifiers = afq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, afq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (afq *AppFiatQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := afq.querySpec()
	if len(afq.modifiers) > 0 {
		_spec.Modifiers = afq.modifiers
	}
	_spec.Node.Columns = afq.ctx.Fields
	if len(afq.ctx.Fields) > 0 {
		_spec.Unique = afq.ctx.Unique != nil && *afq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, afq.driver, _spec)
}

func (afq *AppFiatQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(appfiat.Table, appfiat.Columns, sqlgraph.NewFieldSpec(appfiat.FieldID, field.TypeUint32))
	_spec.From = afq.sql
	if unique := afq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if afq.path != nil {
		_spec.Unique = true
	}
	if fields := afq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appfiat.FieldID)
		for i := range fields {
			if fields[i] != appfiat.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := afq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := afq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := afq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := afq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (afq *AppFiatQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(afq.driver.Dialect())
	t1 := builder.Table(appfiat.Table)
	columns := afq.ctx.Fields
	if len(columns) == 0 {
		columns = appfiat.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if afq.sql != nil {
		selector = afq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if afq.ctx.Unique != nil && *afq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range afq.modifiers {
		m(selector)
	}
	for _, p := range afq.predicates {
		p(selector)
	}
	for _, p := range afq.order {
		p(selector)
	}
	if offset := afq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := afq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (afq *AppFiatQuery) ForUpdate(opts ...sql.LockOption) *AppFiatQuery {
	if afq.driver.Dialect() == dialect.Postgres {
		afq.Unique(false)
	}
	afq.modifiers = append(afq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return afq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (afq *AppFiatQuery) ForShare(opts ...sql.LockOption) *AppFiatQuery {
	if afq.driver.Dialect() == dialect.Postgres {
		afq.Unique(false)
	}
	afq.modifiers = append(afq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return afq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (afq *AppFiatQuery) Modify(modifiers ...func(s *sql.Selector)) *AppFiatSelect {
	afq.modifiers = append(afq.modifiers, modifiers...)
	return afq.Select()
}

// AppFiatGroupBy is the group-by builder for AppFiat entities.
type AppFiatGroupBy struct {
	selector
	build *AppFiatQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (afgb *AppFiatGroupBy) Aggregate(fns ...AggregateFunc) *AppFiatGroupBy {
	afgb.fns = append(afgb.fns, fns...)
	return afgb
}

// Scan applies the selector query and scans the result into the given value.
func (afgb *AppFiatGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, afgb.build.ctx, ent.OpQueryGroupBy)
	if err := afgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppFiatQuery, *AppFiatGroupBy](ctx, afgb.build, afgb, afgb.build.inters, v)
}

func (afgb *AppFiatGroupBy) sqlScan(ctx context.Context, root *AppFiatQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(afgb.fns))
	for _, fn := range afgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*afgb.flds)+len(afgb.fns))
		for _, f := range *afgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*afgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := afgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AppFiatSelect is the builder for selecting fields of AppFiat entities.
type AppFiatSelect struct {
	*AppFiatQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (afs *AppFiatSelect) Aggregate(fns ...AggregateFunc) *AppFiatSelect {
	afs.fns = append(afs.fns, fns...)
	return afs
}

// Scan applies the selector query and scans the result into the given value.
func (afs *AppFiatSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, afs.ctx, ent.OpQuerySelect)
	if err := afs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppFiatQuery, *AppFiatSelect](ctx, afs.AppFiatQuery, afs, afs.inters, v)
}

func (afs *AppFiatSelect) sqlScan(ctx context.Context, root *AppFiatQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(afs.fns))
	for _, fn := range afs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*afs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := afs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (afs *AppFiatSelect) Modify(modifiers ...func(s *sql.Selector)) *AppFiatSelect {
	afs.modifiers = append(afs.modifiers, modifiers...)
	return afs
}
