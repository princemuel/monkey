package token

type TokenKind string

const (
	ILLEGAL TokenKind = "ILLEGAL"
	EOF     TokenKind = "EOF"

	// Identifiers + literals
	IDENT TokenKind = "IDENT" // add, foobar, x, y, ...
	INT   TokenKind = "INT"   // 1343456

	// Operators
	ASSIGN   TokenKind = "="
	PLUS     TokenKind = "+"
	MINUS    TokenKind = "-"
	BANG     TokenKind = "!"
	ASTERISK TokenKind = "*"
	SLASH    TokenKind = "/"

	LT TokenKind = "<"
	GT TokenKind = ">"

	EQ     TokenKind = "=="
	NOT_EQ TokenKind = "!="

	// Delimiters
	COMMA     TokenKind = ","
	SEMICOLON TokenKind = ";"

	LPAREN TokenKind = "("
	RPAREN TokenKind = ")"
	LBRACE TokenKind = "{"
	RBRACE TokenKind = "}"

	// Keywords
	FUNCTION TokenKind = "FUNCTION"
	LET      TokenKind = "LET"
	TRUE     TokenKind = "TRUE"
	FALSE    TokenKind = "FALSE"
	IF       TokenKind = "IF"
	ELSE     TokenKind = "ELSE"
	RETURN   TokenKind = "RETURN"
)

type Token struct {
	Kind    TokenKind
	Literal string
}

var keywords = map[string]TokenKind{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent returns the keyword kind for ident, or IDENT if it isn't a keyword.
func LookupIdent(ident string) TokenKind {
	if kind, ok := keywords[ident]; ok {
		return kind
	}
	return IDENT
}
