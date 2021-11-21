// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/project"
	"github.com/fogo-sh/grackdb/ent/projectassociation"
)

// ProjectAssociationCreate is the builder for creating a ProjectAssociation entity.
type ProjectAssociationCreate struct {
	config
	mutation *ProjectAssociationMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (pac *ProjectAssociationCreate) SetType(pr projectassociation.Type) *ProjectAssociationCreate {
	pac.mutation.SetType(pr)
	return pac
}

// SetParentID sets the "parent" edge to the Project entity by ID.
func (pac *ProjectAssociationCreate) SetParentID(id int) *ProjectAssociationCreate {
	pac.mutation.SetParentID(id)
	return pac
}

// SetParent sets the "parent" edge to the Project entity.
func (pac *ProjectAssociationCreate) SetParent(p *Project) *ProjectAssociationCreate {
	return pac.SetParentID(p.ID)
}

// SetChildID sets the "child" edge to the Project entity by ID.
func (pac *ProjectAssociationCreate) SetChildID(id int) *ProjectAssociationCreate {
	pac.mutation.SetChildID(id)
	return pac
}

// SetChild sets the "child" edge to the Project entity.
func (pac *ProjectAssociationCreate) SetChild(p *Project) *ProjectAssociationCreate {
	return pac.SetChildID(p.ID)
}

// Mutation returns the ProjectAssociationMutation object of the builder.
func (pac *ProjectAssociationCreate) Mutation() *ProjectAssociationMutation {
	return pac.mutation
}

// Save creates the ProjectAssociation in the database.
func (pac *ProjectAssociationCreate) Save(ctx context.Context) (*ProjectAssociation, error) {
	var (
		err  error
		node *ProjectAssociation
	)
	if len(pac.hooks) == 0 {
		if err = pac.check(); err != nil {
			return nil, err
		}
		node, err = pac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectAssociationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pac.check(); err != nil {
				return nil, err
			}
			pac.mutation = mutation
			if node, err = pac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pac.hooks) - 1; i >= 0; i-- {
			if pac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pac *ProjectAssociationCreate) SaveX(ctx context.Context) *ProjectAssociation {
	v, err := pac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pac *ProjectAssociationCreate) Exec(ctx context.Context) error {
	_, err := pac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pac *ProjectAssociationCreate) ExecX(ctx context.Context) {
	if err := pac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pac *ProjectAssociationCreate) check() error {
	if _, ok := pac.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "type"`)}
	}
	if v, ok := pac.mutation.GetType(); ok {
		if err := projectassociation.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "type": %w`, err)}
		}
	}
	if _, ok := pac.mutation.ParentID(); !ok {
		return &ValidationError{Name: "parent", err: errors.New("ent: missing required edge \"parent\"")}
	}
	if _, ok := pac.mutation.ChildID(); !ok {
		return &ValidationError{Name: "child", err: errors.New("ent: missing required edge \"child\"")}
	}
	return nil
}

func (pac *ProjectAssociationCreate) sqlSave(ctx context.Context) (*ProjectAssociation, error) {
	_node, _spec := pac.createSpec()
	if err := sqlgraph.CreateNode(ctx, pac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pac *ProjectAssociationCreate) createSpec() (*ProjectAssociation, *sqlgraph.CreateSpec) {
	var (
		_node = &ProjectAssociation{config: pac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: projectassociation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: projectassociation.FieldID,
			},
		}
	)
	if value, ok := pac.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: projectassociation.FieldType,
		})
		_node.Type = value
	}
	if nodes := pac.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectassociation.ParentTable,
			Columns: []string{projectassociation.ParentColumn},
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
		_node.project_child_projects = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pac.mutation.ChildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectassociation.ChildTable,
			Columns: []string{projectassociation.ChildColumn},
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
		_node.project_parent_projects = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProjectAssociationCreateBulk is the builder for creating many ProjectAssociation entities in bulk.
type ProjectAssociationCreateBulk struct {
	config
	builders []*ProjectAssociationCreate
}

// Save creates the ProjectAssociation entities in the database.
func (pacb *ProjectAssociationCreateBulk) Save(ctx context.Context) ([]*ProjectAssociation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pacb.builders))
	nodes := make([]*ProjectAssociation, len(pacb.builders))
	mutators := make([]Mutator, len(pacb.builders))
	for i := range pacb.builders {
		func(i int, root context.Context) {
			builder := pacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProjectAssociationMutation)
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
					_, err = mutators[i+1].Mutate(root, pacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pacb *ProjectAssociationCreateBulk) SaveX(ctx context.Context) []*ProjectAssociation {
	v, err := pacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pacb *ProjectAssociationCreateBulk) Exec(ctx context.Context) error {
	_, err := pacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pacb *ProjectAssociationCreateBulk) ExecX(ctx context.Context) {
	if err := pacb.Exec(ctx); err != nil {
		panic(err)
	}
}
