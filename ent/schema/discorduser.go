package schema

import (
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
		field.String("id").
			Unique().
			Immutable().
			NotEmpty(),
		field.String("username"),
		field.String("discriminator").
			MinLen(4).
			MaxLen(4),
	}
}

// Edges of the DiscordAccount.
func (DiscordAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("discord_accounts").
			Required().
			Unique(),
	}
}
