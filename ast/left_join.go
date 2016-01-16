package ast

// LeftJoin represents a left join table relation in the SQL query.
type LeftJoin struct {
	Target

	Left  *Field
	Right *Field
}

// BuildQuery turns the node into an SQL statement.
func (j *LeftJoin) BuildQuery() string {
	return "LEFT JOIN " + j.GetTarget().BuildQuery() + " ON " +
		j.Left.BuildQuery() + "=" + j.Right.BuildQuery()
}
