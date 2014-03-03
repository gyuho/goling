// Package st provides basic string manipulating functions.
package st

import "strings"

// CleanUp cleans up unnecessary characters in string.
// It cleans up the blank characters that carry no meaning in context
// , converts all whitespaces into single whitespace.
// String is immutable, which means the original string would not change.
func CleanUp(str string) string {

	// validID := regexp.MustCompile(`\s{2,}`)
	// func TrimSpace(s string) string
	// slicing off all "leading" and
	// "trailing" white space, as defined by Unicode.
	str = strings.TrimSpace(str)

	// func Fields(s string) []string
	// Fields splits the slice s around each instance
	// of "one or more consecutive white space"
	slice := strings.Fields(str)

	// now join them with a single white space character
	return strings.Join(slice, " ")
}
