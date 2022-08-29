// Code generated by ent, DO NOT EDIT.

package technology

import (
	"fmt"
	"io"
	"strconv"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the technology type in the database.
	Label = "technology"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldColour holds the string denoting the colour field in the database.
	FieldColour = "colour"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// EdgeParentTechnologies holds the string denoting the parent_technologies edge name in mutations.
	EdgeParentTechnologies = "parent_technologies"
	// EdgeChildTechnologies holds the string denoting the child_technologies edge name in mutations.
	EdgeChildTechnologies = "child_technologies"
	// EdgeProjects holds the string denoting the projects edge name in mutations.
	EdgeProjects = "projects"
	// EdgeRepositories holds the string denoting the repositories edge name in mutations.
	EdgeRepositories = "repositories"
	// Table holds the table name of the technology in the database.
	Table = "technologies"
	// ParentTechnologiesTable is the table that holds the parent_technologies relation/edge.
	ParentTechnologiesTable = "technology_associations"
	// ParentTechnologiesInverseTable is the table name for the TechnologyAssociation entity.
	// It exists in this package in order to avoid circular dependency with the "technologyassociation" package.
	ParentTechnologiesInverseTable = "technology_associations"
	// ParentTechnologiesColumn is the table column denoting the parent_technologies relation/edge.
	ParentTechnologiesColumn = "technology_parent_technologies"
	// ChildTechnologiesTable is the table that holds the child_technologies relation/edge.
	ChildTechnologiesTable = "technology_associations"
	// ChildTechnologiesInverseTable is the table name for the TechnologyAssociation entity.
	// It exists in this package in order to avoid circular dependency with the "technologyassociation" package.
	ChildTechnologiesInverseTable = "technology_associations"
	// ChildTechnologiesColumn is the table column denoting the child_technologies relation/edge.
	ChildTechnologiesColumn = "technology_child_technologies"
	// ProjectsTable is the table that holds the projects relation/edge.
	ProjectsTable = "project_technologies"
	// ProjectsInverseTable is the table name for the ProjectTechnology entity.
	// It exists in this package in order to avoid circular dependency with the "projecttechnology" package.
	ProjectsInverseTable = "project_technologies"
	// ProjectsColumn is the table column denoting the projects relation/edge.
	ProjectsColumn = "technology_projects"
	// RepositoriesTable is the table that holds the repositories relation/edge.
	RepositoriesTable = "repository_technologies"
	// RepositoriesInverseTable is the table name for the RepositoryTechnology entity.
	// It exists in this package in order to avoid circular dependency with the "repositorytechnology" package.
	RepositoriesInverseTable = "repository_technologies"
	// RepositoriesColumn is the table column denoting the repositories relation/edge.
	RepositoriesColumn = "technology_repositories"
)

// Columns holds all SQL columns for technology fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldColour,
	FieldType,
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

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeLibrary   Type = "LIBRARY"
	TypeLanguage  Type = "LANGUAGE"
	TypeAlgorithm Type = "ALGORITHM"
	TypeDatabase  Type = "DATABASE"
	TypeProtocol  Type = "PROTOCOL"
	TypeService   Type = "SERVICE"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeLibrary, TypeLanguage, TypeAlgorithm, TypeDatabase, TypeProtocol, TypeService:
		return nil
	default:
		return fmt.Errorf("technology: invalid enum value for type field: %q", _type)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Type) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Type) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Type(str)
	if err := TypeValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Type", str)
	}
	return nil
}
