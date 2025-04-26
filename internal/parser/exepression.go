package parser

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

type AssignStatement struct {
	Statement
	Identifier lexer.Token
	Value      Expression
}

func (s AssignStatement) Literal() string {
	return fmt.Sprintf("%s = %s", s.Identifier.Literal, s.Value.Literal())
}

type Expression struct {
	Node
	Evaluate func()
}

type NumberExpression struct {
	Expression
	Value lexer.Token
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
