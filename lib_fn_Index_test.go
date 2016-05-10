package fStrings

import (
	"strings"
	"testing"
)

func Index(s, sep string) int {
	sepLen := len(sep)
	sLen := len(s)

	if sLen < sepLen {
		return -1
	}
	if s == sep || sep == "" {
		return 0
	}

	for i := 0; i <= sLen-sepLen; i++ {
		if s[i:i+sepLen] == sep {
			return i
		}
	}
	return -1
}

var indexTests = []struct {
	s, sep string
	out    int
}{
	{"", "", 0},
	{"", "a", -1},
	{"a", "", 0},
	{"a", "a", 0},
	{"abcde", "", 0},
	{"abcde", "abcde", 0},
	{"abcde", "bc", 1},
	{"abcde", "ae", -1},
	{"abcde", "c", 2},
	{"abcde", "e", 4},
	{"abcde", "k", -1},
}

func TestLibIndex(t *testing.T) {

	for _, test := range indexTests {
		actual := Index(test.s, test.sep)
		libout := strings.Index(test.s, test.sep)
		if actual != libout {
			t.Errorf("Index(%q, %q) = %d ; %d wanted", test.s, test.sep, actual, libout)
		}

	}
}

func TestIndex(t *testing.T) {
	for _, test := range indexTests {
		actual := Index(test.s, test.sep)
		if actual != test.out {
			t.Errorf("Index(%q, %q) = %d ; %d wanted", test.s, test.sep, actual, test.out)
		}
	}
}

func BenchmarkIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Index("abcdefghijklmnopqrstuvwxyz", "st")
	}
}
