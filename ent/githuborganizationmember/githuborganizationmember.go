// Code generated by entc, DO NOT EDIT.

package githuborganizationmember

import (
	"fmt"
	"io"
	"strconv"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the githuborganizationmember type in the database.
	Label = "github_organization_member"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// EdgeOrganization holds the string denoting the organization edge name in mutations.
	EdgeOrganization = "organization"
	// EdgeAccount holds the string denoting the account edge name in mutations.
	EdgeAccount = "account"
	// Table holds the table name of the githuborganizationmember in the database.
	Table = "github_organization_members"
	// OrganizationTable is the table the holds the organization relation/edge.
	OrganizationTable = "github_organization_members"
	// OrganizationInverseTable is the table name for the GithubOrganization entity.
	// It exists in this package in order to avoid circular dependency with the "githuborganization" package.
	OrganizationInverseTable = "github_organizations"
	// OrganizationColumn is the table column denoting the organization relation/edge.
	OrganizationColumn = "github_organization_members"
	// AccountTable is the table the holds the account relation/edge.
	AccountTable = "github_organization_members"
	// AccountInverseTable is the table name for the GithubAccount entity.
	// It exists in this package in order to avoid circular dependency with the "githubaccount" package.
	AccountInverseTable = "github_accounts"
	// AccountColumn is the table column denoting the account relation/edge.
	AccountColumn = "github_account_organization_memberships"
)

// Columns holds all SQL columns for githuborganizationmember fields.
var Columns = []string{
	FieldID,
	FieldRole,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "github_organization_members"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"github_account_organization_memberships",
	"github_organization_members",
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
)

// Role defines the type for the "role" enum field.
type Role string

// RoleMember is the default value of the Role enum.
const DefaultRole = RoleMember

// Role values.
const (
	RoleAdmin  Role = "ADMIN"
	RoleMember Role = "MEMBER"
)

func (r Role) String() string {
	return string(r)
}

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleAdmin, RoleMember:
		return nil
	default:
		return fmt.Errorf("githuborganizationmember: invalid enum value for role field: %q", r)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (r Role) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(r.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (r *Role) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*r = Role(str)
	if err := RoleValidator(*r); err != nil {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}
