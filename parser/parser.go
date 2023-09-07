package parser

import (
	"fmt"
	"geomys/lexer"
	"geomys/token"
	"geomys/tree"
)

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

type (
	prefixParseFunc func() tree.Expression
	infixParseFunc  func(tree.Expression) tree.Expression
)

type Parser struct {
	l           *lexer.Lexer
	currToken   token.Token
	peekedToken token.Token
	errors      []string

	prefixParseFunc map[token.TokenType]prefixParseFunc
	infixParseFunc  map[token.TokenType]infixParseFunc
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFunc) {
	p.prefixParseFunc[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFunc) {
	p.infixParseFunc[tokenType] = fn
}

func (p *Parser) nextToken() {
	p.currToken = p.peekedToken
	p.peekedToken = p.l.AdvanceToken()
}
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.nextToken()
	p.nextToken()
	p.prefixParseFunc = make(map[token.TokenType]prefixParseFunc)
	p.registerInfix(token.IDENT, p.parseIdentifier)
	return p
}

func (p *Parser) parseIdentifier() tree.Expression {
	return &tree.Identifier{Token: p.currToken, Value: p.currToken.Literal}
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.currToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekedToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}

func (p *Parser) parseLetStatement() *tree.LetStatement {
	statement := &tree.LetStatement{Token: p.currToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	statement.Name = &tree.Identifier{Token: p.currToken, Value: p.currToken.Literal}

	if !p.expectPeek(token.EQ) {
		return nil
	}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

func (p *Parser) parseReturnStatement() *tree.ReturnStatement {
	st := &tree.ReturnStatement{Token: p.currToken}
	p.nextToken()

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return st
}

func (p *Parser) parseExpression(precedence int) tree.Expression {
	prefix := p.prefixParseFunc[p.currToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()
	return leftExp
}

func (p *Parser) parseExpressionStatement() *tree.ExpressionStatement {
	statement := &tree.ExpressionStatement{Token: p.currToken}

	statement.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

func (p *Parser) ParseStatement() tree.Statement {
	switch p.currToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekErrors(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekedToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) ParseProgram() *tree.Program {
	program := &tree.Program{}
	program.Statements = []tree.Statement{}

	for p.currToken.Type != token.EOF {
		statement := p.ParseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}

	return program
}
