// Package ast defines the abstract syntax tree nodes for the CSV spreadsheet language.
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
	Value      Expression
}

func (stmt LetStatement) String() string {
	return fmt.Sprintf("let %s = %s;", stmt.Identifier.String(), stmt.Value)
}

type ExpressionStatement struct {
	Statement
	Token lexer.Token
	Value Expression
}

func (stmt ExpressionStatement) String() string {
	return stmt.Value.String() + ";"
}

type Expression interface {
	Node
}

type IdentifierExpression struct {
	Expression
	Token lexer.Token
}

func (expr IdentifierExpression) String() string {
	return expr.Token.Literal
}

type BooleanExpression struct {
	Expression
	Token lexer.Token
	Value bool
}

func (expr BooleanExpression) String() string {
	return fmt.Sprintf("<bool %v>", expr.Value)
}

type IntExpression struct {
	Expression
	Token lexer.Token
	Value int
}

func (expr IntExpression) String() string {
	return fmt.Sprintf("<int %d>", expr.Value)
}

type FloatExpression struct {
	Expression
	Token lexer.Token
	Value float64
}

func (expr FloatExpression) String() string {
	return fmt.Sprintf("<float %.2f>", expr.Value)
}

type StringExpression struct {
	Expression
	Token lexer.Token
}

func (expr StringExpression) String() string {
	return fmt.Sprintf("<str %s>", expr.Token.Literal)
}

type PrefixExpression struct {
	Expression
	Operator lexer.Token
	Value    Expression
}

func (expr PrefixExpression) String() string {
	return fmt.Sprintf("(%s %s)", expr.Operator.Literal, expr.Value)
}

type InfixExpression struct {
	Expression
	Left     Expression
	Operator lexer.Token
	Right    Expression
}

func (expr InfixExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", expr.Operator.Literal, expr.Left, expr.Right)
}

type CallExpression struct {
	Expression
	Identifier Expression
	Arguments  []Expression
}

func (expr CallExpression) String() string {
	b := strings.Builder{}
	b.WriteString("(")
	b.WriteString(expr.Identifier.String())
	if len(expr.Arguments) > 0 {
		b.WriteString(" ")
	}

	for i, arg := range expr.Arguments {
		b.WriteString(arg.String())
		if i < len(expr.Arguments)-1 {
			b.WriteString(" ")
		}
	}
	b.WriteString(")")

	return b.String()
}

type RangeExpression struct {
	Expression
	Value []string
}

func (expr RangeExpression) String() string {
	return fmt.Sprintf("(: %s)", strings.Join(expr.Value, " "))
}

type Program []Statement
