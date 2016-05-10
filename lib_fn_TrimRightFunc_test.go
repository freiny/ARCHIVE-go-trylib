package fStrings

import (
	"strings"
	"testing"
)

func TrimRightFunc(s string, f func(rune) bool) string {
	end := len(s) - 1
	for i := len(s) - 1; i >= 0; i-- {
		end = i
		if !f(rune(s[i])) {
			return s[:end+1]
		}
	}

	return ""
}

var trimRightFuncTests = []struct {
	s   string
	f   func(rune) bool
	out string
}{

	{"", func(r rune) bool { return r == rune(0) }, ""},
	{"", func(r rune) bool { return r == rune('a') }, ""},
	{"a", func(r rune) bool { return r == rune(0) }, "a"},
	{"a", func(r rune) bool { return r == rune('a') }, ""},
	{"a", func(r rune) bool { return r == rune('b') }, "a"},

	{"aaaaa", func(r rune) bool { return r == rune('a') }, ""},

	{"abcab", func(r rune) bool { return r == rune('a') }, "abcab"},
	{"abcab", func(r rune) bool { return r == rune('c') }, "abcab"},
	{"abcab", func(r rune) bool { return r == rune('b') }, "abca"},
	{"bacba", func(r rune) bool { return r == rune('a') }, "bacb"},
	{"bacba", func(r rune) bool { return r == rune('c') }, "bacba"},
	{"bacba", func(r rune) bool { return r == rune('b') }, "bacba"},
	{"bacba", func(r rune) bool { return r == rune('D') }, "bacba"},

	{"abcab", func(r rune) bool { return r == rune('a') || r == rune('b') }, "abc"},
	{"abcab", func(r rune) bool { return r == rune('b') || r == rune('a') }, "abc"},
	{"abcab", func(r rune) bool { return r == rune('a') || r == rune('D') }, "abcab"},
	{"abcab", func(r rune) bool { return r == rune('D') || r == rune('b') }, "abca"},
	{"bacba", func(r rune) bool { return r == rune('a') || r == rune('b') }, "bac"},
	{"bacba", func(r rune) bool { return r == rune('b') || r == rune('a') }, "bac"},
	{"bacba", func(r rune) bool { return r == rune('a') || r == rune('D') }, "bacb"},
	{"bacba", func(r rune) bool { return r == rune('D') || r == rune('a') }, "bacb"},
	{"ababbaab", func(r rune) bool { return r == rune('a') || r == rune('b') }, ""},
	{"ababbaab", func(r rune) bool { return r == rune('b') || r == rune('a') }, ""},
	{"ababbaab", func(r rune) bool { return r == rune('a') || r == rune('D') || r == rune('b') }, ""},
	{"ababbaab", func(r rune) bool { return r == rune('b') || r == rune('D') || r == rune('a') }, ""},

	{"abcab", func(r rune) bool { return r == rune('a') || r == rune('b') || r == rune('c') }, ""},
}

func TestLibTrimRightFunc(t *testing.T) {
	for _, test := range trimRightFuncTests {
		actual := TrimRightFunc(test.s, test.f)
		libout := strings.TrimRightFunc(test.s, test.f)
		if actual != libout {
			t.Errorf("TrimRightFunc(%q, %T) = %q, %q wanted", test.s, test.f, actual, libout)
		}
	}
}

func TestTrimRightFunc(t *testing.T) {
	for _, test := range trimRightFuncTests {
		actual := TrimRightFunc(test.s, test.f)
		if actual != test.out {
			t.Errorf("TrimRightFunc(%q, %T) = %q, %q wanted", test.s, test.f, actual, test.out)
		}
	}
}

func BenchmarkTrimRightFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrimRightFunc("ababbaab", func(r rune) bool { return r == rune('b') || r == rune('D') || r == rune('a') })
	}

}
