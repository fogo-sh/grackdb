package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/privacy"
	"github.com/fogo-sh/grackdb/ent/rules"
)

// Site holds the schema definition for the Site entity.
type Site struct {
	ent.Schema
}

// Fields of the Site.
func (Site) Fields() []ent.Field {
	return []ent.Field{
		field.String("url").
			Annotations(entgql.OrderField("URL")).
			NotEmpty(),
	}
}

// Edges of the Site.
func (Site) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("sites").
			Annotations(entgql.Bind()).
			Required().
			Unique(),
		edge.From("repository", Repository.Type).
			Ref("sites").
			Annotations(entgql.Bind()).
			Unique(),
	}
}

func (Site) Policy() ent.Policy {
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
