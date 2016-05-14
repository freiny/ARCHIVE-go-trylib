package fStrings

import (
	"strings"
	"testing"
)

func Replace(s, old, new string, n int) string {
	ret := ""

	switch {
	case n == 0:
		return s
	case new == "":
		ret = s
	case old == "" && s == "":
		ret = new
	case n == 1 && old == "":
		ret += new + s
	case old == "":
		ret += new
		for i := 0; i < n-1 || (n == -1 && i < len(s)); i++ {
			ret += string(s[i]) + new
		}
		if n > 1 {
			ret += string(s[n-1:])
		}
	default:
		current := s[:]
		start := 0
		index := 0

		for i := 0; (i < n && n != -1) || (n == -1); i++ {
			current = current[start:]

			index = Index(current, old)
			if index == -1 {
				break
			}

			ret += current[0:index] + new
			start = index + len(old)
		}

		if n > 0 {
			index := Index(current, old)
			if index != -1 {
				ret += current[index+len(old):]
			}
		} else {
			ret += current

		}
	}

	return ret
}

var replaceTests = []struct {
	s, old, new string
	n           int
	out         string
}{

	{"a aa aaa a aa", "aa", "c-d", -1, "a c-d c-da a c-d"},
	{"", "", "", -1, ""},
	{"", "", "b", -1, "b"},
	{"a", "", "b", -1, "bab"},
	{"aaa", "", "b", -1, "bababab"},
	{" a aa ", "", "b", -1, "b bab babab b"},
	{"", "a", "a", -1, ""},
	{"", "a", "b", -1, ""},
	{"a", "a", "a", -1, "a"},
	{"a", "a", "b", -1, "b"},
	{"aa", "a", "b", -1, "bb"},
	{"aa", "a", "bc", -1, "bcbc"},
	{"aa", "aa", "b", -1, "b"},
	{"a aa aa ", "aa", "bc", -1, "a bc bc "},
	{"a aa aa cde", "aa", "bc", -1, "a bc bc cde"},
	{" a aa aa", "a", "aab", -1, " aab aabaab aabaab"},
	{"", "", "", 0, ""},
	{"", "", "b", 0, ""},
	{"a", "", "b", 0, "a"},
	{"aaa", "", "b", 0, "aaa"},
	{" a aa ", "", "b", 0, " a aa "},
	{"", "a", "a", 0, ""},
	{"", "a", "b", 0, ""},
	{"a", "a", "a", 0, "a"},
	{"a", "a", "b", 0, "a"},
	{"aa", "a", "b", 0, "aa"},
	{"aa", "a", "bc", 0, "aa"},
	{"aa", "aa", "b", 0, "aa"},
	{"aa", "aa", "b", 0, "aa"},
	{"aa", "aa", "b", 0, "aa"},
	{"a aa aa ", "aa", "bc", 0, "a aa aa "},
	{" a aa aa", "a", "aab", 0, " a aa aa"},
	{"", "", "", 1, ""},
	{"", "", "b", 1, "b"},
	{"a", "", "b", 1, "ba"},
	{"aaa", "", "b", 1, "baaa"},
	{" a aa ", "", "b", 1, "b a aa "},
	{"", "a", "a", 1, ""},
	{"", "a", "b", 1, ""},
	{"a", "a", "a", 1, "a"},
	{"a", "a", "b", 1, "b"},
	{"aa", "a", "b", 1, "ba"},
	{"aa", "a", "bc", 1, "bca"},
	{"aa", "aa", "b", 1, "b"},
	{"aa", "aa", "b", 1, "b"},
	{"aa", "aa", "b", 1, "b"},
	{"a aa aa ", "aa", "bc", 1, "a bc aa "},
	{" a aa aa", "a", "aab", 1, " aab aa aa"},
	{"", "", "", 2, ""},
	{"", "", "b", 2, "b"},
	{"a", "", "b", 2, "bab"},
	{"aaa", "", "b", 2, "babaa"},
	{" a aa ", "", "b", 2, "b ba aa "},
	{"", "a", "a", 2, ""},
	{"", "a", "b", 2, ""},
	{"a", "a", "a", 2, "a"},
	{"a", "a", "b", 2, "b"},
	{"aa", "a", "b", 2, "bb"},
	{"aa", "a", "bc", 2, "bcbc"},
	{"aa", "aa", "b", 2, "b"},
	{"aa", "aa", "b", 2, "b"},
	{"aa", "aa", "b", 2, "b"},
	{"a aa aa ", "aa", "bc", 2, "a bc bc "},
	{" a aa aa", "a", "aab", 2, " aab aaba aa"},
}

func TestLibReplace(t *testing.T) {
	for _, test := range replaceTests {
		actual := Replace(test.s, test.old, test.new, test.n)
		libout := strings.Replace(test.s, test.old, test.new, test.n)
		if actual != libout {
			t.Errorf("Replace(%q, %q, %q, %d) = %q ; %q wanted", test.s, test.old, test.new, test.n, actual, libout)
		}
	}
}

func TestReplace(t *testing.T) {
	for _, test := range replaceTests {
		actual := Replace(test.s, test.old, test.new, test.n)
		if actual != test.out {
			t.Errorf("Replace(%q, %q, %q, %d) = %q ; %q wanted", test.s, test.old, test.new, test.n, actual, test.out)
		}
	}
}

func BenchmarkReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Replace("a aa aaa a aa", "aa", "c-d", -1)

	}
}
