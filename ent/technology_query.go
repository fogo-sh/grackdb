// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/predicate"
	"github.com/fogo-sh/grackdb/ent/technology"
)

// TechnologyQuery is the builder for querying Technology entities.
type TechnologyQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Technology
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TechnologyQuery builder.
func (tq *TechnologyQuery) Where(ps ...predicate.Technology) *TechnologyQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit adds a limit step to the query.
func (tq *TechnologyQuery) Limit(limit int) *TechnologyQuery {
	tq.limit = &limit
	return tq
}

// Offset adds an offset step to the query.
func (tq *TechnologyQuery) Offset(offset int) *TechnologyQuery {
	tq.offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TechnologyQuery) Unique(unique bool) *TechnologyQuery {
	tq.unique = &unique
	return tq
}

// Order adds an order step to the query.
func (tq *TechnologyQuery) Order(o ...OrderFunc) *TechnologyQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// First returns the first Technology entity from the query.
// Returns a *NotFoundError when no Technology was found.
func (tq *TechnologyQuery) First(ctx context.Context) (*Technology, error) {
	nodes, err := tq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{technology.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TechnologyQuery) FirstX(ctx context.Context) *Technology {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Technology ID from the query.
// Returns a *NotFoundError when no Technology ID was found.
func (tq *TechnologyQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{technology.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TechnologyQuery) FirstIDX(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Technology entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Technology entity is not found.
// Returns a *NotFoundError when no Technology entities are found.
func (tq *TechnologyQuery) Only(ctx context.Context) (*Technology, error) {
	nodes, err := tq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{technology.Label}
	default:
		return nil, &NotSingularError{technology.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TechnologyQuery) OnlyX(ctx context.Context) *Technology {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Technology ID in the query.
// Returns a *NotSingularError when exactly one Technology ID is not found.
// Returns a *NotFoundError when no entities are found.
func (tq *TechnologyQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{technology.Label}
	default:
		err = &NotSingularError{technology.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TechnologyQuery) OnlyIDX(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Technologies.
func (tq *TechnologyQuery) All(ctx context.Context) ([]*Technology, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tq *TechnologyQuery) AllX(ctx context.Context) []*Technology {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Technology IDs.
func (tq *TechnologyQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := tq.Select(technology.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TechnologyQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TechnologyQuery) Count(ctx context.Context) (int, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TechnologyQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TechnologyQuery) Exist(ctx context.Context) (bool, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TechnologyQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TechnologyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TechnologyQuery) Clone() *TechnologyQuery {
	if tq == nil {
		return nil
	}
	return &TechnologyQuery{
		config:     tq.config,
		limit:      tq.limit,
		offset:     tq.offset,
		order:      append([]OrderFunc{}, tq.order...),
		predicates: append([]predicate.Technology{}, tq.predicates...),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Technology.Query().
//		GroupBy(technology.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tq *TechnologyQuery) GroupBy(field string, fields ...string) *TechnologyGroupBy {
	group := &TechnologyGroupBy{config: tq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Technology.Query().
//		Select(technology.FieldName).
//		Scan(ctx, &v)
//
func (tq *TechnologyQuery) Select(field string, fields ...string) *TechnologySelect {
	tq.fields = append([]string{field}, fields...)
	return &TechnologySelect{TechnologyQuery: tq}
}

func (tq *TechnologyQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tq.fields {
		if !technology.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	if technology.Policy == nil {
		return errors.New("ent: uninitialized technology.Policy (forgotten import ent/runtime?)")
	}
	if err := technology.Policy.EvalQuery(ctx, tq); err != nil {
		return err
	}
	return nil
}

func (tq *TechnologyQuery) sqlAll(ctx context.Context) ([]*Technology, error) {
	var (
		nodes = []*Technology{}
		_spec = tq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Technology{config: tq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tq *TechnologyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TechnologyQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tq *TechnologyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   technology.Table,
			Columns: technology.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: technology.FieldID,
			},
		},
		From:   tq.sql,
		Unique: true,
	}
	if unique := tq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, technology.FieldID)
		for i := range fields {
			if fields[i] != technology.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TechnologyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(technology.Table)
	columns := tq.fields
	if len(columns) == 0 {
		columns = technology.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TechnologyGroupBy is the group-by builder for Technology entities.
type TechnologyGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TechnologyGroupBy) Aggregate(fns ...AggregateFunc) *TechnologyGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tgb *TechnologyGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tgb.path(ctx)
	if err != nil {
		return err
	}
	tgb.sql = query
	return tgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tgb *TechnologyGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := tgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (tgb *TechnologyGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TechnologyGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tgb *TechnologyGroupBy) StringsX(ctx context.Context) []string {
	v, err := tgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tgb *TechnologyGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{technology.Label}
	default:
		err = fmt.Errorf("ent: TechnologyGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tgb *TechnologyGroupBy) StringX(ctx context.Context) string {
	v, err := tgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (tgb *TechnologyGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TechnologyGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tgb *TechnologyGroupBy) IntsX(ctx context.Context) []int {
	v, err := tgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tgb *TechnologyGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{technology.Label}
	default:
		err = fmt.Errorf("ent: TechnologyGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tgb *TechnologyGroupBy) IntX(ctx context.Context) int {
	v, err := tgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (tgb *TechnologyGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TechnologyGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tgb *TechnologyGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := tgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tgb *TechnologyGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{technology.Label}
	default:
		err = fmt.Errorf("ent: TechnologyGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tgb *TechnologyGroupBy) Float64X(ctx context.Context) float64 {
	v, err := tgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (tgb *TechnologyGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TechnologyGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tgb *TechnologyGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := tgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tgb *TechnologyGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{technology.Label}
	default:
		err = fmt.Errorf("ent: TechnologyGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tgb *TechnologyGroupBy) BoolX(ctx context.Context) bool {
	v, err := tgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tgb *TechnologyGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tgb.fields {
		if !technology.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tgb *TechnologyGroupBy) sqlQuery() *sql.Selector {
	selector := tgb.sql.Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tgb.fields)+len(tgb.fns))
		for _, f := range tgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tgb.fields...)...)
}

// TechnologySelect is the builder for selecting fields of Technology entities.
type TechnologySelect struct {
	*TechnologyQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TechnologySelect) Scan(ctx context.Context, v interface{}) error {
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	ts.sql = ts.TechnologyQuery.sqlQuery(ctx)
	return ts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ts *TechnologySelect) ScanX(ctx context.Context, v interface{}) {
	if err := ts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ts *TechnologySelect) Strings(ctx context.Context) ([]string, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TechnologySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ts *TechnologySelect) StringsX(ctx context.Context) []string {
	v, err := ts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ts *TechnologySelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{technology.Label}
	default:
		err = fmt.Errorf("ent: TechnologySelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ts *TechnologySelect) StringX(ctx context.Context) string {
	v, err := ts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ts *TechnologySelect) Ints(ctx context.Context) ([]int, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TechnologySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ts *TechnologySelect) IntsX(ctx context.Context) []int {
	v, err := ts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ts *TechnologySelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{technology.Label}
	default:
		err = fmt.Errorf("ent: TechnologySelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ts *TechnologySelect) IntX(ctx context.Context) int {
	v, err := ts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ts *TechnologySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TechnologySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ts *TechnologySelect) Float64sX(ctx context.Context) []float64 {
	v, err := ts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ts *TechnologySelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{technology.Label}
	default:
		err = fmt.Errorf("ent: TechnologySelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ts *TechnologySelect) Float64X(ctx context.Context) float64 {
	v, err := ts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ts *TechnologySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TechnologySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ts *TechnologySelect) BoolsX(ctx context.Context) []bool {
	v, err := ts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ts *TechnologySelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{technology.Label}
	default:
		err = fmt.Errorf("ent: TechnologySelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ts *TechnologySelect) BoolX(ctx context.Context) bool {
	v, err := ts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ts *TechnologySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ts.sql.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
