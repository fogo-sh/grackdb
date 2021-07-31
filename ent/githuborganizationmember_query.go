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
	"github.com/fogo-sh/grackdb/ent/githubaccount"
	"github.com/fogo-sh/grackdb/ent/githuborganization"
	"github.com/fogo-sh/grackdb/ent/githuborganizationmember"
	"github.com/fogo-sh/grackdb/ent/predicate"
)

// GithubOrganizationMemberQuery is the builder for querying GithubOrganizationMember entities.
type GithubOrganizationMemberQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GithubOrganizationMember
	// eager-loading edges.
	withOrganization *GithubOrganizationQuery
	withAccount      *GithubAccountQuery
	withFKs          bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GithubOrganizationMemberQuery builder.
func (gomq *GithubOrganizationMemberQuery) Where(ps ...predicate.GithubOrganizationMember) *GithubOrganizationMemberQuery {
	gomq.predicates = append(gomq.predicates, ps...)
	return gomq
}

// Limit adds a limit step to the query.
func (gomq *GithubOrganizationMemberQuery) Limit(limit int) *GithubOrganizationMemberQuery {
	gomq.limit = &limit
	return gomq
}

// Offset adds an offset step to the query.
func (gomq *GithubOrganizationMemberQuery) Offset(offset int) *GithubOrganizationMemberQuery {
	gomq.offset = &offset
	return gomq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gomq *GithubOrganizationMemberQuery) Unique(unique bool) *GithubOrganizationMemberQuery {
	gomq.unique = &unique
	return gomq
}

// Order adds an order step to the query.
func (gomq *GithubOrganizationMemberQuery) Order(o ...OrderFunc) *GithubOrganizationMemberQuery {
	gomq.order = append(gomq.order, o...)
	return gomq
}

// QueryOrganization chains the current query on the "organization" edge.
func (gomq *GithubOrganizationMemberQuery) QueryOrganization() *GithubOrganizationQuery {
	query := &GithubOrganizationQuery{config: gomq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gomq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gomq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(githuborganizationmember.Table, githuborganizationmember.FieldID, selector),
			sqlgraph.To(githuborganization.Table, githuborganization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, githuborganizationmember.OrganizationTable, githuborganizationmember.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(gomq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAccount chains the current query on the "account" edge.
func (gomq *GithubOrganizationMemberQuery) QueryAccount() *GithubAccountQuery {
	query := &GithubAccountQuery{config: gomq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gomq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gomq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(githuborganizationmember.Table, githuborganizationmember.FieldID, selector),
			sqlgraph.To(githubaccount.Table, githubaccount.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, githuborganizationmember.AccountTable, githuborganizationmember.AccountColumn),
		)
		fromU = sqlgraph.SetNeighbors(gomq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GithubOrganizationMember entity from the query.
// Returns a *NotFoundError when no GithubOrganizationMember was found.
func (gomq *GithubOrganizationMemberQuery) First(ctx context.Context) (*GithubOrganizationMember, error) {
	nodes, err := gomq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{githuborganizationmember.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gomq *GithubOrganizationMemberQuery) FirstX(ctx context.Context) *GithubOrganizationMember {
	node, err := gomq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GithubOrganizationMember ID from the query.
// Returns a *NotFoundError when no GithubOrganizationMember ID was found.
func (gomq *GithubOrganizationMemberQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gomq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{githuborganizationmember.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gomq *GithubOrganizationMemberQuery) FirstIDX(ctx context.Context) int {
	id, err := gomq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GithubOrganizationMember entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one GithubOrganizationMember entity is not found.
// Returns a *NotFoundError when no GithubOrganizationMember entities are found.
func (gomq *GithubOrganizationMemberQuery) Only(ctx context.Context) (*GithubOrganizationMember, error) {
	nodes, err := gomq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{githuborganizationmember.Label}
	default:
		return nil, &NotSingularError{githuborganizationmember.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gomq *GithubOrganizationMemberQuery) OnlyX(ctx context.Context) *GithubOrganizationMember {
	node, err := gomq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GithubOrganizationMember ID in the query.
// Returns a *NotSingularError when exactly one GithubOrganizationMember ID is not found.
// Returns a *NotFoundError when no entities are found.
func (gomq *GithubOrganizationMemberQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gomq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{githuborganizationmember.Label}
	default:
		err = &NotSingularError{githuborganizationmember.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gomq *GithubOrganizationMemberQuery) OnlyIDX(ctx context.Context) int {
	id, err := gomq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GithubOrganizationMembers.
func (gomq *GithubOrganizationMemberQuery) All(ctx context.Context) ([]*GithubOrganizationMember, error) {
	if err := gomq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gomq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gomq *GithubOrganizationMemberQuery) AllX(ctx context.Context) []*GithubOrganizationMember {
	nodes, err := gomq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GithubOrganizationMember IDs.
func (gomq *GithubOrganizationMemberQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := gomq.Select(githuborganizationmember.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gomq *GithubOrganizationMemberQuery) IDsX(ctx context.Context) []int {
	ids, err := gomq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gomq *GithubOrganizationMemberQuery) Count(ctx context.Context) (int, error) {
	if err := gomq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gomq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gomq *GithubOrganizationMemberQuery) CountX(ctx context.Context) int {
	count, err := gomq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gomq *GithubOrganizationMemberQuery) Exist(ctx context.Context) (bool, error) {
	if err := gomq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gomq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gomq *GithubOrganizationMemberQuery) ExistX(ctx context.Context) bool {
	exist, err := gomq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GithubOrganizationMemberQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gomq *GithubOrganizationMemberQuery) Clone() *GithubOrganizationMemberQuery {
	if gomq == nil {
		return nil
	}
	return &GithubOrganizationMemberQuery{
		config:           gomq.config,
		limit:            gomq.limit,
		offset:           gomq.offset,
		order:            append([]OrderFunc{}, gomq.order...),
		predicates:       append([]predicate.GithubOrganizationMember{}, gomq.predicates...),
		withOrganization: gomq.withOrganization.Clone(),
		withAccount:      gomq.withAccount.Clone(),
		// clone intermediate query.
		sql:  gomq.sql.Clone(),
		path: gomq.path,
	}
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (gomq *GithubOrganizationMemberQuery) WithOrganization(opts ...func(*GithubOrganizationQuery)) *GithubOrganizationMemberQuery {
	query := &GithubOrganizationQuery{config: gomq.config}
	for _, opt := range opts {
		opt(query)
	}
	gomq.withOrganization = query
	return gomq
}

// WithAccount tells the query-builder to eager-load the nodes that are connected to
// the "account" edge. The optional arguments are used to configure the query builder of the edge.
func (gomq *GithubOrganizationMemberQuery) WithAccount(opts ...func(*GithubAccountQuery)) *GithubOrganizationMemberQuery {
	query := &GithubAccountQuery{config: gomq.config}
	for _, opt := range opts {
		opt(query)
	}
	gomq.withAccount = query
	return gomq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Role githuborganizationmember.Role `json:"role,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GithubOrganizationMember.Query().
//		GroupBy(githuborganizationmember.FieldRole).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (gomq *GithubOrganizationMemberQuery) GroupBy(field string, fields ...string) *GithubOrganizationMemberGroupBy {
	group := &GithubOrganizationMemberGroupBy{config: gomq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gomq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gomq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Role githuborganizationmember.Role `json:"role,omitempty"`
//	}
//
//	client.GithubOrganizationMember.Query().
//		Select(githuborganizationmember.FieldRole).
//		Scan(ctx, &v)
//
func (gomq *GithubOrganizationMemberQuery) Select(field string, fields ...string) *GithubOrganizationMemberSelect {
	gomq.fields = append([]string{field}, fields...)
	return &GithubOrganizationMemberSelect{GithubOrganizationMemberQuery: gomq}
}

func (gomq *GithubOrganizationMemberQuery) prepareQuery(ctx context.Context) error {
	for _, f := range gomq.fields {
		if !githuborganizationmember.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gomq.path != nil {
		prev, err := gomq.path(ctx)
		if err != nil {
			return err
		}
		gomq.sql = prev
	}
	return nil
}

func (gomq *GithubOrganizationMemberQuery) sqlAll(ctx context.Context) ([]*GithubOrganizationMember, error) {
	var (
		nodes       = []*GithubOrganizationMember{}
		withFKs     = gomq.withFKs
		_spec       = gomq.querySpec()
		loadedTypes = [2]bool{
			gomq.withOrganization != nil,
			gomq.withAccount != nil,
		}
	)
	if gomq.withOrganization != nil || gomq.withAccount != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, githuborganizationmember.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &GithubOrganizationMember{config: gomq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, gomq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := gomq.withOrganization; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*GithubOrganizationMember)
		for i := range nodes {
			if nodes[i].github_organization_members == nil {
				continue
			}
			fk := *nodes[i].github_organization_members
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(githuborganization.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "github_organization_members" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Organization = n
			}
		}
	}

	if query := gomq.withAccount; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*GithubOrganizationMember)
		for i := range nodes {
			if nodes[i].github_account_organization_memberships == nil {
				continue
			}
			fk := *nodes[i].github_account_organization_memberships
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(githubaccount.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "github_account_organization_memberships" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Account = n
			}
		}
	}

	return nodes, nil
}

func (gomq *GithubOrganizationMemberQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gomq.querySpec()
	return sqlgraph.CountNodes(ctx, gomq.driver, _spec)
}

func (gomq *GithubOrganizationMemberQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gomq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (gomq *GithubOrganizationMemberQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   githuborganizationmember.Table,
			Columns: githuborganizationmember.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: githuborganizationmember.FieldID,
			},
		},
		From:   gomq.sql,
		Unique: true,
	}
	if unique := gomq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gomq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, githuborganizationmember.FieldID)
		for i := range fields {
			if fields[i] != githuborganizationmember.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gomq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gomq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gomq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gomq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gomq *GithubOrganizationMemberQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gomq.driver.Dialect())
	t1 := builder.Table(githuborganizationmember.Table)
	columns := gomq.fields
	if len(columns) == 0 {
		columns = githuborganizationmember.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gomq.sql != nil {
		selector = gomq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range gomq.predicates {
		p(selector)
	}
	for _, p := range gomq.order {
		p(selector)
	}
	if offset := gomq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gomq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GithubOrganizationMemberGroupBy is the group-by builder for GithubOrganizationMember entities.
type GithubOrganizationMemberGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gomgb *GithubOrganizationMemberGroupBy) Aggregate(fns ...AggregateFunc) *GithubOrganizationMemberGroupBy {
	gomgb.fns = append(gomgb.fns, fns...)
	return gomgb
}

// Scan applies the group-by query and scans the result into the given value.
func (gomgb *GithubOrganizationMemberGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := gomgb.path(ctx)
	if err != nil {
		return err
	}
	gomgb.sql = query
	return gomgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gomgb *GithubOrganizationMemberGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := gomgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (gomgb *GithubOrganizationMemberGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(gomgb.fields) > 1 {
		return nil, errors.New("ent: GithubOrganizationMemberGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := gomgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gomgb *GithubOrganizationMemberGroupBy) StringsX(ctx context.Context) []string {
	v, err := gomgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gomgb *GithubOrganizationMemberGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gomgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{githuborganizationmember.Label}
	default:
		err = fmt.Errorf("ent: GithubOrganizationMemberGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gomgb *GithubOrganizationMemberGroupBy) StringX(ctx context.Context) string {
	v, err := gomgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (gomgb *GithubOrganizationMemberGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(gomgb.fields) > 1 {
		return nil, errors.New("ent: GithubOrganizationMemberGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := gomgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gomgb *GithubOrganizationMemberGroupBy) IntsX(ctx context.Context) []int {
	v, err := gomgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gomgb *GithubOrganizationMemberGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gomgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{githuborganizationmember.Label}
	default:
		err = fmt.Errorf("ent: GithubOrganizationMemberGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gomgb *GithubOrganizationMemberGroupBy) IntX(ctx context.Context) int {
	v, err := gomgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (gomgb *GithubOrganizationMemberGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(gomgb.fields) > 1 {
		return nil, errors.New("ent: GithubOrganizationMemberGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := gomgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gomgb *GithubOrganizationMemberGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := gomgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gomgb *GithubOrganizationMemberGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gomgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{githuborganizationmember.Label}
	default:
		err = fmt.Errorf("ent: GithubOrganizationMemberGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gomgb *GithubOrganizationMemberGroupBy) Float64X(ctx context.Context) float64 {
	v, err := gomgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (gomgb *GithubOrganizationMemberGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(gomgb.fields) > 1 {
		return nil, errors.New("ent: GithubOrganizationMemberGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := gomgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gomgb *GithubOrganizationMemberGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := gomgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gomgb *GithubOrganizationMemberGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gomgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{githuborganizationmember.Label}
	default:
		err = fmt.Errorf("ent: GithubOrganizationMemberGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gomgb *GithubOrganizationMemberGroupBy) BoolX(ctx context.Context) bool {
	v, err := gomgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gomgb *GithubOrganizationMemberGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range gomgb.fields {
		if !githuborganizationmember.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := gomgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gomgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gomgb *GithubOrganizationMemberGroupBy) sqlQuery() *sql.Selector {
	selector := gomgb.sql.Select()
	aggregation := make([]string, 0, len(gomgb.fns))
	for _, fn := range gomgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(gomgb.fields)+len(gomgb.fns))
		for _, f := range gomgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(gomgb.fields...)...)
}

// GithubOrganizationMemberSelect is the builder for selecting fields of GithubOrganizationMember entities.
type GithubOrganizationMemberSelect struct {
	*GithubOrganizationMemberQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (goms *GithubOrganizationMemberSelect) Scan(ctx context.Context, v interface{}) error {
	if err := goms.prepareQuery(ctx); err != nil {
		return err
	}
	goms.sql = goms.GithubOrganizationMemberQuery.sqlQuery(ctx)
	return goms.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (goms *GithubOrganizationMemberSelect) ScanX(ctx context.Context, v interface{}) {
	if err := goms.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (goms *GithubOrganizationMemberSelect) Strings(ctx context.Context) ([]string, error) {
	if len(goms.fields) > 1 {
		return nil, errors.New("ent: GithubOrganizationMemberSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := goms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (goms *GithubOrganizationMemberSelect) StringsX(ctx context.Context) []string {
	v, err := goms.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (goms *GithubOrganizationMemberSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = goms.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{githuborganizationmember.Label}
	default:
		err = fmt.Errorf("ent: GithubOrganizationMemberSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (goms *GithubOrganizationMemberSelect) StringX(ctx context.Context) string {
	v, err := goms.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (goms *GithubOrganizationMemberSelect) Ints(ctx context.Context) ([]int, error) {
	if len(goms.fields) > 1 {
		return nil, errors.New("ent: GithubOrganizationMemberSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := goms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (goms *GithubOrganizationMemberSelect) IntsX(ctx context.Context) []int {
	v, err := goms.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (goms *GithubOrganizationMemberSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = goms.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{githuborganizationmember.Label}
	default:
		err = fmt.Errorf("ent: GithubOrganizationMemberSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (goms *GithubOrganizationMemberSelect) IntX(ctx context.Context) int {
	v, err := goms.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (goms *GithubOrganizationMemberSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(goms.fields) > 1 {
		return nil, errors.New("ent: GithubOrganizationMemberSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := goms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (goms *GithubOrganizationMemberSelect) Float64sX(ctx context.Context) []float64 {
	v, err := goms.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (goms *GithubOrganizationMemberSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = goms.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{githuborganizationmember.Label}
	default:
		err = fmt.Errorf("ent: GithubOrganizationMemberSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (goms *GithubOrganizationMemberSelect) Float64X(ctx context.Context) float64 {
	v, err := goms.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (goms *GithubOrganizationMemberSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(goms.fields) > 1 {
		return nil, errors.New("ent: GithubOrganizationMemberSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := goms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (goms *GithubOrganizationMemberSelect) BoolsX(ctx context.Context) []bool {
	v, err := goms.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (goms *GithubOrganizationMemberSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = goms.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{githuborganizationmember.Label}
	default:
		err = fmt.Errorf("ent: GithubOrganizationMemberSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (goms *GithubOrganizationMemberSelect) BoolX(ctx context.Context) bool {
	v, err := goms.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (goms *GithubOrganizationMemberSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := goms.sql.Query()
	if err := goms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
