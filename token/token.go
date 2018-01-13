/*
"token/token.go"
Cody Shepherd

Following Thorsten Ball book
*/

package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string // easy to debug, maybe slower on performance
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Identifiers & literals
	IDENT = "IDENT"
	INT   = "INT"

	//Operators
	ASSIGN = "="
	PLUS   = "+"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
