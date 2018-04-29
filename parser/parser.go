package parser

import (
	"fmt"
	"strconv"

	"github.com/matt-snider/monkey/ast"
	"github.com/matt-snider/monkey/lexer"
	"github.com/matt-snider/monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	currToken token.Token
	peekToken token.Token
	errors    []string
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

	for !p.currTokenIs(token.EOF) {
		statement := p.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}

	return &program
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		// TODO: error handling
		return nil
	}
}

func (p *Parser) currTokenIs(t token.TokenType) bool {
	return p.currToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}

func (p *Parser) peekError(t token.TokenType) {
	error := fmt.Sprintf(
		"expected next token to be %s, got %s",
		t, p.peekToken.Type,
	)
	p.errors = append(p.errors, error)
}

/**
 * Identifier
 */
func (p *Parser) parseIdentifier() *ast.Identifier {
	return &ast.Identifier{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}
}

/**
 *  IntegerLiteral
 */
func (p *Parser) parseIntegerLiteral() *ast.IntegerLiteral {
	value, err := strconv.ParseInt(p.currToken.Literal, 0, 64)
	if err != nil {
		error := fmt.Sprintf("could not parse int literal %q",
			p.currToken.Literal)
		p.errors = append(p.errors, error)
		return nil
	}
	return &ast.IntegerLiteral{
		Token: p.currToken,
		Value: value,
	}
}

/**
 * LetStatement
 */
func (p *Parser) parseLetStatement() *ast.LetStatement {
	letStatement := ast.LetStatement{Token: p.currToken}

	if !p.expectPeek(token.IDENT) {
		p.peekError(token.IDENT)
		return nil
	}
	letStatement.Name = &ast.Identifier{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}

	if !p.expectPeek(token.ASSIGN) {
		p.peekError(token.ASSIGN)
		return nil
	}

	// Skip over expression until EOL
	// TODO: letStatement Expression
	for !p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return &letStatement
}

/**
 * ReturnStatement
 */
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	returnStatement := ast.ReturnStatement{Token: p.currToken}
	return &returnStatement
}
