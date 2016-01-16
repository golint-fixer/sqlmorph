package parsing

import (
	"github.com/s2gatev/sqlmorph/ast"
)

const (
	RightWithoutJoinError         = "Expected JOIN following RIGHT."
	RightJoinWithoutTargetError   = "RIGHT JOIN statement must be followed by a target class."
	RightJoinWithoutOnError       = "RIGHT JOIN statement must have an ON clause."
	RightJoinWrongJoinFieldsError = "Wrong join fields in RIGHT JOIN statement."
)

type RightJoinState struct {
	BaseState
}

func (s *RightJoinState) Name() string {
	return "RIGHT JOIN"
}

func (s *RightJoinState) Parse(result ast.Node, tokenizer *Tokenizer) (ast.Node, bool) {
	concrete := result.(ast.WithRelations)

	if token, _ := tokenizer.ReadToken(); token != RIGHT {
		tokenizer.UnreadToken()
		return result, false
	}

	if token, value := tokenizer.ReadToken(); token != JOIN {
		wrongTokenPanic(RightWithoutJoinError, value)
	}

	join := &ast.RightJoin{}
	target := &ast.Target{}

	if token, tableName := tokenizer.ReadToken(); token == LITERAL {
		target.Name = tableName
	} else {
		wrongTokenPanic(RightJoinWithoutTargetError, tableName)
	}

	if token, tableAlias := tokenizer.ReadToken(); token == LITERAL {
		target.Alias = tableAlias
	} else {
		tokenizer.UnreadToken()
	}

	join.SetTarget(target)

	if token, value := tokenizer.ReadToken(); token != ON {
		wrongTokenPanic(RightJoinWithoutOnError, value)
	}

	if token, leftField := tokenizer.ReadToken(); token == LITERAL {
		join.Left = parseField(leftField)
	} else {
		wrongTokenPanic(RightJoinWrongJoinFieldsError, leftField)
	}

	if token, operator := tokenizer.ReadToken(); token != EQUALS {
		wrongTokenPanic(RightJoinWrongJoinFieldsError, operator)
	}

	if token, rightField := tokenizer.ReadToken(); token == LITERAL {
		join.Right = parseField(rightField)
	} else {
		wrongTokenPanic(RightJoinWrongJoinFieldsError, rightField)
	}

	concrete.AddRelation(join)

	return result, true
}
