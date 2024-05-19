package lexer

import (
	"strings"

	"github.com/jesses-code-adventures/go-interpreter/token"
)

func isLetter(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_'
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func isDoubleableSymbol(c byte) bool {
	doubleable := "!="
	return strings.Contains(doubleable, string(c))
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
	if isDigit(l.ch) {
		ident := l.readInteger() // Reads all required chars, we don't call l.readChar() before returning
		tok := token.TokenFromInteger(ident)
		return *tok
	} else if isLetter(l.ch) {
		ident := l.readIdentifier() // Reads all required chars, we don't call l.readChar() before returning
		tok := token.TokenFromIdentifierString(ident)
		return *tok
	} else if isDoubleableSymbol(l.ch) && l.peekChar() == '=' {
		curr := l.ch
		l.readChar()
		tok := token.TokenFromDoubledSymbol(curr, l.ch)
		l.readChar() // Call l.readChar before returning so we don't re-read the same symbol
		return *tok
	} else if token.IsSymbol(l.ch) {
		tok := token.TokenFromChar(l.ch)
		l.readChar() // Call l.readChar before returning so we don't re-read the same symbol
		return *tok
	} else {
		tok := token.TokenFromIllegal(l.ch)
		l.readChar() // Call l.readChar before returning so we don't re-read the same symbol
		return *tok
	}
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

func (l *Lexer) peekChar() byte {
	if l.readPosition >= l.inputLength {
		return 0
	}
	return l.input[l.readPosition]
}
