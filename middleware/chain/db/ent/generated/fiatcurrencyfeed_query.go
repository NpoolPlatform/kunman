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
	"github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/fiatcurrencyfeed"
	"github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/predicate"
)

// FiatCurrencyFeedQuery is the builder for querying FiatCurrencyFeed entities.
type FiatCurrencyFeedQuery struct {
	config
	ctx        *QueryContext
	order      []fiatcurrencyfeed.OrderOption
	inters     []Interceptor
	predicates []predicate.FiatCurrencyFeed
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FiatCurrencyFeedQuery builder.
func (fcfq *FiatCurrencyFeedQuery) Where(ps ...predicate.FiatCurrencyFeed) *FiatCurrencyFeedQuery {
	fcfq.predicates = append(fcfq.predicates, ps...)
	return fcfq
}

// Limit the number of records to be returned by this query.
func (fcfq *FiatCurrencyFeedQuery) Limit(limit int) *FiatCurrencyFeedQuery {
	fcfq.ctx.Limit = &limit
	return fcfq
}

// Offset to start from.
func (fcfq *FiatCurrencyFeedQuery) Offset(offset int) *FiatCurrencyFeedQuery {
	fcfq.ctx.Offset = &offset
	return fcfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fcfq *FiatCurrencyFeedQuery) Unique(unique bool) *FiatCurrencyFeedQuery {
	fcfq.ctx.Unique = &unique
	return fcfq
}

// Order specifies how the records should be ordered.
func (fcfq *FiatCurrencyFeedQuery) Order(o ...fiatcurrencyfeed.OrderOption) *FiatCurrencyFeedQuery {
	fcfq.order = append(fcfq.order, o...)
	return fcfq
}

// First returns the first FiatCurrencyFeed entity from the query.
// Returns a *NotFoundError when no FiatCurrencyFeed was found.
func (fcfq *FiatCurrencyFeedQuery) First(ctx context.Context) (*FiatCurrencyFeed, error) {
	nodes, err := fcfq.Limit(1).All(setContextOp(ctx, fcfq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{fiatcurrencyfeed.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fcfq *FiatCurrencyFeedQuery) FirstX(ctx context.Context) *FiatCurrencyFeed {
	node, err := fcfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FiatCurrencyFeed ID from the query.
// Returns a *NotFoundError when no FiatCurrencyFeed ID was found.
func (fcfq *FiatCurrencyFeedQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = fcfq.Limit(1).IDs(setContextOp(ctx, fcfq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{fiatcurrencyfeed.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fcfq *FiatCurrencyFeedQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := fcfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FiatCurrencyFeed entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FiatCurrencyFeed entity is found.
// Returns a *NotFoundError when no FiatCurrencyFeed entities are found.
func (fcfq *FiatCurrencyFeedQuery) Only(ctx context.Context) (*FiatCurrencyFeed, error) {
	nodes, err := fcfq.Limit(2).All(setContextOp(ctx, fcfq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{fiatcurrencyfeed.Label}
	default:
		return nil, &NotSingularError{fiatcurrencyfeed.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fcfq *FiatCurrencyFeedQuery) OnlyX(ctx context.Context) *FiatCurrencyFeed {
	node, err := fcfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FiatCurrencyFeed ID in the query.
// Returns a *NotSingularError when more than one FiatCurrencyFeed ID is found.
// Returns a *NotFoundError when no entities are found.
func (fcfq *FiatCurrencyFeedQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = fcfq.Limit(2).IDs(setContextOp(ctx, fcfq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{fiatcurrencyfeed.Label}
	default:
		err = &NotSingularError{fiatcurrencyfeed.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fcfq *FiatCurrencyFeedQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := fcfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FiatCurrencyFeeds.
func (fcfq *FiatCurrencyFeedQuery) All(ctx context.Context) ([]*FiatCurrencyFeed, error) {
	ctx = setContextOp(ctx, fcfq.ctx, ent.OpQueryAll)
	if err := fcfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*FiatCurrencyFeed, *FiatCurrencyFeedQuery]()
	return withInterceptors[[]*FiatCurrencyFeed](ctx, fcfq, qr, fcfq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fcfq *FiatCurrencyFeedQuery) AllX(ctx context.Context) []*FiatCurrencyFeed {
	nodes, err := fcfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FiatCurrencyFeed IDs.
func (fcfq *FiatCurrencyFeedQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if fcfq.ctx.Unique == nil && fcfq.path != nil {
		fcfq.Unique(true)
	}
	ctx = setContextOp(ctx, fcfq.ctx, ent.OpQueryIDs)
	if err = fcfq.Select(fiatcurrencyfeed.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fcfq *FiatCurrencyFeedQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := fcfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fcfq *FiatCurrencyFeedQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fcfq.ctx, ent.OpQueryCount)
	if err := fcfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fcfq, querierCount[*FiatCurrencyFeedQuery](), fcfq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fcfq *FiatCurrencyFeedQuery) CountX(ctx context.Context) int {
	count, err := fcfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fcfq *FiatCurrencyFeedQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fcfq.ctx, ent.OpQueryExist)
	switch _, err := fcfq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fcfq *FiatCurrencyFeedQuery) ExistX(ctx context.Context) bool {
	exist, err := fcfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FiatCurrencyFeedQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fcfq *FiatCurrencyFeedQuery) Clone() *FiatCurrencyFeedQuery {
	if fcfq == nil {
		return nil
	}
	return &FiatCurrencyFeedQuery{
		config:     fcfq.config,
		ctx:        fcfq.ctx.Clone(),
		order:      append([]fiatcurrencyfeed.OrderOption{}, fcfq.order...),
		inters:     append([]Interceptor{}, fcfq.inters...),
		predicates: append([]predicate.FiatCurrencyFeed{}, fcfq.predicates...),
		// clone intermediate query.
		sql:       fcfq.sql.Clone(),
		path:      fcfq.path,
		modifiers: append([]func(*sql.Selector){}, fcfq.modifiers...),
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
//	client.FiatCurrencyFeed.Query().
//		GroupBy(fiatcurrencyfeed.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (fcfq *FiatCurrencyFeedQuery) GroupBy(field string, fields ...string) *FiatCurrencyFeedGroupBy {
	fcfq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FiatCurrencyFeedGroupBy{build: fcfq}
	grbuild.flds = &fcfq.ctx.Fields
	grbuild.label = fiatcurrencyfeed.Label
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
//	client.FiatCurrencyFeed.Query().
//		Select(fiatcurrencyfeed.FieldCreatedAt).
//		Scan(ctx, &v)
func (fcfq *FiatCurrencyFeedQuery) Select(fields ...string) *FiatCurrencyFeedSelect {
	fcfq.ctx.Fields = append(fcfq.ctx.Fields, fields...)
	sbuild := &FiatCurrencyFeedSelect{FiatCurrencyFeedQuery: fcfq}
	sbuild.label = fiatcurrencyfeed.Label
	sbuild.flds, sbuild.scan = &fcfq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FiatCurrencyFeedSelect configured with the given aggregations.
func (fcfq *FiatCurrencyFeedQuery) Aggregate(fns ...AggregateFunc) *FiatCurrencyFeedSelect {
	return fcfq.Select().Aggregate(fns...)
}

func (fcfq *FiatCurrencyFeedQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fcfq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fcfq); err != nil {
				return err
			}
		}
	}
	for _, f := range fcfq.ctx.Fields {
		if !fiatcurrencyfeed.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if fcfq.path != nil {
		prev, err := fcfq.path(ctx)
		if err != nil {
			return err
		}
		fcfq.sql = prev
	}
	return nil
}

func (fcfq *FiatCurrencyFeedQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FiatCurrencyFeed, error) {
	var (
		nodes = []*FiatCurrencyFeed{}
		_spec = fcfq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*FiatCurrencyFeed).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &FiatCurrencyFeed{config: fcfq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(fcfq.modifiers) > 0 {
		_spec.Modifiers = fcfq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fcfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (fcfq *FiatCurrencyFeedQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fcfq.querySpec()
	if len(fcfq.modifiers) > 0 {
		_spec.Modifiers = fcfq.modifiers
	}
	_spec.Node.Columns = fcfq.ctx.Fields
	if len(fcfq.ctx.Fields) > 0 {
		_spec.Unique = fcfq.ctx.Unique != nil && *fcfq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fcfq.driver, _spec)
}

func (fcfq *FiatCurrencyFeedQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(fiatcurrencyfeed.Table, fiatcurrencyfeed.Columns, sqlgraph.NewFieldSpec(fiatcurrencyfeed.FieldID, field.TypeUint32))
	_spec.From = fcfq.sql
	if unique := fcfq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fcfq.path != nil {
		_spec.Unique = true
	}
	if fields := fcfq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fiatcurrencyfeed.FieldID)
		for i := range fields {
			if fields[i] != fiatcurrencyfeed.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fcfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fcfq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fcfq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fcfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fcfq *FiatCurrencyFeedQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fcfq.driver.Dialect())
	t1 := builder.Table(fiatcurrencyfeed.Table)
	columns := fcfq.ctx.Fields
	if len(columns) == 0 {
		columns = fiatcurrencyfeed.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fcfq.sql != nil {
		selector = fcfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fcfq.ctx.Unique != nil && *fcfq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range fcfq.modifiers {
		m(selector)
	}
	for _, p := range fcfq.predicates {
		p(selector)
	}
	for _, p := range fcfq.order {
		p(selector)
	}
	if offset := fcfq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fcfq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (fcfq *FiatCurrencyFeedQuery) ForUpdate(opts ...sql.LockOption) *FiatCurrencyFeedQuery {
	if fcfq.driver.Dialect() == dialect.Postgres {
		fcfq.Unique(false)
	}
	fcfq.modifiers = append(fcfq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return fcfq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (fcfq *FiatCurrencyFeedQuery) ForShare(opts ...sql.LockOption) *FiatCurrencyFeedQuery {
	if fcfq.driver.Dialect() == dialect.Postgres {
		fcfq.Unique(false)
	}
	fcfq.modifiers = append(fcfq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return fcfq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (fcfq *FiatCurrencyFeedQuery) Modify(modifiers ...func(s *sql.Selector)) *FiatCurrencyFeedSelect {
	fcfq.modifiers = append(fcfq.modifiers, modifiers...)
	return fcfq.Select()
}

// FiatCurrencyFeedGroupBy is the group-by builder for FiatCurrencyFeed entities.
type FiatCurrencyFeedGroupBy struct {
	selector
	build *FiatCurrencyFeedQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fcfgb *FiatCurrencyFeedGroupBy) Aggregate(fns ...AggregateFunc) *FiatCurrencyFeedGroupBy {
	fcfgb.fns = append(fcfgb.fns, fns...)
	return fcfgb
}

// Scan applies the selector query and scans the result into the given value.
func (fcfgb *FiatCurrencyFeedGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fcfgb.build.ctx, ent.OpQueryGroupBy)
	if err := fcfgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FiatCurrencyFeedQuery, *FiatCurrencyFeedGroupBy](ctx, fcfgb.build, fcfgb, fcfgb.build.inters, v)
}

func (fcfgb *FiatCurrencyFeedGroupBy) sqlScan(ctx context.Context, root *FiatCurrencyFeedQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fcfgb.fns))
	for _, fn := range fcfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fcfgb.flds)+len(fcfgb.fns))
		for _, f := range *fcfgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fcfgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fcfgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FiatCurrencyFeedSelect is the builder for selecting fields of FiatCurrencyFeed entities.
type FiatCurrencyFeedSelect struct {
	*FiatCurrencyFeedQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fcfs *FiatCurrencyFeedSelect) Aggregate(fns ...AggregateFunc) *FiatCurrencyFeedSelect {
	fcfs.fns = append(fcfs.fns, fns...)
	return fcfs
}

// Scan applies the selector query and scans the result into the given value.
func (fcfs *FiatCurrencyFeedSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fcfs.ctx, ent.OpQuerySelect)
	if err := fcfs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FiatCurrencyFeedQuery, *FiatCurrencyFeedSelect](ctx, fcfs.FiatCurrencyFeedQuery, fcfs, fcfs.inters, v)
}

func (fcfs *FiatCurrencyFeedSelect) sqlScan(ctx context.Context, root *FiatCurrencyFeedQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fcfs.fns))
	for _, fn := range fcfs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fcfs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fcfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (fcfs *FiatCurrencyFeedSelect) Modify(modifiers ...func(s *sql.Selector)) *FiatCurrencyFeedSelect {
	fcfs.modifiers = append(fcfs.modifiers, modifiers...)
	return fcfs
}
