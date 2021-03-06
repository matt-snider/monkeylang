package parser

import (
	"testing"

	"github.com/matt-snider/monkey/ast"
	"github.com/matt-snider/monkey/lexer"
)

/**
 * General helpers
 */
func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf(" - error: %q", msg)
	}
	t.FailNow()
}

/**
 * LetStatements
 */
func TestParsingLetStatements(t *testing.T) {
	l := lexer.New(`
		let x = 5;
		let y = 10;
		let foobar = 1000;
	`)
	p := New(l)
	program := p.Parse()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatal("Parse() returned an empty program (nil)")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("Parse() should return %d statements, got=%d",
			3, len(program.Statements))
	}

	// Check statements
	expected := []struct {
		identifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, expectation := range expected {
		statement := program.Statements[i]
		if !testLetStatement(t, statement, expectation.identifier) {
			return
		}
	}
}

func TestParsingErroneousLetStatement(t *testing.T) {
	l := lexer.New(`
		let x;
		let
	`)
	p := New(l)
	p.Parse()

	if len(p.Errors()) != 2 {
		t.Fatalf("Expected %d errors, got %d", 2, len(p.Errors()))
	}

	expectedErrors := []string{
		"expected next token to be =, got ;",
		"expected next token to be IDENT, got EOF",
	}
	for i, expectation := range expectedErrors {
		actual := p.Errors()[i]
		if expectation != actual {
			t.Errorf("Expected parser error %d to be '%s', got '%s'",
				i, expectation, actual)
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, expectedIdentifier string) bool {
	if s == nil {
		t.Errorf("Expected identifier to be %s, got nil instead", expectedIdentifier)
		return false
	}
	if s.TokenLiteral() != "let" {
		t.Errorf("Expected 'let' got=%s", s.TokenLiteral())
		return false
	}

	letStatement, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("Expected ast.LetStatement, got=%T", s)
		return false
	}

	// Check that name matches expectation
	if letStatement.Name.Value != expectedIdentifier {
		t.Errorf("letStatement.Name.Value not '%s', got=%s",
			expectedIdentifier, letStatement.Name.Value)
		return false
	}
	if letStatement.Name.TokenLiteral() != expectedIdentifier {
		t.Errorf("letStatement.TokenLiteral() not '%s', got=%s",
			expectedIdentifier, letStatement.TokenLiteral())
		return false
	}

	return true
}

/**
 * ReturnStatement
 */

func TestParsingReturnStatements(t *testing.T) {
	l := lexer.New(`
		return 5;
		return add(5, 3);
	`)
	p := New(l)
	program := p.Parse()

	if len(program.Statements) != 2 {
		t.Fatalf("Parse() should return a program with %d statements, got %d",
			2, len(program.Statements))
	}

	for i, statement := range program.Statements {
		returnStatement, ok := statement.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("Expected statement %d to be a ReturnStatement, got %T",
				i, statement)
		}
		if returnStatement.TokenLiteral() != "return" {
			t.Errorf("returnStatement.TokenLiteral() not 'return', got %q", returnStatement.TokenLiteral())
		}
	}
}

/**
 * IdentifierExpression
 */

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := New(l)
	program := p.Parse()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program should have %d statements, got %d", 1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] should be an ast.ExpressionStatement, got %T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("expression should be an *ast.Identifier, got %T", stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Fatalf("ident.Value should be 'foobar', got %s", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Fatalf("ident.TokenLiteral should be 'foobar', got %s", ident.TokenLiteral())
	}
}

/**
 * IntegerLiteralExpression
 */

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5"

	l := lexer.New(input)
	p := New(l)
	program := p.Parse()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program should have %d statements, got %d", 1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] should be an ast.ExpressionStatement, got %T", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("expression should be an *ast.IntegerLiteral, got %T", stmt.Expression)
	}
	if literal.Value != 5 {
		t.Fatalf("literal.Value should be 5, got %d", literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Fatalf("literal.TokenLiteral() should be '5', got '%s'", literal.TokenLiteral())
	}
}
