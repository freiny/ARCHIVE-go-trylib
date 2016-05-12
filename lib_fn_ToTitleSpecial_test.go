package fStrings

import (
	"strings"
	"testing"
	"unicode"
)

func ToTitleSpecial(_case unicode.SpecialCase, s string) string {
	return Map(func(r rune) rune { return _case.ToTitle(r) }, s)
}

var toTitleSpecialTests = []struct {
	_case unicode.SpecialCase
	s     string
	out   string
}{
	{unicode.TurkishCase, "", ""},
	{unicode.TurkishCase, " ", " "},
	{unicode.TurkishCase, "-", "-"},

	{unicode.TurkishCase, "i", "İ"},
	{unicode.TurkishCase, "iii", "İİİ"},
	{unicode.TurkishCase, "iii iii iii", "İİİ İİİ İİİ"},
}

func TestLibToTitleSpecial(t *testing.T) {
	for _, test := range toTitleSpecialTests {
		actual := ToTitleSpecial(test._case, test.s)
		libout := strings.ToTitleSpecial(test._case, test.s)
		if actual != libout {
			t.Errorf("ToTitleSpecial(%q) = %q ; %q wanted", test.s, actual, libout)
		}
	}
}

func TestToTitleSpecial(t *testing.T) {
	for _, test := range toTitleSpecialTests {
		actual := ToTitleSpecial(test._case, test.s)
		if actual != test.out {
			t.Errorf("ToTitleSpecial(%q) = %q ; %q wanted", test.s, actual, test.out)
		}
	}
}

func BenchmarkToTitleSpecial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToTitleSpecial(unicode.TurkishCase, " iii iii iii ")
	}
}
