// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/discordbot"
	"github.com/fogo-sh/grackdb/ent/predicate"
)

// DiscordBotDelete is the builder for deleting a DiscordBot entity.
type DiscordBotDelete struct {
	config
	hooks    []Hook
	mutation *DiscordBotMutation
}

// Where adds a new predicate to the DiscordBotDelete builder.
func (dbd *DiscordBotDelete) Where(ps ...predicate.DiscordBot) *DiscordBotDelete {
	dbd.mutation.predicates = append(dbd.mutation.predicates, ps...)
	return dbd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dbd *DiscordBotDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dbd.hooks) == 0 {
		affected, err = dbd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiscordBotMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dbd.mutation = mutation
			affected, err = dbd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dbd.hooks) - 1; i >= 0; i-- {
			mut = dbd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dbd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (dbd *DiscordBotDelete) ExecX(ctx context.Context) int {
	n, err := dbd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dbd *DiscordBotDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: discordbot.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: discordbot.FieldID,
			},
		},
	}
	if ps := dbd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, dbd.driver, _spec)
}

// DiscordBotDeleteOne is the builder for deleting a single DiscordBot entity.
type DiscordBotDeleteOne struct {
	dbd *DiscordBotDelete
}

// Exec executes the deletion query.
func (dbdo *DiscordBotDeleteOne) Exec(ctx context.Context) error {
	n, err := dbdo.dbd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{discordbot.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dbdo *DiscordBotDeleteOne) ExecX(ctx context.Context) {
	dbdo.dbd.ExecX(ctx)
}