package parser

import (
	"github.com/matt-snider/monkey/ast"
	"github.com/matt-snider/monkey/lexer"
	"github.com/matt-snider/monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	currToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens so currToken and peekToken are set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) Parse() *ast.Program {
	program := ast.Program{}

	for p.currToken.Type != token.EOF {
		statement := p.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}

	return &program
}

func (p *Parser) parseStatement() *ast.LetStatement {
	switch p.currToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		// TODO: error handling
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	letStatement := ast.LetStatement{Token: p.currToken}

	if p.peekToken.Type != token.IDENT {
		// TODO: error handling
		return nil
	}
	letStatement.Name = &ast.Identifier{
		Token: p.peekToken,
		Value: p.peekToken.Literal,
	}

	p.nextToken()
	p.nextToken()
	if p.currToken.Type != token.ASSIGN {
		// TODO: error handling
		return nil
	}

	return &letStatement
}
