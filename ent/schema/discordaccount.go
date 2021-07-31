package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DiscordAccount holds the schema definition for the DiscordAccount entity.
type DiscordAccount struct {
	ent.Schema
}

// Fields of the DiscordAccount.
func (DiscordAccount) Fields() []ent.Field {
	return []ent.Field{
		field.String("discord_id").
			Unique().
			Immutable().
			NotEmpty().
			Annotations(
				entgql.OrderField("DISCORD_ID"),
			),
		field.String("username").
			Annotations(
				entgql.OrderField("USERNAME"),
			),
		field.String("discriminator").
			MinLen(4).
			MaxLen(4).
			Annotations(
				entgql.OrderField("DISCRIMINATOR"),
			),
	}
}

// Edges of the DiscordAccount.
func (DiscordAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Annotations(entgql.Bind()).
			Ref("discord_accounts").
			Required().
			Unique(),
	}
}
