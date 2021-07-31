package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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
			Annotations(entgql.Bind()),
		edge.To("github_accounts", GithubAccount.Type).
			Annotations(entgql.Bind()),
	}
}
