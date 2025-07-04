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
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/creditallocated"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/predicate"
)

// CreditAllocatedQuery is the builder for querying CreditAllocated entities.
type CreditAllocatedQuery struct {
	config
	ctx        *QueryContext
	order      []creditallocated.OrderOption
	inters     []Interceptor
	predicates []predicate.CreditAllocated
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CreditAllocatedQuery builder.
func (caq *CreditAllocatedQuery) Where(ps ...predicate.CreditAllocated) *CreditAllocatedQuery {
	caq.predicates = append(caq.predicates, ps...)
	return caq
}

// Limit the number of records to be returned by this query.
func (caq *CreditAllocatedQuery) Limit(limit int) *CreditAllocatedQuery {
	caq.ctx.Limit = &limit
	return caq
}

// Offset to start from.
func (caq *CreditAllocatedQuery) Offset(offset int) *CreditAllocatedQuery {
	caq.ctx.Offset = &offset
	return caq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (caq *CreditAllocatedQuery) Unique(unique bool) *CreditAllocatedQuery {
	caq.ctx.Unique = &unique
	return caq
}

// Order specifies how the records should be ordered.
func (caq *CreditAllocatedQuery) Order(o ...creditallocated.OrderOption) *CreditAllocatedQuery {
	caq.order = append(caq.order, o...)
	return caq
}

// First returns the first CreditAllocated entity from the query.
// Returns a *NotFoundError when no CreditAllocated was found.
func (caq *CreditAllocatedQuery) First(ctx context.Context) (*CreditAllocated, error) {
	nodes, err := caq.Limit(1).All(setContextOp(ctx, caq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{creditallocated.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (caq *CreditAllocatedQuery) FirstX(ctx context.Context) *CreditAllocated {
	node, err := caq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CreditAllocated ID from the query.
// Returns a *NotFoundError when no CreditAllocated ID was found.
func (caq *CreditAllocatedQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = caq.Limit(1).IDs(setContextOp(ctx, caq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{creditallocated.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (caq *CreditAllocatedQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := caq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CreditAllocated entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CreditAllocated entity is found.
// Returns a *NotFoundError when no CreditAllocated entities are found.
func (caq *CreditAllocatedQuery) Only(ctx context.Context) (*CreditAllocated, error) {
	nodes, err := caq.Limit(2).All(setContextOp(ctx, caq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{creditallocated.Label}
	default:
		return nil, &NotSingularError{creditallocated.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (caq *CreditAllocatedQuery) OnlyX(ctx context.Context) *CreditAllocated {
	node, err := caq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CreditAllocated ID in the query.
// Returns a *NotSingularError when more than one CreditAllocated ID is found.
// Returns a *NotFoundError when no entities are found.
func (caq *CreditAllocatedQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = caq.Limit(2).IDs(setContextOp(ctx, caq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{creditallocated.Label}
	default:
		err = &NotSingularError{creditallocated.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (caq *CreditAllocatedQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := caq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CreditAllocateds.
func (caq *CreditAllocatedQuery) All(ctx context.Context) ([]*CreditAllocated, error) {
	ctx = setContextOp(ctx, caq.ctx, ent.OpQueryAll)
	if err := caq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CreditAllocated, *CreditAllocatedQuery]()
	return withInterceptors[[]*CreditAllocated](ctx, caq, qr, caq.inters)
}

// AllX is like All, but panics if an error occurs.
func (caq *CreditAllocatedQuery) AllX(ctx context.Context) []*CreditAllocated {
	nodes, err := caq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CreditAllocated IDs.
func (caq *CreditAllocatedQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if caq.ctx.Unique == nil && caq.path != nil {
		caq.Unique(true)
	}
	ctx = setContextOp(ctx, caq.ctx, ent.OpQueryIDs)
	if err = caq.Select(creditallocated.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (caq *CreditAllocatedQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := caq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (caq *CreditAllocatedQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, caq.ctx, ent.OpQueryCount)
	if err := caq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, caq, querierCount[*CreditAllocatedQuery](), caq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (caq *CreditAllocatedQuery) CountX(ctx context.Context) int {
	count, err := caq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (caq *CreditAllocatedQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, caq.ctx, ent.OpQueryExist)
	switch _, err := caq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (caq *CreditAllocatedQuery) ExistX(ctx context.Context) bool {
	exist, err := caq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CreditAllocatedQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (caq *CreditAllocatedQuery) Clone() *CreditAllocatedQuery {
	if caq == nil {
		return nil
	}
	return &CreditAllocatedQuery{
		config:     caq.config,
		ctx:        caq.ctx.Clone(),
		order:      append([]creditallocated.OrderOption{}, caq.order...),
		inters:     append([]Interceptor{}, caq.inters...),
		predicates: append([]predicate.CreditAllocated{}, caq.predicates...),
		// clone intermediate query.
		sql:       caq.sql.Clone(),
		path:      caq.path,
		modifiers: append([]func(*sql.Selector){}, caq.modifiers...),
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
//	client.CreditAllocated.Query().
//		GroupBy(creditallocated.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (caq *CreditAllocatedQuery) GroupBy(field string, fields ...string) *CreditAllocatedGroupBy {
	caq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CreditAllocatedGroupBy{build: caq}
	grbuild.flds = &caq.ctx.Fields
	grbuild.label = creditallocated.Label
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
//	client.CreditAllocated.Query().
//		Select(creditallocated.FieldCreatedAt).
//		Scan(ctx, &v)
func (caq *CreditAllocatedQuery) Select(fields ...string) *CreditAllocatedSelect {
	caq.ctx.Fields = append(caq.ctx.Fields, fields...)
	sbuild := &CreditAllocatedSelect{CreditAllocatedQuery: caq}
	sbuild.label = creditallocated.Label
	sbuild.flds, sbuild.scan = &caq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CreditAllocatedSelect configured with the given aggregations.
func (caq *CreditAllocatedQuery) Aggregate(fns ...AggregateFunc) *CreditAllocatedSelect {
	return caq.Select().Aggregate(fns...)
}

func (caq *CreditAllocatedQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range caq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, caq); err != nil {
				return err
			}
		}
	}
	for _, f := range caq.ctx.Fields {
		if !creditallocated.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if caq.path != nil {
		prev, err := caq.path(ctx)
		if err != nil {
			return err
		}
		caq.sql = prev
	}
	return nil
}

func (caq *CreditAllocatedQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CreditAllocated, error) {
	var (
		nodes = []*CreditAllocated{}
		_spec = caq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CreditAllocated).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CreditAllocated{config: caq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(caq.modifiers) > 0 {
		_spec.Modifiers = caq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, caq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (caq *CreditAllocatedQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := caq.querySpec()
	if len(caq.modifiers) > 0 {
		_spec.Modifiers = caq.modifiers
	}
	_spec.Node.Columns = caq.ctx.Fields
	if len(caq.ctx.Fields) > 0 {
		_spec.Unique = caq.ctx.Unique != nil && *caq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, caq.driver, _spec)
}

func (caq *CreditAllocatedQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(creditallocated.Table, creditallocated.Columns, sqlgraph.NewFieldSpec(creditallocated.FieldID, field.TypeUint32))
	_spec.From = caq.sql
	if unique := caq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if caq.path != nil {
		_spec.Unique = true
	}
	if fields := caq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, creditallocated.FieldID)
		for i := range fields {
			if fields[i] != creditallocated.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := caq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := caq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := caq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := caq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (caq *CreditAllocatedQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(caq.driver.Dialect())
	t1 := builder.Table(creditallocated.Table)
	columns := caq.ctx.Fields
	if len(columns) == 0 {
		columns = creditallocated.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if caq.sql != nil {
		selector = caq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if caq.ctx.Unique != nil && *caq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range caq.modifiers {
		m(selector)
	}
	for _, p := range caq.predicates {
		p(selector)
	}
	for _, p := range caq.order {
		p(selector)
	}
	if offset := caq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := caq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (caq *CreditAllocatedQuery) ForUpdate(opts ...sql.LockOption) *CreditAllocatedQuery {
	if caq.driver.Dialect() == dialect.Postgres {
		caq.Unique(false)
	}
	caq.modifiers = append(caq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return caq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (caq *CreditAllocatedQuery) ForShare(opts ...sql.LockOption) *CreditAllocatedQuery {
	if caq.driver.Dialect() == dialect.Postgres {
		caq.Unique(false)
	}
	caq.modifiers = append(caq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return caq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (caq *CreditAllocatedQuery) Modify(modifiers ...func(s *sql.Selector)) *CreditAllocatedSelect {
	caq.modifiers = append(caq.modifiers, modifiers...)
	return caq.Select()
}

// CreditAllocatedGroupBy is the group-by builder for CreditAllocated entities.
type CreditAllocatedGroupBy struct {
	selector
	build *CreditAllocatedQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cagb *CreditAllocatedGroupBy) Aggregate(fns ...AggregateFunc) *CreditAllocatedGroupBy {
	cagb.fns = append(cagb.fns, fns...)
	return cagb
}

// Scan applies the selector query and scans the result into the given value.
func (cagb *CreditAllocatedGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cagb.build.ctx, ent.OpQueryGroupBy)
	if err := cagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CreditAllocatedQuery, *CreditAllocatedGroupBy](ctx, cagb.build, cagb, cagb.build.inters, v)
}

func (cagb *CreditAllocatedGroupBy) sqlScan(ctx context.Context, root *CreditAllocatedQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cagb.fns))
	for _, fn := range cagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cagb.flds)+len(cagb.fns))
		for _, f := range *cagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CreditAllocatedSelect is the builder for selecting fields of CreditAllocated entities.
type CreditAllocatedSelect struct {
	*CreditAllocatedQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cas *CreditAllocatedSelect) Aggregate(fns ...AggregateFunc) *CreditAllocatedSelect {
	cas.fns = append(cas.fns, fns...)
	return cas
}

// Scan applies the selector query and scans the result into the given value.
func (cas *CreditAllocatedSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cas.ctx, ent.OpQuerySelect)
	if err := cas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CreditAllocatedQuery, *CreditAllocatedSelect](ctx, cas.CreditAllocatedQuery, cas, cas.inters, v)
}

func (cas *CreditAllocatedSelect) sqlScan(ctx context.Context, root *CreditAllocatedQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cas.fns))
	for _, fn := range cas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cas *CreditAllocatedSelect) Modify(modifiers ...func(s *sql.Selector)) *CreditAllocatedSelect {
	cas.modifiers = append(cas.modifiers, modifiers...)
	return cas
}
