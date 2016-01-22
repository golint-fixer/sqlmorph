package parse_test

import (
	"fmt"
	"reflect"

	"github.com/s2gatev/sqlmorph/ast"
	"github.com/s2gatev/sqlmorph/parse"
)

func ExampleSelectParsing() {
	query := `SELECT u.Name FROM User u WHERE u.Age=21`

	expected := &ast.Select{
		Fields: []*ast.Field{
			&ast.Field{Target: "u", Name: "Name"},
		},
		Filters: ast.Filters{
			&ast.EqualsFilter{
				Field: &ast.Field{Target: "u", Name: "Age"},
				Value: "21",
			},
		},
		Target: ast.Target{Name: "User", Alias: "u"},
	}
	actual, _ := parse.NewParser().ParseString(query)

	fmt.Print(reflect.DeepEqual(actual, expected))

	// Output:
	// true
}
