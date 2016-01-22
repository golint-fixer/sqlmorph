package parse

import (
	"github.com/s2gatev/sqlmorph/ast"
)

const UpdateWithoutTargetError = "UPDATE statement must be followed by a target class."

// UpdateState parses UPDATE SQL clauses along with the target table.
// UPDATE User u ...
type UpdateState struct {
	BaseState
}

func (s *UpdateState) Name() string {
	return "UPDATE"
}

func (s *UpdateState) Parse(result ast.Node, tokenizer *Tokenizer) (ast.Node, bool) {
	concrete := &ast.Update{}

	if token, _ := tokenizer.ReadToken(); token != UPDATE {
		tokenizer.UnreadToken()
		return result, false
	}

	target := &ast.Target{}

	if token, tableName := tokenizer.ReadToken(); token == LITERAL {
		target.Name = tableName
	} else {
		wrongTokenPanic(UpdateWithoutTargetError, tableName)
	}

	if token, tableAlias := tokenizer.ReadToken(); token == LITERAL {
		target.Alias = tableAlias
	} else {
		tokenizer.UnreadToken()
	}

	concrete.SetTarget(target)

	return concrete, true
}
