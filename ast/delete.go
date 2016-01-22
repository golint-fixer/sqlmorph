package ast

// Delete represents a DELETE SQL query AST node.
type Delete struct {
	Target
	Filters
}
