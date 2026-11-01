package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

// Parser builds an AST from the lexer's tokens, using one token of lookahead.
type Parser struct {
	lex    *lexer.Lexer
	cur    token.Token // token under examination
	peek   token.Token // the token after cur
	errors []string
}

func New(lex *lexer.Lexer) *Parser {
	p := &Parser{lex: lex, errors: []string{}}

	// Prime the two-token window. The first call fills peek; the second
	// shifts it into cur and fills peek again.
	p.advance()
	p.advance()
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenKind) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peek.Kind)
	p.errors = append(p.errors, msg)
}

// advance slides the window forward by one token.
func (p *Parser) advance() {
	p.cur = p.peek
	p.peek = p.lex.NextToken()
}

// ParseProgram parses the whole input into a Program.
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}

	for !p.currentTokIs(token.EOF) {
		if stmt := p.parseStatement(); stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.advance()
	}
	return program
}

// parseStatement parses the statement starting at p.cur.
// It returns nil if the current token doesn't begin a known statement.
//
// Convention for every parseX method: on return, p.cur is the LAST token
// of what was parsed, so the caller's advance() steps onto the next one.
func (p *Parser) parseStatement() ast.Statement {
	switch p.cur.Kind {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	}
	return nil
}

func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := &ast.ReturnStatement{Token: p.cur}
	p.advance()

	// TODO: parse the expression instead of skipping it
	p.skipExpression()
	return stmt
}

// parseLetStatement parses `let <ident> = <expression>;`.
func (p *Parser) parseLetStatement() ast.Statement {
	stmt := &ast.LetStatement{Token: p.cur}

	if !p.advanceIf(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.cur, Value: p.cur.Literal}

	if !p.advanceIf(token.ASSIGN) {
		return nil
	}

	// TODO: parse the expression instead of skipping it
	p.skipExpression()

	return stmt
}

// skipExpression is a placeholder: it advances to the end of the statement
// (the semicolon, or EOF if the input ends without one).
func (p *Parser) skipExpression() {
	for !p.currentTokIs(token.SEMICOLON) && !p.currentTokIs(token.EOF) {
		p.advance()
	}
}

// advanceIf moves forward only if the next token is of the given kind,
// and reports whether it did.
func (p *Parser) advanceIf(kind token.TokenKind) bool {
	if !p.peekTokenIs(kind) {
		p.peekError(kind)
		return false
	}
	p.advance()
	return true
}

func (p *Parser) currentTokIs(kind token.TokenKind) bool { return p.cur.Kind == kind }
func (p *Parser) peekTokenIs(kind token.TokenKind) bool  { return p.peek.Kind == kind }
