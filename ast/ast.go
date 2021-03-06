package ast

import (
	"bytes"

	"github.com/matt-snider/monkey/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

/**
 * Program
 */

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var buf bytes.Buffer
	for _, statement := range p.Statements {
		buf.WriteString(statement.String())
	}
	return buf.String()
}

/**
 * Identifier
 */

type Identifier struct {
	Token token.Token
	Value string
}

func (id *Identifier) expressionNode() {}

func (id *Identifier) TokenLiteral() string {
	return id.Token.Literal
}

func (id *Identifier) String() string {
	return id.Value
}

/**
 * IntegerLiteral
 */

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.TokenLiteral()
}

/**
 * ExpressionStatement
 */

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

/**
 * LetStatement
 */

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var buf bytes.Buffer
	buf.WriteString(ls.TokenLiteral())
	buf.WriteString(" ")
	buf.WriteString(ls.Name.String())
	buf.WriteString(" = ")
	if ls.Value != nil {
		buf.WriteString(ls.Value.String())
	}
	buf.WriteString(";")
	return buf.String()
}

/**
 * ReturnStatement
 */

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var buf bytes.Buffer
	buf.WriteString(rs.TokenLiteral())
	buf.WriteString(" ")
	if rs.Value != nil {
		buf.WriteString(rs.Value.String())
	}
	buf.WriteString(";")
	return buf.String()
}
