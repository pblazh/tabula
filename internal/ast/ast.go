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

type LetStatement struct {
	Node
	Identifier IdentifierExpression
	Value      Node
}

func (stmt LetStatement) String() string {
	return fmt.Sprintf("let %s = %s;", stmt.Identifier.String(), stmt.Value)
}

type FmtStatement struct {
	Node
	Identifier IdentifierExpression
	Value      Node
}

func (stmt FmtStatement) String() string {
	return fmt.Sprintf("fmt %s = %s;", stmt.Identifier.String(), stmt.Value)
}

type ExpressionStatement struct {
	Node
	Token lexer.Token
	Value Node
}

func (stmt ExpressionStatement) String() string {
	return stmt.Value.String() + ";"
}

type IncludeStatement struct {
	Node
	Token lexer.Token
	Path  string
}

func (stmt IncludeStatement) String() string {
	return fmt.Sprintf("#include \"%s\";", stmt.Path)
}

type IdentifierExpression struct {
	Node
	Token lexer.Token
	Value string
}

func (expr IdentifierExpression) String() string {
	return expr.Value
}

type BooleanExpression struct {
	Node
	Token lexer.Token
	Value bool
}

func (expr BooleanExpression) String() string {
	return fmt.Sprintf("%v", expr.Value)
}

type IntExpression struct {
	Node
	Token lexer.Token
	Value int
}

func (expr IntExpression) String() string {
	return fmt.Sprintf("%d", expr.Value)
}

type FloatExpression struct {
	Node
	Token lexer.Token
	Value float64
}

func (expr FloatExpression) String() string {
	return fmt.Sprintf("%.2f", expr.Value)
}

type StringExpression struct {
	Node
	Token lexer.Token
	Value string
}

func (expr StringExpression) String() string {
	return fmt.Sprintf("\"%s\"", expr.Value)
}

type DateExpression struct {
	Node
	Token lexer.Token
	Value time.Time
}

func (expr DateExpression) String() string {
	return fmt.Sprintf("<%s>", expr.Value.Format("2006-01-02 15:04:05"))
}

type PrefixExpression struct {
	Node
	Token    lexer.Token
	Operator lexer.Token
	Value    Node
}

func (expr PrefixExpression) String() string {
	return fmt.Sprintf("%s%s", expr.Operator.Literal, expr.Value)
}

type InfixExpression struct {
	Node
	Token    lexer.Token
	Left     Node
	Operator lexer.Token
	Right    Node
}

func (expr InfixExpression) String() string {
	return fmt.Sprintf("%s %s %s", expr.Left, expr.Operator.Literal, expr.Right)
}

type CallExpression struct {
	Node
	Token      lexer.Token
	Identifier Node
	Arguments  []Node
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
	Node
	Token lexer.Token
	Value []string
}

func (expr RangeExpression) String() string {
	return fmt.Sprintf("[%s]", strings.Join(expr.Value, ", "))
}

type Program []Node

func (p Program) String() string {
	var sb strings.Builder
	for i := range p {
		sb.WriteString(p[i].String())
	}
	return sb.String()
}
