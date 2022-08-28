// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/fogo-sh/grackdb/ent/githubaccount"
	"github.com/fogo-sh/grackdb/ent/githuborganization"
	"github.com/fogo-sh/grackdb/ent/githuborganizationmember"
)

// GithubOrganizationMember is the model entity for the GithubOrganizationMember schema.
type GithubOrganizationMember struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Role holds the value of the "role" field.
	Role githuborganizationmember.Role `json:"role,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GithubOrganizationMemberQuery when eager-loading is set.
	Edges                                   GithubOrganizationMemberEdges `json:"edges"`
	github_account_organization_memberships *int
	github_organization_members             *int
}

// GithubOrganizationMemberEdges holds the relations/edges for other nodes in the graph.
type GithubOrganizationMemberEdges struct {
	// Organization holds the value of the organization edge.
	Organization *GithubOrganization `json:"organization,omitempty"`
	// Account holds the value of the account edge.
	Account *GithubAccount `json:"account,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GithubOrganizationMemberEdges) OrganizationOrErr() (*GithubOrganization, error) {
	if e.loadedTypes[0] {
		if e.Organization == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: githuborganization.Label}
		}
		return e.Organization, nil
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// AccountOrErr returns the Account value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GithubOrganizationMemberEdges) AccountOrErr() (*GithubAccount, error) {
	if e.loadedTypes[1] {
		if e.Account == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: githubaccount.Label}
		}
		return e.Account, nil
	}
	return nil, &NotLoadedError{edge: "account"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GithubOrganizationMember) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case githuborganizationmember.FieldID:
			values[i] = new(sql.NullInt64)
		case githuborganizationmember.FieldRole:
			values[i] = new(sql.NullString)
		case githuborganizationmember.ForeignKeys[0]: // github_account_organization_memberships
			values[i] = new(sql.NullInt64)
		case githuborganizationmember.ForeignKeys[1]: // github_organization_members
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GithubOrganizationMember", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GithubOrganizationMember fields.
func (gom *GithubOrganizationMember) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case githuborganizationmember.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gom.ID = int(value.Int64)
		case githuborganizationmember.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				gom.Role = githuborganizationmember.Role(value.String)
			}
		case githuborganizationmember.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field github_account_organization_memberships", value)
			} else if value.Valid {
				gom.github_account_organization_memberships = new(int)
				*gom.github_account_organization_memberships = int(value.Int64)
			}
		case githuborganizationmember.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field github_organization_members", value)
			} else if value.Valid {
				gom.github_organization_members = new(int)
				*gom.github_organization_members = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOrganization queries the "organization" edge of the GithubOrganizationMember entity.
func (gom *GithubOrganizationMember) QueryOrganization() *GithubOrganizationQuery {
	return (&GithubOrganizationMemberClient{config: gom.config}).QueryOrganization(gom)
}

// QueryAccount queries the "account" edge of the GithubOrganizationMember entity.
func (gom *GithubOrganizationMember) QueryAccount() *GithubAccountQuery {
	return (&GithubOrganizationMemberClient{config: gom.config}).QueryAccount(gom)
}

// Update returns a builder for updating this GithubOrganizationMember.
// Note that you need to call GithubOrganizationMember.Unwrap() before calling this method if this GithubOrganizationMember
// was returned from a transaction, and the transaction was committed or rolled back.
func (gom *GithubOrganizationMember) Update() *GithubOrganizationMemberUpdateOne {
	return (&GithubOrganizationMemberClient{config: gom.config}).UpdateOne(gom)
}

// Unwrap unwraps the GithubOrganizationMember entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gom *GithubOrganizationMember) Unwrap() *GithubOrganizationMember {
	_tx, ok := gom.config.driver.(*txDriver)
	if !ok {
		panic("ent: GithubOrganizationMember is not a transactional entity")
	}
	gom.config.driver = _tx.drv
	return gom
}

// String implements the fmt.Stringer.
func (gom *GithubOrganizationMember) String() string {
	var builder strings.Builder
	builder.WriteString("GithubOrganizationMember(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gom.ID))
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", gom.Role))
	builder.WriteByte(')')
	return builder.String()
}

// GithubOrganizationMembers is a parsable slice of GithubOrganizationMember.
type GithubOrganizationMembers []*GithubOrganizationMember

func (gom GithubOrganizationMembers) config(cfg config) {
	for _i := range gom {
		gom[_i].config = cfg
	}
}
