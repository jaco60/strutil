// strutil is a collection of functions on strings
package strutil

// Reverse returns a reversed version of its argument
func Reverse(in string) string {
	out := []rune(in)
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return string(out)
}
