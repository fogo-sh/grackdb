// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/fogo-sh/grackdb/ent/githubaccount"
	"github.com/fogo-sh/grackdb/ent/githuborganization"
	"github.com/fogo-sh/grackdb/ent/project"
	"github.com/fogo-sh/grackdb/ent/repository"
)

// Repository is the model entity for the Repository schema.
type Repository struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RepositoryQuery when eager-loading is set.
	Edges                            RepositoryEdges `json:"edges"`
	github_account_repositories      *int
	github_organization_repositories *int
	project_repositories             *int
}

// RepositoryEdges holds the relations/edges for other nodes in the graph.
type RepositoryEdges struct {
	// Project holds the value of the project edge.
	Project *Project `json:"project,omitempty"`
	// GithubAccount holds the value of the github_account edge.
	GithubAccount *GithubAccount `json:"github_account,omitempty"`
	// GithubOrganization holds the value of the github_organization edge.
	GithubOrganization *GithubOrganization `json:"github_organization,omitempty"`
	// DiscordBots holds the value of the discord_bots edge.
	DiscordBots []*DiscordBot `json:"discord_bots,omitempty"`
	// Sites holds the value of the sites edge.
	Sites []*Site `json:"sites,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RepositoryEdges) ProjectOrErr() (*Project, error) {
	if e.loadedTypes[0] {
		if e.Project == nil {
			// The edge project was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: project.Label}
		}
		return e.Project, nil
	}
	return nil, &NotLoadedError{edge: "project"}
}

// GithubAccountOrErr returns the GithubAccount value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RepositoryEdges) GithubAccountOrErr() (*GithubAccount, error) {
	if e.loadedTypes[1] {
		if e.GithubAccount == nil {
			// The edge github_account was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: githubaccount.Label}
		}
		return e.GithubAccount, nil
	}
	return nil, &NotLoadedError{edge: "github_account"}
}

// GithubOrganizationOrErr returns the GithubOrganization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RepositoryEdges) GithubOrganizationOrErr() (*GithubOrganization, error) {
	if e.loadedTypes[2] {
		if e.GithubOrganization == nil {
			// The edge github_organization was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: githuborganization.Label}
		}
		return e.GithubOrganization, nil
	}
	return nil, &NotLoadedError{edge: "github_organization"}
}

// DiscordBotsOrErr returns the DiscordBots value or an error if the edge
// was not loaded in eager-loading.
func (e RepositoryEdges) DiscordBotsOrErr() ([]*DiscordBot, error) {
	if e.loadedTypes[3] {
		return e.DiscordBots, nil
	}
	return nil, &NotLoadedError{edge: "discord_bots"}
}

// SitesOrErr returns the Sites value or an error if the edge
// was not loaded in eager-loading.
func (e RepositoryEdges) SitesOrErr() ([]*Site, error) {
	if e.loadedTypes[4] {
		return e.Sites, nil
	}
	return nil, &NotLoadedError{edge: "sites"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Repository) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case repository.FieldID:
			values[i] = new(sql.NullInt64)
		case repository.FieldName, repository.FieldDescription:
			values[i] = new(sql.NullString)
		case repository.ForeignKeys[0]: // github_account_repositories
			values[i] = new(sql.NullInt64)
		case repository.ForeignKeys[1]: // github_organization_repositories
			values[i] = new(sql.NullInt64)
		case repository.ForeignKeys[2]: // project_repositories
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Repository", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Repository fields.
func (r *Repository) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case repository.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case repository.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		case repository.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				r.Description = new(string)
				*r.Description = value.String
			}
		case repository.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field github_account_repositories", value)
			} else if value.Valid {
				r.github_account_repositories = new(int)
				*r.github_account_repositories = int(value.Int64)
			}
		case repository.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field github_organization_repositories", value)
			} else if value.Valid {
				r.github_organization_repositories = new(int)
				*r.github_organization_repositories = int(value.Int64)
			}
		case repository.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field project_repositories", value)
			} else if value.Valid {
				r.project_repositories = new(int)
				*r.project_repositories = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryProject queries the "project" edge of the Repository entity.
func (r *Repository) QueryProject() *ProjectQuery {
	return (&RepositoryClient{config: r.config}).QueryProject(r)
}

// QueryGithubAccount queries the "github_account" edge of the Repository entity.
func (r *Repository) QueryGithubAccount() *GithubAccountQuery {
	return (&RepositoryClient{config: r.config}).QueryGithubAccount(r)
}

// QueryGithubOrganization queries the "github_organization" edge of the Repository entity.
func (r *Repository) QueryGithubOrganization() *GithubOrganizationQuery {
	return (&RepositoryClient{config: r.config}).QueryGithubOrganization(r)
}

// QueryDiscordBots queries the "discord_bots" edge of the Repository entity.
func (r *Repository) QueryDiscordBots() *DiscordBotQuery {
	return (&RepositoryClient{config: r.config}).QueryDiscordBots(r)
}

// QuerySites queries the "sites" edge of the Repository entity.
func (r *Repository) QuerySites() *SiteQuery {
	return (&RepositoryClient{config: r.config}).QuerySites(r)
}

// Update returns a builder for updating this Repository.
// Note that you need to call Repository.Unwrap() before calling this method if this Repository
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Repository) Update() *RepositoryUpdateOne {
	return (&RepositoryClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Repository entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Repository) Unwrap() *Repository {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Repository is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Repository) String() string {
	var builder strings.Builder
	builder.WriteString("Repository(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", name=")
	builder.WriteString(r.Name)
	if v := r.Description; v != nil {
		builder.WriteString(", description=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Repositories is a parsable slice of Repository.
type Repositories []*Repository

func (r Repositories) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
