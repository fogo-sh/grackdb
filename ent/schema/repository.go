package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/privacy"
	"github.com/fogo-sh/grackdb/ent/rules"
)

// Repository holds the schema definition for the Repository entity.
type Repository struct {
	ent.Schema
}

// Fields of the Repository.
func (Repository) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Annotations(entgql.OrderField("NAME")),
		field.String("description").
			Optional().
			Nillable().
			Annotations(entgql.OrderField("DESCRIPTION")),
	}
}

// Edges of the Repository.
func (Repository) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("repositories").
			Annotations(entgql.Bind()).
			Required().
			Unique(),
		edge.From("github_account", GithubAccount.Type).
			Ref("repositories").
			Annotations(entgql.MapsTo("githubAccount")).
			Unique(),
		edge.From("github_organization", GithubOrganization.Type).
			Ref("repositories").
			Annotations(entgql.MapsTo("githubOrganization")).
			Unique(),
		edge.To("discord_bots", DiscordBot.Type).
			Annotations(entgql.MapsTo("discordBots")),
		edge.To("sites", Site.Type).
			Annotations(entgql.Bind()),
	}
}

func (Repository) Policy() ent.Policy {
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
