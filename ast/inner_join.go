package ast

// InnerJoin represents an inner join table relation in the SQL query.
type InnerJoin struct {
	Target

	Left  *Field
	Right *Field
}

func (j *InnerJoin) BuildQuery() string {
	return "INNER JOIN " + j.GetTarget().BuildQuery() + " ON " +
		j.Left.BuildQuery() + "=" + j.Right.BuildQuery()
}
