package ast

type Visitor struct{}

func NewVisitor() *Visitor {
	return &Visitor{}
}

func (v *Visitor) Visit(node Node, handle func(Node)) {
	handle(node)

	switch concrete := node.(type) {
	case *Select:
		for _, field := range concrete.Fields {
			v.Visit(field, handle)
		}
		for _, join := range concrete.JoinTables {
			v.Visit(join, handle)
		}
		for _, condition := range concrete.Conditions {
			v.Visit(condition, handle)
		}
	case *Update:
		for _, field := range concrete.Fields {
			v.Visit(field, handle)
		}
		for _, condition := range concrete.Conditions {
			v.Visit(condition, handle)
		}
	case *Delete:
		for _, condition := range concrete.Conditions {
			v.Visit(condition, handle)
		}
	case *EqualsCondition:
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
