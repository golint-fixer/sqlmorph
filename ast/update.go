package ast

// Update represents an Update SQL query AST node.
type Update struct {
	Target
	Filters
	Fields
}

// BuildQuery turns the node into an SQL statement.
func (u *Update) BuildQuery() string {
	query := "UPDATE " + u.GetTarget().BuildQuery()

	fieldsPart := ""
	for index, field := range u.Fields {
		if index != 0 {
			fieldsPart += ", "
		}
		fieldsPart += field.BuildQuery()
	}
	query += " SET " + fieldsPart

	// Build WHERE part.
	if len(u.Filters) > 0 {
		jointFilters := ""
		for index, filter := range u.Filters {
			if index != 0 {
				jointFilters += " AND "
			}
			jointFilters += filter.BuildQuery()
		}
		query += " WHERE " + jointFilters
	}

	return query
}
