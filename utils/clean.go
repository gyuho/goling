package utils // import "github.com/gyuho/goling/utils"

import (
	"regexp"
	"strings"
)

// CleanUp cleans up unnecessary characters in string.
// It cleans up the blank characters that carry no meaning in context,
// converts all white spaces into single whitespace.
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

// Expand expands the apostrophe phrases.
func Expand(str string) string {
	// str = strings.Replace(strings.ToLower(str), "'d", " would", -1)
	str = strings.Replace(str, "'d", " would", -1)
	str = strings.Replace(str, "isn't", "is not", -1)
	str = strings.Replace(str, "aren't", "are not", -1)
	str = strings.Replace(str, "was't", "was not", -1)
	str = strings.Replace(str, "weren't", "were not", -1)

	str = strings.Replace(str, "'ve", " have", -1)
	str = strings.Replace(str, "'re", " are", -1)
	str = strings.Replace(str, "'m", " am", -1)
	str = strings.Replace(str, "it's", "it is", -1)
	str = strings.Replace(str, "what's", "what is", -1)
	str = strings.Replace(str, "'ll", " will", -1)

	str = strings.Replace(str, "won't", "will not", -1)
	str = strings.Replace(str, "can't", "can not", -1)
	str = strings.Replace(str, "mustn't", "must not", -1)

	str = strings.Replace(str, "haven't", "have not", -1)
	str = strings.Replace(str, "hasn't", "has not", -1)

	str = strings.Replace(str, "dn't", "d not", -1)
	str = strings.Replace(str, "don't", "do not", -1)
	str = strings.Replace(str, "doesn't", "does not", -1)
	str = strings.Replace(str, "didn't", "did not", -1)

	return str
}

// ReplaceNonAlpha replaces all non-alphabetic characters.
func ReplaceNonAlpha(str, replace string) string {
	// alphabetic (== [A-Za-z])
	// \s is a white space character
	validID := regexp.MustCompile(`[^[:alpha:]\s]`)
	return validID.ReplaceAllString(str, replace)
}

// ReplaceNonAlnum replaces all alphanumeric characters.
func ReplaceNonAlnum(str, replace string) string {
	// alphanumeric (== [0-9A-Za-z])
	// \s is a white space character
	validID := regexp.MustCompile(`[^[:alnum:]\s]`)
	return validID.ReplaceAllString(str, replace)
}

// SplitToWords splits string into a slice of words.
func SplitToWords(str string) []string {
	str = Expand(str)
	validID := regexp.MustCompile(`\"`)
	str = validID.ReplaceAllString(str, "")
	return strings.Split(CleanUp(str), " ")
}

// SplitToSentences splits a string into sentences, removing all puctuation.
// It returns an slice of sentences.
func SplitToSentences(str string) []string {
	validID1 := regexp.MustCompile(`\"`)
	str = validID1.ReplaceAllString(str, "")

	validID2 := regexp.MustCompile(`[.]`)
	str = validID2.ReplaceAllString(str, ".___")

	validID3 := regexp.MustCompile(`[?]`)
	str = validID3.ReplaceAllString(str, "?___")

	validID4 := regexp.MustCompile(`[!]`)
	str = validID4.ReplaceAllString(str, "!___")

	slice := strings.Split(str, "___")
	// to clean up the empty strings
	result := []string{}
	for _, value := range slice {
		value = CleanUp(value)
		if value != "" {
			result = append(result, value)
		}
	}
	return result
}

// CountCharacter returns the total number of single characters.
func CountCharacter(str string) int {
	str = CleanUp(str)
	slice := strings.Split(str, "")
	return len(slice)
}

// CountWord returns the total number of words.
func CountWord(str string) int {
	slice := SplitToWords(str)
	return len(slice)
}

// CountSentence return the total number of sentences.
func CountSentence(str string) int {
	slice := SplitToSentences(str)
	return len(slice)
}
