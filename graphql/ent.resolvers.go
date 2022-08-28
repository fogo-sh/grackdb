package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/fogo-sh/grackdb/ent"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Nodes - nodes"))
}

// DiscordAccounts is the resolver for the discordAccounts field.
func (r *queryResolver) DiscordAccounts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.DiscordAccountOrder, where *ent.DiscordAccountWhereInput) (*ent.DiscordAccountConnection, error) {
	return r.client.DiscordAccount.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithDiscordAccountOrder(orderBy),
			ent.WithDiscordAccountFilter(where.Filter),
		)
}

// DiscordBots is the resolver for the discordBots field.
func (r *queryResolver) DiscordBots(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.DiscordBotWhereInput) (*ent.DiscordBotConnection, error) {
	return r.client.DiscordBot.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithDiscordBotFilter(where.Filter),
		)
}

// GithubAccounts is the resolver for the githubAccounts field.
func (r *queryResolver) GithubAccounts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.GithubAccountOrder, where *ent.GithubAccountWhereInput) (*ent.GithubAccountConnection, error) {
	return r.client.GithubAccount.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithGithubAccountOrder(orderBy),
			ent.WithGithubAccountFilter(where.Filter),
		)
}

// GithubOrganizations is the resolver for the githubOrganizations field.
func (r *queryResolver) GithubOrganizations(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.GithubOrganizationOrder, where *ent.GithubOrganizationWhereInput) (*ent.GithubOrganizationConnection, error) {
	return r.client.GithubOrganization.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithGithubOrganizationOrder(orderBy),
			ent.WithGithubOrganizationFilter(where.Filter),
		)
}

// GithubOrganizationMembers is the resolver for the githubOrganizationMembers field.
func (r *queryResolver) GithubOrganizationMembers(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.GithubOrganizationMemberOrder, where *ent.GithubOrganizationMemberWhereInput) (*ent.GithubOrganizationMemberConnection, error) {
	return r.client.GithubOrganizationMember.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithGithubOrganizationMemberOrder(orderBy),
			ent.WithGithubOrganizationMemberFilter(where.Filter),
		)
}

// Projects is the resolver for the projects field.
func (r *queryResolver) Projects(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ProjectOrder, where *ent.ProjectWhereInput) (*ent.ProjectConnection, error) {
	return r.client.Project.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithProjectOrder(orderBy),
			ent.WithProjectFilter(where.Filter),
		)
}

// ProjectAssociations is the resolver for the projectAssociations field.
func (r *queryResolver) ProjectAssociations(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ProjectAssociationOrder, where *ent.ProjectAssociationWhereInput) (*ent.ProjectAssociationConnection, error) {
	return r.client.ProjectAssociation.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithProjectAssociationOrder(orderBy),
			ent.WithProjectAssociationFilter(where.Filter),
		)
}

// ProjectContributors is the resolver for the projectContributors field.
func (r *queryResolver) ProjectContributors(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ProjectContributorOrder, where *ent.ProjectContributorWhereInput) (*ent.ProjectContributorConnection, error) {
	return r.client.ProjectContributor.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithProjectContributorOrder(orderBy),
			ent.WithProjectContributorFilter(where.Filter),
		)
}

// ProjectTechnologies is the resolver for the projectTechnologies field.
func (r *queryResolver) ProjectTechnologies(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ProjectTechnologyOrder, where *ent.ProjectTechnologyWhereInput) (*ent.ProjectTechnologyConnection, error) {
	return r.client.ProjectTechnology.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithProjectTechnologyOrder(orderBy),
			ent.WithProjectTechnologyFilter(where.Filter),
		)
}

// Repositories is the resolver for the repositories field.
func (r *queryResolver) Repositories(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.RepositoryOrder, where *ent.RepositoryWhereInput) (*ent.RepositoryConnection, error) {
	return r.client.Repository.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithRepositoryOrder(orderBy),
			ent.WithRepositoryFilter(where.Filter),
		)
}

// RepositoryTechnologies is the resolver for the repositoryTechnologies field.
func (r *queryResolver) RepositoryTechnologies(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.RepositoryTechnologyOrder, where *ent.RepositoryTechnologyWhereInput) (*ent.RepositoryTechnologyConnection, error) {
	return r.client.RepositoryTechnology.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithRepositoryTechnologyOrder(orderBy),
			ent.WithRepositoryTechnologyFilter(where.Filter),
		)
}

// Sites is the resolver for the sites field.
func (r *queryResolver) Sites(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.SiteOrder, where *ent.SiteWhereInput) (*ent.SiteConnection, error) {
	return r.client.Site.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithSiteOrder(orderBy),
			ent.WithSiteFilter(where.Filter),
		)
}

// Technologies is the resolver for the technologies field.
func (r *queryResolver) Technologies(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TechnologyOrder, where *ent.TechnologyWhereInput) (*ent.TechnologyConnection, error) {
	return r.client.Technology.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithTechnologyOrder(orderBy),
			ent.WithTechnologyFilter(where.Filter),
		)
}

// TechnologyAssociations is the resolver for the technologyAssociations field.
func (r *queryResolver) TechnologyAssociations(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TechnologyAssociationOrder, where *ent.TechnologyAssociationWhereInput) (*ent.TechnologyAssociationConnection, error) {
	return r.client.TechnologyAssociation.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithTechnologyAssociationOrder(orderBy),
			ent.WithTechnologyAssociationFilter(where.Filter),
		)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return r.client.User.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithUserOrder(orderBy),
			ent.WithUserFilter(where.Filter),
		)
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
