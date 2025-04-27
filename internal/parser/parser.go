package parser

import (
	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

type Parser struct {
	lex *lexer.Lexer
	cur lexer.Lexem
	nex lexer.Lexem
}

func New(lex *lexer.Lexer) *Parser {
	return &Parser{
		lex: lex,
	}
}

func (p *Parser) Parse() ast.Program {
	program := make([]ast.Statement, 0)

	p.advance()
	p.advance()

outer:
	for p.cur.Type != lexer.EOF {
		switch p.cur.Type {
		case lexer.LET:
			program = append(program, p.parseLetStatement())
		default:
			break outer
		}
	}

	return program
}

func (p *Parser) advance() {
	p.cur = p.nex
	p.nex = p.lex.Next()
}

func (p *Parser) parseLetStatement() ast.Statement {
	p.advance()

	statement := ast.LetStatement{
		Identifier: ast.IdentifierExpression{Value: p.cur},
	}
	p.advance()
	p.advance()
	statement.Value = p.parseExpression()

	return statement
}

func (p *Parser) parseExpression() ast.Expression {
	val := p.cur
	p.advance()
	return ast.NumberExpression{
		Value: val,
	}
}
