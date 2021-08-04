// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/fogo-sh/grackdb/ent/project"
)

// Project is the model entity for the Project schema.
type Project struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
	// StartDate holds the value of the "start_date" field.
	StartDate time.Time `json:"start_date,omitempty"`
	// EndDate holds the value of the "end_date" field.
	EndDate *time.Time `json:"end_date,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProjectQuery when eager-loading is set.
	Edges ProjectEdges `json:"edges"`
}

// ProjectEdges holds the relations/edges for other nodes in the graph.
type ProjectEdges struct {
	// Contributors holds the value of the contributors edge.
	Contributors []*ProjectContributor `json:"contributors,omitempty"`
	// ParentProjects holds the value of the parent_projects edge.
	ParentProjects []*ProjectAssociation `json:"parent_projects,omitempty"`
	// ChildProjects holds the value of the child_projects edge.
	ChildProjects []*ProjectAssociation `json:"child_projects,omitempty"`
	// Repositories holds the value of the repositories edge.
	Repositories []*Repository `json:"repositories,omitempty"`
	// DiscordBots holds the value of the discord_bots edge.
	DiscordBots []*DiscordBot `json:"discord_bots,omitempty"`
	// Sites holds the value of the sites edge.
	Sites []*Site `json:"sites,omitempty"`
	// Technologies holds the value of the technologies edge.
	Technologies []*ProjectTechnology `json:"technologies,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [7]bool
}

// ContributorsOrErr returns the Contributors value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) ContributorsOrErr() ([]*ProjectContributor, error) {
	if e.loadedTypes[0] {
		return e.Contributors, nil
	}
	return nil, &NotLoadedError{edge: "contributors"}
}

// ParentProjectsOrErr returns the ParentProjects value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) ParentProjectsOrErr() ([]*ProjectAssociation, error) {
	if e.loadedTypes[1] {
		return e.ParentProjects, nil
	}
	return nil, &NotLoadedError{edge: "parent_projects"}
}

// ChildProjectsOrErr returns the ChildProjects value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) ChildProjectsOrErr() ([]*ProjectAssociation, error) {
	if e.loadedTypes[2] {
		return e.ChildProjects, nil
	}
	return nil, &NotLoadedError{edge: "child_projects"}
}

// RepositoriesOrErr returns the Repositories value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) RepositoriesOrErr() ([]*Repository, error) {
	if e.loadedTypes[3] {
		return e.Repositories, nil
	}
	return nil, &NotLoadedError{edge: "repositories"}
}

// DiscordBotsOrErr returns the DiscordBots value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) DiscordBotsOrErr() ([]*DiscordBot, error) {
	if e.loadedTypes[4] {
		return e.DiscordBots, nil
	}
	return nil, &NotLoadedError{edge: "discord_bots"}
}

// SitesOrErr returns the Sites value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) SitesOrErr() ([]*Site, error) {
	if e.loadedTypes[5] {
		return e.Sites, nil
	}
	return nil, &NotLoadedError{edge: "sites"}
}

// TechnologiesOrErr returns the Technologies value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) TechnologiesOrErr() ([]*ProjectTechnology, error) {
	if e.loadedTypes[6] {
		return e.Technologies, nil
	}
	return nil, &NotLoadedError{edge: "technologies"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Project) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case project.FieldID:
			values[i] = new(sql.NullInt64)
		case project.FieldName, project.FieldDescription:
			values[i] = new(sql.NullString)
		case project.FieldStartDate, project.FieldEndDate:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Project", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Project fields.
func (pr *Project) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case project.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case project.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case project.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pr.Description = new(string)
				*pr.Description = value.String
			}
		case project.FieldStartDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_date", values[i])
			} else if value.Valid {
				pr.StartDate = value.Time
			}
		case project.FieldEndDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_date", values[i])
			} else if value.Valid {
				pr.EndDate = new(time.Time)
				*pr.EndDate = value.Time
			}
		}
	}
	return nil
}

// QueryContributors queries the "contributors" edge of the Project entity.
func (pr *Project) QueryContributors() *ProjectContributorQuery {
	return (&ProjectClient{config: pr.config}).QueryContributors(pr)
}

// QueryParentProjects queries the "parent_projects" edge of the Project entity.
func (pr *Project) QueryParentProjects() *ProjectAssociationQuery {
	return (&ProjectClient{config: pr.config}).QueryParentProjects(pr)
}

// QueryChildProjects queries the "child_projects" edge of the Project entity.
func (pr *Project) QueryChildProjects() *ProjectAssociationQuery {
	return (&ProjectClient{config: pr.config}).QueryChildProjects(pr)
}

// QueryRepositories queries the "repositories" edge of the Project entity.
func (pr *Project) QueryRepositories() *RepositoryQuery {
	return (&ProjectClient{config: pr.config}).QueryRepositories(pr)
}

// QueryDiscordBots queries the "discord_bots" edge of the Project entity.
func (pr *Project) QueryDiscordBots() *DiscordBotQuery {
	return (&ProjectClient{config: pr.config}).QueryDiscordBots(pr)
}

// QuerySites queries the "sites" edge of the Project entity.
func (pr *Project) QuerySites() *SiteQuery {
	return (&ProjectClient{config: pr.config}).QuerySites(pr)
}

// QueryTechnologies queries the "technologies" edge of the Project entity.
func (pr *Project) QueryTechnologies() *ProjectTechnologyQuery {
	return (&ProjectClient{config: pr.config}).QueryTechnologies(pr)
}

// Update returns a builder for updating this Project.
// Note that you need to call Project.Unwrap() before calling this method if this Project
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Project) Update() *ProjectUpdateOne {
	return (&ProjectClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Project entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Project) Unwrap() *Project {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Project is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Project) String() string {
	var builder strings.Builder
	builder.WriteString("Project(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", name=")
	builder.WriteString(pr.Name)
	if v := pr.Description; v != nil {
		builder.WriteString(", description=")
		builder.WriteString(*v)
	}
	builder.WriteString(", start_date=")
	builder.WriteString(pr.StartDate.Format(time.ANSIC))
	if v := pr.EndDate; v != nil {
		builder.WriteString(", end_date=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Projects is a parsable slice of Project.
type Projects []*Project

func (pr Projects) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
