// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/discordaccount"
	"github.com/fogo-sh/grackdb/ent/discordbot"
	"github.com/fogo-sh/grackdb/ent/predicate"
	"github.com/fogo-sh/grackdb/ent/project"
	"github.com/fogo-sh/grackdb/ent/repository"
)

// DiscordBotUpdate is the builder for updating DiscordBot entities.
type DiscordBotUpdate struct {
	config
	hooks    []Hook
	mutation *DiscordBotMutation
}

// Where adds a new predicate for the DiscordBotUpdate builder.
func (dbu *DiscordBotUpdate) Where(ps ...predicate.DiscordBot) *DiscordBotUpdate {
	dbu.mutation.predicates = append(dbu.mutation.predicates, ps...)
	return dbu
}

// SetAccountID sets the "account" edge to the DiscordAccount entity by ID.
func (dbu *DiscordBotUpdate) SetAccountID(id int) *DiscordBotUpdate {
	dbu.mutation.SetAccountID(id)
	return dbu
}

// SetAccount sets the "account" edge to the DiscordAccount entity.
func (dbu *DiscordBotUpdate) SetAccount(d *DiscordAccount) *DiscordBotUpdate {
	return dbu.SetAccountID(d.ID)
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (dbu *DiscordBotUpdate) SetProjectID(id int) *DiscordBotUpdate {
	dbu.mutation.SetProjectID(id)
	return dbu
}

// SetProject sets the "project" edge to the Project entity.
func (dbu *DiscordBotUpdate) SetProject(p *Project) *DiscordBotUpdate {
	return dbu.SetProjectID(p.ID)
}

// SetRepositoryID sets the "repository" edge to the Repository entity by ID.
func (dbu *DiscordBotUpdate) SetRepositoryID(id int) *DiscordBotUpdate {
	dbu.mutation.SetRepositoryID(id)
	return dbu
}

// SetNillableRepositoryID sets the "repository" edge to the Repository entity by ID if the given value is not nil.
func (dbu *DiscordBotUpdate) SetNillableRepositoryID(id *int) *DiscordBotUpdate {
	if id != nil {
		dbu = dbu.SetRepositoryID(*id)
	}
	return dbu
}

// SetRepository sets the "repository" edge to the Repository entity.
func (dbu *DiscordBotUpdate) SetRepository(r *Repository) *DiscordBotUpdate {
	return dbu.SetRepositoryID(r.ID)
}

// Mutation returns the DiscordBotMutation object of the builder.
func (dbu *DiscordBotUpdate) Mutation() *DiscordBotMutation {
	return dbu.mutation
}

// ClearAccount clears the "account" edge to the DiscordAccount entity.
func (dbu *DiscordBotUpdate) ClearAccount() *DiscordBotUpdate {
	dbu.mutation.ClearAccount()
	return dbu
}

// ClearProject clears the "project" edge to the Project entity.
func (dbu *DiscordBotUpdate) ClearProject() *DiscordBotUpdate {
	dbu.mutation.ClearProject()
	return dbu
}

// ClearRepository clears the "repository" edge to the Repository entity.
func (dbu *DiscordBotUpdate) ClearRepository() *DiscordBotUpdate {
	dbu.mutation.ClearRepository()
	return dbu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dbu *DiscordBotUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dbu.hooks) == 0 {
		if err = dbu.check(); err != nil {
			return 0, err
		}
		affected, err = dbu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiscordBotMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dbu.check(); err != nil {
				return 0, err
			}
			dbu.mutation = mutation
			affected, err = dbu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dbu.hooks) - 1; i >= 0; i-- {
			mut = dbu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dbu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (dbu *DiscordBotUpdate) SaveX(ctx context.Context) int {
	affected, err := dbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dbu *DiscordBotUpdate) Exec(ctx context.Context) error {
	_, err := dbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dbu *DiscordBotUpdate) ExecX(ctx context.Context) {
	if err := dbu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dbu *DiscordBotUpdate) check() error {
	if _, ok := dbu.mutation.AccountID(); dbu.mutation.AccountCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"account\"")
	}
	if _, ok := dbu.mutation.ProjectID(); dbu.mutation.ProjectCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"project\"")
	}
	return nil
}

func (dbu *DiscordBotUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   discordbot.Table,
			Columns: discordbot.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: discordbot.FieldID,
			},
		},
	}
	if ps := dbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if dbu.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   discordbot.AccountTable,
			Columns: []string{discordbot.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: discordaccount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dbu.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   discordbot.AccountTable,
			Columns: []string{discordbot.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: discordaccount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if dbu.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   discordbot.ProjectTable,
			Columns: []string{discordbot.ProjectColumn},
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
	if nodes := dbu.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   discordbot.ProjectTable,
			Columns: []string{discordbot.ProjectColumn},
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
	if dbu.mutation.RepositoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   discordbot.RepositoryTable,
			Columns: []string{discordbot.RepositoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: repository.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dbu.mutation.RepositoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   discordbot.RepositoryTable,
			Columns: []string{discordbot.RepositoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: repository.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{discordbot.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// DiscordBotUpdateOne is the builder for updating a single DiscordBot entity.
type DiscordBotUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DiscordBotMutation
}

// SetAccountID sets the "account" edge to the DiscordAccount entity by ID.
func (dbuo *DiscordBotUpdateOne) SetAccountID(id int) *DiscordBotUpdateOne {
	dbuo.mutation.SetAccountID(id)
	return dbuo
}

// SetAccount sets the "account" edge to the DiscordAccount entity.
func (dbuo *DiscordBotUpdateOne) SetAccount(d *DiscordAccount) *DiscordBotUpdateOne {
	return dbuo.SetAccountID(d.ID)
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (dbuo *DiscordBotUpdateOne) SetProjectID(id int) *DiscordBotUpdateOne {
	dbuo.mutation.SetProjectID(id)
	return dbuo
}

// SetProject sets the "project" edge to the Project entity.
func (dbuo *DiscordBotUpdateOne) SetProject(p *Project) *DiscordBotUpdateOne {
	return dbuo.SetProjectID(p.ID)
}

// SetRepositoryID sets the "repository" edge to the Repository entity by ID.
func (dbuo *DiscordBotUpdateOne) SetRepositoryID(id int) *DiscordBotUpdateOne {
	dbuo.mutation.SetRepositoryID(id)
	return dbuo
}

// SetNillableRepositoryID sets the "repository" edge to the Repository entity by ID if the given value is not nil.
func (dbuo *DiscordBotUpdateOne) SetNillableRepositoryID(id *int) *DiscordBotUpdateOne {
	if id != nil {
		dbuo = dbuo.SetRepositoryID(*id)
	}
	return dbuo
}

// SetRepository sets the "repository" edge to the Repository entity.
func (dbuo *DiscordBotUpdateOne) SetRepository(r *Repository) *DiscordBotUpdateOne {
	return dbuo.SetRepositoryID(r.ID)
}

// Mutation returns the DiscordBotMutation object of the builder.
func (dbuo *DiscordBotUpdateOne) Mutation() *DiscordBotMutation {
	return dbuo.mutation
}

// ClearAccount clears the "account" edge to the DiscordAccount entity.
func (dbuo *DiscordBotUpdateOne) ClearAccount() *DiscordBotUpdateOne {
	dbuo.mutation.ClearAccount()
	return dbuo
}

// ClearProject clears the "project" edge to the Project entity.
func (dbuo *DiscordBotUpdateOne) ClearProject() *DiscordBotUpdateOne {
	dbuo.mutation.ClearProject()
	return dbuo
}

// ClearRepository clears the "repository" edge to the Repository entity.
func (dbuo *DiscordBotUpdateOne) ClearRepository() *DiscordBotUpdateOne {
	dbuo.mutation.ClearRepository()
	return dbuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dbuo *DiscordBotUpdateOne) Select(field string, fields ...string) *DiscordBotUpdateOne {
	dbuo.fields = append([]string{field}, fields...)
	return dbuo
}

// Save executes the query and returns the updated DiscordBot entity.
func (dbuo *DiscordBotUpdateOne) Save(ctx context.Context) (*DiscordBot, error) {
	var (
		err  error
		node *DiscordBot
	)
	if len(dbuo.hooks) == 0 {
		if err = dbuo.check(); err != nil {
			return nil, err
		}
		node, err = dbuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiscordBotMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dbuo.check(); err != nil {
				return nil, err
			}
			dbuo.mutation = mutation
			node, err = dbuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dbuo.hooks) - 1; i >= 0; i-- {
			mut = dbuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dbuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (dbuo *DiscordBotUpdateOne) SaveX(ctx context.Context) *DiscordBot {
	node, err := dbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dbuo *DiscordBotUpdateOne) Exec(ctx context.Context) error {
	_, err := dbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dbuo *DiscordBotUpdateOne) ExecX(ctx context.Context) {
	if err := dbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dbuo *DiscordBotUpdateOne) check() error {
	if _, ok := dbuo.mutation.AccountID(); dbuo.mutation.AccountCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"account\"")
	}
	if _, ok := dbuo.mutation.ProjectID(); dbuo.mutation.ProjectCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"project\"")
	}
	return nil
}

func (dbuo *DiscordBotUpdateOne) sqlSave(ctx context.Context) (_node *DiscordBot, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   discordbot.Table,
			Columns: discordbot.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: discordbot.FieldID,
			},
		},
	}
	id, ok := dbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing DiscordBot.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := dbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, discordbot.FieldID)
		for _, f := range fields {
			if !discordbot.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != discordbot.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if dbuo.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   discordbot.AccountTable,
			Columns: []string{discordbot.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: discordaccount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dbuo.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   discordbot.AccountTable,
			Columns: []string{discordbot.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: discordaccount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if dbuo.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   discordbot.ProjectTable,
			Columns: []string{discordbot.ProjectColumn},
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
	if nodes := dbuo.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   discordbot.ProjectTable,
			Columns: []string{discordbot.ProjectColumn},
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
	if dbuo.mutation.RepositoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   discordbot.RepositoryTable,
			Columns: []string{discordbot.RepositoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: repository.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dbuo.mutation.RepositoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   discordbot.RepositoryTable,
			Columns: []string{discordbot.RepositoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: repository.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &DiscordBot{config: dbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{discordbot.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}