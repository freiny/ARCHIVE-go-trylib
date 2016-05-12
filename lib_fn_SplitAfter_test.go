package fStrings

import (
	"reflect"
	"strings"
	"testing"
)

func SplitAfter(s, sep string) []string {
	switch {
	case s == "" && sep == "":
		return make([]string, 0)
	case sep == "" || Index(s, sep) == -1:
		return []string{s}
	}

	a := SplitN(s, sep, -1)

	ret := make([]string, 0, 256)

	for i := 0; i < len(a)-1; i++ {
		ret = append(ret, a[i]+sep)
	}
	if len(a) > 1 {
		ret = append(ret, a[len(a)-1])
	}
	return ret
}

var splitAfterTests = []struct {
	s, sep string
	out    []string
}{
	{"", "", []string{}},
	{"a", "", []string{"a"}},
	{"", "a", []string{""}},
	{"a", "a", []string{"a", ""}},
	{"a", "b", []string{"a"}},
	{"a-", "-", []string{"a-", ""}},
	{"-a", "-", []string{"-", "a"}},
	{"a-b", "-", []string{"a-", "b"}},
	{"a-b-", "-", []string{"a-", "b-", ""}},
	{"-a-b", "-", []string{"-", "a-", "b"}},
	{"a-b-c-d-e", "-", []string{"a-", "b-", "c-", "d-", "e"}},
	{"a-:-b-:-c-:-d-:-e", "-:-", []string{"a-:-", "b-:-", "c-:-", "d-:-", "e"}},
	{"a-:-b-:-c-:-d-:-e-:-", "-:-", []string{"a-:-", "b-:-", "c-:-", "d-:-", "e-:-", ""}},
}

func TestLibSplitAfter(t *testing.T) {
	for _, test := range splitAfterTests {
		actual := SplitAfter(test.s, test.sep)
		libout := strings.SplitAfter(test.s, test.sep)
		if !reflect.DeepEqual(actual, libout) {
			t.Errorf("SplitAfter(%q, %q) = %q ; %q", test.s, test.sep, actual, libout)
		}
	}
}

func TestSplitAfter(t *testing.T) {
	for _, test := range splitAfterTests {
		actual := SplitAfter(test.s, test.sep)
		if !reflect.DeepEqual(actual, test.out) {
			t.Errorf("SplitAfter(%q, %q) = %q ; %q", test.s, test.sep, actual, test.out)
		}
	}
}

func BenchmarkSplitAfter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SplitAfter("a-:-b-:-c-:-d-:-e", "-:-")
	}
}
