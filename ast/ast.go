package ast

import "go/token"

type Node interface {
	TokenLiteral() string
}

type Program struct {
	Statements []Statement
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()

	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}
