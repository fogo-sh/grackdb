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

func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input ent.UpdateUserInput) (*ent.User, error) {
	return ent.FromContext(ctx).User.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*ent.User, error) {
	user, err := ent.FromContext(ctx).User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, ent.FromContext(ctx).User.
		DeleteOne(user).
		Exec(ctx)
}

func (r *mutationResolver) CreateDiscordAccount(ctx context.Context, input ent.CreateDiscordAccountInput) (*ent.DiscordAccount, error) {
	return ent.FromContext(ctx).DiscordAccount.
		Create().
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) UpdateDiscordAccount(ctx context.Context, id int, input ent.UpdateDiscordAccountInput) (*ent.DiscordAccount, error) {
	return ent.FromContext(ctx).DiscordAccount.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) DeleteDiscordAccount(ctx context.Context, id int) (*ent.DiscordAccount, error) {
	discordAccount, err := ent.FromContext(ctx).DiscordAccount.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return discordAccount, ent.FromContext(ctx).DiscordAccount.
		DeleteOne(discordAccount).
		Exec(ctx)
}

func (r *mutationResolver) CreateDiscordBot(ctx context.Context, input ent.CreateDiscordBotInput) (*ent.DiscordBot, error) {
	return ent.FromContext(ctx).DiscordBot.
		Create().
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) UpdateDiscordBot(ctx context.Context, id int, input ent.UpdateDiscordBotInput) (*ent.DiscordBot, error) {
	return ent.FromContext(ctx).DiscordBot.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) DeleteDiscordBot(ctx context.Context, id int) (*ent.DiscordBot, error) {
	discordBot, err := ent.FromContext(ctx).DiscordBot.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return discordBot, ent.FromContext(ctx).DiscordBot.
		DeleteOne(discordBot).
		Exec(ctx)
}

func (r *mutationResolver) CreateGithubAccount(ctx context.Context, input ent.CreateGithubAccountInput) (*ent.GithubAccount, error) {
	return ent.FromContext(ctx).GithubAccount.
		Create().
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) UpdateGithubAccount(ctx context.Context, id int, input ent.UpdateGithubAccountInput) (*ent.GithubAccount, error) {
	return ent.FromContext(ctx).GithubAccount.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) DeleteGithubAccount(ctx context.Context, id int) (*ent.GithubAccount, error) {
	githubAccount, err := ent.FromContext(ctx).GithubAccount.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return githubAccount, ent.FromContext(ctx).GithubAccount.
		DeleteOne(githubAccount).
		Exec(ctx)
}

func (r *mutationResolver) CreateGithubOrganization(ctx context.Context, input ent.CreateGithubOrganizationInput) (*ent.GithubOrganization, error) {
	return ent.FromContext(ctx).GithubOrganization.
		Create().
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) UpdateGithubOrganization(ctx context.Context, id int, input ent.UpdateGithubOrganizationInput) (*ent.GithubOrganization, error) {
	return ent.FromContext(ctx).GithubOrganization.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) DeleteGithubOrganization(ctx context.Context, id int) (*ent.GithubOrganization, error) {
	githubOrganization, err := ent.FromContext(ctx).GithubOrganization.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return githubOrganization, ent.FromContext(ctx).GithubOrganization.
		DeleteOne(githubOrganization).
		Exec(ctx)
}

func (r *mutationResolver) CreateGithubOrganizationMember(ctx context.Context, input ent.CreateGithubOrganizationMemberInput) (*ent.GithubOrganizationMember, error) {
	return ent.FromContext(ctx).GithubOrganizationMember.
		Create().
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) UpdateGithubOrganizationMember(ctx context.Context, id int, input ent.UpdateGithubOrganizationMemberInput) (*ent.GithubOrganizationMember, error) {
	return ent.FromContext(ctx).GithubOrganizationMember.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) DeleteGithubOrganizationMember(ctx context.Context, id int) (*ent.GithubOrganizationMember, error) {
	githubOrganizationMember, err := ent.FromContext(ctx).GithubOrganizationMember.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return githubOrganizationMember, ent.FromContext(ctx).GithubOrganizationMember.
		DeleteOne(githubOrganizationMember).
		Exec(ctx)
}

func (r *mutationResolver) CreateProject(ctx context.Context, input ent.CreateProjectInput) (*ent.Project, error) {
	return ent.FromContext(ctx).Project.
		Create().
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) UpdateProject(ctx context.Context, id int, input ent.UpdateProjectInput) (*ent.Project, error) {
	return ent.FromContext(ctx).Project.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) DeleteProject(ctx context.Context, id int) (*ent.Project, error) {
	project, err := ent.FromContext(ctx).Project.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return project, ent.FromContext(ctx).Project.
		DeleteOne(project).
		Exec(ctx)
}

func (r *mutationResolver) CreateProjectContributor(ctx context.Context, input ent.CreateProjectContributorInput) (*ent.ProjectContributor, error) {
	return ent.FromContext(ctx).ProjectContributor.
		Create().
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) UpdateProjectContributor(ctx context.Context, id int, input ent.UpdateProjectContributorInput) (*ent.ProjectContributor, error) {
	return ent.FromContext(ctx).ProjectContributor.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) DeleteProjectContributor(ctx context.Context, id int) (*ent.ProjectContributor, error) {
	projectContributor, err := ent.FromContext(ctx).ProjectContributor.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return projectContributor, ent.FromContext(ctx).ProjectContributor.
		DeleteOne(projectContributor).
		Exec(ctx)
}

func (r *mutationResolver) CreateProjectAssociation(ctx context.Context, input ent.CreateProjectAssociationInput) (*ent.ProjectAssociation, error) {
	return ent.FromContext(ctx).ProjectAssociation.
		Create().
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) UpdateProjectAssociation(ctx context.Context, id int, input ent.UpdateProjectAssociationInput) (*ent.ProjectAssociation, error) {
	return ent.FromContext(ctx).ProjectAssociation.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) DeleteProjectAssociation(ctx context.Context, id int) (*ent.ProjectAssociation, error) {
	projectAssociation, err := ent.FromContext(ctx).ProjectAssociation.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return projectAssociation, ent.FromContext(ctx).ProjectAssociation.
		DeleteOne(projectAssociation).
		Exec(ctx)
}

func (r *mutationResolver) CreateRepository(ctx context.Context, input ent.CreateRepositoryInput) (*ent.Repository, error) {
	return ent.FromContext(ctx).Repository.
		Create().
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) UpdateRepository(ctx context.Context, id int, input ent.UpdateRepositoryInput) (*ent.Repository, error) {
	return ent.FromContext(ctx).Repository.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) DeleteRepository(ctx context.Context, id int) (*ent.Repository, error) {
	repository, err := ent.FromContext(ctx).Repository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return repository, ent.FromContext(ctx).Repository.
		DeleteOne(repository).
		Exec(ctx)
}

func (r *mutationResolver) CreateSite(ctx context.Context, input ent.CreateSiteInput) (*ent.Site, error) {
	return ent.FromContext(ctx).Site.
		Create().
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) UpdateSite(ctx context.Context, id int, input ent.UpdateSiteInput) (*ent.Site, error) {
	return ent.FromContext(ctx).Site.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) DeleteSite(ctx context.Context, id int) (*ent.Site, error) {
	site, err := ent.FromContext(ctx).Site.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return site, ent.FromContext(ctx).Site.
		DeleteOne(site).
		Exec(ctx)
}

func (r *mutationResolver) CreateTechnology(ctx context.Context, input ent.CreateTechnologyInput) (*ent.Technology, error) {
	return ent.FromContext(ctx).Technology.
		Create().
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) UpdateTechnology(ctx context.Context, id int, input ent.UpdateTechnologyInput) (*ent.Technology, error) {
	return ent.FromContext(ctx).Technology.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) DeleteTechnology(ctx context.Context, id int) (*ent.Technology, error) {
	technology, err := ent.FromContext(ctx).Technology.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return technology, ent.FromContext(ctx).Technology.
		DeleteOne(technology).
		Exec(ctx)
}

func (r *mutationResolver) CreateTechnologyAssociation(ctx context.Context, input ent.CreateTechnologyAssociationInput) (*ent.TechnologyAssociation, error) {
	return ent.FromContext(ctx).TechnologyAssociation.
		Create().
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) UpdateTechnologyAssociation(ctx context.Context, id int, input ent.UpdateTechnologyAssociationInput) (*ent.TechnologyAssociation, error) {
	return ent.FromContext(ctx).TechnologyAssociation.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) DeleteTechnologyAssociation(ctx context.Context, id int) (*ent.TechnologyAssociation, error) {
	technologyAssociation, err := ent.FromContext(ctx).TechnologyAssociation.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return technologyAssociation, ent.FromContext(ctx).TechnologyAssociation.
		DeleteOne(technologyAssociation).
		Exec(ctx)
}

func (r *mutationResolver) CreateProjectTechnology(ctx context.Context, input ent.CreateProjectTechnologyInput) (*ent.ProjectTechnology, error) {
	return ent.FromContext(ctx).ProjectTechnology.
		Create().
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) UpdateProjectTechnology(ctx context.Context, id int, input ent.UpdateProjectTechnologyInput) (*ent.ProjectTechnology, error) {
	return ent.FromContext(ctx).ProjectTechnology.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) DeleteProjectTechnology(ctx context.Context, id int) (*ent.ProjectTechnology, error) {
	projectTechnology, err := ent.FromContext(ctx).ProjectTechnology.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return projectTechnology, ent.FromContext(ctx).ProjectTechnology.
		DeleteOne(projectTechnology).
		Exec(ctx)
}

func (r *mutationResolver) CreateRepositoryTechnology(ctx context.Context, input ent.CreateRepositoryTechnologyInput) (*ent.RepositoryTechnology, error) {
	return ent.FromContext(ctx).RepositoryTechnology.
		Create().
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) UpdateRepositoryTechnology(ctx context.Context, id int, input ent.UpdateRepositoryTechnologyInput) (*ent.RepositoryTechnology, error) {
	return ent.FromContext(ctx).RepositoryTechnology.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

func (r *mutationResolver) DeleteRepositoryTechnology(ctx context.Context, id int) (*ent.RepositoryTechnology, error) {
	repositoryTechnology, err := ent.FromContext(ctx).RepositoryTechnology.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return repositoryTechnology, ent.FromContext(ctx).RepositoryTechnology.
		DeleteOne(repositoryTechnology).
		Exec(ctx)
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

func (r *queryResolver) ProjectTechnologies(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ProjectTechnologyOrder, where *ent.ProjectTechnologyWhereInput) (*ent.ProjectTechnologyConnection, error) {
	return r.client.ProjectTechnology.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithProjectTechnologyOrder(orderBy),
			ent.WithProjectTechnologyFilter(where.Filter),
		)
}

func (r *queryResolver) RepositoryTechnologies(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.RepositoryTechnologyOrder, where *ent.RepositoryTechnologyWhereInput) (*ent.RepositoryTechnologyConnection, error) {
	return r.client.RepositoryTechnology.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithRepositoryTechnologyOrder(orderBy),
			ent.WithRepositoryTechnologyFilter(where.Filter),
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
