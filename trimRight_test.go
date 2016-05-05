package fStrings

import (
	"strings"
	"testing"
)

func TrimRight(s, cutset string) string {

	inCutset := map[byte]bool{}
	for i := 0; i < len(cutset); i++ {
		inCutset[cutset[i]] = true
	}

	sLen := len(s)
	end := 0
	for i := sLen; i > 0; i-- {
		end = i
		if !inCutset[s[i-1]] {
			break
		}
	}
	if end == 1 && inCutset[s[0]] {
		return ""
	}

	if end > 0 {
		return s[:end]
	}
	return ""

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
		libOut := strings.TrimRight(test.s, test.cutset)
		if actual != libOut {
			t.Errorf("TrimRight(%q, %q) = %q ; %q wanted", test.s, test.cutset, actual, libOut)
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
