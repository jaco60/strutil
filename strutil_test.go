package strutil

import (
	"testing"
)

func TestReverse(t *testing.T) {
	dataset := []string{
		"0123456789",
		"Il était une fois",
		"あいうえお",
		"天地玄黃宇宙洪荒",
	}

	expected := []string{
		"9876543210",
		"siof enu tiaté lI",
		"おえういあ",
		"荒洪宙宇黃玄地天",
	}

	for i, word := range dataset {
		got, want := Reverse(word), expected[i]
		if got != want {
			t.Errorf("got : %s, want : %s\n", got, want)
		}
	}
}
