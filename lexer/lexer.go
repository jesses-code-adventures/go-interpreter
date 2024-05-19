package lexer

import "github.com/jesses-code-adventures/go-interpreter/token"

type Lexer struct {
    input string
    inputLength int
    position int // current position in input (points to current char)
    readPosition int // current reading position (points to after current char)
    ch byte // current char value. only supports ASCII, if we want to support UTF-8/UNICODE we must use a rune.
}

func NewLexer(input string) *Lexer {
    l := &Lexer{input: input, inputLength: len(input)}
    l.readChar()
    return l
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

func (l *Lexer) NextToken() token.Token {
    tok := token.TokenFromChar(l.ch)
    l.readChar()
    return tok
}
