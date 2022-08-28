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

// TechnologyAssociation holds the schema definition for the TechnologyAssociation entity.
type TechnologyAssociation struct {
	ent.Schema
}

// Fields of the TechnologyAssociation.
func (TechnologyAssociation) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			NamedValues(
				"WrittenIn", "WRITTEN_IN",
				"Implements", "IMPLEMENTS",
				"Uses", "USES",
			).
			Annotations(entgql.OrderField("TYPE")),
	}
}

// Edges of the TechnologyAssociation.
func (TechnologyAssociation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent", Technology.Type).
			Ref("child_technologies").
			Annotations(entgql.Bind()).
			Required().
			Unique(),
		edge.From("child", Technology.Type).
			Ref("parent_technologies").
			Annotations(entgql.Bind()).
			Required().
			Unique(),
	}
}

func (TechnologyAssociation) Policy() ent.Policy {
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

func (TechnologyAssociation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField("technologyAssociations"),
	}
}
