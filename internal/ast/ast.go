package ast

import (
	"fmt"
	"strings"

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

type IntExpression struct {
	Expression
	Right lexer.Lexem
	Value int
}

func (s IntExpression) String() string {
	return fmt.Sprintf("<int %d>", s.Value)
}

type FloatExpression struct {
	Expression
	Right lexer.Lexem
	Value float64
}

func (s FloatExpression) String() string {
	return fmt.Sprintf("<float %.2f>", s.Value)
}

type StringExpression struct {
	Expression
	Right lexer.Lexem
}

func (s StringExpression) String() string {
	return fmt.Sprintf("<str %s>", s.Right.Literal)
}

type PrefixExpression struct {
	Expression
	Operator lexer.Lexem
	Right    Expression
}

func (s PrefixExpression) String() string {
	return fmt.Sprintf("(%s %s)", s.Operator.Literal, s.Right)
}

type InfixExpression struct {
	Expression
	Left     Expression
	Operator lexer.Lexem
	Right    Expression
}

func (s InfixExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", s.Operator.Literal, s.Left, s.Right)
}

type CallExpression struct {
	Expression
	Identifier Expression
	Arguments  []Expression
}

func (s CallExpression) String() string {
	b := strings.Builder{}
	b.WriteString("(")
	b.WriteString(s.Identifier.String())
	if len(s.Arguments) > 0 {
		b.WriteString(" ")
	}

	for i, arg := range s.Arguments {
		b.WriteString(arg.String())
		if i < len(s.Arguments)-1 {
			b.WriteString(" ")
		}
	}
	b.WriteString(")")

	return b.String()
}

type Program []Statement
