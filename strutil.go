// strutil is a collection of functions on strings
package strutil

import (
	"regexp"
	"strings"
)

// Reverse returns a reversed version of its argument
func Reverse(in string) string {
	out := []rune(in)
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return string(out)
}

// Title returns a version of its argument where each word is capitalised
func Title(in string) string {
	re := regexp.MustCompile("[^\\p{Z}]+")
	words := re.FindAllString(in, -1)
	out := []string{}
	for _, word := range words {
		out = append(out, strings.Title(word))
	}
	return strings.Join(out, " ")
}
