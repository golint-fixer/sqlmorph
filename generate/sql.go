package generate

import (
	"bytes"

	"github.com/s2gatev/sqlmorph/ast"
)

type SQLGenerator struct{}

func (s *SQLGenerator) Generate(node ast.Node) string {
	var content bytes.Buffer
	switch c := node.(type) {
	case *ast.Select:
		s.generateSelect(&content, c)
	case *ast.Update:
		s.generateUpdate(&content, c)
	case *ast.Delete:
		s.generateDelete(&content, c)
	case *ast.Field:
		s.generateField(&content, c)
	case *ast.EqualsFilter:
		s.generateEqualsFilter(&content, c)
	case *ast.InnerJoin:
		s.generateInnerJoin(&content, c)
	case *ast.LeftJoin:
		s.generateLeftJoin(&content, c)
	case *ast.RightJoin:
		s.generateRightJoin(&content, c)
	case *ast.Target:
		s.generateTarget(&content, c)
	}
	return content.String()
}

func (s *SQLGenerator) generateSelect(content *bytes.Buffer, n *ast.Select) {
	content.WriteString("SELECT ")
	for index, field := range n.Fields {
		if index > 0 {
			content.WriteString(", ")
		}
		s.generateField(content, field)
	}

	// Build FROM part.
	content.WriteString(" FROM ")
	s.generateTarget(content, n.GetTarget())

	for _, join := range n.Relations {
		content.WriteString(" " + s.Generate(join))
	}

	// Build WHERE part.
	if len(n.Filters) > 0 {
		jointFilters := ""
		for index, filter := range n.Filters {
			if index != 0 {
				jointFilters += " AND "
			}
			jointFilters += s.Generate(filter)
		}
		content.WriteString(" WHERE " + jointFilters)
	}

	// Build LIMIT part.
	if n.GetLimit() != "" {
		content.WriteString(" LIMIT " + n.GetLimit())
	}

	// Build OFFSET part.
	if n.GetOffset() != "" {
		content.WriteString(" OFFSET " + n.GetOffset())
	}
}

func (s *SQLGenerator) generateUpdate(content *bytes.Buffer, n *ast.Update) {
	content.WriteString("UPDATE " + s.Generate(n.GetTarget()))

	fieldsPart := ""
	for index, field := range n.Fields {
		if index != 0 {
			fieldsPart += ", "
		}
		fieldsPart += s.Generate(field)
	}
	content.WriteString(" SET " + fieldsPart)

	// Build WHERE part.
	if len(n.Filters) > 0 {
		jointFilters := ""
		for index, filter := range n.Filters {
			if index != 0 {
				jointFilters += " AND "
			}
			jointFilters += s.Generate(filter)
		}
		content.WriteString(" WHERE " + jointFilters)
	}
}

func (s *SQLGenerator) generateDelete(content *bytes.Buffer, n *ast.Delete) {
	content.WriteString("DELETE FROM ")
	content.WriteString(s.Generate(n.GetTarget()))

	if len(n.Filters) > 0 {
		jointFilters := ""
		for index, filter := range n.Filters {
			if index != 0 {
				jointFilters += " AND "
			}
			jointFilters += s.Generate(filter)
		}
		content.WriteString(" WHERE " + jointFilters)
	}
}

func (s *SQLGenerator) generateField(content *bytes.Buffer, n *ast.Field) {
	if n.Target != "" {
		content.WriteString(n.Target + ".")
	}
	content.WriteString(n.Name)
	if n.Value != "" {
		content.WriteString("=" + n.Value)
	}
}

func (s *SQLGenerator) generateEqualsFilter(content *bytes.Buffer, n *ast.EqualsFilter) {
	content.WriteString(s.Generate(n.Field))
	content.WriteString("=")
	content.WriteString(n.Value)
}

func (s *SQLGenerator) generateInnerJoin(content *bytes.Buffer, n *ast.InnerJoin) {
	content.WriteString("INNER JOIN ")
	content.WriteString(s.Generate(n.GetTarget()))
	content.WriteString(" ON ")
	content.WriteString(s.Generate(n.Left))
	content.WriteString("=")
	content.WriteString(s.Generate(n.Right))
}

func (s *SQLGenerator) generateLeftJoin(content *bytes.Buffer, n *ast.LeftJoin) {
	content.WriteString("LEFT JOIN ")
	content.WriteString(s.Generate(n.GetTarget()))
	content.WriteString(" ON ")
	content.WriteString(s.Generate(n.Left))
	content.WriteString("=")
	content.WriteString(s.Generate(n.Right))
}

func (s *SQLGenerator) generateRightJoin(content *bytes.Buffer, n *ast.RightJoin) {
	content.WriteString("RIGHT JOIN ")
	content.WriteString(s.Generate(n.GetTarget()))
	content.WriteString(" ON ")
	content.WriteString(s.Generate(n.Left))
	content.WriteString("=")
	content.WriteString(s.Generate(n.Right))
}

func (s *SQLGenerator) generateTarget(content *bytes.Buffer, n *ast.Target) {
	content.WriteString(n.Name)
	if n.Alias != "" {
		content.WriteString(" " + n.Alias)
	}
}
