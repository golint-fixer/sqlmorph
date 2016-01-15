package parsing

import (
	"github.com/s2gatev/sqlmorph/ast"
)

const OffsetWithoutNumberError = "OFFSET statement must be followed by a number."

// OffsetState parses OFFSET SQL clauses along with the value.
// ... OFFSET 20 ...
type OffsetState struct {
	BaseState
}

func (s *OffsetState) Name() string {
	return "OFFSET"
}

func (s *OffsetState) Parse(result ast.Node, tokenizer *Tokenizer) (ast.Node, bool) {
	concrete := result.(ast.Offsetable)

	if token, _ := tokenizer.ReadToken(); token != OFFSET {
		tokenizer.UnreadToken()
		return result, false
	}

	if token, offset := tokenizer.ReadToken(); token == LITERAL {
		concrete.SetOffset(offset)
	} else {
		wrongTokenPanic(OffsetWithoutNumberError, offset)
	}

	return result, true
}
