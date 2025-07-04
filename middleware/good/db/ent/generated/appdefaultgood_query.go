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
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appdefaultgood"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/predicate"
)

// AppDefaultGoodQuery is the builder for querying AppDefaultGood entities.
type AppDefaultGoodQuery struct {
	config
	ctx        *QueryContext
	order      []appdefaultgood.OrderOption
	inters     []Interceptor
	predicates []predicate.AppDefaultGood
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppDefaultGoodQuery builder.
func (adgq *AppDefaultGoodQuery) Where(ps ...predicate.AppDefaultGood) *AppDefaultGoodQuery {
	adgq.predicates = append(adgq.predicates, ps...)
	return adgq
}

// Limit the number of records to be returned by this query.
func (adgq *AppDefaultGoodQuery) Limit(limit int) *AppDefaultGoodQuery {
	adgq.ctx.Limit = &limit
	return adgq
}

// Offset to start from.
func (adgq *AppDefaultGoodQuery) Offset(offset int) *AppDefaultGoodQuery {
	adgq.ctx.Offset = &offset
	return adgq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (adgq *AppDefaultGoodQuery) Unique(unique bool) *AppDefaultGoodQuery {
	adgq.ctx.Unique = &unique
	return adgq
}

// Order specifies how the records should be ordered.
func (adgq *AppDefaultGoodQuery) Order(o ...appdefaultgood.OrderOption) *AppDefaultGoodQuery {
	adgq.order = append(adgq.order, o...)
	return adgq
}

// First returns the first AppDefaultGood entity from the query.
// Returns a *NotFoundError when no AppDefaultGood was found.
func (adgq *AppDefaultGoodQuery) First(ctx context.Context) (*AppDefaultGood, error) {
	nodes, err := adgq.Limit(1).All(setContextOp(ctx, adgq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appdefaultgood.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (adgq *AppDefaultGoodQuery) FirstX(ctx context.Context) *AppDefaultGood {
	node, err := adgq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppDefaultGood ID from the query.
// Returns a *NotFoundError when no AppDefaultGood ID was found.
func (adgq *AppDefaultGoodQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = adgq.Limit(1).IDs(setContextOp(ctx, adgq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appdefaultgood.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (adgq *AppDefaultGoodQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := adgq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppDefaultGood entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppDefaultGood entity is found.
// Returns a *NotFoundError when no AppDefaultGood entities are found.
func (adgq *AppDefaultGoodQuery) Only(ctx context.Context) (*AppDefaultGood, error) {
	nodes, err := adgq.Limit(2).All(setContextOp(ctx, adgq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appdefaultgood.Label}
	default:
		return nil, &NotSingularError{appdefaultgood.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (adgq *AppDefaultGoodQuery) OnlyX(ctx context.Context) *AppDefaultGood {
	node, err := adgq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppDefaultGood ID in the query.
// Returns a *NotSingularError when more than one AppDefaultGood ID is found.
// Returns a *NotFoundError when no entities are found.
func (adgq *AppDefaultGoodQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = adgq.Limit(2).IDs(setContextOp(ctx, adgq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appdefaultgood.Label}
	default:
		err = &NotSingularError{appdefaultgood.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (adgq *AppDefaultGoodQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := adgq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppDefaultGoods.
func (adgq *AppDefaultGoodQuery) All(ctx context.Context) ([]*AppDefaultGood, error) {
	ctx = setContextOp(ctx, adgq.ctx, ent.OpQueryAll)
	if err := adgq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AppDefaultGood, *AppDefaultGoodQuery]()
	return withInterceptors[[]*AppDefaultGood](ctx, adgq, qr, adgq.inters)
}

// AllX is like All, but panics if an error occurs.
func (adgq *AppDefaultGoodQuery) AllX(ctx context.Context) []*AppDefaultGood {
	nodes, err := adgq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppDefaultGood IDs.
func (adgq *AppDefaultGoodQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if adgq.ctx.Unique == nil && adgq.path != nil {
		adgq.Unique(true)
	}
	ctx = setContextOp(ctx, adgq.ctx, ent.OpQueryIDs)
	if err = adgq.Select(appdefaultgood.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (adgq *AppDefaultGoodQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := adgq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (adgq *AppDefaultGoodQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, adgq.ctx, ent.OpQueryCount)
	if err := adgq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, adgq, querierCount[*AppDefaultGoodQuery](), adgq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (adgq *AppDefaultGoodQuery) CountX(ctx context.Context) int {
	count, err := adgq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (adgq *AppDefaultGoodQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, adgq.ctx, ent.OpQueryExist)
	switch _, err := adgq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (adgq *AppDefaultGoodQuery) ExistX(ctx context.Context) bool {
	exist, err := adgq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppDefaultGoodQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (adgq *AppDefaultGoodQuery) Clone() *AppDefaultGoodQuery {
	if adgq == nil {
		return nil
	}
	return &AppDefaultGoodQuery{
		config:     adgq.config,
		ctx:        adgq.ctx.Clone(),
		order:      append([]appdefaultgood.OrderOption{}, adgq.order...),
		inters:     append([]Interceptor{}, adgq.inters...),
		predicates: append([]predicate.AppDefaultGood{}, adgq.predicates...),
		// clone intermediate query.
		sql:       adgq.sql.Clone(),
		path:      adgq.path,
		modifiers: append([]func(*sql.Selector){}, adgq.modifiers...),
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
//	client.AppDefaultGood.Query().
//		GroupBy(appdefaultgood.FieldEntID).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (adgq *AppDefaultGoodQuery) GroupBy(field string, fields ...string) *AppDefaultGoodGroupBy {
	adgq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AppDefaultGoodGroupBy{build: adgq}
	grbuild.flds = &adgq.ctx.Fields
	grbuild.label = appdefaultgood.Label
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
//	client.AppDefaultGood.Query().
//		Select(appdefaultgood.FieldEntID).
//		Scan(ctx, &v)
func (adgq *AppDefaultGoodQuery) Select(fields ...string) *AppDefaultGoodSelect {
	adgq.ctx.Fields = append(adgq.ctx.Fields, fields...)
	sbuild := &AppDefaultGoodSelect{AppDefaultGoodQuery: adgq}
	sbuild.label = appdefaultgood.Label
	sbuild.flds, sbuild.scan = &adgq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AppDefaultGoodSelect configured with the given aggregations.
func (adgq *AppDefaultGoodQuery) Aggregate(fns ...AggregateFunc) *AppDefaultGoodSelect {
	return adgq.Select().Aggregate(fns...)
}

func (adgq *AppDefaultGoodQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range adgq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, adgq); err != nil {
				return err
			}
		}
	}
	for _, f := range adgq.ctx.Fields {
		if !appdefaultgood.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if adgq.path != nil {
		prev, err := adgq.path(ctx)
		if err != nil {
			return err
		}
		adgq.sql = prev
	}
	return nil
}

func (adgq *AppDefaultGoodQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppDefaultGood, error) {
	var (
		nodes = []*AppDefaultGood{}
		_spec = adgq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AppDefaultGood).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AppDefaultGood{config: adgq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(adgq.modifiers) > 0 {
		_spec.Modifiers = adgq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, adgq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (adgq *AppDefaultGoodQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := adgq.querySpec()
	if len(adgq.modifiers) > 0 {
		_spec.Modifiers = adgq.modifiers
	}
	_spec.Node.Columns = adgq.ctx.Fields
	if len(adgq.ctx.Fields) > 0 {
		_spec.Unique = adgq.ctx.Unique != nil && *adgq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, adgq.driver, _spec)
}

func (adgq *AppDefaultGoodQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(appdefaultgood.Table, appdefaultgood.Columns, sqlgraph.NewFieldSpec(appdefaultgood.FieldID, field.TypeUint32))
	_spec.From = adgq.sql
	if unique := adgq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if adgq.path != nil {
		_spec.Unique = true
	}
	if fields := adgq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appdefaultgood.FieldID)
		for i := range fields {
			if fields[i] != appdefaultgood.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := adgq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := adgq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := adgq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := adgq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (adgq *AppDefaultGoodQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(adgq.driver.Dialect())
	t1 := builder.Table(appdefaultgood.Table)
	columns := adgq.ctx.Fields
	if len(columns) == 0 {
		columns = appdefaultgood.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if adgq.sql != nil {
		selector = adgq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if adgq.ctx.Unique != nil && *adgq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range adgq.modifiers {
		m(selector)
	}
	for _, p := range adgq.predicates {
		p(selector)
	}
	for _, p := range adgq.order {
		p(selector)
	}
	if offset := adgq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := adgq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (adgq *AppDefaultGoodQuery) ForUpdate(opts ...sql.LockOption) *AppDefaultGoodQuery {
	if adgq.driver.Dialect() == dialect.Postgres {
		adgq.Unique(false)
	}
	adgq.modifiers = append(adgq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return adgq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (adgq *AppDefaultGoodQuery) ForShare(opts ...sql.LockOption) *AppDefaultGoodQuery {
	if adgq.driver.Dialect() == dialect.Postgres {
		adgq.Unique(false)
	}
	adgq.modifiers = append(adgq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return adgq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (adgq *AppDefaultGoodQuery) Modify(modifiers ...func(s *sql.Selector)) *AppDefaultGoodSelect {
	adgq.modifiers = append(adgq.modifiers, modifiers...)
	return adgq.Select()
}

// AppDefaultGoodGroupBy is the group-by builder for AppDefaultGood entities.
type AppDefaultGoodGroupBy struct {
	selector
	build *AppDefaultGoodQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (adggb *AppDefaultGoodGroupBy) Aggregate(fns ...AggregateFunc) *AppDefaultGoodGroupBy {
	adggb.fns = append(adggb.fns, fns...)
	return adggb
}

// Scan applies the selector query and scans the result into the given value.
func (adggb *AppDefaultGoodGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, adggb.build.ctx, ent.OpQueryGroupBy)
	if err := adggb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppDefaultGoodQuery, *AppDefaultGoodGroupBy](ctx, adggb.build, adggb, adggb.build.inters, v)
}

func (adggb *AppDefaultGoodGroupBy) sqlScan(ctx context.Context, root *AppDefaultGoodQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(adggb.fns))
	for _, fn := range adggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*adggb.flds)+len(adggb.fns))
		for _, f := range *adggb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*adggb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := adggb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AppDefaultGoodSelect is the builder for selecting fields of AppDefaultGood entities.
type AppDefaultGoodSelect struct {
	*AppDefaultGoodQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (adgs *AppDefaultGoodSelect) Aggregate(fns ...AggregateFunc) *AppDefaultGoodSelect {
	adgs.fns = append(adgs.fns, fns...)
	return adgs
}

// Scan applies the selector query and scans the result into the given value.
func (adgs *AppDefaultGoodSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, adgs.ctx, ent.OpQuerySelect)
	if err := adgs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppDefaultGoodQuery, *AppDefaultGoodSelect](ctx, adgs.AppDefaultGoodQuery, adgs, adgs.inters, v)
}

func (adgs *AppDefaultGoodSelect) sqlScan(ctx context.Context, root *AppDefaultGoodQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(adgs.fns))
	for _, fn := range adgs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*adgs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := adgs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (adgs *AppDefaultGoodSelect) Modify(modifiers ...func(s *sql.Selector)) *AppDefaultGoodSelect {
	adgs.modifiers = append(adgs.modifiers, modifiers...)
	return adgs
}
