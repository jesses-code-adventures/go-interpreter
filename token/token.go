package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL" // Unknown token or character
	EOF     = "EOF"     // End of file - lexer can stop

	// Identifiers, literals
	IDENT = "IDENT" // User-defined identifier such as add, foobar, x, y etc
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

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
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
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
