package ast

// InnerJoin represents an inner join table relation in the SQL query.
type InnerJoin struct {
	Target

	Left  *Field
	Right *Field
}
