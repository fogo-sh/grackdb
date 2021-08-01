// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/fogo-sh/grackdb/ent/migrate"

	"github.com/fogo-sh/grackdb/ent/discordaccount"
	"github.com/fogo-sh/grackdb/ent/githubaccount"
	"github.com/fogo-sh/grackdb/ent/githuborganization"
	"github.com/fogo-sh/grackdb/ent/githuborganizationmember"
	"github.com/fogo-sh/grackdb/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// DiscordAccount is the client for interacting with the DiscordAccount builders.
	DiscordAccount *DiscordAccountClient
	// GithubAccount is the client for interacting with the GithubAccount builders.
	GithubAccount *GithubAccountClient
	// GithubOrganization is the client for interacting with the GithubOrganization builders.
	GithubOrganization *GithubOrganizationClient
	// GithubOrganizationMember is the client for interacting with the GithubOrganizationMember builders.
	GithubOrganizationMember *GithubOrganizationMemberClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// additional fields for node api
	tables tables
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.DiscordAccount = NewDiscordAccountClient(c.config)
	c.GithubAccount = NewGithubAccountClient(c.config)
	c.GithubOrganization = NewGithubOrganizationClient(c.config)
	c.GithubOrganizationMember = NewGithubOrganizationMemberClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:                      ctx,
		config:                   cfg,
		DiscordAccount:           NewDiscordAccountClient(cfg),
		GithubAccount:            NewGithubAccountClient(cfg),
		GithubOrganization:       NewGithubOrganizationClient(cfg),
		GithubOrganizationMember: NewGithubOrganizationMemberClient(cfg),
		User:                     NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:                   cfg,
		DiscordAccount:           NewDiscordAccountClient(cfg),
		GithubAccount:            NewGithubAccountClient(cfg),
		GithubOrganization:       NewGithubOrganizationClient(cfg),
		GithubOrganizationMember: NewGithubOrganizationMemberClient(cfg),
		User:                     NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		DiscordAccount.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.DiscordAccount.Use(hooks...)
	c.GithubAccount.Use(hooks...)
	c.GithubOrganization.Use(hooks...)
	c.GithubOrganizationMember.Use(hooks...)
	c.User.Use(hooks...)
}

// DiscordAccountClient is a client for the DiscordAccount schema.
type DiscordAccountClient struct {
	config
}

// NewDiscordAccountClient returns a client for the DiscordAccount from the given config.
func NewDiscordAccountClient(c config) *DiscordAccountClient {
	return &DiscordAccountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `discordaccount.Hooks(f(g(h())))`.
func (c *DiscordAccountClient) Use(hooks ...Hook) {
	c.hooks.DiscordAccount = append(c.hooks.DiscordAccount, hooks...)
}

// Create returns a create builder for DiscordAccount.
func (c *DiscordAccountClient) Create() *DiscordAccountCreate {
	mutation := newDiscordAccountMutation(c.config, OpCreate)
	return &DiscordAccountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of DiscordAccount entities.
func (c *DiscordAccountClient) CreateBulk(builders ...*DiscordAccountCreate) *DiscordAccountCreateBulk {
	return &DiscordAccountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for DiscordAccount.
func (c *DiscordAccountClient) Update() *DiscordAccountUpdate {
	mutation := newDiscordAccountMutation(c.config, OpUpdate)
	return &DiscordAccountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DiscordAccountClient) UpdateOne(da *DiscordAccount) *DiscordAccountUpdateOne {
	mutation := newDiscordAccountMutation(c.config, OpUpdateOne, withDiscordAccount(da))
	return &DiscordAccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DiscordAccountClient) UpdateOneID(id int) *DiscordAccountUpdateOne {
	mutation := newDiscordAccountMutation(c.config, OpUpdateOne, withDiscordAccountID(id))
	return &DiscordAccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for DiscordAccount.
func (c *DiscordAccountClient) Delete() *DiscordAccountDelete {
	mutation := newDiscordAccountMutation(c.config, OpDelete)
	return &DiscordAccountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DiscordAccountClient) DeleteOne(da *DiscordAccount) *DiscordAccountDeleteOne {
	return c.DeleteOneID(da.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DiscordAccountClient) DeleteOneID(id int) *DiscordAccountDeleteOne {
	builder := c.Delete().Where(discordaccount.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DiscordAccountDeleteOne{builder}
}

// Query returns a query builder for DiscordAccount.
func (c *DiscordAccountClient) Query() *DiscordAccountQuery {
	return &DiscordAccountQuery{
		config: c.config,
	}
}

// Get returns a DiscordAccount entity by its id.
func (c *DiscordAccountClient) Get(ctx context.Context, id int) (*DiscordAccount, error) {
	return c.Query().Where(discordaccount.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DiscordAccountClient) GetX(ctx context.Context, id int) *DiscordAccount {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a DiscordAccount.
func (c *DiscordAccountClient) QueryOwner(da *DiscordAccount) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := da.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(discordaccount.Table, discordaccount.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, discordaccount.OwnerTable, discordaccount.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(da.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DiscordAccountClient) Hooks() []Hook {
	hooks := c.hooks.DiscordAccount
	return append(hooks[:len(hooks):len(hooks)], discordaccount.Hooks[:]...)
}

// GithubAccountClient is a client for the GithubAccount schema.
type GithubAccountClient struct {
	config
}

// NewGithubAccountClient returns a client for the GithubAccount from the given config.
func NewGithubAccountClient(c config) *GithubAccountClient {
	return &GithubAccountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `githubaccount.Hooks(f(g(h())))`.
func (c *GithubAccountClient) Use(hooks ...Hook) {
	c.hooks.GithubAccount = append(c.hooks.GithubAccount, hooks...)
}

// Create returns a create builder for GithubAccount.
func (c *GithubAccountClient) Create() *GithubAccountCreate {
	mutation := newGithubAccountMutation(c.config, OpCreate)
	return &GithubAccountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of GithubAccount entities.
func (c *GithubAccountClient) CreateBulk(builders ...*GithubAccountCreate) *GithubAccountCreateBulk {
	return &GithubAccountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for GithubAccount.
func (c *GithubAccountClient) Update() *GithubAccountUpdate {
	mutation := newGithubAccountMutation(c.config, OpUpdate)
	return &GithubAccountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GithubAccountClient) UpdateOne(ga *GithubAccount) *GithubAccountUpdateOne {
	mutation := newGithubAccountMutation(c.config, OpUpdateOne, withGithubAccount(ga))
	return &GithubAccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GithubAccountClient) UpdateOneID(id int) *GithubAccountUpdateOne {
	mutation := newGithubAccountMutation(c.config, OpUpdateOne, withGithubAccountID(id))
	return &GithubAccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for GithubAccount.
func (c *GithubAccountClient) Delete() *GithubAccountDelete {
	mutation := newGithubAccountMutation(c.config, OpDelete)
	return &GithubAccountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *GithubAccountClient) DeleteOne(ga *GithubAccount) *GithubAccountDeleteOne {
	return c.DeleteOneID(ga.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *GithubAccountClient) DeleteOneID(id int) *GithubAccountDeleteOne {
	builder := c.Delete().Where(githubaccount.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GithubAccountDeleteOne{builder}
}

// Query returns a query builder for GithubAccount.
func (c *GithubAccountClient) Query() *GithubAccountQuery {
	return &GithubAccountQuery{
		config: c.config,
	}
}

// Get returns a GithubAccount entity by its id.
func (c *GithubAccountClient) Get(ctx context.Context, id int) (*GithubAccount, error) {
	return c.Query().Where(githubaccount.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GithubAccountClient) GetX(ctx context.Context, id int) *GithubAccount {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a GithubAccount.
func (c *GithubAccountClient) QueryOwner(ga *GithubAccount) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ga.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(githubaccount.Table, githubaccount.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, githubaccount.OwnerTable, githubaccount.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(ga.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOrganizationMemberships queries the organization_memberships edge of a GithubAccount.
func (c *GithubAccountClient) QueryOrganizationMemberships(ga *GithubAccount) *GithubOrganizationMemberQuery {
	query := &GithubOrganizationMemberQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ga.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(githubaccount.Table, githubaccount.FieldID, id),
			sqlgraph.To(githuborganizationmember.Table, githuborganizationmember.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, githubaccount.OrganizationMembershipsTable, githubaccount.OrganizationMembershipsColumn),
		)
		fromV = sqlgraph.Neighbors(ga.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GithubAccountClient) Hooks() []Hook {
	hooks := c.hooks.GithubAccount
	return append(hooks[:len(hooks):len(hooks)], githubaccount.Hooks[:]...)
}

// GithubOrganizationClient is a client for the GithubOrganization schema.
type GithubOrganizationClient struct {
	config
}

// NewGithubOrganizationClient returns a client for the GithubOrganization from the given config.
func NewGithubOrganizationClient(c config) *GithubOrganizationClient {
	return &GithubOrganizationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `githuborganization.Hooks(f(g(h())))`.
func (c *GithubOrganizationClient) Use(hooks ...Hook) {
	c.hooks.GithubOrganization = append(c.hooks.GithubOrganization, hooks...)
}

// Create returns a create builder for GithubOrganization.
func (c *GithubOrganizationClient) Create() *GithubOrganizationCreate {
	mutation := newGithubOrganizationMutation(c.config, OpCreate)
	return &GithubOrganizationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of GithubOrganization entities.
func (c *GithubOrganizationClient) CreateBulk(builders ...*GithubOrganizationCreate) *GithubOrganizationCreateBulk {
	return &GithubOrganizationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for GithubOrganization.
func (c *GithubOrganizationClient) Update() *GithubOrganizationUpdate {
	mutation := newGithubOrganizationMutation(c.config, OpUpdate)
	return &GithubOrganizationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GithubOrganizationClient) UpdateOne(_go *GithubOrganization) *GithubOrganizationUpdateOne {
	mutation := newGithubOrganizationMutation(c.config, OpUpdateOne, withGithubOrganization(_go))
	return &GithubOrganizationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GithubOrganizationClient) UpdateOneID(id int) *GithubOrganizationUpdateOne {
	mutation := newGithubOrganizationMutation(c.config, OpUpdateOne, withGithubOrganizationID(id))
	return &GithubOrganizationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for GithubOrganization.
func (c *GithubOrganizationClient) Delete() *GithubOrganizationDelete {
	mutation := newGithubOrganizationMutation(c.config, OpDelete)
	return &GithubOrganizationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *GithubOrganizationClient) DeleteOne(_go *GithubOrganization) *GithubOrganizationDeleteOne {
	return c.DeleteOneID(_go.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *GithubOrganizationClient) DeleteOneID(id int) *GithubOrganizationDeleteOne {
	builder := c.Delete().Where(githuborganization.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GithubOrganizationDeleteOne{builder}
}

// Query returns a query builder for GithubOrganization.
func (c *GithubOrganizationClient) Query() *GithubOrganizationQuery {
	return &GithubOrganizationQuery{
		config: c.config,
	}
}

// Get returns a GithubOrganization entity by its id.
func (c *GithubOrganizationClient) Get(ctx context.Context, id int) (*GithubOrganization, error) {
	return c.Query().Where(githuborganization.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GithubOrganizationClient) GetX(ctx context.Context, id int) *GithubOrganization {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryMembers queries the members edge of a GithubOrganization.
func (c *GithubOrganizationClient) QueryMembers(_go *GithubOrganization) *GithubOrganizationMemberQuery {
	query := &GithubOrganizationMemberQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := _go.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(githuborganization.Table, githuborganization.FieldID, id),
			sqlgraph.To(githuborganizationmember.Table, githuborganizationmember.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, githuborganization.MembersTable, githuborganization.MembersColumn),
		)
		fromV = sqlgraph.Neighbors(_go.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GithubOrganizationClient) Hooks() []Hook {
	hooks := c.hooks.GithubOrganization
	return append(hooks[:len(hooks):len(hooks)], githuborganization.Hooks[:]...)
}

// GithubOrganizationMemberClient is a client for the GithubOrganizationMember schema.
type GithubOrganizationMemberClient struct {
	config
}

// NewGithubOrganizationMemberClient returns a client for the GithubOrganizationMember from the given config.
func NewGithubOrganizationMemberClient(c config) *GithubOrganizationMemberClient {
	return &GithubOrganizationMemberClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `githuborganizationmember.Hooks(f(g(h())))`.
func (c *GithubOrganizationMemberClient) Use(hooks ...Hook) {
	c.hooks.GithubOrganizationMember = append(c.hooks.GithubOrganizationMember, hooks...)
}

// Create returns a create builder for GithubOrganizationMember.
func (c *GithubOrganizationMemberClient) Create() *GithubOrganizationMemberCreate {
	mutation := newGithubOrganizationMemberMutation(c.config, OpCreate)
	return &GithubOrganizationMemberCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of GithubOrganizationMember entities.
func (c *GithubOrganizationMemberClient) CreateBulk(builders ...*GithubOrganizationMemberCreate) *GithubOrganizationMemberCreateBulk {
	return &GithubOrganizationMemberCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for GithubOrganizationMember.
func (c *GithubOrganizationMemberClient) Update() *GithubOrganizationMemberUpdate {
	mutation := newGithubOrganizationMemberMutation(c.config, OpUpdate)
	return &GithubOrganizationMemberUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GithubOrganizationMemberClient) UpdateOne(gom *GithubOrganizationMember) *GithubOrganizationMemberUpdateOne {
	mutation := newGithubOrganizationMemberMutation(c.config, OpUpdateOne, withGithubOrganizationMember(gom))
	return &GithubOrganizationMemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GithubOrganizationMemberClient) UpdateOneID(id int) *GithubOrganizationMemberUpdateOne {
	mutation := newGithubOrganizationMemberMutation(c.config, OpUpdateOne, withGithubOrganizationMemberID(id))
	return &GithubOrganizationMemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for GithubOrganizationMember.
func (c *GithubOrganizationMemberClient) Delete() *GithubOrganizationMemberDelete {
	mutation := newGithubOrganizationMemberMutation(c.config, OpDelete)
	return &GithubOrganizationMemberDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *GithubOrganizationMemberClient) DeleteOne(gom *GithubOrganizationMember) *GithubOrganizationMemberDeleteOne {
	return c.DeleteOneID(gom.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *GithubOrganizationMemberClient) DeleteOneID(id int) *GithubOrganizationMemberDeleteOne {
	builder := c.Delete().Where(githuborganizationmember.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GithubOrganizationMemberDeleteOne{builder}
}

// Query returns a query builder for GithubOrganizationMember.
func (c *GithubOrganizationMemberClient) Query() *GithubOrganizationMemberQuery {
	return &GithubOrganizationMemberQuery{
		config: c.config,
	}
}

// Get returns a GithubOrganizationMember entity by its id.
func (c *GithubOrganizationMemberClient) Get(ctx context.Context, id int) (*GithubOrganizationMember, error) {
	return c.Query().Where(githuborganizationmember.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GithubOrganizationMemberClient) GetX(ctx context.Context, id int) *GithubOrganizationMember {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOrganization queries the organization edge of a GithubOrganizationMember.
func (c *GithubOrganizationMemberClient) QueryOrganization(gom *GithubOrganizationMember) *GithubOrganizationQuery {
	query := &GithubOrganizationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gom.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(githuborganizationmember.Table, githuborganizationmember.FieldID, id),
			sqlgraph.To(githuborganization.Table, githuborganization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, githuborganizationmember.OrganizationTable, githuborganizationmember.OrganizationColumn),
		)
		fromV = sqlgraph.Neighbors(gom.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAccount queries the account edge of a GithubOrganizationMember.
func (c *GithubOrganizationMemberClient) QueryAccount(gom *GithubOrganizationMember) *GithubAccountQuery {
	query := &GithubAccountQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gom.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(githuborganizationmember.Table, githuborganizationmember.FieldID, id),
			sqlgraph.To(githubaccount.Table, githubaccount.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, githuborganizationmember.AccountTable, githuborganizationmember.AccountColumn),
		)
		fromV = sqlgraph.Neighbors(gom.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GithubOrganizationMemberClient) Hooks() []Hook {
	hooks := c.hooks.GithubOrganizationMember
	return append(hooks[:len(hooks):len(hooks)], githuborganizationmember.Hooks[:]...)
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDiscordAccounts queries the discord_accounts edge of a User.
func (c *UserClient) QueryDiscordAccounts(u *User) *DiscordAccountQuery {
	query := &DiscordAccountQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(discordaccount.Table, discordaccount.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.DiscordAccountsTable, user.DiscordAccountsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGithubAccounts queries the github_accounts edge of a User.
func (c *UserClient) QueryGithubAccounts(u *User) *GithubAccountQuery {
	query := &GithubAccountQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(githubaccount.Table, githubaccount.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.GithubAccountsTable, user.GithubAccountsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	hooks := c.hooks.User
	return append(hooks[:len(hooks):len(hooks)], user.Hooks[:]...)
}
