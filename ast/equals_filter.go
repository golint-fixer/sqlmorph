package ast

// EqualsFilter represents an equality filter.
type EqualsFilter struct {
	Field *Field
	Value string
}

// BuildQuery turns the node into an SQL statement.
func (c *EqualsFilter) BuildQuery() string {
	return c.Field.BuildQuery() + "=" + c.Value
}
