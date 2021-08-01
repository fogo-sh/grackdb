package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/privacy"
	"github.com/fogo-sh/grackdb/ent/rules"
)

// ProjectContributor holds the schema definition for the ProjectContributor entity.
type ProjectContributor struct {
	ent.Schema
}

// Fields of the ProjectContributor.
func (ProjectContributor) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("role").
			NamedValues(
				"Owner", "OWNER",
				"Contributor", "CONTRIBUTOR",
			).
			Annotations(entgql.OrderField("ROLE")),
	}
}

// Edges of the ProjectContributor.
func (ProjectContributor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("contributors").
			Annotations(entgql.Bind()).
			Unique(),
		edge.From("user", User.Type).
			Ref("project_contributions").
			Annotations(entgql.Bind()).
			Unique(),
	}
}

func (ProjectContributor) Policy() ent.Policy {
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
