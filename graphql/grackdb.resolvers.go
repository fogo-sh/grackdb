package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/fogo-sh/grackdb"
	"github.com/fogo-sh/grackdb/ent"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	return ent.FromContext(ctx).User.
		Create().
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) CreateDiscordAccount(ctx context.Context, input ent.CreateDiscordAccountInput) (*ent.DiscordAccount, error) {
	return ent.FromContext(ctx).DiscordAccount.
		Create().
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) CreateDiscordBot(ctx context.Context, input ent.CreateDiscordBotInput) (*ent.DiscordBot, error) {
	return ent.FromContext(ctx).DiscordBot.
		Create().
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) CreateGithubAccount(ctx context.Context, input ent.CreateGithubAccountInput) (*ent.GithubAccount, error) {
	return ent.FromContext(ctx).GithubAccount.
		Create().
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) CreateGithubOrganization(ctx context.Context, input ent.CreateGithubOrganizationInput) (*ent.GithubOrganization, error) {
	return ent.FromContext(ctx).GithubOrganization.
		Create().
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) CreateGithubOrganizationMember(ctx context.Context, input ent.CreateGithubOrganizationMemberInput) (*ent.GithubOrganizationMember, error) {
	return ent.FromContext(ctx).GithubOrganizationMember.
		Create().
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) CreateProject(ctx context.Context, input ent.CreateProjectInput) (*ent.Project, error) {
	return ent.FromContext(ctx).Project.
		Create().
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) CreateProjectContributor(ctx context.Context, input ent.CreateProjectContributorInput) (*ent.ProjectContributor, error) {
	return ent.FromContext(ctx).ProjectContributor.
		Create().
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) CreateProjectAssociation(ctx context.Context, input ent.CreateProjectAssociationInput) (*ent.ProjectAssociation, error) {
	return ent.FromContext(ctx).ProjectAssociation.
		Create().
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) CreateRepository(ctx context.Context, input ent.CreateRepositoryInput) (*ent.Repository, error) {
	return ent.FromContext(ctx).Repository.
		Create().
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) CreateSite(ctx context.Context, input ent.CreateSiteInput) (*ent.Site, error) {
	return ent.FromContext(ctx).Site.
		Create().
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) CreateTechnology(ctx context.Context, input ent.CreateTechnologyInput) (*ent.Technology, error) {
	return ent.FromContext(ctx).Technology.
		Create().
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) CreateTechnologyAssociation(ctx context.Context, input ent.CreateTechnologyAssociationInput) (*ent.TechnologyAssociation, error) {
	return ent.FromContext(ctx).TechnologyAssociation.
		Create().
		SetInput(input).
		Save(ctx)
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

func (r *queryResolver) DiscordBots(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.DiscordBotWhereInput) (*ent.DiscordBotConnection, error) {
	return r.client.DiscordBot.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithDiscordBotFilter(where.Filter),
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

func (r *queryResolver) Projects(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ProjectOrder, where *ent.ProjectWhereInput) (*ent.ProjectConnection, error) {
	return r.client.Project.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithProjectOrder(orderBy),
			ent.WithProjectFilter(where.Filter),
		)
}

func (r *queryResolver) ProjectContributors(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ProjectContributorOrder, where *ent.ProjectContributorWhereInput) (*ent.ProjectContributorConnection, error) {
	return r.client.ProjectContributor.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithProjectContributorOrder(orderBy),
			ent.WithProjectContributorFilter(where.Filter),
		)
}

func (r *queryResolver) ProjectAssociations(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ProjectAssociationOrder, where *ent.ProjectAssociationWhereInput) (*ent.ProjectAssociationConnection, error) {
	return r.client.ProjectAssociation.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithProjectAssociationOrder(orderBy),
			ent.WithProjectAssociationFilter(where.Filter),
		)
}

func (r *queryResolver) Repositories(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.RepositoryOrder, where *ent.RepositoryWhereInput) (*ent.RepositoryConnection, error) {
	return r.client.Repository.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithRepositoryOrder(orderBy),
			ent.WithRepositoryFilter(where.Filter),
		)
}

func (r *queryResolver) Sites(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.SiteOrder, where *ent.SiteWhereInput) (*ent.SiteConnection, error) {
	return r.client.Site.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithSiteOrder(orderBy),
			ent.WithSiteFilter(where.Filter),
		)
}

func (r *queryResolver) Technologies(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TechnologyOrder, where *ent.TechnologyWhereInput) (*ent.TechnologyConnection, error) {
	return r.client.Technology.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithTechnologyOrder(orderBy),
			ent.WithTechnologyFilter(where.Filter),
		)
}

func (r *queryResolver) TechnologyAssociations(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TechnologyAssociationOrder, where *ent.TechnologyAssociationWhereInput) (*ent.TechnologyAssociationConnection, error) {
	return r.client.TechnologyAssociation.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithTechnologyAssociationOrder(orderBy),
			ent.WithTechnologyAssociationFilter(where.Filter),
		)
}

func (r *queryResolver) AvailableAuthProviders(ctx context.Context) ([]*AuthProvider, error) {
	availableAuthProviders := []*AuthProvider{}

	if grackdb.AppConfig.DiscordClientId != "" {
		availableAuthProviders = append(availableAuthProviders, &AuthProvider{
			Type: AuthProviderTypeDiscord,
			URL:  "/oauth/discord/auth",
		})
	}

	if grackdb.AppConfig.GithubClientId != "" {
		availableAuthProviders = append(availableAuthProviders, &AuthProvider{
			Type: AuthProviderTypeGithub,
			URL:  "/oauth/github/auth",
		})
	}

	return availableAuthProviders, nil
}

func (r *queryResolver) CurrentUser(ctx context.Context) (*ent.User, error) {
	ginCtx, err := grackdb.GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("error fetching context: %w", err)
	}
	return ginCtx.Value("user").(*ent.User), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
