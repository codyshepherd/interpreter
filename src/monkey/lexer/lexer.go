/*
lexer/lexer.go

Cody Shepherd

Following the Thorsten Ball book
*/

package lexer

import "monkey/token"

type Lexer struct {
	input        string
	runes        []rune // my addition to allow handling unicode
	position     int    // current position in input (points to current char)
	readPosition int    // current reading position in input (after current char)
	ch           rune
	//ch           byte // current char under examination (Note: byte type means only ASCII supported)
	// would need to use `rune` type to support UTF-8, and change some other stuff (future project?)
}

// Return a new Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input, runes: []rune(input)}
	l.readChar()
	return l
}

// Give us the next character and advance our position in the input string.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.runes) {
		l.ch = 0
	} else {
		l.ch = l.runes[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
