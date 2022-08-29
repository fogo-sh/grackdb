// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/discordaccount"
	"github.com/fogo-sh/grackdb/ent/predicate"
)

// DiscordAccountDelete is the builder for deleting a DiscordAccount entity.
type DiscordAccountDelete struct {
	config
	hooks    []Hook
	mutation *DiscordAccountMutation
}

// Where appends a list predicates to the DiscordAccountDelete builder.
func (dad *DiscordAccountDelete) Where(ps ...predicate.DiscordAccount) *DiscordAccountDelete {
	dad.mutation.Where(ps...)
	return dad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dad *DiscordAccountDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dad.hooks) == 0 {
		affected, err = dad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiscordAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dad.mutation = mutation
			affected, err = dad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dad.hooks) - 1; i >= 0; i-- {
			if dad.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (dad *DiscordAccountDelete) ExecX(ctx context.Context) int {
	n, err := dad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dad *DiscordAccountDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: discordaccount.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: discordaccount.FieldID,
			},
		},
	}
	if ps := dad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// DiscordAccountDeleteOne is the builder for deleting a single DiscordAccount entity.
type DiscordAccountDeleteOne struct {
	dad *DiscordAccountDelete
}

// Exec executes the deletion query.
func (dado *DiscordAccountDeleteOne) Exec(ctx context.Context) error {
	n, err := dado.dad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{discordaccount.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dado *DiscordAccountDeleteOne) ExecX(ctx context.Context) {
	dado.dad.ExecX(ctx)
}
