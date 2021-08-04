package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/privacy"
	"github.com/fogo-sh/grackdb/ent/rules"
)

// ProjectTechnology holds the schema definition for the ProjectTechnology entity.
type ProjectTechnology struct {
	ent.Schema
}

// Fields of the ProjectTechnology.
func (ProjectTechnology) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			NamedValues(
				"WrittenIn", "WRITTEN_IN",
				"Implements", "IMPLEMENTS",
				"Uses", "USES",
				"Contains", "CONTAINS",
			).
			Annotations(entgql.OrderField("TYPE")),
	}
}

// Edges of the ProjectTechnology.
func (ProjectTechnology) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("technologies").
			Annotations(entgql.Bind()).
			Required().
			Unique(),
		edge.From("technology", Technology.Type).
			Ref("projects").
			Annotations(entgql.Bind()).
			Required().
			Unique(),
	}
}

func (ProjectTechnology) Policy() ent.Policy {
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
