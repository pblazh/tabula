// Package parser provides parsing functionality for the CSV spreadsheet language,
// converting list of tokens into an abstract syntax tree.
package parser

import (
	"strconv"
	"strings"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

type (
	prefixParse func() (ast.Expression, error)
	infixParse  func(ast.Expression) (ast.Expression, error)
)

type Parser struct {
	lex *lexer.Lexer
	cur lexer.Token
	nex lexer.Token

	prefixParsers map[lexer.TokenType]prefixParse
	infixParsers  map[lexer.TokenType]infixParse
	identifiers   []string
}

func New(lex *lexer.Lexer) *Parser {
	parser := &Parser{
		lex:           lex,
		prefixParsers: make(map[lexer.TokenType]prefixParse),
		infixParsers:  make(map[lexer.TokenType]infixParse),
		identifiers:   make([]string, 0),
	}

	parser.registerPrefix(lexer.IDENT, parser.parseIdentifier)
	parser.registerPrefix(lexer.INT, parser.parseInt)
	parser.registerPrefix(lexer.FLOAT, parser.parseFloat)
	parser.registerPrefix(lexer.TRUE, parser.parseBool)
	parser.registerPrefix(lexer.FALSE, parser.parseBool)
	parser.registerPrefix(lexer.STRING, parser.parseString)
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

func (p *Parser) currentPrecedence() int {
	if p, ok := precedences[p.cur.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) nextTokenIs(typ lexer.TokenType) bool {
	return p.nex.Type == typ
}

func (p *Parser) Parse() (ast.Program, []string, error) {
	program := make([]ast.Statement, 0)
	p.identifiers = make([]string, 0) // Reset identifiers for each parse

	err := p.advance()
	if err != nil {
		return nil, nil, err
	}
	err = p.advance()
	if err != nil {
		return nil, nil, err
	}

	for p.cur.Type != lexer.EOF {
		switch p.cur.Type {
		case lexer.LET:
			stm, err := p.parseLetStatement()
			if err != nil {
				return nil, nil, err
			}
			program = append(program, stm)
		case lexer.FMT:
			stm, err := p.parseFmtStatement()
			if err != nil {
				return nil, nil, err
			}
			program = append(program, stm)
		default:
			stm, err := p.parseExpressionStatement()
			if err != nil {
				return nil, nil, err
			}
			program = append(program, stm)
		}
	}

	return program, p.identifiers, nil
}

func (p *Parser) advance() error {
	p.cur = p.nex

	nex, err := p.lex.Next()
	if err != nil {
		return err
	}
	p.nex = nex
	return nil
}

func (p *Parser) expectCurrentToken(typ lexer.TokenType) bool {
	return p.cur.Type == typ
}

func (p *Parser) registerPrefix(l lexer.TokenType, parse prefixParse) {
	p.prefixParsers[l] = parse
}

func (p *Parser) registerInfix(l lexer.TokenType, parse infixParse) {
	p.infixParsers[l] = parse
}

func (p *Parser) parseLetStatement() (ast.Statement, error) {
	err := p.advance()
	if err != nil {
		return nil, err
	}

	if !p.expectCurrentToken(lexer.IDENT) {
		return nil, ErrExpectedIdentifier(p.cur.Literal, p.cur.Position)
	}

	// parse range expression

	// Add let statement identifier to the list
	p.identifiers = append(p.identifiers, p.cur.Literal)

	statement := ast.LetStatement{
		Identifier: ast.IdentifierExpression{Token: p.cur, Value: p.cur.Literal},
	}
	err = p.advance()
	if err != nil {
		return nil, err
	}

	if !p.expectCurrentToken(lexer.ASSIGN) {
		return nil, ErrExpectedToken("=", p.cur)
	}

	err = p.advance()
	if err != nil {
		return nil, err
	}
	expression, err := p.parseExpression(LOWEST)
	if err != nil {
		return nil, err
	}
	statement.Value = expression
	if p.expectCurrentToken(lexer.SEMI) {
		err = p.advance()
		if err != nil {
			return nil, err
		}
	}

	return statement, nil
}

func (p *Parser) parseFmtStatement() (ast.Statement, error) {
	err := p.advance()
	if err != nil {
		return nil, err
	}

	if !p.expectCurrentToken(lexer.IDENT) {
		return nil, ErrExpectedIdentifier(p.cur.Literal, p.cur.Position)
	}

	// Add fmt statement identifier to the list
	p.identifiers = append(p.identifiers, p.cur.Literal)

	statement := ast.FmtStatement{
		Identifier: ast.IdentifierExpression{Token: p.cur, Value: p.cur.Literal},
	}
	err = p.advance()
	if err != nil {
		return nil, err
	}

	if !p.expectCurrentToken(lexer.ASSIGN) {
		return nil, ErrExpectedToken("=", p.cur)
	}

	err = p.advance()
	if err != nil {
		return nil, err
	}
	expression, err := p.parseExpression(LOWEST)
	if err != nil {
		return nil, err
	}
	statement.Value = expression
	if p.expectCurrentToken(lexer.SEMI) {
		err = p.advance()
		if err != nil {
			return nil, err
		}
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
	if p.expectCurrentToken(lexer.SEMI) {
		err := p.advance()
		if err != nil {
			return nil, err
		}
	}

	return statement, nil
}

func (p *Parser) parseExpression(precedence int) (ast.Expression, error) {
	prefix := p.prefixParsers[p.cur.Type]

	if prefix == nil {
		return nil, ErrUnexpectedToken(p.cur.Literal, p.cur.Position)
	}

	leftExpr, err := prefix()
	if err != nil {
		return nil, err
	}

	for !p.nextTokenIs(lexer.SEMI) && precedence < p.currentPrecedence() {
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
	// Add identifier to the list (convert cell identifiers to uppercase)
	literal := p.cur.Literal
	token := p.cur
	if ast.IsCellIdentifier(literal) {
		literal = strings.ToUpper(literal)
		token.Literal = literal // Also update the token for the AST
	}
	p.identifiers = append(p.identifiers, literal)

	expr := ast.IdentifierExpression{Token: token, Value: literal}
	err := p.advance()
	if err != nil {
		return nil, err
	}
	return expr, nil
}

func (p *Parser) parseInt() (ast.Expression, error) {
	value, err := strconv.Atoi(p.cur.Literal)
	if err != nil {
		return nil, err
	}
	expr := ast.IntExpression{Token: p.cur, Value: value}
	err = p.advance()
	if err != nil {
		return nil, err
	}
	return expr, nil
}

func (p *Parser) parseFloat() (ast.Expression, error) {
	value, err := strconv.ParseFloat(p.cur.Literal, 64)
	if err != nil {
		return nil, err
	}

	expr := ast.FloatExpression{Token: p.cur, Value: value}
	err = p.advance()
	if err != nil {
		return nil, err
	}
	return expr, nil
}

func (p *Parser) parseBool() (ast.Expression, error) {
	expr := ast.BooleanExpression{Token: p.cur, Value: p.cur.Type == lexer.TRUE}
	err := p.advance()
	if err != nil {
		return nil, err
	}
	return expr, nil
}

func (p *Parser) parseString() (ast.Expression, error) {
	expr := ast.StringExpression{Value: p.cur.Literal[1 : len(p.cur.Literal)-1], Token: p.cur}
	err := p.advance()
	if err != nil {
		return nil, err
	}
	return expr, nil
}

func (p *Parser) parseLparen() (ast.Expression, error) {
	err := p.advance()
	if err != nil {
		return nil, err
	}
	expression, err := p.parseExpression(LOWEST)
	if err != nil {
		return nil, err
	}

	if !p.expectCurrentToken(lexer.RPAREN) {
		return nil, ErrExpectedRightParen(p.cur)
	}
	err = p.advance()
	if err != nil {
		return nil, err
	}

	return expression, nil
}

func (p *Parser) parsePrefix() (ast.Expression, error) {
	if _, ok := p.prefixParsers[p.cur.Type]; !ok {
		return nil, ErrExpectedPrefix(p.cur)
	}

	prefix := ast.PrefixExpression{Token: p.cur, Operator: p.cur}
	err := p.advance()
	if err != nil {
		return nil, err
	}
	expression, err := p.parseExpression(PREFIX)
	if err != nil {
		return nil, err
	}

	prefix.Value = expression
	return prefix, nil
}

func (p *Parser) parseInfix(left ast.Expression) (ast.Expression, error) {
	// Check if this is a range operator (:)
	if p.cur.Type == lexer.COLUMN {
		return p.parseRange(left)
	}

	expression := ast.InfixExpression{
		Token:    p.cur,
		Operator: p.cur,
		Left:     left,
	}

	precedence := p.currentPrecedence()
	err := p.advance()
	if err != nil {
		return nil, err
	}
	right, err := p.parseExpression(precedence)
	if err != nil {
		return nil, err
	}
	expression.Right = right

	return expression, nil
}

func (p *Parser) parseRange(left ast.Expression) (ast.Expression, error) {
	// store the ':' token
	colonToken := p.cur
	// advance past the ':'
	if err := p.advance(); err != nil {
		return nil, err
	}

	// parse the right side
	right, err := p.parseExpression(LOWEST)
	if err != nil {
		return nil, err
	}

	leftIdent := left.(ast.IdentifierExpression)
	rightIdent := right.(ast.IdentifierExpression)

	cells, err := ast.ExpandRange(leftIdent.Value, rightIdent.Value)
	if err != nil {
		return nil, err
	}

	return ast.RangeExpression{Token: colonToken, Value: cells}, nil
}

func (p *Parser) parseCallExpression(left ast.Expression) (ast.Expression, error) {
	identifier := left.(ast.IdentifierExpression)
	expr := ast.CallExpression{Identifier: identifier, Token: identifier.Token}
	arguments, err := p.parseCallArguments()
	if err != nil {
		return nil, err
	}

	expandedArguments := make([]ast.Expression, 0, len(arguments))
	for _, arg := range arguments {
		if rangeExpr, ok := arg.(ast.RangeExpression); ok {
			for _, expandedArg := range rangeExpr.Value {
				expandedArguments = append(expandedArguments, ast.IdentifierExpression{Value: expandedArg, Token: rangeExpr.Token})
			}
		} else {
			expandedArguments = append(expandedArguments, arg)
		}
	}

	expr.Arguments = expandedArguments
	return expr, nil
}

func (p *Parser) parseCallArguments() ([]ast.Expression, error) {
	err := p.advance()
	if err != nil {
		return nil, err
	}

	arguments := []ast.Expression{}
	if p.expectCurrentToken(lexer.RPAREN) {
		err = p.advance()
		if err != nil {
			return nil, err
		}
		return arguments, nil
	}

	expr, err := p.parseExpression(LOWEST)
	if err != nil {
		return nil, err
	}
	arguments = append(arguments, expr)

	for p.expectCurrentToken(lexer.COMMA) {
		err := p.advance()
		if err != nil {
			return nil, err
		}
		expr, err := p.parseExpression(LOWEST)
		if err != nil {
			return nil, err
		}
		arguments = append(arguments, expr)
	}

	err = p.advance()
	if err != nil {
		return nil, err
	}
	return arguments, nil
}
