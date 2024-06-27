package lexer

import (
	"monkeyLang/token"
)

type Lexer struct {
	input        string // input (source code) type string.
	position     int    // points to the current position of input string which under examination (current char).
	readPosition int    // points to the next position which would be examined and crrent reading position of input string (next char).
	ch           byte   // current char under examination.
}

/*
Initialize a new Lexer and call readchar so that the character can be examined from the source code (input).
*/
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

/*
Sets the current position of the source code which is examined and increase the readPosition so later the next char  can be examine later.
*/
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) { // If readPosition reaches to end of the input string then there is nothing to read.
		l.ch = 0

	} else { // If readPosition in the boundary of the input string then it would be read.
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

/*
Helper function is use to return the peek value (byte) of the lexer which is readPosition.
*/
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0

	} else {
		return l.input[l.readPosition]
	}
}

/*
NextToken returns the token which would be examined and increase readPosition when readchar is called.
*/
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace() // Skipping the whitespaces until there is no whitespace.

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}

		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}

	case ';':
		tok = newToken(token.SEMICOLON, l.ch)

	case ',':
		tok = newToken(token.COMMA, l.ch)

	case '(':
		tok = newToken(token.LPAREN, l.ch)

	case ')':
		tok = newToken(token.RPAREN, l.ch)

	case '{':
		tok = newToken(token.LBRACE, l.ch)

	case '}':
		tok = newToken(token.RBRACE, l.ch)

	case '+':
		tok = newToken(token.PLUS, l.ch)

	case '-':
		tok = newToken(token.MINUS, l.ch)

	case '*':
		tok = newToken(token.ASTERIK, l.ch)

	case '/':
		tok = newToken(token.SLASH, l.ch)

	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NQ, Literal: string(ch) + string(l.ch)}

		} else {
			tok = newToken(token.BANG, l.ch)
		}

	case '<':
		tok = newToken(token.LT, l.ch)

	case '>':
		tok = newToken(token.GT, l.ch)

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok

		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok

		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()

	return tok
}

/*
Helper function checks if the character is a valid letter or not.
And returns true or false (bool).
*/
func isLetter(ch byte) bool {
	return 'a' <= ch && 'z' >= ch || 'A' <= ch && 'Z' >= ch || ch == '_'
}

/*
Helper function checks if the character is a valid digit or not.
And returns true or false (bool).
*/
func isDigit(ch byte) bool {
	return '0' <= ch && '9' >= ch
}

/*
Helper function reads the indentifier.
And returns the identifier (string).
*/
func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

/*
Helper function reads the number.
And returns the number (string).
*/
func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

/*
Helper function skips the white space.
*/
func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

/*
Helper function to convert type of ch (byte) to (string).
And returns Token.
*/
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
