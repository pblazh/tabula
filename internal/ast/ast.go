package ast

import (
	"fmt"

	"github.com/pblazh/csvss/internal/lexer"
)

type Node interface {
	String() string
}

type Statement interface {
	Node
}

type LetStatement struct {
	Statement
	Identifier IdentifierExpression
	Value      Expression
}

func (s LetStatement) String() string {
	return fmt.Sprintf("let %s = %s;", s.Identifier.String(), s.Value.String())
}

type ExpressionStatement struct {
	Statement
	Token lexer.Lexem
	Value Expression
}

func (s ExpressionStatement) String() string {
	return s.Value.String()
}

type Expression interface {
	Node
}

type IdentifierExpression struct {
	Expression
	Value lexer.Lexem
}

func (s IdentifierExpression) String() string {
	return s.Value.Literal
}

type NumberExpression struct {
	Expression
	Value lexer.Lexem
}

func (s NumberExpression) String() string {
	return s.Value.Literal
}

type PrefixExpression struct {
	Expression
	Operator lexer.Lexem
	Value    Expression
}

func (s PrefixExpression) String() string {
	return fmt.Sprintf("%s%s", s.Operator.Literal, s.Value.String())
}

type SumExpression struct {
	Expression
	Left  Expression
	Right Expression
}

func (s SumExpression) String() string {
	return fmt.Sprintf("%s + %s", s.Left.String(), s.Right.String())
}

type Program []Statement
