package lexer

type Lexer struct {
	input        string
	position     int
	readPosition int
	char         byte
}
