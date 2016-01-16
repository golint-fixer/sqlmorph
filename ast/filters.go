package ast

// WithFilters is an AST node with filters.
type WithFilters interface {

	// AddFilter adds a filter to the node.
	AddFilter(Filter)

	// GetFilters returns the filters of the node.
	GetFilters() []Filter
}

// Filters represents the filters of an AST node.
type Filters []Filter

// AddFilter adds a filter to the node.
func (f *Filters) AddFilter(filter Filter) {
	*f = append(*f, filter)
}

// GetFilters returns the filters of the node.
func (f *Filters) GetFilters() []Filter {
	return []Filter(*f)
}
