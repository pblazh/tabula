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

var precedences = map[lexer.TokenType]int{
	lexer.EQUAL:            EQUALS,
	lexer.NOT_EQUAL:        EQUALS,
	lexer.LESS:             LESSGREATER,
	lexer.GREATER:          LESSGREATER,
	lexer.GREATER_OR_EQUAL: LESSGREATER,
	lexer.LESS_OR_EQUAL:    LESSGREATER,
	lexer.PLUS:             SUM,
	lexer.MINUS:            SUM,
	lexer.DIV:              PRODUCT,
	lexer.MULT:             PRODUCT,
	lexer.REM:              PRODUCT,
	lexer.LPAREN:           CALL,
	lexer.COLUMN:           CALL,
}
