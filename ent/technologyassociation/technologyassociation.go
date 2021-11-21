// Code generated by entc, DO NOT EDIT.

package technologyassociation

import (
	"fmt"
	"io"
	"strconv"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the technologyassociation type in the database.
	Label = "technology_association"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChild holds the string denoting the child edge name in mutations.
	EdgeChild = "child"
	// Table holds the table name of the technologyassociation in the database.
	Table = "technology_associations"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "technology_associations"
	// ParentInverseTable is the table name for the Technology entity.
	// It exists in this package in order to avoid circular dependency with the "technology" package.
	ParentInverseTable = "technologies"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "technology_child_technologies"
	// ChildTable is the table that holds the child relation/edge.
	ChildTable = "technology_associations"
	// ChildInverseTable is the table name for the Technology entity.
	// It exists in this package in order to avoid circular dependency with the "technology" package.
	ChildInverseTable = "technologies"
	// ChildColumn is the table column denoting the child relation/edge.
	ChildColumn = "technology_parent_technologies"
)

// Columns holds all SQL columns for technologyassociation fields.
var Columns = []string{
	FieldID,
	FieldType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "technology_associations"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"technology_parent_technologies",
	"technology_child_technologies",
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

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeWrittenIn  Type = "WRITTEN_IN"
	TypeImplements Type = "IMPLEMENTS"
	TypeUses       Type = "USES"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeWrittenIn, TypeImplements, TypeUses:
		return nil
	default:
		return fmt.Errorf("technologyassociation: invalid enum value for type field: %q", _type)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (_type Type) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(_type.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (_type *Type) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*_type = Type(str)
	if err := TypeValidator(*_type); err != nil {
		return fmt.Errorf("%s is not a valid Type", str)
	}
	return nil
}
