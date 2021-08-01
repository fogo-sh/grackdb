package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/privacy"
	"github.com/fogo-sh/grackdb/ent/rules"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Unique().
			Annotations(
				entgql.OrderField("USERNAME"),
			),
		field.String("avatar_url").
			Optional().
			Nillable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("discord_accounts", DiscordAccount.Type).
			Annotations(entgql.MapsTo("discordAccounts")),
		edge.To("github_accounts", GithubAccount.Type).
			Annotations(entgql.MapsTo("githubAccounts")),
		edge.To("project_contributions", ProjectContributor.Type).
			Annotations(entgql.MapsTo("projectContributions")),
	}
}

func (User) Policy() ent.Policy {
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
