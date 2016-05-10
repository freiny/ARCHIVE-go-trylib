package fStrings

import (
	"strings"
	"testing"
)

func TrimRight(s, cutset string) string {
	return TrimRightFunc(s, func(r rune) bool { return ContainsAny(string(r), cutset) })
}

var trimRightTests = []struct {
	s, cutset string
	out       string
}{
	{"", "", ""},
	{"", "a", ""},
	{"a", "", "a"},
	{"a", "a", ""},
	{"a", "b", "a"},
	{"abcab", "a", "abcab"},
	{"abcab", "c", "abcab"},
	{"abcab", "b", "abca"},
	{"bacba", "a", "bacb"},
	{"bacba", "c", "bacba"},
	{"bacba", "b", "bacba"},
	{"bacba", "D", "bacba"},
	{"abcab", "ab", "abc"},
	{"abcab", "ba", "abc"},
	{"abcab", "aD", "abcab"},
	{"abcab", "Db", "abca"},
	{"bacba", "ab", "bac"},
	{"bacba", "ba", "bac"},
	{"bacba", "aD", "bacb"},
	{"bacba", "Da", "bacb"},
	{"ababbaab", "ab", ""},
	{"ababbaab", "aDb", ""},
	{"ababbaab", "ba", ""},
	{"ababbaab", "bDa", ""},
}

func TestLibTrimRight(t *testing.T) {
	for _, test := range trimRightTests {
		actual := TrimRight(test.s, test.cutset)
		libout := strings.TrimRight(test.s, test.cutset)
		if actual != libout {
			t.Errorf("TrimRight(%q, %q) = %q ; %q wanted", test.s, test.cutset, actual, libout)
		}
	}
}

func TestTrimRight(t *testing.T) {
	for _, test := range trimRightTests {
		actual := TrimRight(test.s, test.cutset)
		if actual != test.out {
			t.Errorf("TrimRight(%q, %q) = %q ; %q wanted", test.s, test.cutset, actual, test.out)
		}
	}
}

func BenchmarkTrimRight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrimRight(" _ : h_ :_ w__ :: :_ ", "_ :")
	}
}
