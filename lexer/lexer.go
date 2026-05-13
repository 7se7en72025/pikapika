package lexer

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type TokenKind int

const (
	TokenInt TokenKind = iota
	TokenFloat
	TokenBool
	TokenNull
	TokenIdent
	TokenKeyword
	TokenIn
	TokenIs
	TokenMatch
	TokenPlus
	TokenMinus
	TokenStar
	TokenSlash
	TokenEq
	TokenNotEq
	TokenLt
	TokenGt
	TokenLtEq
	TokenGtEq
	TokenAssign
	TokenArrow
	TokenBang
	TokenAmpAmp
	TokenPipePipe
	TokenPipe
	TokenLParen
	TokenRParen
	TokenLBrace
	TokenRBrace
	TokenLBracket
	TokenRBracket
	TokenComma
	TokenSemicolon
	TokenColon
	TokenDot
	TokenDotDot
	TokenDotDotDot
	TokenEOF
	TokenIllegal
	TokenWord
	TokenString
	TokenComment
)

type Token struct {
	Kind   TokenKind
	Value  any
	Line   int
	Column int
	Start  int
	End    int
}

type LexError struct {
	Message string
	Line    int
	Column  int
}

func (e *LexError) Error() string {
	return fmt.Sprintf("[line %d, col %d] %s", e.Line, e.Column, e.Message)
}

type Lexer struct {
	source string
	runes  []rune
	index  int
	line   int
	column int
	offset int
}

func New(input string) *Lexer {
	return &Lexer{
		source: input,
		runes:  []rune(input),
		line:   1,
		column: 1,
		offset: 0,
	}
}

func (l *Lexer) peek() rune {
	if l.index >= len(l.runes) {
		return 0
	}
	return l.runes[l.index]
}

func (l *Lexer) peekN(n int) rune {
	idx := l.index + n
	if idx < 0 || idx >= len(l.runes) {
		return 0
	}
	return l.runes[idx]
}

func (l *Lexer) advance() rune {
	if l.index >= len(l.runes) {
		return 0
	}
	r := l.runes[l.index]
	l.index++
	l.offset += utf8.RuneLen(r)
	if r == '\n' {
		l.line++
		l.column = 1
	} else {
		l.column++
	}
	return r
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.peek()) {
		l.advance()
	}
}

func (l *Lexer) readWhile(pred func(rune) bool) string {
	var builder strings.Builder
	for current := l.peek(); current != 0 && pred(current); current = l.peek() {
		builder.WriteRune(l.advance())
	}
	return builder.String()
}
