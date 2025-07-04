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
	"github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/orderbenefit"
	"github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/predicate"
)

// OrderBenefitQuery is the builder for querying OrderBenefit entities.
type OrderBenefitQuery struct {
	config
	ctx        *QueryContext
	order      []orderbenefit.OrderOption
	inters     []Interceptor
	predicates []predicate.OrderBenefit
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderBenefitQuery builder.
func (obq *OrderBenefitQuery) Where(ps ...predicate.OrderBenefit) *OrderBenefitQuery {
	obq.predicates = append(obq.predicates, ps...)
	return obq
}

// Limit the number of records to be returned by this query.
func (obq *OrderBenefitQuery) Limit(limit int) *OrderBenefitQuery {
	obq.ctx.Limit = &limit
	return obq
}

// Offset to start from.
func (obq *OrderBenefitQuery) Offset(offset int) *OrderBenefitQuery {
	obq.ctx.Offset = &offset
	return obq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (obq *OrderBenefitQuery) Unique(unique bool) *OrderBenefitQuery {
	obq.ctx.Unique = &unique
	return obq
}

// Order specifies how the records should be ordered.
func (obq *OrderBenefitQuery) Order(o ...orderbenefit.OrderOption) *OrderBenefitQuery {
	obq.order = append(obq.order, o...)
	return obq
}

// First returns the first OrderBenefit entity from the query.
// Returns a *NotFoundError when no OrderBenefit was found.
func (obq *OrderBenefitQuery) First(ctx context.Context) (*OrderBenefit, error) {
	nodes, err := obq.Limit(1).All(setContextOp(ctx, obq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orderbenefit.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (obq *OrderBenefitQuery) FirstX(ctx context.Context) *OrderBenefit {
	node, err := obq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderBenefit ID from the query.
// Returns a *NotFoundError when no OrderBenefit ID was found.
func (obq *OrderBenefitQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = obq.Limit(1).IDs(setContextOp(ctx, obq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orderbenefit.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (obq *OrderBenefitQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := obq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderBenefit entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrderBenefit entity is found.
// Returns a *NotFoundError when no OrderBenefit entities are found.
func (obq *OrderBenefitQuery) Only(ctx context.Context) (*OrderBenefit, error) {
	nodes, err := obq.Limit(2).All(setContextOp(ctx, obq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orderbenefit.Label}
	default:
		return nil, &NotSingularError{orderbenefit.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (obq *OrderBenefitQuery) OnlyX(ctx context.Context) *OrderBenefit {
	node, err := obq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderBenefit ID in the query.
// Returns a *NotSingularError when more than one OrderBenefit ID is found.
// Returns a *NotFoundError when no entities are found.
func (obq *OrderBenefitQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = obq.Limit(2).IDs(setContextOp(ctx, obq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orderbenefit.Label}
	default:
		err = &NotSingularError{orderbenefit.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (obq *OrderBenefitQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := obq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderBenefits.
func (obq *OrderBenefitQuery) All(ctx context.Context) ([]*OrderBenefit, error) {
	ctx = setContextOp(ctx, obq.ctx, ent.OpQueryAll)
	if err := obq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OrderBenefit, *OrderBenefitQuery]()
	return withInterceptors[[]*OrderBenefit](ctx, obq, qr, obq.inters)
}

// AllX is like All, but panics if an error occurs.
func (obq *OrderBenefitQuery) AllX(ctx context.Context) []*OrderBenefit {
	nodes, err := obq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderBenefit IDs.
func (obq *OrderBenefitQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if obq.ctx.Unique == nil && obq.path != nil {
		obq.Unique(true)
	}
	ctx = setContextOp(ctx, obq.ctx, ent.OpQueryIDs)
	if err = obq.Select(orderbenefit.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (obq *OrderBenefitQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := obq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (obq *OrderBenefitQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, obq.ctx, ent.OpQueryCount)
	if err := obq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, obq, querierCount[*OrderBenefitQuery](), obq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (obq *OrderBenefitQuery) CountX(ctx context.Context) int {
	count, err := obq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (obq *OrderBenefitQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, obq.ctx, ent.OpQueryExist)
	switch _, err := obq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (obq *OrderBenefitQuery) ExistX(ctx context.Context) bool {
	exist, err := obq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderBenefitQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (obq *OrderBenefitQuery) Clone() *OrderBenefitQuery {
	if obq == nil {
		return nil
	}
	return &OrderBenefitQuery{
		config:     obq.config,
		ctx:        obq.ctx.Clone(),
		order:      append([]orderbenefit.OrderOption{}, obq.order...),
		inters:     append([]Interceptor{}, obq.inters...),
		predicates: append([]predicate.OrderBenefit{}, obq.predicates...),
		// clone intermediate query.
		sql:       obq.sql.Clone(),
		path:      obq.path,
		modifiers: append([]func(*sql.Selector){}, obq.modifiers...),
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
//	client.OrderBenefit.Query().
//		GroupBy(orderbenefit.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (obq *OrderBenefitQuery) GroupBy(field string, fields ...string) *OrderBenefitGroupBy {
	obq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrderBenefitGroupBy{build: obq}
	grbuild.flds = &obq.ctx.Fields
	grbuild.label = orderbenefit.Label
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
//	client.OrderBenefit.Query().
//		Select(orderbenefit.FieldCreatedAt).
//		Scan(ctx, &v)
func (obq *OrderBenefitQuery) Select(fields ...string) *OrderBenefitSelect {
	obq.ctx.Fields = append(obq.ctx.Fields, fields...)
	sbuild := &OrderBenefitSelect{OrderBenefitQuery: obq}
	sbuild.label = orderbenefit.Label
	sbuild.flds, sbuild.scan = &obq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrderBenefitSelect configured with the given aggregations.
func (obq *OrderBenefitQuery) Aggregate(fns ...AggregateFunc) *OrderBenefitSelect {
	return obq.Select().Aggregate(fns...)
}

func (obq *OrderBenefitQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range obq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, obq); err != nil {
				return err
			}
		}
	}
	for _, f := range obq.ctx.Fields {
		if !orderbenefit.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if obq.path != nil {
		prev, err := obq.path(ctx)
		if err != nil {
			return err
		}
		obq.sql = prev
	}
	return nil
}

func (obq *OrderBenefitQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrderBenefit, error) {
	var (
		nodes = []*OrderBenefit{}
		_spec = obq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrderBenefit).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrderBenefit{config: obq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(obq.modifiers) > 0 {
		_spec.Modifiers = obq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, obq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (obq *OrderBenefitQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := obq.querySpec()
	if len(obq.modifiers) > 0 {
		_spec.Modifiers = obq.modifiers
	}
	_spec.Node.Columns = obq.ctx.Fields
	if len(obq.ctx.Fields) > 0 {
		_spec.Unique = obq.ctx.Unique != nil && *obq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, obq.driver, _spec)
}

func (obq *OrderBenefitQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(orderbenefit.Table, orderbenefit.Columns, sqlgraph.NewFieldSpec(orderbenefit.FieldID, field.TypeUint32))
	_spec.From = obq.sql
	if unique := obq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if obq.path != nil {
		_spec.Unique = true
	}
	if fields := obq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderbenefit.FieldID)
		for i := range fields {
			if fields[i] != orderbenefit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := obq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := obq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := obq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := obq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (obq *OrderBenefitQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(obq.driver.Dialect())
	t1 := builder.Table(orderbenefit.Table)
	columns := obq.ctx.Fields
	if len(columns) == 0 {
		columns = orderbenefit.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if obq.sql != nil {
		selector = obq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if obq.ctx.Unique != nil && *obq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range obq.modifiers {
		m(selector)
	}
	for _, p := range obq.predicates {
		p(selector)
	}
	for _, p := range obq.order {
		p(selector)
	}
	if offset := obq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := obq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (obq *OrderBenefitQuery) ForUpdate(opts ...sql.LockOption) *OrderBenefitQuery {
	if obq.driver.Dialect() == dialect.Postgres {
		obq.Unique(false)
	}
	obq.modifiers = append(obq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return obq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (obq *OrderBenefitQuery) ForShare(opts ...sql.LockOption) *OrderBenefitQuery {
	if obq.driver.Dialect() == dialect.Postgres {
		obq.Unique(false)
	}
	obq.modifiers = append(obq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return obq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (obq *OrderBenefitQuery) Modify(modifiers ...func(s *sql.Selector)) *OrderBenefitSelect {
	obq.modifiers = append(obq.modifiers, modifiers...)
	return obq.Select()
}

// OrderBenefitGroupBy is the group-by builder for OrderBenefit entities.
type OrderBenefitGroupBy struct {
	selector
	build *OrderBenefitQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (obgb *OrderBenefitGroupBy) Aggregate(fns ...AggregateFunc) *OrderBenefitGroupBy {
	obgb.fns = append(obgb.fns, fns...)
	return obgb
}

// Scan applies the selector query and scans the result into the given value.
func (obgb *OrderBenefitGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, obgb.build.ctx, ent.OpQueryGroupBy)
	if err := obgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrderBenefitQuery, *OrderBenefitGroupBy](ctx, obgb.build, obgb, obgb.build.inters, v)
}

func (obgb *OrderBenefitGroupBy) sqlScan(ctx context.Context, root *OrderBenefitQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(obgb.fns))
	for _, fn := range obgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*obgb.flds)+len(obgb.fns))
		for _, f := range *obgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*obgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := obgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrderBenefitSelect is the builder for selecting fields of OrderBenefit entities.
type OrderBenefitSelect struct {
	*OrderBenefitQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (obs *OrderBenefitSelect) Aggregate(fns ...AggregateFunc) *OrderBenefitSelect {
	obs.fns = append(obs.fns, fns...)
	return obs
}

// Scan applies the selector query and scans the result into the given value.
func (obs *OrderBenefitSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, obs.ctx, ent.OpQuerySelect)
	if err := obs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrderBenefitQuery, *OrderBenefitSelect](ctx, obs.OrderBenefitQuery, obs, obs.inters, v)
}

func (obs *OrderBenefitSelect) sqlScan(ctx context.Context, root *OrderBenefitQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(obs.fns))
	for _, fn := range obs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*obs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := obs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (obs *OrderBenefitSelect) Modify(modifiers ...func(s *sql.Selector)) *OrderBenefitSelect {
	obs.modifiers = append(obs.modifiers, modifiers...)
	return obs
}
