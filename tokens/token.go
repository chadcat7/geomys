package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	IDENT   = "IDENT"
	INT     = "INT"
	FLOAT   = "FLOAT"

	EQUALS   = "="
	PLUS     = "+"
	MINUS    = "-"
	DIVID    = "/"
	MULTIPLY = "*"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"

	LBRACKET = "{"
	RBRACKET = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
