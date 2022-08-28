// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/githuborganization"
	"github.com/fogo-sh/grackdb/ent/githuborganizationmember"
	"github.com/fogo-sh/grackdb/ent/repository"
)

// GithubOrganizationCreate is the builder for creating a GithubOrganization entity.
type GithubOrganizationCreate struct {
	config
	mutation *GithubOrganizationMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (goc *GithubOrganizationCreate) SetName(s string) *GithubOrganizationCreate {
	goc.mutation.SetName(s)
	return goc
}

// SetDisplayName sets the "display_name" field.
func (goc *GithubOrganizationCreate) SetDisplayName(s string) *GithubOrganizationCreate {
	goc.mutation.SetDisplayName(s)
	return goc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (goc *GithubOrganizationCreate) SetNillableDisplayName(s *string) *GithubOrganizationCreate {
	if s != nil {
		goc.SetDisplayName(*s)
	}
	return goc
}

// AddMemberIDs adds the "members" edge to the GithubOrganizationMember entity by IDs.
func (goc *GithubOrganizationCreate) AddMemberIDs(ids ...int) *GithubOrganizationCreate {
	goc.mutation.AddMemberIDs(ids...)
	return goc
}

// AddMembers adds the "members" edges to the GithubOrganizationMember entity.
func (goc *GithubOrganizationCreate) AddMembers(g ...*GithubOrganizationMember) *GithubOrganizationCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return goc.AddMemberIDs(ids...)
}

// AddRepositoryIDs adds the "repositories" edge to the Repository entity by IDs.
func (goc *GithubOrganizationCreate) AddRepositoryIDs(ids ...int) *GithubOrganizationCreate {
	goc.mutation.AddRepositoryIDs(ids...)
	return goc
}

// AddRepositories adds the "repositories" edges to the Repository entity.
func (goc *GithubOrganizationCreate) AddRepositories(r ...*Repository) *GithubOrganizationCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return goc.AddRepositoryIDs(ids...)
}

// Mutation returns the GithubOrganizationMutation object of the builder.
func (goc *GithubOrganizationCreate) Mutation() *GithubOrganizationMutation {
	return goc.mutation
}

// Save creates the GithubOrganization in the database.
func (goc *GithubOrganizationCreate) Save(ctx context.Context) (*GithubOrganization, error) {
	var (
		err  error
		node *GithubOrganization
	)
	if len(goc.hooks) == 0 {
		if err = goc.check(); err != nil {
			return nil, err
		}
		node, err = goc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GithubOrganizationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = goc.check(); err != nil {
				return nil, err
			}
			goc.mutation = mutation
			if node, err = goc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(goc.hooks) - 1; i >= 0; i-- {
			if goc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = goc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, goc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*GithubOrganization)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GithubOrganizationMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (goc *GithubOrganizationCreate) SaveX(ctx context.Context) *GithubOrganization {
	v, err := goc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (goc *GithubOrganizationCreate) Exec(ctx context.Context) error {
	_, err := goc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (goc *GithubOrganizationCreate) ExecX(ctx context.Context) {
	if err := goc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (goc *GithubOrganizationCreate) check() error {
	if _, ok := goc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "GithubOrganization.name"`)}
	}
	if v, ok := goc.mutation.Name(); ok {
		if err := githuborganization.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "GithubOrganization.name": %w`, err)}
		}
	}
	return nil
}

func (goc *GithubOrganizationCreate) sqlSave(ctx context.Context) (*GithubOrganization, error) {
	_node, _spec := goc.createSpec()
	if err := sqlgraph.CreateNode(ctx, goc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (goc *GithubOrganizationCreate) createSpec() (*GithubOrganization, *sqlgraph.CreateSpec) {
	var (
		_node = &GithubOrganization{config: goc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: githuborganization.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: githuborganization.FieldID,
			},
		}
	)
	if value, ok := goc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: githuborganization.FieldName,
		})
		_node.Name = value
	}
	if value, ok := goc.mutation.DisplayName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: githuborganization.FieldDisplayName,
		})
		_node.DisplayName = value
	}
	if nodes := goc.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   githuborganization.MembersTable,
			Columns: []string{githuborganization.MembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: githuborganizationmember.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := goc.mutation.RepositoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   githuborganization.RepositoriesTable,
			Columns: []string{githuborganization.RepositoriesColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GithubOrganizationCreateBulk is the builder for creating many GithubOrganization entities in bulk.
type GithubOrganizationCreateBulk struct {
	config
	builders []*GithubOrganizationCreate
}

// Save creates the GithubOrganization entities in the database.
func (gocb *GithubOrganizationCreateBulk) Save(ctx context.Context) ([]*GithubOrganization, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gocb.builders))
	nodes := make([]*GithubOrganization, len(gocb.builders))
	mutators := make([]Mutator, len(gocb.builders))
	for i := range gocb.builders {
		func(i int, root context.Context) {
			builder := gocb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GithubOrganizationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, gocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, gocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gocb *GithubOrganizationCreateBulk) SaveX(ctx context.Context) []*GithubOrganization {
	v, err := gocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gocb *GithubOrganizationCreateBulk) Exec(ctx context.Context) error {
	_, err := gocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gocb *GithubOrganizationCreateBulk) ExecX(ctx context.Context) {
	if err := gocb.Exec(ctx); err != nil {
		panic(err)
	}
}
