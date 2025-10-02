package main

import (
	"errors"
	"testing"
)

func compare_slices[E comparable](t *testing.T, w []E, a []E) {
	if len(w) != len(a) {
		t.Fatalf("actual list does not have correct number of elements, want: %v, actual: %v", w, a)
	}

	for i, we := range w {
		ae := a[i]

		if we != ae {
			t.Errorf("actual list's %d element is not correct, want: %v, actual: %v", i, we, ae)
		}
	}
}

func TestTokenize(t *testing.T) {
	sql := "select * from testtable"
	want := []Token{
		{SELECT, "select"},
		{IDENTIFIER, "*"},
		{FROM, "from"},
		{IDENTIFIER, "testtable"},
		{EOF, ""},
	}
	l := NewLexer(sql)
	actual, err := l.tokenize()

	if err != nil {
		t.Fatalf("received unexpected error, %v", err)
	}
	compare_slices(t, want, actual)
}

func TestIsLetter(t *testing.T) {
	var tests = []struct {
		name string
		in   rune
		want bool
	}{
		{"a", 'a', true},
		{"4", '4', false},
		{"P", 'P', true},
		{"*", '*', true},
		{"0", '0', false},
		{"space", ' ', false},
		{"newline", '\n', false},
		{"tab", '\t', false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := isLetter(tt.in)

			if tt.want != a {
				t.Errorf("%s: got: %v, want: %v", tt.name, a, tt.want)
			}
		})
	}
}

func TestIsDigit(t *testing.T) {
	var tests = []struct {
		name string
		in   rune
		want bool
	}{
		{"a", 'a', false},
		{"4", '4', true},
		{"P", 'P', false},
		{"*", '*', false},
		{"0", '0', true},
		{"1", '1', true},
		{"space", ' ', false},
		{"newline", '\n', false},
		{"tab", '\t', false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := isDigit(tt.in)

			if tt.want != a {
				t.Errorf("%s: got: %v, want: %v", tt.name, a, tt.want)
			}
		})
	}
}

func TestIsWhitespace(t *testing.T) {
	var tests = []struct {
		name string
		in   rune
		want bool
	}{
		{"a", 'a', false},
		{"4", '4', false},
		{"P", 'P', false},
		{"*", '*', false},
		{"0", '0', false},
		{"1", '1', false},
		{"space", ' ', true},
		{"newline", '\n', true},
		{"tab", '\t', true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := isWhitespace(tt.in)

			if tt.want != a {
				t.Errorf("%s: got: %v, want: %v", tt.name, a, tt.want)
			}
		})
	}
}

func TestReadIdentifier(t *testing.T) {
	tests := []struct {
		name        string
		in          []rune
		want_n      int
		want_string string
	}{
		{"abcd", []rune("abcd"), 4, "abcd"},
		{"x", []rune{'x'}, 1, "x"},
		{"name,anothername", []rune("name,anothername"), 4, "name"},
		{"abc ijk", []rune("abc"), 3, "abc"},
		{"a_name_1, a_name_2", []rune("a_name_1, a_name_2"), 8, "a_name_1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a_n, a_string, err := readIdentifier(tt.in)

			if err != nil {
				t.Fatalf("%s: unexpected error, %v", tt.name, err)
			}

			if tt.want_n != a_n {
				t.Fatalf("%s: got: %d, want: %d", tt.name, a_n, tt.want_n)
			}

			if tt.want_string != a_string {
				t.Fatalf("%s: got: %s, want: %s", tt.name, a_string, tt.want_string)
			}
		})
	}
}

func TestReadIdentifierError(t *testing.T) {
	_, _, err := readIdentifier([]rune("1name"))

	if !errors.Is(err, ErrInvalidIdentifier) {
		t.Fatalf("expected ErrInvalidIdentifier error, got %v", err)
	}

	var lexErr *LexError
	if !errors.As(err, &lexErr) {
		t.Fatalf("expected a LexError, got %v", err)
	}
}
