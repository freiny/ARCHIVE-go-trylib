package fStrings

import (
	"strings"
	"testing"
)

func HasPrefix(s, prefix string) bool {

	if prefix == "" {
		return true
	}
	if len(s) >= len(prefix) && s[0:len(prefix)] == prefix {
		return true
	}
	return false
}

var hasPrefixTests = []struct {
	s, prefix string
	out       bool
}{
	{"", "", true},
	{"", "a", false},
	{"", "abc", false},

	{"a", "", true},
	{"a", "a", true},
	{"a", "b", false},
	{"a", "abc", false},

	{"ab", "", true},
	{"ab", "a", true},
	{"ab", "b", false},
	{"ab", "ab", true},
	{"ab", "ba", false},
	{"ab", "abc", false},

	{"abcde", "", true},
	{"abcde", "a", true},
	{"abcde", "b", false},
	{"abcde", "ab", true},
	{"abcde", "ba", false},
	{"abcde", "abcde", true},
	{"abcde", "abcdef", false},
}

func TestLibHasPrefix(t *testing.T) {
	for _, test := range hasPrefixTests {
		actual := HasPrefix(test.s, test.prefix)
		libout := strings.HasPrefix(test.s, test.prefix)
		if actual != libout {
			t.Errorf("HasPrefix(%q, %q) = %t ; %t wanted", test.s, test.prefix, actual, libout)
		}
	}

}

func TestHasPrefix(t *testing.T) {
	for _, test := range hasPrefixTests {
		actual := HasPrefix(test.s, test.prefix)
		if actual != test.out {
			t.Errorf("HasPrefix(%q, %q) = %t ; %t wanted", test.s, test.prefix, actual, test.out)
		}
	}
}

func BenchmarkHasPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HasPrefix("abcde", "ab")
	}
}
