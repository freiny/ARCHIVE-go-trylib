package fStrings

import (
	"strings"
	"testing"
)

func TrimLeft(s, cutset string) string {
	return TrimLeftFunc(s, func(r rune) bool { return ContainsAny(string(r), cutset) })
}

var trimLeftTests = []struct {
	s, cutset string
	out       string
}{
	{"", "", ""},
	{"", "a", ""},
	{"a", "", "a"},
	{"a", "a", ""},
	{"a", "b", "a"},
	{"abcab", "a", "bcab"},
	{"abcab", "c", "abcab"},
	{"abcab", "b", "abcab"},
	{"bacba", "a", "bacba"},
	{"bacba", "c", "bacba"},
	{"bacba", "b", "acba"},
	{"bacba", "D", "bacba"},
	{"abcab", "ab", "cab"},
	{"abcab", "ba", "cab"},
	{"abcab", "aD", "bcab"},
	{"abcab", "Db", "abcab"},
	{"bacba", "ab", "cba"},
	{"bacba", "ba", "cba"},
	{"bacba", "aD", "bacba"},
	{"bacba", "Da", "bacba"},
	{"ababbaab", "ab", ""},
	{"ababbaab", "aDb", ""},
	{"ababbaab", "ba", ""},
	{"ababbaab", "bDa", ""},
}

func TestLibTrimLeft(t *testing.T) {
	for _, test := range trimLeftTests {
		actual := TrimLeft(test.s, test.cutset)
		libout := strings.TrimLeft(test.s, test.cutset)
		if actual != libout {
			t.Errorf("TrimLeft(%q, %q) = %q ; %q wanted", test.s, test.cutset, actual, libout)
		}
	}
}

func TestTrimLeft(t *testing.T) {
	for _, test := range trimLeftTests {
		actual := TrimLeft(test.s, test.cutset)
		if actual != test.out {
			t.Errorf("TrimLeft(%q, %q) = %q ; %q wanted", test.s, test.cutset, actual, test.out)
		}
	}
}

func BenchmarkTrimLeft(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrimLeft(" _ : h_ :_ w__ :: :_ ", "_ :")
	}
}
