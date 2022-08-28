// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DiscordAccountsColumns holds the columns for the "discord_accounts" table.
	DiscordAccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "discord_id", Type: field.TypeString, Unique: true},
		{Name: "username", Type: field.TypeString},
		{Name: "discriminator", Type: field.TypeString, Size: 4},
		{Name: "discord_bot_account", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "user_discord_accounts", Type: field.TypeInt, Nullable: true},
	}
	// DiscordAccountsTable holds the schema information for the "discord_accounts" table.
	DiscordAccountsTable = &schema.Table{
		Name:       "discord_accounts",
		Columns:    DiscordAccountsColumns,
		PrimaryKey: []*schema.Column{DiscordAccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "discord_accounts_discord_bots_account",
				Columns:    []*schema.Column{DiscordAccountsColumns[4]},
				RefColumns: []*schema.Column{DiscordBotsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "discord_accounts_users_discord_accounts",
				Columns:    []*schema.Column{DiscordAccountsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DiscordBotsColumns holds the columns for the "discord_bots" table.
	DiscordBotsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "project_discord_bots", Type: field.TypeInt},
		{Name: "repository_discord_bots", Type: field.TypeInt, Nullable: true},
	}
	// DiscordBotsTable holds the schema information for the "discord_bots" table.
	DiscordBotsTable = &schema.Table{
		Name:       "discord_bots",
		Columns:    DiscordBotsColumns,
		PrimaryKey: []*schema.Column{DiscordBotsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "discord_bots_projects_discord_bots",
				Columns:    []*schema.Column{DiscordBotsColumns[1]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "discord_bots_repositories_discord_bots",
				Columns:    []*schema.Column{DiscordBotsColumns[2]},
				RefColumns: []*schema.Column{RepositoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GithubAccountsColumns holds the columns for the "github_accounts" table.
	GithubAccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "user_github_accounts", Type: field.TypeInt},
	}
	// GithubAccountsTable holds the schema information for the "github_accounts" table.
	GithubAccountsTable = &schema.Table{
		Name:       "github_accounts",
		Columns:    GithubAccountsColumns,
		PrimaryKey: []*schema.Column{GithubAccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "github_accounts_users_github_accounts",
				Columns:    []*schema.Column{GithubAccountsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// GithubOrganizationsColumns holds the columns for the "github_organizations" table.
	GithubOrganizationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "display_name", Type: field.TypeString, Nullable: true},
	}
	// GithubOrganizationsTable holds the schema information for the "github_organizations" table.
	GithubOrganizationsTable = &schema.Table{
		Name:       "github_organizations",
		Columns:    GithubOrganizationsColumns,
		PrimaryKey: []*schema.Column{GithubOrganizationsColumns[0]},
	}
	// GithubOrganizationMembersColumns holds the columns for the "github_organization_members" table.
	GithubOrganizationMembersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"ADMIN", "MEMBER"}, Default: "MEMBER"},
		{Name: "github_account_organization_memberships", Type: field.TypeInt},
		{Name: "github_organization_members", Type: field.TypeInt},
	}
	// GithubOrganizationMembersTable holds the schema information for the "github_organization_members" table.
	GithubOrganizationMembersTable = &schema.Table{
		Name:       "github_organization_members",
		Columns:    GithubOrganizationMembersColumns,
		PrimaryKey: []*schema.Column{GithubOrganizationMembersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "github_organization_members_github_accounts_organization_memberships",
				Columns:    []*schema.Column{GithubOrganizationMembersColumns[2]},
				RefColumns: []*schema.Column{GithubAccountsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "github_organization_members_github_organizations_members",
				Columns:    []*schema.Column{GithubOrganizationMembersColumns[3]},
				RefColumns: []*schema.Column{GithubOrganizationsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "start_date", Type: field.TypeTime},
		{Name: "end_date", Type: field.TypeTime, Nullable: true},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:       "projects",
		Columns:    ProjectsColumns,
		PrimaryKey: []*schema.Column{ProjectsColumns[0]},
	}
	// ProjectAssociationsColumns holds the columns for the "project_associations" table.
	ProjectAssociationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"BASED_OFF", "REPLACES", "INSPIRED_BY", "RELATED"}},
		{Name: "project_parent_projects", Type: field.TypeInt},
		{Name: "project_child_projects", Type: field.TypeInt},
	}
	// ProjectAssociationsTable holds the schema information for the "project_associations" table.
	ProjectAssociationsTable = &schema.Table{
		Name:       "project_associations",
		Columns:    ProjectAssociationsColumns,
		PrimaryKey: []*schema.Column{ProjectAssociationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "project_associations_projects_parent_projects",
				Columns:    []*schema.Column{ProjectAssociationsColumns[2]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "project_associations_projects_child_projects",
				Columns:    []*schema.Column{ProjectAssociationsColumns[3]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ProjectContributorsColumns holds the columns for the "project_contributors" table.
	ProjectContributorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"OWNER", "CONTRIBUTOR"}},
		{Name: "project_contributors", Type: field.TypeInt},
		{Name: "user_project_contributions", Type: field.TypeInt},
	}
	// ProjectContributorsTable holds the schema information for the "project_contributors" table.
	ProjectContributorsTable = &schema.Table{
		Name:       "project_contributors",
		Columns:    ProjectContributorsColumns,
		PrimaryKey: []*schema.Column{ProjectContributorsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "project_contributors_projects_contributors",
				Columns:    []*schema.Column{ProjectContributorsColumns[2]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "project_contributors_users_project_contributions",
				Columns:    []*schema.Column{ProjectContributorsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ProjectTechnologiesColumns holds the columns for the "project_technologies" table.
	ProjectTechnologiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"WRITTEN_IN", "IMPLEMENTS", "USES", "CONTAINS"}},
		{Name: "project_technologies", Type: field.TypeInt},
		{Name: "technology_projects", Type: field.TypeInt},
	}
	// ProjectTechnologiesTable holds the schema information for the "project_technologies" table.
	ProjectTechnologiesTable = &schema.Table{
		Name:       "project_technologies",
		Columns:    ProjectTechnologiesColumns,
		PrimaryKey: []*schema.Column{ProjectTechnologiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "project_technologies_projects_technologies",
				Columns:    []*schema.Column{ProjectTechnologiesColumns[2]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "project_technologies_technologies_projects",
				Columns:    []*schema.Column{ProjectTechnologiesColumns[3]},
				RefColumns: []*schema.Column{TechnologiesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// RepositoriesColumns holds the columns for the "repositories" table.
	RepositoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "github_account_repositories", Type: field.TypeInt, Nullable: true},
		{Name: "github_organization_repositories", Type: field.TypeInt, Nullable: true},
		{Name: "project_repositories", Type: field.TypeInt},
	}
	// RepositoriesTable holds the schema information for the "repositories" table.
	RepositoriesTable = &schema.Table{
		Name:       "repositories",
		Columns:    RepositoriesColumns,
		PrimaryKey: []*schema.Column{RepositoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "repositories_github_accounts_repositories",
				Columns:    []*schema.Column{RepositoriesColumns[3]},
				RefColumns: []*schema.Column{GithubAccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "repositories_github_organizations_repositories",
				Columns:    []*schema.Column{RepositoriesColumns[4]},
				RefColumns: []*schema.Column{GithubOrganizationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "repositories_projects_repositories",
				Columns:    []*schema.Column{RepositoriesColumns[5]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// RepositoryTechnologiesColumns holds the columns for the "repository_technologies" table.
	RepositoryTechnologiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"WRITTEN_IN", "IMPLEMENTS", "USES", "CONTAINS"}},
		{Name: "repository_technologies", Type: field.TypeInt},
		{Name: "technology_repositories", Type: field.TypeInt},
	}
	// RepositoryTechnologiesTable holds the schema information for the "repository_technologies" table.
	RepositoryTechnologiesTable = &schema.Table{
		Name:       "repository_technologies",
		Columns:    RepositoryTechnologiesColumns,
		PrimaryKey: []*schema.Column{RepositoryTechnologiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "repository_technologies_repositories_technologies",
				Columns:    []*schema.Column{RepositoryTechnologiesColumns[2]},
				RefColumns: []*schema.Column{RepositoriesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "repository_technologies_technologies_repositories",
				Columns:    []*schema.Column{RepositoryTechnologiesColumns[3]},
				RefColumns: []*schema.Column{TechnologiesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// SitesColumns holds the columns for the "sites" table.
	SitesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "url", Type: field.TypeString},
		{Name: "project_sites", Type: field.TypeInt},
		{Name: "repository_sites", Type: field.TypeInt, Nullable: true},
	}
	// SitesTable holds the schema information for the "sites" table.
	SitesTable = &schema.Table{
		Name:       "sites",
		Columns:    SitesColumns,
		PrimaryKey: []*schema.Column{SitesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sites_projects_sites",
				Columns:    []*schema.Column{SitesColumns[2]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "sites_repositories_sites",
				Columns:    []*schema.Column{SitesColumns[3]},
				RefColumns: []*schema.Column{RepositoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TechnologiesColumns holds the columns for the "technologies" table.
	TechnologiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "colour", Type: field.TypeString, Nullable: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"LIBRARY", "LANGUAGE", "ALGORITHM", "DATABASE", "PROTOCOL", "SERVICE"}},
	}
	// TechnologiesTable holds the schema information for the "technologies" table.
	TechnologiesTable = &schema.Table{
		Name:       "technologies",
		Columns:    TechnologiesColumns,
		PrimaryKey: []*schema.Column{TechnologiesColumns[0]},
	}
	// TechnologyAssociationsColumns holds the columns for the "technology_associations" table.
	TechnologyAssociationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"WRITTEN_IN", "IMPLEMENTS", "USES"}},
		{Name: "technology_parent_technologies", Type: field.TypeInt},
		{Name: "technology_child_technologies", Type: field.TypeInt},
	}
	// TechnologyAssociationsTable holds the schema information for the "technology_associations" table.
	TechnologyAssociationsTable = &schema.Table{
		Name:       "technology_associations",
		Columns:    TechnologyAssociationsColumns,
		PrimaryKey: []*schema.Column{TechnologyAssociationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "technology_associations_technologies_parent_technologies",
				Columns:    []*schema.Column{TechnologyAssociationsColumns[2]},
				RefColumns: []*schema.Column{TechnologiesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "technology_associations_technologies_child_technologies",
				Columns:    []*schema.Column{TechnologyAssociationsColumns[3]},
				RefColumns: []*schema.Column{TechnologiesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "avatar_url", Type: field.TypeString, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DiscordAccountsTable,
		DiscordBotsTable,
		GithubAccountsTable,
		GithubOrganizationsTable,
		GithubOrganizationMembersTable,
		ProjectsTable,
		ProjectAssociationsTable,
		ProjectContributorsTable,
		ProjectTechnologiesTable,
		RepositoriesTable,
		RepositoryTechnologiesTable,
		SitesTable,
		TechnologiesTable,
		TechnologyAssociationsTable,
		UsersTable,
	}
)

func init() {
	DiscordAccountsTable.ForeignKeys[0].RefTable = DiscordBotsTable
	DiscordAccountsTable.ForeignKeys[1].RefTable = UsersTable
	DiscordBotsTable.ForeignKeys[0].RefTable = ProjectsTable
	DiscordBotsTable.ForeignKeys[1].RefTable = RepositoriesTable
	GithubAccountsTable.ForeignKeys[0].RefTable = UsersTable
	GithubOrganizationMembersTable.ForeignKeys[0].RefTable = GithubAccountsTable
	GithubOrganizationMembersTable.ForeignKeys[1].RefTable = GithubOrganizationsTable
	ProjectAssociationsTable.ForeignKeys[0].RefTable = ProjectsTable
	ProjectAssociationsTable.ForeignKeys[1].RefTable = ProjectsTable
	ProjectContributorsTable.ForeignKeys[0].RefTable = ProjectsTable
	ProjectContributorsTable.ForeignKeys[1].RefTable = UsersTable
	ProjectTechnologiesTable.ForeignKeys[0].RefTable = ProjectsTable
	ProjectTechnologiesTable.ForeignKeys[1].RefTable = TechnologiesTable
	RepositoriesTable.ForeignKeys[0].RefTable = GithubAccountsTable
	RepositoriesTable.ForeignKeys[1].RefTable = GithubOrganizationsTable
	RepositoriesTable.ForeignKeys[2].RefTable = ProjectsTable
	RepositoryTechnologiesTable.ForeignKeys[0].RefTable = RepositoriesTable
	RepositoryTechnologiesTable.ForeignKeys[1].RefTable = TechnologiesTable
	SitesTable.ForeignKeys[0].RefTable = ProjectsTable
	SitesTable.ForeignKeys[1].RefTable = RepositoriesTable
	TechnologyAssociationsTable.ForeignKeys[0].RefTable = TechnologiesTable
	TechnologyAssociationsTable.ForeignKeys[1].RefTable = TechnologiesTable
}
