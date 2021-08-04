// Code generated by entc, DO NOT EDIT.

package repository

import (
	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the repository type in the database.
	Label = "repository"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// EdgeGithubAccount holds the string denoting the github_account edge name in mutations.
	EdgeGithubAccount = "github_account"
	// EdgeGithubOrganization holds the string denoting the github_organization edge name in mutations.
	EdgeGithubOrganization = "github_organization"
	// EdgeDiscordBots holds the string denoting the discord_bots edge name in mutations.
	EdgeDiscordBots = "discord_bots"
	// Table holds the table name of the repository in the database.
	Table = "repositories"
	// ProjectTable is the table the holds the project relation/edge.
	ProjectTable = "repositories"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "projects"
	// ProjectColumn is the table column denoting the project relation/edge.
	ProjectColumn = "project_repositories"
	// GithubAccountTable is the table the holds the github_account relation/edge.
	GithubAccountTable = "repositories"
	// GithubAccountInverseTable is the table name for the GithubAccount entity.
	// It exists in this package in order to avoid circular dependency with the "githubaccount" package.
	GithubAccountInverseTable = "github_accounts"
	// GithubAccountColumn is the table column denoting the github_account relation/edge.
	GithubAccountColumn = "github_account_repositories"
	// GithubOrganizationTable is the table the holds the github_organization relation/edge.
	GithubOrganizationTable = "repositories"
	// GithubOrganizationInverseTable is the table name for the GithubOrganization entity.
	// It exists in this package in order to avoid circular dependency with the "githuborganization" package.
	GithubOrganizationInverseTable = "github_organizations"
	// GithubOrganizationColumn is the table column denoting the github_organization relation/edge.
	GithubOrganizationColumn = "github_organization_repositories"
	// DiscordBotsTable is the table the holds the discord_bots relation/edge.
	DiscordBotsTable = "discord_bots"
	// DiscordBotsInverseTable is the table name for the DiscordBot entity.
	// It exists in this package in order to avoid circular dependency with the "discordbot" package.
	DiscordBotsInverseTable = "discord_bots"
	// DiscordBotsColumn is the table column denoting the discord_bots relation/edge.
	DiscordBotsColumn = "repository_discord_bots"
)

// Columns holds all SQL columns for repository fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "repositories"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"github_account_repositories",
	"github_organization_repositories",
	"project_repositories",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/fogo-sh/grackdb/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)
