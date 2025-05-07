package parser

import "github.com/pblazh/csvss/internal/lexer"

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

var precedences = map[lexer.LexemType]int{
	lexer.EQUAL:     EQUALS,
	lexer.NOT_EQUAL: EQUALS,
	lexer.LESS:      LESSGREATER,
	lexer.GREATER:   LESSGREATER,
	lexer.PLUS:      SUM,
	lexer.MINUS:     SUM,
	lexer.DIV:       PRODUCT,
	lexer.MULT:      PRODUCT,
}
