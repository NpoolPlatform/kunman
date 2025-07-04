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
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/goodachievement"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/predicate"
)

// GoodAchievementQuery is the builder for querying GoodAchievement entities.
type GoodAchievementQuery struct {
	config
	ctx        *QueryContext
	order      []goodachievement.OrderOption
	inters     []Interceptor
	predicates []predicate.GoodAchievement
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GoodAchievementQuery builder.
func (gaq *GoodAchievementQuery) Where(ps ...predicate.GoodAchievement) *GoodAchievementQuery {
	gaq.predicates = append(gaq.predicates, ps...)
	return gaq
}

// Limit the number of records to be returned by this query.
func (gaq *GoodAchievementQuery) Limit(limit int) *GoodAchievementQuery {
	gaq.ctx.Limit = &limit
	return gaq
}

// Offset to start from.
func (gaq *GoodAchievementQuery) Offset(offset int) *GoodAchievementQuery {
	gaq.ctx.Offset = &offset
	return gaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gaq *GoodAchievementQuery) Unique(unique bool) *GoodAchievementQuery {
	gaq.ctx.Unique = &unique
	return gaq
}

// Order specifies how the records should be ordered.
func (gaq *GoodAchievementQuery) Order(o ...goodachievement.OrderOption) *GoodAchievementQuery {
	gaq.order = append(gaq.order, o...)
	return gaq
}

// First returns the first GoodAchievement entity from the query.
// Returns a *NotFoundError when no GoodAchievement was found.
func (gaq *GoodAchievementQuery) First(ctx context.Context) (*GoodAchievement, error) {
	nodes, err := gaq.Limit(1).All(setContextOp(ctx, gaq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{goodachievement.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gaq *GoodAchievementQuery) FirstX(ctx context.Context) *GoodAchievement {
	node, err := gaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GoodAchievement ID from the query.
// Returns a *NotFoundError when no GoodAchievement ID was found.
func (gaq *GoodAchievementQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = gaq.Limit(1).IDs(setContextOp(ctx, gaq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{goodachievement.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gaq *GoodAchievementQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := gaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GoodAchievement entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GoodAchievement entity is found.
// Returns a *NotFoundError when no GoodAchievement entities are found.
func (gaq *GoodAchievementQuery) Only(ctx context.Context) (*GoodAchievement, error) {
	nodes, err := gaq.Limit(2).All(setContextOp(ctx, gaq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{goodachievement.Label}
	default:
		return nil, &NotSingularError{goodachievement.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gaq *GoodAchievementQuery) OnlyX(ctx context.Context) *GoodAchievement {
	node, err := gaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GoodAchievement ID in the query.
// Returns a *NotSingularError when more than one GoodAchievement ID is found.
// Returns a *NotFoundError when no entities are found.
func (gaq *GoodAchievementQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = gaq.Limit(2).IDs(setContextOp(ctx, gaq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{goodachievement.Label}
	default:
		err = &NotSingularError{goodachievement.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gaq *GoodAchievementQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := gaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GoodAchievements.
func (gaq *GoodAchievementQuery) All(ctx context.Context) ([]*GoodAchievement, error) {
	ctx = setContextOp(ctx, gaq.ctx, ent.OpQueryAll)
	if err := gaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GoodAchievement, *GoodAchievementQuery]()
	return withInterceptors[[]*GoodAchievement](ctx, gaq, qr, gaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gaq *GoodAchievementQuery) AllX(ctx context.Context) []*GoodAchievement {
	nodes, err := gaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GoodAchievement IDs.
func (gaq *GoodAchievementQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if gaq.ctx.Unique == nil && gaq.path != nil {
		gaq.Unique(true)
	}
	ctx = setContextOp(ctx, gaq.ctx, ent.OpQueryIDs)
	if err = gaq.Select(goodachievement.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gaq *GoodAchievementQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := gaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gaq *GoodAchievementQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, gaq.ctx, ent.OpQueryCount)
	if err := gaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gaq, querierCount[*GoodAchievementQuery](), gaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gaq *GoodAchievementQuery) CountX(ctx context.Context) int {
	count, err := gaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gaq *GoodAchievementQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, gaq.ctx, ent.OpQueryExist)
	switch _, err := gaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gaq *GoodAchievementQuery) ExistX(ctx context.Context) bool {
	exist, err := gaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GoodAchievementQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gaq *GoodAchievementQuery) Clone() *GoodAchievementQuery {
	if gaq == nil {
		return nil
	}
	return &GoodAchievementQuery{
		config:     gaq.config,
		ctx:        gaq.ctx.Clone(),
		order:      append([]goodachievement.OrderOption{}, gaq.order...),
		inters:     append([]Interceptor{}, gaq.inters...),
		predicates: append([]predicate.GoodAchievement{}, gaq.predicates...),
		// clone intermediate query.
		sql:       gaq.sql.Clone(),
		path:      gaq.path,
		modifiers: append([]func(*sql.Selector){}, gaq.modifiers...),
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
//	client.GoodAchievement.Query().
//		GroupBy(goodachievement.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (gaq *GoodAchievementQuery) GroupBy(field string, fields ...string) *GoodAchievementGroupBy {
	gaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GoodAchievementGroupBy{build: gaq}
	grbuild.flds = &gaq.ctx.Fields
	grbuild.label = goodachievement.Label
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
//	client.GoodAchievement.Query().
//		Select(goodachievement.FieldCreatedAt).
//		Scan(ctx, &v)
func (gaq *GoodAchievementQuery) Select(fields ...string) *GoodAchievementSelect {
	gaq.ctx.Fields = append(gaq.ctx.Fields, fields...)
	sbuild := &GoodAchievementSelect{GoodAchievementQuery: gaq}
	sbuild.label = goodachievement.Label
	sbuild.flds, sbuild.scan = &gaq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GoodAchievementSelect configured with the given aggregations.
func (gaq *GoodAchievementQuery) Aggregate(fns ...AggregateFunc) *GoodAchievementSelect {
	return gaq.Select().Aggregate(fns...)
}

func (gaq *GoodAchievementQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gaq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gaq); err != nil {
				return err
			}
		}
	}
	for _, f := range gaq.ctx.Fields {
		if !goodachievement.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if gaq.path != nil {
		prev, err := gaq.path(ctx)
		if err != nil {
			return err
		}
		gaq.sql = prev
	}
	return nil
}

func (gaq *GoodAchievementQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GoodAchievement, error) {
	var (
		nodes = []*GoodAchievement{}
		_spec = gaq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GoodAchievement).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GoodAchievement{config: gaq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(gaq.modifiers) > 0 {
		_spec.Modifiers = gaq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (gaq *GoodAchievementQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gaq.querySpec()
	if len(gaq.modifiers) > 0 {
		_spec.Modifiers = gaq.modifiers
	}
	_spec.Node.Columns = gaq.ctx.Fields
	if len(gaq.ctx.Fields) > 0 {
		_spec.Unique = gaq.ctx.Unique != nil && *gaq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, gaq.driver, _spec)
}

func (gaq *GoodAchievementQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(goodachievement.Table, goodachievement.Columns, sqlgraph.NewFieldSpec(goodachievement.FieldID, field.TypeUint32))
	_spec.From = gaq.sql
	if unique := gaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if gaq.path != nil {
		_spec.Unique = true
	}
	if fields := gaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodachievement.FieldID)
		for i := range fields {
			if fields[i] != goodachievement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gaq *GoodAchievementQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gaq.driver.Dialect())
	t1 := builder.Table(goodachievement.Table)
	columns := gaq.ctx.Fields
	if len(columns) == 0 {
		columns = goodachievement.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gaq.sql != nil {
		selector = gaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gaq.ctx.Unique != nil && *gaq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range gaq.modifiers {
		m(selector)
	}
	for _, p := range gaq.predicates {
		p(selector)
	}
	for _, p := range gaq.order {
		p(selector)
	}
	if offset := gaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (gaq *GoodAchievementQuery) ForUpdate(opts ...sql.LockOption) *GoodAchievementQuery {
	if gaq.driver.Dialect() == dialect.Postgres {
		gaq.Unique(false)
	}
	gaq.modifiers = append(gaq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return gaq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (gaq *GoodAchievementQuery) ForShare(opts ...sql.LockOption) *GoodAchievementQuery {
	if gaq.driver.Dialect() == dialect.Postgres {
		gaq.Unique(false)
	}
	gaq.modifiers = append(gaq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return gaq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gaq *GoodAchievementQuery) Modify(modifiers ...func(s *sql.Selector)) *GoodAchievementSelect {
	gaq.modifiers = append(gaq.modifiers, modifiers...)
	return gaq.Select()
}

// GoodAchievementGroupBy is the group-by builder for GoodAchievement entities.
type GoodAchievementGroupBy struct {
	selector
	build *GoodAchievementQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gagb *GoodAchievementGroupBy) Aggregate(fns ...AggregateFunc) *GoodAchievementGroupBy {
	gagb.fns = append(gagb.fns, fns...)
	return gagb
}

// Scan applies the selector query and scans the result into the given value.
func (gagb *GoodAchievementGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gagb.build.ctx, ent.OpQueryGroupBy)
	if err := gagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GoodAchievementQuery, *GoodAchievementGroupBy](ctx, gagb.build, gagb, gagb.build.inters, v)
}

func (gagb *GoodAchievementGroupBy) sqlScan(ctx context.Context, root *GoodAchievementQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(gagb.fns))
	for _, fn := range gagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*gagb.flds)+len(gagb.fns))
		for _, f := range *gagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*gagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GoodAchievementSelect is the builder for selecting fields of GoodAchievement entities.
type GoodAchievementSelect struct {
	*GoodAchievementQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gas *GoodAchievementSelect) Aggregate(fns ...AggregateFunc) *GoodAchievementSelect {
	gas.fns = append(gas.fns, fns...)
	return gas
}

// Scan applies the selector query and scans the result into the given value.
func (gas *GoodAchievementSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gas.ctx, ent.OpQuerySelect)
	if err := gas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GoodAchievementQuery, *GoodAchievementSelect](ctx, gas.GoodAchievementQuery, gas, gas.inters, v)
}

func (gas *GoodAchievementSelect) sqlScan(ctx context.Context, root *GoodAchievementQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gas.fns))
	for _, fn := range gas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gas *GoodAchievementSelect) Modify(modifiers ...func(s *sql.Selector)) *GoodAchievementSelect {
	gas.modifiers = append(gas.modifiers, modifiers...)
	return gas
}
