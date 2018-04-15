package lexer

import "github.com/matt-snider/monkey/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.eatWhitespace()
	tok.Literal = string(l.ch)
	switch l.ch {
	case '=':
		tok.Type = token.ASSIGN
	case '+':
		tok.Type = token.PLUS
	case ',':
		tok.Type = token.COMMA
	case ';':
		tok.Type = token.SEMICOLON
	case '(':
		tok.Type = token.LPAREN
	case ')':
		tok.Type = token.RPAREN
	case '{':
		tok.Type = token.LBRACE
	case '}':
		tok.Type = token.RBRACE
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
		} else if isNumber(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
		} else {
			tok.Type = token.ILLEGAL
		}
		l.readPosition -= 1
		l.position -= 1
	}
	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) readIdentifier() string {
	var ident string
	for isLetter(l.ch) {
		ident += string(l.ch)
		l.readChar()
	}
	return ident
}

func (l *Lexer) readNumber() string {
	var num string
	for isNumber(l.ch) {
		num += string(l.ch)
		l.readChar()
	}
	return num
}

func (l *Lexer) eatWhitespace() {
	for isWhitespace(l.ch) {
		l.readChar()
	}
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isNumber(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
