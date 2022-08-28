package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"

	"github.com/fogo-sh/grackdb/ent/privacy"
	"github.com/fogo-sh/grackdb/ent/rules"
)

// DiscordBot holds the schema definition for the DiscordBot entity.
type DiscordBot struct {
	ent.Schema
}

// Fields of the DiscordBot.
func (DiscordBot) Fields() []ent.Field {
	return nil
}

// Edges of the DiscordBot.
func (DiscordBot) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("account", DiscordAccount.Type).
			Annotations(entgql.Bind()).
			Required().
			Unique(),
		edge.From("project", Project.Type).
			Ref("discord_bots").
			Annotations(entgql.Bind()).
			Required().
			Unique(),
		edge.From("repository", Repository.Type).
			Ref("discord_bots").
			Annotations(entgql.Bind()).
			Unique(),
	}
}

func (DiscordBot) Policy() ent.Policy {
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

func (DiscordBot) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField("discordBots"),
	}
}
