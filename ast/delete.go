package ast

// Delete represents an DELETE SQL query AST node.
type Delete struct {
	Target
	Filters
}

func NewDelete() *Delete {
	return &Delete{}
}

func (d *Delete) BuildQuery() string {
	query := "DELETE FROM " + d.GetTarget().BuildQuery()

	if len(d.Filters) > 0 {
		jointFilters := ""
		for index, filter := range d.Filters {
			if index != 0 {
				jointFilters += " AND "
			}
			jointFilters += filter.BuildQuery()
		}
		query += " WHERE " + jointFilters
	}

	return query
}
