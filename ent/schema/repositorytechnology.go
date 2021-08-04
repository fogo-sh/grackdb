package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/privacy"
	"github.com/fogo-sh/grackdb/ent/rules"
)

// RepositoryTechnology holds the schema definition for the RepositoryTechnology entity.
type RepositoryTechnology struct {
	ent.Schema
}

// Fields of the RepositoryTechnology.
func (RepositoryTechnology) Fields() []ent.Field {
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

// Edges of the RepositoryTechnology.
func (RepositoryTechnology) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("repository", Repository.Type).
			Ref("technologies").
			Annotations(entgql.Bind()).
			Required().
			Unique(),
		edge.From("technology", Technology.Type).
			Ref("repositories").
			Annotations(entgql.Bind()).
			Required().
			Unique(),
	}
}

func (RepositoryTechnology) Policy() ent.Policy {
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
