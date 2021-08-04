// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/githubaccount"
	"github.com/fogo-sh/grackdb/ent/githuborganization"
	"github.com/fogo-sh/grackdb/ent/githuborganizationmember"
)

// GithubOrganizationMemberCreate is the builder for creating a GithubOrganizationMember entity.
type GithubOrganizationMemberCreate struct {
	config
	mutation *GithubOrganizationMemberMutation
	hooks    []Hook
}

// SetRole sets the "role" field.
func (gomc *GithubOrganizationMemberCreate) SetRole(gi githuborganizationmember.Role) *GithubOrganizationMemberCreate {
	gomc.mutation.SetRole(gi)
	return gomc
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (gomc *GithubOrganizationMemberCreate) SetNillableRole(gi *githuborganizationmember.Role) *GithubOrganizationMemberCreate {
	if gi != nil {
		gomc.SetRole(*gi)
	}
	return gomc
}

// SetOrganizationID sets the "organization" edge to the GithubOrganization entity by ID.
func (gomc *GithubOrganizationMemberCreate) SetOrganizationID(id int) *GithubOrganizationMemberCreate {
	gomc.mutation.SetOrganizationID(id)
	return gomc
}

// SetOrganization sets the "organization" edge to the GithubOrganization entity.
func (gomc *GithubOrganizationMemberCreate) SetOrganization(g *GithubOrganization) *GithubOrganizationMemberCreate {
	return gomc.SetOrganizationID(g.ID)
}

// SetAccountID sets the "account" edge to the GithubAccount entity by ID.
func (gomc *GithubOrganizationMemberCreate) SetAccountID(id int) *GithubOrganizationMemberCreate {
	gomc.mutation.SetAccountID(id)
	return gomc
}

// SetAccount sets the "account" edge to the GithubAccount entity.
func (gomc *GithubOrganizationMemberCreate) SetAccount(g *GithubAccount) *GithubOrganizationMemberCreate {
	return gomc.SetAccountID(g.ID)
}

// Mutation returns the GithubOrganizationMemberMutation object of the builder.
func (gomc *GithubOrganizationMemberCreate) Mutation() *GithubOrganizationMemberMutation {
	return gomc.mutation
}

// Save creates the GithubOrganizationMember in the database.
func (gomc *GithubOrganizationMemberCreate) Save(ctx context.Context) (*GithubOrganizationMember, error) {
	var (
		err  error
		node *GithubOrganizationMember
	)
	gomc.defaults()
	if len(gomc.hooks) == 0 {
		if err = gomc.check(); err != nil {
			return nil, err
		}
		node, err = gomc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GithubOrganizationMemberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gomc.check(); err != nil {
				return nil, err
			}
			gomc.mutation = mutation
			if node, err = gomc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gomc.hooks) - 1; i >= 0; i-- {
			mut = gomc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gomc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gomc *GithubOrganizationMemberCreate) SaveX(ctx context.Context) *GithubOrganizationMember {
	v, err := gomc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (gomc *GithubOrganizationMemberCreate) defaults() {
	if _, ok := gomc.mutation.Role(); !ok {
		v := githuborganizationmember.DefaultRole
		gomc.mutation.SetRole(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gomc *GithubOrganizationMemberCreate) check() error {
	if _, ok := gomc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New("ent: missing required field \"role\"")}
	}
	if v, ok := gomc.mutation.Role(); ok {
		if err := githuborganizationmember.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf("ent: validator failed for field \"role\": %w", err)}
		}
	}
	if _, ok := gomc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization", err: errors.New("ent: missing required edge \"organization\"")}
	}
	if _, ok := gomc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account", err: errors.New("ent: missing required edge \"account\"")}
	}
	return nil
}

func (gomc *GithubOrganizationMemberCreate) sqlSave(ctx context.Context) (*GithubOrganizationMember, error) {
	_node, _spec := gomc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gomc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (gomc *GithubOrganizationMemberCreate) createSpec() (*GithubOrganizationMember, *sqlgraph.CreateSpec) {
	var (
		_node = &GithubOrganizationMember{config: gomc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: githuborganizationmember.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: githuborganizationmember.FieldID,
			},
		}
	)
	if value, ok := gomc.mutation.Role(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: githuborganizationmember.FieldRole,
		})
		_node.Role = value
	}
	if nodes := gomc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   githuborganizationmember.OrganizationTable,
			Columns: []string{githuborganizationmember.OrganizationColumn},
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
		_node.github_organization_members = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gomc.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   githuborganizationmember.AccountTable,
			Columns: []string{githuborganizationmember.AccountColumn},
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
		_node.github_account_organization_memberships = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GithubOrganizationMemberCreateBulk is the builder for creating many GithubOrganizationMember entities in bulk.
type GithubOrganizationMemberCreateBulk struct {
	config
	builders []*GithubOrganizationMemberCreate
}

// Save creates the GithubOrganizationMember entities in the database.
func (gomcb *GithubOrganizationMemberCreateBulk) Save(ctx context.Context) ([]*GithubOrganizationMember, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gomcb.builders))
	nodes := make([]*GithubOrganizationMember, len(gomcb.builders))
	mutators := make([]Mutator, len(gomcb.builders))
	for i := range gomcb.builders {
		func(i int, root context.Context) {
			builder := gomcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GithubOrganizationMemberMutation)
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
					_, err = mutators[i+1].Mutate(root, gomcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gomcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, gomcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gomcb *GithubOrganizationMemberCreateBulk) SaveX(ctx context.Context) []*GithubOrganizationMember {
	v, err := gomcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
