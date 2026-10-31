package lexer

import "monkey/token"

// what the lexer "sees" once it runs off the end of the input.
const eof byte = 0

// this maps characters that always form a one-character token.
var singleCharTokens = [256]token.TokenKind{
	'+': token.PLUS,
	'-': token.MINUS,
	'*': token.ASTERISK,
	'/': token.SLASH,
	'<': token.LT,
	'>': token.GT,
	',': token.COMMA,
	';': token.SEMICOLON,
	'(': token.LPAREN,
	')': token.RPAREN,
	'{': token.LBRACE,
	'}': token.RBRACE,
}

type Lexer struct {
	input   string
	pos     int  // index of ch in input
	nextPos int  // index of the byte after ch
	ch      byte // byte under examination (eof when exhausted)
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.advance() // prime ch with the first byte
	return l
}

// NextToken returns the next token and leaves the lexer on the first
// byte after it. Every scanning helper below keeps that invariant.
func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()

	switch {
	case l.ch == eof:
		return token.Token{Kind: token.EOF}
	case isAsciiAlphabet(l.ch):
		return l.scanIdentifier()
	case isAsciiDigit(l.ch):
		return l.scanNumber()
	case l.ch == '=':
		return l.scanOneOrTwoChar('=', token.EQ, token.ASSIGN)
	case l.ch == '!':
		return l.scanOneOrTwoChar('=', token.NOT_EQ, token.BANG)
	}

	kind := singleCharTokens[l.ch]
	if kind == "" {
		kind = token.ILLEGAL
	}
	return l.consumeChar(kind)
}

func (l *Lexer) scanIdentifier() token.Token {
	literal := l.readWhile(isAsciiAlphabet)
	return token.Token{Kind: token.LookupIdent(literal), Literal: literal}
}

func (l *Lexer) scanNumber() token.Token {
	return token.Token{Kind: token.INT, Literal: l.readWhile(isAsciiDigit)}
}

// This handles operators like "=" vs "==" and "!" vs "!=".
//
// If the next byte is `second`, it emits the two-char token (pairKind), otherwise it emits the one-char token (singleKind).
func (l *Lexer) scanOneOrTwoChar(second byte, pairKind, singleKind token.TokenKind) token.Token {
	if l.peek() != second {
		return l.consumeChar(singleKind)
	}

	start := l.pos
	l.advanceBy(2)
	return token.Token{Kind: pairKind, Literal: l.input[start:l.pos]}
}

// consumeChar emits the current byte as a token of the given kind.
func (l *Lexer) consumeChar(kind token.TokenKind) token.Token {
	tok := token.Token{Kind: kind, Literal: string(l.ch)}
	l.advance()
	return tok
}

// advanceBy moves the cursor forward n bytes.
func (l *Lexer) advanceBy(n int) {
	// Go 1.22+; use `for i := 0; i < n; i++` on older versions
	for range n {
		l.advance()
	}
}

// advance moves the cursor forward by 1 byte.
func (l *Lexer) advance() {
	l.ch = l.peek()
	l.pos = l.nextPos
	l.nextPos++
}

// peek looks at the next byte without consuming it.
func (l *Lexer) peek() byte {
	if l.nextPos >= len(l.input) {
		return eof
	}
	return l.input[l.nextPos]
}

// readWhile consumes bytes while the predicate holds, and returns what it consumed.
func (l *Lexer) readWhile(predicate func(byte) bool) string {
	start := l.pos
	for predicate(l.ch) {
		l.advance()
	}
	return l.input[start:l.pos]
}

func (l *Lexer) skipWhitespace() {
	for isAsciiWhitespace(l.ch) {
		l.advance()
	}
}

func isAsciiWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isAsciiDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isAsciiAlphabet(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
