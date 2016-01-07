package parsing_test

import (
	"fmt"
	"reflect"

	"github.com/s2gatev/sqlmorph/ast"
	"github.com/s2gatev/sqlmorph/parsing"
)

func ExampleSelectParsing() {
	query := `SELECT u.Name FROM User u WHERE u.Age=21`

	expected := &ast.Select{
		Fields: []*ast.Field{
			&ast.Field{Target: "u", Name: "Name"},
		},
		Conditions: []*ast.EqualsCondition{
			&ast.EqualsCondition{
				Field: &ast.Field{Target: "u", Name: "Age"},
				Value: "21",
			},
		},
		Table: &ast.Table{Name: "User", Alias: "u"},
	}
	actual, _ := parsing.NewParser().ParseString(query)

	fmt.Print(reflect.DeepEqual(actual, expected))

	// Output:
	// true
}
