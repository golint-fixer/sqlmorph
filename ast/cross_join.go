package ast

// CrossJoin represents an cross join table relation in the SQL query.
type CrossJoin struct {
	Target
}

// BuildQuery turns the node into an SQL statement.
func (j *CrossJoin) BuildQuery() string {
	return "CROSS JOIN " + j.GetTarget().BuildQuery()
}
