// Package ast defines the abstract syntax tree nodes for the CSV spreadsheet language.
package ast

import (
	"fmt"
	"strings"
	"time"

	"github.com/pblazh/tabula/internal/lexer"
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

type FmtStatement struct {
	Statement
	Identifier IdentifierExpression
	Value      Expression
}

func (stmt FmtStatement) String() string {
	return fmt.Sprintf("fmt %s = %s;", stmt.Identifier.String(), stmt.Value)
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
	Value string
}

func (expr IdentifierExpression) String() string {
	return expr.Value
}

type BooleanExpression struct {
	Expression
	Token lexer.Token
	Value bool
}

func (expr BooleanExpression) String() string {
	return fmt.Sprintf("%v", expr.Value)
}

type IntExpression struct {
	Expression
	Token lexer.Token
	Value int
}

func (expr IntExpression) String() string {
	return fmt.Sprintf("%d", expr.Value)
}

type FloatExpression struct {
	Expression
	Token lexer.Token
	Value float64
}

func (expr FloatExpression) String() string {
	return fmt.Sprintf("%.2f", expr.Value)
}

type StringExpression struct {
	Expression
	Token lexer.Token
	Value string
}

func (expr StringExpression) String() string {
	return fmt.Sprintf("\"%s\"", expr.Value)
}

type DateExpression struct {
	Expression
	Token lexer.Token
	Value time.Time
}

func (expr DateExpression) String() string {
	return fmt.Sprintf("<%s>", expr.Value.Format("2006-01-02 15:04:05"))
}

type PrefixExpression struct {
	Expression
	Token    lexer.Token
	Operator lexer.Token
	Value    Expression
}

func (expr PrefixExpression) String() string {
	return fmt.Sprintf("%s%s", expr.Operator.Literal, expr.Value)
}

type InfixExpression struct {
	Expression
	Token    lexer.Token
	Left     Expression
	Operator lexer.Token
	Right    Expression
}

func (expr InfixExpression) String() string {
	return fmt.Sprintf("%s %s %s", expr.Left, expr.Operator.Literal, expr.Right)
}

type CallExpression struct {
	Expression
	Token      lexer.Token
	Identifier Expression
	Arguments  []Expression
}

func (expr CallExpression) String() string {
	b := strings.Builder{}
	b.WriteString(expr.Identifier.String())
	b.WriteString("(")

	for i, arg := range expr.Arguments {
		b.WriteString(arg.String())
		if i < len(expr.Arguments)-1 {
			b.WriteString(", ")
		}
	}
	b.WriteString(")")

	return b.String()
}

type RangeExpression struct {
	Expression
	Token lexer.Token
	Value []string
}

func (expr RangeExpression) String() string {
	return fmt.Sprintf("[%s]", strings.Join(expr.Value, ", "))
}

type Program []Statement
