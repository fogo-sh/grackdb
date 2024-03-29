// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/predicate"
	"github.com/fogo-sh/grackdb/ent/projecttechnology"
	"github.com/fogo-sh/grackdb/ent/repositorytechnology"
	"github.com/fogo-sh/grackdb/ent/technology"
	"github.com/fogo-sh/grackdb/ent/technologyassociation"
)

// TechnologyQuery is the builder for querying Technology entities.
type TechnologyQuery struct {
	config
	limit                       *int
	offset                      *int
	unique                      *bool
	order                       []OrderFunc
	fields                      []string
	predicates                  []predicate.Technology
	withParentTechnologies      *TechnologyAssociationQuery
	withChildTechnologies       *TechnologyAssociationQuery
	withProjects                *ProjectTechnologyQuery
	withRepositories            *RepositoryTechnologyQuery
	modifiers                   []func(*sql.Selector)
	loadTotal                   []func(context.Context, []*Technology) error
	withNamedParentTechnologies map[string]*TechnologyAssociationQuery
	withNamedChildTechnologies  map[string]*TechnologyAssociationQuery
	withNamedProjects           map[string]*ProjectTechnologyQuery
	withNamedRepositories       map[string]*RepositoryTechnologyQuery
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

// QueryParentTechnologies chains the current query on the "parent_technologies" edge.
func (tq *TechnologyQuery) QueryParentTechnologies() *TechnologyAssociationQuery {
	query := &TechnologyAssociationQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(technology.Table, technology.FieldID, selector),
			sqlgraph.To(technologyassociation.Table, technologyassociation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, technology.ParentTechnologiesTable, technology.ParentTechnologiesColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildTechnologies chains the current query on the "child_technologies" edge.
func (tq *TechnologyQuery) QueryChildTechnologies() *TechnologyAssociationQuery {
	query := &TechnologyAssociationQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(technology.Table, technology.FieldID, selector),
			sqlgraph.To(technologyassociation.Table, technologyassociation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, technology.ChildTechnologiesTable, technology.ChildTechnologiesColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProjects chains the current query on the "projects" edge.
func (tq *TechnologyQuery) QueryProjects() *ProjectTechnologyQuery {
	query := &ProjectTechnologyQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(technology.Table, technology.FieldID, selector),
			sqlgraph.To(projecttechnology.Table, projecttechnology.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, technology.ProjectsTable, technology.ProjectsColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRepositories chains the current query on the "repositories" edge.
func (tq *TechnologyQuery) QueryRepositories() *RepositoryTechnologyQuery {
	query := &RepositoryTechnologyQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(technology.Table, technology.FieldID, selector),
			sqlgraph.To(repositorytechnology.Table, repositorytechnology.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, technology.RepositoriesTable, technology.RepositoriesColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
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
// Returns a *NotSingularError when more than one Technology entity is found.
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
// Returns a *NotSingularError when more than one Technology ID is found.
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
		config:                 tq.config,
		limit:                  tq.limit,
		offset:                 tq.offset,
		order:                  append([]OrderFunc{}, tq.order...),
		predicates:             append([]predicate.Technology{}, tq.predicates...),
		withParentTechnologies: tq.withParentTechnologies.Clone(),
		withChildTechnologies:  tq.withChildTechnologies.Clone(),
		withProjects:           tq.withProjects.Clone(),
		withRepositories:       tq.withRepositories.Clone(),
		// clone intermediate query.
		sql:    tq.sql.Clone(),
		path:   tq.path,
		unique: tq.unique,
	}
}

// WithParentTechnologies tells the query-builder to eager-load the nodes that are connected to
// the "parent_technologies" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TechnologyQuery) WithParentTechnologies(opts ...func(*TechnologyAssociationQuery)) *TechnologyQuery {
	query := &TechnologyAssociationQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withParentTechnologies = query
	return tq
}

// WithChildTechnologies tells the query-builder to eager-load the nodes that are connected to
// the "child_technologies" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TechnologyQuery) WithChildTechnologies(opts ...func(*TechnologyAssociationQuery)) *TechnologyQuery {
	query := &TechnologyAssociationQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withChildTechnologies = query
	return tq
}

// WithProjects tells the query-builder to eager-load the nodes that are connected to
// the "projects" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TechnologyQuery) WithProjects(opts ...func(*ProjectTechnologyQuery)) *TechnologyQuery {
	query := &ProjectTechnologyQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withProjects = query
	return tq
}

// WithRepositories tells the query-builder to eager-load the nodes that are connected to
// the "repositories" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TechnologyQuery) WithRepositories(opts ...func(*RepositoryTechnologyQuery)) *TechnologyQuery {
	query := &RepositoryTechnologyQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withRepositories = query
	return tq
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
func (tq *TechnologyQuery) GroupBy(field string, fields ...string) *TechnologyGroupBy {
	grbuild := &TechnologyGroupBy{config: tq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tq.sqlQuery(ctx), nil
	}
	grbuild.label = technology.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
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
func (tq *TechnologyQuery) Select(fields ...string) *TechnologySelect {
	tq.fields = append(tq.fields, fields...)
	selbuild := &TechnologySelect{TechnologyQuery: tq}
	selbuild.label = technology.Label
	selbuild.flds, selbuild.scan = &tq.fields, selbuild.Scan
	return selbuild
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

func (tq *TechnologyQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Technology, error) {
	var (
		nodes       = []*Technology{}
		_spec       = tq.querySpec()
		loadedTypes = [4]bool{
			tq.withParentTechnologies != nil,
			tq.withChildTechnologies != nil,
			tq.withProjects != nil,
			tq.withRepositories != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Technology).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Technology{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withParentTechnologies; query != nil {
		if err := tq.loadParentTechnologies(ctx, query, nodes,
			func(n *Technology) { n.Edges.ParentTechnologies = []*TechnologyAssociation{} },
			func(n *Technology, e *TechnologyAssociation) {
				n.Edges.ParentTechnologies = append(n.Edges.ParentTechnologies, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := tq.withChildTechnologies; query != nil {
		if err := tq.loadChildTechnologies(ctx, query, nodes,
			func(n *Technology) { n.Edges.ChildTechnologies = []*TechnologyAssociation{} },
			func(n *Technology, e *TechnologyAssociation) {
				n.Edges.ChildTechnologies = append(n.Edges.ChildTechnologies, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := tq.withProjects; query != nil {
		if err := tq.loadProjects(ctx, query, nodes,
			func(n *Technology) { n.Edges.Projects = []*ProjectTechnology{} },
			func(n *Technology, e *ProjectTechnology) { n.Edges.Projects = append(n.Edges.Projects, e) }); err != nil {
			return nil, err
		}
	}
	if query := tq.withRepositories; query != nil {
		if err := tq.loadRepositories(ctx, query, nodes,
			func(n *Technology) { n.Edges.Repositories = []*RepositoryTechnology{} },
			func(n *Technology, e *RepositoryTechnology) { n.Edges.Repositories = append(n.Edges.Repositories, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range tq.withNamedParentTechnologies {
		if err := tq.loadParentTechnologies(ctx, query, nodes,
			func(n *Technology) { n.appendNamedParentTechnologies(name) },
			func(n *Technology, e *TechnologyAssociation) { n.appendNamedParentTechnologies(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range tq.withNamedChildTechnologies {
		if err := tq.loadChildTechnologies(ctx, query, nodes,
			func(n *Technology) { n.appendNamedChildTechnologies(name) },
			func(n *Technology, e *TechnologyAssociation) { n.appendNamedChildTechnologies(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range tq.withNamedProjects {
		if err := tq.loadProjects(ctx, query, nodes,
			func(n *Technology) { n.appendNamedProjects(name) },
			func(n *Technology, e *ProjectTechnology) { n.appendNamedProjects(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range tq.withNamedRepositories {
		if err := tq.loadRepositories(ctx, query, nodes,
			func(n *Technology) { n.appendNamedRepositories(name) },
			func(n *Technology, e *RepositoryTechnology) { n.appendNamedRepositories(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range tq.loadTotal {
		if err := tq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TechnologyQuery) loadParentTechnologies(ctx context.Context, query *TechnologyAssociationQuery, nodes []*Technology, init func(*Technology), assign func(*Technology, *TechnologyAssociation)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Technology)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.TechnologyAssociation(func(s *sql.Selector) {
		s.Where(sql.InValues(technology.ParentTechnologiesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.technology_parent_technologies
		if fk == nil {
			return fmt.Errorf(`foreign-key "technology_parent_technologies" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "technology_parent_technologies" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (tq *TechnologyQuery) loadChildTechnologies(ctx context.Context, query *TechnologyAssociationQuery, nodes []*Technology, init func(*Technology), assign func(*Technology, *TechnologyAssociation)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Technology)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.TechnologyAssociation(func(s *sql.Selector) {
		s.Where(sql.InValues(technology.ChildTechnologiesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.technology_child_technologies
		if fk == nil {
			return fmt.Errorf(`foreign-key "technology_child_technologies" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "technology_child_technologies" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (tq *TechnologyQuery) loadProjects(ctx context.Context, query *ProjectTechnologyQuery, nodes []*Technology, init func(*Technology), assign func(*Technology, *ProjectTechnology)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Technology)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.ProjectTechnology(func(s *sql.Selector) {
		s.Where(sql.InValues(technology.ProjectsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.technology_projects
		if fk == nil {
			return fmt.Errorf(`foreign-key "technology_projects" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "technology_projects" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (tq *TechnologyQuery) loadRepositories(ctx context.Context, query *RepositoryTechnologyQuery, nodes []*Technology, init func(*Technology), assign func(*Technology, *RepositoryTechnology)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Technology)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.RepositoryTechnology(func(s *sql.Selector) {
		s.Where(sql.InValues(technology.RepositoriesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.technology_repositories
		if fk == nil {
			return fmt.Errorf(`foreign-key "technology_repositories" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "technology_repositories" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (tq *TechnologyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	_spec.Node.Columns = tq.fields
	if len(tq.fields) > 0 {
		_spec.Unique = tq.unique != nil && *tq.unique
	}
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
	if tq.unique != nil && *tq.unique {
		selector.Distinct()
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

// WithNamedParentTechnologies tells the query-builder to eager-load the nodes that are connected to the "parent_technologies"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (tq *TechnologyQuery) WithNamedParentTechnologies(name string, opts ...func(*TechnologyAssociationQuery)) *TechnologyQuery {
	query := &TechnologyAssociationQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	if tq.withNamedParentTechnologies == nil {
		tq.withNamedParentTechnologies = make(map[string]*TechnologyAssociationQuery)
	}
	tq.withNamedParentTechnologies[name] = query
	return tq
}

// WithNamedChildTechnologies tells the query-builder to eager-load the nodes that are connected to the "child_technologies"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (tq *TechnologyQuery) WithNamedChildTechnologies(name string, opts ...func(*TechnologyAssociationQuery)) *TechnologyQuery {
	query := &TechnologyAssociationQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	if tq.withNamedChildTechnologies == nil {
		tq.withNamedChildTechnologies = make(map[string]*TechnologyAssociationQuery)
	}
	tq.withNamedChildTechnologies[name] = query
	return tq
}

// WithNamedProjects tells the query-builder to eager-load the nodes that are connected to the "projects"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (tq *TechnologyQuery) WithNamedProjects(name string, opts ...func(*ProjectTechnologyQuery)) *TechnologyQuery {
	query := &ProjectTechnologyQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	if tq.withNamedProjects == nil {
		tq.withNamedProjects = make(map[string]*ProjectTechnologyQuery)
	}
	tq.withNamedProjects[name] = query
	return tq
}

// WithNamedRepositories tells the query-builder to eager-load the nodes that are connected to the "repositories"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (tq *TechnologyQuery) WithNamedRepositories(name string, opts ...func(*RepositoryTechnologyQuery)) *TechnologyQuery {
	query := &RepositoryTechnologyQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	if tq.withNamedRepositories == nil {
		tq.withNamedRepositories = make(map[string]*RepositoryTechnologyQuery)
	}
	tq.withNamedRepositories[name] = query
	return tq
}

// TechnologyGroupBy is the group-by builder for Technology entities.
type TechnologyGroupBy struct {
	config
	selector
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
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tgb.fields...)...)
}

// TechnologySelect is the builder for selecting fields of Technology entities.
type TechnologySelect struct {
	*TechnologyQuery
	selector
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

func (ts *TechnologySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ts.sql.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
