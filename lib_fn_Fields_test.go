package fStrings

import (
	"reflect"
	"strings"
	"testing"
	"unicode"
)

func Fields(s string) []string {
	return FieldsFunc(s, unicode.IsSpace)
}

var fieldsTests = []struct {
	s   string
	out []string
}{
	{"", []string{}},
	{" ", []string{}},
	{"-", []string{"-"}},

	{"a", []string{"a"}},
	{"a ", []string{"a"}},
	{" a", []string{"a"}},

	{"aaa", []string{"aaa"}},

	{" aaa bbb ccc ", []string{"aaa", "bbb", "ccc"}},
	{"aaa bbb ccc ", []string{"aaa", "bbb", "ccc"}},
	{" aaa bbb ccc", []string{"aaa", "bbb", "ccc"}},
	{"aaa bbb ccc", []string{"aaa", "bbb", "ccc"}},
	{" \taaa \n\t bbb \v ccc \n", []string{"aaa", "bbb", "ccc"}},
}

func TestLibFields(t *testing.T) {
	for _, test := range fieldsTests {
		actual := Fields(test.s)
		libout := strings.Fields(test.s)
		if !reflect.DeepEqual(actual, libout) {
			t.Errorf("Fields(%q) = %q ; %q wanted", test.s, actual, libout)
		}
	}
}

func TestFields(t *testing.T) {
	for _, test := range fieldsTests {
		actual := Fields(test.s)
		if !reflect.DeepEqual(actual, test.out) {
			t.Errorf("Fields(%q) = %q ; %q wanted", test.s, actual, test.out)
		}
	}
}

func BenchmarkFields(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fields(" aaa bbb ccc ")
	}
}
