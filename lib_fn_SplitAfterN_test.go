package fStrings

import (
	"reflect"
	"strings"
	"testing"
)

func SplitAfterN(s, sep string, n int) []string {
	switch {
	case s == "" && sep == "":
		return make([]string, 0)
	case sep == "" || Index(s, sep) == -1:
		return []string{s}
	}

	a := SplitN(s, sep, n)

	ret := make([]string, 0, 256)

	for i := 0; i < len(a)-1; i++ {
		ret = append(ret, a[i]+sep)
	}
	if len(a) > 1 {
		ret = append(ret, a[len(a)-1])
	}
	return ret
}

var splitAfterNTests = []struct {
	s, sep string
	n      int
	out    []string
}{
	{"", "", 2, []string{}},
	{"a", "", 2, []string{"a"}},
	{"", "a", 2, []string{""}},
	{"a", "a", 2, []string{"a", ""}},
	{"a", "b", 2, []string{"a"}},
	{"a-", "-", 2, []string{"a-", ""}},
	{"-a", "-", 2, []string{"-", "a"}},
	{"a-b", "-", 2, []string{"a-", "b"}},
	{"a-b-", "-", 2, []string{"a-", "b-"}},
	{"-a-b", "-", 2, []string{"-", "a-b"}},
	{"a-b-c-d-e", "-", 2, []string{"a-", "b-c-d-e"}},
	{"a-:-b-:-c-:-d-:-e", "-:-", 2, []string{"a-:-", "b-:-c-:-d-:-e"}},
	{"a-:-b-:-c-:-d-:-e-:-", "-:-", 2, []string{"a-:-", "b-:-c-:-d-:-e-:-"}},
}

func TestLibSplitAfterN(t *testing.T) {
	for _, test := range splitAfterNTests {
		actual := SplitAfterN(test.s, test.sep, test.n)
		libout := strings.SplitAfterN(test.s, test.sep, test.n)
		if !reflect.DeepEqual(actual, libout) {
			t.Errorf("SplitAfterN(%q, %q, %d) = %q ; %q", test.s, test.sep, test.n, actual, libout)
		}
	}
}

func TestSplitAfterN(t *testing.T) {
	for _, test := range splitAfterNTests {
		actual := SplitAfterN(test.s, test.sep, test.n)
		if !reflect.DeepEqual(actual, test.out) {
			t.Errorf("SplitAfterN(%q, %q, %d) = %q ; %q", test.s, test.sep, test.n, actual, test.out)
		}
	}
}

func BenchmarkSplitAfterN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SplitAfterN("a-:-b-:-c-:-d-:-e", "-:-", 3)
	}
}
