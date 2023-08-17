package lexer

import (
	"geomys/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	char         byte
}

func checkForAlphabet(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func checkForDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) Advance() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1 // add 1
}

func NewLex(input string) *Lexer {
	l := &Lexer{input: input}
	l.Advance()
	return l
}

func (l *Lexer) readIndentifier() string {
	position := l.position
	for checkForAlphabet(l.char) {
		l.Advance()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for checkForDigit(l.char) {
		l.Advance()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipSpaces() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.Advance()
	}
}

// Advance but doestn actually advance only returns
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) AdvanceToken() token.Token {
	var tok token.Token
	l.skipSpaces()
	switch l.char {
	case '=': // using '' for characters
		if l.peekChar() == '=' {
			ch := l.char
			l.Advance()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.char)}
		} else {
			tok = token.NewTok(token.EQUALS, l.char)
		}
	case ';':
		tok = token.NewTok(token.SEMICOLON, l.char)
	case '(':
		tok = token.NewTok(token.LPAREN, l.char)
	case ')':
		tok = token.NewTok(token.RPAREN, l.char)
	case ',':
		tok = token.NewTok(token.COMMA, l.char)
	case '+':
		tok = token.NewTok(token.PLUS, l.char)
	case '-':
		tok = token.NewTok(token.MINUS, l.char)
	case '/':
		tok = token.NewTok(token.DIVIDE, l.char)
	case '*':
		tok = token.NewTok(token.MULTIPLY, l.char)
	case '{':
		tok = token.NewTok(token.LBRACKET, l.char)
	case '}':
		tok = token.NewTok(token.RBRACKET, l.char)
	case '!':
		if l.peekChar() == '=' {
			ch := l.char
			l.Advance()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.char)}
		} else {
			tok = token.NewTok(token.BANG, l.char)
		}
	case '<':
		tok = token.NewTok(token.LT, l.char)
	case '>':
		tok = token.NewTok(token.GT, l.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if checkForAlphabet(l.char) {
			tok.Literal = l.readIndentifier()
			tok.Type = token.LookIdent(tok.Literal)
			return tok
		} else if checkForDigit(l.char) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = token.NewTok(token.ILLEGAL, l.char)
		}
	}
	l.Advance()
	return tok
}
