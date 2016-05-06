package fStrings

import (
	"strings"
	"testing"
)

func TrimSuffix(s, suffix string) string {
	sLen := len(s)
	sufLen := len(suffix)

	switch {
	case suffix == "" || sLen < sufLen:
		return s
	case sLen >= sufLen:
		if s[sLen-sufLen:] == suffix {
			return s[:sLen-sufLen]
		}
		return s
	default:
		return ""
	}
}

var trimSuffixTests = []struct {
	s, suffix string
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
	{"ab", "a", "ab"},
	{"ab", "b", "a"},
	{"ab", "ab", ""},
	{"ab", "ba", "ab"},
	{"ab", "abc", "ab"},

	{"abcde", "", "abcde"},
	{"abcde", "e", "abcd"},
	{"abcde", "d", "abcde"},
	{"abcde", "de", "abc"},
	{"abcde", "ed", "abcde"},
	{"abcde", "abcde", ""},
	{"abcde", "_abcde", "abcde"},
}

func TestLibTrimSuffix(t *testing.T) {
	for _, test := range trimSuffixTests {
		actual := TrimSuffix(test.s, test.suffix)
		libOut := strings.TrimSuffix(test.s, test.suffix)
		if actual != libOut {
			t.Errorf("TrimSuffix(%q, %q) = %q ; %q wanted", test.s, test.suffix, actual, libOut)
		}
	}
}

func TestTrimSuffix(t *testing.T) {
	for _, test := range trimSuffixTests {
		actual := TrimSuffix(test.s, test.suffix)
		if actual != test.out {
			t.Errorf("TrimSuffix(%q, %q) = %q ; %q wanted", test.s, test.suffix, actual, test.out)
		}
	}
}

func BenchmarkTrimSuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrimSuffix("abcde", "de")
	}
}
