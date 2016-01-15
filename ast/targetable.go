package ast

// Targetable is an AST node with table.
type Targetable interface {

	// SetTarget sets the target of the node.
	SetTarget(*Target)

	// GetTarget gets the target of the node.
	GetTarget() *Target
}

type Target struct {
	Name  string
	Alias string
}

func (t *Target) BuildQuery() string {
	query := t.Name
	if t.Alias != "" {
		query += " " + t.Alias
	}
	return query
}

func (t *Target) SetTarget(target *Target) {
	*t = *target
}

func (t *Target) GetTarget() *Target {
	return t
}
