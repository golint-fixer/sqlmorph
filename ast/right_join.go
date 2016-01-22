package ast

// RightJoin represents a right join table relation in the SQL query.
type RightJoin struct {
	Target

	Left  *Field
	Right *Field
}
