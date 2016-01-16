package ast

// Node is a generic representation of an element in the SQL AST.
type Node interface {

	// BuildQuery turns the node into an SQL statement.
	BuildQuery() string
}
