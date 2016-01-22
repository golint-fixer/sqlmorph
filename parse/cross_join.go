package parse

import (
	"github.com/s2gatev/sqlmorph/ast"
)

const (
	CrossWithoutJoinError       = "Expected JOIN following CROSS."
	CrossJoinWithoutTargetError = "CROSS JOIN statement must be followed by a target class."
)

type CrossJoinState struct {
	BaseState
}

func (s *CrossJoinState) Name() string {
	return "CROSS JOIN"
}

func (s *CrossJoinState) Parse(result ast.Node, tokenizer *Tokenizer) (ast.Node, bool) {
	concrete := result.(ast.WithRelations)

	if token, _ := tokenizer.ReadToken(); token != CROSS {
		tokenizer.UnreadToken()
		return result, false
	}

	if token, value := tokenizer.ReadToken(); token != JOIN {
		wrongTokenPanic(CrossWithoutJoinError, value)
	}

	join := &ast.CrossJoin{}
	target := &ast.Target{}

	if token, tableName := tokenizer.ReadToken(); token == LITERAL {
		target.Name = tableName
	} else {
		wrongTokenPanic(CrossJoinWithoutTargetError, tableName)
	}

	if token, tableAlias := tokenizer.ReadToken(); token == LITERAL {
		target.Alias = tableAlias
	} else {
		tokenizer.UnreadToken()
	}

	join.SetTarget(target)

	concrete.AddRelation(join)

	return result, true
}
