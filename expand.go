package goling

import "strings"

// ExpandApostrophe expands the apostrophe phrases.
// And convert them to lower case letters.
func ExpandApostrophe(str string) string {
	// assignment between string is not "copy"
	// even if str1 is longer than str2
	// like str1 := "Hello", str2 = ""
	// str1 = str2 makes str1 ""
	str = strings.Replace(strings.ToLower(str), "'d", " would", -1)

	// If n < 0, there is no limit on the number of replacements.
	str = strings.Replace(str, "isn't", "is not", -1)
	str = strings.Replace(str, "aren't", "are not", -1)
	str = strings.Replace(str, "was't", "was not", -1)
	str = strings.Replace(str, "weren't", "were not", -1)

	str = strings.Replace(str, "'ve", " have", -1)
	str = strings.Replace(str, "'re", " are", -1)
	str = strings.Replace(str, "'m", " am", -1)
	str = strings.Replace(str, "t's", "t is", -1)
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
