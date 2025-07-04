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
	"github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/goodstatement"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated/predicate"
)

// GoodStatementQuery is the builder for querying GoodStatement entities.
type GoodStatementQuery struct {
	config
	ctx        *QueryContext
	order      []goodstatement.OrderOption
	inters     []Interceptor
	predicates []predicate.GoodStatement
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GoodStatementQuery builder.
func (gsq *GoodStatementQuery) Where(ps ...predicate.GoodStatement) *GoodStatementQuery {
	gsq.predicates = append(gsq.predicates, ps...)
	return gsq
}

// Limit the number of records to be returned by this query.
func (gsq *GoodStatementQuery) Limit(limit int) *GoodStatementQuery {
	gsq.ctx.Limit = &limit
	return gsq
}

// Offset to start from.
func (gsq *GoodStatementQuery) Offset(offset int) *GoodStatementQuery {
	gsq.ctx.Offset = &offset
	return gsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gsq *GoodStatementQuery) Unique(unique bool) *GoodStatementQuery {
	gsq.ctx.Unique = &unique
	return gsq
}

// Order specifies how the records should be ordered.
func (gsq *GoodStatementQuery) Order(o ...goodstatement.OrderOption) *GoodStatementQuery {
	gsq.order = append(gsq.order, o...)
	return gsq
}

// First returns the first GoodStatement entity from the query.
// Returns a *NotFoundError when no GoodStatement was found.
func (gsq *GoodStatementQuery) First(ctx context.Context) (*GoodStatement, error) {
	nodes, err := gsq.Limit(1).All(setContextOp(ctx, gsq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{goodstatement.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gsq *GoodStatementQuery) FirstX(ctx context.Context) *GoodStatement {
	node, err := gsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GoodStatement ID from the query.
// Returns a *NotFoundError when no GoodStatement ID was found.
func (gsq *GoodStatementQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = gsq.Limit(1).IDs(setContextOp(ctx, gsq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{goodstatement.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gsq *GoodStatementQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := gsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GoodStatement entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GoodStatement entity is found.
// Returns a *NotFoundError when no GoodStatement entities are found.
func (gsq *GoodStatementQuery) Only(ctx context.Context) (*GoodStatement, error) {
	nodes, err := gsq.Limit(2).All(setContextOp(ctx, gsq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{goodstatement.Label}
	default:
		return nil, &NotSingularError{goodstatement.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gsq *GoodStatementQuery) OnlyX(ctx context.Context) *GoodStatement {
	node, err := gsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GoodStatement ID in the query.
// Returns a *NotSingularError when more than one GoodStatement ID is found.
// Returns a *NotFoundError when no entities are found.
func (gsq *GoodStatementQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = gsq.Limit(2).IDs(setContextOp(ctx, gsq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{goodstatement.Label}
	default:
		err = &NotSingularError{goodstatement.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gsq *GoodStatementQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := gsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GoodStatements.
func (gsq *GoodStatementQuery) All(ctx context.Context) ([]*GoodStatement, error) {
	ctx = setContextOp(ctx, gsq.ctx, ent.OpQueryAll)
	if err := gsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GoodStatement, *GoodStatementQuery]()
	return withInterceptors[[]*GoodStatement](ctx, gsq, qr, gsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gsq *GoodStatementQuery) AllX(ctx context.Context) []*GoodStatement {
	nodes, err := gsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GoodStatement IDs.
func (gsq *GoodStatementQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if gsq.ctx.Unique == nil && gsq.path != nil {
		gsq.Unique(true)
	}
	ctx = setContextOp(ctx, gsq.ctx, ent.OpQueryIDs)
	if err = gsq.Select(goodstatement.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gsq *GoodStatementQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := gsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gsq *GoodStatementQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, gsq.ctx, ent.OpQueryCount)
	if err := gsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gsq, querierCount[*GoodStatementQuery](), gsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gsq *GoodStatementQuery) CountX(ctx context.Context) int {
	count, err := gsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gsq *GoodStatementQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, gsq.ctx, ent.OpQueryExist)
	switch _, err := gsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gsq *GoodStatementQuery) ExistX(ctx context.Context) bool {
	exist, err := gsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GoodStatementQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gsq *GoodStatementQuery) Clone() *GoodStatementQuery {
	if gsq == nil {
		return nil
	}
	return &GoodStatementQuery{
		config:     gsq.config,
		ctx:        gsq.ctx.Clone(),
		order:      append([]goodstatement.OrderOption{}, gsq.order...),
		inters:     append([]Interceptor{}, gsq.inters...),
		predicates: append([]predicate.GoodStatement{}, gsq.predicates...),
		// clone intermediate query.
		sql:       gsq.sql.Clone(),
		path:      gsq.path,
		modifiers: append([]func(*sql.Selector){}, gsq.modifiers...),
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
//	client.GoodStatement.Query().
//		GroupBy(goodstatement.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (gsq *GoodStatementQuery) GroupBy(field string, fields ...string) *GoodStatementGroupBy {
	gsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GoodStatementGroupBy{build: gsq}
	grbuild.flds = &gsq.ctx.Fields
	grbuild.label = goodstatement.Label
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
//	client.GoodStatement.Query().
//		Select(goodstatement.FieldCreatedAt).
//		Scan(ctx, &v)
func (gsq *GoodStatementQuery) Select(fields ...string) *GoodStatementSelect {
	gsq.ctx.Fields = append(gsq.ctx.Fields, fields...)
	sbuild := &GoodStatementSelect{GoodStatementQuery: gsq}
	sbuild.label = goodstatement.Label
	sbuild.flds, sbuild.scan = &gsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GoodStatementSelect configured with the given aggregations.
func (gsq *GoodStatementQuery) Aggregate(fns ...AggregateFunc) *GoodStatementSelect {
	return gsq.Select().Aggregate(fns...)
}

func (gsq *GoodStatementQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gsq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gsq); err != nil {
				return err
			}
		}
	}
	for _, f := range gsq.ctx.Fields {
		if !goodstatement.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if gsq.path != nil {
		prev, err := gsq.path(ctx)
		if err != nil {
			return err
		}
		gsq.sql = prev
	}
	return nil
}

func (gsq *GoodStatementQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GoodStatement, error) {
	var (
		nodes = []*GoodStatement{}
		_spec = gsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GoodStatement).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GoodStatement{config: gsq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(gsq.modifiers) > 0 {
		_spec.Modifiers = gsq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (gsq *GoodStatementQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gsq.querySpec()
	if len(gsq.modifiers) > 0 {
		_spec.Modifiers = gsq.modifiers
	}
	_spec.Node.Columns = gsq.ctx.Fields
	if len(gsq.ctx.Fields) > 0 {
		_spec.Unique = gsq.ctx.Unique != nil && *gsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, gsq.driver, _spec)
}

func (gsq *GoodStatementQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(goodstatement.Table, goodstatement.Columns, sqlgraph.NewFieldSpec(goodstatement.FieldID, field.TypeUint32))
	_spec.From = gsq.sql
	if unique := gsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if gsq.path != nil {
		_spec.Unique = true
	}
	if fields := gsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodstatement.FieldID)
		for i := range fields {
			if fields[i] != goodstatement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gsq *GoodStatementQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gsq.driver.Dialect())
	t1 := builder.Table(goodstatement.Table)
	columns := gsq.ctx.Fields
	if len(columns) == 0 {
		columns = goodstatement.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gsq.sql != nil {
		selector = gsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gsq.ctx.Unique != nil && *gsq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range gsq.modifiers {
		m(selector)
	}
	for _, p := range gsq.predicates {
		p(selector)
	}
	for _, p := range gsq.order {
		p(selector)
	}
	if offset := gsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (gsq *GoodStatementQuery) ForUpdate(opts ...sql.LockOption) *GoodStatementQuery {
	if gsq.driver.Dialect() == dialect.Postgres {
		gsq.Unique(false)
	}
	gsq.modifiers = append(gsq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return gsq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (gsq *GoodStatementQuery) ForShare(opts ...sql.LockOption) *GoodStatementQuery {
	if gsq.driver.Dialect() == dialect.Postgres {
		gsq.Unique(false)
	}
	gsq.modifiers = append(gsq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return gsq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gsq *GoodStatementQuery) Modify(modifiers ...func(s *sql.Selector)) *GoodStatementSelect {
	gsq.modifiers = append(gsq.modifiers, modifiers...)
	return gsq.Select()
}

// GoodStatementGroupBy is the group-by builder for GoodStatement entities.
type GoodStatementGroupBy struct {
	selector
	build *GoodStatementQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gsgb *GoodStatementGroupBy) Aggregate(fns ...AggregateFunc) *GoodStatementGroupBy {
	gsgb.fns = append(gsgb.fns, fns...)
	return gsgb
}

// Scan applies the selector query and scans the result into the given value.
func (gsgb *GoodStatementGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gsgb.build.ctx, ent.OpQueryGroupBy)
	if err := gsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GoodStatementQuery, *GoodStatementGroupBy](ctx, gsgb.build, gsgb, gsgb.build.inters, v)
}

func (gsgb *GoodStatementGroupBy) sqlScan(ctx context.Context, root *GoodStatementQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(gsgb.fns))
	for _, fn := range gsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*gsgb.flds)+len(gsgb.fns))
		for _, f := range *gsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*gsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GoodStatementSelect is the builder for selecting fields of GoodStatement entities.
type GoodStatementSelect struct {
	*GoodStatementQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gss *GoodStatementSelect) Aggregate(fns ...AggregateFunc) *GoodStatementSelect {
	gss.fns = append(gss.fns, fns...)
	return gss
}

// Scan applies the selector query and scans the result into the given value.
func (gss *GoodStatementSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gss.ctx, ent.OpQuerySelect)
	if err := gss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GoodStatementQuery, *GoodStatementSelect](ctx, gss.GoodStatementQuery, gss, gss.inters, v)
}

func (gss *GoodStatementSelect) sqlScan(ctx context.Context, root *GoodStatementQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gss.fns))
	for _, fn := range gss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gss *GoodStatementSelect) Modify(modifiers ...func(s *sql.Selector)) *GoodStatementSelect {
	gss.modifiers = append(gss.modifiers, modifiers...)
	return gss
}
