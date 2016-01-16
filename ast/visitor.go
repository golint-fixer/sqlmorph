package ast

// Visitor is an SQL AST node visitor.
type Visitor struct{}

// NewVisitor creates a new SQL AST node visitor.
func NewVisitor() *Visitor {
	return &Visitor{}
}

// Visit visits an SQL AST node and passes it to a handle.
func (v *Visitor) Visit(node Node, handle func(Node)) {
	handle(node)

	switch concrete := node.(type) {
	case *Select:
		for _, field := range concrete.Fields {
			v.Visit(field, handle)
		}
		for _, join := range concrete.Relations {
			v.Visit(join, handle)
		}
		for _, filter := range concrete.Filters {
			v.Visit(filter, handle)
		}
	case *Update:
		for _, field := range concrete.Fields {
			v.Visit(field, handle)
		}
		for _, filter := range concrete.Filters {
			v.Visit(filter, handle)
		}
	case *Delete:
		for _, filter := range concrete.Filters {
			v.Visit(filter, handle)
		}
	case *EqualsFilter:
		v.Visit(concrete.Field, handle)
	case *InnerJoin:
		v.Visit(concrete.Left, handle)
		v.Visit(concrete.Right, handle)
	case *LeftJoin:
		v.Visit(concrete.Left, handle)
		v.Visit(concrete.Right, handle)
	case *RightJoin:
		v.Visit(concrete.Left, handle)
		v.Visit(concrete.Right, handle)
	}
}
