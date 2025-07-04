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
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/delegatedstaking"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/predicate"
)

// DelegatedStakingQuery is the builder for querying DelegatedStaking entities.
type DelegatedStakingQuery struct {
	config
	ctx        *QueryContext
	order      []delegatedstaking.OrderOption
	inters     []Interceptor
	predicates []predicate.DelegatedStaking
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DelegatedStakingQuery builder.
func (dsq *DelegatedStakingQuery) Where(ps ...predicate.DelegatedStaking) *DelegatedStakingQuery {
	dsq.predicates = append(dsq.predicates, ps...)
	return dsq
}

// Limit the number of records to be returned by this query.
func (dsq *DelegatedStakingQuery) Limit(limit int) *DelegatedStakingQuery {
	dsq.ctx.Limit = &limit
	return dsq
}

// Offset to start from.
func (dsq *DelegatedStakingQuery) Offset(offset int) *DelegatedStakingQuery {
	dsq.ctx.Offset = &offset
	return dsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dsq *DelegatedStakingQuery) Unique(unique bool) *DelegatedStakingQuery {
	dsq.ctx.Unique = &unique
	return dsq
}

// Order specifies how the records should be ordered.
func (dsq *DelegatedStakingQuery) Order(o ...delegatedstaking.OrderOption) *DelegatedStakingQuery {
	dsq.order = append(dsq.order, o...)
	return dsq
}

// First returns the first DelegatedStaking entity from the query.
// Returns a *NotFoundError when no DelegatedStaking was found.
func (dsq *DelegatedStakingQuery) First(ctx context.Context) (*DelegatedStaking, error) {
	nodes, err := dsq.Limit(1).All(setContextOp(ctx, dsq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{delegatedstaking.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dsq *DelegatedStakingQuery) FirstX(ctx context.Context) *DelegatedStaking {
	node, err := dsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DelegatedStaking ID from the query.
// Returns a *NotFoundError when no DelegatedStaking ID was found.
func (dsq *DelegatedStakingQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = dsq.Limit(1).IDs(setContextOp(ctx, dsq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{delegatedstaking.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dsq *DelegatedStakingQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := dsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DelegatedStaking entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DelegatedStaking entity is found.
// Returns a *NotFoundError when no DelegatedStaking entities are found.
func (dsq *DelegatedStakingQuery) Only(ctx context.Context) (*DelegatedStaking, error) {
	nodes, err := dsq.Limit(2).All(setContextOp(ctx, dsq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{delegatedstaking.Label}
	default:
		return nil, &NotSingularError{delegatedstaking.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dsq *DelegatedStakingQuery) OnlyX(ctx context.Context) *DelegatedStaking {
	node, err := dsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DelegatedStaking ID in the query.
// Returns a *NotSingularError when more than one DelegatedStaking ID is found.
// Returns a *NotFoundError when no entities are found.
func (dsq *DelegatedStakingQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = dsq.Limit(2).IDs(setContextOp(ctx, dsq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{delegatedstaking.Label}
	default:
		err = &NotSingularError{delegatedstaking.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dsq *DelegatedStakingQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := dsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DelegatedStakings.
func (dsq *DelegatedStakingQuery) All(ctx context.Context) ([]*DelegatedStaking, error) {
	ctx = setContextOp(ctx, dsq.ctx, ent.OpQueryAll)
	if err := dsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*DelegatedStaking, *DelegatedStakingQuery]()
	return withInterceptors[[]*DelegatedStaking](ctx, dsq, qr, dsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dsq *DelegatedStakingQuery) AllX(ctx context.Context) []*DelegatedStaking {
	nodes, err := dsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DelegatedStaking IDs.
func (dsq *DelegatedStakingQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if dsq.ctx.Unique == nil && dsq.path != nil {
		dsq.Unique(true)
	}
	ctx = setContextOp(ctx, dsq.ctx, ent.OpQueryIDs)
	if err = dsq.Select(delegatedstaking.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dsq *DelegatedStakingQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := dsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dsq *DelegatedStakingQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dsq.ctx, ent.OpQueryCount)
	if err := dsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dsq, querierCount[*DelegatedStakingQuery](), dsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dsq *DelegatedStakingQuery) CountX(ctx context.Context) int {
	count, err := dsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dsq *DelegatedStakingQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dsq.ctx, ent.OpQueryExist)
	switch _, err := dsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dsq *DelegatedStakingQuery) ExistX(ctx context.Context) bool {
	exist, err := dsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DelegatedStakingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dsq *DelegatedStakingQuery) Clone() *DelegatedStakingQuery {
	if dsq == nil {
		return nil
	}
	return &DelegatedStakingQuery{
		config:     dsq.config,
		ctx:        dsq.ctx.Clone(),
		order:      append([]delegatedstaking.OrderOption{}, dsq.order...),
		inters:     append([]Interceptor{}, dsq.inters...),
		predicates: append([]predicate.DelegatedStaking{}, dsq.predicates...),
		// clone intermediate query.
		sql:       dsq.sql.Clone(),
		path:      dsq.path,
		modifiers: append([]func(*sql.Selector){}, dsq.modifiers...),
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
//	client.DelegatedStaking.Query().
//		GroupBy(delegatedstaking.FieldEntID).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (dsq *DelegatedStakingQuery) GroupBy(field string, fields ...string) *DelegatedStakingGroupBy {
	dsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DelegatedStakingGroupBy{build: dsq}
	grbuild.flds = &dsq.ctx.Fields
	grbuild.label = delegatedstaking.Label
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
//	client.DelegatedStaking.Query().
//		Select(delegatedstaking.FieldEntID).
//		Scan(ctx, &v)
func (dsq *DelegatedStakingQuery) Select(fields ...string) *DelegatedStakingSelect {
	dsq.ctx.Fields = append(dsq.ctx.Fields, fields...)
	sbuild := &DelegatedStakingSelect{DelegatedStakingQuery: dsq}
	sbuild.label = delegatedstaking.Label
	sbuild.flds, sbuild.scan = &dsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DelegatedStakingSelect configured with the given aggregations.
func (dsq *DelegatedStakingQuery) Aggregate(fns ...AggregateFunc) *DelegatedStakingSelect {
	return dsq.Select().Aggregate(fns...)
}

func (dsq *DelegatedStakingQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dsq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dsq); err != nil {
				return err
			}
		}
	}
	for _, f := range dsq.ctx.Fields {
		if !delegatedstaking.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if dsq.path != nil {
		prev, err := dsq.path(ctx)
		if err != nil {
			return err
		}
		dsq.sql = prev
	}
	return nil
}

func (dsq *DelegatedStakingQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DelegatedStaking, error) {
	var (
		nodes = []*DelegatedStaking{}
		_spec = dsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DelegatedStaking).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DelegatedStaking{config: dsq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(dsq.modifiers) > 0 {
		_spec.Modifiers = dsq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (dsq *DelegatedStakingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dsq.querySpec()
	if len(dsq.modifiers) > 0 {
		_spec.Modifiers = dsq.modifiers
	}
	_spec.Node.Columns = dsq.ctx.Fields
	if len(dsq.ctx.Fields) > 0 {
		_spec.Unique = dsq.ctx.Unique != nil && *dsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dsq.driver, _spec)
}

func (dsq *DelegatedStakingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(delegatedstaking.Table, delegatedstaking.Columns, sqlgraph.NewFieldSpec(delegatedstaking.FieldID, field.TypeUint32))
	_spec.From = dsq.sql
	if unique := dsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dsq.path != nil {
		_spec.Unique = true
	}
	if fields := dsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, delegatedstaking.FieldID)
		for i := range fields {
			if fields[i] != delegatedstaking.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dsq *DelegatedStakingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dsq.driver.Dialect())
	t1 := builder.Table(delegatedstaking.Table)
	columns := dsq.ctx.Fields
	if len(columns) == 0 {
		columns = delegatedstaking.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dsq.sql != nil {
		selector = dsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dsq.ctx.Unique != nil && *dsq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range dsq.modifiers {
		m(selector)
	}
	for _, p := range dsq.predicates {
		p(selector)
	}
	for _, p := range dsq.order {
		p(selector)
	}
	if offset := dsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (dsq *DelegatedStakingQuery) ForUpdate(opts ...sql.LockOption) *DelegatedStakingQuery {
	if dsq.driver.Dialect() == dialect.Postgres {
		dsq.Unique(false)
	}
	dsq.modifiers = append(dsq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return dsq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (dsq *DelegatedStakingQuery) ForShare(opts ...sql.LockOption) *DelegatedStakingQuery {
	if dsq.driver.Dialect() == dialect.Postgres {
		dsq.Unique(false)
	}
	dsq.modifiers = append(dsq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return dsq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (dsq *DelegatedStakingQuery) Modify(modifiers ...func(s *sql.Selector)) *DelegatedStakingSelect {
	dsq.modifiers = append(dsq.modifiers, modifiers...)
	return dsq.Select()
}

// DelegatedStakingGroupBy is the group-by builder for DelegatedStaking entities.
type DelegatedStakingGroupBy struct {
	selector
	build *DelegatedStakingQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dsgb *DelegatedStakingGroupBy) Aggregate(fns ...AggregateFunc) *DelegatedStakingGroupBy {
	dsgb.fns = append(dsgb.fns, fns...)
	return dsgb
}

// Scan applies the selector query and scans the result into the given value.
func (dsgb *DelegatedStakingGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dsgb.build.ctx, ent.OpQueryGroupBy)
	if err := dsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DelegatedStakingQuery, *DelegatedStakingGroupBy](ctx, dsgb.build, dsgb, dsgb.build.inters, v)
}

func (dsgb *DelegatedStakingGroupBy) sqlScan(ctx context.Context, root *DelegatedStakingQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dsgb.fns))
	for _, fn := range dsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dsgb.flds)+len(dsgb.fns))
		for _, f := range *dsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DelegatedStakingSelect is the builder for selecting fields of DelegatedStaking entities.
type DelegatedStakingSelect struct {
	*DelegatedStakingQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (dss *DelegatedStakingSelect) Aggregate(fns ...AggregateFunc) *DelegatedStakingSelect {
	dss.fns = append(dss.fns, fns...)
	return dss
}

// Scan applies the selector query and scans the result into the given value.
func (dss *DelegatedStakingSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dss.ctx, ent.OpQuerySelect)
	if err := dss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DelegatedStakingQuery, *DelegatedStakingSelect](ctx, dss.DelegatedStakingQuery, dss, dss.inters, v)
}

func (dss *DelegatedStakingSelect) sqlScan(ctx context.Context, root *DelegatedStakingQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(dss.fns))
	for _, fn := range dss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*dss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (dss *DelegatedStakingSelect) Modify(modifiers ...func(s *sql.Selector)) *DelegatedStakingSelect {
	dss.modifiers = append(dss.modifiers, modifiers...)
	return dss
}
