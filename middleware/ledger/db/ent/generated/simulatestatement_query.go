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
	"github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/predicate"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/simulatestatement"
)

// SimulateStatementQuery is the builder for querying SimulateStatement entities.
type SimulateStatementQuery struct {
	config
	ctx        *QueryContext
	order      []simulatestatement.OrderOption
	inters     []Interceptor
	predicates []predicate.SimulateStatement
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SimulateStatementQuery builder.
func (ssq *SimulateStatementQuery) Where(ps ...predicate.SimulateStatement) *SimulateStatementQuery {
	ssq.predicates = append(ssq.predicates, ps...)
	return ssq
}

// Limit the number of records to be returned by this query.
func (ssq *SimulateStatementQuery) Limit(limit int) *SimulateStatementQuery {
	ssq.ctx.Limit = &limit
	return ssq
}

// Offset to start from.
func (ssq *SimulateStatementQuery) Offset(offset int) *SimulateStatementQuery {
	ssq.ctx.Offset = &offset
	return ssq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ssq *SimulateStatementQuery) Unique(unique bool) *SimulateStatementQuery {
	ssq.ctx.Unique = &unique
	return ssq
}

// Order specifies how the records should be ordered.
func (ssq *SimulateStatementQuery) Order(o ...simulatestatement.OrderOption) *SimulateStatementQuery {
	ssq.order = append(ssq.order, o...)
	return ssq
}

// First returns the first SimulateStatement entity from the query.
// Returns a *NotFoundError when no SimulateStatement was found.
func (ssq *SimulateStatementQuery) First(ctx context.Context) (*SimulateStatement, error) {
	nodes, err := ssq.Limit(1).All(setContextOp(ctx, ssq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{simulatestatement.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ssq *SimulateStatementQuery) FirstX(ctx context.Context) *SimulateStatement {
	node, err := ssq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SimulateStatement ID from the query.
// Returns a *NotFoundError when no SimulateStatement ID was found.
func (ssq *SimulateStatementQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = ssq.Limit(1).IDs(setContextOp(ctx, ssq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{simulatestatement.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ssq *SimulateStatementQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := ssq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SimulateStatement entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SimulateStatement entity is found.
// Returns a *NotFoundError when no SimulateStatement entities are found.
func (ssq *SimulateStatementQuery) Only(ctx context.Context) (*SimulateStatement, error) {
	nodes, err := ssq.Limit(2).All(setContextOp(ctx, ssq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{simulatestatement.Label}
	default:
		return nil, &NotSingularError{simulatestatement.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ssq *SimulateStatementQuery) OnlyX(ctx context.Context) *SimulateStatement {
	node, err := ssq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SimulateStatement ID in the query.
// Returns a *NotSingularError when more than one SimulateStatement ID is found.
// Returns a *NotFoundError when no entities are found.
func (ssq *SimulateStatementQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = ssq.Limit(2).IDs(setContextOp(ctx, ssq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{simulatestatement.Label}
	default:
		err = &NotSingularError{simulatestatement.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ssq *SimulateStatementQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := ssq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SimulateStatements.
func (ssq *SimulateStatementQuery) All(ctx context.Context) ([]*SimulateStatement, error) {
	ctx = setContextOp(ctx, ssq.ctx, ent.OpQueryAll)
	if err := ssq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SimulateStatement, *SimulateStatementQuery]()
	return withInterceptors[[]*SimulateStatement](ctx, ssq, qr, ssq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ssq *SimulateStatementQuery) AllX(ctx context.Context) []*SimulateStatement {
	nodes, err := ssq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SimulateStatement IDs.
func (ssq *SimulateStatementQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if ssq.ctx.Unique == nil && ssq.path != nil {
		ssq.Unique(true)
	}
	ctx = setContextOp(ctx, ssq.ctx, ent.OpQueryIDs)
	if err = ssq.Select(simulatestatement.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ssq *SimulateStatementQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := ssq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ssq *SimulateStatementQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ssq.ctx, ent.OpQueryCount)
	if err := ssq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ssq, querierCount[*SimulateStatementQuery](), ssq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ssq *SimulateStatementQuery) CountX(ctx context.Context) int {
	count, err := ssq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ssq *SimulateStatementQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ssq.ctx, ent.OpQueryExist)
	switch _, err := ssq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ssq *SimulateStatementQuery) ExistX(ctx context.Context) bool {
	exist, err := ssq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SimulateStatementQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ssq *SimulateStatementQuery) Clone() *SimulateStatementQuery {
	if ssq == nil {
		return nil
	}
	return &SimulateStatementQuery{
		config:     ssq.config,
		ctx:        ssq.ctx.Clone(),
		order:      append([]simulatestatement.OrderOption{}, ssq.order...),
		inters:     append([]Interceptor{}, ssq.inters...),
		predicates: append([]predicate.SimulateStatement{}, ssq.predicates...),
		// clone intermediate query.
		sql:       ssq.sql.Clone(),
		path:      ssq.path,
		modifiers: append([]func(*sql.Selector){}, ssq.modifiers...),
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
//	client.SimulateStatement.Query().
//		GroupBy(simulatestatement.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (ssq *SimulateStatementQuery) GroupBy(field string, fields ...string) *SimulateStatementGroupBy {
	ssq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SimulateStatementGroupBy{build: ssq}
	grbuild.flds = &ssq.ctx.Fields
	grbuild.label = simulatestatement.Label
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
//	client.SimulateStatement.Query().
//		Select(simulatestatement.FieldCreatedAt).
//		Scan(ctx, &v)
func (ssq *SimulateStatementQuery) Select(fields ...string) *SimulateStatementSelect {
	ssq.ctx.Fields = append(ssq.ctx.Fields, fields...)
	sbuild := &SimulateStatementSelect{SimulateStatementQuery: ssq}
	sbuild.label = simulatestatement.Label
	sbuild.flds, sbuild.scan = &ssq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SimulateStatementSelect configured with the given aggregations.
func (ssq *SimulateStatementQuery) Aggregate(fns ...AggregateFunc) *SimulateStatementSelect {
	return ssq.Select().Aggregate(fns...)
}

func (ssq *SimulateStatementQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ssq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ssq); err != nil {
				return err
			}
		}
	}
	for _, f := range ssq.ctx.Fields {
		if !simulatestatement.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if ssq.path != nil {
		prev, err := ssq.path(ctx)
		if err != nil {
			return err
		}
		ssq.sql = prev
	}
	return nil
}

func (ssq *SimulateStatementQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SimulateStatement, error) {
	var (
		nodes = []*SimulateStatement{}
		_spec = ssq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SimulateStatement).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SimulateStatement{config: ssq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(ssq.modifiers) > 0 {
		_spec.Modifiers = ssq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ssq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ssq *SimulateStatementQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ssq.querySpec()
	if len(ssq.modifiers) > 0 {
		_spec.Modifiers = ssq.modifiers
	}
	_spec.Node.Columns = ssq.ctx.Fields
	if len(ssq.ctx.Fields) > 0 {
		_spec.Unique = ssq.ctx.Unique != nil && *ssq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ssq.driver, _spec)
}

func (ssq *SimulateStatementQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(simulatestatement.Table, simulatestatement.Columns, sqlgraph.NewFieldSpec(simulatestatement.FieldID, field.TypeUint32))
	_spec.From = ssq.sql
	if unique := ssq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ssq.path != nil {
		_spec.Unique = true
	}
	if fields := ssq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, simulatestatement.FieldID)
		for i := range fields {
			if fields[i] != simulatestatement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ssq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ssq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ssq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ssq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ssq *SimulateStatementQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ssq.driver.Dialect())
	t1 := builder.Table(simulatestatement.Table)
	columns := ssq.ctx.Fields
	if len(columns) == 0 {
		columns = simulatestatement.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ssq.sql != nil {
		selector = ssq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ssq.ctx.Unique != nil && *ssq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range ssq.modifiers {
		m(selector)
	}
	for _, p := range ssq.predicates {
		p(selector)
	}
	for _, p := range ssq.order {
		p(selector)
	}
	if offset := ssq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ssq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (ssq *SimulateStatementQuery) ForUpdate(opts ...sql.LockOption) *SimulateStatementQuery {
	if ssq.driver.Dialect() == dialect.Postgres {
		ssq.Unique(false)
	}
	ssq.modifiers = append(ssq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return ssq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (ssq *SimulateStatementQuery) ForShare(opts ...sql.LockOption) *SimulateStatementQuery {
	if ssq.driver.Dialect() == dialect.Postgres {
		ssq.Unique(false)
	}
	ssq.modifiers = append(ssq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return ssq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ssq *SimulateStatementQuery) Modify(modifiers ...func(s *sql.Selector)) *SimulateStatementSelect {
	ssq.modifiers = append(ssq.modifiers, modifiers...)
	return ssq.Select()
}

// SimulateStatementGroupBy is the group-by builder for SimulateStatement entities.
type SimulateStatementGroupBy struct {
	selector
	build *SimulateStatementQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ssgb *SimulateStatementGroupBy) Aggregate(fns ...AggregateFunc) *SimulateStatementGroupBy {
	ssgb.fns = append(ssgb.fns, fns...)
	return ssgb
}

// Scan applies the selector query and scans the result into the given value.
func (ssgb *SimulateStatementGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ssgb.build.ctx, ent.OpQueryGroupBy)
	if err := ssgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SimulateStatementQuery, *SimulateStatementGroupBy](ctx, ssgb.build, ssgb, ssgb.build.inters, v)
}

func (ssgb *SimulateStatementGroupBy) sqlScan(ctx context.Context, root *SimulateStatementQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ssgb.fns))
	for _, fn := range ssgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ssgb.flds)+len(ssgb.fns))
		for _, f := range *ssgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ssgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ssgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SimulateStatementSelect is the builder for selecting fields of SimulateStatement entities.
type SimulateStatementSelect struct {
	*SimulateStatementQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sss *SimulateStatementSelect) Aggregate(fns ...AggregateFunc) *SimulateStatementSelect {
	sss.fns = append(sss.fns, fns...)
	return sss
}

// Scan applies the selector query and scans the result into the given value.
func (sss *SimulateStatementSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sss.ctx, ent.OpQuerySelect)
	if err := sss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SimulateStatementQuery, *SimulateStatementSelect](ctx, sss.SimulateStatementQuery, sss, sss.inters, v)
}

func (sss *SimulateStatementSelect) sqlScan(ctx context.Context, root *SimulateStatementQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sss.fns))
	for _, fn := range sss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sss *SimulateStatementSelect) Modify(modifiers ...func(s *sql.Selector)) *SimulateStatementSelect {
	sss.modifiers = append(sss.modifiers, modifiers...)
	return sss
}
