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
	"github.com/fogo-sh/grackdb/ent/discordaccount"
	"github.com/fogo-sh/grackdb/ent/discordbot"
	"github.com/fogo-sh/grackdb/ent/predicate"
	"github.com/fogo-sh/grackdb/ent/user"
)

// DiscordAccountQuery is the builder for querying DiscordAccount entities.
type DiscordAccountQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.DiscordAccount
	// eager-loading edges.
	withOwner *UserQuery
	withBot   *DiscordBotQuery
	withFKs   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DiscordAccountQuery builder.
func (daq *DiscordAccountQuery) Where(ps ...predicate.DiscordAccount) *DiscordAccountQuery {
	daq.predicates = append(daq.predicates, ps...)
	return daq
}

// Limit adds a limit step to the query.
func (daq *DiscordAccountQuery) Limit(limit int) *DiscordAccountQuery {
	daq.limit = &limit
	return daq
}

// Offset adds an offset step to the query.
func (daq *DiscordAccountQuery) Offset(offset int) *DiscordAccountQuery {
	daq.offset = &offset
	return daq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (daq *DiscordAccountQuery) Unique(unique bool) *DiscordAccountQuery {
	daq.unique = &unique
	return daq
}

// Order adds an order step to the query.
func (daq *DiscordAccountQuery) Order(o ...OrderFunc) *DiscordAccountQuery {
	daq.order = append(daq.order, o...)
	return daq
}

// QueryOwner chains the current query on the "owner" edge.
func (daq *DiscordAccountQuery) QueryOwner() *UserQuery {
	query := &UserQuery{config: daq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := daq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := daq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(discordaccount.Table, discordaccount.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, discordaccount.OwnerTable, discordaccount.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(daq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBot chains the current query on the "bot" edge.
func (daq *DiscordAccountQuery) QueryBot() *DiscordBotQuery {
	query := &DiscordBotQuery{config: daq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := daq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := daq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(discordaccount.Table, discordaccount.FieldID, selector),
			sqlgraph.To(discordbot.Table, discordbot.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, discordaccount.BotTable, discordaccount.BotColumn),
		)
		fromU = sqlgraph.SetNeighbors(daq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DiscordAccount entity from the query.
// Returns a *NotFoundError when no DiscordAccount was found.
func (daq *DiscordAccountQuery) First(ctx context.Context) (*DiscordAccount, error) {
	nodes, err := daq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{discordaccount.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (daq *DiscordAccountQuery) FirstX(ctx context.Context) *DiscordAccount {
	node, err := daq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DiscordAccount ID from the query.
// Returns a *NotFoundError when no DiscordAccount ID was found.
func (daq *DiscordAccountQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = daq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{discordaccount.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (daq *DiscordAccountQuery) FirstIDX(ctx context.Context) int {
	id, err := daq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DiscordAccount entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one DiscordAccount entity is not found.
// Returns a *NotFoundError when no DiscordAccount entities are found.
func (daq *DiscordAccountQuery) Only(ctx context.Context) (*DiscordAccount, error) {
	nodes, err := daq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{discordaccount.Label}
	default:
		return nil, &NotSingularError{discordaccount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (daq *DiscordAccountQuery) OnlyX(ctx context.Context) *DiscordAccount {
	node, err := daq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DiscordAccount ID in the query.
// Returns a *NotSingularError when exactly one DiscordAccount ID is not found.
// Returns a *NotFoundError when no entities are found.
func (daq *DiscordAccountQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = daq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{discordaccount.Label}
	default:
		err = &NotSingularError{discordaccount.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (daq *DiscordAccountQuery) OnlyIDX(ctx context.Context) int {
	id, err := daq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DiscordAccounts.
func (daq *DiscordAccountQuery) All(ctx context.Context) ([]*DiscordAccount, error) {
	if err := daq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return daq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (daq *DiscordAccountQuery) AllX(ctx context.Context) []*DiscordAccount {
	nodes, err := daq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DiscordAccount IDs.
func (daq *DiscordAccountQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := daq.Select(discordaccount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (daq *DiscordAccountQuery) IDsX(ctx context.Context) []int {
	ids, err := daq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (daq *DiscordAccountQuery) Count(ctx context.Context) (int, error) {
	if err := daq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return daq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (daq *DiscordAccountQuery) CountX(ctx context.Context) int {
	count, err := daq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (daq *DiscordAccountQuery) Exist(ctx context.Context) (bool, error) {
	if err := daq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return daq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (daq *DiscordAccountQuery) ExistX(ctx context.Context) bool {
	exist, err := daq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DiscordAccountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (daq *DiscordAccountQuery) Clone() *DiscordAccountQuery {
	if daq == nil {
		return nil
	}
	return &DiscordAccountQuery{
		config:     daq.config,
		limit:      daq.limit,
		offset:     daq.offset,
		order:      append([]OrderFunc{}, daq.order...),
		predicates: append([]predicate.DiscordAccount{}, daq.predicates...),
		withOwner:  daq.withOwner.Clone(),
		withBot:    daq.withBot.Clone(),
		// clone intermediate query.
		sql:  daq.sql.Clone(),
		path: daq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (daq *DiscordAccountQuery) WithOwner(opts ...func(*UserQuery)) *DiscordAccountQuery {
	query := &UserQuery{config: daq.config}
	for _, opt := range opts {
		opt(query)
	}
	daq.withOwner = query
	return daq
}

// WithBot tells the query-builder to eager-load the nodes that are connected to
// the "bot" edge. The optional arguments are used to configure the query builder of the edge.
func (daq *DiscordAccountQuery) WithBot(opts ...func(*DiscordBotQuery)) *DiscordAccountQuery {
	query := &DiscordBotQuery{config: daq.config}
	for _, opt := range opts {
		opt(query)
	}
	daq.withBot = query
	return daq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DiscordID string `json:"discord_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DiscordAccount.Query().
//		GroupBy(discordaccount.FieldDiscordID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (daq *DiscordAccountQuery) GroupBy(field string, fields ...string) *DiscordAccountGroupBy {
	group := &DiscordAccountGroupBy{config: daq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := daq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return daq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DiscordID string `json:"discord_id,omitempty"`
//	}
//
//	client.DiscordAccount.Query().
//		Select(discordaccount.FieldDiscordID).
//		Scan(ctx, &v)
//
func (daq *DiscordAccountQuery) Select(fields ...string) *DiscordAccountSelect {
	daq.fields = append(daq.fields, fields...)
	return &DiscordAccountSelect{DiscordAccountQuery: daq}
}

func (daq *DiscordAccountQuery) prepareQuery(ctx context.Context) error {
	for _, f := range daq.fields {
		if !discordaccount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if daq.path != nil {
		prev, err := daq.path(ctx)
		if err != nil {
			return err
		}
		daq.sql = prev
	}
	if discordaccount.Policy == nil {
		return errors.New("ent: uninitialized discordaccount.Policy (forgotten import ent/runtime?)")
	}
	if err := discordaccount.Policy.EvalQuery(ctx, daq); err != nil {
		return err
	}
	return nil
}

func (daq *DiscordAccountQuery) sqlAll(ctx context.Context) ([]*DiscordAccount, error) {
	var (
		nodes       = []*DiscordAccount{}
		withFKs     = daq.withFKs
		_spec       = daq.querySpec()
		loadedTypes = [2]bool{
			daq.withOwner != nil,
			daq.withBot != nil,
		}
	)
	if daq.withOwner != nil || daq.withBot != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, discordaccount.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &DiscordAccount{config: daq.config}
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
	if err := sqlgraph.QueryNodes(ctx, daq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := daq.withOwner; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*DiscordAccount)
		for i := range nodes {
			if nodes[i].user_discord_accounts == nil {
				continue
			}
			fk := *nodes[i].user_discord_accounts
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
				return nil, fmt.Errorf(`unexpected foreign-key "user_discord_accounts" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Owner = n
			}
		}
	}

	if query := daq.withBot; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*DiscordAccount)
		for i := range nodes {
			if nodes[i].discord_bot_account == nil {
				continue
			}
			fk := *nodes[i].discord_bot_account
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(discordbot.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "discord_bot_account" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Bot = n
			}
		}
	}

	return nodes, nil
}

func (daq *DiscordAccountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := daq.querySpec()
	return sqlgraph.CountNodes(ctx, daq.driver, _spec)
}

func (daq *DiscordAccountQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := daq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (daq *DiscordAccountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   discordaccount.Table,
			Columns: discordaccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: discordaccount.FieldID,
			},
		},
		From:   daq.sql,
		Unique: true,
	}
	if unique := daq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := daq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, discordaccount.FieldID)
		for i := range fields {
			if fields[i] != discordaccount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := daq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := daq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := daq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := daq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (daq *DiscordAccountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(daq.driver.Dialect())
	t1 := builder.Table(discordaccount.Table)
	columns := daq.fields
	if len(columns) == 0 {
		columns = discordaccount.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if daq.sql != nil {
		selector = daq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range daq.predicates {
		p(selector)
	}
	for _, p := range daq.order {
		p(selector)
	}
	if offset := daq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := daq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DiscordAccountGroupBy is the group-by builder for DiscordAccount entities.
type DiscordAccountGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dagb *DiscordAccountGroupBy) Aggregate(fns ...AggregateFunc) *DiscordAccountGroupBy {
	dagb.fns = append(dagb.fns, fns...)
	return dagb
}

// Scan applies the group-by query and scans the result into the given value.
func (dagb *DiscordAccountGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dagb.path(ctx)
	if err != nil {
		return err
	}
	dagb.sql = query
	return dagb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dagb *DiscordAccountGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := dagb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (dagb *DiscordAccountGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(dagb.fields) > 1 {
		return nil, errors.New("ent: DiscordAccountGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := dagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dagb *DiscordAccountGroupBy) StringsX(ctx context.Context) []string {
	v, err := dagb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dagb *DiscordAccountGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dagb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discordaccount.Label}
	default:
		err = fmt.Errorf("ent: DiscordAccountGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dagb *DiscordAccountGroupBy) StringX(ctx context.Context) string {
	v, err := dagb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (dagb *DiscordAccountGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(dagb.fields) > 1 {
		return nil, errors.New("ent: DiscordAccountGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := dagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dagb *DiscordAccountGroupBy) IntsX(ctx context.Context) []int {
	v, err := dagb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dagb *DiscordAccountGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dagb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discordaccount.Label}
	default:
		err = fmt.Errorf("ent: DiscordAccountGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dagb *DiscordAccountGroupBy) IntX(ctx context.Context) int {
	v, err := dagb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (dagb *DiscordAccountGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(dagb.fields) > 1 {
		return nil, errors.New("ent: DiscordAccountGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := dagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dagb *DiscordAccountGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := dagb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dagb *DiscordAccountGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dagb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discordaccount.Label}
	default:
		err = fmt.Errorf("ent: DiscordAccountGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dagb *DiscordAccountGroupBy) Float64X(ctx context.Context) float64 {
	v, err := dagb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (dagb *DiscordAccountGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(dagb.fields) > 1 {
		return nil, errors.New("ent: DiscordAccountGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := dagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dagb *DiscordAccountGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := dagb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dagb *DiscordAccountGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dagb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discordaccount.Label}
	default:
		err = fmt.Errorf("ent: DiscordAccountGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dagb *DiscordAccountGroupBy) BoolX(ctx context.Context) bool {
	v, err := dagb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dagb *DiscordAccountGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dagb.fields {
		if !discordaccount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dagb *DiscordAccountGroupBy) sqlQuery() *sql.Selector {
	selector := dagb.sql.Select()
	aggregation := make([]string, 0, len(dagb.fns))
	for _, fn := range dagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dagb.fields)+len(dagb.fns))
		for _, f := range dagb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dagb.fields...)...)
}

// DiscordAccountSelect is the builder for selecting fields of DiscordAccount entities.
type DiscordAccountSelect struct {
	*DiscordAccountQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (das *DiscordAccountSelect) Scan(ctx context.Context, v interface{}) error {
	if err := das.prepareQuery(ctx); err != nil {
		return err
	}
	das.sql = das.DiscordAccountQuery.sqlQuery(ctx)
	return das.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (das *DiscordAccountSelect) ScanX(ctx context.Context, v interface{}) {
	if err := das.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (das *DiscordAccountSelect) Strings(ctx context.Context) ([]string, error) {
	if len(das.fields) > 1 {
		return nil, errors.New("ent: DiscordAccountSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := das.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (das *DiscordAccountSelect) StringsX(ctx context.Context) []string {
	v, err := das.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (das *DiscordAccountSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = das.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discordaccount.Label}
	default:
		err = fmt.Errorf("ent: DiscordAccountSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (das *DiscordAccountSelect) StringX(ctx context.Context) string {
	v, err := das.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (das *DiscordAccountSelect) Ints(ctx context.Context) ([]int, error) {
	if len(das.fields) > 1 {
		return nil, errors.New("ent: DiscordAccountSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := das.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (das *DiscordAccountSelect) IntsX(ctx context.Context) []int {
	v, err := das.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (das *DiscordAccountSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = das.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discordaccount.Label}
	default:
		err = fmt.Errorf("ent: DiscordAccountSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (das *DiscordAccountSelect) IntX(ctx context.Context) int {
	v, err := das.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (das *DiscordAccountSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(das.fields) > 1 {
		return nil, errors.New("ent: DiscordAccountSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := das.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (das *DiscordAccountSelect) Float64sX(ctx context.Context) []float64 {
	v, err := das.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (das *DiscordAccountSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = das.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discordaccount.Label}
	default:
		err = fmt.Errorf("ent: DiscordAccountSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (das *DiscordAccountSelect) Float64X(ctx context.Context) float64 {
	v, err := das.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (das *DiscordAccountSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(das.fields) > 1 {
		return nil, errors.New("ent: DiscordAccountSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := das.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (das *DiscordAccountSelect) BoolsX(ctx context.Context) []bool {
	v, err := das.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (das *DiscordAccountSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = das.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discordaccount.Label}
	default:
		err = fmt.Errorf("ent: DiscordAccountSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (das *DiscordAccountSelect) BoolX(ctx context.Context) bool {
	v, err := das.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (das *DiscordAccountSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := das.sql.Query()
	if err := das.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
