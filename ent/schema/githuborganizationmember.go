package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GithubOrganizationMember holds the schema definition for the GithubOrganizationMember entity.
type GithubOrganizationMember struct {
	ent.Schema
}

// Fields of the GithubOrganizationMember.
func (GithubOrganizationMember) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("role").
			NamedValues(
				"Admin", "ADMIN",
				"Member", "MEMBER",
			).
			Default("MEMBER").
			Annotations(
				entgql.OrderField("ROLE"),
			),
	}
}

// Edges of the GithubOrganizationMember.
func (GithubOrganizationMember) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organization", GithubOrganization.Type).
			Annotations(entgql.Bind()).
			Ref("members").
			Unique(),
		edge.From("account", GithubAccount.Type).
			Annotations(entgql.Bind()).
			Ref("organization_memberships").
			Unique(),
	}
}
