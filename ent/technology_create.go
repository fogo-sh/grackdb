// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/projecttechnology"
	"github.com/fogo-sh/grackdb/ent/repositorytechnology"
	"github.com/fogo-sh/grackdb/ent/technology"
	"github.com/fogo-sh/grackdb/ent/technologyassociation"
)

// TechnologyCreate is the builder for creating a Technology entity.
type TechnologyCreate struct {
	config
	mutation *TechnologyMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (tc *TechnologyCreate) SetName(s string) *TechnologyCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetDescription sets the "description" field.
func (tc *TechnologyCreate) SetDescription(s string) *TechnologyCreate {
	tc.mutation.SetDescription(s)
	return tc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tc *TechnologyCreate) SetNillableDescription(s *string) *TechnologyCreate {
	if s != nil {
		tc.SetDescription(*s)
	}
	return tc
}

// SetColour sets the "colour" field.
func (tc *TechnologyCreate) SetColour(s string) *TechnologyCreate {
	tc.mutation.SetColour(s)
	return tc
}

// SetNillableColour sets the "colour" field if the given value is not nil.
func (tc *TechnologyCreate) SetNillableColour(s *string) *TechnologyCreate {
	if s != nil {
		tc.SetColour(*s)
	}
	return tc
}

// SetType sets the "type" field.
func (tc *TechnologyCreate) SetType(t technology.Type) *TechnologyCreate {
	tc.mutation.SetType(t)
	return tc
}

// AddParentTechnologyIDs adds the "parent_technologies" edge to the TechnologyAssociation entity by IDs.
func (tc *TechnologyCreate) AddParentTechnologyIDs(ids ...int) *TechnologyCreate {
	tc.mutation.AddParentTechnologyIDs(ids...)
	return tc
}

// AddParentTechnologies adds the "parent_technologies" edges to the TechnologyAssociation entity.
func (tc *TechnologyCreate) AddParentTechnologies(t ...*TechnologyAssociation) *TechnologyCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddParentTechnologyIDs(ids...)
}

// AddChildTechnologyIDs adds the "child_technologies" edge to the TechnologyAssociation entity by IDs.
func (tc *TechnologyCreate) AddChildTechnologyIDs(ids ...int) *TechnologyCreate {
	tc.mutation.AddChildTechnologyIDs(ids...)
	return tc
}

// AddChildTechnologies adds the "child_technologies" edges to the TechnologyAssociation entity.
func (tc *TechnologyCreate) AddChildTechnologies(t ...*TechnologyAssociation) *TechnologyCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddChildTechnologyIDs(ids...)
}

// AddProjectIDs adds the "projects" edge to the ProjectTechnology entity by IDs.
func (tc *TechnologyCreate) AddProjectIDs(ids ...int) *TechnologyCreate {
	tc.mutation.AddProjectIDs(ids...)
	return tc
}

// AddProjects adds the "projects" edges to the ProjectTechnology entity.
func (tc *TechnologyCreate) AddProjects(p ...*ProjectTechnology) *TechnologyCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tc.AddProjectIDs(ids...)
}

// AddRepositoryIDs adds the "repositories" edge to the RepositoryTechnology entity by IDs.
func (tc *TechnologyCreate) AddRepositoryIDs(ids ...int) *TechnologyCreate {
	tc.mutation.AddRepositoryIDs(ids...)
	return tc
}

// AddRepositories adds the "repositories" edges to the RepositoryTechnology entity.
func (tc *TechnologyCreate) AddRepositories(r ...*RepositoryTechnology) *TechnologyCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tc.AddRepositoryIDs(ids...)
}

// Mutation returns the TechnologyMutation object of the builder.
func (tc *TechnologyCreate) Mutation() *TechnologyMutation {
	return tc.mutation
}

// Save creates the Technology in the database.
func (tc *TechnologyCreate) Save(ctx context.Context) (*Technology, error) {
	var (
		err  error
		node *Technology
	)
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TechnologyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TechnologyCreate) SaveX(ctx context.Context) *Technology {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TechnologyCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TechnologyCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TechnologyCreate) check() error {
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := tc.mutation.Name(); ok {
		if err := technology.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := tc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "type"`)}
	}
	if v, ok := tc.mutation.GetType(); ok {
		if err := technology.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "type": %w`, err)}
		}
	}
	return nil
}

func (tc *TechnologyCreate) sqlSave(ctx context.Context) (*Technology, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tc *TechnologyCreate) createSpec() (*Technology, *sqlgraph.CreateSpec) {
	var (
		_node = &Technology{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: technology.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: technology.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: technology.FieldName,
		})
		_node.Name = value
	}
	if value, ok := tc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: technology.FieldDescription,
		})
		_node.Description = &value
	}
	if value, ok := tc.mutation.Colour(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: technology.FieldColour,
		})
		_node.Colour = &value
	}
	if value, ok := tc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: technology.FieldType,
		})
		_node.Type = value
	}
	if nodes := tc.mutation.ParentTechnologiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   technology.ParentTechnologiesTable,
			Columns: []string{technology.ParentTechnologiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: technologyassociation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ChildTechnologiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   technology.ChildTechnologiesTable,
			Columns: []string{technology.ChildTechnologiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: technologyassociation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   technology.ProjectsTable,
			Columns: []string{technology.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: projecttechnology.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.RepositoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   technology.RepositoriesTable,
			Columns: []string{technology.RepositoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: repositorytechnology.FieldID,
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

// TechnologyCreateBulk is the builder for creating many Technology entities in bulk.
type TechnologyCreateBulk struct {
	config
	builders []*TechnologyCreate
}

// Save creates the Technology entities in the database.
func (tcb *TechnologyCreateBulk) Save(ctx context.Context) ([]*Technology, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Technology, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TechnologyMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TechnologyCreateBulk) SaveX(ctx context.Context) []*Technology {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TechnologyCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TechnologyCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
