package token

import "strings"

type TokenType string

const (
	ILLEGAL = "ILLEGAL" // Unknown token or character
	EOF     = "EOF"     // End of file - lexer can stop

	// Identifiers, literals
	IDENT = "IDENT" // User-defined identifier such as add, foobar, x, y etc
	INT   = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	NE       = "!="
	EQ       = "=="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	LPAREN    = "("
	RPAREN    = ")"
	LSQUIRLY  = "{"
	RSQUIRLY  = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	RETURN   = "RETURN"
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	FOR      = "FOR"
	BREAK    = "BREAK"
)

func IsSymbol(c byte) bool {
	return strings.Contains("=+-!*/<>,;:(){}", string(c)) || c == 0
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"false":  FALSE,
	"true":   TRUE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"break":  BREAK,
	"for":    FOR,
}

func lookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

type Token struct {
	Type    TokenType
	Literal string
}

func newToken(tokenType TokenType, literal byte) Token {
	return Token{tokenType, string(literal)}
}

func TokenFromChar(c byte) *Token {
	var tok Token
	switch c {
	case '=':
		tok = newToken(ASSIGN, c)
		return &tok
	case '+':
		tok = newToken(PLUS, c)
		return &tok
	case ',':
		tok = newToken(COMMA, c)
		return &tok
	case ';':
		tok = newToken(SEMICOLON, c)
		return &tok
	case ':':
		tok = newToken(COLON, c)
		return &tok
	case '(':
		tok = newToken(LPAREN, c)
		return &tok
	case ')':
		tok = newToken(RPAREN, c)
		return &tok
	case '{':
		tok = newToken(LSQUIRLY, c)
		return &tok
	case '}':
		tok = newToken(RSQUIRLY, c)
		return &tok
	case '-':
		tok = newToken(MINUS, c)
		return &tok
	case '!':
		tok = newToken(BANG, c)
		return &tok
	case '*':
		tok = newToken(ASTERISK, c)
		return &tok
	case '/':
		tok = newToken(SLASH, c)
		return &tok
	case '<':
		tok = newToken(LT, c)
		return &tok
	case '>':
		tok = newToken(GT, c)
		return &tok
	case 0:
		tok = Token{EOF, ""}
		return &tok
	}
	return nil
}

func TokenFromIdentifierString(s string) *Token {
	tokenType := lookupIdent(s)
	return &Token{tokenType, s}
}

func TokenFromInteger(s string) *Token {
	return &Token{INT, s}
}

func TokenFromIllegal(c byte) *Token {
	tok := newToken(ILLEGAL, c)
	return &tok
}

func TokenFromDoubledSymbol(c1 byte, c2 byte) *Token {
	bytes := make([]byte, 0)
	bytes = append(bytes, c1)
	bytes = append(bytes, c2)
	combined := string(bytes)
	var tok Token
	if c1 == '!' {
		tok = Token{NE, combined}
	} else {
		tok = Token{EQ, combined}
	}
	return &tok
}
