// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/technology"
	"github.com/fogo-sh/grackdb/ent/technologyassociation"
)

// TechnologyAssociationCreate is the builder for creating a TechnologyAssociation entity.
type TechnologyAssociationCreate struct {
	config
	mutation *TechnologyAssociationMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (tac *TechnologyAssociationCreate) SetType(t technologyassociation.Type) *TechnologyAssociationCreate {
	tac.mutation.SetType(t)
	return tac
}

// SetParentID sets the "parent" edge to the Technology entity by ID.
func (tac *TechnologyAssociationCreate) SetParentID(id int) *TechnologyAssociationCreate {
	tac.mutation.SetParentID(id)
	return tac
}

// SetParent sets the "parent" edge to the Technology entity.
func (tac *TechnologyAssociationCreate) SetParent(t *Technology) *TechnologyAssociationCreate {
	return tac.SetParentID(t.ID)
}

// SetChildID sets the "child" edge to the Technology entity by ID.
func (tac *TechnologyAssociationCreate) SetChildID(id int) *TechnologyAssociationCreate {
	tac.mutation.SetChildID(id)
	return tac
}

// SetChild sets the "child" edge to the Technology entity.
func (tac *TechnologyAssociationCreate) SetChild(t *Technology) *TechnologyAssociationCreate {
	return tac.SetChildID(t.ID)
}

// Mutation returns the TechnologyAssociationMutation object of the builder.
func (tac *TechnologyAssociationCreate) Mutation() *TechnologyAssociationMutation {
	return tac.mutation
}

// Save creates the TechnologyAssociation in the database.
func (tac *TechnologyAssociationCreate) Save(ctx context.Context) (*TechnologyAssociation, error) {
	var (
		err  error
		node *TechnologyAssociation
	)
	if len(tac.hooks) == 0 {
		if err = tac.check(); err != nil {
			return nil, err
		}
		node, err = tac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TechnologyAssociationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tac.check(); err != nil {
				return nil, err
			}
			tac.mutation = mutation
			if node, err = tac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tac.hooks) - 1; i >= 0; i-- {
			if tac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tac.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TechnologyAssociation)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TechnologyAssociationMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tac *TechnologyAssociationCreate) SaveX(ctx context.Context) *TechnologyAssociation {
	v, err := tac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tac *TechnologyAssociationCreate) Exec(ctx context.Context) error {
	_, err := tac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tac *TechnologyAssociationCreate) ExecX(ctx context.Context) {
	if err := tac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tac *TechnologyAssociationCreate) check() error {
	if _, ok := tac.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "TechnologyAssociation.type"`)}
	}
	if v, ok := tac.mutation.GetType(); ok {
		if err := technologyassociation.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "TechnologyAssociation.type": %w`, err)}
		}
	}
	if _, ok := tac.mutation.ParentID(); !ok {
		return &ValidationError{Name: "parent", err: errors.New(`ent: missing required edge "TechnologyAssociation.parent"`)}
	}
	if _, ok := tac.mutation.ChildID(); !ok {
		return &ValidationError{Name: "child", err: errors.New(`ent: missing required edge "TechnologyAssociation.child"`)}
	}
	return nil
}

func (tac *TechnologyAssociationCreate) sqlSave(ctx context.Context) (*TechnologyAssociation, error) {
	_node, _spec := tac.createSpec()
	if err := sqlgraph.CreateNode(ctx, tac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tac *TechnologyAssociationCreate) createSpec() (*TechnologyAssociation, *sqlgraph.CreateSpec) {
	var (
		_node = &TechnologyAssociation{config: tac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: technologyassociation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: technologyassociation.FieldID,
			},
		}
	)
	if value, ok := tac.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: technologyassociation.FieldType,
		})
		_node.Type = value
	}
	if nodes := tac.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   technologyassociation.ParentTable,
			Columns: []string{technologyassociation.ParentColumn},
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
		_node.technology_child_technologies = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tac.mutation.ChildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   technologyassociation.ChildTable,
			Columns: []string{technologyassociation.ChildColumn},
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
		_node.technology_parent_technologies = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TechnologyAssociationCreateBulk is the builder for creating many TechnologyAssociation entities in bulk.
type TechnologyAssociationCreateBulk struct {
	config
	builders []*TechnologyAssociationCreate
}

// Save creates the TechnologyAssociation entities in the database.
func (tacb *TechnologyAssociationCreateBulk) Save(ctx context.Context) ([]*TechnologyAssociation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tacb.builders))
	nodes := make([]*TechnologyAssociation, len(tacb.builders))
	mutators := make([]Mutator, len(tacb.builders))
	for i := range tacb.builders {
		func(i int, root context.Context) {
			builder := tacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TechnologyAssociationMutation)
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
					_, err = mutators[i+1].Mutate(root, tacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tacb *TechnologyAssociationCreateBulk) SaveX(ctx context.Context) []*TechnologyAssociation {
	v, err := tacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tacb *TechnologyAssociationCreateBulk) Exec(ctx context.Context) error {
	_, err := tacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tacb *TechnologyAssociationCreateBulk) ExecX(ctx context.Context) {
	if err := tacb.Exec(ctx); err != nil {
		panic(err)
	}
}
