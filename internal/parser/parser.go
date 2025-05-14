package parser

import (
	"errors"
	"fmt"
	"strconv"

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
	parser.registerPrefix(lexer.INT, parser.parseInt)
	parser.registerPrefix(lexer.FLOAT, parser.parseFloat)
	parser.registerPrefix(lexer.TRUE, parser.parseBool)
	parser.registerPrefix(lexer.FALSE, parser.parseBool)
	parser.registerPrefix(lexer.LPAREN, parser.parseLparen)

	parser.registerPrefix(lexer.MINUS, parser.parsePrefix)
	parser.registerPrefix(lexer.NOT, parser.parsePrefix)

	parser.registerInfix(lexer.PLUS, parser.parseInfix)
	parser.registerInfix(lexer.MINUS, parser.parseInfix)
	parser.registerInfix(lexer.MULT, parser.parseInfix)
	parser.registerInfix(lexer.DIV, parser.parseInfix)
	parser.registerInfix(lexer.EQUAL, parser.parseInfix)
	parser.registerInfix(lexer.NOT_EQUAL, parser.parseInfix)
	parser.registerInfix(lexer.LESS, parser.parseInfix)
	parser.registerInfix(lexer.GREATER, parser.parseInfix)
	parser.registerInfix(lexer.COLUMN, parser.parseInfix)
	parser.registerInfix(lexer.LPAREN, parser.parseCallExpression)

	return parser
}

func (p *Parser) curretnPrecedence() int {
	if p, ok := precedences[p.cur.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) nextTokenIs(typ lexer.LexemType) bool {
	return p.nex.Type == typ
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
		Identifier: ast.IdentifierExpression{Right: p.cur},
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
	statement.Right = expression
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

	statement.Right = expression
	if p.expectCurLexem(lexer.SEMI) {
		p.advance()
	}

	return statement, nil
}

func (p *Parser) parseExpression(precendence int) (ast.Expression, error) {
	prefix := p.prefixParsers[p.cur.Type]

	if prefix == nil {
		return nil, errors.New("parser not found for " + p.cur.Literal)
	}

	leftExpr, err := prefix()
	if err != nil {
		return nil, err
	}

	for !p.nextTokenIs(lexer.SEMI) && precendence < p.curretnPrecedence() {
		infix := p.infixParsers[p.cur.Type]
		if infix == nil {
			return leftExpr, nil
		}

		leftExpr, err = infix(leftExpr)
		if err != nil {
			return nil, err
		}
	}

	return leftExpr, err
}

func (p *Parser) parseIdentifier() (ast.Expression, error) {
	expr := ast.IdentifierExpression{Right: p.cur}
	p.advance()
	return expr, nil
}

func (p *Parser) parseInt() (ast.Expression, error) {
	value, err := strconv.Atoi(p.cur.Literal)
	if err != nil {
		return nil, err
	}
	expr := ast.IntExpression{Right: p.cur, Value: value}
	p.advance()
	return expr, nil
}

func (p *Parser) parseFloat() (ast.Expression, error) {
	value, err := strconv.ParseFloat(p.cur.Literal, 32)
	if err != nil {
		return nil, err
	}

	expr := ast.FloatExpression{Right: p.cur, Value: value}
	p.advance()
	return expr, nil
}

func (p *Parser) parseBool() (ast.Expression, error) {
	expr := ast.BooleanExpression{Right: p.cur, Value: p.cur.Type == lexer.TRUE}
	p.advance()
	return expr, nil
}

func (p *Parser) parseLparen() (ast.Expression, error) {
	p.advance()
	expression, err := p.parseExpression(LOWEST)
	if err != nil {
		return nil, err
	}

	if !p.expectCurLexem(lexer.RPAREN) {
		return nil, fmt.Errorf("expected right paren, but got %v", p.cur)
	}
	p.advance()

	return expression, nil
}

func (p *Parser) parsePrefix() (ast.Expression, error) {
	if _, ok := p.prefixParsers[p.cur.Type]; !ok {
		return nil, fmt.Errorf("expected prefix, but got %v", p.cur)
	}

	prefix := ast.PrefixExpression{Operator: p.cur}
	p.advance()
	expression, err := p.parseExpression(PREFIX)
	if err != nil {
		return nil, err
	}

	prefix.Right = expression
	return prefix, nil
}

func (p *Parser) parseInfix(left ast.Expression) (ast.Expression, error) {
	expression := &ast.InfixExpression{
		Operator: p.cur,
		Left:     left,
	}

	precedence := p.curretnPrecedence()
	p.advance()
	right, err := p.parseExpression(precedence)
	if err != nil {
		return nil, err
	}
	expression.Right = right

	return expression, nil
}

func (p *Parser) parseCallExpression(left ast.Expression) (ast.Expression, error) {
	expr := ast.CallExpression{Identifier: left}
	arguments, err := p.parseCallArguments()
	if err != nil {
		return nil, err
	}
	expr.Arguments = arguments
	return expr, nil
}

func (p *Parser) parseCallArguments() ([]ast.Expression, error) {
	p.advance()

	arguments := []ast.Expression{}
	if p.expectCurLexem(lexer.RPAREN) {
		p.advance()
		return arguments, nil
	}

	expr, err := p.parseExpression(LOWEST)
	if err != nil {
		return nil, err
	}
	arguments = append(arguments, expr)

	for p.expectCurLexem(lexer.COMA) {
		p.advance()
		expr, err := p.parseExpression(LOWEST)
		if err != nil {
			return nil, err
		}
		arguments = append(arguments, expr)
	}

	p.advance()
	return arguments, nil
}
