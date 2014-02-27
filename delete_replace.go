package goling

import "regexp"

// ReplaceNonAlphaWithSpace removes all non-alphabetic characters.
// It replaces them with a single whitespace character.
// It returns the new version of input string, in lower case.
func ReplaceNonAlphaWithSpace(str string) string {
	str = ExpandApostrophe(str)
	// alphabetic (== [A-Za-z])
	// \s is a white space character
	validID := regexp.MustCompile(`[^[:alpha:]\s]`)
	return validID.ReplaceAllString(str, " ")
}

// DeleteNonAlpha removes all non-alphabetic characters.
// It returns the new version of input string, in lower case.
func DeleteNonAlpha(str string) string {
	str = ExpandApostrophe(str)
	// alphabetic (== [A-Za-z])
	// \s is a white space character
	validID := regexp.MustCompile(`[^[:alpha:]\s]`)
	return validID.ReplaceAllString(str, "")
}

// ReplaceNonAlnumWithSpace removes all alphanumeric characters.
// It replaces them with a single whitespace character.
// It returns the new version of input string, in lower case.
func ReplaceNonAlnumWithSpace(str string) string {
	str = ExpandApostrophe(str)
	// alphanumeric (== [0-9A-Za-z])
	// \s is a white space character
	validID := regexp.MustCompile(`[^[:alnum:]\s]`)
	return validID.ReplaceAllString(str, " ")
}

// DeleteNonAlnum removes all alphanumeric characters.
// It returns the new version of input string, in lower case.
func DeleteNonAlnum(str string) string {
	str = ExpandApostrophe(str)
	// alphanumeric (== [0-9A-Za-z])
	// \s is a white space character
	validID := regexp.MustCompile(`[^[:alnum:]\s]`)
	return validID.ReplaceAllString(str, "")
}
