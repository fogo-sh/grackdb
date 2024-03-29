// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/predicate"
	"github.com/fogo-sh/grackdb/ent/project"
	"github.com/fogo-sh/grackdb/ent/projecttechnology"
	"github.com/fogo-sh/grackdb/ent/technology"
)

// ProjectTechnologyUpdate is the builder for updating ProjectTechnology entities.
type ProjectTechnologyUpdate struct {
	config
	hooks    []Hook
	mutation *ProjectTechnologyMutation
}

// Where appends a list predicates to the ProjectTechnologyUpdate builder.
func (ptu *ProjectTechnologyUpdate) Where(ps ...predicate.ProjectTechnology) *ProjectTechnologyUpdate {
	ptu.mutation.Where(ps...)
	return ptu
}

// SetType sets the "type" field.
func (ptu *ProjectTechnologyUpdate) SetType(pr projecttechnology.Type) *ProjectTechnologyUpdate {
	ptu.mutation.SetType(pr)
	return ptu
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (ptu *ProjectTechnologyUpdate) SetProjectID(id int) *ProjectTechnologyUpdate {
	ptu.mutation.SetProjectID(id)
	return ptu
}

// SetProject sets the "project" edge to the Project entity.
func (ptu *ProjectTechnologyUpdate) SetProject(p *Project) *ProjectTechnologyUpdate {
	return ptu.SetProjectID(p.ID)
}

// SetTechnologyID sets the "technology" edge to the Technology entity by ID.
func (ptu *ProjectTechnologyUpdate) SetTechnologyID(id int) *ProjectTechnologyUpdate {
	ptu.mutation.SetTechnologyID(id)
	return ptu
}

// SetTechnology sets the "technology" edge to the Technology entity.
func (ptu *ProjectTechnologyUpdate) SetTechnology(t *Technology) *ProjectTechnologyUpdate {
	return ptu.SetTechnologyID(t.ID)
}

// Mutation returns the ProjectTechnologyMutation object of the builder.
func (ptu *ProjectTechnologyUpdate) Mutation() *ProjectTechnologyMutation {
	return ptu.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (ptu *ProjectTechnologyUpdate) ClearProject() *ProjectTechnologyUpdate {
	ptu.mutation.ClearProject()
	return ptu
}

// ClearTechnology clears the "technology" edge to the Technology entity.
func (ptu *ProjectTechnologyUpdate) ClearTechnology() *ProjectTechnologyUpdate {
	ptu.mutation.ClearTechnology()
	return ptu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ptu *ProjectTechnologyUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ptu.hooks) == 0 {
		if err = ptu.check(); err != nil {
			return 0, err
		}
		affected, err = ptu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectTechnologyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ptu.check(); err != nil {
				return 0, err
			}
			ptu.mutation = mutation
			affected, err = ptu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ptu.hooks) - 1; i >= 0; i-- {
			if ptu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ptu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ptu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ptu *ProjectTechnologyUpdate) SaveX(ctx context.Context) int {
	affected, err := ptu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ptu *ProjectTechnologyUpdate) Exec(ctx context.Context) error {
	_, err := ptu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptu *ProjectTechnologyUpdate) ExecX(ctx context.Context) {
	if err := ptu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptu *ProjectTechnologyUpdate) check() error {
	if v, ok := ptu.mutation.GetType(); ok {
		if err := projecttechnology.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "ProjectTechnology.type": %w`, err)}
		}
	}
	if _, ok := ptu.mutation.ProjectID(); ptu.mutation.ProjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProjectTechnology.project"`)
	}
	if _, ok := ptu.mutation.TechnologyID(); ptu.mutation.TechnologyCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProjectTechnology.technology"`)
	}
	return nil
}

func (ptu *ProjectTechnologyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   projecttechnology.Table,
			Columns: projecttechnology.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: projecttechnology.FieldID,
			},
		},
	}
	if ps := ptu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ptu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: projecttechnology.FieldType,
		})
	}
	if ptu.mutation.ProjectCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptu.mutation.ProjectIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ptu.mutation.TechnologyCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptu.mutation.TechnologyIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ptu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projecttechnology.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ProjectTechnologyUpdateOne is the builder for updating a single ProjectTechnology entity.
type ProjectTechnologyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProjectTechnologyMutation
}

// SetType sets the "type" field.
func (ptuo *ProjectTechnologyUpdateOne) SetType(pr projecttechnology.Type) *ProjectTechnologyUpdateOne {
	ptuo.mutation.SetType(pr)
	return ptuo
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (ptuo *ProjectTechnologyUpdateOne) SetProjectID(id int) *ProjectTechnologyUpdateOne {
	ptuo.mutation.SetProjectID(id)
	return ptuo
}

// SetProject sets the "project" edge to the Project entity.
func (ptuo *ProjectTechnologyUpdateOne) SetProject(p *Project) *ProjectTechnologyUpdateOne {
	return ptuo.SetProjectID(p.ID)
}

// SetTechnologyID sets the "technology" edge to the Technology entity by ID.
func (ptuo *ProjectTechnologyUpdateOne) SetTechnologyID(id int) *ProjectTechnologyUpdateOne {
	ptuo.mutation.SetTechnologyID(id)
	return ptuo
}

// SetTechnology sets the "technology" edge to the Technology entity.
func (ptuo *ProjectTechnologyUpdateOne) SetTechnology(t *Technology) *ProjectTechnologyUpdateOne {
	return ptuo.SetTechnologyID(t.ID)
}

// Mutation returns the ProjectTechnologyMutation object of the builder.
func (ptuo *ProjectTechnologyUpdateOne) Mutation() *ProjectTechnologyMutation {
	return ptuo.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (ptuo *ProjectTechnologyUpdateOne) ClearProject() *ProjectTechnologyUpdateOne {
	ptuo.mutation.ClearProject()
	return ptuo
}

// ClearTechnology clears the "technology" edge to the Technology entity.
func (ptuo *ProjectTechnologyUpdateOne) ClearTechnology() *ProjectTechnologyUpdateOne {
	ptuo.mutation.ClearTechnology()
	return ptuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ptuo *ProjectTechnologyUpdateOne) Select(field string, fields ...string) *ProjectTechnologyUpdateOne {
	ptuo.fields = append([]string{field}, fields...)
	return ptuo
}

// Save executes the query and returns the updated ProjectTechnology entity.
func (ptuo *ProjectTechnologyUpdateOne) Save(ctx context.Context) (*ProjectTechnology, error) {
	var (
		err  error
		node *ProjectTechnology
	)
	if len(ptuo.hooks) == 0 {
		if err = ptuo.check(); err != nil {
			return nil, err
		}
		node, err = ptuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectTechnologyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ptuo.check(); err != nil {
				return nil, err
			}
			ptuo.mutation = mutation
			node, err = ptuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ptuo.hooks) - 1; i >= 0; i-- {
			if ptuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ptuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ptuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (ptuo *ProjectTechnologyUpdateOne) SaveX(ctx context.Context) *ProjectTechnology {
	node, err := ptuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ptuo *ProjectTechnologyUpdateOne) Exec(ctx context.Context) error {
	_, err := ptuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptuo *ProjectTechnologyUpdateOne) ExecX(ctx context.Context) {
	if err := ptuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptuo *ProjectTechnologyUpdateOne) check() error {
	if v, ok := ptuo.mutation.GetType(); ok {
		if err := projecttechnology.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "ProjectTechnology.type": %w`, err)}
		}
	}
	if _, ok := ptuo.mutation.ProjectID(); ptuo.mutation.ProjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProjectTechnology.project"`)
	}
	if _, ok := ptuo.mutation.TechnologyID(); ptuo.mutation.TechnologyCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProjectTechnology.technology"`)
	}
	return nil
}

func (ptuo *ProjectTechnologyUpdateOne) sqlSave(ctx context.Context) (_node *ProjectTechnology, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   projecttechnology.Table,
			Columns: projecttechnology.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: projecttechnology.FieldID,
			},
		},
	}
	id, ok := ptuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProjectTechnology.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ptuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, projecttechnology.FieldID)
		for _, f := range fields {
			if !projecttechnology.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != projecttechnology.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ptuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ptuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: projecttechnology.FieldType,
		})
	}
	if ptuo.mutation.ProjectCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptuo.mutation.ProjectIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ptuo.mutation.TechnologyCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptuo.mutation.TechnologyIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProjectTechnology{config: ptuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ptuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projecttechnology.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
