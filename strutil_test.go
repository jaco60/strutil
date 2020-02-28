package strutil

import (
	"testing"
)

func TestReverse(t *testing.T) {
	dataset := map[string]string{
		"0123456789":        "9876543210",
		"Il était une fois": "siof enu tiaté lI",
		"あいうえお":             "おえういあ",
		"天地玄黃宇宙洪荒":          "荒洪宙宇黃玄地天",
	}

	for word, want := range dataset {
		if got := Reverse(word); got != want {
			t.Errorf("got : %s, want : %s\n", got, want)
		}
	}
}
