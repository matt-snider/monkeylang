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
		"expected token of type =, got ;",
		"expected token of type IDENT, got EOF",
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
