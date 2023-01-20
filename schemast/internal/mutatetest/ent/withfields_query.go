// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/contrib/schemast/internal/mutatetest/ent/predicate"
	"entgo.io/contrib/schemast/internal/mutatetest/ent/withfields"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WithFieldsQuery is the builder for querying WithFields entities.
type WithFieldsQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.WithFields
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WithFieldsQuery builder.
func (wfq *WithFieldsQuery) Where(ps ...predicate.WithFields) *WithFieldsQuery {
	wfq.predicates = append(wfq.predicates, ps...)
	return wfq
}

// Limit the number of records to be returned by this query.
func (wfq *WithFieldsQuery) Limit(limit int) *WithFieldsQuery {
	wfq.ctx.Limit = &limit
	return wfq
}

// Offset to start from.
func (wfq *WithFieldsQuery) Offset(offset int) *WithFieldsQuery {
	wfq.ctx.Offset = &offset
	return wfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wfq *WithFieldsQuery) Unique(unique bool) *WithFieldsQuery {
	wfq.ctx.Unique = &unique
	return wfq
}

// Order specifies how the records should be ordered.
func (wfq *WithFieldsQuery) Order(o ...OrderFunc) *WithFieldsQuery {
	wfq.order = append(wfq.order, o...)
	return wfq
}

// First returns the first WithFields entity from the query.
// Returns a *NotFoundError when no WithFields was found.
func (wfq *WithFieldsQuery) First(ctx context.Context) (*WithFields, error) {
	nodes, err := wfq.Limit(1).All(setContextOp(ctx, wfq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{withfields.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wfq *WithFieldsQuery) FirstX(ctx context.Context) *WithFields {
	node, err := wfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WithFields ID from the query.
// Returns a *NotFoundError when no WithFields ID was found.
func (wfq *WithFieldsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wfq.Limit(1).IDs(setContextOp(ctx, wfq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{withfields.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wfq *WithFieldsQuery) FirstIDX(ctx context.Context) int {
	id, err := wfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WithFields entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WithFields entity is found.
// Returns a *NotFoundError when no WithFields entities are found.
func (wfq *WithFieldsQuery) Only(ctx context.Context) (*WithFields, error) {
	nodes, err := wfq.Limit(2).All(setContextOp(ctx, wfq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{withfields.Label}
	default:
		return nil, &NotSingularError{withfields.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wfq *WithFieldsQuery) OnlyX(ctx context.Context) *WithFields {
	node, err := wfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WithFields ID in the query.
// Returns a *NotSingularError when more than one WithFields ID is found.
// Returns a *NotFoundError when no entities are found.
func (wfq *WithFieldsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wfq.Limit(2).IDs(setContextOp(ctx, wfq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{withfields.Label}
	default:
		err = &NotSingularError{withfields.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wfq *WithFieldsQuery) OnlyIDX(ctx context.Context) int {
	id, err := wfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WithFieldsSlice.
func (wfq *WithFieldsQuery) All(ctx context.Context) ([]*WithFields, error) {
	ctx = setContextOp(ctx, wfq.ctx, "All")
	if err := wfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WithFields, *WithFieldsQuery]()
	return withInterceptors[[]*WithFields](ctx, wfq, qr, wfq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wfq *WithFieldsQuery) AllX(ctx context.Context) []*WithFields {
	nodes, err := wfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WithFields IDs.
func (wfq *WithFieldsQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	ctx = setContextOp(ctx, wfq.ctx, "IDs")
	if err := wfq.Select(withfields.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wfq *WithFieldsQuery) IDsX(ctx context.Context) []int {
	ids, err := wfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wfq *WithFieldsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wfq.ctx, "Count")
	if err := wfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wfq, querierCount[*WithFieldsQuery](), wfq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wfq *WithFieldsQuery) CountX(ctx context.Context) int {
	count, err := wfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wfq *WithFieldsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wfq.ctx, "Exist")
	switch _, err := wfq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wfq *WithFieldsQuery) ExistX(ctx context.Context) bool {
	exist, err := wfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WithFieldsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wfq *WithFieldsQuery) Clone() *WithFieldsQuery {
	if wfq == nil {
		return nil
	}
	return &WithFieldsQuery{
		config:     wfq.config,
		ctx:        wfq.ctx.Clone(),
		order:      append([]OrderFunc{}, wfq.order...),
		inters:     append([]Interceptor{}, wfq.inters...),
		predicates: append([]predicate.WithFields{}, wfq.predicates...),
		// clone intermediate query.
		sql:  wfq.sql.Clone(),
		path: wfq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Existing string `json:"existing,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.WithFields.Query().
//		GroupBy(withfields.FieldExisting).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (wfq *WithFieldsQuery) GroupBy(field string, fields ...string) *WithFieldsGroupBy {
	wfq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WithFieldsGroupBy{build: wfq}
	grbuild.flds = &wfq.ctx.Fields
	grbuild.label = withfields.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Existing string `json:"existing,omitempty"`
//	}
//
//	client.WithFields.Query().
//		Select(withfields.FieldExisting).
//		Scan(ctx, &v)
func (wfq *WithFieldsQuery) Select(fields ...string) *WithFieldsSelect {
	wfq.ctx.Fields = append(wfq.ctx.Fields, fields...)
	sbuild := &WithFieldsSelect{WithFieldsQuery: wfq}
	sbuild.label = withfields.Label
	sbuild.flds, sbuild.scan = &wfq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WithFieldsSelect configured with the given aggregations.
func (wfq *WithFieldsQuery) Aggregate(fns ...AggregateFunc) *WithFieldsSelect {
	return wfq.Select().Aggregate(fns...)
}

func (wfq *WithFieldsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wfq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wfq); err != nil {
				return err
			}
		}
	}
	for _, f := range wfq.ctx.Fields {
		if !withfields.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wfq.path != nil {
		prev, err := wfq.path(ctx)
		if err != nil {
			return err
		}
		wfq.sql = prev
	}
	return nil
}

func (wfq *WithFieldsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WithFields, error) {
	var (
		nodes = []*WithFields{}
		_spec = wfq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WithFields).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WithFields{config: wfq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (wfq *WithFieldsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wfq.querySpec()
	_spec.Node.Columns = wfq.ctx.Fields
	if len(wfq.ctx.Fields) > 0 {
		_spec.Unique = wfq.ctx.Unique != nil && *wfq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wfq.driver, _spec)
}

func (wfq *WithFieldsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   withfields.Table,
			Columns: withfields.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: withfields.FieldID,
			},
		},
		From:   wfq.sql,
		Unique: true,
	}
	if unique := wfq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := wfq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, withfields.FieldID)
		for i := range fields {
			if fields[i] != withfields.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wfq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wfq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wfq *WithFieldsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wfq.driver.Dialect())
	t1 := builder.Table(withfields.Table)
	columns := wfq.ctx.Fields
	if len(columns) == 0 {
		columns = withfields.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wfq.sql != nil {
		selector = wfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wfq.ctx.Unique != nil && *wfq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range wfq.predicates {
		p(selector)
	}
	for _, p := range wfq.order {
		p(selector)
	}
	if offset := wfq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wfq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithFieldsGroupBy is the group-by builder for WithFields entities.
type WithFieldsGroupBy struct {
	selector
	build *WithFieldsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wfgb *WithFieldsGroupBy) Aggregate(fns ...AggregateFunc) *WithFieldsGroupBy {
	wfgb.fns = append(wfgb.fns, fns...)
	return wfgb
}

// Scan applies the selector query and scans the result into the given value.
func (wfgb *WithFieldsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wfgb.build.ctx, "GroupBy")
	if err := wfgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WithFieldsQuery, *WithFieldsGroupBy](ctx, wfgb.build, wfgb, wfgb.build.inters, v)
}

func (wfgb *WithFieldsGroupBy) sqlScan(ctx context.Context, root *WithFieldsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wfgb.fns))
	for _, fn := range wfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wfgb.flds)+len(wfgb.fns))
		for _, f := range *wfgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wfgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wfgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WithFieldsSelect is the builder for selecting fields of WithFields entities.
type WithFieldsSelect struct {
	*WithFieldsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (wfs *WithFieldsSelect) Aggregate(fns ...AggregateFunc) *WithFieldsSelect {
	wfs.fns = append(wfs.fns, fns...)
	return wfs
}

// Scan applies the selector query and scans the result into the given value.
func (wfs *WithFieldsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wfs.ctx, "Select")
	if err := wfs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WithFieldsQuery, *WithFieldsSelect](ctx, wfs.WithFieldsQuery, wfs, wfs.inters, v)
}

func (wfs *WithFieldsSelect) sqlScan(ctx context.Context, root *WithFieldsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(wfs.fns))
	for _, fn := range wfs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*wfs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
