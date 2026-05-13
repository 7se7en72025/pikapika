package lexer

import "fmt"

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
