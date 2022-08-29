package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/fogo-sh/grackdb"
	"github.com/fogo-sh/grackdb/ent"
	"github.com/golang-jwt/jwt"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	return ent.FromContext(ctx).User.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input ent.UpdateUserInput) (*ent.User, error) {
	return ent.FromContext(ctx).User.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*ent.User, error) {
	user, err := ent.FromContext(ctx).User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, ent.FromContext(ctx).User.
		DeleteOne(user).
		Exec(ctx)
}

// CreateDiscordAccount is the resolver for the createDiscordAccount field.
func (r *mutationResolver) CreateDiscordAccount(ctx context.Context, input ent.CreateDiscordAccountInput) (*ent.DiscordAccount, error) {
	return ent.FromContext(ctx).DiscordAccount.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdateDiscordAccount is the resolver for the updateDiscordAccount field.
func (r *mutationResolver) UpdateDiscordAccount(ctx context.Context, id int, input ent.UpdateDiscordAccountInput) (*ent.DiscordAccount, error) {
	return ent.FromContext(ctx).DiscordAccount.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

// DeleteDiscordAccount is the resolver for the deleteDiscordAccount field.
func (r *mutationResolver) DeleteDiscordAccount(ctx context.Context, id int) (*ent.DiscordAccount, error) {
	discordAccount, err := ent.FromContext(ctx).DiscordAccount.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return discordAccount, ent.FromContext(ctx).DiscordAccount.
		DeleteOne(discordAccount).
		Exec(ctx)
}

// CreateDiscordBot is the resolver for the createDiscordBot field.
func (r *mutationResolver) CreateDiscordBot(ctx context.Context, input ent.CreateDiscordBotInput) (*ent.DiscordBot, error) {
	return ent.FromContext(ctx).DiscordBot.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdateDiscordBot is the resolver for the updateDiscordBot field.
func (r *mutationResolver) UpdateDiscordBot(ctx context.Context, id int, input ent.UpdateDiscordBotInput) (*ent.DiscordBot, error) {
	return ent.FromContext(ctx).DiscordBot.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

// DeleteDiscordBot is the resolver for the deleteDiscordBot field.
func (r *mutationResolver) DeleteDiscordBot(ctx context.Context, id int) (*ent.DiscordBot, error) {
	discordBot, err := ent.FromContext(ctx).DiscordBot.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return discordBot, ent.FromContext(ctx).DiscordBot.
		DeleteOne(discordBot).
		Exec(ctx)
}

// CreateGithubAccount is the resolver for the createGithubAccount field.
func (r *mutationResolver) CreateGithubAccount(ctx context.Context, input ent.CreateGithubAccountInput) (*ent.GithubAccount, error) {
	return ent.FromContext(ctx).GithubAccount.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdateGithubAccount is the resolver for the updateGithubAccount field.
func (r *mutationResolver) UpdateGithubAccount(ctx context.Context, id int, input ent.UpdateGithubAccountInput) (*ent.GithubAccount, error) {
	return ent.FromContext(ctx).GithubAccount.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

// DeleteGithubAccount is the resolver for the deleteGithubAccount field.
func (r *mutationResolver) DeleteGithubAccount(ctx context.Context, id int) (*ent.GithubAccount, error) {
	githubAccount, err := ent.FromContext(ctx).GithubAccount.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return githubAccount, ent.FromContext(ctx).GithubAccount.
		DeleteOne(githubAccount).
		Exec(ctx)
}

// CreateGithubOrganization is the resolver for the createGithubOrganization field.
func (r *mutationResolver) CreateGithubOrganization(ctx context.Context, input ent.CreateGithubOrganizationInput) (*ent.GithubOrganization, error) {
	return ent.FromContext(ctx).GithubOrganization.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdateGithubOrganization is the resolver for the updateGithubOrganization field.
func (r *mutationResolver) UpdateGithubOrganization(ctx context.Context, id int, input ent.UpdateGithubOrganizationInput) (*ent.GithubOrganization, error) {
	return ent.FromContext(ctx).GithubOrganization.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

// DeleteGithubOrganization is the resolver for the deleteGithubOrganization field.
func (r *mutationResolver) DeleteGithubOrganization(ctx context.Context, id int) (*ent.GithubOrganization, error) {
	githubOrganization, err := ent.FromContext(ctx).GithubOrganization.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return githubOrganization, ent.FromContext(ctx).GithubOrganization.
		DeleteOne(githubOrganization).
		Exec(ctx)
}

// CreateGithubOrganizationMember is the resolver for the createGithubOrganizationMember field.
func (r *mutationResolver) CreateGithubOrganizationMember(ctx context.Context, input ent.CreateGithubOrganizationMemberInput) (*ent.GithubOrganizationMember, error) {
	return ent.FromContext(ctx).GithubOrganizationMember.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdateGithubOrganizationMember is the resolver for the updateGithubOrganizationMember field.
func (r *mutationResolver) UpdateGithubOrganizationMember(ctx context.Context, id int, input ent.UpdateGithubOrganizationMemberInput) (*ent.GithubOrganizationMember, error) {
	return ent.FromContext(ctx).GithubOrganizationMember.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

// DeleteGithubOrganizationMember is the resolver for the deleteGithubOrganizationMember field.
func (r *mutationResolver) DeleteGithubOrganizationMember(ctx context.Context, id int) (*ent.GithubOrganizationMember, error) {
	githubOrganizationMember, err := ent.FromContext(ctx).GithubOrganizationMember.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return githubOrganizationMember, ent.FromContext(ctx).GithubOrganizationMember.
		DeleteOne(githubOrganizationMember).
		Exec(ctx)
}

// CreateProject is the resolver for the createProject field.
func (r *mutationResolver) CreateProject(ctx context.Context, input ent.CreateProjectInput) (*ent.Project, error) {
	return ent.FromContext(ctx).Project.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdateProject is the resolver for the updateProject field.
func (r *mutationResolver) UpdateProject(ctx context.Context, id int, input ent.UpdateProjectInput) (*ent.Project, error) {
	return ent.FromContext(ctx).Project.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

// DeleteProject is the resolver for the deleteProject field.
func (r *mutationResolver) DeleteProject(ctx context.Context, id int) (*ent.Project, error) {
	project, err := ent.FromContext(ctx).Project.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return project, ent.FromContext(ctx).Project.
		DeleteOne(project).
		Exec(ctx)
}

// CreateProjectContributor is the resolver for the createProjectContributor field.
func (r *mutationResolver) CreateProjectContributor(ctx context.Context, input ent.CreateProjectContributorInput) (*ent.ProjectContributor, error) {
	return ent.FromContext(ctx).ProjectContributor.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdateProjectContributor is the resolver for the updateProjectContributor field.
func (r *mutationResolver) UpdateProjectContributor(ctx context.Context, id int, input ent.UpdateProjectContributorInput) (*ent.ProjectContributor, error) {
	return ent.FromContext(ctx).ProjectContributor.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

// DeleteProjectContributor is the resolver for the deleteProjectContributor field.
func (r *mutationResolver) DeleteProjectContributor(ctx context.Context, id int) (*ent.ProjectContributor, error) {
	projectContributor, err := ent.FromContext(ctx).ProjectContributor.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return projectContributor, ent.FromContext(ctx).ProjectContributor.
		DeleteOne(projectContributor).
		Exec(ctx)
}

// CreateProjectAssociation is the resolver for the createProjectAssociation field.
func (r *mutationResolver) CreateProjectAssociation(ctx context.Context, input ent.CreateProjectAssociationInput) (*ent.ProjectAssociation, error) {
	return ent.FromContext(ctx).ProjectAssociation.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdateProjectAssociation is the resolver for the updateProjectAssociation field.
func (r *mutationResolver) UpdateProjectAssociation(ctx context.Context, id int, input ent.UpdateProjectAssociationInput) (*ent.ProjectAssociation, error) {
	return ent.FromContext(ctx).ProjectAssociation.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

// DeleteProjectAssociation is the resolver for the deleteProjectAssociation field.
func (r *mutationResolver) DeleteProjectAssociation(ctx context.Context, id int) (*ent.ProjectAssociation, error) {
	projectAssociation, err := ent.FromContext(ctx).ProjectAssociation.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return projectAssociation, ent.FromContext(ctx).ProjectAssociation.
		DeleteOne(projectAssociation).
		Exec(ctx)
}

// CreateRepository is the resolver for the createRepository field.
func (r *mutationResolver) CreateRepository(ctx context.Context, input ent.CreateRepositoryInput) (*ent.Repository, error) {
	return ent.FromContext(ctx).Repository.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdateRepository is the resolver for the updateRepository field.
func (r *mutationResolver) UpdateRepository(ctx context.Context, id int, input ent.UpdateRepositoryInput) (*ent.Repository, error) {
	return ent.FromContext(ctx).Repository.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

// DeleteRepository is the resolver for the deleteRepository field.
func (r *mutationResolver) DeleteRepository(ctx context.Context, id int) (*ent.Repository, error) {
	repository, err := ent.FromContext(ctx).Repository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return repository, ent.FromContext(ctx).Repository.
		DeleteOne(repository).
		Exec(ctx)
}

// CreateSite is the resolver for the createSite field.
func (r *mutationResolver) CreateSite(ctx context.Context, input ent.CreateSiteInput) (*ent.Site, error) {
	return ent.FromContext(ctx).Site.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdateSite is the resolver for the updateSite field.
func (r *mutationResolver) UpdateSite(ctx context.Context, id int, input ent.UpdateSiteInput) (*ent.Site, error) {
	return ent.FromContext(ctx).Site.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

// DeleteSite is the resolver for the deleteSite field.
func (r *mutationResolver) DeleteSite(ctx context.Context, id int) (*ent.Site, error) {
	site, err := ent.FromContext(ctx).Site.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return site, ent.FromContext(ctx).Site.
		DeleteOne(site).
		Exec(ctx)
}

// CreateTechnology is the resolver for the createTechnology field.
func (r *mutationResolver) CreateTechnology(ctx context.Context, input ent.CreateTechnologyInput) (*ent.Technology, error) {
	return ent.FromContext(ctx).Technology.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdateTechnology is the resolver for the updateTechnology field.
func (r *mutationResolver) UpdateTechnology(ctx context.Context, id int, input ent.UpdateTechnologyInput) (*ent.Technology, error) {
	return ent.FromContext(ctx).Technology.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

// DeleteTechnology is the resolver for the deleteTechnology field.
func (r *mutationResolver) DeleteTechnology(ctx context.Context, id int) (*ent.Technology, error) {
	technology, err := ent.FromContext(ctx).Technology.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return technology, ent.FromContext(ctx).Technology.
		DeleteOne(technology).
		Exec(ctx)
}

// CreateTechnologyAssociation is the resolver for the createTechnologyAssociation field.
func (r *mutationResolver) CreateTechnologyAssociation(ctx context.Context, input ent.CreateTechnologyAssociationInput) (*ent.TechnologyAssociation, error) {
	return ent.FromContext(ctx).TechnologyAssociation.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdateTechnologyAssociation is the resolver for the updateTechnologyAssociation field.
func (r *mutationResolver) UpdateTechnologyAssociation(ctx context.Context, id int, input ent.UpdateTechnologyAssociationInput) (*ent.TechnologyAssociation, error) {
	return ent.FromContext(ctx).TechnologyAssociation.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

// DeleteTechnologyAssociation is the resolver for the deleteTechnologyAssociation field.
func (r *mutationResolver) DeleteTechnologyAssociation(ctx context.Context, id int) (*ent.TechnologyAssociation, error) {
	technologyAssociation, err := ent.FromContext(ctx).TechnologyAssociation.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return technologyAssociation, ent.FromContext(ctx).TechnologyAssociation.
		DeleteOne(technologyAssociation).
		Exec(ctx)
}

// CreateProjectTechnology is the resolver for the createProjectTechnology field.
func (r *mutationResolver) CreateProjectTechnology(ctx context.Context, input ent.CreateProjectTechnologyInput) (*ent.ProjectTechnology, error) {
	return ent.FromContext(ctx).ProjectTechnology.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdateProjectTechnology is the resolver for the updateProjectTechnology field.
func (r *mutationResolver) UpdateProjectTechnology(ctx context.Context, id int, input ent.UpdateProjectTechnologyInput) (*ent.ProjectTechnology, error) {
	return ent.FromContext(ctx).ProjectTechnology.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

// DeleteProjectTechnology is the resolver for the deleteProjectTechnology field.
func (r *mutationResolver) DeleteProjectTechnology(ctx context.Context, id int) (*ent.ProjectTechnology, error) {
	projectTechnology, err := ent.FromContext(ctx).ProjectTechnology.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return projectTechnology, ent.FromContext(ctx).ProjectTechnology.
		DeleteOne(projectTechnology).
		Exec(ctx)
}

// CreateRepositoryTechnology is the resolver for the createRepositoryTechnology field.
func (r *mutationResolver) CreateRepositoryTechnology(ctx context.Context, input ent.CreateRepositoryTechnologyInput) (*ent.RepositoryTechnology, error) {
	return ent.FromContext(ctx).RepositoryTechnology.
		Create().
		SetInput(input).
		Save(ctx)
}

// UpdateRepositoryTechnology is the resolver for the updateRepositoryTechnology field.
func (r *mutationResolver) UpdateRepositoryTechnology(ctx context.Context, id int, input ent.UpdateRepositoryTechnologyInput) (*ent.RepositoryTechnology, error) {
	return ent.FromContext(ctx).RepositoryTechnology.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

// DeleteRepositoryTechnology is the resolver for the deleteRepositoryTechnology field.
func (r *mutationResolver) DeleteRepositoryTechnology(ctx context.Context, id int) (*ent.RepositoryTechnology, error) {
	repositoryTechnology, err := ent.FromContext(ctx).RepositoryTechnology.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return repositoryTechnology, ent.FromContext(ctx).RepositoryTechnology.
		DeleteOne(repositoryTechnology).
		Exec(ctx)
}

// AssumeDevelopmentUser is the resolver for the assumeDevelopmentUser field.
func (r *mutationResolver) AssumeDevelopmentUser(ctx context.Context, id int) (*ent.User, error) {
	if !grackdb.AppConfig.DevelopmentMode {
		return nil, errors.New("assumeDevelopmentUser is only available in development mode")
	}

	user, err := ent.FromContext(ctx).User.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error fetching user: %w", err)
	}

	unsignedToken := jwt.NewWithClaims(
		jwt.SigningMethodHS512,
		jwt.MapClaims{
			"iat": time.Now().UTC().Unix(),
			"sub": user.ID,
		},
	)

	tokenString, err := unsignedToken.SignedString([]byte(grackdb.AppConfig.JwtSigningSecret))

	if err != nil {
		return nil, fmt.Errorf("error signing token: %w", err)
	}

	ginContext, err := grackdb.GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting gin context: %w", err)
	}

	ginContext.SetCookie("jwt", tokenString, 0, "/", "", false, false)

	return user, nil
}

// AvailableAuthProviders is the resolver for the availableAuthProviders field.
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

// CurrentUser is the resolver for the currentUser field.
func (r *queryResolver) CurrentUser(ctx context.Context) (*ent.User, error) {
	ginCtx, err := grackdb.GinContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("error fetching context: %w", err)
	}
	return ginCtx.Value("user").(*ent.User), nil
}

// DevelopmentMode is the resolver for the developmentMode field.
func (r *queryResolver) DevelopmentMode(ctx context.Context) (bool, error) {
	return grackdb.AppConfig.DevelopmentMode, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
