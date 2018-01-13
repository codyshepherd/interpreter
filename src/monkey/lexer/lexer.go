/*
lexer/lexer.go

Cody Shepherd

Following the Thorsten Ball book
*/

package lexer

import (
	"monkey/token"
	"unicode"
)

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

	l.skipWhitespace()

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
	default:
		if isLetter(l.ch) { // case to check for identifier string
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok // readIdentifier calls readChar() so we want to early return here to avoid an extra call to readChar()
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else { // case to handle unknown input
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// helper for instantiating new Token object
func newToken(tokenType token.TokenType, ch rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// reads an identifier string and advances lexer's positions until it encounters
// a non-letter character (rune)
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return string(l.runes[position:l.position])
}

// checks whether its argument is a letter
// changing this function will have a large impact on parsable language
// Note: gonna allow unicode symbols for now and see where that takes us
func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || unicode.IsSymbol(ch)
}

// eats up whitespace by calling readChar on spaces, tabs, or newlines
// replacing guard from book with a call to unicode library, since I'm importing it already
func (l *Lexer) skipWhitespace() {
	//for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
	for unicode.IsSpace(l.ch) {
		l.readChar()
	}
}

// same as readIdentifier except for numbers instead of strings
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return string(l.runes[position:l.position])
}

// This is just a wrapper for the unicode call... extra function overhead but might allow
// for customization later
func isDigit(ch rune) bool {
	//return '0' <= ch && ch <= '9'
	return unicode.IsDigit(ch)
}
