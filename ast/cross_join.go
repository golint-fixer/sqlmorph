package ast

// CrossJoin represents an inner join table relation in the SQL query.
type CrossJoin struct {
	Target
}

func (j *CrossJoin) BuildQuery() string {
	return "CROSS JOIN " + j.GetTarget().BuildQuery()
}
