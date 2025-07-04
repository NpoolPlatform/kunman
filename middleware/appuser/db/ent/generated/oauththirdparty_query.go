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
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/oauththirdparty"
	"github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/predicate"
)

// OAuthThirdPartyQuery is the builder for querying OAuthThirdParty entities.
type OAuthThirdPartyQuery struct {
	config
	ctx        *QueryContext
	order      []oauththirdparty.OrderOption
	inters     []Interceptor
	predicates []predicate.OAuthThirdParty
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OAuthThirdPartyQuery builder.
func (otpq *OAuthThirdPartyQuery) Where(ps ...predicate.OAuthThirdParty) *OAuthThirdPartyQuery {
	otpq.predicates = append(otpq.predicates, ps...)
	return otpq
}

// Limit the number of records to be returned by this query.
func (otpq *OAuthThirdPartyQuery) Limit(limit int) *OAuthThirdPartyQuery {
	otpq.ctx.Limit = &limit
	return otpq
}

// Offset to start from.
func (otpq *OAuthThirdPartyQuery) Offset(offset int) *OAuthThirdPartyQuery {
	otpq.ctx.Offset = &offset
	return otpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (otpq *OAuthThirdPartyQuery) Unique(unique bool) *OAuthThirdPartyQuery {
	otpq.ctx.Unique = &unique
	return otpq
}

// Order specifies how the records should be ordered.
func (otpq *OAuthThirdPartyQuery) Order(o ...oauththirdparty.OrderOption) *OAuthThirdPartyQuery {
	otpq.order = append(otpq.order, o...)
	return otpq
}

// First returns the first OAuthThirdParty entity from the query.
// Returns a *NotFoundError when no OAuthThirdParty was found.
func (otpq *OAuthThirdPartyQuery) First(ctx context.Context) (*OAuthThirdParty, error) {
	nodes, err := otpq.Limit(1).All(setContextOp(ctx, otpq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{oauththirdparty.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (otpq *OAuthThirdPartyQuery) FirstX(ctx context.Context) *OAuthThirdParty {
	node, err := otpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OAuthThirdParty ID from the query.
// Returns a *NotFoundError when no OAuthThirdParty ID was found.
func (otpq *OAuthThirdPartyQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = otpq.Limit(1).IDs(setContextOp(ctx, otpq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{oauththirdparty.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (otpq *OAuthThirdPartyQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := otpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OAuthThirdParty entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OAuthThirdParty entity is found.
// Returns a *NotFoundError when no OAuthThirdParty entities are found.
func (otpq *OAuthThirdPartyQuery) Only(ctx context.Context) (*OAuthThirdParty, error) {
	nodes, err := otpq.Limit(2).All(setContextOp(ctx, otpq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{oauththirdparty.Label}
	default:
		return nil, &NotSingularError{oauththirdparty.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (otpq *OAuthThirdPartyQuery) OnlyX(ctx context.Context) *OAuthThirdParty {
	node, err := otpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OAuthThirdParty ID in the query.
// Returns a *NotSingularError when more than one OAuthThirdParty ID is found.
// Returns a *NotFoundError when no entities are found.
func (otpq *OAuthThirdPartyQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = otpq.Limit(2).IDs(setContextOp(ctx, otpq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{oauththirdparty.Label}
	default:
		err = &NotSingularError{oauththirdparty.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (otpq *OAuthThirdPartyQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := otpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OAuthThirdParties.
func (otpq *OAuthThirdPartyQuery) All(ctx context.Context) ([]*OAuthThirdParty, error) {
	ctx = setContextOp(ctx, otpq.ctx, ent.OpQueryAll)
	if err := otpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OAuthThirdParty, *OAuthThirdPartyQuery]()
	return withInterceptors[[]*OAuthThirdParty](ctx, otpq, qr, otpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (otpq *OAuthThirdPartyQuery) AllX(ctx context.Context) []*OAuthThirdParty {
	nodes, err := otpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OAuthThirdParty IDs.
func (otpq *OAuthThirdPartyQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if otpq.ctx.Unique == nil && otpq.path != nil {
		otpq.Unique(true)
	}
	ctx = setContextOp(ctx, otpq.ctx, ent.OpQueryIDs)
	if err = otpq.Select(oauththirdparty.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (otpq *OAuthThirdPartyQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := otpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (otpq *OAuthThirdPartyQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, otpq.ctx, ent.OpQueryCount)
	if err := otpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, otpq, querierCount[*OAuthThirdPartyQuery](), otpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (otpq *OAuthThirdPartyQuery) CountX(ctx context.Context) int {
	count, err := otpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (otpq *OAuthThirdPartyQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, otpq.ctx, ent.OpQueryExist)
	switch _, err := otpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (otpq *OAuthThirdPartyQuery) ExistX(ctx context.Context) bool {
	exist, err := otpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OAuthThirdPartyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (otpq *OAuthThirdPartyQuery) Clone() *OAuthThirdPartyQuery {
	if otpq == nil {
		return nil
	}
	return &OAuthThirdPartyQuery{
		config:     otpq.config,
		ctx:        otpq.ctx.Clone(),
		order:      append([]oauththirdparty.OrderOption{}, otpq.order...),
		inters:     append([]Interceptor{}, otpq.inters...),
		predicates: append([]predicate.OAuthThirdParty{}, otpq.predicates...),
		// clone intermediate query.
		sql:       otpq.sql.Clone(),
		path:      otpq.path,
		modifiers: append([]func(*sql.Selector){}, otpq.modifiers...),
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
//	client.OAuthThirdParty.Query().
//		GroupBy(oauththirdparty.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (otpq *OAuthThirdPartyQuery) GroupBy(field string, fields ...string) *OAuthThirdPartyGroupBy {
	otpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OAuthThirdPartyGroupBy{build: otpq}
	grbuild.flds = &otpq.ctx.Fields
	grbuild.label = oauththirdparty.Label
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
//	client.OAuthThirdParty.Query().
//		Select(oauththirdparty.FieldCreatedAt).
//		Scan(ctx, &v)
func (otpq *OAuthThirdPartyQuery) Select(fields ...string) *OAuthThirdPartySelect {
	otpq.ctx.Fields = append(otpq.ctx.Fields, fields...)
	sbuild := &OAuthThirdPartySelect{OAuthThirdPartyQuery: otpq}
	sbuild.label = oauththirdparty.Label
	sbuild.flds, sbuild.scan = &otpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OAuthThirdPartySelect configured with the given aggregations.
func (otpq *OAuthThirdPartyQuery) Aggregate(fns ...AggregateFunc) *OAuthThirdPartySelect {
	return otpq.Select().Aggregate(fns...)
}

func (otpq *OAuthThirdPartyQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range otpq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, otpq); err != nil {
				return err
			}
		}
	}
	for _, f := range otpq.ctx.Fields {
		if !oauththirdparty.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if otpq.path != nil {
		prev, err := otpq.path(ctx)
		if err != nil {
			return err
		}
		otpq.sql = prev
	}
	return nil
}

func (otpq *OAuthThirdPartyQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OAuthThirdParty, error) {
	var (
		nodes = []*OAuthThirdParty{}
		_spec = otpq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OAuthThirdParty).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OAuthThirdParty{config: otpq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(otpq.modifiers) > 0 {
		_spec.Modifiers = otpq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, otpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (otpq *OAuthThirdPartyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := otpq.querySpec()
	if len(otpq.modifiers) > 0 {
		_spec.Modifiers = otpq.modifiers
	}
	_spec.Node.Columns = otpq.ctx.Fields
	if len(otpq.ctx.Fields) > 0 {
		_spec.Unique = otpq.ctx.Unique != nil && *otpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, otpq.driver, _spec)
}

func (otpq *OAuthThirdPartyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(oauththirdparty.Table, oauththirdparty.Columns, sqlgraph.NewFieldSpec(oauththirdparty.FieldID, field.TypeUint32))
	_spec.From = otpq.sql
	if unique := otpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if otpq.path != nil {
		_spec.Unique = true
	}
	if fields := otpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, oauththirdparty.FieldID)
		for i := range fields {
			if fields[i] != oauththirdparty.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := otpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := otpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := otpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := otpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (otpq *OAuthThirdPartyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(otpq.driver.Dialect())
	t1 := builder.Table(oauththirdparty.Table)
	columns := otpq.ctx.Fields
	if len(columns) == 0 {
		columns = oauththirdparty.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if otpq.sql != nil {
		selector = otpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if otpq.ctx.Unique != nil && *otpq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range otpq.modifiers {
		m(selector)
	}
	for _, p := range otpq.predicates {
		p(selector)
	}
	for _, p := range otpq.order {
		p(selector)
	}
	if offset := otpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := otpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (otpq *OAuthThirdPartyQuery) ForUpdate(opts ...sql.LockOption) *OAuthThirdPartyQuery {
	if otpq.driver.Dialect() == dialect.Postgres {
		otpq.Unique(false)
	}
	otpq.modifiers = append(otpq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return otpq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (otpq *OAuthThirdPartyQuery) ForShare(opts ...sql.LockOption) *OAuthThirdPartyQuery {
	if otpq.driver.Dialect() == dialect.Postgres {
		otpq.Unique(false)
	}
	otpq.modifiers = append(otpq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return otpq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (otpq *OAuthThirdPartyQuery) Modify(modifiers ...func(s *sql.Selector)) *OAuthThirdPartySelect {
	otpq.modifiers = append(otpq.modifiers, modifiers...)
	return otpq.Select()
}

// OAuthThirdPartyGroupBy is the group-by builder for OAuthThirdParty entities.
type OAuthThirdPartyGroupBy struct {
	selector
	build *OAuthThirdPartyQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (otpgb *OAuthThirdPartyGroupBy) Aggregate(fns ...AggregateFunc) *OAuthThirdPartyGroupBy {
	otpgb.fns = append(otpgb.fns, fns...)
	return otpgb
}

// Scan applies the selector query and scans the result into the given value.
func (otpgb *OAuthThirdPartyGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, otpgb.build.ctx, ent.OpQueryGroupBy)
	if err := otpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OAuthThirdPartyQuery, *OAuthThirdPartyGroupBy](ctx, otpgb.build, otpgb, otpgb.build.inters, v)
}

func (otpgb *OAuthThirdPartyGroupBy) sqlScan(ctx context.Context, root *OAuthThirdPartyQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(otpgb.fns))
	for _, fn := range otpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*otpgb.flds)+len(otpgb.fns))
		for _, f := range *otpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*otpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := otpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OAuthThirdPartySelect is the builder for selecting fields of OAuthThirdParty entities.
type OAuthThirdPartySelect struct {
	*OAuthThirdPartyQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (otps *OAuthThirdPartySelect) Aggregate(fns ...AggregateFunc) *OAuthThirdPartySelect {
	otps.fns = append(otps.fns, fns...)
	return otps
}

// Scan applies the selector query and scans the result into the given value.
func (otps *OAuthThirdPartySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, otps.ctx, ent.OpQuerySelect)
	if err := otps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OAuthThirdPartyQuery, *OAuthThirdPartySelect](ctx, otps.OAuthThirdPartyQuery, otps, otps.inters, v)
}

func (otps *OAuthThirdPartySelect) sqlScan(ctx context.Context, root *OAuthThirdPartyQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(otps.fns))
	for _, fn := range otps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*otps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := otps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (otps *OAuthThirdPartySelect) Modify(modifiers ...func(s *sql.Selector)) *OAuthThirdPartySelect {
	otps.modifiers = append(otps.modifiers, modifiers...)
	return otps
}
