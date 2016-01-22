package ast

// LeftJoin represents a left join table relation in the SQL query.
type LeftJoin struct {
	Target

	Left  *Field
	Right *Field
}
