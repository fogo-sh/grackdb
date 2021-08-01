package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GithubAccount holds the schema definition for the GithubAccount entity.
type GithubAccount struct {
	ent.Schema
}

// Fields of the GithubAccount.
func (GithubAccount) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Unique().
			NotEmpty().
			Annotations(
				entgql.OrderField("USERNAME"),
			),
	}
}

// Edges of the GithubAccount.
func (GithubAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("github_accounts").
			Annotations(entgql.Bind()).
			Required().
			Unique(),
		edge.To("organization_memberships", GithubOrganizationMember.Type).
			Annotations(entgql.MapsTo("organizationMemberships")),
	}
}
