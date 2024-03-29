// Code generated by ent, DO NOT EDIT.

package githuborganization

import (
	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the githuborganization type in the database.
	Label = "github_organization"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDisplayName holds the string denoting the display_name field in the database.
	FieldDisplayName = "display_name"
	// EdgeMembers holds the string denoting the members edge name in mutations.
	EdgeMembers = "members"
	// EdgeRepositories holds the string denoting the repositories edge name in mutations.
	EdgeRepositories = "repositories"
	// Table holds the table name of the githuborganization in the database.
	Table = "github_organizations"
	// MembersTable is the table that holds the members relation/edge.
	MembersTable = "github_organization_members"
	// MembersInverseTable is the table name for the GithubOrganizationMember entity.
	// It exists in this package in order to avoid circular dependency with the "githuborganizationmember" package.
	MembersInverseTable = "github_organization_members"
	// MembersColumn is the table column denoting the members relation/edge.
	MembersColumn = "github_organization_members"
	// RepositoriesTable is the table that holds the repositories relation/edge.
	RepositoriesTable = "repositories"
	// RepositoriesInverseTable is the table name for the Repository entity.
	// It exists in this package in order to avoid circular dependency with the "repository" package.
	RepositoriesInverseTable = "repositories"
	// RepositoriesColumn is the table column denoting the repositories relation/edge.
	RepositoriesColumn = "github_organization_repositories"
)

// Columns holds all SQL columns for githuborganization fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDisplayName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)
