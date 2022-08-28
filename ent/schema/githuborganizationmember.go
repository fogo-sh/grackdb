package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/fogo-sh/grackdb/ent/privacy"
	"github.com/fogo-sh/grackdb/ent/rules"
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
			Ref("members").
			Annotations(entgql.Bind()).
			Required().
			Unique(),
		edge.From("account", GithubAccount.Type).
			Ref("organization_memberships").
			Annotations(entgql.Bind()).
			Required().
			Unique(),
	}
}

func (GithubOrganizationMember) Policy() ent.Policy {
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

func (GithubOrganizationMember) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField("githubOrganizationMembers"),
	}
}
