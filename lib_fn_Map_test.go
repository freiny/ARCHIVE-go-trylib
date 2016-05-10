package fStrings

import (
	"strings"
	"testing"
)

func Map(mapping func(rune) rune, s string) string {
	ret := ""
	for i := 0; i < len(s); i++ {
		m := mapping(rune(s[i]))
		if m >= 0 {
			ret += string(m)
		}
	}
	return ret
}

var mapTests = []struct {
	mapping func(rune) rune
	s       string
	out     string
}{
	{func(r rune) rune {
		return r + 1
	}, "", ""},

	{func(r rune) rune {
		return r + 1
	}, "a", "b"},

	{func(r rune) rune {
		return r + 1
	}, "abcde", "bcdef"},

	{func(r rune) rune {
		return r - 32
	}, "abcDEFghi", "ABC$%&GHI"},

	{func(r rune) rune {
		return r - 1000
	}, "aabbcc", ""},
}

func TestLibMap(t *testing.T) {
	for _, test := range mapTests {
		actual := Map(test.mapping, test.s)
		libout := strings.Map(test.mapping, test.s)
		if actual != libout {
			t.Errorf("Map(%T, %q) = %q ; %q wanted", test.mapping, test.s, actual, libout)
		}
	}
}

func TestMap(t *testing.T) {
	for _, test := range mapTests {
		actual := Map(test.mapping, test.s)
		if actual != test.out {
			t.Errorf("Map(%T, %q) = %q ; %q wanted", test.mapping, test.s, actual, test.out)
		}
	}
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Map(func(r rune) rune { return r - 32 }, "abcdefghijklmnopqrstuvwxyz")
	}
}
