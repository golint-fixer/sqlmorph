package ast

// Container is an AST nodes with children.
type Container interface {

	//  AddField adds a Field to the node.
	AddField(*Field)

	GetFields() []*Field
}

type Fields []*Field

func (f *Fields) AddField(field *Field) {
	*f = append(*f, field)
}

func (f *Fields) GetFields() []*Field {
	return []*Field(*f)
}
