package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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
			NotEmpty(),
		field.String("display_name").
			Optional(),
	}
}

// Edges of the GithubOrganization.
func (GithubOrganization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("members", GithubOrganizationMember.Type),
	}
}
