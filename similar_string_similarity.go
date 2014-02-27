package goling

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

// StringSimilarity returns the similarity of two texts
// using cosine similarity and Levenshtein distance.
// CosineSimilarity: the higher, the more similar(range 0 to 1).
// Levenshtein: the lower, the more similar(any integer number).
func StringSimilarity(str1, str2 string) float64 {

	// First remove unnecessary characters
	// 	punctuation, affix
	nstr1 := SuffixEnglish(str1)
	nstr2 := SuffixEnglish(str2)

	// CosineSimilarity: higher result means higher similarity.
	// range from 0 to 1, so scale them up to 1000
	measure1 := CosineSimilarity(nstr1, nstr2)

	// LevenshteinSentence: lower result means higher similarity.
	measure2 := (1.0 / float64(LevenshteinDistance(nstr1, nstr2))) * 1000

	str1nogap := DeleteSpace(nstr1)
	str2nogap := DeleteSpace(nstr2)
	measure3 := (1.0 / float64(LevenshteinDistance(str1nogap, str2nogap))) * 1000

	return math.Sqrt(measure1*measure2*measure3) * 100
}

// SuffixEnglish returns the stem, in English.
// It receives only one word.
func SuffixEnglish(str string) string {

	ns := ReplaceNonAlnumWithSpace(strings.ToLower(str))

	// convert string to []byte
	tm := []byte(ns)

	var result []byte

	// start with the longest suffix
	switch {

	case bytes.HasSuffix(tm, []byte("ational")) || bytes.HasSuffix(tm, []byte("ization")):
		result = tm[:len(tm)-5]

	case bytes.HasSuffix(tm, []byte("mming")) || bytes.HasSuffix(tm, []byte("izer")) || bytes.HasSuffix(tm, []byte("fulness")) || bytes.HasSuffix(tm, []byte("iveness")) || bytes.HasSuffix(tm, []byte("ousness")) || bytes.HasSuffix(tm, []byte("aliti")) || bytes.HasSuffix(tm, []byte("iviti")):
		result = tm[:len(tm)-4]

	case bytes.HasSuffix(tm, []byte("ing")) || bytes.HasSuffix(tm, []byte("biliti")) || bytes.HasSuffix(tm, []byte("ies")) || bytes.HasSuffix(tm, []byte("alli")) || bytes.HasSuffix(tm, []byte("ation")) || bytes.HasSuffix(tm, []byte("alism")):
		result = tm[:len(tm)-3]

	case bytes.HasSuffix(tm, []byte("sses")) || bytes.HasSuffix(tm, []byte("ies")) || bytes.HasSuffix(tm, []byte("ator")):
		result = tm[:len(tm)-2]
		// delete the last two characters

	case bytes.HasSuffix(tm, []byte("ss")):
		result = tm

	case !bytes.HasSuffix(tm, []byte("is")) && bytes.HasSuffix(tm, []byte("s")):
		result = tm[:len(tm)-1]

	case !bytes.HasSuffix(tm, []byte("need")) && bytes.HasSuffix(tm, []byte("ed")):
		result = tm[:len(tm)-2]

	default:
		result = tm
		// No stem found!
	}

	// convert it back to string format
	return fmt.Sprintf("%s", result)
}
