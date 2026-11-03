package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

var precedences = map[token.TokenKind]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.LPAREN:   CALL,
}

type (
	prefixParseFn func() ast.Expr
	infixParseFn  func(ast.Expr) ast.Expr
)

// Parser builds an AST from the lexer's tokens, using one token of lookahead.
type Parser struct {
	lex    *lexer.Lexer
	cur    token.Token // token under examination
	peek   token.Token // the token after cur
	errors []string

	prefixParseFns map[token.TokenKind]prefixParseFn
	infixParseFns  map[token.TokenKind]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{lex: l, errors: []string{}}

	p.prefixParseFns = make(map[token.TokenKind]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	// Prime the two-token window. The first call fills peek;
	// the second shifts it into cur and fills peek again.
	p.next()
	p.next()
	return p
}

// next slides the window forward by one token.
func (p *Parser) next() {
	p.cur = p.peek
	p.peek = p.lex.NextToken()
}

func (p *Parser) currentTokIs(kind token.TokenKind) bool { return p.cur.Kind == kind }
func (p *Parser) peekTokenIs(kind token.TokenKind) bool  { return p.peek.Kind == kind }

// nextIf moves forward only if the next token is of the given kind,
// and reports whether it did.
func (p *Parser) nextIf(kind token.TokenKind) bool {
	if !p.peekTokenIs(kind) {
		p.peekError(kind)
		return false
	}
	p.next()
	return true
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenKind) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peek.Kind)
	p.errors = append(p.errors, msg)
}

func (p *Parser) noPrefixParseFnError(t token.TokenKind) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

// ParseProgram parses the whole input into a Program.
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}

	for !p.currentTokIs(token.EOF) {
		if stmt := p.parseStatement(); stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.next()
	}
	return program
}

// parseStatement parses the statement starting at p.cur.
// It returns nil if the current token doesn't begin a known statement.
//
// Convention for every parseX method: on return, p.cur is the LAST token
// of what was parsed, so the caller's advance() steps onto the next one.
func (p *Parser) parseStatement() ast.Stmt {
	switch p.cur.Kind {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	}
	return p.parseExpressionStatement()
}

// parseLetStatement parses `let <ident> = <expression>;`.
func (p *Parser) parseLetStatement() ast.Stmt {
	stmt := &ast.LetStmt{Token: p.cur}

	if !p.nextIf(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.cur, Value: p.cur.Literal}

	if !p.nextIf(token.ASSIGN) {
		return nil
	}

	// TODO: parse the expression instead of skipping it
	p.skipExpression()

	return stmt
}

func (p *Parser) parseReturnStatement() ast.Stmt {
	stmt := &ast.ReturnStatement{Token: p.cur}
	p.next()

	// TODO: parse the expression instead of skipping it
	p.skipExpression()
	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.cur}
	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.next()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expr {
	prefix := p.prefixParseFns[p.cur.Kind]
	if prefix == nil {
		return nil
	}
	lhs := prefix()
	return lhs
}

func (p *Parser) parseIdentifier() ast.Expr {
	return &ast.Identifier{Token: p.cur, Value: p.cur.Literal}
}

// skipExpression is a placeholder: it advances to the end of the statement
// (the semicolon, or EOF if the input ends without one).
func (p *Parser) skipExpression() {
	for !p.currentTokIs(token.SEMICOLON) && !p.currentTokIs(token.EOF) {
		p.next()
	}
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peek.Kind]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.cur.Kind]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) registerPrefix(tokenKind token.TokenKind, fn prefixParseFn) {
	p.prefixParseFns[tokenKind] = fn
}
func (p *Parser) registerInfix(tokenKind token.TokenKind, fn infixParseFn) {
	p.infixParseFns[tokenKind] = fn
}
