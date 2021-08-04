// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/discordbot"
	"github.com/fogo-sh/grackdb/ent/githubaccount"
	"github.com/fogo-sh/grackdb/ent/githuborganization"
	"github.com/fogo-sh/grackdb/ent/predicate"
	"github.com/fogo-sh/grackdb/ent/project"
	"github.com/fogo-sh/grackdb/ent/repository"
	"github.com/fogo-sh/grackdb/ent/site"
)

// RepositoryUpdate is the builder for updating Repository entities.
type RepositoryUpdate struct {
	config
	hooks    []Hook
	mutation *RepositoryMutation
}

// Where adds a new predicate for the RepositoryUpdate builder.
func (ru *RepositoryUpdate) Where(ps ...predicate.Repository) *RepositoryUpdate {
	ru.mutation.predicates = append(ru.mutation.predicates, ps...)
	return ru
}

// SetName sets the "name" field.
func (ru *RepositoryUpdate) SetName(s string) *RepositoryUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetDescription sets the "description" field.
func (ru *RepositoryUpdate) SetDescription(s string) *RepositoryUpdate {
	ru.mutation.SetDescription(s)
	return ru
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ru *RepositoryUpdate) SetNillableDescription(s *string) *RepositoryUpdate {
	if s != nil {
		ru.SetDescription(*s)
	}
	return ru
}

// ClearDescription clears the value of the "description" field.
func (ru *RepositoryUpdate) ClearDescription() *RepositoryUpdate {
	ru.mutation.ClearDescription()
	return ru
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (ru *RepositoryUpdate) SetProjectID(id int) *RepositoryUpdate {
	ru.mutation.SetProjectID(id)
	return ru
}

// SetProject sets the "project" edge to the Project entity.
func (ru *RepositoryUpdate) SetProject(p *Project) *RepositoryUpdate {
	return ru.SetProjectID(p.ID)
}

// SetGithubAccountID sets the "github_account" edge to the GithubAccount entity by ID.
func (ru *RepositoryUpdate) SetGithubAccountID(id int) *RepositoryUpdate {
	ru.mutation.SetGithubAccountID(id)
	return ru
}

// SetNillableGithubAccountID sets the "github_account" edge to the GithubAccount entity by ID if the given value is not nil.
func (ru *RepositoryUpdate) SetNillableGithubAccountID(id *int) *RepositoryUpdate {
	if id != nil {
		ru = ru.SetGithubAccountID(*id)
	}
	return ru
}

// SetGithubAccount sets the "github_account" edge to the GithubAccount entity.
func (ru *RepositoryUpdate) SetGithubAccount(g *GithubAccount) *RepositoryUpdate {
	return ru.SetGithubAccountID(g.ID)
}

// SetGithubOrganizationID sets the "github_organization" edge to the GithubOrganization entity by ID.
func (ru *RepositoryUpdate) SetGithubOrganizationID(id int) *RepositoryUpdate {
	ru.mutation.SetGithubOrganizationID(id)
	return ru
}

// SetNillableGithubOrganizationID sets the "github_organization" edge to the GithubOrganization entity by ID if the given value is not nil.
func (ru *RepositoryUpdate) SetNillableGithubOrganizationID(id *int) *RepositoryUpdate {
	if id != nil {
		ru = ru.SetGithubOrganizationID(*id)
	}
	return ru
}

// SetGithubOrganization sets the "github_organization" edge to the GithubOrganization entity.
func (ru *RepositoryUpdate) SetGithubOrganization(g *GithubOrganization) *RepositoryUpdate {
	return ru.SetGithubOrganizationID(g.ID)
}

// AddDiscordBotIDs adds the "discord_bots" edge to the DiscordBot entity by IDs.
func (ru *RepositoryUpdate) AddDiscordBotIDs(ids ...int) *RepositoryUpdate {
	ru.mutation.AddDiscordBotIDs(ids...)
	return ru
}

// AddDiscordBots adds the "discord_bots" edges to the DiscordBot entity.
func (ru *RepositoryUpdate) AddDiscordBots(d ...*DiscordBot) *RepositoryUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ru.AddDiscordBotIDs(ids...)
}

// AddSiteIDs adds the "sites" edge to the Site entity by IDs.
func (ru *RepositoryUpdate) AddSiteIDs(ids ...int) *RepositoryUpdate {
	ru.mutation.AddSiteIDs(ids...)
	return ru
}

// AddSites adds the "sites" edges to the Site entity.
func (ru *RepositoryUpdate) AddSites(s ...*Site) *RepositoryUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ru.AddSiteIDs(ids...)
}

// Mutation returns the RepositoryMutation object of the builder.
func (ru *RepositoryUpdate) Mutation() *RepositoryMutation {
	return ru.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (ru *RepositoryUpdate) ClearProject() *RepositoryUpdate {
	ru.mutation.ClearProject()
	return ru
}

// ClearGithubAccount clears the "github_account" edge to the GithubAccount entity.
func (ru *RepositoryUpdate) ClearGithubAccount() *RepositoryUpdate {
	ru.mutation.ClearGithubAccount()
	return ru
}

// ClearGithubOrganization clears the "github_organization" edge to the GithubOrganization entity.
func (ru *RepositoryUpdate) ClearGithubOrganization() *RepositoryUpdate {
	ru.mutation.ClearGithubOrganization()
	return ru
}

// ClearDiscordBots clears all "discord_bots" edges to the DiscordBot entity.
func (ru *RepositoryUpdate) ClearDiscordBots() *RepositoryUpdate {
	ru.mutation.ClearDiscordBots()
	return ru
}

// RemoveDiscordBotIDs removes the "discord_bots" edge to DiscordBot entities by IDs.
func (ru *RepositoryUpdate) RemoveDiscordBotIDs(ids ...int) *RepositoryUpdate {
	ru.mutation.RemoveDiscordBotIDs(ids...)
	return ru
}

// RemoveDiscordBots removes "discord_bots" edges to DiscordBot entities.
func (ru *RepositoryUpdate) RemoveDiscordBots(d ...*DiscordBot) *RepositoryUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ru.RemoveDiscordBotIDs(ids...)
}

// ClearSites clears all "sites" edges to the Site entity.
func (ru *RepositoryUpdate) ClearSites() *RepositoryUpdate {
	ru.mutation.ClearSites()
	return ru
}

// RemoveSiteIDs removes the "sites" edge to Site entities by IDs.
func (ru *RepositoryUpdate) RemoveSiteIDs(ids ...int) *RepositoryUpdate {
	ru.mutation.RemoveSiteIDs(ids...)
	return ru
}

// RemoveSites removes "sites" edges to Site entities.
func (ru *RepositoryUpdate) RemoveSites(s ...*Site) *RepositoryUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ru.RemoveSiteIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RepositoryUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RepositoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RepositoryUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RepositoryUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RepositoryUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RepositoryUpdate) check() error {
	if v, ok := ru.mutation.Name(); ok {
		if err := repository.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := ru.mutation.ProjectID(); ru.mutation.ProjectCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"project\"")
	}
	return nil
}

func (ru *RepositoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   repository.Table,
			Columns: repository.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: repository.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repository.FieldName,
		})
	}
	if value, ok := ru.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repository.FieldDescription,
		})
	}
	if ru.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: repository.FieldDescription,
		})
	}
	if ru.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repository.ProjectTable,
			Columns: []string{repository.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repository.ProjectTable,
			Columns: []string{repository.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.GithubAccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repository.GithubAccountTable,
			Columns: []string{repository.GithubAccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: githubaccount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.GithubAccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repository.GithubAccountTable,
			Columns: []string{repository.GithubAccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: githubaccount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.GithubOrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repository.GithubOrganizationTable,
			Columns: []string{repository.GithubOrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: githuborganization.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.GithubOrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repository.GithubOrganizationTable,
			Columns: []string{repository.GithubOrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: githuborganization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.DiscordBotsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.DiscordBotsTable,
			Columns: []string{repository.DiscordBotsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: discordbot.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedDiscordBotsIDs(); len(nodes) > 0 && !ru.mutation.DiscordBotsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.DiscordBotsTable,
			Columns: []string{repository.DiscordBotsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: discordbot.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.DiscordBotsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.DiscordBotsTable,
			Columns: []string{repository.DiscordBotsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: discordbot.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.SitesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.SitesTable,
			Columns: []string{repository.SitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: site.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedSitesIDs(); len(nodes) > 0 && !ru.mutation.SitesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.SitesTable,
			Columns: []string{repository.SitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: site.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.SitesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.SitesTable,
			Columns: []string{repository.SitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: site.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{repository.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RepositoryUpdateOne is the builder for updating a single Repository entity.
type RepositoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RepositoryMutation
}

// SetName sets the "name" field.
func (ruo *RepositoryUpdateOne) SetName(s string) *RepositoryUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetDescription sets the "description" field.
func (ruo *RepositoryUpdateOne) SetDescription(s string) *RepositoryUpdateOne {
	ruo.mutation.SetDescription(s)
	return ruo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ruo *RepositoryUpdateOne) SetNillableDescription(s *string) *RepositoryUpdateOne {
	if s != nil {
		ruo.SetDescription(*s)
	}
	return ruo
}

// ClearDescription clears the value of the "description" field.
func (ruo *RepositoryUpdateOne) ClearDescription() *RepositoryUpdateOne {
	ruo.mutation.ClearDescription()
	return ruo
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (ruo *RepositoryUpdateOne) SetProjectID(id int) *RepositoryUpdateOne {
	ruo.mutation.SetProjectID(id)
	return ruo
}

// SetProject sets the "project" edge to the Project entity.
func (ruo *RepositoryUpdateOne) SetProject(p *Project) *RepositoryUpdateOne {
	return ruo.SetProjectID(p.ID)
}

// SetGithubAccountID sets the "github_account" edge to the GithubAccount entity by ID.
func (ruo *RepositoryUpdateOne) SetGithubAccountID(id int) *RepositoryUpdateOne {
	ruo.mutation.SetGithubAccountID(id)
	return ruo
}

// SetNillableGithubAccountID sets the "github_account" edge to the GithubAccount entity by ID if the given value is not nil.
func (ruo *RepositoryUpdateOne) SetNillableGithubAccountID(id *int) *RepositoryUpdateOne {
	if id != nil {
		ruo = ruo.SetGithubAccountID(*id)
	}
	return ruo
}

// SetGithubAccount sets the "github_account" edge to the GithubAccount entity.
func (ruo *RepositoryUpdateOne) SetGithubAccount(g *GithubAccount) *RepositoryUpdateOne {
	return ruo.SetGithubAccountID(g.ID)
}

// SetGithubOrganizationID sets the "github_organization" edge to the GithubOrganization entity by ID.
func (ruo *RepositoryUpdateOne) SetGithubOrganizationID(id int) *RepositoryUpdateOne {
	ruo.mutation.SetGithubOrganizationID(id)
	return ruo
}

// SetNillableGithubOrganizationID sets the "github_organization" edge to the GithubOrganization entity by ID if the given value is not nil.
func (ruo *RepositoryUpdateOne) SetNillableGithubOrganizationID(id *int) *RepositoryUpdateOne {
	if id != nil {
		ruo = ruo.SetGithubOrganizationID(*id)
	}
	return ruo
}

// SetGithubOrganization sets the "github_organization" edge to the GithubOrganization entity.
func (ruo *RepositoryUpdateOne) SetGithubOrganization(g *GithubOrganization) *RepositoryUpdateOne {
	return ruo.SetGithubOrganizationID(g.ID)
}

// AddDiscordBotIDs adds the "discord_bots" edge to the DiscordBot entity by IDs.
func (ruo *RepositoryUpdateOne) AddDiscordBotIDs(ids ...int) *RepositoryUpdateOne {
	ruo.mutation.AddDiscordBotIDs(ids...)
	return ruo
}

// AddDiscordBots adds the "discord_bots" edges to the DiscordBot entity.
func (ruo *RepositoryUpdateOne) AddDiscordBots(d ...*DiscordBot) *RepositoryUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ruo.AddDiscordBotIDs(ids...)
}

// AddSiteIDs adds the "sites" edge to the Site entity by IDs.
func (ruo *RepositoryUpdateOne) AddSiteIDs(ids ...int) *RepositoryUpdateOne {
	ruo.mutation.AddSiteIDs(ids...)
	return ruo
}

// AddSites adds the "sites" edges to the Site entity.
func (ruo *RepositoryUpdateOne) AddSites(s ...*Site) *RepositoryUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ruo.AddSiteIDs(ids...)
}

// Mutation returns the RepositoryMutation object of the builder.
func (ruo *RepositoryUpdateOne) Mutation() *RepositoryMutation {
	return ruo.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (ruo *RepositoryUpdateOne) ClearProject() *RepositoryUpdateOne {
	ruo.mutation.ClearProject()
	return ruo
}

// ClearGithubAccount clears the "github_account" edge to the GithubAccount entity.
func (ruo *RepositoryUpdateOne) ClearGithubAccount() *RepositoryUpdateOne {
	ruo.mutation.ClearGithubAccount()
	return ruo
}

// ClearGithubOrganization clears the "github_organization" edge to the GithubOrganization entity.
func (ruo *RepositoryUpdateOne) ClearGithubOrganization() *RepositoryUpdateOne {
	ruo.mutation.ClearGithubOrganization()
	return ruo
}

// ClearDiscordBots clears all "discord_bots" edges to the DiscordBot entity.
func (ruo *RepositoryUpdateOne) ClearDiscordBots() *RepositoryUpdateOne {
	ruo.mutation.ClearDiscordBots()
	return ruo
}

// RemoveDiscordBotIDs removes the "discord_bots" edge to DiscordBot entities by IDs.
func (ruo *RepositoryUpdateOne) RemoveDiscordBotIDs(ids ...int) *RepositoryUpdateOne {
	ruo.mutation.RemoveDiscordBotIDs(ids...)
	return ruo
}

// RemoveDiscordBots removes "discord_bots" edges to DiscordBot entities.
func (ruo *RepositoryUpdateOne) RemoveDiscordBots(d ...*DiscordBot) *RepositoryUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ruo.RemoveDiscordBotIDs(ids...)
}

// ClearSites clears all "sites" edges to the Site entity.
func (ruo *RepositoryUpdateOne) ClearSites() *RepositoryUpdateOne {
	ruo.mutation.ClearSites()
	return ruo
}

// RemoveSiteIDs removes the "sites" edge to Site entities by IDs.
func (ruo *RepositoryUpdateOne) RemoveSiteIDs(ids ...int) *RepositoryUpdateOne {
	ruo.mutation.RemoveSiteIDs(ids...)
	return ruo
}

// RemoveSites removes "sites" edges to Site entities.
func (ruo *RepositoryUpdateOne) RemoveSites(s ...*Site) *RepositoryUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ruo.RemoveSiteIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RepositoryUpdateOne) Select(field string, fields ...string) *RepositoryUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Repository entity.
func (ruo *RepositoryUpdateOne) Save(ctx context.Context) (*Repository, error) {
	var (
		err  error
		node *Repository
	)
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RepositoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RepositoryUpdateOne) SaveX(ctx context.Context) *Repository {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RepositoryUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RepositoryUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RepositoryUpdateOne) check() error {
	if v, ok := ruo.mutation.Name(); ok {
		if err := repository.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := ruo.mutation.ProjectID(); ruo.mutation.ProjectCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"project\"")
	}
	return nil
}

func (ruo *RepositoryUpdateOne) sqlSave(ctx context.Context) (_node *Repository, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   repository.Table,
			Columns: repository.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: repository.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Repository.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, repository.FieldID)
		for _, f := range fields {
			if !repository.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != repository.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repository.FieldName,
		})
	}
	if value, ok := ruo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repository.FieldDescription,
		})
	}
	if ruo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: repository.FieldDescription,
		})
	}
	if ruo.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repository.ProjectTable,
			Columns: []string{repository.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repository.ProjectTable,
			Columns: []string{repository.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.GithubAccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repository.GithubAccountTable,
			Columns: []string{repository.GithubAccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: githubaccount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.GithubAccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repository.GithubAccountTable,
			Columns: []string{repository.GithubAccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: githubaccount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.GithubOrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repository.GithubOrganizationTable,
			Columns: []string{repository.GithubOrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: githuborganization.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.GithubOrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   repository.GithubOrganizationTable,
			Columns: []string{repository.GithubOrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: githuborganization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.DiscordBotsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.DiscordBotsTable,
			Columns: []string{repository.DiscordBotsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: discordbot.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedDiscordBotsIDs(); len(nodes) > 0 && !ruo.mutation.DiscordBotsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.DiscordBotsTable,
			Columns: []string{repository.DiscordBotsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: discordbot.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.DiscordBotsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.DiscordBotsTable,
			Columns: []string{repository.DiscordBotsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: discordbot.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.SitesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.SitesTable,
			Columns: []string{repository.SitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: site.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedSitesIDs(); len(nodes) > 0 && !ruo.mutation.SitesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.SitesTable,
			Columns: []string{repository.SitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: site.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.SitesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   repository.SitesTable,
			Columns: []string{repository.SitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: site.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Repository{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{repository.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}