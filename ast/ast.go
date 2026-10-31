package ast

import "monkey/token"

// Node is anything in the syntax tree.
type Node interface {
	TokenLiteral() string
}

// A Node that performs an action without producing a value (let, return, ...).
//
// isStatement is a marker method. It has no behavior and exists only so the compiler can tell a `[Statement]“ from an `[Expression]“.
type Statement interface {
	Node
	isStatement()
}

// A Node that produces a value (5, x, a + b, ...). isExpression is a marker method.
type Expression interface {
	Node
	isExpression()
}

type Program struct{ Statements []Statement }

func (p *Program) TokenLiteral() string {
	if len(p.Statements) == 0 {
		return ""
	}

	return p.Statements[0].TokenLiteral()
}

type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) isStatement()         {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) isExpression()        {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
