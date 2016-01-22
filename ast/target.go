package ast

// WithTarget is an AST node with table.
type WithTarget interface {

	// SetTarget sets the target of the node.
	SetTarget(*Target)

	// GetTarget returns the target of the node.
	GetTarget() *Target
}

// Target represents the target of an AST node.
type Target struct {
	Name  string
	Alias string
}

// SetTarget sets the target of the node.
func (t *Target) SetTarget(target *Target) {
	*t = *target
}

// GetTarget returns the target of the node.
func (t *Target) GetTarget() *Target {
	return t
}
