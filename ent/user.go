// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/fogo-sh/grackdb/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// AvatarURL holds the value of the "avatar_url" field.
	AvatarURL *string `json:"avatar_url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// DiscordAccounts holds the value of the discord_accounts edge.
	DiscordAccounts []*DiscordAccount `json:"discord_accounts,omitempty"`
	// GithubAccounts holds the value of the github_accounts edge.
	GithubAccounts []*GithubAccount `json:"github_accounts,omitempty"`
	// ProjectContributions holds the value of the project_contributions edge.
	ProjectContributions []*ProjectContributor `json:"project_contributions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// DiscordAccountsOrErr returns the DiscordAccounts value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) DiscordAccountsOrErr() ([]*DiscordAccount, error) {
	if e.loadedTypes[0] {
		return e.DiscordAccounts, nil
	}
	return nil, &NotLoadedError{edge: "discord_accounts"}
}

// GithubAccountsOrErr returns the GithubAccounts value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) GithubAccountsOrErr() ([]*GithubAccount, error) {
	if e.loadedTypes[1] {
		return e.GithubAccounts, nil
	}
	return nil, &NotLoadedError{edge: "github_accounts"}
}

// ProjectContributionsOrErr returns the ProjectContributions value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ProjectContributionsOrErr() ([]*ProjectContributor, error) {
	if e.loadedTypes[2] {
		return e.ProjectContributions, nil
	}
	return nil, &NotLoadedError{edge: "project_contributions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldUsername, user.FieldAvatarURL:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldAvatarURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_url", values[i])
			} else if value.Valid {
				u.AvatarURL = new(string)
				*u.AvatarURL = value.String
			}
		}
	}
	return nil
}

// QueryDiscordAccounts queries the "discord_accounts" edge of the User entity.
func (u *User) QueryDiscordAccounts() *DiscordAccountQuery {
	return (&UserClient{config: u.config}).QueryDiscordAccounts(u)
}

// QueryGithubAccounts queries the "github_accounts" edge of the User entity.
func (u *User) QueryGithubAccounts() *GithubAccountQuery {
	return (&UserClient{config: u.config}).QueryGithubAccounts(u)
}

// QueryProjectContributions queries the "project_contributions" edge of the User entity.
func (u *User) QueryProjectContributions() *ProjectContributorQuery {
	return (&UserClient{config: u.config}).QueryProjectContributions(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", username=")
	builder.WriteString(u.Username)
	if v := u.AvatarURL; v != nil {
		builder.WriteString(", avatar_url=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
