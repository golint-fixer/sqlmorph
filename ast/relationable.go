package ast

// Relationable is an AST node with join table.
type Relationable interface {
	AddRelation(Relation)
}

type Relations []Relation

func (r *Relations) AddRelation(relation Relation) {
	*r = append(*r, relation)
}
