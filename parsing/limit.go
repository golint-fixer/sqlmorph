package parsing

import (
	"github.com/s2gatev/sqlmorph/ast"
)

const LimitWithoutNumberError = "LIMIT statement must be followed by a number."

// LimitState parses LIMIT SQL clauses along with the value.
// ... LIMIT 10 ...
type LimitState struct {
	BaseState
}

func (s *LimitState) Name() string {
	return "LIMIT"
}

func (s *LimitState) Parse(result ast.Node, tokenizer *Tokenizer) (ast.Node, bool) {
	concrete := result.(ast.WithLimit)

	if token, _ := tokenizer.ReadToken(); token != LIMIT {
		tokenizer.UnreadToken()
		return result, false
	}

	if token, limit := tokenizer.ReadToken(); token == LITERAL {
		concrete.SetLimit(limit)
	} else {
		wrongTokenPanic(LimitWithoutNumberError, limit)
	}

	return result, true
}
