package fStrings

import (
	"strings"
	"testing"
)

func Trim(s, cutset string) string {
	ret := TrimLeft(s, cutset)
	return TrimRight(ret, cutset)
}

var trimTests = []struct {
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
	{"abcab", "b", "abca"},
	{"bacba", "a", "bacb"},
	{"bacba", "c", "bacba"},
	{"bacba", "b", "acba"},
	{"bacba", "D", "bacba"},
	{"abcab", "ab", "c"},
	{"abcab", "ba", "c"},
	{"abcab", "aD", "bcab"},
	{"abcab", "Db", "abca"},
	{"bacba", "ab", "c"},
	{"bacba", "ba", "c"},
	{"bacba", "aD", "bacb"},
	{"bacba", "Da", "bacb"},
	{"ababbaab", "ab", ""},
	{"ababbaab", "aDb", ""},
	{"ababbaab", "ba", ""},
	{"ababbaab", "bDa", ""},
}

func TestLibTrim(t *testing.T) {
	for _, test := range trimTests {
		actual := Trim(test.s, test.cutset)
		libout := strings.Trim(test.s, test.cutset)
		if actual != libout {
			t.Errorf("Trim(%q, %q) = %q ; %q wanted", test.s, test.cutset, actual, libout)
		}
	}
}

func TestTrim(t *testing.T) {
	for _, test := range trimTests {
		actual := Trim(test.s, test.cutset)
		if actual != test.out {
			t.Errorf("Trim(%q, %q) = %q ; %q wanted", test.s, test.cutset, actual, test.out)
		}
	}
}

func BenchmarkTrim(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Trim(" _ : h_ :_ w__ :: :_ ", "_ :")
	}
}
