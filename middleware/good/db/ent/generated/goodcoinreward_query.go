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
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodcoinreward"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/predicate"
)

// GoodCoinRewardQuery is the builder for querying GoodCoinReward entities.
type GoodCoinRewardQuery struct {
	config
	ctx        *QueryContext
	order      []goodcoinreward.OrderOption
	inters     []Interceptor
	predicates []predicate.GoodCoinReward
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GoodCoinRewardQuery builder.
func (gcrq *GoodCoinRewardQuery) Where(ps ...predicate.GoodCoinReward) *GoodCoinRewardQuery {
	gcrq.predicates = append(gcrq.predicates, ps...)
	return gcrq
}

// Limit the number of records to be returned by this query.
func (gcrq *GoodCoinRewardQuery) Limit(limit int) *GoodCoinRewardQuery {
	gcrq.ctx.Limit = &limit
	return gcrq
}

// Offset to start from.
func (gcrq *GoodCoinRewardQuery) Offset(offset int) *GoodCoinRewardQuery {
	gcrq.ctx.Offset = &offset
	return gcrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gcrq *GoodCoinRewardQuery) Unique(unique bool) *GoodCoinRewardQuery {
	gcrq.ctx.Unique = &unique
	return gcrq
}

// Order specifies how the records should be ordered.
func (gcrq *GoodCoinRewardQuery) Order(o ...goodcoinreward.OrderOption) *GoodCoinRewardQuery {
	gcrq.order = append(gcrq.order, o...)
	return gcrq
}

// First returns the first GoodCoinReward entity from the query.
// Returns a *NotFoundError when no GoodCoinReward was found.
func (gcrq *GoodCoinRewardQuery) First(ctx context.Context) (*GoodCoinReward, error) {
	nodes, err := gcrq.Limit(1).All(setContextOp(ctx, gcrq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{goodcoinreward.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gcrq *GoodCoinRewardQuery) FirstX(ctx context.Context) *GoodCoinReward {
	node, err := gcrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GoodCoinReward ID from the query.
// Returns a *NotFoundError when no GoodCoinReward ID was found.
func (gcrq *GoodCoinRewardQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = gcrq.Limit(1).IDs(setContextOp(ctx, gcrq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{goodcoinreward.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gcrq *GoodCoinRewardQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := gcrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GoodCoinReward entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GoodCoinReward entity is found.
// Returns a *NotFoundError when no GoodCoinReward entities are found.
func (gcrq *GoodCoinRewardQuery) Only(ctx context.Context) (*GoodCoinReward, error) {
	nodes, err := gcrq.Limit(2).All(setContextOp(ctx, gcrq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{goodcoinreward.Label}
	default:
		return nil, &NotSingularError{goodcoinreward.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gcrq *GoodCoinRewardQuery) OnlyX(ctx context.Context) *GoodCoinReward {
	node, err := gcrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GoodCoinReward ID in the query.
// Returns a *NotSingularError when more than one GoodCoinReward ID is found.
// Returns a *NotFoundError when no entities are found.
func (gcrq *GoodCoinRewardQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = gcrq.Limit(2).IDs(setContextOp(ctx, gcrq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{goodcoinreward.Label}
	default:
		err = &NotSingularError{goodcoinreward.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gcrq *GoodCoinRewardQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := gcrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GoodCoinRewards.
func (gcrq *GoodCoinRewardQuery) All(ctx context.Context) ([]*GoodCoinReward, error) {
	ctx = setContextOp(ctx, gcrq.ctx, ent.OpQueryAll)
	if err := gcrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GoodCoinReward, *GoodCoinRewardQuery]()
	return withInterceptors[[]*GoodCoinReward](ctx, gcrq, qr, gcrq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gcrq *GoodCoinRewardQuery) AllX(ctx context.Context) []*GoodCoinReward {
	nodes, err := gcrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GoodCoinReward IDs.
func (gcrq *GoodCoinRewardQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if gcrq.ctx.Unique == nil && gcrq.path != nil {
		gcrq.Unique(true)
	}
	ctx = setContextOp(ctx, gcrq.ctx, ent.OpQueryIDs)
	if err = gcrq.Select(goodcoinreward.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gcrq *GoodCoinRewardQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := gcrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gcrq *GoodCoinRewardQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, gcrq.ctx, ent.OpQueryCount)
	if err := gcrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gcrq, querierCount[*GoodCoinRewardQuery](), gcrq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gcrq *GoodCoinRewardQuery) CountX(ctx context.Context) int {
	count, err := gcrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gcrq *GoodCoinRewardQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, gcrq.ctx, ent.OpQueryExist)
	switch _, err := gcrq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gcrq *GoodCoinRewardQuery) ExistX(ctx context.Context) bool {
	exist, err := gcrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GoodCoinRewardQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gcrq *GoodCoinRewardQuery) Clone() *GoodCoinRewardQuery {
	if gcrq == nil {
		return nil
	}
	return &GoodCoinRewardQuery{
		config:     gcrq.config,
		ctx:        gcrq.ctx.Clone(),
		order:      append([]goodcoinreward.OrderOption{}, gcrq.order...),
		inters:     append([]Interceptor{}, gcrq.inters...),
		predicates: append([]predicate.GoodCoinReward{}, gcrq.predicates...),
		// clone intermediate query.
		sql:       gcrq.sql.Clone(),
		path:      gcrq.path,
		modifiers: append([]func(*sql.Selector){}, gcrq.modifiers...),
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
//	client.GoodCoinReward.Query().
//		GroupBy(goodcoinreward.FieldEntID).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (gcrq *GoodCoinRewardQuery) GroupBy(field string, fields ...string) *GoodCoinRewardGroupBy {
	gcrq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GoodCoinRewardGroupBy{build: gcrq}
	grbuild.flds = &gcrq.ctx.Fields
	grbuild.label = goodcoinreward.Label
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
//	client.GoodCoinReward.Query().
//		Select(goodcoinreward.FieldEntID).
//		Scan(ctx, &v)
func (gcrq *GoodCoinRewardQuery) Select(fields ...string) *GoodCoinRewardSelect {
	gcrq.ctx.Fields = append(gcrq.ctx.Fields, fields...)
	sbuild := &GoodCoinRewardSelect{GoodCoinRewardQuery: gcrq}
	sbuild.label = goodcoinreward.Label
	sbuild.flds, sbuild.scan = &gcrq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GoodCoinRewardSelect configured with the given aggregations.
func (gcrq *GoodCoinRewardQuery) Aggregate(fns ...AggregateFunc) *GoodCoinRewardSelect {
	return gcrq.Select().Aggregate(fns...)
}

func (gcrq *GoodCoinRewardQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gcrq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gcrq); err != nil {
				return err
			}
		}
	}
	for _, f := range gcrq.ctx.Fields {
		if !goodcoinreward.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if gcrq.path != nil {
		prev, err := gcrq.path(ctx)
		if err != nil {
			return err
		}
		gcrq.sql = prev
	}
	return nil
}

func (gcrq *GoodCoinRewardQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GoodCoinReward, error) {
	var (
		nodes = []*GoodCoinReward{}
		_spec = gcrq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GoodCoinReward).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GoodCoinReward{config: gcrq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(gcrq.modifiers) > 0 {
		_spec.Modifiers = gcrq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gcrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (gcrq *GoodCoinRewardQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gcrq.querySpec()
	if len(gcrq.modifiers) > 0 {
		_spec.Modifiers = gcrq.modifiers
	}
	_spec.Node.Columns = gcrq.ctx.Fields
	if len(gcrq.ctx.Fields) > 0 {
		_spec.Unique = gcrq.ctx.Unique != nil && *gcrq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, gcrq.driver, _spec)
}

func (gcrq *GoodCoinRewardQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(goodcoinreward.Table, goodcoinreward.Columns, sqlgraph.NewFieldSpec(goodcoinreward.FieldID, field.TypeUint32))
	_spec.From = gcrq.sql
	if unique := gcrq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if gcrq.path != nil {
		_spec.Unique = true
	}
	if fields := gcrq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodcoinreward.FieldID)
		for i := range fields {
			if fields[i] != goodcoinreward.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gcrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gcrq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gcrq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gcrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gcrq *GoodCoinRewardQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gcrq.driver.Dialect())
	t1 := builder.Table(goodcoinreward.Table)
	columns := gcrq.ctx.Fields
	if len(columns) == 0 {
		columns = goodcoinreward.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gcrq.sql != nil {
		selector = gcrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gcrq.ctx.Unique != nil && *gcrq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range gcrq.modifiers {
		m(selector)
	}
	for _, p := range gcrq.predicates {
		p(selector)
	}
	for _, p := range gcrq.order {
		p(selector)
	}
	if offset := gcrq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gcrq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (gcrq *GoodCoinRewardQuery) ForUpdate(opts ...sql.LockOption) *GoodCoinRewardQuery {
	if gcrq.driver.Dialect() == dialect.Postgres {
		gcrq.Unique(false)
	}
	gcrq.modifiers = append(gcrq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return gcrq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (gcrq *GoodCoinRewardQuery) ForShare(opts ...sql.LockOption) *GoodCoinRewardQuery {
	if gcrq.driver.Dialect() == dialect.Postgres {
		gcrq.Unique(false)
	}
	gcrq.modifiers = append(gcrq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return gcrq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gcrq *GoodCoinRewardQuery) Modify(modifiers ...func(s *sql.Selector)) *GoodCoinRewardSelect {
	gcrq.modifiers = append(gcrq.modifiers, modifiers...)
	return gcrq.Select()
}

// GoodCoinRewardGroupBy is the group-by builder for GoodCoinReward entities.
type GoodCoinRewardGroupBy struct {
	selector
	build *GoodCoinRewardQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gcrgb *GoodCoinRewardGroupBy) Aggregate(fns ...AggregateFunc) *GoodCoinRewardGroupBy {
	gcrgb.fns = append(gcrgb.fns, fns...)
	return gcrgb
}

// Scan applies the selector query and scans the result into the given value.
func (gcrgb *GoodCoinRewardGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gcrgb.build.ctx, ent.OpQueryGroupBy)
	if err := gcrgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GoodCoinRewardQuery, *GoodCoinRewardGroupBy](ctx, gcrgb.build, gcrgb, gcrgb.build.inters, v)
}

func (gcrgb *GoodCoinRewardGroupBy) sqlScan(ctx context.Context, root *GoodCoinRewardQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(gcrgb.fns))
	for _, fn := range gcrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*gcrgb.flds)+len(gcrgb.fns))
		for _, f := range *gcrgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*gcrgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gcrgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GoodCoinRewardSelect is the builder for selecting fields of GoodCoinReward entities.
type GoodCoinRewardSelect struct {
	*GoodCoinRewardQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gcrs *GoodCoinRewardSelect) Aggregate(fns ...AggregateFunc) *GoodCoinRewardSelect {
	gcrs.fns = append(gcrs.fns, fns...)
	return gcrs
}

// Scan applies the selector query and scans the result into the given value.
func (gcrs *GoodCoinRewardSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gcrs.ctx, ent.OpQuerySelect)
	if err := gcrs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GoodCoinRewardQuery, *GoodCoinRewardSelect](ctx, gcrs.GoodCoinRewardQuery, gcrs, gcrs.inters, v)
}

func (gcrs *GoodCoinRewardSelect) sqlScan(ctx context.Context, root *GoodCoinRewardQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gcrs.fns))
	for _, fn := range gcrs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gcrs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gcrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gcrs *GoodCoinRewardSelect) Modify(modifiers ...func(s *sql.Selector)) *GoodCoinRewardSelect {
	gcrs.modifiers = append(gcrs.modifiers, modifiers...)
	return gcrs
}
