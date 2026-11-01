package ast

import "monkey/token"

// Node is anything in the syntax tree.
type Node interface {
	// TokenLiteral returns the literal of the token this node was built from.
	TokenLiteral() string
}

// Statement is a Node that performs an action without producing a value
// (let, return, ...).
//
// isStatement is a marker method: it has no behavior and exists only so the
// compiler can tell statements from expressions.
type Statement interface {
	Node
	isStatement()
}

// Expression is a Node that produces a value (5, x, a + b, ...).
//
// isExpression is a marker method, like isStatement.
type Expression interface {
	Node
	isExpression()
}

// Program is the root of every AST: a sequence of top-level statements.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) == 0 {
		return ""
	}
	return p.Statements[0].TokenLiteral()
}

// LetStatement is `let <ident> = <value>;`.
type LetStatement struct {
	Token token.Token // the LET token
	Name  *Identifier
	Value Expression
}

func (*LetStatement) isStatement() {}

func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type ReturnStatement struct {
	Token       token.Token // the RETURN token
	ReturnValue Expression
}

func (*ReturnStatement) isStatement() {}

func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// Identifier is a ident, such as `x` or `add`.
type Identifier struct {
	Token token.Token // the IDENT token
	Value string
}

func (*Identifier) isExpression() {}

func (id *Identifier) TokenLiteral() string { return id.Token.Literal }

// Compile-time checks that each type satisfies the interface it claims to.
var (
	_ Node      = (*Program)(nil)
	_ Statement = (*LetStatement)(nil)
	_ Statement = (*ReturnStatement)(nil)
	// _ Statement  = (*ExpressionStatement)(nil)
	_ Expression = (*Identifier)(nil)
	// _ Expression = (*IntegerLiteral)(nil)
)
