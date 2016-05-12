package fStrings

import (
	"reflect"
	"strings"
	"testing"
)

func Split(s, sep string) []string {
	return SplitN(s, sep, -1)
}

var splitTests = []struct {
	s, sep string
	out    []string
}{
	{"", "", []string{}},
	{"a", "", []string{"a"}},
	{"", "a", []string{""}},
	{"a", "a", []string{"", ""}},
	{"a", "b", []string{"a"}},
	{"a-", "-", []string{"a", ""}},
	{"-a", "-", []string{"", "a"}},
	{"a-b", "-", []string{"a", "b"}},
	{"a-b-", "-", []string{"a", "b", ""}},
	{"-a-b", "-", []string{"", "a", "b"}},
	{"a-b-c-d-e", "-", []string{"a", "b", "c", "d", "e"}},
	{"a-:-b-:-c-:-d-:-e", "-:-", []string{"a", "b", "c", "d", "e"}},
}

func TestLibSplit(t *testing.T) {
	for _, test := range splitTests {
		actual := Split(test.s, test.sep)
		libout := strings.Split(test.s, test.sep)
		if !reflect.DeepEqual(actual, libout) {
			t.Errorf("Split(%q, %q) = %q ; %q", test.s, test.sep, actual, libout)
		}
	}
}

func TestSplit(t *testing.T) {
	for _, test := range splitTests {
		actual := Split(test.s, test.sep)
		if !reflect.DeepEqual(actual, test.out) {
			t.Errorf("Split(%q, %q) = %q ; %q", test.s, test.sep, actual, test.out)
		}
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a-:-b-:-c-:-d-:-e", "-:-")
	}
}
