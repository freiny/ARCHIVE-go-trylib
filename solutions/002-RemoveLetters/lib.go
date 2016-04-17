package lib

import "bytes"

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
