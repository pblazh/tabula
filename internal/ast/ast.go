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
	Right      Expression
}

func (s LetStatement) String() string {
	return fmt.Sprintf("let %s = %s;", s.Identifier.String(), s.Right)
}

type ExpressionStatement struct {
	Statement
	Token lexer.Lexem
	Right Expression
}

func (s ExpressionStatement) String() string {
	return s.Right.String() + ";"
}

type Expression interface {
	Node
}

type IdentifierExpression struct {
	Expression
	Right lexer.Lexem
}

func (s IdentifierExpression) String() string {
	return s.Right.Literal
}

type BooleanExpression struct {
	Expression
	Right lexer.Lexem
	Value bool
}

func (s BooleanExpression) String() string {
	return fmt.Sprintf("<bool %v>", s.Value)
}

type NumberExpression struct {
	Expression
	Right lexer.Lexem
}

func (s NumberExpression) String() string {
	return s.Right.Literal
}

type PrefixExpression struct {
	Expression
	Operator lexer.Lexem
	Right    Expression
}

func (s PrefixExpression) String() string {
	return fmt.Sprintf("%s%s", s.Operator.Literal, s.Right)
}

type SumExpression struct {
	Expression
	Left  Expression
	Right Expression
}

func (s SumExpression) String() string {
	return fmt.Sprintf("%s + %s", s.Left.String(), s.Right.String())
}

type InfixExpression struct {
	Expression
	Left     Expression
	Operator lexer.Lexem
	Right    Expression
}

func (s InfixExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", s.Left, s.Operator.Literal, s.Right)
}

type Program []Statement
