package parser

import (
	"fmt"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

type (
	prefixParse func() (ast.Expression, error)
	infixParse  func(ast.Expression) (ast.Expression, error)
)

type Parser struct {
	lex *lexer.Lexer
	cur lexer.Lexem
	nex lexer.Lexem

	prefixParsers map[lexer.LexemType]prefixParse
	infixParsers  map[lexer.LexemType]infixParse
}

func New(lex *lexer.Lexer) *Parser {
	parser := &Parser{
		lex:           lex,
		prefixParsers: make(map[lexer.LexemType]prefixParse),
		infixParsers:  make(map[lexer.LexemType]infixParse),
	}

	parser.registerPrefix(lexer.IDENT, parser.parseIdentifier)
	parser.registerPrefix(lexer.NUMBER, parser.parseNumber)
	parser.registerPrefix(lexer.MINUS, parser.parsePrefix)
	parser.registerPrefix(lexer.NOT, parser.parsePrefix)

	return parser
}

func (p *Parser) Parse() ast.Program {
	program := make([]ast.Statement, 0)

	p.advance()
	p.advance()

	for p.cur.Type != lexer.EOF {
		switch p.cur.Type {
		case lexer.LET:
			stm, err := p.parseLetStatement()
			if err != nil {
				panic(err)
			}
			program = append(program, stm)
		default:
			stm, err := p.parseExpressionStatement()
			if err != nil {
				panic(err)
			}
			program = append(program, stm)
		}
	}

	return program
}

func (p *Parser) advance() {
	p.cur = p.nex
	p.nex = p.lex.Next()
}

func (p *Parser) expectCurLexem(typ lexer.LexemType) bool {
	return p.cur.Type == typ
}

func (p *Parser) registerPrefix(l lexer.LexemType, parse prefixParse) {
	p.prefixParsers[l] = parse
}

func (p *Parser) registerInfix(l lexer.LexemType, parse infixParse) {
	p.infixParsers[l] = parse
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
	expression, err := p.parseExpression(LOWEST)
	if err != nil {
		return nil, err
	}
	statement.Value = expression
	if p.expectCurLexem(lexer.SEMI) {
		p.advance()
	}

	return statement, nil
}

func (p *Parser) parseExpressionStatement() (ast.Statement, error) {
	statement := ast.ExpressionStatement{
		Token: p.cur,
	}

	expression, err := p.parseExpression(LOWEST)
	if err != nil {
		return nil, err
	}

	statement.Value = expression
	if p.expectCurLexem(lexer.SEMI) {
		p.advance()
	}

	return statement, nil
}

func (p *Parser) parseExpression(precendence int) (ast.Expression, error) {
	prefix := p.prefixParsers[p.cur.Type]
	if prefix == nil {
		return nil, nil
	}

	leftExpr, err := prefix()
	return leftExpr, err
}

func (p *Parser) parseIdentifier() (ast.Expression, error) {
	if p.cur.Type != lexer.IDENT {
		return nil, fmt.Errorf("expected an identifier, but got %v", p.cur)
	}
	expr := ast.IdentifierExpression{Value: p.cur}
	p.advance()
	return expr, nil
}

func (p *Parser) parseNumber() (ast.Expression, error) {
	if p.cur.Type != lexer.NUMBER {
		return nil, fmt.Errorf("expected an identifier, but got %v", p.cur)
	}

	expr := ast.NumberExpression{Value: p.cur}
	p.advance()
	return expr, nil
}

func (p *Parser) parsePrefix() (ast.Expression, error) {
	if p.cur.Type != lexer.MINUS {
		return nil, fmt.Errorf("expected -, but got %v", p.cur)
	}

	prefix := ast.PrefixExpression{Operator: p.cur}
	p.advance()
	expression, err := p.parseExpression(LOWEST)
	if err != nil {
		return nil, err
	}

	prefix.Value = expression
	return prefix, nil
}
