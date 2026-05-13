package lexer

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
