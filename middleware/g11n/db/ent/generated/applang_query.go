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
	"github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated/applang"
	"github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated/predicate"
)

// AppLangQuery is the builder for querying AppLang entities.
type AppLangQuery struct {
	config
	ctx        *QueryContext
	order      []applang.OrderOption
	inters     []Interceptor
	predicates []predicate.AppLang
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppLangQuery builder.
func (alq *AppLangQuery) Where(ps ...predicate.AppLang) *AppLangQuery {
	alq.predicates = append(alq.predicates, ps...)
	return alq
}

// Limit the number of records to be returned by this query.
func (alq *AppLangQuery) Limit(limit int) *AppLangQuery {
	alq.ctx.Limit = &limit
	return alq
}

// Offset to start from.
func (alq *AppLangQuery) Offset(offset int) *AppLangQuery {
	alq.ctx.Offset = &offset
	return alq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (alq *AppLangQuery) Unique(unique bool) *AppLangQuery {
	alq.ctx.Unique = &unique
	return alq
}

// Order specifies how the records should be ordered.
func (alq *AppLangQuery) Order(o ...applang.OrderOption) *AppLangQuery {
	alq.order = append(alq.order, o...)
	return alq
}

// First returns the first AppLang entity from the query.
// Returns a *NotFoundError when no AppLang was found.
func (alq *AppLangQuery) First(ctx context.Context) (*AppLang, error) {
	nodes, err := alq.Limit(1).All(setContextOp(ctx, alq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{applang.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (alq *AppLangQuery) FirstX(ctx context.Context) *AppLang {
	node, err := alq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppLang ID from the query.
// Returns a *NotFoundError when no AppLang ID was found.
func (alq *AppLangQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = alq.Limit(1).IDs(setContextOp(ctx, alq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{applang.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (alq *AppLangQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := alq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppLang entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppLang entity is found.
// Returns a *NotFoundError when no AppLang entities are found.
func (alq *AppLangQuery) Only(ctx context.Context) (*AppLang, error) {
	nodes, err := alq.Limit(2).All(setContextOp(ctx, alq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{applang.Label}
	default:
		return nil, &NotSingularError{applang.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (alq *AppLangQuery) OnlyX(ctx context.Context) *AppLang {
	node, err := alq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppLang ID in the query.
// Returns a *NotSingularError when more than one AppLang ID is found.
// Returns a *NotFoundError when no entities are found.
func (alq *AppLangQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = alq.Limit(2).IDs(setContextOp(ctx, alq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{applang.Label}
	default:
		err = &NotSingularError{applang.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (alq *AppLangQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := alq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppLangs.
func (alq *AppLangQuery) All(ctx context.Context) ([]*AppLang, error) {
	ctx = setContextOp(ctx, alq.ctx, ent.OpQueryAll)
	if err := alq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AppLang, *AppLangQuery]()
	return withInterceptors[[]*AppLang](ctx, alq, qr, alq.inters)
}

// AllX is like All, but panics if an error occurs.
func (alq *AppLangQuery) AllX(ctx context.Context) []*AppLang {
	nodes, err := alq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppLang IDs.
func (alq *AppLangQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if alq.ctx.Unique == nil && alq.path != nil {
		alq.Unique(true)
	}
	ctx = setContextOp(ctx, alq.ctx, ent.OpQueryIDs)
	if err = alq.Select(applang.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (alq *AppLangQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := alq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (alq *AppLangQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, alq.ctx, ent.OpQueryCount)
	if err := alq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, alq, querierCount[*AppLangQuery](), alq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (alq *AppLangQuery) CountX(ctx context.Context) int {
	count, err := alq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (alq *AppLangQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, alq.ctx, ent.OpQueryExist)
	switch _, err := alq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (alq *AppLangQuery) ExistX(ctx context.Context) bool {
	exist, err := alq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppLangQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (alq *AppLangQuery) Clone() *AppLangQuery {
	if alq == nil {
		return nil
	}
	return &AppLangQuery{
		config:     alq.config,
		ctx:        alq.ctx.Clone(),
		order:      append([]applang.OrderOption{}, alq.order...),
		inters:     append([]Interceptor{}, alq.inters...),
		predicates: append([]predicate.AppLang{}, alq.predicates...),
		// clone intermediate query.
		sql:       alq.sql.Clone(),
		path:      alq.path,
		modifiers: append([]func(*sql.Selector){}, alq.modifiers...),
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
//	client.AppLang.Query().
//		GroupBy(applang.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (alq *AppLangQuery) GroupBy(field string, fields ...string) *AppLangGroupBy {
	alq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AppLangGroupBy{build: alq}
	grbuild.flds = &alq.ctx.Fields
	grbuild.label = applang.Label
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
//	client.AppLang.Query().
//		Select(applang.FieldCreatedAt).
//		Scan(ctx, &v)
func (alq *AppLangQuery) Select(fields ...string) *AppLangSelect {
	alq.ctx.Fields = append(alq.ctx.Fields, fields...)
	sbuild := &AppLangSelect{AppLangQuery: alq}
	sbuild.label = applang.Label
	sbuild.flds, sbuild.scan = &alq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AppLangSelect configured with the given aggregations.
func (alq *AppLangQuery) Aggregate(fns ...AggregateFunc) *AppLangSelect {
	return alq.Select().Aggregate(fns...)
}

func (alq *AppLangQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range alq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, alq); err != nil {
				return err
			}
		}
	}
	for _, f := range alq.ctx.Fields {
		if !applang.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if alq.path != nil {
		prev, err := alq.path(ctx)
		if err != nil {
			return err
		}
		alq.sql = prev
	}
	return nil
}

func (alq *AppLangQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppLang, error) {
	var (
		nodes = []*AppLang{}
		_spec = alq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AppLang).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AppLang{config: alq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(alq.modifiers) > 0 {
		_spec.Modifiers = alq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, alq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (alq *AppLangQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := alq.querySpec()
	if len(alq.modifiers) > 0 {
		_spec.Modifiers = alq.modifiers
	}
	_spec.Node.Columns = alq.ctx.Fields
	if len(alq.ctx.Fields) > 0 {
		_spec.Unique = alq.ctx.Unique != nil && *alq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, alq.driver, _spec)
}

func (alq *AppLangQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(applang.Table, applang.Columns, sqlgraph.NewFieldSpec(applang.FieldID, field.TypeUint32))
	_spec.From = alq.sql
	if unique := alq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if alq.path != nil {
		_spec.Unique = true
	}
	if fields := alq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, applang.FieldID)
		for i := range fields {
			if fields[i] != applang.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := alq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := alq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := alq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := alq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (alq *AppLangQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(alq.driver.Dialect())
	t1 := builder.Table(applang.Table)
	columns := alq.ctx.Fields
	if len(columns) == 0 {
		columns = applang.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if alq.sql != nil {
		selector = alq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if alq.ctx.Unique != nil && *alq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range alq.modifiers {
		m(selector)
	}
	for _, p := range alq.predicates {
		p(selector)
	}
	for _, p := range alq.order {
		p(selector)
	}
	if offset := alq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := alq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (alq *AppLangQuery) ForUpdate(opts ...sql.LockOption) *AppLangQuery {
	if alq.driver.Dialect() == dialect.Postgres {
		alq.Unique(false)
	}
	alq.modifiers = append(alq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return alq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (alq *AppLangQuery) ForShare(opts ...sql.LockOption) *AppLangQuery {
	if alq.driver.Dialect() == dialect.Postgres {
		alq.Unique(false)
	}
	alq.modifiers = append(alq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return alq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (alq *AppLangQuery) Modify(modifiers ...func(s *sql.Selector)) *AppLangSelect {
	alq.modifiers = append(alq.modifiers, modifiers...)
	return alq.Select()
}

// AppLangGroupBy is the group-by builder for AppLang entities.
type AppLangGroupBy struct {
	selector
	build *AppLangQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (algb *AppLangGroupBy) Aggregate(fns ...AggregateFunc) *AppLangGroupBy {
	algb.fns = append(algb.fns, fns...)
	return algb
}

// Scan applies the selector query and scans the result into the given value.
func (algb *AppLangGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, algb.build.ctx, ent.OpQueryGroupBy)
	if err := algb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppLangQuery, *AppLangGroupBy](ctx, algb.build, algb, algb.build.inters, v)
}

func (algb *AppLangGroupBy) sqlScan(ctx context.Context, root *AppLangQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(algb.fns))
	for _, fn := range algb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*algb.flds)+len(algb.fns))
		for _, f := range *algb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*algb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := algb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AppLangSelect is the builder for selecting fields of AppLang entities.
type AppLangSelect struct {
	*AppLangQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (als *AppLangSelect) Aggregate(fns ...AggregateFunc) *AppLangSelect {
	als.fns = append(als.fns, fns...)
	return als
}

// Scan applies the selector query and scans the result into the given value.
func (als *AppLangSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, als.ctx, ent.OpQuerySelect)
	if err := als.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppLangQuery, *AppLangSelect](ctx, als.AppLangQuery, als, als.inters, v)
}

func (als *AppLangSelect) sqlScan(ctx context.Context, root *AppLangQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(als.fns))
	for _, fn := range als.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*als.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := als.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (als *AppLangSelect) Modify(modifiers ...func(s *sql.Selector)) *AppLangSelect {
	als.modifiers = append(als.modifiers, modifiers...)
	return als
}
