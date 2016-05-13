package fStrings

import (
	"reflect"
	"strings"
	"testing"
)

func SplitN(s, sep string, n int) []string {
	switch {
	case n == 0:
		var ret []string
		return ret
	case s == "" && sep == "":
		return make([]string, 0)
	case n == 1:
		return []string{s}
	case s == sep:
		return []string{"", ""}
	case sep == "" || Index(s, sep) == -1:
		return []string{s}

	case n != 0:
		ret := make([]string, 0, 256)
		i := 1
		start := 0

		for index := strings.Index(s[start:], sep); index != -1; index = strings.Index(s[start:], sep) {
			if i == n {
				break
			}
			i++
			ret = append(ret, s[start:start+index])
			start += index + len(sep)
		}
		ret = append(ret, s[start:])
		return ret
	default:
		return []string{s}
	}

}

var splitNTests = []struct {
	s, sep string
	n      int
	out    []string
}{
	{"", "", -1, []string{}},
	{"", "", 0, none},
	{"", "", 1, []string{}},
	{"", "", 2, []string{}},
	{"", "", 3, []string{}},

	{"a", "", -1, []string{"a"}},
	{"a", "", 0, none},
	{"a", "", 1, []string{"a"}},
	{"a", "", 2, []string{"a"}},
	{"a", "", 3, []string{"a"}},

	{"", "a", -1, []string{""}},
	{"", "a", 0, none},
	{"", "a", 1, []string{""}},
	{"", "a", 2, []string{""}},
	{"", "a", 3, []string{""}},

	{"a", "a", -1, []string{"", ""}},
	{"a", "a", 0, none},
	{"a", "a", 1, []string{"a"}},
	{"-", "-", 1, []string{"-"}},
	{"a", "a", 2, []string{"", ""}},
	{"a", "a", 3, []string{"", ""}},

	{"a", "b", -1, []string{"a"}},
	{"a", "b", 0, none},
	{"a", "b", 1, []string{"a"}},
	{"a", "b", 2, []string{"a"}},
	{"a", "b", 3, []string{"a"}},

	{"a-", "-", -1, []string{"a", ""}},
	{"a-", "-", 0, none},
	{"a-", "-", 1, []string{"a-"}},
	{"a-", "-", 2, []string{"a", ""}},
	{"a-", "-", 3, []string{"a", ""}},

	{"-a", "-", -1, []string{"", "a"}},
	{"-a", "-", 0, none},
	{"-a", "-", 1, []string{"-a"}},
	{"-a", "-", 2, []string{"", "a"}},
	{"-a", "-", 3, []string{"", "a"}},

	{"a-b", "-", -1, []string{"a", "b"}},
	{"a-b", "-", 0, none},
	{"a-b", "-", 1, []string{"a-b"}},
	{"a-b", "-", 2, []string{"a", "b"}},
	{"a-b", "-", 3, []string{"a", "b"}},

	{"a-b", "a-b", 1, []string{"a-b"}},
	{"a-b", "a-b", 3, []string{"", ""}},

	{"a-b-", "-", -1, []string{"a", "b", ""}},
	{"a-b-", "-", 0, none},
	{"a-b-", "-", 1, []string{"a-b-"}},
	{"a-b-", "-", 2, []string{"a", "b-"}},
	{"a-b-", "-", 3, []string{"a", "b", ""}},

	{"-a-b", "-", -1, []string{"", "a", "b"}},
	{"-a-b", "-", 0, none},
	{"-a-b", "-", 1, []string{"-a-b"}},
	{"-a-b", "-", 2, []string{"", "a-b"}},
	{"-a-b", "-", 3, []string{"", "a", "b"}},

	{"a-b-c-d-e", "-", -1, []string{"a", "b", "c", "d", "e"}},
	{"a-b-c-d-e", "-", 0, none},
	{"a-b-c-d-e", "-", 1, []string{"a-b-c-d-e"}},
	{"a-b-c-d-e", "-", 2, []string{"a", "b-c-d-e"}},
	{"a-b-c-d-e", "-", 3, []string{"a", "b", "c-d-e"}},

	{"a-:-b-:-c-:-d-:-e", "-:-", -2, []string{"a", "b", "c", "d", "e"}},
	{"a-:-b-:-c-:-d-:-e", "-:-", -1, []string{"a", "b", "c", "d", "e"}},
	{"a-:-b-:-c-:-d-:-e", "-:-", 0, none},
	{"a-:-b-:-c-:-d-:-e", "-:-", 1, []string{"a-:-b-:-c-:-d-:-e"}},
	{"a-:-b-:-c-:-d-:-e", "-:-", 2, []string{"a", "b-:-c-:-d-:-e"}},
	{"a-:-b-:-c-:-d-:-e", "-:-", 3, []string{"a", "b", "c-:-d-:-e"}},
}

func TestLibSplitN(t *testing.T) {
	for _, test := range splitNTests {
		actual := SplitN(test.s, test.sep, test.n)
		libout := strings.SplitN(test.s, test.sep, test.n)
		if !reflect.DeepEqual(actual, libout) {
			t.Errorf("SplitN(%q, %q, %d) = %q; %q wanted", test.s, test.sep, test.n, actual, libout)
		}
	}
}

func TestSplitN(t *testing.T) {
	for _, test := range splitNTests {
		actual := SplitN(test.s, test.sep, test.n)
		if !reflect.DeepEqual(actual, test.out) {
			t.Errorf("SplitN(%q, %q, %d) = %q; %q wanted", test.s, test.sep, test.n, actual, test.out)
		}
	}
}

func BenchmarkSplitN(b *testing.B) {
}
