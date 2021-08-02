// Code generated by entc, DO NOT EDIT.

package githubaccount

import (
	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the githubaccount type in the database.
	Label = "github_account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeOrganizationMemberships holds the string denoting the organization_memberships edge name in mutations.
	EdgeOrganizationMemberships = "organization_memberships"
	// EdgeRepositories holds the string denoting the repositories edge name in mutations.
	EdgeRepositories = "repositories"
	// Table holds the table name of the githubaccount in the database.
	Table = "github_accounts"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "github_accounts"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_github_accounts"
	// OrganizationMembershipsTable is the table the holds the organization_memberships relation/edge.
	OrganizationMembershipsTable = "github_organization_members"
	// OrganizationMembershipsInverseTable is the table name for the GithubOrganizationMember entity.
	// It exists in this package in order to avoid circular dependency with the "githuborganizationmember" package.
	OrganizationMembershipsInverseTable = "github_organization_members"
	// OrganizationMembershipsColumn is the table column denoting the organization_memberships relation/edge.
	OrganizationMembershipsColumn = "github_account_organization_memberships"
	// RepositoriesTable is the table the holds the repositories relation/edge.
	RepositoriesTable = "repositories"
	// RepositoriesInverseTable is the table name for the Repository entity.
	// It exists in this package in order to avoid circular dependency with the "repository" package.
	RepositoriesInverseTable = "repositories"
	// RepositoriesColumn is the table column denoting the repositories relation/edge.
	RepositoriesColumn = "github_account_repositories"
)

// Columns holds all SQL columns for githubaccount fields.
var Columns = []string{
	FieldID,
	FieldUsername,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "github_accounts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_github_accounts",
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
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
)
