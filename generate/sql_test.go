package generate_test

import (
	"testing"

	. "github.com/s2gatev/sqlmorph/ast"
	. "github.com/s2gatev/sqlmorph/generate"
)

func TestGenerateSelect(t *testing.T) {
	runTests(t, []test{
		{
			AST: &Select{
				Fields: []*Field{
					&Field{Name: "Name", Target: "u"},
					&Field{Name: "Location", Target: "a"},
				},
				Target: Target{Name: "User", Alias: "u"},
				Relations: Relations{
					&RightJoin{
						Target: Target{Name: "Address", Alias: "a"},
						Left:   &Field{Name: "ID", Target: "u"},
						Right:  &Field{Name: "UserID", Target: "a"},
					},
				},
			},
			Expected: "SELECT u.Name, a.Location FROM User u RIGHT JOIN Address a ON u.ID=a.UserID",
		},
	})
}

type test struct {
	AST      Node
	Expected string
}

func runTests(t *testing.T, tests []test) {
	generator := SQLGenerator{}

	execute := func(test test) {
		actual := generator.Generate(test.AST)
		if actual != test.Expected {
			t.Errorf("Generated query is not correct.\n"+
				"Expected: %v\n"+
				"Actual: %v\n", test.Expected, actual)
		}
	}

	for _, test := range tests {
		execute(test)
	}
}
