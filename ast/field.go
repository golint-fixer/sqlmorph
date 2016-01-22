package ast

// Field represents a struct field in the SQL query.
type Field struct {
	Target string
	Name   string
	Value  string
}
