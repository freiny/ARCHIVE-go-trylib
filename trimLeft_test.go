package fStrings

import (
	"strings"
	"testing"
)

func TrimLeft(s, cutset string) string {

	inCutset := map[byte]bool{}
	for i := 0; i < len(cutset); i++ {
		inCutset[cutset[i]] = true
	}

	sLen := len(s)
	start := 0
	for i := 0; i < sLen; i++ {
		start = i
		if !inCutset[s[i]] {
			break
		}
	}
	if start == sLen-1 && inCutset[s[sLen-1]] {
		start++
	}

	if start <= sLen {
		return s[start:]
	}
	return ""

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
		libOut := strings.TrimLeft(test.s, test.cutset)
		if actual != libOut {
			t.Errorf("TrimLeft(%q, %q) = %q ; %q wanted", test.s, test.cutset, actual, libOut)
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
