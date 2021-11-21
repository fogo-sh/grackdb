// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/discordaccount"
	"github.com/fogo-sh/grackdb/ent/discordbot"
	"github.com/fogo-sh/grackdb/ent/user"
)

// DiscordAccountCreate is the builder for creating a DiscordAccount entity.
type DiscordAccountCreate struct {
	config
	mutation *DiscordAccountMutation
	hooks    []Hook
}

// SetDiscordID sets the "discord_id" field.
func (dac *DiscordAccountCreate) SetDiscordID(s string) *DiscordAccountCreate {
	dac.mutation.SetDiscordID(s)
	return dac
}

// SetUsername sets the "username" field.
func (dac *DiscordAccountCreate) SetUsername(s string) *DiscordAccountCreate {
	dac.mutation.SetUsername(s)
	return dac
}

// SetDiscriminator sets the "discriminator" field.
func (dac *DiscordAccountCreate) SetDiscriminator(s string) *DiscordAccountCreate {
	dac.mutation.SetDiscriminator(s)
	return dac
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (dac *DiscordAccountCreate) SetOwnerID(id int) *DiscordAccountCreate {
	dac.mutation.SetOwnerID(id)
	return dac
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (dac *DiscordAccountCreate) SetNillableOwnerID(id *int) *DiscordAccountCreate {
	if id != nil {
		dac = dac.SetOwnerID(*id)
	}
	return dac
}

// SetOwner sets the "owner" edge to the User entity.
func (dac *DiscordAccountCreate) SetOwner(u *User) *DiscordAccountCreate {
	return dac.SetOwnerID(u.ID)
}

// SetBotID sets the "bot" edge to the DiscordBot entity by ID.
func (dac *DiscordAccountCreate) SetBotID(id int) *DiscordAccountCreate {
	dac.mutation.SetBotID(id)
	return dac
}

// SetNillableBotID sets the "bot" edge to the DiscordBot entity by ID if the given value is not nil.
func (dac *DiscordAccountCreate) SetNillableBotID(id *int) *DiscordAccountCreate {
	if id != nil {
		dac = dac.SetBotID(*id)
	}
	return dac
}

// SetBot sets the "bot" edge to the DiscordBot entity.
func (dac *DiscordAccountCreate) SetBot(d *DiscordBot) *DiscordAccountCreate {
	return dac.SetBotID(d.ID)
}

// Mutation returns the DiscordAccountMutation object of the builder.
func (dac *DiscordAccountCreate) Mutation() *DiscordAccountMutation {
	return dac.mutation
}

// Save creates the DiscordAccount in the database.
func (dac *DiscordAccountCreate) Save(ctx context.Context) (*DiscordAccount, error) {
	var (
		err  error
		node *DiscordAccount
	)
	if len(dac.hooks) == 0 {
		if err = dac.check(); err != nil {
			return nil, err
		}
		node, err = dac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiscordAccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dac.check(); err != nil {
				return nil, err
			}
			dac.mutation = mutation
			if node, err = dac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dac.hooks) - 1; i >= 0; i-- {
			if dac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dac *DiscordAccountCreate) SaveX(ctx context.Context) *DiscordAccount {
	v, err := dac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dac *DiscordAccountCreate) Exec(ctx context.Context) error {
	_, err := dac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dac *DiscordAccountCreate) ExecX(ctx context.Context) {
	if err := dac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dac *DiscordAccountCreate) check() error {
	if _, ok := dac.mutation.DiscordID(); !ok {
		return &ValidationError{Name: "discord_id", err: errors.New(`ent: missing required field "discord_id"`)}
	}
	if v, ok := dac.mutation.DiscordID(); ok {
		if err := discordaccount.DiscordIDValidator(v); err != nil {
			return &ValidationError{Name: "discord_id", err: fmt.Errorf(`ent: validator failed for field "discord_id": %w`, err)}
		}
	}
	if _, ok := dac.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "username"`)}
	}
	if _, ok := dac.mutation.Discriminator(); !ok {
		return &ValidationError{Name: "discriminator", err: errors.New(`ent: missing required field "discriminator"`)}
	}
	if v, ok := dac.mutation.Discriminator(); ok {
		if err := discordaccount.DiscriminatorValidator(v); err != nil {
			return &ValidationError{Name: "discriminator", err: fmt.Errorf(`ent: validator failed for field "discriminator": %w`, err)}
		}
	}
	return nil
}

func (dac *DiscordAccountCreate) sqlSave(ctx context.Context) (*DiscordAccount, error) {
	_node, _spec := dac.createSpec()
	if err := sqlgraph.CreateNode(ctx, dac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (dac *DiscordAccountCreate) createSpec() (*DiscordAccount, *sqlgraph.CreateSpec) {
	var (
		_node = &DiscordAccount{config: dac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: discordaccount.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: discordaccount.FieldID,
			},
		}
	)
	if value, ok := dac.mutation.DiscordID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: discordaccount.FieldDiscordID,
		})
		_node.DiscordID = value
	}
	if value, ok := dac.mutation.Username(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: discordaccount.FieldUsername,
		})
		_node.Username = value
	}
	if value, ok := dac.mutation.Discriminator(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: discordaccount.FieldDiscriminator,
		})
		_node.Discriminator = value
	}
	if nodes := dac.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   discordaccount.OwnerTable,
			Columns: []string{discordaccount.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_discord_accounts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dac.mutation.BotIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   discordaccount.BotTable,
			Columns: []string{discordaccount.BotColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: discordbot.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.discord_bot_account = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DiscordAccountCreateBulk is the builder for creating many DiscordAccount entities in bulk.
type DiscordAccountCreateBulk struct {
	config
	builders []*DiscordAccountCreate
}

// Save creates the DiscordAccount entities in the database.
func (dacb *DiscordAccountCreateBulk) Save(ctx context.Context) ([]*DiscordAccount, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dacb.builders))
	nodes := make([]*DiscordAccount, len(dacb.builders))
	mutators := make([]Mutator, len(dacb.builders))
	for i := range dacb.builders {
		func(i int, root context.Context) {
			builder := dacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DiscordAccountMutation)
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
					_, err = mutators[i+1].Mutate(root, dacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dacb *DiscordAccountCreateBulk) SaveX(ctx context.Context) []*DiscordAccount {
	v, err := dacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dacb *DiscordAccountCreateBulk) Exec(ctx context.Context) error {
	_, err := dacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dacb *DiscordAccountCreateBulk) ExecX(ctx context.Context) {
	if err := dacb.Exec(ctx); err != nil {
		panic(err)
	}
}
