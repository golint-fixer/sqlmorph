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

// BuildQuery turns the node into an SQL statement.
func (t *Target) BuildQuery() string {
	query := t.Name
	if t.Alias != "" {
		query += " " + t.Alias
	}
	return query
}

// SetTarget sets the target of the node.
func (t *Target) SetTarget(target *Target) {
	*t = *target
}

// GetTarget returns the target of the node.
func (t *Target) GetTarget() *Target {
	return t
}
