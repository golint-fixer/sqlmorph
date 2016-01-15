package ast

// LeftJoin represents an inner join table relation in the SQL query.
type LeftJoin struct {
	Target

	Left  *Field
	Right *Field
}

func (j *LeftJoin) BuildQuery() string {
	return "LEFT JOIN " + j.GetTarget().BuildQuery() + " ON " +
		j.Left.BuildQuery() + "=" + j.Right.BuildQuery()
}
