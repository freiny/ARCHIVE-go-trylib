package fStrings

import (
	"strings"
	"testing"
)

func IndexRune(s string, r rune) int {
	return Index(s, string(r))
}

var indexRuneTests = []struct {
	s   string
	r   rune
	out int
}{
	{"", rune(0), -1},
	{"a", rune(0), -1},
	{"abc", rune(0), -1},
	{"a", 'a', 0},
	{"abcde", 'a', 0},
	{"abcde", 'c', 2},
	{"abcde", 'e', 4},
	{"abcde", 'D', -1},
}

func TestLibIndexRune(t *testing.T) {
	for _, test := range indexRuneTests {
		actual := IndexRune(test.s, test.r)
		libout := strings.IndexRune(test.s, test.r)
		if actual != libout {
			t.Errorf("IndexRune(%q, %q) = %d ; %d wanted", test.s, test.r, actual, libout)
		}
	}
}

func TestIndexRune(t *testing.T) {
	for _, test := range indexRuneTests {
		actual := IndexRune(test.s, test.r)
		if actual != test.out {
			t.Errorf("IndexRune(%q, %q) = %d ; %d wanted", test.s, test.r, actual, test.out)
		}
	}
}

func BenchmarkIndexRune(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IndexRune("abcdefghijklmnopqrstuvwxyz", 'x')
	}
}
