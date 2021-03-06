// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/fogo-sh/grackdb/ent/technology"
)

// Technology is the model entity for the Technology schema.
type Technology struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
	// Colour holds the value of the "colour" field.
	Colour *string `json:"colour,omitempty"`
	// Type holds the value of the "type" field.
	Type technology.Type `json:"type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TechnologyQuery when eager-loading is set.
	Edges TechnologyEdges `json:"edges"`
}

// TechnologyEdges holds the relations/edges for other nodes in the graph.
type TechnologyEdges struct {
	// ParentTechnologies holds the value of the parent_technologies edge.
	ParentTechnologies []*TechnologyAssociation `json:"parent_technologies,omitempty"`
	// ChildTechnologies holds the value of the child_technologies edge.
	ChildTechnologies []*TechnologyAssociation `json:"child_technologies,omitempty"`
	// Projects holds the value of the projects edge.
	Projects []*ProjectTechnology `json:"projects,omitempty"`
	// Repositories holds the value of the repositories edge.
	Repositories []*RepositoryTechnology `json:"repositories,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ParentTechnologiesOrErr returns the ParentTechnologies value or an error if the edge
// was not loaded in eager-loading.
func (e TechnologyEdges) ParentTechnologiesOrErr() ([]*TechnologyAssociation, error) {
	if e.loadedTypes[0] {
		return e.ParentTechnologies, nil
	}
	return nil, &NotLoadedError{edge: "parent_technologies"}
}

// ChildTechnologiesOrErr returns the ChildTechnologies value or an error if the edge
// was not loaded in eager-loading.
func (e TechnologyEdges) ChildTechnologiesOrErr() ([]*TechnologyAssociation, error) {
	if e.loadedTypes[1] {
		return e.ChildTechnologies, nil
	}
	return nil, &NotLoadedError{edge: "child_technologies"}
}

// ProjectsOrErr returns the Projects value or an error if the edge
// was not loaded in eager-loading.
func (e TechnologyEdges) ProjectsOrErr() ([]*ProjectTechnology, error) {
	if e.loadedTypes[2] {
		return e.Projects, nil
	}
	return nil, &NotLoadedError{edge: "projects"}
}

// RepositoriesOrErr returns the Repositories value or an error if the edge
// was not loaded in eager-loading.
func (e TechnologyEdges) RepositoriesOrErr() ([]*RepositoryTechnology, error) {
	if e.loadedTypes[3] {
		return e.Repositories, nil
	}
	return nil, &NotLoadedError{edge: "repositories"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Technology) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case technology.FieldID:
			values[i] = new(sql.NullInt64)
		case technology.FieldName, technology.FieldDescription, technology.FieldColour, technology.FieldType:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Technology", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Technology fields.
func (t *Technology) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case technology.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case technology.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case technology.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				t.Description = new(string)
				*t.Description = value.String
			}
		case technology.FieldColour:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field colour", values[i])
			} else if value.Valid {
				t.Colour = new(string)
				*t.Colour = value.String
			}
		case technology.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				t.Type = technology.Type(value.String)
			}
		}
	}
	return nil
}

// QueryParentTechnologies queries the "parent_technologies" edge of the Technology entity.
func (t *Technology) QueryParentTechnologies() *TechnologyAssociationQuery {
	return (&TechnologyClient{config: t.config}).QueryParentTechnologies(t)
}

// QueryChildTechnologies queries the "child_technologies" edge of the Technology entity.
func (t *Technology) QueryChildTechnologies() *TechnologyAssociationQuery {
	return (&TechnologyClient{config: t.config}).QueryChildTechnologies(t)
}

// QueryProjects queries the "projects" edge of the Technology entity.
func (t *Technology) QueryProjects() *ProjectTechnologyQuery {
	return (&TechnologyClient{config: t.config}).QueryProjects(t)
}

// QueryRepositories queries the "repositories" edge of the Technology entity.
func (t *Technology) QueryRepositories() *RepositoryTechnologyQuery {
	return (&TechnologyClient{config: t.config}).QueryRepositories(t)
}

// Update returns a builder for updating this Technology.
// Note that you need to call Technology.Unwrap() before calling this method if this Technology
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Technology) Update() *TechnologyUpdateOne {
	return (&TechnologyClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Technology entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Technology) Unwrap() *Technology {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Technology is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Technology) String() string {
	var builder strings.Builder
	builder.WriteString("Technology(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", name=")
	builder.WriteString(t.Name)
	if v := t.Description; v != nil {
		builder.WriteString(", description=")
		builder.WriteString(*v)
	}
	if v := t.Colour; v != nil {
		builder.WriteString(", colour=")
		builder.WriteString(*v)
	}
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", t.Type))
	builder.WriteByte(')')
	return builder.String()
}

// Technologies is a parsable slice of Technology.
type Technologies []*Technology

func (t Technologies) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
