package main

import (
	"errors"
	"fmt"
)

type TokenType int

const (
	SELECT TokenType = iota
	IDENTIFIER
	COMMA
	FROM
	WHERE
	GT
	NUMBER
	STRING
	EQ
	LT
	EOF
)

type Token struct {
	kind  TokenType
	value string
}

var (
	ErrInvalidIdentifier = errors.New("invalid identifier")
)

type LexError struct {
	Message string
	Pos     int
	Err     error
}

func (e *LexError) Error() string {
	return fmt.Sprintf("lex error at position %d: %s", e.Pos, e.Message)
}

func (e *LexError) Unwrap() error { return e.Err }

type Lexer struct {
	input string
	runes []rune
	pos   int
}

func NewLexer(query string) Lexer {
	return Lexer{
		input: query,
		runes: []rune(query),
		pos:   0,
	}
}

func (l *Lexer) tokenize() ([]Token, error) {
	tokens := make([]Token, 0)

	// for l.pos < len(l.runes) {
	// 	c := l.runes[l.pos]

	// 	switch {
	// 	case isWhitespace(c):
	// 		l.pos += 1
	// 	case isLetter(c):
	// 		n, ident, err := readIdentifier(l.runes[l.pos:])

	// 		if err != nil {
	// 			return tokens, err
	// 		}

	// 		l.pos += n

	// 	}
	// }
	return tokens, nil
}

func isLetter(c rune) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '*' {
		return true
	}
	return false
}

func isDigit(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func isWhitespace(c rune) bool {
	if c == ' ' || c == '\n' || c == '\t' {
		return true
	}
	return false
}

func readIdentifier(runes []rune) (int, string, error) {
	ident := make([]rune, 0)
	i := 0

	if len(runes) > 0 && isDigit(runes[0]) {
		return i, "", &LexError{Message: "Identifier cannot start with a digit", Pos: 0, Err: ErrInvalidIdentifier}
	}

	for _, c := range runes {
		if isLetter(c) || isDigit(c) || c == '_' {
			ident = append(ident, c)
			i += 1
		} else {
			break
		}
	}

	return i, string(ident), nil
}
