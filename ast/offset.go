package ast

// WithOffset is an AST node with offset.
type WithOffset interface {

	// SetOffset sets the offset of the node.
	SetOffset(string)

	// GetOffset returns the offset of the node.
	GetOffset() string
}

// Offset is an offset of an AST node.
type Offset string

// SetOffset sets the offset of the node.
func (o *Offset) SetOffset(offset string) {
	*o = Offset(offset)
}

// GetOffset returns the offset of the node.
func (o *Offset) GetOffset() string {
	return string(*o)
}
