package fStrings

import (
	"reflect"
	"strings"
	"testing"
	"unicode"
)

func FieldsFunc(s string, f func(rune) bool) []string {
	r := make([]string, 0, 256)

	field := ""
	for _, cp := range s {
		switch {
		case f(cp) && field != "":
			r = append(r, field)
			field = ""
		case !f(cp):
			field += string(cp)
		}
	}

	if field != "" {
		r = append(r, field)
	}

	return r

}

var fieldsFuncTests = []struct {
	s   string
	f   func(rune) bool
	out []string
}{
	{"", func(r rune) bool { return unicode.IsSpace(r) }, []string{}},
	{" ", func(r rune) bool { return unicode.IsSpace(r) }, []string{}},
	{"-", func(r rune) bool { return unicode.IsSpace(r) }, []string{"-"}},

	{"a", func(r rune) bool { return unicode.IsSpace(r) }, []string{"a"}},
	{"a ", func(r rune) bool { return unicode.IsSpace(r) }, []string{"a"}},
	{" a", func(r rune) bool { return unicode.IsSpace(r) }, []string{"a"}},

	{"aaa", func(r rune) bool { return unicode.IsSpace(r) }, []string{"aaa"}},

	{" aaa bbb ccc ", func(r rune) bool { return unicode.IsSpace(r) }, []string{"aaa", "bbb", "ccc"}},
	{"aaa bbb ccc ", func(r rune) bool { return unicode.IsSpace(r) }, []string{"aaa", "bbb", "ccc"}},
	{" aaa bbb ccc", func(r rune) bool { return unicode.IsSpace(r) }, []string{"aaa", "bbb", "ccc"}},
	{"aaa bbb ccc", func(r rune) bool { return unicode.IsSpace(r) }, []string{"aaa", "bbb", "ccc"}},
	{" \taaa \n\t bbb \v ccc \n", func(r rune) bool { return unicode.IsSpace(r) }, []string{"aaa", "bbb", "ccc"}},
}

func TestLibFieldsFunc(t *testing.T) {
	for _, test := range fieldsFuncTests {
		actual := FieldsFunc(test.s, test.f)
		libout := strings.FieldsFunc(test.s, test.f)
		if !reflect.DeepEqual(actual, libout) {
			t.Errorf("FieldsFunc(%q, %T) = %q ; %q wanted", test.s, test.f, actual, libout)
		}
	}
}

func TestFieldsFunc(t *testing.T) {
	for _, test := range fieldsFuncTests {
		actual := FieldsFunc(test.s, test.f)
		if !reflect.DeepEqual(actual, test.out) {
			t.Errorf("FieldsFunc(%q, %T) = %q ; %q wanted", test.s, test.f, actual, test.out)
		}
	}
}

func BenchmarkFieldsFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FieldsFunc(" aaa bbb ccc ", unicode.IsSpace)
	}
}
