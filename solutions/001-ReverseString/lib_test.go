package lib

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	m := map[string]string{
		"asdf":        "fdsa",
		"vU _!12 ?3P": "P3? 21!_ Uv",
		"":            "",
	}
	for in, out := range m {
		got := Reverse(in)
		if got != out {
			t.Error("fn(", in, ") =>", got, "!=", out)
		} else {
			fmt.Println("PASS: fn(", in, ") =>", got, "==", out)
		}
	}

}
