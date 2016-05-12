package fStrings

import (
	"strings"
	"testing"
	"unicode"
)

func Title(s string) string {
	ret := ""

	prevIsSpace := true

	for _, cp := range s {
		if prevIsSpace && !unicode.IsSpace(cp) {
			ret += ToUpper(string(cp))
		} else {
			ret += string(cp)
		}

		prevIsSpace = unicode.IsSpace(cp)
	}

	return ret
}

var titleTests = []struct {
	s   string
	out string
}{
	{"", ""},
	{" ", " "},
	{"-", "-"},
	{"a", "A"},

	{"a b c", "A B C"},
	{"a B c", "A B C"},

	{"a ", "A "},
	{" a", " A"},
	{" a ", " A "},

	{"aaa ", "Aaa "},
	{" aaa", " Aaa"},
	{" aaa ", " Aaa "},

	{"aaa bbb ccc", "Aaa Bbb Ccc"},
	{" aaa bbb ccc ", " Aaa Bbb Ccc "},
}

func TestLibTitle(t *testing.T) {
	for _, test := range titleTests {
		actual := Title(test.s)
		libout := strings.Title(test.s)
		if actual != libout {
			t.Errorf("Title(%q) = %q ; %q wanted", test.s, actual, libout)
		}
	}
}

func TestTitle(t *testing.T) {
	for _, test := range titleTests {
		actual := Title(test.s)
		if actual != test.out {
			t.Errorf("Title(%q) = %q ; %q wanted", test.s, actual, test.out)
		}
	}
}

func BenchmarkTitle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Title(" aaa bbb ccc ")
	}
}
