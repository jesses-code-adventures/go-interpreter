package token

type TokenType string

const (
    ILLEGAL = "ILLEGAL" // Unknown token or character
    EOF = "EOF" // End of file - lexer can stop

    // Identifiers, literals
    IDENT = "IDENT" // User-defined identifier such as add, foobar, x, y etc
    INT = "INT"

    // Operators
    ASSIGN = "="
    PLUS = "+"

    // Delimiters
    COMMA = ","
    SEMICOLON = ";"
    COLON = ":"
    LPAREN = "("
    RPAREN = ")"
    LSQUIRLY = "{"
    RSQUIRLY = "}"

    // Keywords
    FUNCTION = "FUNCTION"
    LET = "LET"
)

type Token struct {
	Type    TokenType
	Literal string
}

func newToken(tokenType TokenType, literal byte) Token {
    return Token{tokenType, string(literal)}
}

func TokenFromChar(c byte) Token {
    switch c {
	case '=':
	    return newToken(ASSIGN, c)
	case '+':
	    return newToken(PLUS, c)
	case ',':
	    return newToken(COMMA, c)
	case ';':
	    return newToken(SEMICOLON, c)
	case ':':
	    return newToken(COLON, c)
	case '(':
	    return newToken(LPAREN, c)
	case ')':
	    return newToken(RPAREN, c)
	case '{':
	    return newToken(LSQUIRLY, c)
	case '}':
	    return newToken(RSQUIRLY, c)
	case 0:
	    return Token{EOF, ""}
    }
    // TODO: Maybe this could be better
    return Token{}
}
