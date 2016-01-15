package ast

// Filterable is an AST node with filters.
type Filterable interface {

	// AddFilter adds a filter to the node.
	AddFilter(Filter)

	// GetFilters gets the filters of the node.
	GetFilters() []Filter
}

type Filter interface {
	Node
}

type Filters []Filter

func (f *Filters) AddFilter(filter Filter) {
	*f = append(*f, filter)
}

func (f *Filters) GetFilters() []Filter {
	return []Filter(*f)
}
