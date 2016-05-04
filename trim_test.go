package fStrings

import (
	"strings"
	"testing"
)

func Trim(s, cutset string) string {
	var a, b int
	m := map[byte]bool{}
	for i := 0; i < len(cutset); i++ {
		m[cutset[i]] = true
	}
	for i := 0; i < len(s); i++ {
		if !m[s[i]] {
			a = i
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if !m[s[i]] {
			b = i + 1
			break
		}
	}

	if a == b || a < b {
		return s[a:b]
	}
	return ""

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
		libOut := strings.Trim(test.s, test.cutset)
		if actual != libOut {
			t.Errorf("Trim(%q, %q) = %q ; %q wanted", test.s, test.cutset, actual, libOut)
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
