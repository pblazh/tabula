package parser

import (
	"fmt"

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
			stm, err := p.parseLetStatement()
			if err != nil {
				panic(err)
			}
			program = append(program, stm)
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

func (p *Parser) parseLetStatement() (ast.Statement, error) {
	p.advance()
	if !p.expectCurLexem(lexer.IDENT) {
		return nil, fmt.Errorf("expected an identifier, but got %v", p.cur)
	}

	statement := ast.LetStatement{
		Identifier: ast.IdentifierExpression{Value: p.cur},
	}
	p.advance()
	if !p.expectCurLexem(lexer.ASSIGN) {
		return nil, fmt.Errorf("expected =, but got %v", p.cur)
	}

	p.advance()
	expression, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	statement.Value = expression

	return statement, nil
}

func (p *Parser) expectCurLexem(typ lexer.LexemType) bool {
	return p.cur.Type == typ
}

func (p *Parser) parseExpression() (ast.Expression, error) {
	val := p.cur

	for !p.expectCurLexem(lexer.SEMI) {
		p.advance()
	}
	return ast.NumberExpression{
		Value: val,
	}, nil
}
