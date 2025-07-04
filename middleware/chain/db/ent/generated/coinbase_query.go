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
	"github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/coinbase"
	"github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/predicate"
)

// CoinBaseQuery is the builder for querying CoinBase entities.
type CoinBaseQuery struct {
	config
	ctx        *QueryContext
	order      []coinbase.OrderOption
	inters     []Interceptor
	predicates []predicate.CoinBase
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CoinBaseQuery builder.
func (cbq *CoinBaseQuery) Where(ps ...predicate.CoinBase) *CoinBaseQuery {
	cbq.predicates = append(cbq.predicates, ps...)
	return cbq
}

// Limit the number of records to be returned by this query.
func (cbq *CoinBaseQuery) Limit(limit int) *CoinBaseQuery {
	cbq.ctx.Limit = &limit
	return cbq
}

// Offset to start from.
func (cbq *CoinBaseQuery) Offset(offset int) *CoinBaseQuery {
	cbq.ctx.Offset = &offset
	return cbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cbq *CoinBaseQuery) Unique(unique bool) *CoinBaseQuery {
	cbq.ctx.Unique = &unique
	return cbq
}

// Order specifies how the records should be ordered.
func (cbq *CoinBaseQuery) Order(o ...coinbase.OrderOption) *CoinBaseQuery {
	cbq.order = append(cbq.order, o...)
	return cbq
}

// First returns the first CoinBase entity from the query.
// Returns a *NotFoundError when no CoinBase was found.
func (cbq *CoinBaseQuery) First(ctx context.Context) (*CoinBase, error) {
	nodes, err := cbq.Limit(1).All(setContextOp(ctx, cbq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{coinbase.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cbq *CoinBaseQuery) FirstX(ctx context.Context) *CoinBase {
	node, err := cbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CoinBase ID from the query.
// Returns a *NotFoundError when no CoinBase ID was found.
func (cbq *CoinBaseQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = cbq.Limit(1).IDs(setContextOp(ctx, cbq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{coinbase.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cbq *CoinBaseQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := cbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CoinBase entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CoinBase entity is found.
// Returns a *NotFoundError when no CoinBase entities are found.
func (cbq *CoinBaseQuery) Only(ctx context.Context) (*CoinBase, error) {
	nodes, err := cbq.Limit(2).All(setContextOp(ctx, cbq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{coinbase.Label}
	default:
		return nil, &NotSingularError{coinbase.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cbq *CoinBaseQuery) OnlyX(ctx context.Context) *CoinBase {
	node, err := cbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CoinBase ID in the query.
// Returns a *NotSingularError when more than one CoinBase ID is found.
// Returns a *NotFoundError when no entities are found.
func (cbq *CoinBaseQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = cbq.Limit(2).IDs(setContextOp(ctx, cbq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{coinbase.Label}
	default:
		err = &NotSingularError{coinbase.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cbq *CoinBaseQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := cbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CoinBases.
func (cbq *CoinBaseQuery) All(ctx context.Context) ([]*CoinBase, error) {
	ctx = setContextOp(ctx, cbq.ctx, ent.OpQueryAll)
	if err := cbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CoinBase, *CoinBaseQuery]()
	return withInterceptors[[]*CoinBase](ctx, cbq, qr, cbq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cbq *CoinBaseQuery) AllX(ctx context.Context) []*CoinBase {
	nodes, err := cbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CoinBase IDs.
func (cbq *CoinBaseQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if cbq.ctx.Unique == nil && cbq.path != nil {
		cbq.Unique(true)
	}
	ctx = setContextOp(ctx, cbq.ctx, ent.OpQueryIDs)
	if err = cbq.Select(coinbase.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cbq *CoinBaseQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := cbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cbq *CoinBaseQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cbq.ctx, ent.OpQueryCount)
	if err := cbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cbq, querierCount[*CoinBaseQuery](), cbq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cbq *CoinBaseQuery) CountX(ctx context.Context) int {
	count, err := cbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cbq *CoinBaseQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cbq.ctx, ent.OpQueryExist)
	switch _, err := cbq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cbq *CoinBaseQuery) ExistX(ctx context.Context) bool {
	exist, err := cbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CoinBaseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cbq *CoinBaseQuery) Clone() *CoinBaseQuery {
	if cbq == nil {
		return nil
	}
	return &CoinBaseQuery{
		config:     cbq.config,
		ctx:        cbq.ctx.Clone(),
		order:      append([]coinbase.OrderOption{}, cbq.order...),
		inters:     append([]Interceptor{}, cbq.inters...),
		predicates: append([]predicate.CoinBase{}, cbq.predicates...),
		// clone intermediate query.
		sql:       cbq.sql.Clone(),
		path:      cbq.path,
		modifiers: append([]func(*sql.Selector){}, cbq.modifiers...),
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
//	client.CoinBase.Query().
//		GroupBy(coinbase.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (cbq *CoinBaseQuery) GroupBy(field string, fields ...string) *CoinBaseGroupBy {
	cbq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CoinBaseGroupBy{build: cbq}
	grbuild.flds = &cbq.ctx.Fields
	grbuild.label = coinbase.Label
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
//	client.CoinBase.Query().
//		Select(coinbase.FieldCreatedAt).
//		Scan(ctx, &v)
func (cbq *CoinBaseQuery) Select(fields ...string) *CoinBaseSelect {
	cbq.ctx.Fields = append(cbq.ctx.Fields, fields...)
	sbuild := &CoinBaseSelect{CoinBaseQuery: cbq}
	sbuild.label = coinbase.Label
	sbuild.flds, sbuild.scan = &cbq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CoinBaseSelect configured with the given aggregations.
func (cbq *CoinBaseQuery) Aggregate(fns ...AggregateFunc) *CoinBaseSelect {
	return cbq.Select().Aggregate(fns...)
}

func (cbq *CoinBaseQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cbq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cbq); err != nil {
				return err
			}
		}
	}
	for _, f := range cbq.ctx.Fields {
		if !coinbase.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if cbq.path != nil {
		prev, err := cbq.path(ctx)
		if err != nil {
			return err
		}
		cbq.sql = prev
	}
	return nil
}

func (cbq *CoinBaseQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CoinBase, error) {
	var (
		nodes = []*CoinBase{}
		_spec = cbq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CoinBase).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CoinBase{config: cbq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(cbq.modifiers) > 0 {
		_spec.Modifiers = cbq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (cbq *CoinBaseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cbq.querySpec()
	if len(cbq.modifiers) > 0 {
		_spec.Modifiers = cbq.modifiers
	}
	_spec.Node.Columns = cbq.ctx.Fields
	if len(cbq.ctx.Fields) > 0 {
		_spec.Unique = cbq.ctx.Unique != nil && *cbq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cbq.driver, _spec)
}

func (cbq *CoinBaseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(coinbase.Table, coinbase.Columns, sqlgraph.NewFieldSpec(coinbase.FieldID, field.TypeUint32))
	_spec.From = cbq.sql
	if unique := cbq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cbq.path != nil {
		_spec.Unique = true
	}
	if fields := cbq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coinbase.FieldID)
		for i := range fields {
			if fields[i] != coinbase.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cbq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cbq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cbq *CoinBaseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cbq.driver.Dialect())
	t1 := builder.Table(coinbase.Table)
	columns := cbq.ctx.Fields
	if len(columns) == 0 {
		columns = coinbase.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cbq.sql != nil {
		selector = cbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cbq.ctx.Unique != nil && *cbq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range cbq.modifiers {
		m(selector)
	}
	for _, p := range cbq.predicates {
		p(selector)
	}
	for _, p := range cbq.order {
		p(selector)
	}
	if offset := cbq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cbq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (cbq *CoinBaseQuery) ForUpdate(opts ...sql.LockOption) *CoinBaseQuery {
	if cbq.driver.Dialect() == dialect.Postgres {
		cbq.Unique(false)
	}
	cbq.modifiers = append(cbq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return cbq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (cbq *CoinBaseQuery) ForShare(opts ...sql.LockOption) *CoinBaseQuery {
	if cbq.driver.Dialect() == dialect.Postgres {
		cbq.Unique(false)
	}
	cbq.modifiers = append(cbq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return cbq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cbq *CoinBaseQuery) Modify(modifiers ...func(s *sql.Selector)) *CoinBaseSelect {
	cbq.modifiers = append(cbq.modifiers, modifiers...)
	return cbq.Select()
}

// CoinBaseGroupBy is the group-by builder for CoinBase entities.
type CoinBaseGroupBy struct {
	selector
	build *CoinBaseQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cbgb *CoinBaseGroupBy) Aggregate(fns ...AggregateFunc) *CoinBaseGroupBy {
	cbgb.fns = append(cbgb.fns, fns...)
	return cbgb
}

// Scan applies the selector query and scans the result into the given value.
func (cbgb *CoinBaseGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cbgb.build.ctx, ent.OpQueryGroupBy)
	if err := cbgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CoinBaseQuery, *CoinBaseGroupBy](ctx, cbgb.build, cbgb, cbgb.build.inters, v)
}

func (cbgb *CoinBaseGroupBy) sqlScan(ctx context.Context, root *CoinBaseQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cbgb.fns))
	for _, fn := range cbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cbgb.flds)+len(cbgb.fns))
		for _, f := range *cbgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cbgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cbgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CoinBaseSelect is the builder for selecting fields of CoinBase entities.
type CoinBaseSelect struct {
	*CoinBaseQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cbs *CoinBaseSelect) Aggregate(fns ...AggregateFunc) *CoinBaseSelect {
	cbs.fns = append(cbs.fns, fns...)
	return cbs
}

// Scan applies the selector query and scans the result into the given value.
func (cbs *CoinBaseSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cbs.ctx, ent.OpQuerySelect)
	if err := cbs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CoinBaseQuery, *CoinBaseSelect](ctx, cbs.CoinBaseQuery, cbs, cbs.inters, v)
}

func (cbs *CoinBaseSelect) sqlScan(ctx context.Context, root *CoinBaseQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cbs.fns))
	for _, fn := range cbs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cbs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cbs *CoinBaseSelect) Modify(modifiers ...func(s *sql.Selector)) *CoinBaseSelect {
	cbs.modifiers = append(cbs.modifiers, modifiers...)
	return cbs
}
