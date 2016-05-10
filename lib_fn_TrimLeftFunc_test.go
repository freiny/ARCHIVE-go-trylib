package fStrings

import (
	"strings"
	"testing"
)

func TrimLeftFunc(s string, f func(rune) bool) string {
	start := 0
	for i := 0; i < len(s); i++ {
		start = i
		if !f(rune(s[i])) {
			return s[start:]
		}
	}
	return ""
}

var trimLeftFuncTests = []struct {
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

	{"abcab", func(r rune) bool { return r == rune('a') }, "bcab"},
	{"abcab", func(r rune) bool { return r == rune('c') }, "abcab"},
	{"abcab", func(r rune) bool { return r == rune('b') }, "abcab"},
	{"bacba", func(r rune) bool { return r == rune('a') }, "bacba"},
	{"bacba", func(r rune) bool { return r == rune('c') }, "bacba"},
	{"bacba", func(r rune) bool { return r == rune('b') }, "acba"},
	{"bacba", func(r rune) bool { return r == rune('D') }, "bacba"},

	{"abcab", func(r rune) bool { return r == rune('a') || r == rune('b') }, "cab"},
	{"abcab", func(r rune) bool { return r == rune('b') || r == rune('a') }, "cab"},
	{"abcab", func(r rune) bool { return r == rune('a') || r == rune('D') }, "bcab"},
	{"abcab", func(r rune) bool { return r == rune('D') || r == rune('b') }, "abcab"},
	{"bacba", func(r rune) bool { return r == rune('a') || r == rune('b') }, "cba"},
	{"bacba", func(r rune) bool { return r == rune('b') || r == rune('a') }, "cba"},
	{"bacba", func(r rune) bool { return r == rune('a') || r == rune('D') }, "bacba"},
	{"bacba", func(r rune) bool { return r == rune('D') || r == rune('a') }, "bacba"},
	{"ababbaab", func(r rune) bool { return r == rune('a') || r == rune('b') }, ""},
	{"ababbaab", func(r rune) bool { return r == rune('b') || r == rune('a') }, ""},
	{"ababbaab", func(r rune) bool { return r == rune('a') || r == rune('D') || r == rune('b') }, ""},
	{"ababbaab", func(r rune) bool { return r == rune('b') || r == rune('D') || r == rune('a') }, ""},

	{"abcab", func(r rune) bool { return r == rune('a') || r == rune('b') || r == rune('c') }, ""},
}

func TestLibTrimLeftFunc(t *testing.T) {
	for _, test := range trimLeftFuncTests {
		actual := TrimLeftFunc(test.s, test.f)
		libout := strings.TrimLeftFunc(test.s, test.f)
		// if true {
		if actual != libout {
			t.Errorf("TrimLeftFunc(%q, %T) = %q, %q wanted", test.s, test.f, actual, libout)
		}
	}
}

func TestTrimLeftFunc(t *testing.T) {
	for _, test := range trimLeftFuncTests {
		actual := TrimLeftFunc(test.s, test.f)
		if actual != test.out {
			t.Errorf("TrimLeftFunc(%q, %T) = %q, %q wanted", test.s, test.f, actual, test.out)
		}
	}
}

func BenchmarkTrimLeftFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrimLeftFunc("ababbaab", func(r rune) bool { return r == rune('b') || r == rune('D') || r == rune('a') })
	}

}
