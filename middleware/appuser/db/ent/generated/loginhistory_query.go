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
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/loginhistory"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/predicate"
)

// LoginHistoryQuery is the builder for querying LoginHistory entities.
type LoginHistoryQuery struct {
	config
	ctx        *QueryContext
	order      []loginhistory.OrderOption
	inters     []Interceptor
	predicates []predicate.LoginHistory
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LoginHistoryQuery builder.
func (lhq *LoginHistoryQuery) Where(ps ...predicate.LoginHistory) *LoginHistoryQuery {
	lhq.predicates = append(lhq.predicates, ps...)
	return lhq
}

// Limit the number of records to be returned by this query.
func (lhq *LoginHistoryQuery) Limit(limit int) *LoginHistoryQuery {
	lhq.ctx.Limit = &limit
	return lhq
}

// Offset to start from.
func (lhq *LoginHistoryQuery) Offset(offset int) *LoginHistoryQuery {
	lhq.ctx.Offset = &offset
	return lhq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lhq *LoginHistoryQuery) Unique(unique bool) *LoginHistoryQuery {
	lhq.ctx.Unique = &unique
	return lhq
}

// Order specifies how the records should be ordered.
func (lhq *LoginHistoryQuery) Order(o ...loginhistory.OrderOption) *LoginHistoryQuery {
	lhq.order = append(lhq.order, o...)
	return lhq
}

// First returns the first LoginHistory entity from the query.
// Returns a *NotFoundError when no LoginHistory was found.
func (lhq *LoginHistoryQuery) First(ctx context.Context) (*LoginHistory, error) {
	nodes, err := lhq.Limit(1).All(setContextOp(ctx, lhq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{loginhistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lhq *LoginHistoryQuery) FirstX(ctx context.Context) *LoginHistory {
	node, err := lhq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LoginHistory ID from the query.
// Returns a *NotFoundError when no LoginHistory ID was found.
func (lhq *LoginHistoryQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = lhq.Limit(1).IDs(setContextOp(ctx, lhq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{loginhistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lhq *LoginHistoryQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := lhq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LoginHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LoginHistory entity is found.
// Returns a *NotFoundError when no LoginHistory entities are found.
func (lhq *LoginHistoryQuery) Only(ctx context.Context) (*LoginHistory, error) {
	nodes, err := lhq.Limit(2).All(setContextOp(ctx, lhq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{loginhistory.Label}
	default:
		return nil, &NotSingularError{loginhistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lhq *LoginHistoryQuery) OnlyX(ctx context.Context) *LoginHistory {
	node, err := lhq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LoginHistory ID in the query.
// Returns a *NotSingularError when more than one LoginHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (lhq *LoginHistoryQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = lhq.Limit(2).IDs(setContextOp(ctx, lhq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{loginhistory.Label}
	default:
		err = &NotSingularError{loginhistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lhq *LoginHistoryQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := lhq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LoginHistories.
func (lhq *LoginHistoryQuery) All(ctx context.Context) ([]*LoginHistory, error) {
	ctx = setContextOp(ctx, lhq.ctx, ent.OpQueryAll)
	if err := lhq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LoginHistory, *LoginHistoryQuery]()
	return withInterceptors[[]*LoginHistory](ctx, lhq, qr, lhq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lhq *LoginHistoryQuery) AllX(ctx context.Context) []*LoginHistory {
	nodes, err := lhq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LoginHistory IDs.
func (lhq *LoginHistoryQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if lhq.ctx.Unique == nil && lhq.path != nil {
		lhq.Unique(true)
	}
	ctx = setContextOp(ctx, lhq.ctx, ent.OpQueryIDs)
	if err = lhq.Select(loginhistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lhq *LoginHistoryQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := lhq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lhq *LoginHistoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lhq.ctx, ent.OpQueryCount)
	if err := lhq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lhq, querierCount[*LoginHistoryQuery](), lhq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lhq *LoginHistoryQuery) CountX(ctx context.Context) int {
	count, err := lhq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lhq *LoginHistoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lhq.ctx, ent.OpQueryExist)
	switch _, err := lhq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lhq *LoginHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := lhq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LoginHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lhq *LoginHistoryQuery) Clone() *LoginHistoryQuery {
	if lhq == nil {
		return nil
	}
	return &LoginHistoryQuery{
		config:     lhq.config,
		ctx:        lhq.ctx.Clone(),
		order:      append([]loginhistory.OrderOption{}, lhq.order...),
		inters:     append([]Interceptor{}, lhq.inters...),
		predicates: append([]predicate.LoginHistory{}, lhq.predicates...),
		// clone intermediate query.
		sql:       lhq.sql.Clone(),
		path:      lhq.path,
		modifiers: append([]func(*sql.Selector){}, lhq.modifiers...),
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
//	client.LoginHistory.Query().
//		GroupBy(loginhistory.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (lhq *LoginHistoryQuery) GroupBy(field string, fields ...string) *LoginHistoryGroupBy {
	lhq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LoginHistoryGroupBy{build: lhq}
	grbuild.flds = &lhq.ctx.Fields
	grbuild.label = loginhistory.Label
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
//	client.LoginHistory.Query().
//		Select(loginhistory.FieldCreatedAt).
//		Scan(ctx, &v)
func (lhq *LoginHistoryQuery) Select(fields ...string) *LoginHistorySelect {
	lhq.ctx.Fields = append(lhq.ctx.Fields, fields...)
	sbuild := &LoginHistorySelect{LoginHistoryQuery: lhq}
	sbuild.label = loginhistory.Label
	sbuild.flds, sbuild.scan = &lhq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LoginHistorySelect configured with the given aggregations.
func (lhq *LoginHistoryQuery) Aggregate(fns ...AggregateFunc) *LoginHistorySelect {
	return lhq.Select().Aggregate(fns...)
}

func (lhq *LoginHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lhq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lhq); err != nil {
				return err
			}
		}
	}
	for _, f := range lhq.ctx.Fields {
		if !loginhistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if lhq.path != nil {
		prev, err := lhq.path(ctx)
		if err != nil {
			return err
		}
		lhq.sql = prev
	}
	return nil
}

func (lhq *LoginHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LoginHistory, error) {
	var (
		nodes = []*LoginHistory{}
		_spec = lhq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LoginHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LoginHistory{config: lhq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(lhq.modifiers) > 0 {
		_spec.Modifiers = lhq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lhq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (lhq *LoginHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lhq.querySpec()
	if len(lhq.modifiers) > 0 {
		_spec.Modifiers = lhq.modifiers
	}
	_spec.Node.Columns = lhq.ctx.Fields
	if len(lhq.ctx.Fields) > 0 {
		_spec.Unique = lhq.ctx.Unique != nil && *lhq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lhq.driver, _spec)
}

func (lhq *LoginHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(loginhistory.Table, loginhistory.Columns, sqlgraph.NewFieldSpec(loginhistory.FieldID, field.TypeUint32))
	_spec.From = lhq.sql
	if unique := lhq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lhq.path != nil {
		_spec.Unique = true
	}
	if fields := lhq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, loginhistory.FieldID)
		for i := range fields {
			if fields[i] != loginhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lhq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lhq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lhq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lhq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lhq *LoginHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lhq.driver.Dialect())
	t1 := builder.Table(loginhistory.Table)
	columns := lhq.ctx.Fields
	if len(columns) == 0 {
		columns = loginhistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lhq.sql != nil {
		selector = lhq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lhq.ctx.Unique != nil && *lhq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range lhq.modifiers {
		m(selector)
	}
	for _, p := range lhq.predicates {
		p(selector)
	}
	for _, p := range lhq.order {
		p(selector)
	}
	if offset := lhq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lhq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (lhq *LoginHistoryQuery) ForUpdate(opts ...sql.LockOption) *LoginHistoryQuery {
	if lhq.driver.Dialect() == dialect.Postgres {
		lhq.Unique(false)
	}
	lhq.modifiers = append(lhq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return lhq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (lhq *LoginHistoryQuery) ForShare(opts ...sql.LockOption) *LoginHistoryQuery {
	if lhq.driver.Dialect() == dialect.Postgres {
		lhq.Unique(false)
	}
	lhq.modifiers = append(lhq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return lhq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (lhq *LoginHistoryQuery) Modify(modifiers ...func(s *sql.Selector)) *LoginHistorySelect {
	lhq.modifiers = append(lhq.modifiers, modifiers...)
	return lhq.Select()
}

// LoginHistoryGroupBy is the group-by builder for LoginHistory entities.
type LoginHistoryGroupBy struct {
	selector
	build *LoginHistoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lhgb *LoginHistoryGroupBy) Aggregate(fns ...AggregateFunc) *LoginHistoryGroupBy {
	lhgb.fns = append(lhgb.fns, fns...)
	return lhgb
}

// Scan applies the selector query and scans the result into the given value.
func (lhgb *LoginHistoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lhgb.build.ctx, ent.OpQueryGroupBy)
	if err := lhgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LoginHistoryQuery, *LoginHistoryGroupBy](ctx, lhgb.build, lhgb, lhgb.build.inters, v)
}

func (lhgb *LoginHistoryGroupBy) sqlScan(ctx context.Context, root *LoginHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lhgb.fns))
	for _, fn := range lhgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lhgb.flds)+len(lhgb.fns))
		for _, f := range *lhgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lhgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lhgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LoginHistorySelect is the builder for selecting fields of LoginHistory entities.
type LoginHistorySelect struct {
	*LoginHistoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (lhs *LoginHistorySelect) Aggregate(fns ...AggregateFunc) *LoginHistorySelect {
	lhs.fns = append(lhs.fns, fns...)
	return lhs
}

// Scan applies the selector query and scans the result into the given value.
func (lhs *LoginHistorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lhs.ctx, ent.OpQuerySelect)
	if err := lhs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LoginHistoryQuery, *LoginHistorySelect](ctx, lhs.LoginHistoryQuery, lhs, lhs.inters, v)
}

func (lhs *LoginHistorySelect) sqlScan(ctx context.Context, root *LoginHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(lhs.fns))
	for _, fn := range lhs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*lhs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lhs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (lhs *LoginHistorySelect) Modify(modifiers ...func(s *sql.Selector)) *LoginHistorySelect {
	lhs.modifiers = append(lhs.modifiers, modifiers...)
	return lhs
}
