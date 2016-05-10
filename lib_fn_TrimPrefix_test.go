package fStrings

import (
	"strings"
	"testing"
)

func TrimPrefix(s, prefix string) string {
	if HasPrefix(s, prefix) {
		return s[len(prefix):]
	}
	return s
}

var trimPrefixTests = []struct {
	s, prefix string
	out       string
}{
	{"", "", ""},
	{"", "a", ""},
	{"", "abc", ""},

	{"a", "", "a"},
	{"a", "a", ""},
	{"a", "b", "a"},
	{"a", "abc", "a"},

	{"ab", "", "ab"},
	{"ab", "a", "b"},
	{"ab", "b", "ab"},
	{"ab", "ab", ""},
	{"ab", "ba", "ab"},
	{"ab", "abc", "ab"},

	{"abcde", "", "abcde"},
	{"abcde", "a", "bcde"},
	{"abcde", "b", "abcde"},
	{"abcde", "ab", "cde"},
	{"abcde", "ba", "abcde"},
	{"abcde", "abcde", ""},
	{"abcde", "abcdef", "abcde"},
}

func TestLibTrimPrefix(t *testing.T) {
	for _, test := range trimPrefixTests {
		actual := TrimPrefix(test.s, test.prefix)
		libout := strings.TrimPrefix(test.s, test.prefix)
		if actual != libout {
			t.Errorf("TrimPrefix(%q, %q) = %q ; %q wanted", test.s, test.prefix, actual, libout)
		}
	}
}

func TestTrimPrefix(t *testing.T) {
	for _, test := range trimPrefixTests {
		actual := TrimPrefix(test.s, test.prefix)
		if actual != test.out {
			t.Errorf("TrimPrefix(%q, %q) = %q ; %q wanted", test.s, test.prefix, actual, test.out)
		}
	}
}

func BenchmarkTrimPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrimPrefix("abcde", "ab")
	}
}
