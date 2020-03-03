package strutil

import (
	"testing"
)

func TestLength(t *testing.T) {
	dataset := map[string]int{
		"":         0,
		"bonjour":  7,
		"œ":        1,
		"Il était": 8,
	}
	for word, want := range dataset {
		if got := Length(word); got != want {
			t.Errorf("got : %d, want : %d\n", got, want)
		}
	}
}

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

func TestTitle(t *testing.T) {
	dataset := map[string]string{
		"once upon   a time":    "Once Upon A Time",
		"Il était une fois":     "Il Était Une Fois",
		"il a dit : viens là !": "Il A Dit : Viens Là !",
	}
	for phrase, want := range dataset {
		if got := Title(phrase); got != want {
			t.Errorf("got : %s, want : %s\n", got, want)
		}
	}
}
