package ast

// WithRelations is an AST node with relations.
type WithRelations interface {

	// AddRelations adds a relation to the AST node.
	AddRelation(Relation)

	// GetRelations returns the relations of the AST node.
	GetRelations() []Relation
}

// Relations represents the relations of an AST node.
type Relations []Relation

// AddRelations adds a relation to the AST node.
func (r *Relations) AddRelation(relation Relation) {
	*r = append(*r, relation)
}

// GetRelations returns the relations of the AST node.
func (r *Relations) GetRelations() []Relation {
	return []Relation(*r)
}
