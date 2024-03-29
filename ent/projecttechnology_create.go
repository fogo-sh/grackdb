// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/project"
	"github.com/fogo-sh/grackdb/ent/projecttechnology"
	"github.com/fogo-sh/grackdb/ent/technology"
)

// ProjectTechnologyCreate is the builder for creating a ProjectTechnology entity.
type ProjectTechnologyCreate struct {
	config
	mutation *ProjectTechnologyMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (ptc *ProjectTechnologyCreate) SetType(pr projecttechnology.Type) *ProjectTechnologyCreate {
	ptc.mutation.SetType(pr)
	return ptc
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (ptc *ProjectTechnologyCreate) SetProjectID(id int) *ProjectTechnologyCreate {
	ptc.mutation.SetProjectID(id)
	return ptc
}

// SetProject sets the "project" edge to the Project entity.
func (ptc *ProjectTechnologyCreate) SetProject(p *Project) *ProjectTechnologyCreate {
	return ptc.SetProjectID(p.ID)
}

// SetTechnologyID sets the "technology" edge to the Technology entity by ID.
func (ptc *ProjectTechnologyCreate) SetTechnologyID(id int) *ProjectTechnologyCreate {
	ptc.mutation.SetTechnologyID(id)
	return ptc
}

// SetTechnology sets the "technology" edge to the Technology entity.
func (ptc *ProjectTechnologyCreate) SetTechnology(t *Technology) *ProjectTechnologyCreate {
	return ptc.SetTechnologyID(t.ID)
}

// Mutation returns the ProjectTechnologyMutation object of the builder.
func (ptc *ProjectTechnologyCreate) Mutation() *ProjectTechnologyMutation {
	return ptc.mutation
}

// Save creates the ProjectTechnology in the database.
func (ptc *ProjectTechnologyCreate) Save(ctx context.Context) (*ProjectTechnology, error) {
	var (
		err  error
		node *ProjectTechnology
	)
	if len(ptc.hooks) == 0 {
		if err = ptc.check(); err != nil {
			return nil, err
		}
		node, err = ptc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectTechnologyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ptc.check(); err != nil {
				return nil, err
			}
			ptc.mutation = mutation
			if node, err = ptc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ptc.hooks) - 1; i >= 0; i-- {
			if ptc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ptc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ptc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ProjectTechnology)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ProjectTechnologyMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ptc *ProjectTechnologyCreate) SaveX(ctx context.Context) *ProjectTechnology {
	v, err := ptc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ptc *ProjectTechnologyCreate) Exec(ctx context.Context) error {
	_, err := ptc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptc *ProjectTechnologyCreate) ExecX(ctx context.Context) {
	if err := ptc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptc *ProjectTechnologyCreate) check() error {
	if _, ok := ptc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "ProjectTechnology.type"`)}
	}
	if v, ok := ptc.mutation.GetType(); ok {
		if err := projecttechnology.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "ProjectTechnology.type": %w`, err)}
		}
	}
	if _, ok := ptc.mutation.ProjectID(); !ok {
		return &ValidationError{Name: "project", err: errors.New(`ent: missing required edge "ProjectTechnology.project"`)}
	}
	if _, ok := ptc.mutation.TechnologyID(); !ok {
		return &ValidationError{Name: "technology", err: errors.New(`ent: missing required edge "ProjectTechnology.technology"`)}
	}
	return nil
}

func (ptc *ProjectTechnologyCreate) sqlSave(ctx context.Context) (*ProjectTechnology, error) {
	_node, _spec := ptc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ptc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ptc *ProjectTechnologyCreate) createSpec() (*ProjectTechnology, *sqlgraph.CreateSpec) {
	var (
		_node = &ProjectTechnology{config: ptc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: projecttechnology.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: projecttechnology.FieldID,
			},
		}
	)
	if value, ok := ptc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: projecttechnology.FieldType,
		})
		_node.Type = value
	}
	if nodes := ptc.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttechnology.ProjectTable,
			Columns: []string{projecttechnology.ProjectColumn},
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
		_node.project_technologies = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ptc.mutation.TechnologyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projecttechnology.TechnologyTable,
			Columns: []string{projecttechnology.TechnologyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: technology.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.technology_projects = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProjectTechnologyCreateBulk is the builder for creating many ProjectTechnology entities in bulk.
type ProjectTechnologyCreateBulk struct {
	config
	builders []*ProjectTechnologyCreate
}

// Save creates the ProjectTechnology entities in the database.
func (ptcb *ProjectTechnologyCreateBulk) Save(ctx context.Context) ([]*ProjectTechnology, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ptcb.builders))
	nodes := make([]*ProjectTechnology, len(ptcb.builders))
	mutators := make([]Mutator, len(ptcb.builders))
	for i := range ptcb.builders {
		func(i int, root context.Context) {
			builder := ptcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProjectTechnologyMutation)
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
					_, err = mutators[i+1].Mutate(root, ptcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ptcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ptcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ptcb *ProjectTechnologyCreateBulk) SaveX(ctx context.Context) []*ProjectTechnology {
	v, err := ptcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ptcb *ProjectTechnologyCreateBulk) Exec(ctx context.Context) error {
	_, err := ptcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptcb *ProjectTechnologyCreateBulk) ExecX(ctx context.Context) {
	if err := ptcb.Exec(ctx); err != nil {
		panic(err)
	}
}
