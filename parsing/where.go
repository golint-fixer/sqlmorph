package parsing

import (
	"github.com/s2gatev/sqlmorph/ast"
)

const WhereWithoutConditionsError = "WHERE statement must be followed by condition list."

// WhereState parses WHERE SQL clauses along with their conditions.
// ... WHERE u.Age=? ...
type WhereState struct {
	BaseState
}

func (s *WhereState) Name() string {
	return "WHERE"
}

// Parse tries to parse a WHERE token from the tokenizer.
func (s *WhereState) Parse(result ast.Node, tokenizer *Tokenizer) (ast.Node, bool) {
	concrete := result.(ast.WithFilters)

	if token, _ := tokenizer.ReadToken(); token != WHERE {
		tokenizer.UnreadToken()
		return result, false
	}

	// Parse WHERE conditions.
	for {
		filter := &ast.EqualsFilter{}

		if token, field := tokenizer.ReadToken(); token == LITERAL {
			filter.Field = parseField(field)
		} else {
			wrongTokenPanic(WhereWithoutConditionsError, field)
		}

		if token, operator := tokenizer.ReadToken(); token != EQUALS {
			wrongTokenPanic(WhereWithoutConditionsError, operator)
		}

		if token, value := tokenizer.ReadToken(); token == LITERAL || token == PLACEHOLDER {
			filter.Value = value
		} else {
			wrongTokenPanic(WhereWithoutConditionsError, value)
		}

		concrete.AddFilter(filter)

		if token, _ := tokenizer.ReadToken(); token != AND {
			tokenizer.UnreadToken()
			break
		}
	}

	return result, true
}
