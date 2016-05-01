package challenges_tests

import (
	"bytes"
	"testing"
)

// RemoveLetters removes letters from string
func RemoveLetters(s string, exclude string) string {

	m := map[rune]bool{}

	for _, v := range exclude {
		m[v] = true
	}

	var b bytes.Buffer
	for _, v := range s {
		if !m[v] {
			b.WriteString(string(v))
		}
	}

	return b.String()
}

func TestRemoveLetters(t *testing.T) {

	removeLettersTests := []struct {
		s       string
		exclude string
		out     string
	}{
		{"abcde", "", "abcde"},
		{"abcde", "abcde", ""},
		{"abcdef", "bcdf", "ae"},
		{"A_8!b.e?kd", "!_?", "A8b.ekd"},
	}

	for _, test := range removeLettersTests {
		actual := RemoveLetters(test.s, test.exclude)
		if actual != test.out {
			t.Errorf("RemoveLetters(%q,%q)= %v; want %v", test.s, test.exclude, actual, test.out)
		}
	}

}

func BenchmarkRemoveLetters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RemoveLetters("abcdefghijklmnopqrstuvwxyz", "aeiou")
	}
}
