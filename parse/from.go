package parse

import (
	"github.com/s2gatev/sqlmorph/ast"
)

const FromWithoutTargetError = "FROM statement must be followed by a target class."

// FromState parses FROM SQL clauses along with the table name and alias.
// ... FROM User u ...
type FromState struct {
	BaseState
}

func (s *FromState) Name() string {
	return "FROM"
}

func (s *FromState) Parse(result ast.Node, tokenizer *Tokenizer) (ast.Node, bool) {
	concrete := result.(ast.WithTarget)

	if token, _ := tokenizer.ReadToken(); token != FROM {
		tokenizer.UnreadToken()
		return result, false
	}

	target := &ast.Target{}

	if token, tableName := tokenizer.ReadToken(); token == LITERAL {
		target.Name = tableName
	} else {
		wrongTokenPanic(FromWithoutTargetError, tableName)
	}

	if token, tableAlias := tokenizer.ReadToken(); token == LITERAL {
		target.Alias = tableAlias
	} else {
		tokenizer.UnreadToken()
	}

	concrete.SetTarget(target)

	return result, true
}
