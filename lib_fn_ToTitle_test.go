package fStrings

import (
	"strings"
	"testing"
)

func ToTitle(s string) string {
	ret := ""
	for _, cp := range s {
		ret += ToUpper(string(cp))
	}

	return ret
}

var toTitleTests = []struct {
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

	{"aaa ", "AAA "},
	{" aaa", " AAA"},
	{" aaa ", " AAA "},

	{"aaa bbb ccc", "AAA BBB CCC"},
	{" aaa bbb ccc ", " AAA BBB CCC "},
}

func TestLibToTitle(t *testing.T) {
	for _, test := range toTitleTests {
		actual := ToTitle(test.s)
		libout := strings.ToTitle(test.s)
		if actual != libout {
			t.Errorf("ToTitle(%q) = %q ; %q wanted", test.s, actual, libout)
		}
	}
}

func TestToTitle(t *testing.T) {
	for _, test := range toTitleTests {
		actual := ToTitle(test.s)
		if actual != test.out {
			t.Errorf("ToTitle(%q) = %q ; %q wanted", test.s, actual, test.out)
		}
	}
}

func BenchmarkToTitle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToTitle(" aaa bbb ccc ")
	}
}
