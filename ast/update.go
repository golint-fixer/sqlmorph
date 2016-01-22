package ast

// Update represents an Update SQL query AST node.
type Update struct {
	Target
	Filters
	Fields
}
