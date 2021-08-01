package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/privacy"
	"github.com/fogo-sh/grackdb/ent/rules"
)

// DiscordAccount holds the schema definition for the DiscordAccount entity.
type DiscordAccount struct {
	ent.Schema
}

// Fields of the DiscordAccount.
func (DiscordAccount) Fields() []ent.Field {
	return []ent.Field{
		field.String("discord_id").
			Unique().
			Immutable().
			NotEmpty().
			Annotations(
				entgql.OrderField("DISCORD_ID"),
			),
		field.String("username").
			Annotations(
				entgql.OrderField("USERNAME"),
			),
		field.String("discriminator").
			MinLen(4).
			MaxLen(4).
			Annotations(
				entgql.OrderField("DISCRIMINATOR"),
			),
	}
}

// Edges of the DiscordAccount.
func (DiscordAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("discord_accounts").
			Annotations(entgql.Bind()).
			Required().
			Unique(),
	}
}

func (DiscordAccount) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rules.DenyIfNotAuthenticated(),
			privacy.AlwaysAllowRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}
