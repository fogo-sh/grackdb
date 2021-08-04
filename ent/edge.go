// Code generated by entc, DO NOT EDIT.

package ent

import "context"

func (da *DiscordAccount) Owner(ctx context.Context) (*User, error) {
	result, err := da.Edges.OwnerOrErr()
	if IsNotLoaded(err) {
		result, err = da.QueryOwner().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (da *DiscordAccount) Bot(ctx context.Context) (*DiscordBot, error) {
	result, err := da.Edges.BotOrErr()
	if IsNotLoaded(err) {
		result, err = da.QueryBot().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (db *DiscordBot) Account(ctx context.Context) (*DiscordAccount, error) {
	result, err := db.Edges.AccountOrErr()
	if IsNotLoaded(err) {
		result, err = db.QueryAccount().Only(ctx)
	}
	return result, err
}

func (db *DiscordBot) Project(ctx context.Context) (*Project, error) {
	result, err := db.Edges.ProjectOrErr()
	if IsNotLoaded(err) {
		result, err = db.QueryProject().Only(ctx)
	}
	return result, err
}

func (db *DiscordBot) Repository(ctx context.Context) (*Repository, error) {
	result, err := db.Edges.RepositoryOrErr()
	if IsNotLoaded(err) {
		result, err = db.QueryRepository().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (ga *GithubAccount) Owner(ctx context.Context) (*User, error) {
	result, err := ga.Edges.OwnerOrErr()
	if IsNotLoaded(err) {
		result, err = ga.QueryOwner().Only(ctx)
	}
	return result, err
}

func (ga *GithubAccount) OrganizationMemberships(ctx context.Context) ([]*GithubOrganizationMember, error) {
	result, err := ga.Edges.OrganizationMembershipsOrErr()
	if IsNotLoaded(err) {
		result, err = ga.QueryOrganizationMemberships().All(ctx)
	}
	return result, err
}

func (ga *GithubAccount) Repositories(ctx context.Context) ([]*Repository, error) {
	result, err := ga.Edges.RepositoriesOrErr()
	if IsNotLoaded(err) {
		result, err = ga.QueryRepositories().All(ctx)
	}
	return result, err
}

func (_go *GithubOrganization) Members(ctx context.Context) ([]*GithubOrganizationMember, error) {
	result, err := _go.Edges.MembersOrErr()
	if IsNotLoaded(err) {
		result, err = _go.QueryMembers().All(ctx)
	}
	return result, err
}

func (_go *GithubOrganization) Repositories(ctx context.Context) ([]*Repository, error) {
	result, err := _go.Edges.RepositoriesOrErr()
	if IsNotLoaded(err) {
		result, err = _go.QueryRepositories().All(ctx)
	}
	return result, err
}

func (gom *GithubOrganizationMember) Organization(ctx context.Context) (*GithubOrganization, error) {
	result, err := gom.Edges.OrganizationOrErr()
	if IsNotLoaded(err) {
		result, err = gom.QueryOrganization().Only(ctx)
	}
	return result, err
}

func (gom *GithubOrganizationMember) Account(ctx context.Context) (*GithubAccount, error) {
	result, err := gom.Edges.AccountOrErr()
	if IsNotLoaded(err) {
		result, err = gom.QueryAccount().Only(ctx)
	}
	return result, err
}

func (pr *Project) Contributors(ctx context.Context) ([]*ProjectContributor, error) {
	result, err := pr.Edges.ContributorsOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryContributors().All(ctx)
	}
	return result, err
}

func (pr *Project) ParentProjects(ctx context.Context) ([]*ProjectAssociation, error) {
	result, err := pr.Edges.ParentProjectsOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryParentProjects().All(ctx)
	}
	return result, err
}

func (pr *Project) ChildProjects(ctx context.Context) ([]*ProjectAssociation, error) {
	result, err := pr.Edges.ChildProjectsOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryChildProjects().All(ctx)
	}
	return result, err
}

func (pr *Project) Repositories(ctx context.Context) ([]*Repository, error) {
	result, err := pr.Edges.RepositoriesOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryRepositories().All(ctx)
	}
	return result, err
}

func (pr *Project) DiscordBots(ctx context.Context) ([]*DiscordBot, error) {
	result, err := pr.Edges.DiscordBotsOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryDiscordBots().All(ctx)
	}
	return result, err
}

func (pr *Project) Sites(ctx context.Context) ([]*Site, error) {
	result, err := pr.Edges.SitesOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QuerySites().All(ctx)
	}
	return result, err
}

func (pr *Project) Technologies(ctx context.Context) ([]*ProjectTechnology, error) {
	result, err := pr.Edges.TechnologiesOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryTechnologies().All(ctx)
	}
	return result, err
}

func (pa *ProjectAssociation) Parent(ctx context.Context) (*Project, error) {
	result, err := pa.Edges.ParentOrErr()
	if IsNotLoaded(err) {
		result, err = pa.QueryParent().Only(ctx)
	}
	return result, err
}

func (pa *ProjectAssociation) Child(ctx context.Context) (*Project, error) {
	result, err := pa.Edges.ChildOrErr()
	if IsNotLoaded(err) {
		result, err = pa.QueryChild().Only(ctx)
	}
	return result, err
}

func (pc *ProjectContributor) Project(ctx context.Context) (*Project, error) {
	result, err := pc.Edges.ProjectOrErr()
	if IsNotLoaded(err) {
		result, err = pc.QueryProject().Only(ctx)
	}
	return result, err
}

func (pc *ProjectContributor) User(ctx context.Context) (*User, error) {
	result, err := pc.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = pc.QueryUser().Only(ctx)
	}
	return result, err
}

func (pt *ProjectTechnology) Project(ctx context.Context) (*Project, error) {
	result, err := pt.Edges.ProjectOrErr()
	if IsNotLoaded(err) {
		result, err = pt.QueryProject().Only(ctx)
	}
	return result, err
}

func (pt *ProjectTechnology) Technology(ctx context.Context) (*Technology, error) {
	result, err := pt.Edges.TechnologyOrErr()
	if IsNotLoaded(err) {
		result, err = pt.QueryTechnology().Only(ctx)
	}
	return result, err
}

func (r *Repository) Project(ctx context.Context) (*Project, error) {
	result, err := r.Edges.ProjectOrErr()
	if IsNotLoaded(err) {
		result, err = r.QueryProject().Only(ctx)
	}
	return result, err
}

func (r *Repository) GithubAccount(ctx context.Context) (*GithubAccount, error) {
	result, err := r.Edges.GithubAccountOrErr()
	if IsNotLoaded(err) {
		result, err = r.QueryGithubAccount().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (r *Repository) GithubOrganization(ctx context.Context) (*GithubOrganization, error) {
	result, err := r.Edges.GithubOrganizationOrErr()
	if IsNotLoaded(err) {
		result, err = r.QueryGithubOrganization().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (r *Repository) DiscordBots(ctx context.Context) ([]*DiscordBot, error) {
	result, err := r.Edges.DiscordBotsOrErr()
	if IsNotLoaded(err) {
		result, err = r.QueryDiscordBots().All(ctx)
	}
	return result, err
}

func (r *Repository) Sites(ctx context.Context) ([]*Site, error) {
	result, err := r.Edges.SitesOrErr()
	if IsNotLoaded(err) {
		result, err = r.QuerySites().All(ctx)
	}
	return result, err
}

func (s *Site) Project(ctx context.Context) (*Project, error) {
	result, err := s.Edges.ProjectOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryProject().Only(ctx)
	}
	return result, err
}

func (s *Site) Repository(ctx context.Context) (*Repository, error) {
	result, err := s.Edges.RepositoryOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryRepository().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (t *Technology) ParentTechnologies(ctx context.Context) ([]*TechnologyAssociation, error) {
	result, err := t.Edges.ParentTechnologiesOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryParentTechnologies().All(ctx)
	}
	return result, err
}

func (t *Technology) ChildTechnologies(ctx context.Context) ([]*TechnologyAssociation, error) {
	result, err := t.Edges.ChildTechnologiesOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryChildTechnologies().All(ctx)
	}
	return result, err
}

func (t *Technology) Projects(ctx context.Context) ([]*ProjectTechnology, error) {
	result, err := t.Edges.ProjectsOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryProjects().All(ctx)
	}
	return result, err
}

func (ta *TechnologyAssociation) Parent(ctx context.Context) (*Technology, error) {
	result, err := ta.Edges.ParentOrErr()
	if IsNotLoaded(err) {
		result, err = ta.QueryParent().Only(ctx)
	}
	return result, err
}

func (ta *TechnologyAssociation) Child(ctx context.Context) (*Technology, error) {
	result, err := ta.Edges.ChildOrErr()
	if IsNotLoaded(err) {
		result, err = ta.QueryChild().Only(ctx)
	}
	return result, err
}

func (u *User) DiscordAccounts(ctx context.Context) ([]*DiscordAccount, error) {
	result, err := u.Edges.DiscordAccountsOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryDiscordAccounts().All(ctx)
	}
	return result, err
}

func (u *User) GithubAccounts(ctx context.Context) ([]*GithubAccount, error) {
	result, err := u.Edges.GithubAccountsOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryGithubAccounts().All(ctx)
	}
	return result, err
}

func (u *User) ProjectContributions(ctx context.Context) ([]*ProjectContributor, error) {
	result, err := u.Edges.ProjectContributionsOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryProjectContributions().All(ctx)
	}
	return result, err
}
