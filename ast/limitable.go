package ast

// Limitable is an AST node with limit value.
type Limitable interface {

	// SetLimit sets the limit of the node.
	SetLimit(string)

	// GetLimit gets the limit of the node.
	GetLimit() string
}

// Limit is a limit property of an AST node.
type Limit string

func (l *Limit) SetLimit(limit string) {
	*l = Limit(limit)
}

func (l *Limit) GetLimit() string {
	return string(*l)
}
