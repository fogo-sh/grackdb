// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/githubaccount"
	"github.com/fogo-sh/grackdb/ent/githuborganization"
	"github.com/fogo-sh/grackdb/ent/githuborganizationmember"
	"github.com/fogo-sh/grackdb/ent/predicate"
)

// GithubOrganizationMemberUpdate is the builder for updating GithubOrganizationMember entities.
type GithubOrganizationMemberUpdate struct {
	config
	hooks    []Hook
	mutation *GithubOrganizationMemberMutation
}

// Where appends a list predicates to the GithubOrganizationMemberUpdate builder.
func (gomu *GithubOrganizationMemberUpdate) Where(ps ...predicate.GithubOrganizationMember) *GithubOrganizationMemberUpdate {
	gomu.mutation.Where(ps...)
	return gomu
}

// SetRole sets the "role" field.
func (gomu *GithubOrganizationMemberUpdate) SetRole(gi githuborganizationmember.Role) *GithubOrganizationMemberUpdate {
	gomu.mutation.SetRole(gi)
	return gomu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (gomu *GithubOrganizationMemberUpdate) SetNillableRole(gi *githuborganizationmember.Role) *GithubOrganizationMemberUpdate {
	if gi != nil {
		gomu.SetRole(*gi)
	}
	return gomu
}

// SetOrganizationID sets the "organization" edge to the GithubOrganization entity by ID.
func (gomu *GithubOrganizationMemberUpdate) SetOrganizationID(id int) *GithubOrganizationMemberUpdate {
	gomu.mutation.SetOrganizationID(id)
	return gomu
}

// SetOrganization sets the "organization" edge to the GithubOrganization entity.
func (gomu *GithubOrganizationMemberUpdate) SetOrganization(g *GithubOrganization) *GithubOrganizationMemberUpdate {
	return gomu.SetOrganizationID(g.ID)
}

// SetAccountID sets the "account" edge to the GithubAccount entity by ID.
func (gomu *GithubOrganizationMemberUpdate) SetAccountID(id int) *GithubOrganizationMemberUpdate {
	gomu.mutation.SetAccountID(id)
	return gomu
}

// SetAccount sets the "account" edge to the GithubAccount entity.
func (gomu *GithubOrganizationMemberUpdate) SetAccount(g *GithubAccount) *GithubOrganizationMemberUpdate {
	return gomu.SetAccountID(g.ID)
}

// Mutation returns the GithubOrganizationMemberMutation object of the builder.
func (gomu *GithubOrganizationMemberUpdate) Mutation() *GithubOrganizationMemberMutation {
	return gomu.mutation
}

// ClearOrganization clears the "organization" edge to the GithubOrganization entity.
func (gomu *GithubOrganizationMemberUpdate) ClearOrganization() *GithubOrganizationMemberUpdate {
	gomu.mutation.ClearOrganization()
	return gomu
}

// ClearAccount clears the "account" edge to the GithubAccount entity.
func (gomu *GithubOrganizationMemberUpdate) ClearAccount() *GithubOrganizationMemberUpdate {
	gomu.mutation.ClearAccount()
	return gomu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gomu *GithubOrganizationMemberUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gomu.hooks) == 0 {
		if err = gomu.check(); err != nil {
			return 0, err
		}
		affected, err = gomu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GithubOrganizationMemberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gomu.check(); err != nil {
				return 0, err
			}
			gomu.mutation = mutation
			affected, err = gomu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gomu.hooks) - 1; i >= 0; i-- {
			if gomu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gomu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gomu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gomu *GithubOrganizationMemberUpdate) SaveX(ctx context.Context) int {
	affected, err := gomu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gomu *GithubOrganizationMemberUpdate) Exec(ctx context.Context) error {
	_, err := gomu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gomu *GithubOrganizationMemberUpdate) ExecX(ctx context.Context) {
	if err := gomu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gomu *GithubOrganizationMemberUpdate) check() error {
	if v, ok := gomu.mutation.Role(); ok {
		if err := githuborganizationmember.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "GithubOrganizationMember.role": %w`, err)}
		}
	}
	if _, ok := gomu.mutation.OrganizationID(); gomu.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GithubOrganizationMember.organization"`)
	}
	if _, ok := gomu.mutation.AccountID(); gomu.mutation.AccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GithubOrganizationMember.account"`)
	}
	return nil
}

func (gomu *GithubOrganizationMemberUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   githuborganizationmember.Table,
			Columns: githuborganizationmember.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: githuborganizationmember.FieldID,
			},
		},
	}
	if ps := gomu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gomu.mutation.Role(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: githuborganizationmember.FieldRole,
		})
	}
	if gomu.mutation.OrganizationCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gomu.mutation.OrganizationIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gomu.mutation.AccountCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gomu.mutation.AccountIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gomu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{githuborganizationmember.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// GithubOrganizationMemberUpdateOne is the builder for updating a single GithubOrganizationMember entity.
type GithubOrganizationMemberUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GithubOrganizationMemberMutation
}

// SetRole sets the "role" field.
func (gomuo *GithubOrganizationMemberUpdateOne) SetRole(gi githuborganizationmember.Role) *GithubOrganizationMemberUpdateOne {
	gomuo.mutation.SetRole(gi)
	return gomuo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (gomuo *GithubOrganizationMemberUpdateOne) SetNillableRole(gi *githuborganizationmember.Role) *GithubOrganizationMemberUpdateOne {
	if gi != nil {
		gomuo.SetRole(*gi)
	}
	return gomuo
}

// SetOrganizationID sets the "organization" edge to the GithubOrganization entity by ID.
func (gomuo *GithubOrganizationMemberUpdateOne) SetOrganizationID(id int) *GithubOrganizationMemberUpdateOne {
	gomuo.mutation.SetOrganizationID(id)
	return gomuo
}

// SetOrganization sets the "organization" edge to the GithubOrganization entity.
func (gomuo *GithubOrganizationMemberUpdateOne) SetOrganization(g *GithubOrganization) *GithubOrganizationMemberUpdateOne {
	return gomuo.SetOrganizationID(g.ID)
}

// SetAccountID sets the "account" edge to the GithubAccount entity by ID.
func (gomuo *GithubOrganizationMemberUpdateOne) SetAccountID(id int) *GithubOrganizationMemberUpdateOne {
	gomuo.mutation.SetAccountID(id)
	return gomuo
}

// SetAccount sets the "account" edge to the GithubAccount entity.
func (gomuo *GithubOrganizationMemberUpdateOne) SetAccount(g *GithubAccount) *GithubOrganizationMemberUpdateOne {
	return gomuo.SetAccountID(g.ID)
}

// Mutation returns the GithubOrganizationMemberMutation object of the builder.
func (gomuo *GithubOrganizationMemberUpdateOne) Mutation() *GithubOrganizationMemberMutation {
	return gomuo.mutation
}

// ClearOrganization clears the "organization" edge to the GithubOrganization entity.
func (gomuo *GithubOrganizationMemberUpdateOne) ClearOrganization() *GithubOrganizationMemberUpdateOne {
	gomuo.mutation.ClearOrganization()
	return gomuo
}

// ClearAccount clears the "account" edge to the GithubAccount entity.
func (gomuo *GithubOrganizationMemberUpdateOne) ClearAccount() *GithubOrganizationMemberUpdateOne {
	gomuo.mutation.ClearAccount()
	return gomuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gomuo *GithubOrganizationMemberUpdateOne) Select(field string, fields ...string) *GithubOrganizationMemberUpdateOne {
	gomuo.fields = append([]string{field}, fields...)
	return gomuo
}

// Save executes the query and returns the updated GithubOrganizationMember entity.
func (gomuo *GithubOrganizationMemberUpdateOne) Save(ctx context.Context) (*GithubOrganizationMember, error) {
	var (
		err  error
		node *GithubOrganizationMember
	)
	if len(gomuo.hooks) == 0 {
		if err = gomuo.check(); err != nil {
			return nil, err
		}
		node, err = gomuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GithubOrganizationMemberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gomuo.check(); err != nil {
				return nil, err
			}
			gomuo.mutation = mutation
			node, err = gomuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gomuo.hooks) - 1; i >= 0; i-- {
			if gomuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gomuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, gomuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*GithubOrganizationMember)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GithubOrganizationMemberMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (gomuo *GithubOrganizationMemberUpdateOne) SaveX(ctx context.Context) *GithubOrganizationMember {
	node, err := gomuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gomuo *GithubOrganizationMemberUpdateOne) Exec(ctx context.Context) error {
	_, err := gomuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gomuo *GithubOrganizationMemberUpdateOne) ExecX(ctx context.Context) {
	if err := gomuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gomuo *GithubOrganizationMemberUpdateOne) check() error {
	if v, ok := gomuo.mutation.Role(); ok {
		if err := githuborganizationmember.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "GithubOrganizationMember.role": %w`, err)}
		}
	}
	if _, ok := gomuo.mutation.OrganizationID(); gomuo.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GithubOrganizationMember.organization"`)
	}
	if _, ok := gomuo.mutation.AccountID(); gomuo.mutation.AccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GithubOrganizationMember.account"`)
	}
	return nil
}

func (gomuo *GithubOrganizationMemberUpdateOne) sqlSave(ctx context.Context) (_node *GithubOrganizationMember, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   githuborganizationmember.Table,
			Columns: githuborganizationmember.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: githuborganizationmember.FieldID,
			},
		},
	}
	id, ok := gomuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GithubOrganizationMember.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gomuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, githuborganizationmember.FieldID)
		for _, f := range fields {
			if !githuborganizationmember.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != githuborganizationmember.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gomuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gomuo.mutation.Role(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: githuborganizationmember.FieldRole,
		})
	}
	if gomuo.mutation.OrganizationCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gomuo.mutation.OrganizationIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gomuo.mutation.AccountCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gomuo.mutation.AccountIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &GithubOrganizationMember{config: gomuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gomuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{githuborganizationmember.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
