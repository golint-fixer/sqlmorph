package parse

import (
	"github.com/s2gatev/sqlmorph/ast"
)

const (
	InnerWithoutJoinError         = "Expected JOIN following INNER."
	InnerJoinWithoutTargetError   = "INNER JOIN statement must be followed by a target class."
	InnerJoinWithoutOnError       = "INNER JOIN statement must have an ON clause."
	InnerJoinWrongJoinFieldsError = "Wrong join fields in INNER JOIN statement."
)

type InnerJoinState struct {
	BaseState
}

func (s *InnerJoinState) Name() string {
	return "INNER JOIN"
}

func (s *InnerJoinState) Parse(result ast.Node, tokenizer *Tokenizer) (ast.Node, bool) {
	concrete := result.(ast.WithRelations)

	if token, _ := tokenizer.ReadToken(); token != INNER {
		tokenizer.UnreadToken()
		return result, false
	}

	if token, value := tokenizer.ReadToken(); token != JOIN {
		wrongTokenPanic(InnerWithoutJoinError, value)
	}

	join := &ast.InnerJoin{}
	target := &ast.Target{}

	if token, tableName := tokenizer.ReadToken(); token == LITERAL {
		target.Name = tableName
	} else {
		wrongTokenPanic(InnerJoinWithoutTargetError, tableName)
	}

	if token, tableAlias := tokenizer.ReadToken(); token == LITERAL {
		target.Alias = tableAlias
	} else {
		tokenizer.UnreadToken()
	}

	join.SetTarget(target)

	if token, value := tokenizer.ReadToken(); token != ON {
		wrongTokenPanic(InnerJoinWithoutOnError, value)
	}

	if token, leftField := tokenizer.ReadToken(); token == LITERAL {
		join.Left = parseField(leftField)
	} else {
		wrongTokenPanic(InnerJoinWrongJoinFieldsError, leftField)
	}

	if token, operator := tokenizer.ReadToken(); token != EQUALS {
		wrongTokenPanic(InnerJoinWrongJoinFieldsError, operator)
	}

	if token, rightField := tokenizer.ReadToken(); token == LITERAL {
		join.Right = parseField(rightField)
	} else {
		wrongTokenPanic(InnerJoinWrongJoinFieldsError, rightField)
	}

	concrete.AddRelation(join)

	return result, true
}
