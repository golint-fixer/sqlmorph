package ast

// RightJoin represents an inner join table relation in the SQL query.
type RightJoin struct {
	Target

	Left  *Field
	Right *Field
}

func (j *RightJoin) BuildQuery() string {
	return "RIGHT JOIN " + j.GetTarget().BuildQuery() + " ON " +
		j.Left.BuildQuery() + "=" + j.Right.BuildQuery()
}
