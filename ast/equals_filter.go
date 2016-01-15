package ast

// EqualsFilter represents an equality field-value relation.
type EqualsFilter struct {
	Field *Field
	Value string
}

func (c *EqualsFilter) BuildQuery() string {
	return c.Field.BuildQuery() + "=" + c.Value
}
