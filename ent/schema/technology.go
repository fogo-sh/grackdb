package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/privacy"
	"github.com/fogo-sh/grackdb/ent/rules"
)

// Technology holds the schema definition for the Technology entity.
type Technology struct {
	ent.Schema
}

// Fields of the Technology.
func (Technology) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Annotations(entgql.OrderField("NAME")).
			NotEmpty().
			Unique(),
		field.String("description").
			Annotations(entgql.OrderField("DESCRIPTION")).
			Optional().
			Nillable(),
		field.String("colour").
			Annotations(entgql.OrderField("COLOUR")).
			Optional().
			Nillable(),
		field.Enum("type").
			NamedValues(
				"Library", "LIBRARY",
				"Language", "LANGUAGE",
				"Algorithm", "ALGORITHM",
				"Database", "DATABASE",
				"Protocol", "PROTOCOL",
				"Service", "SERVICE",
			).
			Annotations(entgql.OrderField("TYPE")),
	}
}

// Edges of the Technology.
func (Technology) Edges() []ent.Edge {
	return nil
}

func (Technology) Policy() ent.Policy {
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
