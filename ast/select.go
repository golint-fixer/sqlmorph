package ast

// Select represents a SQL SELECT statement.
type Select struct {
	Target
	Limit
	Offset
	Relations
	Filters
	Fields
}

// BuildQuery turns the node into an SQL statement.
func (ss *Select) BuildQuery() string {
	query := ""

	// Build SELECT part.
	fieldsPart := ""
	for index, field := range ss.Fields {
		if index != 0 {
			fieldsPart += ", "
		}
		fieldsPart += field.BuildQuery()
	}
	query += "SELECT " + fieldsPart

	// Build FROM part.
	query += " FROM " + ss.GetTarget().BuildQuery()

	for _, join := range ss.Relations {
		query += " " + join.BuildQuery()
	}

	// Build WHERE part.
	if len(ss.Filters) > 0 {
		jointFilters := ""
		for index, filter := range ss.Filters {
			if index != 0 {
				jointFilters += " AND "
			}
			jointFilters += filter.BuildQuery()
		}
		query += " WHERE " + jointFilters
	}

	// Build LIMIT part.
	if ss.GetLimit() != "" {
		query += " LIMIT " + ss.GetLimit()
	}

	// Build OFFSET part.
	if ss.GetOffset() != "" {
		query += " OFFSET " + ss.GetOffset()
	}
	return query
}
