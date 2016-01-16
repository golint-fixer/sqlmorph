package ast

// WithFields is an AST nodes with fields.
type WithFields interface {

	// AddField adds a field to the node.
	AddField(*Field)

	// GetFields returns the fields of the node.
	GetFields() []*Field
}

// Fields represents the fields of an AST node.
type Fields []*Field

// AddField adds a field to the node.
func (f *Fields) AddField(field *Field) {
	*f = append(*f, field)
}

// GetFields returns the fields of the node.
func (f *Fields) GetFields() []*Field {
	return []*Field(*f)
}
