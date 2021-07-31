package schema

import (
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
			NotEmpty(),
	}
}

// Edges of the GithubAccount.
func (GithubAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("github_accounts").
			Required().
			Unique(),
		edge.To("organization_memberships", GithubOrganizationMember.Type),
	}
}
