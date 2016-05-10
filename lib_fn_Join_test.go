package fStrings

import (
	"strings"
	"testing"
)

func Join(a []string, sep string) string {
	if len(a) == 0 {
		return ""
	}
	if len(a) == 1 {
		return a[0]
	}

	var s string
	for i := 0; i < len(a)-1; i++ {
		s += a[i] + sep
	}
	return s + a[len(a)-1]
}

var joinTests = []struct {
	a   []string
	sep string
	out string
}{
	{[]string{}, "", ""},
	{[]string{}, "@", ""},
	{[]string{""}, "", ""},
	{[]string{""}, "@", ""},
	{[]string{"a"}, "", "a"},
	{[]string{"a"}, "@", "a"},
	{[]string{"abc"}, "", "abc"},
	{[]string{"abc"}, "@", "abc"},
	{[]string{"a", "b"}, "@", "a@b"},
	{[]string{"a", "bc"}, "@", "a@bc"},
	{[]string{"ab", "c"}, "@", "ab@c"},
	{[]string{"a", "b", "c"}, "@", "a@b@c"},
	{[]string{"abc", "de", "f"}, "@", "abc@de@f"},
}

func TestLibJoin(t *testing.T) {
	for _, test := range joinTests {
		actual := Join(test.a, test.sep)
		libout := strings.Join(test.a, test.sep)
		if actual != libout {
			t.Errorf("Join(%q, %q) = %q ; %q wanted", test.a, test.sep, actual, libout)
		}
	}
}

func TestJoin(t *testing.T) {
	for _, test := range joinTests {
		actual := Join(test.a, test.sep)
		if actual != test.out {
			t.Errorf("Join(%q, %q) = %q ; %q wanted", test.a, test.sep, actual, test.out)
		}
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join([]string{"ab", "cd", "ef", "gh"}, "-*-")
	}
}
