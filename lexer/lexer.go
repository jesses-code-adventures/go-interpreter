package lexer

import "github.com/jesses-code-adventures/go-interpreter/token"

func isLetter(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_'
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

type Lexer struct {
	input        string
	inputLength  int
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position (points to after current char)
	ch           byte // current char value. only supports ASCII, if we want to support UTF-8/UNICODE we must use a rune.
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input, inputLength: len(input)}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	l.eatWhitespace()
	tok := token.TokenFromChar(l.ch)
	if tok == nil {
		if isDigit(l.ch) {
			ident := l.readInteger()
			tok = token.TokenFromInteger(ident)
			return *tok
		} else if isLetter(l.ch) {
			ident := l.readIdentifier()
			tok = token.TokenFromIdentifierString(ident)
			return *tok
		} else {
			tok = token.TokenFromIllegal(l.ch)
		}
	}
	l.readChar()
	return *tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= l.inputLength {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readIdentifier() string {
	startPosition := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[startPosition:l.position]
}

func (l *Lexer) readInteger() string {
	startPosition := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[startPosition:l.position]
}

func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
