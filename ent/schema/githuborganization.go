package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/privacy"
	"github.com/fogo-sh/grackdb/ent/rules"
)

// GithubOrganization holds the schema definition for the GithubOrganization entity.
type GithubOrganization struct {
	ent.Schema
}

// Fields of the GithubOrganization.
func (GithubOrganization) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			NotEmpty().
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.String("display_name").
			Optional().
			Annotations(
				entgql.OrderField("DISPLAY_NAME"),
			),
	}
}

// Edges of the GithubOrganization.
func (GithubOrganization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("members", GithubOrganizationMember.Type).
			Annotations(entgql.Bind()),
		edge.To("repositories", Repository.Type).
			Annotations(entgql.Bind()),
	}
}

func (GithubOrganization) Policy() ent.Policy {
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
