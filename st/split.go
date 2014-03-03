package st

import (
	"regexp"
	"strings"
)

// SplitWord splits a string into words.
// It is punctionation-inclusive.
// It returns an slice of words.
func SplitWord(str string) []string {
	str = ExpandApostrophe(str)
	validID := regexp.MustCompile(`\"`)
	str = validID.ReplaceAllString(str, "")

	slice := strings.Split(str, " ")

	var result []string
	for _, value := range slice {
		value = CleanUp(value)
		if value != "" {
			result = append(result, value)
		}
	}
	return result
}

// GetWords extracts only words, removing all puctuation.
// It returns an slice of words.
func GetWords(str string) []string {
	str = ExpandApostrophe(str)
	str = ReplaceNonAlnumWithSpace(str)
	return strings.Fields(str)
}

// SplitSentRaw splits a string into sentences.
// It is punctionation-inclusive.
// It returns an slice of sentences.
func SplitSentRaw(str string) []string {
	validID_1 := regexp.MustCompile(`\"`)
	str = validID_1.ReplaceAllString(str, "")

	validID_2 := regexp.MustCompile(`[.]`)
	str = validID_2.ReplaceAllString(str, ".___")

	validID_3 := regexp.MustCompile(`[?]`)
	str = validID_3.ReplaceAllString(str, "?___")

	validID_4 := regexp.MustCompile(`[!]`)
	str = validID_4.ReplaceAllString(str, "!___")

	slice := strings.Split(str, "___")

	var result []string
	for _, value := range slice {
		value = CleanUp(value)
		if value != "" {
			result = append(result, value)
		}
	}
	return result
}

// SplitSent splits a string into sentences, removing all puctuation.
// It returns an slice of sentences.
func SplitSent(str string) []string {
	validID_1 := regexp.MustCompile(`[.!?]`)
	str = validID_1.ReplaceAllString(str, "___")

	validID_2 := regexp.MustCompile(`\"`)
	str = validID_2.ReplaceAllString(str, "")

	slice := strings.Split(str, "___")

	// to clean up the empty strings
	var result []string
	for _, value := range slice {
		value = CleanUp(value)
		if value != "" {
			result = append(result, value)
		}
	}
	return result
}

// SplitParag splits texts into paragraphs.
// It returns an slice of paragraphs.
func SplitParag(str string) []string {
	slice := strings.Split(str, "\n")

	// to clean up the empty strings
	var result []string
	for _, value := range slice {
		value = CleanUp(value)
		if value != "" {
			result = append(result, value)
		}
	}
	return result
}
