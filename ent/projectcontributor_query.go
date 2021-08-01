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
	"github.com/fogo-sh/grackdb/ent/project"
	"github.com/fogo-sh/grackdb/ent/projectcontributor"
	"github.com/fogo-sh/grackdb/ent/user"
)

// ProjectContributorQuery is the builder for querying ProjectContributor entities.
type ProjectContributorQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ProjectContributor
	// eager-loading edges.
	withProject *ProjectQuery
	withUser    *UserQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProjectContributorQuery builder.
func (pcq *ProjectContributorQuery) Where(ps ...predicate.ProjectContributor) *ProjectContributorQuery {
	pcq.predicates = append(pcq.predicates, ps...)
	return pcq
}

// Limit adds a limit step to the query.
func (pcq *ProjectContributorQuery) Limit(limit int) *ProjectContributorQuery {
	pcq.limit = &limit
	return pcq
}

// Offset adds an offset step to the query.
func (pcq *ProjectContributorQuery) Offset(offset int) *ProjectContributorQuery {
	pcq.offset = &offset
	return pcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pcq *ProjectContributorQuery) Unique(unique bool) *ProjectContributorQuery {
	pcq.unique = &unique
	return pcq
}

// Order adds an order step to the query.
func (pcq *ProjectContributorQuery) Order(o ...OrderFunc) *ProjectContributorQuery {
	pcq.order = append(pcq.order, o...)
	return pcq
}

// QueryProject chains the current query on the "project" edge.
func (pcq *ProjectContributorQuery) QueryProject() *ProjectQuery {
	query := &ProjectQuery{config: pcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(projectcontributor.Table, projectcontributor.FieldID, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, projectcontributor.ProjectTable, projectcontributor.ProjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(pcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (pcq *ProjectContributorQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: pcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(projectcontributor.Table, projectcontributor.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, projectcontributor.UserTable, projectcontributor.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(pcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProjectContributor entity from the query.
// Returns a *NotFoundError when no ProjectContributor was found.
func (pcq *ProjectContributorQuery) First(ctx context.Context) (*ProjectContributor, error) {
	nodes, err := pcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{projectcontributor.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pcq *ProjectContributorQuery) FirstX(ctx context.Context) *ProjectContributor {
	node, err := pcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProjectContributor ID from the query.
// Returns a *NotFoundError when no ProjectContributor ID was found.
func (pcq *ProjectContributorQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{projectcontributor.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pcq *ProjectContributorQuery) FirstIDX(ctx context.Context) int {
	id, err := pcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProjectContributor entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ProjectContributor entity is not found.
// Returns a *NotFoundError when no ProjectContributor entities are found.
func (pcq *ProjectContributorQuery) Only(ctx context.Context) (*ProjectContributor, error) {
	nodes, err := pcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{projectcontributor.Label}
	default:
		return nil, &NotSingularError{projectcontributor.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pcq *ProjectContributorQuery) OnlyX(ctx context.Context) *ProjectContributor {
	node, err := pcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProjectContributor ID in the query.
// Returns a *NotSingularError when exactly one ProjectContributor ID is not found.
// Returns a *NotFoundError when no entities are found.
func (pcq *ProjectContributorQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{projectcontributor.Label}
	default:
		err = &NotSingularError{projectcontributor.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pcq *ProjectContributorQuery) OnlyIDX(ctx context.Context) int {
	id, err := pcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProjectContributors.
func (pcq *ProjectContributorQuery) All(ctx context.Context) ([]*ProjectContributor, error) {
	if err := pcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pcq *ProjectContributorQuery) AllX(ctx context.Context) []*ProjectContributor {
	nodes, err := pcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProjectContributor IDs.
func (pcq *ProjectContributorQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pcq.Select(projectcontributor.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pcq *ProjectContributorQuery) IDsX(ctx context.Context) []int {
	ids, err := pcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pcq *ProjectContributorQuery) Count(ctx context.Context) (int, error) {
	if err := pcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pcq *ProjectContributorQuery) CountX(ctx context.Context) int {
	count, err := pcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pcq *ProjectContributorQuery) Exist(ctx context.Context) (bool, error) {
	if err := pcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pcq *ProjectContributorQuery) ExistX(ctx context.Context) bool {
	exist, err := pcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProjectContributorQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pcq *ProjectContributorQuery) Clone() *ProjectContributorQuery {
	if pcq == nil {
		return nil
	}
	return &ProjectContributorQuery{
		config:      pcq.config,
		limit:       pcq.limit,
		offset:      pcq.offset,
		order:       append([]OrderFunc{}, pcq.order...),
		predicates:  append([]predicate.ProjectContributor{}, pcq.predicates...),
		withProject: pcq.withProject.Clone(),
		withUser:    pcq.withUser.Clone(),
		// clone intermediate query.
		sql:  pcq.sql.Clone(),
		path: pcq.path,
	}
}

// WithProject tells the query-builder to eager-load the nodes that are connected to
// the "project" edge. The optional arguments are used to configure the query builder of the edge.
func (pcq *ProjectContributorQuery) WithProject(opts ...func(*ProjectQuery)) *ProjectContributorQuery {
	query := &ProjectQuery{config: pcq.config}
	for _, opt := range opts {
		opt(query)
	}
	pcq.withProject = query
	return pcq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (pcq *ProjectContributorQuery) WithUser(opts ...func(*UserQuery)) *ProjectContributorQuery {
	query := &UserQuery{config: pcq.config}
	for _, opt := range opts {
		opt(query)
	}
	pcq.withUser = query
	return pcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Role projectcontributor.Role `json:"role,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ProjectContributor.Query().
//		GroupBy(projectcontributor.FieldRole).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pcq *ProjectContributorQuery) GroupBy(field string, fields ...string) *ProjectContributorGroupBy {
	group := &ProjectContributorGroupBy{config: pcq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pcq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Role projectcontributor.Role `json:"role,omitempty"`
//	}
//
//	client.ProjectContributor.Query().
//		Select(projectcontributor.FieldRole).
//		Scan(ctx, &v)
//
func (pcq *ProjectContributorQuery) Select(field string, fields ...string) *ProjectContributorSelect {
	pcq.fields = append([]string{field}, fields...)
	return &ProjectContributorSelect{ProjectContributorQuery: pcq}
}

func (pcq *ProjectContributorQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pcq.fields {
		if !projectcontributor.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pcq.path != nil {
		prev, err := pcq.path(ctx)
		if err != nil {
			return err
		}
		pcq.sql = prev
	}
	if projectcontributor.Policy == nil {
		return errors.New("ent: uninitialized projectcontributor.Policy (forgotten import ent/runtime?)")
	}
	if err := projectcontributor.Policy.EvalQuery(ctx, pcq); err != nil {
		return err
	}
	return nil
}

func (pcq *ProjectContributorQuery) sqlAll(ctx context.Context) ([]*ProjectContributor, error) {
	var (
		nodes       = []*ProjectContributor{}
		withFKs     = pcq.withFKs
		_spec       = pcq.querySpec()
		loadedTypes = [2]bool{
			pcq.withProject != nil,
			pcq.withUser != nil,
		}
	)
	if pcq.withProject != nil || pcq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, projectcontributor.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ProjectContributor{config: pcq.config}
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
	if err := sqlgraph.QueryNodes(ctx, pcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pcq.withProject; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ProjectContributor)
		for i := range nodes {
			if nodes[i].project_contributors == nil {
				continue
			}
			fk := *nodes[i].project_contributors
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(project.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "project_contributors" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Project = n
			}
		}
	}

	if query := pcq.withUser; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ProjectContributor)
		for i := range nodes {
			if nodes[i].user_project_contributions == nil {
				continue
			}
			fk := *nodes[i].user_project_contributions
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_project_contributions" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	return nodes, nil
}

func (pcq *ProjectContributorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pcq.querySpec()
	return sqlgraph.CountNodes(ctx, pcq.driver, _spec)
}

func (pcq *ProjectContributorQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pcq *ProjectContributorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   projectcontributor.Table,
			Columns: projectcontributor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: projectcontributor.FieldID,
			},
		},
		From:   pcq.sql,
		Unique: true,
	}
	if unique := pcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, projectcontributor.FieldID)
		for i := range fields {
			if fields[i] != projectcontributor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pcq *ProjectContributorQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pcq.driver.Dialect())
	t1 := builder.Table(projectcontributor.Table)
	columns := pcq.fields
	if len(columns) == 0 {
		columns = projectcontributor.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pcq.sql != nil {
		selector = pcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range pcq.predicates {
		p(selector)
	}
	for _, p := range pcq.order {
		p(selector)
	}
	if offset := pcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProjectContributorGroupBy is the group-by builder for ProjectContributor entities.
type ProjectContributorGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pcgb *ProjectContributorGroupBy) Aggregate(fns ...AggregateFunc) *ProjectContributorGroupBy {
	pcgb.fns = append(pcgb.fns, fns...)
	return pcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pcgb *ProjectContributorGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pcgb.path(ctx)
	if err != nil {
		return err
	}
	pcgb.sql = query
	return pcgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pcgb *ProjectContributorGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pcgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (pcgb *ProjectContributorGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pcgb.fields) > 1 {
		return nil, errors.New("ent: ProjectContributorGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pcgb *ProjectContributorGroupBy) StringsX(ctx context.Context) []string {
	v, err := pcgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pcgb *ProjectContributorGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pcgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projectcontributor.Label}
	default:
		err = fmt.Errorf("ent: ProjectContributorGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pcgb *ProjectContributorGroupBy) StringX(ctx context.Context) string {
	v, err := pcgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (pcgb *ProjectContributorGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pcgb.fields) > 1 {
		return nil, errors.New("ent: ProjectContributorGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pcgb *ProjectContributorGroupBy) IntsX(ctx context.Context) []int {
	v, err := pcgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pcgb *ProjectContributorGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pcgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projectcontributor.Label}
	default:
		err = fmt.Errorf("ent: ProjectContributorGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pcgb *ProjectContributorGroupBy) IntX(ctx context.Context) int {
	v, err := pcgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (pcgb *ProjectContributorGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pcgb.fields) > 1 {
		return nil, errors.New("ent: ProjectContributorGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pcgb *ProjectContributorGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pcgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pcgb *ProjectContributorGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pcgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projectcontributor.Label}
	default:
		err = fmt.Errorf("ent: ProjectContributorGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pcgb *ProjectContributorGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pcgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (pcgb *ProjectContributorGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pcgb.fields) > 1 {
		return nil, errors.New("ent: ProjectContributorGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pcgb *ProjectContributorGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pcgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pcgb *ProjectContributorGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pcgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projectcontributor.Label}
	default:
		err = fmt.Errorf("ent: ProjectContributorGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pcgb *ProjectContributorGroupBy) BoolX(ctx context.Context) bool {
	v, err := pcgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pcgb *ProjectContributorGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pcgb.fields {
		if !projectcontributor.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pcgb *ProjectContributorGroupBy) sqlQuery() *sql.Selector {
	selector := pcgb.sql.Select()
	aggregation := make([]string, 0, len(pcgb.fns))
	for _, fn := range pcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pcgb.fields)+len(pcgb.fns))
		for _, f := range pcgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pcgb.fields...)...)
}

// ProjectContributorSelect is the builder for selecting fields of ProjectContributor entities.
type ProjectContributorSelect struct {
	*ProjectContributorQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pcs *ProjectContributorSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pcs.prepareQuery(ctx); err != nil {
		return err
	}
	pcs.sql = pcs.ProjectContributorQuery.sqlQuery(ctx)
	return pcs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pcs *ProjectContributorSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pcs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (pcs *ProjectContributorSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pcs.fields) > 1 {
		return nil, errors.New("ent: ProjectContributorSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pcs *ProjectContributorSelect) StringsX(ctx context.Context) []string {
	v, err := pcs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (pcs *ProjectContributorSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pcs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projectcontributor.Label}
	default:
		err = fmt.Errorf("ent: ProjectContributorSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pcs *ProjectContributorSelect) StringX(ctx context.Context) string {
	v, err := pcs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (pcs *ProjectContributorSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pcs.fields) > 1 {
		return nil, errors.New("ent: ProjectContributorSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pcs *ProjectContributorSelect) IntsX(ctx context.Context) []int {
	v, err := pcs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (pcs *ProjectContributorSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pcs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projectcontributor.Label}
	default:
		err = fmt.Errorf("ent: ProjectContributorSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pcs *ProjectContributorSelect) IntX(ctx context.Context) int {
	v, err := pcs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (pcs *ProjectContributorSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pcs.fields) > 1 {
		return nil, errors.New("ent: ProjectContributorSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pcs *ProjectContributorSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pcs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (pcs *ProjectContributorSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pcs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projectcontributor.Label}
	default:
		err = fmt.Errorf("ent: ProjectContributorSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pcs *ProjectContributorSelect) Float64X(ctx context.Context) float64 {
	v, err := pcs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (pcs *ProjectContributorSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pcs.fields) > 1 {
		return nil, errors.New("ent: ProjectContributorSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pcs *ProjectContributorSelect) BoolsX(ctx context.Context) []bool {
	v, err := pcs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (pcs *ProjectContributorSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pcs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{projectcontributor.Label}
	default:
		err = fmt.Errorf("ent: ProjectContributorSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pcs *ProjectContributorSelect) BoolX(ctx context.Context) bool {
	v, err := pcs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pcs *ProjectContributorSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pcs.sql.Query()
	if err := pcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
