package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/fogo-sh/grackdb/ent/privacy"
	"github.com/fogo-sh/grackdb/ent/rules"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Annotations(entgql.OrderField("NAME")),
		field.String("description").
			Optional().
			Nillable().
			Annotations(entgql.OrderField("DESCRIPTION")),
		field.Time("start_date").
			Annotations(entgql.OrderField("START_DATE")),
		field.Time("end_date").
			Optional().
			Nillable().
			Annotations(entgql.OrderField("END_DATE")),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("contributors", ProjectContributor.Type).
			Annotations(entgql.Bind()),
		edge.To("parent_projects", ProjectAssociation.Type).
			Annotations(entgql.MapsTo("parentProjects")),
		edge.To("child_projects", ProjectAssociation.Type).
			Annotations(entgql.MapsTo("childProjects")),
		edge.To("repositories", Repository.Type).
			Annotations(entgql.Bind()),
		edge.To("discord_bots", DiscordBot.Type).
			Annotations(entgql.MapsTo("discordBots")),
		edge.To("sites", Site.Type).
			Annotations(entgql.Bind()),
	}
}

func (Project) Policy() ent.Policy {
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
