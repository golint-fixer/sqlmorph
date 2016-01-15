package ast

// Offsetable is an AST node with offset value.
type Offsetable interface {

	// SetOffset sets the offset value of the node.
	SetOffset(string)

	// GetOffset gets the offset value of the node.
	GetOffset() string
}

// Offset is an offset property of an AST node.
type Offset string

func (o *Offset) SetOffset(offset string) {
	*o = Offset(offset)
}

func (o *Offset) GetOffset() string {
	return string(*o)
}
