package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/fogo-sh/grackdb/ent"
)

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return r.client.User.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithUserOrder(orderBy),
			ent.WithUserFilter(where.Filter),
		)
}

func (r *queryResolver) DiscordAccounts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.DiscordAccountOrder, where *ent.DiscordAccountWhereInput) (*ent.DiscordAccountConnection, error) {
	return r.client.DiscordAccount.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithDiscordAccountOrder(orderBy),
			ent.WithDiscordAccountFilter(where.Filter),
		)
}

func (r *queryResolver) GithubAccounts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.GithubAccountOrder, where *ent.GithubAccountWhereInput) (*ent.GithubAccountConnection, error) {
	return r.client.GithubAccount.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithGithubAccountOrder(orderBy),
			ent.WithGithubAccountFilter(where.Filter),
		)
}

func (r *queryResolver) GithubOrganizationMembers(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.GithubOrganizationMemberOrder, where *ent.GithubOrganizationMemberWhereInput) (*ent.GithubOrganizationMemberConnection, error) {
	return r.client.GithubOrganizationMember.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithGithubOrganizationMemberOrder(orderBy),
			ent.WithGithubOrganizationMemberFilter(where.Filter),
		)
}

func (r *queryResolver) GithubOrganizations(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.GithubOrganizationOrder, where *ent.GithubOrganizationWhereInput) (*ent.GithubOrganizationConnection, error) {
	return r.client.GithubOrganization.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithGithubOrganizationOrder(orderBy),
			ent.WithGithubOrganizationFilter(where.Filter),
		)
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
