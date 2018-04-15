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
	switch l.ch {
	case '=':
		tok = simpleToken(token.ASSIGN, l.ch)
	case '+':
		tok = simpleToken(token.PLUS, l.ch)
	case ',':
		tok = simpleToken(token.COMMA, l.ch)
	case ';':
		tok = simpleToken(token.SEMICOLON, l.ch)
	case '(':
		tok = simpleToken(token.LPAREN, l.ch)
	case ')':
		tok = simpleToken(token.RPAREN, l.ch)
	case '{':
		tok = simpleToken(token.LBRACE, l.ch)
	case '}':
		tok = simpleToken(token.RBRACE, l.ch)
	case 0:
		tok = newToken(token.EOF, "")
	default:
		if isLetter(l.ch) {
			literal := l.readIdentifier()
			tok = newToken(token.LookupIdentifier(literal), literal)
		} else if isNumber(l.ch) {
			tok = newToken(token.INT, l.readNumber())
		} else {
			tok = simpleToken(token.ILLEGAL, l.ch)
		}
	}

	// Only advance lexer position if we have just read a token
	// that does not manage its own position
	// i.e. keywords, ints, etc do final increment already
	if !managesOwnPosition(tok.Type) {
		l.readChar()
	}
	return tok
}

func simpleToken(tokenType token.TokenType, ch byte) token.Token {
	return newToken(tokenType, string(ch))
}

func newToken(tokenType token.TokenType, lit string) token.Token {
	return token.Token{Type: tokenType, Literal: lit}
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
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isNumber(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
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

func managesOwnPosition(tokenType token.TokenType) bool {
	return token.IsKeyword(tokenType) || tokenType == token.IDENT || tokenType == token.INT
}
