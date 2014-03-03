// Package simi calculates string similarities.
package simi

import (
	"math"

	"github.com/gyuho/goling/st"
)

// StringSimilarity returns the similarity of two texts
// using cosine similarity and Levenshtein distance.
// CosineSimilarity: the higher, the more similar(range 0 to 1).
// Levenshtein: the lower, the more similar(any integer number).
func StringSimilarity(str1, str2 string) float64 {

	// First remove unnecessary characters: punctuation, affix
	nstr1 := st.Stem(str1)
	nstr2 := st.Stem(str2)

	// CosineSimilarity: higher result means higher similarity.
	// range from 0 to 1, so scale them up to 1000
	measure1 := CosineSimilarity(nstr1, nstr2)

	// LevenshteinSentence: lower result means higher similarity.
	measure2 := (1.0 / float64(LevenshteinDistance(nstr1, nstr2))) * 1000

	str1nogap := st.DeleteSpace(nstr1)
	str2nogap := st.DeleteSpace(nstr2)
	measure3 := (1.0 / float64(LevenshteinDistance(str1nogap, str2nogap))) * 1000

	return math.Sqrt(measure1*measure2*measure3) * 100
}
