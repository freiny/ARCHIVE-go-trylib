package fStrings

import (
	"strings"
	"testing"
)

func ContainsAny(s, chars string) bool {
	if IndexAny(s, chars) != -1 {
		return true
	}
	return false
}

var containsAnyTests = []struct {
	s, chars string
	out      bool
}{
	{"", "", false},
	{"", "a", false},
	{"a", "", false},
	{"a", "a", true},
	{"abcde", "abcde", true},
	{"abcde", "c", true},
	{"abcde", "dcb", true},
	{"abcde", "bde", true},
	{"abcde", "wxyzbcdq", true},
	{"abcde", "qxyz", false},
}

func TestLibContainsAny(t *testing.T) {
	for _, test := range containsAnyTests {
		actual := ContainsAny(test.s, test.chars)
		libout := strings.ContainsAny(test.s, test.chars)
		if actual != libout {
			t.Errorf("ContainsAny(%q, %q) = %t ; %t wanted", test.s, test.chars, actual, libout)
		}
	}

}

func TestContainsAny(t *testing.T) {
	for _, test := range containsAnyTests {
		actual := ContainsAny(test.s, test.chars)
		if actual != test.out {
			t.Errorf("ContainsAny(%q, %q) = %t ; %t wanted", test.s, test.chars, actual, test.out)
		}
	}
}

func BenchmarkContainsAny(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ContainsAny("abcdefghijklmnopqrstuvwxyz", "ABCDEFgHI")
	}
}
