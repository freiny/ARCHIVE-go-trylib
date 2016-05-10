package fStrings

import (
	"strings"
	"testing"
)

func ContainsRune(s string, r rune) bool {
	if IndexRune(s, r) != -1 {
		return true
	}
	return false
}

var containsRuneTests = []struct {
	s   string
	r   rune
	out bool
}{
	{"", rune(0), false},
	{"a", rune(0), false},
	{"abc", rune(0), false},
	{"a", 'a', true},
	{"abcde", 'a', true},
	{"abcde", 'c', true},
	{"abcde", 'e', true},
	{"abcde", 'D', false},
}

func TestLibContainsRune(t *testing.T) {
	for _, test := range containsRuneTests {
		actual := ContainsRune(test.s, test.r)
		libout := strings.ContainsRune(test.s, test.r)
		if actual != libout {
			t.Errorf("ContainsRune(%q, %q) = %t ; %t wanted", test.s, test.r, actual, libout)
		}
	}
}

func TestContainsRune(t *testing.T) {
	for _, test := range containsRuneTests {
		actual := ContainsRune(test.s, test.r)
		if actual != test.out {
			t.Errorf("ContainsRune(%q, %q) = %t ; %t wanted", test.s, test.r, actual, test.out)
		}
	}
}

func BenchmarkContainsRune(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ContainsRune("abcdefghijklmnopqrstuvwxyz", 'x')
	}
}
