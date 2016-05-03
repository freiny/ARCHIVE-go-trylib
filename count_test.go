package fStrings

import (
	"strings"
	"testing"
)

func Count(s, sep string) int {
	if s == "" && sep == "" {
		return 1
	}
	if sep == "" {
		return len(s) + 1
	}
	count := 0
	sepLen := len(sep)
	sp := s[0:len(s)]
	for {
		switch {
		case len(sp) == sepLen:
			if sp == sep {
				count++
			}
			return count
		case len(sp) < sepLen:
			return count
		case len(sp) > sepLen:
			if sp[0:sepLen] == sep {
				count++
				sp = sp[sepLen:len(sp)]
			} else {
				sp = sp[1:len(sp)]
			}
		}
	}
}

var countTests = []struct {
	s, sep string
	out    int
}{
	{"", "", 1},
	{"a", "", 2},
	{"ab", "", 3},
	{"", "a", 0},
	{"a", "a", 1},
	{"a", "b", 0},
	{"a", "aaa", 0},
	{"abcde", "a", 1},
	{"abcde", "c", 1},
	{"abcde", "e", 1},
	{"abcde", "ab", 1},
	{"abcde", "bc", 1},
	{"abcde", "de", 1},
	{"aaa aaaa aaa", "aa", 4},
	{"aaa aaaa aaa", "aaaaa", 0},
}

func TestLibCount(t *testing.T) {
	for _, test := range countTests {
		actual := Count(test.s, test.sep)
		libOut := strings.Count(test.s, test.sep)
		if actual != libOut {
			t.Errorf("Count(%q, %q) = %d ; %d wanted", test.s, test.sep, actual, libOut)
		}
	}
}

func TestCount(t *testing.T) {
	for _, test := range countTests {
		actual := Count(test.s, test.sep)
		if actual != test.out {
			t.Errorf("Count(%q, %q) = %d ; %d wanted", test.s, test.sep, actual, test.out)
		}
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("aaba aaabab aaaa", "aa")
	}
}
