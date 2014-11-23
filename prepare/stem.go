package st

import "strings"

// Stem returns the stem, in English.
// It receives only one word.
func Stem(str string) string {
	str = ReplaceNonAlnumWithSpace(strings.ToLower(str))

	// start with the longest suffix
	switch {

	case strings.HasSuffix(str, "ization"):
		str = str[:len(str)-7]

	case strings.HasSuffix(str, "mming") || strings.HasSuffix(str, "izer") || strings.HasSuffix(str, "fulness") || strings.HasSuffix(str, "iveness") || strings.HasSuffix(str, "ousness") || strings.HasSuffix(str, "aliti") || strings.HasSuffix(str, "iviti"):
		str = str[:len(str)-4]

	case strings.HasSuffix(str, "ing") || strings.HasSuffix(str, "biliti") || strings.HasSuffix(str, "ational") || strings.HasSuffix(str, "ies") || strings.HasSuffix(str, "alli") || strings.HasSuffix(str, "ation") || strings.HasSuffix(str, "alism"):
		str = str[:len(str)-3]

	case strings.HasSuffix(str, "sses") || strings.HasSuffix(str, "ies") || strings.HasSuffix(str, "ator"):
		str = str[:len(str)-2]
		// delete the last two characters

	case strings.HasSuffix(str, "ss"):
		str = str

	case !strings.HasSuffix(str, "is") && strings.HasSuffix(str, "s"):
		str = str[:len(str)-1]

	case !strings.HasSuffix(str, "need") && strings.HasSuffix(str, "ed"):
		str = str[:len(str)-2]

	default:
		// No stem found!
	}

	// convert it back to string format
	return str
}

// StemText gets stem from text.
func StemText(str string) string {
	nstr := ReplaceNonAlnumWithSpace(strings.ToLower(str))
	nslice := GetWords(nstr)
	var result []string
	for _, elem := range nslice {
		result = append(result, Stem(elem))
	}
	return strings.Join(result, " ")
}
