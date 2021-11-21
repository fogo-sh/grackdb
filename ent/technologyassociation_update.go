// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/predicate"
	"github.com/fogo-sh/grackdb/ent/technology"
	"github.com/fogo-sh/grackdb/ent/technologyassociation"
)

// TechnologyAssociationUpdate is the builder for updating TechnologyAssociation entities.
type TechnologyAssociationUpdate struct {
	config
	hooks    []Hook
	mutation *TechnologyAssociationMutation
}

// Where appends a list predicates to the TechnologyAssociationUpdate builder.
func (tau *TechnologyAssociationUpdate) Where(ps ...predicate.TechnologyAssociation) *TechnologyAssociationUpdate {
	tau.mutation.Where(ps...)
	return tau
}

// SetType sets the "type" field.
func (tau *TechnologyAssociationUpdate) SetType(t technologyassociation.Type) *TechnologyAssociationUpdate {
	tau.mutation.SetType(t)
	return tau
}

// SetParentID sets the "parent" edge to the Technology entity by ID.
func (tau *TechnologyAssociationUpdate) SetParentID(id int) *TechnologyAssociationUpdate {
	tau.mutation.SetParentID(id)
	return tau
}

// SetParent sets the "parent" edge to the Technology entity.
func (tau *TechnologyAssociationUpdate) SetParent(t *Technology) *TechnologyAssociationUpdate {
	return tau.SetParentID(t.ID)
}

// SetChildID sets the "child" edge to the Technology entity by ID.
func (tau *TechnologyAssociationUpdate) SetChildID(id int) *TechnologyAssociationUpdate {
	tau.mutation.SetChildID(id)
	return tau
}

// SetChild sets the "child" edge to the Technology entity.
func (tau *TechnologyAssociationUpdate) SetChild(t *Technology) *TechnologyAssociationUpdate {
	return tau.SetChildID(t.ID)
}

// Mutation returns the TechnologyAssociationMutation object of the builder.
func (tau *TechnologyAssociationUpdate) Mutation() *TechnologyAssociationMutation {
	return tau.mutation
}

// ClearParent clears the "parent" edge to the Technology entity.
func (tau *TechnologyAssociationUpdate) ClearParent() *TechnologyAssociationUpdate {
	tau.mutation.ClearParent()
	return tau
}

// ClearChild clears the "child" edge to the Technology entity.
func (tau *TechnologyAssociationUpdate) ClearChild() *TechnologyAssociationUpdate {
	tau.mutation.ClearChild()
	return tau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tau *TechnologyAssociationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tau.hooks) == 0 {
		if err = tau.check(); err != nil {
			return 0, err
		}
		affected, err = tau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TechnologyAssociationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tau.check(); err != nil {
				return 0, err
			}
			tau.mutation = mutation
			affected, err = tau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tau.hooks) - 1; i >= 0; i-- {
			if tau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tau *TechnologyAssociationUpdate) SaveX(ctx context.Context) int {
	affected, err := tau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tau *TechnologyAssociationUpdate) Exec(ctx context.Context) error {
	_, err := tau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tau *TechnologyAssociationUpdate) ExecX(ctx context.Context) {
	if err := tau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tau *TechnologyAssociationUpdate) check() error {
	if v, ok := tau.mutation.GetType(); ok {
		if err := technologyassociation.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if _, ok := tau.mutation.ParentID(); tau.mutation.ParentCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"parent\"")
	}
	if _, ok := tau.mutation.ChildID(); tau.mutation.ChildCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"child\"")
	}
	return nil
}

func (tau *TechnologyAssociationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   technologyassociation.Table,
			Columns: technologyassociation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: technologyassociation.FieldID,
			},
		},
	}
	if ps := tau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tau.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: technologyassociation.FieldType,
		})
	}
	if tau.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tau.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tau.mutation.ChildCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tau.mutation.ChildIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{technologyassociation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TechnologyAssociationUpdateOne is the builder for updating a single TechnologyAssociation entity.
type TechnologyAssociationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TechnologyAssociationMutation
}

// SetType sets the "type" field.
func (tauo *TechnologyAssociationUpdateOne) SetType(t technologyassociation.Type) *TechnologyAssociationUpdateOne {
	tauo.mutation.SetType(t)
	return tauo
}

// SetParentID sets the "parent" edge to the Technology entity by ID.
func (tauo *TechnologyAssociationUpdateOne) SetParentID(id int) *TechnologyAssociationUpdateOne {
	tauo.mutation.SetParentID(id)
	return tauo
}

// SetParent sets the "parent" edge to the Technology entity.
func (tauo *TechnologyAssociationUpdateOne) SetParent(t *Technology) *TechnologyAssociationUpdateOne {
	return tauo.SetParentID(t.ID)
}

// SetChildID sets the "child" edge to the Technology entity by ID.
func (tauo *TechnologyAssociationUpdateOne) SetChildID(id int) *TechnologyAssociationUpdateOne {
	tauo.mutation.SetChildID(id)
	return tauo
}

// SetChild sets the "child" edge to the Technology entity.
func (tauo *TechnologyAssociationUpdateOne) SetChild(t *Technology) *TechnologyAssociationUpdateOne {
	return tauo.SetChildID(t.ID)
}

// Mutation returns the TechnologyAssociationMutation object of the builder.
func (tauo *TechnologyAssociationUpdateOne) Mutation() *TechnologyAssociationMutation {
	return tauo.mutation
}

// ClearParent clears the "parent" edge to the Technology entity.
func (tauo *TechnologyAssociationUpdateOne) ClearParent() *TechnologyAssociationUpdateOne {
	tauo.mutation.ClearParent()
	return tauo
}

// ClearChild clears the "child" edge to the Technology entity.
func (tauo *TechnologyAssociationUpdateOne) ClearChild() *TechnologyAssociationUpdateOne {
	tauo.mutation.ClearChild()
	return tauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tauo *TechnologyAssociationUpdateOne) Select(field string, fields ...string) *TechnologyAssociationUpdateOne {
	tauo.fields = append([]string{field}, fields...)
	return tauo
}

// Save executes the query and returns the updated TechnologyAssociation entity.
func (tauo *TechnologyAssociationUpdateOne) Save(ctx context.Context) (*TechnologyAssociation, error) {
	var (
		err  error
		node *TechnologyAssociation
	)
	if len(tauo.hooks) == 0 {
		if err = tauo.check(); err != nil {
			return nil, err
		}
		node, err = tauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TechnologyAssociationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tauo.check(); err != nil {
				return nil, err
			}
			tauo.mutation = mutation
			node, err = tauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tauo.hooks) - 1; i >= 0; i-- {
			if tauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tauo *TechnologyAssociationUpdateOne) SaveX(ctx context.Context) *TechnologyAssociation {
	node, err := tauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tauo *TechnologyAssociationUpdateOne) Exec(ctx context.Context) error {
	_, err := tauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tauo *TechnologyAssociationUpdateOne) ExecX(ctx context.Context) {
	if err := tauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tauo *TechnologyAssociationUpdateOne) check() error {
	if v, ok := tauo.mutation.GetType(); ok {
		if err := technologyassociation.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if _, ok := tauo.mutation.ParentID(); tauo.mutation.ParentCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"parent\"")
	}
	if _, ok := tauo.mutation.ChildID(); tauo.mutation.ChildCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"child\"")
	}
	return nil
}

func (tauo *TechnologyAssociationUpdateOne) sqlSave(ctx context.Context) (_node *TechnologyAssociation, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   technologyassociation.Table,
			Columns: technologyassociation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: technologyassociation.FieldID,
			},
		},
	}
	id, ok := tauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing TechnologyAssociation.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := tauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, technologyassociation.FieldID)
		for _, f := range fields {
			if !technologyassociation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != technologyassociation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tauo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: technologyassociation.FieldType,
		})
	}
	if tauo.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tauo.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tauo.mutation.ChildCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tauo.mutation.ChildIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TechnologyAssociation{config: tauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{technologyassociation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
