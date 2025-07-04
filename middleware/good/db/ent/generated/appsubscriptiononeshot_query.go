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
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appsubscriptiononeshot"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/predicate"
)

// AppSubscriptionOneShotQuery is the builder for querying AppSubscriptionOneShot entities.
type AppSubscriptionOneShotQuery struct {
	config
	ctx        *QueryContext
	order      []appsubscriptiononeshot.OrderOption
	inters     []Interceptor
	predicates []predicate.AppSubscriptionOneShot
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppSubscriptionOneShotQuery builder.
func (asosq *AppSubscriptionOneShotQuery) Where(ps ...predicate.AppSubscriptionOneShot) *AppSubscriptionOneShotQuery {
	asosq.predicates = append(asosq.predicates, ps...)
	return asosq
}

// Limit the number of records to be returned by this query.
func (asosq *AppSubscriptionOneShotQuery) Limit(limit int) *AppSubscriptionOneShotQuery {
	asosq.ctx.Limit = &limit
	return asosq
}

// Offset to start from.
func (asosq *AppSubscriptionOneShotQuery) Offset(offset int) *AppSubscriptionOneShotQuery {
	asosq.ctx.Offset = &offset
	return asosq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (asosq *AppSubscriptionOneShotQuery) Unique(unique bool) *AppSubscriptionOneShotQuery {
	asosq.ctx.Unique = &unique
	return asosq
}

// Order specifies how the records should be ordered.
func (asosq *AppSubscriptionOneShotQuery) Order(o ...appsubscriptiononeshot.OrderOption) *AppSubscriptionOneShotQuery {
	asosq.order = append(asosq.order, o...)
	return asosq
}

// First returns the first AppSubscriptionOneShot entity from the query.
// Returns a *NotFoundError when no AppSubscriptionOneShot was found.
func (asosq *AppSubscriptionOneShotQuery) First(ctx context.Context) (*AppSubscriptionOneShot, error) {
	nodes, err := asosq.Limit(1).All(setContextOp(ctx, asosq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appsubscriptiononeshot.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (asosq *AppSubscriptionOneShotQuery) FirstX(ctx context.Context) *AppSubscriptionOneShot {
	node, err := asosq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppSubscriptionOneShot ID from the query.
// Returns a *NotFoundError when no AppSubscriptionOneShot ID was found.
func (asosq *AppSubscriptionOneShotQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = asosq.Limit(1).IDs(setContextOp(ctx, asosq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appsubscriptiononeshot.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (asosq *AppSubscriptionOneShotQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := asosq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppSubscriptionOneShot entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppSubscriptionOneShot entity is found.
// Returns a *NotFoundError when no AppSubscriptionOneShot entities are found.
func (asosq *AppSubscriptionOneShotQuery) Only(ctx context.Context) (*AppSubscriptionOneShot, error) {
	nodes, err := asosq.Limit(2).All(setContextOp(ctx, asosq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appsubscriptiononeshot.Label}
	default:
		return nil, &NotSingularError{appsubscriptiononeshot.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (asosq *AppSubscriptionOneShotQuery) OnlyX(ctx context.Context) *AppSubscriptionOneShot {
	node, err := asosq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppSubscriptionOneShot ID in the query.
// Returns a *NotSingularError when more than one AppSubscriptionOneShot ID is found.
// Returns a *NotFoundError when no entities are found.
func (asosq *AppSubscriptionOneShotQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = asosq.Limit(2).IDs(setContextOp(ctx, asosq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appsubscriptiononeshot.Label}
	default:
		err = &NotSingularError{appsubscriptiononeshot.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (asosq *AppSubscriptionOneShotQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := asosq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppSubscriptionOneShots.
func (asosq *AppSubscriptionOneShotQuery) All(ctx context.Context) ([]*AppSubscriptionOneShot, error) {
	ctx = setContextOp(ctx, asosq.ctx, ent.OpQueryAll)
	if err := asosq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AppSubscriptionOneShot, *AppSubscriptionOneShotQuery]()
	return withInterceptors[[]*AppSubscriptionOneShot](ctx, asosq, qr, asosq.inters)
}

// AllX is like All, but panics if an error occurs.
func (asosq *AppSubscriptionOneShotQuery) AllX(ctx context.Context) []*AppSubscriptionOneShot {
	nodes, err := asosq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppSubscriptionOneShot IDs.
func (asosq *AppSubscriptionOneShotQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if asosq.ctx.Unique == nil && asosq.path != nil {
		asosq.Unique(true)
	}
	ctx = setContextOp(ctx, asosq.ctx, ent.OpQueryIDs)
	if err = asosq.Select(appsubscriptiononeshot.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (asosq *AppSubscriptionOneShotQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := asosq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (asosq *AppSubscriptionOneShotQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, asosq.ctx, ent.OpQueryCount)
	if err := asosq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, asosq, querierCount[*AppSubscriptionOneShotQuery](), asosq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (asosq *AppSubscriptionOneShotQuery) CountX(ctx context.Context) int {
	count, err := asosq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (asosq *AppSubscriptionOneShotQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, asosq.ctx, ent.OpQueryExist)
	switch _, err := asosq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (asosq *AppSubscriptionOneShotQuery) ExistX(ctx context.Context) bool {
	exist, err := asosq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppSubscriptionOneShotQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (asosq *AppSubscriptionOneShotQuery) Clone() *AppSubscriptionOneShotQuery {
	if asosq == nil {
		return nil
	}
	return &AppSubscriptionOneShotQuery{
		config:     asosq.config,
		ctx:        asosq.ctx.Clone(),
		order:      append([]appsubscriptiononeshot.OrderOption{}, asosq.order...),
		inters:     append([]Interceptor{}, asosq.inters...),
		predicates: append([]predicate.AppSubscriptionOneShot{}, asosq.predicates...),
		// clone intermediate query.
		sql:       asosq.sql.Clone(),
		path:      asosq.path,
		modifiers: append([]func(*sql.Selector){}, asosq.modifiers...),
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
//	client.AppSubscriptionOneShot.Query().
//		GroupBy(appsubscriptiononeshot.FieldEntID).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (asosq *AppSubscriptionOneShotQuery) GroupBy(field string, fields ...string) *AppSubscriptionOneShotGroupBy {
	asosq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AppSubscriptionOneShotGroupBy{build: asosq}
	grbuild.flds = &asosq.ctx.Fields
	grbuild.label = appsubscriptiononeshot.Label
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
//	client.AppSubscriptionOneShot.Query().
//		Select(appsubscriptiononeshot.FieldEntID).
//		Scan(ctx, &v)
func (asosq *AppSubscriptionOneShotQuery) Select(fields ...string) *AppSubscriptionOneShotSelect {
	asosq.ctx.Fields = append(asosq.ctx.Fields, fields...)
	sbuild := &AppSubscriptionOneShotSelect{AppSubscriptionOneShotQuery: asosq}
	sbuild.label = appsubscriptiononeshot.Label
	sbuild.flds, sbuild.scan = &asosq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AppSubscriptionOneShotSelect configured with the given aggregations.
func (asosq *AppSubscriptionOneShotQuery) Aggregate(fns ...AggregateFunc) *AppSubscriptionOneShotSelect {
	return asosq.Select().Aggregate(fns...)
}

func (asosq *AppSubscriptionOneShotQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range asosq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, asosq); err != nil {
				return err
			}
		}
	}
	for _, f := range asosq.ctx.Fields {
		if !appsubscriptiononeshot.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if asosq.path != nil {
		prev, err := asosq.path(ctx)
		if err != nil {
			return err
		}
		asosq.sql = prev
	}
	return nil
}

func (asosq *AppSubscriptionOneShotQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppSubscriptionOneShot, error) {
	var (
		nodes = []*AppSubscriptionOneShot{}
		_spec = asosq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AppSubscriptionOneShot).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AppSubscriptionOneShot{config: asosq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(asosq.modifiers) > 0 {
		_spec.Modifiers = asosq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, asosq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (asosq *AppSubscriptionOneShotQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := asosq.querySpec()
	if len(asosq.modifiers) > 0 {
		_spec.Modifiers = asosq.modifiers
	}
	_spec.Node.Columns = asosq.ctx.Fields
	if len(asosq.ctx.Fields) > 0 {
		_spec.Unique = asosq.ctx.Unique != nil && *asosq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, asosq.driver, _spec)
}

func (asosq *AppSubscriptionOneShotQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(appsubscriptiononeshot.Table, appsubscriptiononeshot.Columns, sqlgraph.NewFieldSpec(appsubscriptiononeshot.FieldID, field.TypeUint32))
	_spec.From = asosq.sql
	if unique := asosq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if asosq.path != nil {
		_spec.Unique = true
	}
	if fields := asosq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appsubscriptiononeshot.FieldID)
		for i := range fields {
			if fields[i] != appsubscriptiononeshot.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := asosq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := asosq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := asosq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := asosq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (asosq *AppSubscriptionOneShotQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(asosq.driver.Dialect())
	t1 := builder.Table(appsubscriptiononeshot.Table)
	columns := asosq.ctx.Fields
	if len(columns) == 0 {
		columns = appsubscriptiononeshot.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if asosq.sql != nil {
		selector = asosq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if asosq.ctx.Unique != nil && *asosq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range asosq.modifiers {
		m(selector)
	}
	for _, p := range asosq.predicates {
		p(selector)
	}
	for _, p := range asosq.order {
		p(selector)
	}
	if offset := asosq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := asosq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (asosq *AppSubscriptionOneShotQuery) ForUpdate(opts ...sql.LockOption) *AppSubscriptionOneShotQuery {
	if asosq.driver.Dialect() == dialect.Postgres {
		asosq.Unique(false)
	}
	asosq.modifiers = append(asosq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return asosq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (asosq *AppSubscriptionOneShotQuery) ForShare(opts ...sql.LockOption) *AppSubscriptionOneShotQuery {
	if asosq.driver.Dialect() == dialect.Postgres {
		asosq.Unique(false)
	}
	asosq.modifiers = append(asosq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return asosq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (asosq *AppSubscriptionOneShotQuery) Modify(modifiers ...func(s *sql.Selector)) *AppSubscriptionOneShotSelect {
	asosq.modifiers = append(asosq.modifiers, modifiers...)
	return asosq.Select()
}

// AppSubscriptionOneShotGroupBy is the group-by builder for AppSubscriptionOneShot entities.
type AppSubscriptionOneShotGroupBy struct {
	selector
	build *AppSubscriptionOneShotQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (asosgb *AppSubscriptionOneShotGroupBy) Aggregate(fns ...AggregateFunc) *AppSubscriptionOneShotGroupBy {
	asosgb.fns = append(asosgb.fns, fns...)
	return asosgb
}

// Scan applies the selector query and scans the result into the given value.
func (asosgb *AppSubscriptionOneShotGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, asosgb.build.ctx, ent.OpQueryGroupBy)
	if err := asosgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppSubscriptionOneShotQuery, *AppSubscriptionOneShotGroupBy](ctx, asosgb.build, asosgb, asosgb.build.inters, v)
}

func (asosgb *AppSubscriptionOneShotGroupBy) sqlScan(ctx context.Context, root *AppSubscriptionOneShotQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(asosgb.fns))
	for _, fn := range asosgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*asosgb.flds)+len(asosgb.fns))
		for _, f := range *asosgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*asosgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := asosgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AppSubscriptionOneShotSelect is the builder for selecting fields of AppSubscriptionOneShot entities.
type AppSubscriptionOneShotSelect struct {
	*AppSubscriptionOneShotQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (asoss *AppSubscriptionOneShotSelect) Aggregate(fns ...AggregateFunc) *AppSubscriptionOneShotSelect {
	asoss.fns = append(asoss.fns, fns...)
	return asoss
}

// Scan applies the selector query and scans the result into the given value.
func (asoss *AppSubscriptionOneShotSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, asoss.ctx, ent.OpQuerySelect)
	if err := asoss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppSubscriptionOneShotQuery, *AppSubscriptionOneShotSelect](ctx, asoss.AppSubscriptionOneShotQuery, asoss, asoss.inters, v)
}

func (asoss *AppSubscriptionOneShotSelect) sqlScan(ctx context.Context, root *AppSubscriptionOneShotQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(asoss.fns))
	for _, fn := range asoss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*asoss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := asoss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (asoss *AppSubscriptionOneShotSelect) Modify(modifiers ...func(s *sql.Selector)) *AppSubscriptionOneShotSelect {
	asoss.modifiers = append(asoss.modifiers, modifiers...)
	return asoss
}
