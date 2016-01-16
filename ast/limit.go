package ast

// WithLimit is an AST node with limit value.
type WithLimit interface {

	// SetLimit sets the limit of the node.
	SetLimit(string)

	// GetLimit returns the limit of the node.
	GetLimit() string
}

// Limit is a limit of an AST node.
type Limit string

// SetLimit sets the limit of the node.
func (l *Limit) SetLimit(limit string) {
	*l = Limit(limit)
}

// GetLimit returns the limit of the node.
func (l *Limit) GetLimit() string {
	return string(*l)
}
