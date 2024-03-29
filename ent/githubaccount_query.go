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
	"github.com/fogo-sh/grackdb/ent/githubaccount"
	"github.com/fogo-sh/grackdb/ent/githuborganizationmember"
	"github.com/fogo-sh/grackdb/ent/predicate"
	"github.com/fogo-sh/grackdb/ent/repository"
	"github.com/fogo-sh/grackdb/ent/user"
)

// GithubAccountQuery is the builder for querying GithubAccount entities.
type GithubAccountQuery struct {
	config
	limit                            *int
	offset                           *int
	unique                           *bool
	order                            []OrderFunc
	fields                           []string
	predicates                       []predicate.GithubAccount
	withOwner                        *UserQuery
	withOrganizationMemberships      *GithubOrganizationMemberQuery
	withRepositories                 *RepositoryQuery
	withFKs                          bool
	modifiers                        []func(*sql.Selector)
	loadTotal                        []func(context.Context, []*GithubAccount) error
	withNamedOrganizationMemberships map[string]*GithubOrganizationMemberQuery
	withNamedRepositories            map[string]*RepositoryQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GithubAccountQuery builder.
func (gaq *GithubAccountQuery) Where(ps ...predicate.GithubAccount) *GithubAccountQuery {
	gaq.predicates = append(gaq.predicates, ps...)
	return gaq
}

// Limit adds a limit step to the query.
func (gaq *GithubAccountQuery) Limit(limit int) *GithubAccountQuery {
	gaq.limit = &limit
	return gaq
}

// Offset adds an offset step to the query.
func (gaq *GithubAccountQuery) Offset(offset int) *GithubAccountQuery {
	gaq.offset = &offset
	return gaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gaq *GithubAccountQuery) Unique(unique bool) *GithubAccountQuery {
	gaq.unique = &unique
	return gaq
}

// Order adds an order step to the query.
func (gaq *GithubAccountQuery) Order(o ...OrderFunc) *GithubAccountQuery {
	gaq.order = append(gaq.order, o...)
	return gaq
}

// QueryOwner chains the current query on the "owner" edge.
func (gaq *GithubAccountQuery) QueryOwner() *UserQuery {
	query := &UserQuery{config: gaq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(githubaccount.Table, githubaccount.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, githubaccount.OwnerTable, githubaccount.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(gaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrganizationMemberships chains the current query on the "organization_memberships" edge.
func (gaq *GithubAccountQuery) QueryOrganizationMemberships() *GithubOrganizationMemberQuery {
	query := &GithubOrganizationMemberQuery{config: gaq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(githubaccount.Table, githubaccount.FieldID, selector),
			sqlgraph.To(githuborganizationmember.Table, githuborganizationmember.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, githubaccount.OrganizationMembershipsTable, githubaccount.OrganizationMembershipsColumn),
		)
		fromU = sqlgraph.SetNeighbors(gaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRepositories chains the current query on the "repositories" edge.
func (gaq *GithubAccountQuery) QueryRepositories() *RepositoryQuery {
	query := &RepositoryQuery{config: gaq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(githubaccount.Table, githubaccount.FieldID, selector),
			sqlgraph.To(repository.Table, repository.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, githubaccount.RepositoriesTable, githubaccount.RepositoriesColumn),
		)
		fromU = sqlgraph.SetNeighbors(gaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GithubAccount entity from the query.
// Returns a *NotFoundError when no GithubAccount was found.
func (gaq *GithubAccountQuery) First(ctx context.Context) (*GithubAccount, error) {
	nodes, err := gaq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{githubaccount.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gaq *GithubAccountQuery) FirstX(ctx context.Context) *GithubAccount {
	node, err := gaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GithubAccount ID from the query.
// Returns a *NotFoundError when no GithubAccount ID was found.
func (gaq *GithubAccountQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gaq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{githubaccount.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gaq *GithubAccountQuery) FirstIDX(ctx context.Context) int {
	id, err := gaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GithubAccount entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GithubAccount entity is found.
// Returns a *NotFoundError when no GithubAccount entities are found.
func (gaq *GithubAccountQuery) Only(ctx context.Context) (*GithubAccount, error) {
	nodes, err := gaq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{githubaccount.Label}
	default:
		return nil, &NotSingularError{githubaccount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gaq *GithubAccountQuery) OnlyX(ctx context.Context) *GithubAccount {
	node, err := gaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GithubAccount ID in the query.
// Returns a *NotSingularError when more than one GithubAccount ID is found.
// Returns a *NotFoundError when no entities are found.
func (gaq *GithubAccountQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gaq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{githubaccount.Label}
	default:
		err = &NotSingularError{githubaccount.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gaq *GithubAccountQuery) OnlyIDX(ctx context.Context) int {
	id, err := gaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GithubAccounts.
func (gaq *GithubAccountQuery) All(ctx context.Context) ([]*GithubAccount, error) {
	if err := gaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gaq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gaq *GithubAccountQuery) AllX(ctx context.Context) []*GithubAccount {
	nodes, err := gaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GithubAccount IDs.
func (gaq *GithubAccountQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := gaq.Select(githubaccount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gaq *GithubAccountQuery) IDsX(ctx context.Context) []int {
	ids, err := gaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gaq *GithubAccountQuery) Count(ctx context.Context) (int, error) {
	if err := gaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gaq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gaq *GithubAccountQuery) CountX(ctx context.Context) int {
	count, err := gaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gaq *GithubAccountQuery) Exist(ctx context.Context) (bool, error) {
	if err := gaq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gaq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gaq *GithubAccountQuery) ExistX(ctx context.Context) bool {
	exist, err := gaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GithubAccountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gaq *GithubAccountQuery) Clone() *GithubAccountQuery {
	if gaq == nil {
		return nil
	}
	return &GithubAccountQuery{
		config:                      gaq.config,
		limit:                       gaq.limit,
		offset:                      gaq.offset,
		order:                       append([]OrderFunc{}, gaq.order...),
		predicates:                  append([]predicate.GithubAccount{}, gaq.predicates...),
		withOwner:                   gaq.withOwner.Clone(),
		withOrganizationMemberships: gaq.withOrganizationMemberships.Clone(),
		withRepositories:            gaq.withRepositories.Clone(),
		// clone intermediate query.
		sql:    gaq.sql.Clone(),
		path:   gaq.path,
		unique: gaq.unique,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (gaq *GithubAccountQuery) WithOwner(opts ...func(*UserQuery)) *GithubAccountQuery {
	query := &UserQuery{config: gaq.config}
	for _, opt := range opts {
		opt(query)
	}
	gaq.withOwner = query
	return gaq
}

// WithOrganizationMemberships tells the query-builder to eager-load the nodes that are connected to
// the "organization_memberships" edge. The optional arguments are used to configure the query builder of the edge.
func (gaq *GithubAccountQuery) WithOrganizationMemberships(opts ...func(*GithubOrganizationMemberQuery)) *GithubAccountQuery {
	query := &GithubOrganizationMemberQuery{config: gaq.config}
	for _, opt := range opts {
		opt(query)
	}
	gaq.withOrganizationMemberships = query
	return gaq
}

// WithRepositories tells the query-builder to eager-load the nodes that are connected to
// the "repositories" edge. The optional arguments are used to configure the query builder of the edge.
func (gaq *GithubAccountQuery) WithRepositories(opts ...func(*RepositoryQuery)) *GithubAccountQuery {
	query := &RepositoryQuery{config: gaq.config}
	for _, opt := range opts {
		opt(query)
	}
	gaq.withRepositories = query
	return gaq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Username string `json:"username,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GithubAccount.Query().
//		GroupBy(githubaccount.FieldUsername).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (gaq *GithubAccountQuery) GroupBy(field string, fields ...string) *GithubAccountGroupBy {
	grbuild := &GithubAccountGroupBy{config: gaq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gaq.sqlQuery(ctx), nil
	}
	grbuild.label = githubaccount.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Username string `json:"username,omitempty"`
//	}
//
//	client.GithubAccount.Query().
//		Select(githubaccount.FieldUsername).
//		Scan(ctx, &v)
func (gaq *GithubAccountQuery) Select(fields ...string) *GithubAccountSelect {
	gaq.fields = append(gaq.fields, fields...)
	selbuild := &GithubAccountSelect{GithubAccountQuery: gaq}
	selbuild.label = githubaccount.Label
	selbuild.flds, selbuild.scan = &gaq.fields, selbuild.Scan
	return selbuild
}

func (gaq *GithubAccountQuery) prepareQuery(ctx context.Context) error {
	for _, f := range gaq.fields {
		if !githubaccount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gaq.path != nil {
		prev, err := gaq.path(ctx)
		if err != nil {
			return err
		}
		gaq.sql = prev
	}
	if githubaccount.Policy == nil {
		return errors.New("ent: uninitialized githubaccount.Policy (forgotten import ent/runtime?)")
	}
	if err := githubaccount.Policy.EvalQuery(ctx, gaq); err != nil {
		return err
	}
	return nil
}

func (gaq *GithubAccountQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GithubAccount, error) {
	var (
		nodes       = []*GithubAccount{}
		withFKs     = gaq.withFKs
		_spec       = gaq.querySpec()
		loadedTypes = [3]bool{
			gaq.withOwner != nil,
			gaq.withOrganizationMemberships != nil,
			gaq.withRepositories != nil,
		}
	)
	if gaq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, githubaccount.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*GithubAccount).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &GithubAccount{config: gaq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
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
	if query := gaq.withOwner; query != nil {
		if err := gaq.loadOwner(ctx, query, nodes, nil,
			func(n *GithubAccount, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	if query := gaq.withOrganizationMemberships; query != nil {
		if err := gaq.loadOrganizationMemberships(ctx, query, nodes,
			func(n *GithubAccount) { n.Edges.OrganizationMemberships = []*GithubOrganizationMember{} },
			func(n *GithubAccount, e *GithubOrganizationMember) {
				n.Edges.OrganizationMemberships = append(n.Edges.OrganizationMemberships, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := gaq.withRepositories; query != nil {
		if err := gaq.loadRepositories(ctx, query, nodes,
			func(n *GithubAccount) { n.Edges.Repositories = []*Repository{} },
			func(n *GithubAccount, e *Repository) { n.Edges.Repositories = append(n.Edges.Repositories, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range gaq.withNamedOrganizationMemberships {
		if err := gaq.loadOrganizationMemberships(ctx, query, nodes,
			func(n *GithubAccount) { n.appendNamedOrganizationMemberships(name) },
			func(n *GithubAccount, e *GithubOrganizationMember) { n.appendNamedOrganizationMemberships(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range gaq.withNamedRepositories {
		if err := gaq.loadRepositories(ctx, query, nodes,
			func(n *GithubAccount) { n.appendNamedRepositories(name) },
			func(n *GithubAccount, e *Repository) { n.appendNamedRepositories(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range gaq.loadTotal {
		if err := gaq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (gaq *GithubAccountQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*GithubAccount, init func(*GithubAccount), assign func(*GithubAccount, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*GithubAccount)
	for i := range nodes {
		if nodes[i].user_github_accounts == nil {
			continue
		}
		fk := *nodes[i].user_github_accounts
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_github_accounts" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (gaq *GithubAccountQuery) loadOrganizationMemberships(ctx context.Context, query *GithubOrganizationMemberQuery, nodes []*GithubAccount, init func(*GithubAccount), assign func(*GithubAccount, *GithubOrganizationMember)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*GithubAccount)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.GithubOrganizationMember(func(s *sql.Selector) {
		s.Where(sql.InValues(githubaccount.OrganizationMembershipsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.github_account_organization_memberships
		if fk == nil {
			return fmt.Errorf(`foreign-key "github_account_organization_memberships" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "github_account_organization_memberships" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (gaq *GithubAccountQuery) loadRepositories(ctx context.Context, query *RepositoryQuery, nodes []*GithubAccount, init func(*GithubAccount), assign func(*GithubAccount, *Repository)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*GithubAccount)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.InValues(githubaccount.RepositoriesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.github_account_repositories
		if fk == nil {
			return fmt.Errorf(`foreign-key "github_account_repositories" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "github_account_repositories" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (gaq *GithubAccountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gaq.querySpec()
	if len(gaq.modifiers) > 0 {
		_spec.Modifiers = gaq.modifiers
	}
	_spec.Node.Columns = gaq.fields
	if len(gaq.fields) > 0 {
		_spec.Unique = gaq.unique != nil && *gaq.unique
	}
	return sqlgraph.CountNodes(ctx, gaq.driver, _spec)
}

func (gaq *GithubAccountQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gaq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (gaq *GithubAccountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   githubaccount.Table,
			Columns: githubaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: githubaccount.FieldID,
			},
		},
		From:   gaq.sql,
		Unique: true,
	}
	if unique := gaq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gaq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, githubaccount.FieldID)
		for i := range fields {
			if fields[i] != githubaccount.FieldID {
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
	if limit := gaq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gaq.offset; offset != nil {
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

func (gaq *GithubAccountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gaq.driver.Dialect())
	t1 := builder.Table(githubaccount.Table)
	columns := gaq.fields
	if len(columns) == 0 {
		columns = githubaccount.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gaq.sql != nil {
		selector = gaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gaq.unique != nil && *gaq.unique {
		selector.Distinct()
	}
	for _, p := range gaq.predicates {
		p(selector)
	}
	for _, p := range gaq.order {
		p(selector)
	}
	if offset := gaq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gaq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedOrganizationMemberships tells the query-builder to eager-load the nodes that are connected to the "organization_memberships"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (gaq *GithubAccountQuery) WithNamedOrganizationMemberships(name string, opts ...func(*GithubOrganizationMemberQuery)) *GithubAccountQuery {
	query := &GithubOrganizationMemberQuery{config: gaq.config}
	for _, opt := range opts {
		opt(query)
	}
	if gaq.withNamedOrganizationMemberships == nil {
		gaq.withNamedOrganizationMemberships = make(map[string]*GithubOrganizationMemberQuery)
	}
	gaq.withNamedOrganizationMemberships[name] = query
	return gaq
}

// WithNamedRepositories tells the query-builder to eager-load the nodes that are connected to the "repositories"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (gaq *GithubAccountQuery) WithNamedRepositories(name string, opts ...func(*RepositoryQuery)) *GithubAccountQuery {
	query := &RepositoryQuery{config: gaq.config}
	for _, opt := range opts {
		opt(query)
	}
	if gaq.withNamedRepositories == nil {
		gaq.withNamedRepositories = make(map[string]*RepositoryQuery)
	}
	gaq.withNamedRepositories[name] = query
	return gaq
}

// GithubAccountGroupBy is the group-by builder for GithubAccount entities.
type GithubAccountGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gagb *GithubAccountGroupBy) Aggregate(fns ...AggregateFunc) *GithubAccountGroupBy {
	gagb.fns = append(gagb.fns, fns...)
	return gagb
}

// Scan applies the group-by query and scans the result into the given value.
func (gagb *GithubAccountGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := gagb.path(ctx)
	if err != nil {
		return err
	}
	gagb.sql = query
	return gagb.sqlScan(ctx, v)
}

func (gagb *GithubAccountGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range gagb.fields {
		if !githubaccount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := gagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gagb *GithubAccountGroupBy) sqlQuery() *sql.Selector {
	selector := gagb.sql.Select()
	aggregation := make([]string, 0, len(gagb.fns))
	for _, fn := range gagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(gagb.fields)+len(gagb.fns))
		for _, f := range gagb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(gagb.fields...)...)
}

// GithubAccountSelect is the builder for selecting fields of GithubAccount entities.
type GithubAccountSelect struct {
	*GithubAccountQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gas *GithubAccountSelect) Scan(ctx context.Context, v interface{}) error {
	if err := gas.prepareQuery(ctx); err != nil {
		return err
	}
	gas.sql = gas.GithubAccountQuery.sqlQuery(ctx)
	return gas.sqlScan(ctx, v)
}

func (gas *GithubAccountSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gas.sql.Query()
	if err := gas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
