package fStrings

import (
	"strings"
	"testing"
)

func HasSuffix(s, suffix string) bool {

	if suffix == "" {
		return true
	}
	if len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix {
		return true
	}
	return false

}

var hasSuffixTests = []struct {
	s, suffix string
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
	{"ab", "a", false},
	{"ab", "b", true},
	{"ab", "ab", true},
	{"ab", "ba", false},
	{"ab", "abc", false},

	{"abcde", "", true},
	{"abcde", "e", true},
	{"abcde", "d", false},
	{"abcde", "de", true},
	{"abcde", "ed", false},
	{"abcde", "abcde", true},
	{"abcde", "_abcde", false},
}

func TestLibHasSuffix(t *testing.T) {
	for _, test := range hasSuffixTests {
		actual := HasSuffix(test.s, test.suffix)
		libout := strings.HasSuffix(test.s, test.suffix)
		if actual != libout {
			t.Errorf("HasSuffix(%q, %q) = %t ; %t wanted", test.s, test.suffix, actual, libout)
		}
	}
}

func TestHasSuffix(t *testing.T) {
	for _, test := range hasSuffixTests {
		actual := HasSuffix(test.s, test.suffix)
		if actual != test.out {
			t.Errorf("HasSuffix(%q, %q) = %t ; %t wanted", test.s, test.suffix, actual, test.out)
		}
	}
}

func BenchmarkHasSuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HasSuffix("abcde", "de")
	}
}
