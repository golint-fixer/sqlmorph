package ast

// Select represents a SQL SELECT statement.
type Select struct {
	Target
	Limit
	Offset
	Relations
	Filters
	Fields
}
