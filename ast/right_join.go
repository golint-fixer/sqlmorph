package ast

// RightJoin represents a right join table relation in the SQL query.
type RightJoin struct {
	Target

	Left  *Field
	Right *Field
}

// BuildQuery turns the node into an SQL statement.
func (j *RightJoin) BuildQuery() string {
	return "RIGHT JOIN " + j.GetTarget().BuildQuery() + " ON " +
		j.Left.BuildQuery() + "=" + j.Right.BuildQuery()
}
