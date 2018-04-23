package lexer

import (
	"testing"

	"github.com/matt-snider/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `
		let five = 5;
		let ten = 10;

		let add = fn(x, y) {
			x + y;
		};

		let result = add(five, ten);
		if (five == ten) { }
		if (five != ten) { }
		if (five > ten) { }
		if (five < ten) { }

		if (true) { }
		else { }

		let result3 = 10 - 5 * 5 / 3;
		return !false;
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		// if (five == ten) { }
		{token.IF, "IF"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.EQ, "=="},
		{token.IDENT, "ten"},
		{token.LPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},

		// if (five != ten) { }
		{token.IF, "IF"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.NOT_EQ, "!="},
		{token.IDENT, "ten"},
		{token.LPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},

		// if (five > ten) { }
		{token.IF, "IF"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.GT, ">"},
		{token.IDENT, "ten"},
		{token.LPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},

		// if (five < ten) { }
		{token.IF, "IF"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.LT, "<"},
		{token.IDENT, "ten"},
		{token.LPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},

		// if (true) { } else { }
		{token.IF, "IF"},
		{token.LPAREN, "("},
		{token.TRUE, "TRUE"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.ELSE, "ELSE"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},

		// let result3 = 10 - 5 * 5 / 3;
		{token.LET, "LET"},
		{token.IDENT, "result3"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.MINUS, "-"},
		{token.INT, "5"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SLASH, "/"},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},

		// return !false;
		{token.RETURN, "RETURN"},
		{token.BANG, "!"},
		{token.FALSE, "FALSE"},
		{token.SEMICOLON, ";"},

		{token.EOF, ""},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("TestNextToken[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("TestNextToken[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestIllegalToken(t *testing.T) {
	l := New("~")
	tok := l.NextToken()
	if tok.Type != token.ILLEGAL {
		t.Fatalf("TestIllegalToken - tokentype wrong. expected=ILLEGAL, got=%q", tok.Type)
	}
	if tok.Literal != "~" {
		t.Fatalf("TestIllegalToken - literal wrong. expected=~, got=%q", tok.Literal)
	}
}
