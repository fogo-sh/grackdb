package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/privacy"
	"github.com/fogo-sh/grackdb/ent/rules"
)

// ProjectAssociation holds the schema definition for the ProjectAssociation entity.
type ProjectAssociation struct {
	ent.Schema
}

// Fields of the ProjectAssociation.
func (ProjectAssociation) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			NamedValues(
				"BasedOff", "BASED_OFF",
				"Replaces", "REPLACES",
				"InspiredBy", "INSPIRED_BY",
				"Related", "RELATED",
			).
			Annotations(entgql.OrderField("TYPE")),
	}
}

// Edges of the ProjectAssociation.
func (ProjectAssociation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", Project.Type).
			Ref("child_projects").
			Annotations(entgql.Bind()).
			Required().
			Unique(),
		edge.From("child", Project.Type).
			Ref("parent_projects").
			Annotations(entgql.Bind()).
			Required().
			Unique(),
	}
}

func (ProjectAssociation) Policy() ent.Policy {
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
