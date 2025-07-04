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
	"github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/powerrentalstate"
	"github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/predicate"
)

// PowerRentalStateQuery is the builder for querying PowerRentalState entities.
type PowerRentalStateQuery struct {
	config
	ctx        *QueryContext
	order      []powerrentalstate.OrderOption
	inters     []Interceptor
	predicates []predicate.PowerRentalState
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PowerRentalStateQuery builder.
func (prsq *PowerRentalStateQuery) Where(ps ...predicate.PowerRentalState) *PowerRentalStateQuery {
	prsq.predicates = append(prsq.predicates, ps...)
	return prsq
}

// Limit the number of records to be returned by this query.
func (prsq *PowerRentalStateQuery) Limit(limit int) *PowerRentalStateQuery {
	prsq.ctx.Limit = &limit
	return prsq
}

// Offset to start from.
func (prsq *PowerRentalStateQuery) Offset(offset int) *PowerRentalStateQuery {
	prsq.ctx.Offset = &offset
	return prsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (prsq *PowerRentalStateQuery) Unique(unique bool) *PowerRentalStateQuery {
	prsq.ctx.Unique = &unique
	return prsq
}

// Order specifies how the records should be ordered.
func (prsq *PowerRentalStateQuery) Order(o ...powerrentalstate.OrderOption) *PowerRentalStateQuery {
	prsq.order = append(prsq.order, o...)
	return prsq
}

// First returns the first PowerRentalState entity from the query.
// Returns a *NotFoundError when no PowerRentalState was found.
func (prsq *PowerRentalStateQuery) First(ctx context.Context) (*PowerRentalState, error) {
	nodes, err := prsq.Limit(1).All(setContextOp(ctx, prsq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{powerrentalstate.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (prsq *PowerRentalStateQuery) FirstX(ctx context.Context) *PowerRentalState {
	node, err := prsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PowerRentalState ID from the query.
// Returns a *NotFoundError when no PowerRentalState ID was found.
func (prsq *PowerRentalStateQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = prsq.Limit(1).IDs(setContextOp(ctx, prsq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{powerrentalstate.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (prsq *PowerRentalStateQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := prsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PowerRentalState entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PowerRentalState entity is found.
// Returns a *NotFoundError when no PowerRentalState entities are found.
func (prsq *PowerRentalStateQuery) Only(ctx context.Context) (*PowerRentalState, error) {
	nodes, err := prsq.Limit(2).All(setContextOp(ctx, prsq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{powerrentalstate.Label}
	default:
		return nil, &NotSingularError{powerrentalstate.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (prsq *PowerRentalStateQuery) OnlyX(ctx context.Context) *PowerRentalState {
	node, err := prsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PowerRentalState ID in the query.
// Returns a *NotSingularError when more than one PowerRentalState ID is found.
// Returns a *NotFoundError when no entities are found.
func (prsq *PowerRentalStateQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = prsq.Limit(2).IDs(setContextOp(ctx, prsq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{powerrentalstate.Label}
	default:
		err = &NotSingularError{powerrentalstate.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (prsq *PowerRentalStateQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := prsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PowerRentalStates.
func (prsq *PowerRentalStateQuery) All(ctx context.Context) ([]*PowerRentalState, error) {
	ctx = setContextOp(ctx, prsq.ctx, ent.OpQueryAll)
	if err := prsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PowerRentalState, *PowerRentalStateQuery]()
	return withInterceptors[[]*PowerRentalState](ctx, prsq, qr, prsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (prsq *PowerRentalStateQuery) AllX(ctx context.Context) []*PowerRentalState {
	nodes, err := prsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PowerRentalState IDs.
func (prsq *PowerRentalStateQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if prsq.ctx.Unique == nil && prsq.path != nil {
		prsq.Unique(true)
	}
	ctx = setContextOp(ctx, prsq.ctx, ent.OpQueryIDs)
	if err = prsq.Select(powerrentalstate.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (prsq *PowerRentalStateQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := prsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (prsq *PowerRentalStateQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, prsq.ctx, ent.OpQueryCount)
	if err := prsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, prsq, querierCount[*PowerRentalStateQuery](), prsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (prsq *PowerRentalStateQuery) CountX(ctx context.Context) int {
	count, err := prsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (prsq *PowerRentalStateQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, prsq.ctx, ent.OpQueryExist)
	switch _, err := prsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (prsq *PowerRentalStateQuery) ExistX(ctx context.Context) bool {
	exist, err := prsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PowerRentalStateQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (prsq *PowerRentalStateQuery) Clone() *PowerRentalStateQuery {
	if prsq == nil {
		return nil
	}
	return &PowerRentalStateQuery{
		config:     prsq.config,
		ctx:        prsq.ctx.Clone(),
		order:      append([]powerrentalstate.OrderOption{}, prsq.order...),
		inters:     append([]Interceptor{}, prsq.inters...),
		predicates: append([]predicate.PowerRentalState{}, prsq.predicates...),
		// clone intermediate query.
		sql:       prsq.sql.Clone(),
		path:      prsq.path,
		modifiers: append([]func(*sql.Selector){}, prsq.modifiers...),
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
//	client.PowerRentalState.Query().
//		GroupBy(powerrentalstate.FieldEntID).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (prsq *PowerRentalStateQuery) GroupBy(field string, fields ...string) *PowerRentalStateGroupBy {
	prsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PowerRentalStateGroupBy{build: prsq}
	grbuild.flds = &prsq.ctx.Fields
	grbuild.label = powerrentalstate.Label
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
//	client.PowerRentalState.Query().
//		Select(powerrentalstate.FieldEntID).
//		Scan(ctx, &v)
func (prsq *PowerRentalStateQuery) Select(fields ...string) *PowerRentalStateSelect {
	prsq.ctx.Fields = append(prsq.ctx.Fields, fields...)
	sbuild := &PowerRentalStateSelect{PowerRentalStateQuery: prsq}
	sbuild.label = powerrentalstate.Label
	sbuild.flds, sbuild.scan = &prsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PowerRentalStateSelect configured with the given aggregations.
func (prsq *PowerRentalStateQuery) Aggregate(fns ...AggregateFunc) *PowerRentalStateSelect {
	return prsq.Select().Aggregate(fns...)
}

func (prsq *PowerRentalStateQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range prsq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, prsq); err != nil {
				return err
			}
		}
	}
	for _, f := range prsq.ctx.Fields {
		if !powerrentalstate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if prsq.path != nil {
		prev, err := prsq.path(ctx)
		if err != nil {
			return err
		}
		prsq.sql = prev
	}
	return nil
}

func (prsq *PowerRentalStateQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PowerRentalState, error) {
	var (
		nodes = []*PowerRentalState{}
		_spec = prsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PowerRentalState).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PowerRentalState{config: prsq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(prsq.modifiers) > 0 {
		_spec.Modifiers = prsq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, prsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (prsq *PowerRentalStateQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := prsq.querySpec()
	if len(prsq.modifiers) > 0 {
		_spec.Modifiers = prsq.modifiers
	}
	_spec.Node.Columns = prsq.ctx.Fields
	if len(prsq.ctx.Fields) > 0 {
		_spec.Unique = prsq.ctx.Unique != nil && *prsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, prsq.driver, _spec)
}

func (prsq *PowerRentalStateQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(powerrentalstate.Table, powerrentalstate.Columns, sqlgraph.NewFieldSpec(powerrentalstate.FieldID, field.TypeUint32))
	_spec.From = prsq.sql
	if unique := prsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if prsq.path != nil {
		_spec.Unique = true
	}
	if fields := prsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, powerrentalstate.FieldID)
		for i := range fields {
			if fields[i] != powerrentalstate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := prsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := prsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := prsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := prsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (prsq *PowerRentalStateQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(prsq.driver.Dialect())
	t1 := builder.Table(powerrentalstate.Table)
	columns := prsq.ctx.Fields
	if len(columns) == 0 {
		columns = powerrentalstate.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if prsq.sql != nil {
		selector = prsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if prsq.ctx.Unique != nil && *prsq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range prsq.modifiers {
		m(selector)
	}
	for _, p := range prsq.predicates {
		p(selector)
	}
	for _, p := range prsq.order {
		p(selector)
	}
	if offset := prsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := prsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (prsq *PowerRentalStateQuery) ForUpdate(opts ...sql.LockOption) *PowerRentalStateQuery {
	if prsq.driver.Dialect() == dialect.Postgres {
		prsq.Unique(false)
	}
	prsq.modifiers = append(prsq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return prsq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (prsq *PowerRentalStateQuery) ForShare(opts ...sql.LockOption) *PowerRentalStateQuery {
	if prsq.driver.Dialect() == dialect.Postgres {
		prsq.Unique(false)
	}
	prsq.modifiers = append(prsq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return prsq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (prsq *PowerRentalStateQuery) Modify(modifiers ...func(s *sql.Selector)) *PowerRentalStateSelect {
	prsq.modifiers = append(prsq.modifiers, modifiers...)
	return prsq.Select()
}

// PowerRentalStateGroupBy is the group-by builder for PowerRentalState entities.
type PowerRentalStateGroupBy struct {
	selector
	build *PowerRentalStateQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (prsgb *PowerRentalStateGroupBy) Aggregate(fns ...AggregateFunc) *PowerRentalStateGroupBy {
	prsgb.fns = append(prsgb.fns, fns...)
	return prsgb
}

// Scan applies the selector query and scans the result into the given value.
func (prsgb *PowerRentalStateGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, prsgb.build.ctx, ent.OpQueryGroupBy)
	if err := prsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PowerRentalStateQuery, *PowerRentalStateGroupBy](ctx, prsgb.build, prsgb, prsgb.build.inters, v)
}

func (prsgb *PowerRentalStateGroupBy) sqlScan(ctx context.Context, root *PowerRentalStateQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(prsgb.fns))
	for _, fn := range prsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*prsgb.flds)+len(prsgb.fns))
		for _, f := range *prsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*prsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := prsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PowerRentalStateSelect is the builder for selecting fields of PowerRentalState entities.
type PowerRentalStateSelect struct {
	*PowerRentalStateQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (prss *PowerRentalStateSelect) Aggregate(fns ...AggregateFunc) *PowerRentalStateSelect {
	prss.fns = append(prss.fns, fns...)
	return prss
}

// Scan applies the selector query and scans the result into the given value.
func (prss *PowerRentalStateSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, prss.ctx, ent.OpQuerySelect)
	if err := prss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PowerRentalStateQuery, *PowerRentalStateSelect](ctx, prss.PowerRentalStateQuery, prss, prss.inters, v)
}

func (prss *PowerRentalStateSelect) sqlScan(ctx context.Context, root *PowerRentalStateQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(prss.fns))
	for _, fn := range prss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*prss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := prss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (prss *PowerRentalStateSelect) Modify(modifiers ...func(s *sql.Selector)) *PowerRentalStateSelect {
	prss.modifiers = append(prss.modifiers, modifiers...)
	return prss
}
