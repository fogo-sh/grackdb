// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/predicate"
	"github.com/fogo-sh/grackdb/ent/project"
	"github.com/fogo-sh/grackdb/ent/projectassociation"
)

// ProjectAssociationUpdate is the builder for updating ProjectAssociation entities.
type ProjectAssociationUpdate struct {
	config
	hooks    []Hook
	mutation *ProjectAssociationMutation
}

// Where adds a new predicate for the ProjectAssociationUpdate builder.
func (pau *ProjectAssociationUpdate) Where(ps ...predicate.ProjectAssociation) *ProjectAssociationUpdate {
	pau.mutation.predicates = append(pau.mutation.predicates, ps...)
	return pau
}

// SetType sets the "type" field.
func (pau *ProjectAssociationUpdate) SetType(pr projectassociation.Type) *ProjectAssociationUpdate {
	pau.mutation.SetType(pr)
	return pau
}

// SetParentID sets the "parent" edge to the Project entity by ID.
func (pau *ProjectAssociationUpdate) SetParentID(id int) *ProjectAssociationUpdate {
	pau.mutation.SetParentID(id)
	return pau
}

// SetNillableParentID sets the "parent" edge to the Project entity by ID if the given value is not nil.
func (pau *ProjectAssociationUpdate) SetNillableParentID(id *int) *ProjectAssociationUpdate {
	if id != nil {
		pau = pau.SetParentID(*id)
	}
	return pau
}

// SetParent sets the "parent" edge to the Project entity.
func (pau *ProjectAssociationUpdate) SetParent(p *Project) *ProjectAssociationUpdate {
	return pau.SetParentID(p.ID)
}

// SetChildID sets the "child" edge to the Project entity by ID.
func (pau *ProjectAssociationUpdate) SetChildID(id int) *ProjectAssociationUpdate {
	pau.mutation.SetChildID(id)
	return pau
}

// SetNillableChildID sets the "child" edge to the Project entity by ID if the given value is not nil.
func (pau *ProjectAssociationUpdate) SetNillableChildID(id *int) *ProjectAssociationUpdate {
	if id != nil {
		pau = pau.SetChildID(*id)
	}
	return pau
}

// SetChild sets the "child" edge to the Project entity.
func (pau *ProjectAssociationUpdate) SetChild(p *Project) *ProjectAssociationUpdate {
	return pau.SetChildID(p.ID)
}

// Mutation returns the ProjectAssociationMutation object of the builder.
func (pau *ProjectAssociationUpdate) Mutation() *ProjectAssociationMutation {
	return pau.mutation
}

// ClearParent clears the "parent" edge to the Project entity.
func (pau *ProjectAssociationUpdate) ClearParent() *ProjectAssociationUpdate {
	pau.mutation.ClearParent()
	return pau
}

// ClearChild clears the "child" edge to the Project entity.
func (pau *ProjectAssociationUpdate) ClearChild() *ProjectAssociationUpdate {
	pau.mutation.ClearChild()
	return pau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pau *ProjectAssociationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pau.hooks) == 0 {
		if err = pau.check(); err != nil {
			return 0, err
		}
		affected, err = pau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectAssociationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pau.check(); err != nil {
				return 0, err
			}
			pau.mutation = mutation
			affected, err = pau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pau.hooks) - 1; i >= 0; i-- {
			mut = pau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pau *ProjectAssociationUpdate) SaveX(ctx context.Context) int {
	affected, err := pau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pau *ProjectAssociationUpdate) Exec(ctx context.Context) error {
	_, err := pau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pau *ProjectAssociationUpdate) ExecX(ctx context.Context) {
	if err := pau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pau *ProjectAssociationUpdate) check() error {
	if v, ok := pau.mutation.GetType(); ok {
		if err := projectassociation.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	return nil
}

func (pau *ProjectAssociationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   projectassociation.Table,
			Columns: projectassociation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: projectassociation.FieldID,
			},
		},
	}
	if ps := pau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pau.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: projectassociation.FieldType,
		})
	}
	if pau.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pau.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pau.mutation.ChildCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pau.mutation.ChildIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projectassociation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ProjectAssociationUpdateOne is the builder for updating a single ProjectAssociation entity.
type ProjectAssociationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProjectAssociationMutation
}

// SetType sets the "type" field.
func (pauo *ProjectAssociationUpdateOne) SetType(pr projectassociation.Type) *ProjectAssociationUpdateOne {
	pauo.mutation.SetType(pr)
	return pauo
}

// SetParentID sets the "parent" edge to the Project entity by ID.
func (pauo *ProjectAssociationUpdateOne) SetParentID(id int) *ProjectAssociationUpdateOne {
	pauo.mutation.SetParentID(id)
	return pauo
}

// SetNillableParentID sets the "parent" edge to the Project entity by ID if the given value is not nil.
func (pauo *ProjectAssociationUpdateOne) SetNillableParentID(id *int) *ProjectAssociationUpdateOne {
	if id != nil {
		pauo = pauo.SetParentID(*id)
	}
	return pauo
}

// SetParent sets the "parent" edge to the Project entity.
func (pauo *ProjectAssociationUpdateOne) SetParent(p *Project) *ProjectAssociationUpdateOne {
	return pauo.SetParentID(p.ID)
}

// SetChildID sets the "child" edge to the Project entity by ID.
func (pauo *ProjectAssociationUpdateOne) SetChildID(id int) *ProjectAssociationUpdateOne {
	pauo.mutation.SetChildID(id)
	return pauo
}

// SetNillableChildID sets the "child" edge to the Project entity by ID if the given value is not nil.
func (pauo *ProjectAssociationUpdateOne) SetNillableChildID(id *int) *ProjectAssociationUpdateOne {
	if id != nil {
		pauo = pauo.SetChildID(*id)
	}
	return pauo
}

// SetChild sets the "child" edge to the Project entity.
func (pauo *ProjectAssociationUpdateOne) SetChild(p *Project) *ProjectAssociationUpdateOne {
	return pauo.SetChildID(p.ID)
}

// Mutation returns the ProjectAssociationMutation object of the builder.
func (pauo *ProjectAssociationUpdateOne) Mutation() *ProjectAssociationMutation {
	return pauo.mutation
}

// ClearParent clears the "parent" edge to the Project entity.
func (pauo *ProjectAssociationUpdateOne) ClearParent() *ProjectAssociationUpdateOne {
	pauo.mutation.ClearParent()
	return pauo
}

// ClearChild clears the "child" edge to the Project entity.
func (pauo *ProjectAssociationUpdateOne) ClearChild() *ProjectAssociationUpdateOne {
	pauo.mutation.ClearChild()
	return pauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pauo *ProjectAssociationUpdateOne) Select(field string, fields ...string) *ProjectAssociationUpdateOne {
	pauo.fields = append([]string{field}, fields...)
	return pauo
}

// Save executes the query and returns the updated ProjectAssociation entity.
func (pauo *ProjectAssociationUpdateOne) Save(ctx context.Context) (*ProjectAssociation, error) {
	var (
		err  error
		node *ProjectAssociation
	)
	if len(pauo.hooks) == 0 {
		if err = pauo.check(); err != nil {
			return nil, err
		}
		node, err = pauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectAssociationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pauo.check(); err != nil {
				return nil, err
			}
			pauo.mutation = mutation
			node, err = pauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pauo.hooks) - 1; i >= 0; i-- {
			mut = pauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (pauo *ProjectAssociationUpdateOne) SaveX(ctx context.Context) *ProjectAssociation {
	node, err := pauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pauo *ProjectAssociationUpdateOne) Exec(ctx context.Context) error {
	_, err := pauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pauo *ProjectAssociationUpdateOne) ExecX(ctx context.Context) {
	if err := pauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pauo *ProjectAssociationUpdateOne) check() error {
	if v, ok := pauo.mutation.GetType(); ok {
		if err := projectassociation.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	return nil
}

func (pauo *ProjectAssociationUpdateOne) sqlSave(ctx context.Context) (_node *ProjectAssociation, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   projectassociation.Table,
			Columns: projectassociation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: projectassociation.FieldID,
			},
		},
	}
	id, ok := pauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ProjectAssociation.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := pauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, projectassociation.FieldID)
		for _, f := range fields {
			if !projectassociation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != projectassociation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pauo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: projectassociation.FieldType,
		})
	}
	if pauo.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pauo.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pauo.mutation.ChildCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pauo.mutation.ChildIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProjectAssociation{config: pauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projectassociation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
