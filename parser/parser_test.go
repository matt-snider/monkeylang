package parser

import (
	"testing"

	"github.com/matt-snider/monkey/ast"
	"github.com/matt-snider/monkey/lexer"
)

func TestParsingLetStatements(t *testing.T) {
	l := lexer.New(`
		let x = 5;
		let y = 10;
		let foobar = 1000;
	`)
	p := New(l)
	program := p.Parse()
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
