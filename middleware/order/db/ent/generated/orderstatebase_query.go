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
	"github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderstatebase"
	"github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/predicate"
)

// OrderStateBaseQuery is the builder for querying OrderStateBase entities.
type OrderStateBaseQuery struct {
	config
	ctx        *QueryContext
	order      []orderstatebase.OrderOption
	inters     []Interceptor
	predicates []predicate.OrderStateBase
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderStateBaseQuery builder.
func (osbq *OrderStateBaseQuery) Where(ps ...predicate.OrderStateBase) *OrderStateBaseQuery {
	osbq.predicates = append(osbq.predicates, ps...)
	return osbq
}

// Limit the number of records to be returned by this query.
func (osbq *OrderStateBaseQuery) Limit(limit int) *OrderStateBaseQuery {
	osbq.ctx.Limit = &limit
	return osbq
}

// Offset to start from.
func (osbq *OrderStateBaseQuery) Offset(offset int) *OrderStateBaseQuery {
	osbq.ctx.Offset = &offset
	return osbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (osbq *OrderStateBaseQuery) Unique(unique bool) *OrderStateBaseQuery {
	osbq.ctx.Unique = &unique
	return osbq
}

// Order specifies how the records should be ordered.
func (osbq *OrderStateBaseQuery) Order(o ...orderstatebase.OrderOption) *OrderStateBaseQuery {
	osbq.order = append(osbq.order, o...)
	return osbq
}

// First returns the first OrderStateBase entity from the query.
// Returns a *NotFoundError when no OrderStateBase was found.
func (osbq *OrderStateBaseQuery) First(ctx context.Context) (*OrderStateBase, error) {
	nodes, err := osbq.Limit(1).All(setContextOp(ctx, osbq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orderstatebase.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (osbq *OrderStateBaseQuery) FirstX(ctx context.Context) *OrderStateBase {
	node, err := osbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderStateBase ID from the query.
// Returns a *NotFoundError when no OrderStateBase ID was found.
func (osbq *OrderStateBaseQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = osbq.Limit(1).IDs(setContextOp(ctx, osbq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orderstatebase.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (osbq *OrderStateBaseQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := osbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderStateBase entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrderStateBase entity is found.
// Returns a *NotFoundError when no OrderStateBase entities are found.
func (osbq *OrderStateBaseQuery) Only(ctx context.Context) (*OrderStateBase, error) {
	nodes, err := osbq.Limit(2).All(setContextOp(ctx, osbq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orderstatebase.Label}
	default:
		return nil, &NotSingularError{orderstatebase.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (osbq *OrderStateBaseQuery) OnlyX(ctx context.Context) *OrderStateBase {
	node, err := osbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderStateBase ID in the query.
// Returns a *NotSingularError when more than one OrderStateBase ID is found.
// Returns a *NotFoundError when no entities are found.
func (osbq *OrderStateBaseQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = osbq.Limit(2).IDs(setContextOp(ctx, osbq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orderstatebase.Label}
	default:
		err = &NotSingularError{orderstatebase.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (osbq *OrderStateBaseQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := osbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderStateBases.
func (osbq *OrderStateBaseQuery) All(ctx context.Context) ([]*OrderStateBase, error) {
	ctx = setContextOp(ctx, osbq.ctx, ent.OpQueryAll)
	if err := osbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OrderStateBase, *OrderStateBaseQuery]()
	return withInterceptors[[]*OrderStateBase](ctx, osbq, qr, osbq.inters)
}

// AllX is like All, but panics if an error occurs.
func (osbq *OrderStateBaseQuery) AllX(ctx context.Context) []*OrderStateBase {
	nodes, err := osbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderStateBase IDs.
func (osbq *OrderStateBaseQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if osbq.ctx.Unique == nil && osbq.path != nil {
		osbq.Unique(true)
	}
	ctx = setContextOp(ctx, osbq.ctx, ent.OpQueryIDs)
	if err = osbq.Select(orderstatebase.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (osbq *OrderStateBaseQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := osbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (osbq *OrderStateBaseQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, osbq.ctx, ent.OpQueryCount)
	if err := osbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, osbq, querierCount[*OrderStateBaseQuery](), osbq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (osbq *OrderStateBaseQuery) CountX(ctx context.Context) int {
	count, err := osbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (osbq *OrderStateBaseQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, osbq.ctx, ent.OpQueryExist)
	switch _, err := osbq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (osbq *OrderStateBaseQuery) ExistX(ctx context.Context) bool {
	exist, err := osbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderStateBaseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (osbq *OrderStateBaseQuery) Clone() *OrderStateBaseQuery {
	if osbq == nil {
		return nil
	}
	return &OrderStateBaseQuery{
		config:     osbq.config,
		ctx:        osbq.ctx.Clone(),
		order:      append([]orderstatebase.OrderOption{}, osbq.order...),
		inters:     append([]Interceptor{}, osbq.inters...),
		predicates: append([]predicate.OrderStateBase{}, osbq.predicates...),
		// clone intermediate query.
		sql:       osbq.sql.Clone(),
		path:      osbq.path,
		modifiers: append([]func(*sql.Selector){}, osbq.modifiers...),
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
//	client.OrderStateBase.Query().
//		GroupBy(orderstatebase.FieldEntID).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (osbq *OrderStateBaseQuery) GroupBy(field string, fields ...string) *OrderStateBaseGroupBy {
	osbq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrderStateBaseGroupBy{build: osbq}
	grbuild.flds = &osbq.ctx.Fields
	grbuild.label = orderstatebase.Label
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
//	client.OrderStateBase.Query().
//		Select(orderstatebase.FieldEntID).
//		Scan(ctx, &v)
func (osbq *OrderStateBaseQuery) Select(fields ...string) *OrderStateBaseSelect {
	osbq.ctx.Fields = append(osbq.ctx.Fields, fields...)
	sbuild := &OrderStateBaseSelect{OrderStateBaseQuery: osbq}
	sbuild.label = orderstatebase.Label
	sbuild.flds, sbuild.scan = &osbq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrderStateBaseSelect configured with the given aggregations.
func (osbq *OrderStateBaseQuery) Aggregate(fns ...AggregateFunc) *OrderStateBaseSelect {
	return osbq.Select().Aggregate(fns...)
}

func (osbq *OrderStateBaseQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range osbq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, osbq); err != nil {
				return err
			}
		}
	}
	for _, f := range osbq.ctx.Fields {
		if !orderstatebase.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if osbq.path != nil {
		prev, err := osbq.path(ctx)
		if err != nil {
			return err
		}
		osbq.sql = prev
	}
	return nil
}

func (osbq *OrderStateBaseQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrderStateBase, error) {
	var (
		nodes = []*OrderStateBase{}
		_spec = osbq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrderStateBase).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrderStateBase{config: osbq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(osbq.modifiers) > 0 {
		_spec.Modifiers = osbq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, osbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (osbq *OrderStateBaseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := osbq.querySpec()
	if len(osbq.modifiers) > 0 {
		_spec.Modifiers = osbq.modifiers
	}
	_spec.Node.Columns = osbq.ctx.Fields
	if len(osbq.ctx.Fields) > 0 {
		_spec.Unique = osbq.ctx.Unique != nil && *osbq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, osbq.driver, _spec)
}

func (osbq *OrderStateBaseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(orderstatebase.Table, orderstatebase.Columns, sqlgraph.NewFieldSpec(orderstatebase.FieldID, field.TypeUint32))
	_spec.From = osbq.sql
	if unique := osbq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if osbq.path != nil {
		_spec.Unique = true
	}
	if fields := osbq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderstatebase.FieldID)
		for i := range fields {
			if fields[i] != orderstatebase.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := osbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := osbq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := osbq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := osbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (osbq *OrderStateBaseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(osbq.driver.Dialect())
	t1 := builder.Table(orderstatebase.Table)
	columns := osbq.ctx.Fields
	if len(columns) == 0 {
		columns = orderstatebase.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if osbq.sql != nil {
		selector = osbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if osbq.ctx.Unique != nil && *osbq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range osbq.modifiers {
		m(selector)
	}
	for _, p := range osbq.predicates {
		p(selector)
	}
	for _, p := range osbq.order {
		p(selector)
	}
	if offset := osbq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := osbq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (osbq *OrderStateBaseQuery) ForUpdate(opts ...sql.LockOption) *OrderStateBaseQuery {
	if osbq.driver.Dialect() == dialect.Postgres {
		osbq.Unique(false)
	}
	osbq.modifiers = append(osbq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return osbq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (osbq *OrderStateBaseQuery) ForShare(opts ...sql.LockOption) *OrderStateBaseQuery {
	if osbq.driver.Dialect() == dialect.Postgres {
		osbq.Unique(false)
	}
	osbq.modifiers = append(osbq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return osbq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (osbq *OrderStateBaseQuery) Modify(modifiers ...func(s *sql.Selector)) *OrderStateBaseSelect {
	osbq.modifiers = append(osbq.modifiers, modifiers...)
	return osbq.Select()
}

// OrderStateBaseGroupBy is the group-by builder for OrderStateBase entities.
type OrderStateBaseGroupBy struct {
	selector
	build *OrderStateBaseQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (osbgb *OrderStateBaseGroupBy) Aggregate(fns ...AggregateFunc) *OrderStateBaseGroupBy {
	osbgb.fns = append(osbgb.fns, fns...)
	return osbgb
}

// Scan applies the selector query and scans the result into the given value.
func (osbgb *OrderStateBaseGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, osbgb.build.ctx, ent.OpQueryGroupBy)
	if err := osbgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrderStateBaseQuery, *OrderStateBaseGroupBy](ctx, osbgb.build, osbgb, osbgb.build.inters, v)
}

func (osbgb *OrderStateBaseGroupBy) sqlScan(ctx context.Context, root *OrderStateBaseQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(osbgb.fns))
	for _, fn := range osbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*osbgb.flds)+len(osbgb.fns))
		for _, f := range *osbgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*osbgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := osbgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrderStateBaseSelect is the builder for selecting fields of OrderStateBase entities.
type OrderStateBaseSelect struct {
	*OrderStateBaseQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (osbs *OrderStateBaseSelect) Aggregate(fns ...AggregateFunc) *OrderStateBaseSelect {
	osbs.fns = append(osbs.fns, fns...)
	return osbs
}

// Scan applies the selector query and scans the result into the given value.
func (osbs *OrderStateBaseSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, osbs.ctx, ent.OpQuerySelect)
	if err := osbs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrderStateBaseQuery, *OrderStateBaseSelect](ctx, osbs.OrderStateBaseQuery, osbs, osbs.inters, v)
}

func (osbs *OrderStateBaseSelect) sqlScan(ctx context.Context, root *OrderStateBaseQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(osbs.fns))
	for _, fn := range osbs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*osbs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := osbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (osbs *OrderStateBaseSelect) Modify(modifiers ...func(s *sql.Selector)) *OrderStateBaseSelect {
	osbs.modifiers = append(osbs.modifiers, modifiers...)
	return osbs
}
