package lexer

import (
	"fmt"
	"strconv"
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

func (l *Lexer) readNumber() (Token, error) {
	line, column := l.line, l.column
	start := l.offset
	text := l.readWhile(func(r rune) bool { return unicode.IsDigit(r) })

	if l.peek() == '.' {
		if l.peekN(1) == '.' {
			value, err := strconv.ParseInt(text, 10, 64)
			if err != nil {
				return Token{Kind: TokenIllegal, Value: text, Line: line, Column: column, Start: start, End: l.offset}, &LexError{Message: "invalid integer literal", Line: line, Column: column}
			}
			return Token{Kind: TokenInt, Value: value, Line: line, Column: column, Start: start, End: l.offset}, nil
		}

		l.advance()
		if !unicode.IsDigit(l.peek()) {
			return Token{Kind: TokenIllegal, Value: ".", Line: line, Column: column, Start: start, End: l.offset}, &LexError{Message: "invalid float literal", Line: line, Column: column}
		}
		text += "."
		text += l.readWhile(func(r rune) bool { return unicode.IsDigit(r) })
		value, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return Token{Kind: TokenIllegal, Value: text, Line: line, Column: column, Start: start, End: l.offset}, &LexError{Message: "invalid float literal", Line: line, Column: column}
		}
		return Token{Kind: TokenFloat, Value: value, Line: line, Column: column, Start: start, End: l.offset}, nil
	}

	value, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return Token{Kind: TokenIllegal, Value: text, Line: line, Column: column, Start: start, End: l.offset}, &LexError{Message: "invalid integer literal", Line: line, Column: column}
	}

	return Token{Kind: TokenInt, Value: value, Line: line, Column: column, Start: start, End: l.offset}, nil
}

var keywords = map[string]TokenKind{
	"if":     TokenKeyword,
	"else":   TokenKeyword,
	"for":    TokenKeyword,
	"while":  TokenKeyword,
	"return": TokenKeyword,
	"fn":     TokenKeyword,
	"true":   TokenBool,
	"false":  TokenBool,
	"null":   TokenNull,
	"in":     TokenIn,
	"is":     TokenIs,
	"match":  TokenMatch,
}

func isIdentStart(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}

func isIdentPart(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_'
}

func (l *Lexer) readIdent() Token {
	line, column := l.line, l.column
	start := l.offset
	text := l.readWhile(isIdentPart)

	if kind, ok := keywords[text]; ok {
		return Token{Kind: kind, Value: text, Line: line, Column: column, Start: start, End: l.offset}
	}

	return Token{Kind: TokenIdent, Value: text, Line: line, Column: column, Start: start, End: l.offset}
}

func (l *Lexer) readComment() Token {
	line, column := l.line, l.column
	start := l.offset
	l.advance()
	text := l.readWhile(func(r rune) bool { return r != '\n' })
	return Token{Kind: TokenComment, Value: text, Line: line, Column: column, Start: start, End: l.offset}
}

func (l *Lexer) readString() (Token, error) {
	line, column := l.line, l.column
	start := l.offset
	l.advance()

	var builder strings.Builder
	for {
		r := l.peek()
		if r == 0 {
			return Token{Kind: TokenIllegal, Value: builder.String(), Line: line, Column: column, Start: start, End: l.offset}, &LexError{Message: "unterminated string", Line: line, Column: column}
		}
		if r == '"' {
			l.advance()
			break
		}
		if r == '\\' {
			l.advance()
			switch l.peek() {
			case 'n':
				builder.WriteRune('\n')
			case 't':
				builder.WriteRune('\t')
			case '"':
				builder.WriteRune('"')
			case '\\':
				builder.WriteRune('\\')
			default:
				builder.WriteRune(l.peek())
			}
			if l.peek() != 0 {
				l.advance()
			}
			continue
		}
		builder.WriteRune(l.advance())
	}

	return Token{Kind: TokenString, Value: builder.String(), Line: line, Column: column, Start: start, End: l.offset}, nil
}

func (l *Lexer) readWord() Token {
	line, column := l.line, l.column
	start := l.offset

	if isIdentStart(l.peek()) {
		return l.readIdent()
	}

	if unicode.IsDigit(l.peek()) {
		if token, err := l.readNumber(); err == nil {
			return token
		}
	}

	text := l.readWhile(func(r rune) bool {
		return !unicode.IsSpace(r) && !strings.ContainsRune("(){}[],:;.\"", r) && r != '#'
	})

	if text == "" {
		r := l.advance()
		return Token{Kind: TokenIllegal, Value: string(r), Line: line, Column: column, Start: start, End: l.offset}
	}

	return Token{Kind: TokenWord, Value: text, Line: line, Column: column, Start: start, End: l.offset}
}

func (l *Lexer) Next() (Token, error) {
	l.skipWhitespace()
	line, column := l.line, l.column
	start := l.offset

	switch current := l.peek(); current {
	case 0:
		return Token{Kind: TokenEOF, Line: line, Column: column, Start: start, End: start}, nil
	case '#':
		return l.readComment(), nil
	case '"':
		return l.readString()
	case '+':
		l.advance()
		return Token{Kind: TokenPlus, Value: "+", Line: line, Column: column, Start: start, End: l.offset}, nil
	case '-':
		l.advance()
		return Token{Kind: TokenMinus, Value: "-", Line: line, Column: column, Start: start, End: l.offset}, nil
	case '*':
		l.advance()
		return Token{Kind: TokenStar, Value: "*", Line: line, Column: column, Start: start, End: l.offset}, nil
	case '/':
		l.advance()
		return Token{Kind: TokenSlash, Value: "/", Line: line, Column: column, Start: start, End: l.offset}, nil
	case '=':
		l.advance()
		if l.peek() == '=' {
			l.advance()
			return Token{Kind: TokenEq, Value: "==", Line: line, Column: column, Start: start, End: l.offset}, nil
		}
		if l.peek() == '>' {
			l.advance()
			return Token{Kind: TokenArrow, Value: "=>", Line: line, Column: column, Start: start, End: l.offset}, nil
		}
		return Token{Kind: TokenAssign, Value: "=", Line: line, Column: column, Start: start, End: l.offset}, nil
	case '!':
		l.advance()
		if l.peek() == '=' {
			l.advance()
			return Token{Kind: TokenNotEq, Value: "!=", Line: line, Column: column, Start: start, End: l.offset}, nil
		}
		return Token{Kind: TokenBang, Value: "!", Line: line, Column: column, Start: start, End: l.offset}, nil
	case '<':
		l.advance()
		if l.peek() == '=' {
			l.advance()
			return Token{Kind: TokenLtEq, Value: "<=", Line: line, Column: column, Start: start, End: l.offset}, nil
		}
		return Token{Kind: TokenLt, Value: "<", Line: line, Column: column, Start: start, End: l.offset}, nil
	case '>':
		l.advance()
		if l.peek() == '=' {
			l.advance()
			return Token{Kind: TokenGtEq, Value: ">=", Line: line, Column: column, Start: start, End: l.offset}, nil
		}
		return Token{Kind: TokenGt, Value: ">", Line: line, Column: column, Start: start, End: l.offset}, nil
	case '&':
		l.advance()
		if l.peek() == '&' {
			l.advance()
			return Token{Kind: TokenAmpAmp, Value: "&&", Line: line, Column: column, Start: start, End: l.offset}, nil
		}
		return Token{Kind: TokenIllegal, Value: "&", Line: line, Column: column, Start: start, End: l.offset}, &LexError{Message: "unexpected '&'", Line: line, Column: column}
	case '|':
		l.advance()
		if l.peek() == '|' {
			l.advance()
			return Token{Kind: TokenPipePipe, Value: "||", Line: line, Column: column, Start: start, End: l.offset}, nil
		}
		if l.peek() == '>' {
			l.advance()
			return Token{Kind: TokenPipe, Value: "|>", Line: line, Column: column, Start: start, End: l.offset}, nil
		}
		return Token{Kind: TokenIllegal, Value: "|", Line: line, Column: column, Start: start, End: l.offset}, &LexError{Message: "unexpected '|'", Line: line, Column: column}
	case '(':
		l.advance()
		return Token{Kind: TokenLParen, Value: "(", Line: line, Column: column, Start: start, End: l.offset}, nil
	case ')':
		l.advance()
		return Token{Kind: TokenRParen, Value: ")", Line: line, Column: column, Start: start, End: l.offset}, nil
	case '{':
		l.advance()
		return Token{Kind: TokenLBrace, Value: "{", Line: line, Column: column, Start: start, End: l.offset}, nil
	case '}':
		l.advance()
		return Token{Kind: TokenRBrace, Value: "}", Line: line, Column: column, Start: start, End: l.offset}, nil
	case '[':
		l.advance()
		return Token{Kind: TokenLBracket, Value: "[", Line: line, Column: column, Start: start, End: l.offset}, nil
	case ']':
		l.advance()
		return Token{Kind: TokenRBracket, Value: "]", Line: line, Column: column, Start: start, End: l.offset}, nil
	case ',':
		l.advance()
		return Token{Kind: TokenComma, Value: ",", Line: line, Column: column, Start: start, End: l.offset}, nil
	case ';':
		l.advance()
		return Token{Kind: TokenSemicolon, Value: ";", Line: line, Column: column, Start: start, End: l.offset}, nil
	case ':':
		l.advance()
		return Token{Kind: TokenColon, Value: ":", Line: line, Column: column, Start: start, End: l.offset}, nil
	case '.':
		l.advance()
		if l.peek() == '.' {
			l.advance()
			if l.peek() == '.' {
				l.advance()
				return Token{Kind: TokenDotDotDot, Value: "...", Line: line, Column: column, Start: start, End: l.offset}, nil
			}
			return Token{Kind: TokenDotDot, Value: "..", Line: line, Column: column, Start: start, End: l.offset}, nil
		}
		return Token{Kind: TokenDot, Value: ".", Line: line, Column: column, Start: start, End: l.offset}, nil
	default:
		if unicode.IsLetter(current) || current == '_' {
			return l.readIdent(), nil
		}
		if unicode.IsDigit(current) {
			return l.readNumber()
		}
		ch := l.advance()
		return Token{Kind: TokenIllegal, Value: string(ch), Line: line, Column: column, Start: start, End: l.offset}, &LexError{Message: fmt.Sprintf("unexpected character %q", ch), Line: line, Column: column}
	}
}

func Lex(input string) ([]Token, error) {
	le := New(input)
	var tokens []Token

	for {
		token, err := le.Next()
		if token.Kind == TokenEOF {
			tokens = append(tokens, token)
			return tokens, nil
		}
		if err != nil {
			return nil, err
		}
		if token.Kind != TokenComment {
			tokens = append(tokens, token)
		}
	}
}
