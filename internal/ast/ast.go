package ast

import (
	"fmt"

	"github.com/pblazh/csvss/internal/lexer"
)

type Node interface {
	Literal() string
}

type Statement interface {
	Node
}

type LetStatement struct {
	Statement
	Identifier IdentifierExpression
	Value      Expression
}

func (s LetStatement) Literal() string {
	return fmt.Sprintf("let %s = %s", s.Identifier.Literal(), s.Value.Literal())
}

type Expression interface {
	Node
}

type IdentifierExpression struct {
	Expression
	Value lexer.Lexem
}

func (s IdentifierExpression) Literal() string {
	return s.Value.Literal
}

type NumberExpression struct {
	Expression
	Value lexer.Lexem
}

func (s NumberExpression) Literal() string {
	return s.Value.Literal
}

type SumExpression struct {
	Expression
	Left  Expression
	Right Expression
}

func (s SumExpression) Literal() string {
	return fmt.Sprintf("%s + %s", s.Left.Literal(), s.Right.Literal())
}

type Program []Statement
