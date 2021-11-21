// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/githuborganizationmember"
	"github.com/fogo-sh/grackdb/ent/predicate"
)

// GithubOrganizationMemberDelete is the builder for deleting a GithubOrganizationMember entity.
type GithubOrganizationMemberDelete struct {
	config
	hooks    []Hook
	mutation *GithubOrganizationMemberMutation
}

// Where appends a list predicates to the GithubOrganizationMemberDelete builder.
func (gomd *GithubOrganizationMemberDelete) Where(ps ...predicate.GithubOrganizationMember) *GithubOrganizationMemberDelete {
	gomd.mutation.Where(ps...)
	return gomd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gomd *GithubOrganizationMemberDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gomd.hooks) == 0 {
		affected, err = gomd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GithubOrganizationMemberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gomd.mutation = mutation
			affected, err = gomd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gomd.hooks) - 1; i >= 0; i-- {
			if gomd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gomd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gomd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (gomd *GithubOrganizationMemberDelete) ExecX(ctx context.Context) int {
	n, err := gomd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gomd *GithubOrganizationMemberDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: githuborganizationmember.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: githuborganizationmember.FieldID,
			},
		},
	}
	if ps := gomd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, gomd.driver, _spec)
}

// GithubOrganizationMemberDeleteOne is the builder for deleting a single GithubOrganizationMember entity.
type GithubOrganizationMemberDeleteOne struct {
	gomd *GithubOrganizationMemberDelete
}

// Exec executes the deletion query.
func (gomdo *GithubOrganizationMemberDeleteOne) Exec(ctx context.Context) error {
	n, err := gomdo.gomd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{githuborganizationmember.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gomdo *GithubOrganizationMemberDeleteOne) ExecX(ctx context.Context) {
	gomdo.gomd.ExecX(ctx)
}
