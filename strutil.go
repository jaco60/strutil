// Package strutil is a collection of functions on strings
package strutil

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

// Length returns the number of codepoints in a string
func Length(in string) int {
	return utf8.RuneCountInString(in)
}

// Reverse returns a reversed version of a string
func Reverse(in string) string {
	out := []rune(in)
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return string(out)
}

// Title returns a version of a string where each word is capitalised
func Title(in string) string {
	re := regexp.MustCompile("[^\\p{Z}]+")
	words := re.FindAllString(in, -1)
	out := []string{}
	for _, word := range words {
		out = append(out, strings.Title(word))
	}
	return strings.Join(out, " ")
}
