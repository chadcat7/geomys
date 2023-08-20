package parser

import (
	"fmt"
	"geomys/lexer"
	"geomys/token"
	"geomys/tree"
)

type Parser struct {
	l           *lexer.Lexer
	currToken   token.Token
	peekedToken token.Token
	errors      []string
}

func (p *Parser) nextToken() {
	p.currToken = p.peekedToken
	p.peekedToken = p.l.AdvanceToken()
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.nextToken()
	p.nextToken()
	return p
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

func (p *Parser) ParseStatement() tree.Statement {
	switch p.currToken.Type {
	case token.LET:
		return p.parseLetStatement()
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
