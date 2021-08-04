// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/project"
	"github.com/fogo-sh/grackdb/ent/repository"
	"github.com/fogo-sh/grackdb/ent/site"
)

// SiteCreate is the builder for creating a Site entity.
type SiteCreate struct {
	config
	mutation *SiteMutation
	hooks    []Hook
}

// SetURL sets the "url" field.
func (sc *SiteCreate) SetURL(s string) *SiteCreate {
	sc.mutation.SetURL(s)
	return sc
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (sc *SiteCreate) SetProjectID(id int) *SiteCreate {
	sc.mutation.SetProjectID(id)
	return sc
}

// SetProject sets the "project" edge to the Project entity.
func (sc *SiteCreate) SetProject(p *Project) *SiteCreate {
	return sc.SetProjectID(p.ID)
}

// SetRepositoryID sets the "repository" edge to the Repository entity by ID.
func (sc *SiteCreate) SetRepositoryID(id int) *SiteCreate {
	sc.mutation.SetRepositoryID(id)
	return sc
}

// SetNillableRepositoryID sets the "repository" edge to the Repository entity by ID if the given value is not nil.
func (sc *SiteCreate) SetNillableRepositoryID(id *int) *SiteCreate {
	if id != nil {
		sc = sc.SetRepositoryID(*id)
	}
	return sc
}

// SetRepository sets the "repository" edge to the Repository entity.
func (sc *SiteCreate) SetRepository(r *Repository) *SiteCreate {
	return sc.SetRepositoryID(r.ID)
}

// Mutation returns the SiteMutation object of the builder.
func (sc *SiteCreate) Mutation() *SiteMutation {
	return sc.mutation
}

// Save creates the Site in the database.
func (sc *SiteCreate) Save(ctx context.Context) (*Site, error) {
	var (
		err  error
		node *Site
	)
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SiteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SiteCreate) SaveX(ctx context.Context) *Site {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (sc *SiteCreate) check() error {
	if _, ok := sc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New("ent: missing required field \"url\"")}
	}
	if v, ok := sc.mutation.URL(); ok {
		if err := site.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf("ent: validator failed for field \"url\": %w", err)}
		}
	}
	if _, ok := sc.mutation.ProjectID(); !ok {
		return &ValidationError{Name: "project", err: errors.New("ent: missing required edge \"project\"")}
	}
	return nil
}

func (sc *SiteCreate) sqlSave(ctx context.Context) (*Site, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sc *SiteCreate) createSpec() (*Site, *sqlgraph.CreateSpec) {
	var (
		_node = &Site{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: site.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: site.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: site.FieldURL,
		})
		_node.URL = value
	}
	if nodes := sc.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   site.ProjectTable,
			Columns: []string{site.ProjectColumn},
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
		_node.project_sites = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.RepositoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   site.RepositoryTable,
			Columns: []string{site.RepositoryColumn},
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
		_node.repository_sites = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SiteCreateBulk is the builder for creating many Site entities in bulk.
type SiteCreateBulk struct {
	config
	builders []*SiteCreate
}

// Save creates the Site entities in the database.
func (scb *SiteCreateBulk) Save(ctx context.Context) ([]*Site, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Site, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SiteMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SiteCreateBulk) SaveX(ctx context.Context) []*Site {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}